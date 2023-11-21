package core

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"time"

	"golang.org/x/exp/slices"
)

const (
	AddUrlPaymentVersion string = "v2"
	AddUrlVtidVersion    string = "v1"
	AddUrlPayment        string = `paynow`
	AddUrlVtid           string = "paynowid"
	PaynowidServiceType  string = "account,charge,recurring,cardinfo"
	ServiceCommandSearch string = "Search"
	ServiceTypeSearch    string = "search"
	StructNameSearch     string = "Request"
	SearchServer         string = "search"
	DummyServer          string = "test"
	Version              string = "0.1.1"
	TxnVersion           string = "1.0.0"
)

func Execute(requestDto any, responseDto any, merchantConfig *MerchantConfig, ctx context.Context) {
	var tv = reflect.TypeOf(requestDto)
	if tv.Kind() != reflect.Ptr {
		errRes := `{"mstatus":"failure","vResultCode":"MA99","merrMsg":"requestDto Type Kind Error"}`
		_ = json.Unmarshal([]byte(errRes), &responseDto)
		slog.ErrorContext(ctx, "requestDto Type Kind is not reflect.Ptr")
		return
	}

	sendUrl := CreateSendUrl(requestDto, merchantConfig.Host, merchantConfig.DummyRequest)
	if sendUrl == "" {
		errRes := `{"mstatus":"failure","vResultCode":"MA99","merrMsg":"System internal error"}`
		_ = json.Unmarshal([]byte(errRes), &responseDto)
		slog.ErrorContext(ctx, "Failed to create the request destination URL.")
		return
	}

	slog.InfoContext(ctx, "MDK Execute start", "sendUrl", sendUrl)

	AppendConnectParam(requestDto, merchantConfig.MerchantCcId, merchantConfig.DummyRequest, TxnVersion)

	j, marshalErr := json.Marshal(requestDto)
	if marshalErr != nil {
		errRes := `{"mstatus":"failure","vResultCode":"MA99","merrMsg":"System internal error"}`
		_ = json.Unmarshal([]byte(errRes), &responseDto)
		slog.ErrorContext(ctx, "An error occurred when marshaling the requestDto.", "error", marshalErr.Error())
		return
	}

	if merchantConfig.ParameterMaskingEnable {
		maskJson, _ := MaskJson(j)
		slog.InfoContext(ctx, "MDK Request Body", "requestBody", string(maskJson))
	} else {
		slog.InfoContext(ctx, "MDK Request Body", "requestBody", string(j))
	}

	requestBody := AppendAuthHash(merchantConfig.MerchantCcId, string(j), merchantConfig.MerchantSecretKey)

	req, reqErr := http.NewRequest("POST", sendUrl, strings.NewReader(url.QueryEscape(requestBody)))
	if reqErr != nil {
		errRes := `{"mstatus":"failure","vResultCode":"MF02","merrMsg":"Gateway Server Connect Error"}`
		_ = json.Unmarshal([]byte(errRes), &responseDto)
		slog.ErrorContext(ctx, "An error occurred when initializing the http.Request.", "error", reqErr.Error())
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "VeryTransMDK/"+Version+" ("+runtime.Version()+"/"+runtime.GOARCH+")")

	response, sendError := SendRequest(req, merchantConfig.Timeout)

	if sendError != nil {
		var urlErr *url.Error
		errors.As(sendError, &urlErr)
		if urlErr.Timeout() {
			errRes := `{"mstatus":"failure","vResultCode":"MF03","merrMsg":"Connection to server timed out"}`
			_ = json.Unmarshal([]byte(errRes), &responseDto)
			slog.WarnContext(ctx, "Request has timed out.", "error", urlErr.Error())
		} else {
			errRes := `{"mstatus":"failure","vResultCode":"MF03","merrMsg":"Sends an HTTP request error"}`
			_ = json.Unmarshal([]byte(errRes), &responseDto)
			slog.WarnContext(ctx, "An error occurred in the HTTP request.", "error", urlErr.Error())
		}
		return
	}

	switch response.StatusCode {
	case 200:
	case 500:
		errRes := `{"mstatus":"failure","vResultCode":"MF05","merrMsg":"500 Internal Server Error"}`
		_ = json.Unmarshal([]byte(errRes), &responseDto)
		slog.WarnContext(ctx, "The payment gateway responded with an HTTP 500 Internal Server Error.")
		return
	case 502:
		errRes := `{"mstatus":"failure","vResultCode":"MF06","merrMsg":"502 Bad Gateway"}`
		_ = json.Unmarshal([]byte(errRes), &responseDto)
		slog.WarnContext(ctx, "The payment gateway responded with an HTTP 502 Bad Gateway.")
		return
	case 503:
		errRes := `{"mstatus":"failure","vResultCode":"MF07","merrMsg":"503 Service Unavailable"}`
		_ = json.Unmarshal([]byte(errRes), &responseDto)
		slog.WarnContext(ctx, "The payment gateway responded with an HTTP 503 Service Unavailable.")
		return
	default:
		errRes := `{"mstatus":"failure","vResultCode":"MF99","merrMsg":"Type F Internal Error"}`
		_ = json.Unmarshal([]byte(errRes), &responseDto)
		slog.WarnContext(ctx, "The payment gateway responded with an HTTP "+response.Status)
		return
	}

	resByte, readErr := ReadResponse(response.Body)
	if readErr != nil {
		errRes := `{"mstatus":"failure","vResultCode":"MF03","merrMsg":"Read Response Error"}`
		_ = json.Unmarshal([]byte(errRes), &responseDto)
		slog.ErrorContext(ctx, "An error occurred while reading the response body.", "error", readErr.Error())
		return
	}

	var resultJson string
	if merchantConfig.ParameterMaskingEnable {
		maskJson, _ := MaskJson(resByte)
		resultJson = string(maskJson)
	} else {
		resultJson = string(resByte)
	}
	slog.InfoContext(ctx, "MDK Response body received", "responseBody", resultJson, "responseStatus", response.StatusCode)

	setResponseErr := SetResponseProperties(resByte, &responseDto, resultJson)
	if setResponseErr != nil {
		errRes := `{"mstatus":"failure","vResultCode":"MB99","merrMsg":"Type B Internal Error"}`
		_ = json.Unmarshal([]byte(errRes), &responseDto)
		slog.ErrorContext(ctx, "An error occurred when unmarshaling the response body.", "error", setResponseErr.Error())
		return
	}

}

func AppendConnectParam(dto any, merchantCcId string, dummyRequest bool, txnVersion string) {
	rv := reflect.ValueOf(dto).Elem()
	for i := 0; i < rv.NumField(); i++ {
		f := rv.Type().Field(i)
		switch f.Name {
		case "TxnVersion":
			fv := rv.FieldByName(f.Name)
			if fv.CanSet() {
				fv.SetString(txnVersion)
			}
		case "DummyRequest":
			fv := rv.FieldByName(f.Name)
			if fv.CanSet() {
				if dummyRequest {
					fv.SetString("1")
				} else {
					fv.SetString("0")
				}
			}
		case "MerchantCcid":
			fv := rv.FieldByName(f.Name)
			if fv.CanSet() {
				fv.SetString(merchantCcId)
			}
		}
	}
}

func CreateSendUrl(dto any, host string, isDummy bool) string {
	var tv = reflect.TypeOf(dto).Elem()
	pkgPath := tv.PkgPath()
	name := tv.Name()

	arr := strings.Split(pkgPath, "/")
	serviceType := strings.ToLower(arr[len(arr)-1])
	if serviceType == ServiceTypeSearch && name == StructNameSearch {
		name = ServiceCommandSearch + StructNameSearch
	}
	regex, _ := regexp.Compile("Request\\z")
	serviceCommand := regex.ReplaceAllString(name, "")
	serviceName := AddUrlPayment
	version := AddUrlPaymentVersion
	if slices.Contains(strings.Split(PaynowidServiceType, ","), serviceType) {
		serviceName = AddUrlVtid
		version = AddUrlVtidVersion
	}
	if serviceName == AddUrlPayment && serviceCommand == ServiceCommandSearch {
		serviceName += "-" + SearchServer
	}
	if isDummy {
		serviceName = DummyServer + "-" + serviceName
	}

	return host + "/" + serviceName + "/" + version + "/" + serviceCommand + "/" + serviceType
}

func AppendAuthHash(merchantCcId string, body string, merchantSecret string) string {
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(merchantCcId+body+merchantSecret)))
	return "{\"params\":" + body + ",\"authHash\":\"" + hash + "\"}"
}

func SendRequest(req *http.Request, timeout time.Duration) (resp *http.Response, err error) {
	client := &http.Client{
		Transport: &http.Transport{},
		Timeout:   timeout,
	}
	return client.Do(req)
}

func ReadResponse(body io.ReadCloser) ([]byte, error) {
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			return
		}
	}(body)
	return io.ReadAll(body)
}

func SetResponseProperties(b []byte, target any, resultJson string) error {
	var obj map[string]interface{}
	if err := json.Unmarshal(b, &obj); err != nil {
		return err
	}
	if val, ok := obj["result"]; ok {
		r, ok := val.(map[string]interface{})
		if ok {
			result := r           // resultを取り出す
			delete(obj, "result") // objからresultを削除

			// objからrootレベルのキーを取得(payNowIdResponse, viewResultなど)
			var rootKeys []string
			for s := range obj {
				rootKeys = append(rootKeys, s)
			}

			for key := range result {
				if p, ok := result[key]; ok {
					// rootレベルにキー重複がなければresultのプロパティをobjに追加
					if !slices.Contains(rootKeys, key) {
						obj[key] = p
					}
				}
			}
		}
	}

	obj["ResultJson"] = resultJson

	v, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(v, &target); err != nil {
		return err
	}
	return nil
}
