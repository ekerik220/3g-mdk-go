package core

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/ekerik220/3g-mdk-go/dto/Mpi"
	"github.com/ekerik220/3g-mdk-go/dto/PayNowId"
)

func TestGetMaskedValue(t *testing.T) {

	type param struct {
		key      string
		value    string
		expected string
	}

	var params = []param{
		{"test", "test", "test"},
		{"cardNumber", "2342-3423-4234-2341", "234234******2341"},
		{"cardNumber", "2342342342342341", "234234******2341"},
		{"cardNumber", "378282246310005", "378282*****0005"},
		{"cardNumber", "3782-822463-10005", "378282*****0005"},
		{"cardNumber", "36666666666660", "366666****6660"},
		{"cardNumber", "3666-666666-6660", "366666****6660"},
		{"cardNumber", "1234567890", "**********"},
		{"cardNumber", "12345678901", "123456*8901"},
		{"CardNumber", "4111111111111111", "411111******1111"},
		{"cardnumber", "4111111111111111", "411111******1111"},
		{"card_number", "4111111111111111", "4111111111111111"},
		{"mailaddr", "foo@example.com", "***@example.com"},
		{"mailaddr", "f@example.com", "*@example.com"},
		{"mailaddr", "@example.com", "@example.com"},
		{"mailaddr", "@example@.com", "@example@.com"},
		{"mailaddr", "foo.example.com", "foo.example.com"},
		{"mailaddr", "foo@t.co", "***@t.co"},
		{"mailaddr", "foo@c", "***@c"},
		{"mailaddr", "1@c", "*@c"},
		{"mailaddr", "@", "@"},
		{"mailaddr", "1@", "*@"},
		{"mailAddress", "foo@example.com", "***@example.com"},
		{"Mailaddr", "foo@example.com", "***@example.com"},
		{"mail_addr", "foo@example.com", "foo@example.com"},
		{"billingAddressCity", "[渋谷区(Shibuya)]", "**************"},
		{"cardholderName", "TEST TARO", "*********"},
	}

	for p := range params {
		actual := GetMaskedValue(params[p].key, params[p].value)
		if actual != params[p].expected {
			t.Error(params[p], actual)
		}
	}

}

func TestMaskJson(t *testing.T) {

	mpiRequestDto := &Mpi.AuthorizeRequest{
		TxnVersion:        "1.0.0",
		DummyRequest:      "0",
		MerchantCcid:      "A100000000000000000000cc",
		OrderId:           "dummy-order-id",
		ServiceOptionType: "mpi-complete",
		Amount:            "100",
		DeviceChannel:     "02",
		RedirectionUri:    "http://localhost/index.php",
		FirstKanaName:     "タロウ",
		PayNowIdParam: &PayNowId.Param{
			Token: "abcdef01-2345-6789-0abc-def012345678",
			AccountParam: &PayNowId.AccountParam{
				AccountId: "dummy-account-id",
			},
			Memo1: "メモ",
		},
	}

	j1, err := json.Marshal(mpiRequestDto)
	if err != nil {
		t.Fail()
	}

	j2, err := MaskJson(j1)
	if err != nil {
		t.Fail()
	}

	maskedMpiRequestDto := Mpi.AuthorizeRequest{}
	err = json.Unmarshal(j2, &maskedMpiRequestDto)

	if err != nil {
		if maskedMpiRequestDto.MerchantCcid != "A100000000000000000000cc" {
			t.Fail()
		}
		if maskedMpiRequestDto.DummyRequest != "0" {
			t.Fail()
		}
		if maskedMpiRequestDto.TxnVersion != "1.0.0" {
			t.Fail()
		}
		if maskedMpiRequestDto.OrderId != mpiRequestDto.OrderId {
			t.Fail()
		}
		if maskedMpiRequestDto.ServiceOptionType != mpiRequestDto.ServiceOptionType {
			t.Fail()
		}
		if maskedMpiRequestDto.Amount != mpiRequestDto.Amount {
			t.Fail()
		}
		if maskedMpiRequestDto.DeviceChannel != mpiRequestDto.DeviceChannel {
			t.Fail()
		}
		if maskedMpiRequestDto.RedirectionUri != mpiRequestDto.RedirectionUri {
			t.Fail()
		}
		if maskedMpiRequestDto.FirstKanaName != "***" {
			t.Fail()
		}
		if maskedMpiRequestDto.PayNowIdParam.Token != mpiRequestDto.PayNowIdParam.Token {
			t.Fail()
		}
		if maskedMpiRequestDto.PayNowIdParam.AccountParam.AccountId != mpiRequestDto.PayNowIdParam.AccountParam.AccountId {
			t.Fail()
		}
		if maskedMpiRequestDto.PayNowIdParam.Memo1 != mpiRequestDto.PayNowIdParam.Memo1 {
			t.Fail()
		}
	}

}

func TestMaskJson2(t *testing.T) {

	r := `{"payNowIdResponse":{"account":{"accountId":"dummy-account-id","cardInfo":[{"cardExpire":"12/50","cardId":"ABCDEFGHIJKLMN0123456789Z","cardNumber":"411111********11","cardholderName":"TEST TARO","defaultCard":"1"}]},"message":"処理が成功しました。","processId":"1234567890","status":"success"},"result":{"vResultCode":"G001X00100000000","custTxn":"87455438","reqAmount":"100","reqCardExpire":"*****","reqCardNumber":"411111********11","reqHttpAccept":"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7","reqHttpUserAgent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36 Edg/116.0.1938.62","reqRedirectionUri":"http://localhost:8125/","reqSecurityCode":"0000","reqWithCapture":"false","res3dMessageVersion":"2.1.0","resBrandId":"4","resCorporationId":"05","marchTxn":"87455438","merrMsg":"本人認証実行可能です。","mstatus":"success","optionResults":[],"orderId":"dummy-order-id","serviceType":"mpi","txnVersion":"1.0.0"}}`
	j, err := MaskJson([]byte(r))

	if err != nil {
		t.Fail()
	}

	if strings.Contains(string(j), "TEST TARO") {
		t.Fail()
	}

}
