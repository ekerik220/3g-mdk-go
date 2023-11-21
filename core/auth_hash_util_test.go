package core

import (
	"net/url"
	"testing"
)

func TestCheckAuthHash(t *testing.T) {

	merchantCcId := "A100000000000001000000cc"
	merchantPw := "abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789"

	if checkAuthHash(nil, merchantCcId, merchantPw) {
		t.Fail()
	}

	if checkAuthHash(url.Values{}, merchantCcId, merchantPw) {
		t.Fail()
	}

	if checkAuthHash(url.Values{"authParams": {"dummy"}}, merchantCcId, merchantPw) {
		t.Fail()
	}

	if checkAuthHash(url.Values{"vAuthInfo": {"dummy"}}, merchantCcId, merchantPw) {
		t.Fail()
	}

	if checkAuthHash(url.Values{"authParams": {"dummy"}, "vAuthInfo": {"dummy"}}, merchantCcId, merchantPw) {
		t.Fail()
	}

	values := url.Values{
		"RequestId":           {"YqAkr6wfBVEAAAS@SK0AAAAF"},
		"OrderId":             {"order-1654662296"},
		"reqAmount":           {"100"},
		"reqCardNumber":       {"411111********11"},
		"reqCurrencyUnit":     {""},
		"mpiMstatus":          {"success"},
		"vResultCode":         {"G011A00100000000"},
		"cardMstatus":         {"success"},
		"cardTransactionType": {"a"},
		"centerRequestDate":   {"20220608132519"},
		"centerResponseDate":  {"20220608132519"},
		"connectedCenterId":   {"jcn"},
		"acquirerCode":        {"05"},
		"authCode":            {"000000"},
		"tradUrl":             {""},
		"fdResult":            {""},
		"authParams":          {"ZmRSZXN1bHQsYXV0aENvZGUsUmVxdWVzdElkLG1waU1zdGF0dXMsY2VudGVyUmVxdWVzdERhdGUsY2VudGVyUmVzcG9uc2VEYXRlLE9yZGVySWQscmVxQ2FyZE51bWJlcixjYXJkVHJhbnNhY3Rpb25UeXBlLGNhcmRNc3RhdHVzLHRyYWRVcmwscmVxQ3VycmVuY3lVbml0LGNvbm5lY3RlZENlbnRlcklkLHJlcUFtb3VudCxhY3F1aXJlckNvZGUsdlJlc3VsdENvZGU="},
		"vAuthInfo":           {"27faee59e696f68f9da0a485da1e9a7c827e9308b6a73828366a528e7158938c"},
	}

	if !checkAuthHash(values, merchantCcId, merchantPw) {
		t.Fail()
	}

	if checkAuthHash(values, merchantCcId, "dummy") {
		t.Fail()
	}

}
