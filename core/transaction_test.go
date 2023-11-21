package core

import (
	"encoding/json"
	"testing"

	"github.com/ekerik220/3g-mdk-go/dto/Account"
	"github.com/ekerik220/3g-mdk-go/dto/Charge"
	"github.com/ekerik220/3g-mdk-go/dto/Mpi"
	"github.com/ekerik220/3g-mdk-go/dto/PayNowId"
	"github.com/ekerik220/3g-mdk-go/dto/ScoreAtPay"
	"github.com/ekerik220/3g-mdk-go/dto/Search"
)

func TestCreateSendUrl(t *testing.T) {

	host := "https://api.veritrans.co.jp"

	mpiAuthorizeRequest := &Mpi.AuthorizeRequest{}
	if CreateSendUrl(mpiAuthorizeRequest, host, true) != "https://api.veritrans.co.jp/test-paynow/v2/Authorize/mpi" {
		t.Fail()
	}
	if CreateSendUrl(mpiAuthorizeRequest, host, false) != "https://api.veritrans.co.jp/paynow/v2/Authorize/mpi" {
		t.Fail()
	}

	accountAddRequest := &Account.AddRequest{}
	if CreateSendUrl(accountAddRequest, host, true) != "https://api.veritrans.co.jp/test-paynowid/v1/Add/account" {
		t.Fail()
	}
	if CreateSendUrl(accountAddRequest, host, false) != "https://api.veritrans.co.jp/paynowid/v1/Add/account" {
		t.Fail()
	}

	searchRequest := &Search.Request{}
	if CreateSendUrl(searchRequest, host, true) != "https://api.veritrans.co.jp/test-paynow-search/v2/Search/search" {
		t.Fail()
	}
	if CreateSendUrl(searchRequest, host, false) != "https://api.veritrans.co.jp/paynow-search/v2/Search/search" {
		t.Fail()
	}

	chargeRequest := &Charge.AddRequest{}
	if CreateSendUrl(chargeRequest, host, false) != "https://api.veritrans.co.jp/paynowid/v1/Add/charge" {
		t.Fail()
	}

	scoreAtPayRequest := &ScoreAtPay.AuthorizeRequest{}
	if CreateSendUrl(scoreAtPayRequest, host, false) != "https://api.veritrans.co.jp/paynow/v2/Authorize/scoreatpay" {
		t.Fail()
	}
}

func TestAppendConnectParam(t *testing.T) {
	mpiAuthorizeRequest := &Mpi.AuthorizeRequest{}
	AppendConnectParam(mpiAuthorizeRequest, "A100000000000001000000cc", false, "1.0.0")
	if mpiAuthorizeRequest.MerchantCcid != "A100000000000001000000cc" {
		t.Fail()
	}
	if mpiAuthorizeRequest.DummyRequest != "0" {
		t.Fail()
	}
	if mpiAuthorizeRequest.TxnVersion != "1.0.0" {
		t.Fail()
	}
}

func TestMarshal(t *testing.T) {
	merchantCcId := "A100000000000000000000cc"

	mpi := &Mpi.AuthorizeRequest{
		TxnVersion:        "1.0.0",
		DummyRequest:      "0",
		MerchantCcid:      merchantCcId,
		OrderId:           "dummy-order-id",
		ServiceOptionType: "mpi-complete",
		Amount:            "100",
		DeviceChannel:     "02",
		PayNowIdParam: &PayNowId.Param{
			Token: "abcdef01-2345-6789-0abc-def012345678",
			AccountParam: &PayNowId.AccountParam{
				AccountId: "dummy-account-id",
			},
		},
	}

	j, err := json.Marshal(mpi)
	if err != nil {
		t.Fail()
	}
	if string(j) != "{\"txnVersion\":\"1.0.0\",\"dummyRequest\":\"0\",\"merchantCcid\":\"A100000000000000000000cc\",\"payNowIdParam\":{\"accountParam\":{\"accountId\":\"dummy-account-id\"},\"token\":\"abcdef01-2345-6789-0abc-def012345678\"},\"serviceOptionType\":\"mpi-complete\",\"orderId\":\"dummy-order-id\",\"amount\":\"100\",\"deviceChannel\":\"02\"}" {
		t.Fail()
	}
}

func TestAppendAuthHash(t *testing.T) {
	merchantCcId := "A100000000000000000000cc"
	merchantSecret := "abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789"
	j := "{\"txnVersion\":\"1.0.0\",\"dummyRequest\":\"0\",\"merchantCcid\":\"A100000000000000000000cc\",\"payNowIdParam\":{\"accountParam\":{\"accountId\":\"dummy-account-id\"},\"token\":\"abcdef01-2345-6789-0abc-def012345678\"},\"serviceOptionType\":\"mpi-complete\",\"orderId\":\"dummy-order-id\",\"amount\":\"100\",\"deviceChannel\":\"02\"}"

	s := AppendAuthHash(merchantCcId, j, merchantSecret)

	if s != "{\"params\":{\"txnVersion\":\"1.0.0\",\"dummyRequest\":\"0\",\"merchantCcid\":\"A100000000000000000000cc\",\"payNowIdParam\":{\"accountParam\":{\"accountId\":\"dummy-account-id\"},\"token\":\"abcdef01-2345-6789-0abc-def012345678\"},\"serviceOptionType\":\"mpi-complete\",\"orderId\":\"dummy-order-id\",\"amount\":\"100\",\"deviceChannel\":\"02\"},\"authHash\":\"2b110237583f7717ccde4b641e3b31b6364e8885b4967cdb34af61e77d3c82f9\"}" {
		t.Fail()
	}
}

func TestSetResponseProperties(t *testing.T) {
	j := `{"payNowIdResponse":{"account":{"accountId":"account-dummy1660701218995","cardInfo":[{"cardExpire":"12/25","cardId":"2SAH2KRI7IJF2KYCGLVDRVIW5","cardNumber":"411111******1111","cardholderName":"cardholderName","defaultCard":"1"}]},"message":"処理が成功しました。","processId":"3179705046","status":"success"},"result":{"vResultCode":"G001X00100000000","custTxn":"10434958","authRequestDatetime":"Wed Aug 30 16:16:48 JST 2023","authResponseDatetime":"Wed Aug 30 16:16:48 JST 2023","authStartUrl":"https://api.veritrans.co.jp:443/tercerog/webinterface/GWTripartiteNACommandRcv/mpi/GetAuthorizeResult?md=QTEwMDAwMDAwMDAwMDAwMTA2OTk5NWNj-MjAyMzA4MzAwNzE2NDg0NjY*-YWVmMjNmYjFjNDcyOWYwZmFiYzZkNjBmMWZlNDZjMjEwNGQ4NWY2ZDNmYTI4ODJhZTgwZDc2YTkwM2YwMWE4Ng**==ZO7s4LKTIsj8BjRn6DQaEQAAANg","kindCode":"5","reqAmount":"10","reqCardExpire":"*****","reqCardNumber":"411111******1111","reqHttpAccept":"text/*;q=0.3, text/html;q=0.7","reqHttpUserAgent":"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)","res3dMessageVersion":"2.1.0","resBrandId":"4","resCorporationId":"05","resResponseContents":"<!DOCTYPE html>\n<HTML lang=\"ja\">\n<HEAD>\n<META charset=\"utf-8\">\n<STYLE>\n<!--\n.box{\n background: white;\n height: 100%;\n width: 100%;\n text-align: center;\n}\n.spin{\n position: absolute;\n top: 30%;\n left: 0;\n right: 0;\n margin: auto;\n display:none;\n}\n.inner{\n position: absolute;\n left: 50%;\n top: 50%;\n}\n.content{\n transform: translate(-50%,-50%);\n -webkit-transform: translate(-50%,-50%);/*Safari*/\n}\n-->\n</STYLE>\n<SCRIPT>\nvar timerId;\nvar send = false;\nvar submited = false;\nfunction OnLoadEvent(){\n document.getElementById(\"spin\").style.display = \"block\";\n document.getElementById(\"threeDsMethodForm\").submit();\n setBrowserInfo();\n setTimeout(function() {\n  send = true;\n }, 2000);\n timerId = setTimeout(function() {\n  document.getElementById(\"result\").value = \"N\";\n  if (submited == false) {\n   submited = true;\n   document.getElementById(\"downloadForm\").submit();\n  }\n }, 1*10000);\n}\nwindow.addEventListener(\"message\", function(e){\n switch (e.data.action) {\n  case \"endThreeDsMethod\":\n   endThreeDsMethod();\n   break;\n }\n});\nfunction endThreeDsMethod(){\n clearTimeout(timerId);\n if (submited == false) {\n  submited = true;\n  var intervalId = setInterval(function() {\n   if (send == true) {\n    clearInterval(intervalId);\n    document.getElementById(\"downloadForm\").submit();\n   }\n  }, 100);\n }\n}\nfunction setBrowserInfo(){\n document.getElementById( \"javaEnabled\" ).value = window.navigator.javaEnabled();\n document.getElementById( \"javascriptEnabled\" ).value = \"true\";\n document.getElementById( \"language\" ).value = (window.navigator.languages && window.navigator.languages[0])\n   || window.navigator.language || window.navigator.userLanguage || window.navigator.browserLanguage;\n document.getElementById( \"colorDepth\" ).value = screen.colorDepth;\n document.getElementById( \"width\" ).value = screen.width;\n document.getElementById( \"height\" ).value = screen.height;\n document.getElementById( \"timezone\" ).value =  new Date().getTimezoneOffset();\n}\n</SCRIPT>\n</HEAD>\n<BODY onLoad=\"OnLoadEvent();\">\n<DIV id=\"box\" class=\"box\">\n<NOSCRIPT>\n<H4>JavaScriptが無効な場合、ご利用いただけません。</H4>\n</NOSCRIPT>\n<IMG src=\"https://api.veritrans.co.jp:443/tercerog/img/mpi/spin.gif\" id=\"spin\" class=\"spin\">\n<DIV id=\"inner\" class=\"inner\">\n<DIV id=\"content\" class=\"content\">\n<IMG src=\"https://api.veritrans.co.jp:443/tercerog/img/mpi/logo_visa.gif?ID=ZO7s4LKTIsj8BjRn6DQaEQAAANg\">\n</DIV>\n</DIV>\n<FORM id=\"downloadForm\" name=\"downloadForm\" action=\"https://api.veritrans.co.jp:443/tercerog/webinterface/GWTripartiteNACommandRcv/mpi/AuthorizeConfirm?i=71b0ecf6-d1a4-3a46-9b01-1b2be4b7b210&amp;h=ca32716d9c697e7e243de0373e185f3c7d85d670caf86b2a74fad9bf0379b7ef\" method=\"POST\">\n<INPUT type=\"hidden\" name=\"browserJavaEnabled\" id=\"javaEnabled\" >\n<INPUT type=\"hidden\" name=\"browserJavascriptEnabled\" id=\"javascriptEnabled\" value=\"false\">\n<INPUT type=\"hidden\" name=\"browserLanguage\" id=\"language\" >\n<INPUT type=\"hidden\" name=\"browserColorDepth\" id=\"colorDepth\" >\n<INPUT type=\"hidden\" name=\"browserScreenWidth\" id=\"width\" >\n<INPUT type=\"hidden\" name=\"browserScreenHeight\" id=\"height\" >\n<INPUT type=\"hidden\" name=\"browserTZ\" id=\"timezone\" >\n<INPUT type=\"hidden\" name=\"result\" id=\"result\" value=\"Y\" >\n</FORM>\n<FORM id=\"threeDsMethodForm\" action=\"https://api.veritrans.co.jp:443/tercerog/webinterface/GWTripartiteNACommandRcv/mpi/AuthorizeNotify?md=QTEwMDAwMDAwMDAwMDAwMTA2OTk5NWNj-MjAyMzA4MzAwNzE2NDg0NjY*-YWVmMjNmYjFjNDcyOWYwZmFiYzZkNjBmMWZlNDZjMjEwNGQ4NWY2ZDNmYTI4ODJhZTgwZDc2YTkwM2YwMWE4Ng**==ZO7s4LKTIsj8BjRn6DQaEQAAANg\" method=\"POST\" target=\"threeDsMethodFrame\">\n<INPUT type=\"hidden\" id=\"threeDSMethodData\" name=\"threeDSMethodData\" value=\"dummyMethodData\">\n</FORM>\n<IFRAME id=\"threeDsMethodFrame\" name=\"threeDsMethodFrame\" style=\"visibility:hidden;border:none;\" width=\"10\" height=\"10\" sandbox=\"allow-forms allow-same-origin allow-scripts\"></iframe>\n</DIV>\n</BODY>\n</HTML>\n","marchTxn":"10434985","merrMsg":"本人認証実行可能です。","mstatus":"success","optionResults":[],"orderId":"dummy1660701218995","serviceType":"mpi","txnVersion":"1.2.8"}}`
	//sj := `{"result":{"vResultCode":"N001000000000000","orderInfos":{"orderInfo":[{"accountId":"account-dummy163305062","index":0,"lastSuccessTxnType":"Cancel","orderId":"dummy163305062","orderStatus":"end","properOrderInfo":{"cardExpire":"12/25","startTxn":"10435077"},"serviceTypeCd":"card","successDetailTxnType":"race","transactionInfos":{"transactionInfo":[{"amount":"10","command":"Authorize","mstatus":"success","properTransactionInfo":{"cardTransactionType":"ac","centerRequestDate":"20230830163306","centerResponseDate":"20230830163306","connectedCenterId":"jcn","fdResult":"accept","gatewayRequestDate":"20230830163306","gatewayResponseDate":"20230830163306","loopback":"0","pending":"0","reqAcquirerCode":"05","reqAmount":"10","reqCardExpire":"*****","reqCardNumber":"411111******1111","reqItemCode":"0990","reqJpoInformation":"10","reqSecurityCode":"0000","reqWithCapture":"true","resActionCode":"000","resAuthCode":"000000","resCenterErrorCode":"   ","resReturnReferenceNumber":"012345678901","txnKind":"card"},"txnDatetime":"2023-08-30 16:33:06.0","txnId":"10435077","vResultCode":"A001X00100000000"},{"command":"Capture","mstatus":"failure","properTransactionInfo":{"cardTransactionType":"ac","gatewayRequestDate":"20230830163405","gatewayResponseDate":"20230830163405","txnKind":"card"},"txnDatetime":"2023-08-30 16:34:05.383","txnId":"10435089","vResultCode":"NH02000000000000"},{"amount":"5","command":"Cancel","mstatus":"success","properTransactionInfo":{"cardTransactionType":"race","centerRequestDate":"20230830163440","centerResponseDate":"20230830163440","connectedCenterId":"jcn","gatewayRequestDate":"20230830163440","gatewayResponseDate":"20230830163442","loopback":"0","pending":"0","reqAcquirerCode":"05","reqAmount":"5","reqItemCode":"0990","resActionCode":"000","resAuthCode":"000000","resCenterErrorCode":"   ","resReturnReferenceNumber":"012345678901","txnKind":"card"},"txnDatetime":"2023-08-30 16:34:42.092","txnId":"10435091","vResultCode":"A001000000000000"}]}}]},"overMaxCountFlag":false,"searchCount":1,"mstatus":"success","serviceType":"search"}}`

	var resDto Mpi.AuthorizeResponse

	if err := SetResponseProperties([]byte(j), &resDto, "{}"); err != nil {
		t.Fail()
	}

	if resDto.Mstatus != "success" {
		t.Fail()
	}

	if resDto.PayNowIdResponse.Account.AccountId != "account-dummy1660701218995" {
		t.Fail()
	}

}
