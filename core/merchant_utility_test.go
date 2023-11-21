package core

import "testing"

func TestCheckMessage(t *testing.T) {

	secretKey := "abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789"
	contentHmacHeader := "h=HmacSHA256;s=A100000000000001000000cc;v=96afe3ca78dd828e1b4643dacae031422ca252bdb2a109178d4637bd44a7ee99"
	contentHmacHeaderFail := "h=HmacSHA256;s=A100000000000001000000cc;v=96afe3ca78dd828e1b4643dacae031422ca252bdb2a109178d4637bd44a7ee90"
	messageBody := "numberOfNotify=1&pushTime=20230609163134&pushId=100&orderId0000=1657593096&vResultCode0000=G011A00100000000&txnType0000=AuthorizeConfirm&mpiMstatus0000=success&cardMstatus0000=success&dummy0000=1"

	if !CheckMessage(secretKey, messageBody, contentHmacHeader) {
		t.Fail()
	}

	if CheckMessage("", "", "") {
		t.Fail()
	}

	if CheckMessage(secretKey, messageBody, contentHmacHeaderFail) {
		t.Fail()
	}

}
