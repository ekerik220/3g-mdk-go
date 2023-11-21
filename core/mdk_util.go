package core

import (
	"encoding/json"
	"golang.org/x/exp/slices"
	"reflect"
	"strings"
	"unicode/utf8"
)

func MaskJson(b []byte) ([]byte, error) {
	if len(b) > 0 {
		if b[0] != 91 {
			var obj map[string]interface{}
			if err := json.Unmarshal(b, &obj); err != nil {
				return nil, err
			}
			parseMap(obj)
			v, mErr := json.Marshal(obj)
			if mErr != nil {
				return nil, mErr
			}
			return v, nil
		} else {
			var obj []interface{}
			if err := json.Unmarshal(b, &obj); err != nil {
				return nil, err
			}
			parseSlice(obj)
			v, mErr := json.Marshal(obj)
			if mErr != nil {
				return nil, mErr
			}
			return v, nil
		}
	}
	return b, nil
}

func parseMap(obj map[string]interface{}) {
	for key := range obj {
		tv := reflect.TypeOf(obj[key])
		if tv.Kind() == reflect.Map {
			if mv, ok := obj[key].(map[string]interface{}); ok {
				parseMap(mv)
			}
		} else if tv.Kind() == reflect.Slice {
			if mv, ok := obj[key].([]interface{}); ok {
				parseSlice(mv)
			}
		} else if tv.Kind() == reflect.String {
			if v, ok := obj[key].(string); ok {
				obj[key] = GetMaskedValue(key, v)
			}
		}
	}
}

func parseSlice(obj []interface{}) {
	for i := range obj {
		tv := reflect.TypeOf(obj[i])
		if tv.Kind() == reflect.Map {
			if mv, ok := obj[i].(map[string]interface{}); ok {
				parseMap(mv)
			}
		} else if tv.Kind() == reflect.Slice {
			if mv, ok := obj[i].([]interface{}); ok {
				parseSlice(mv)
			}
		}
	}
}

func GetMaskedValue(key string, value string) string {
	if value == "" {
		return ""
	}

	if strings.Contains(strings.ToLower(key), "mailaddr") {
		index := strings.Index(value, "@")
		if index > 0 {
			return strings.Repeat("*", index) + value[index:]
		} else {
			return value
		}
	} else if strings.Contains(strings.ToLower(key), "cardnumber") {
		cardNumber := strings.ReplaceAll(value, "-", "")
		if len(cardNumber) >= 11 {
			return cardNumber[0:6] + strings.Repeat("*", len(cardNumber)-10) + cardNumber[len(cardNumber)-4:]
		} else {
			return strings.Repeat("*", len(cardNumber))
		}
	} else if isNeedMask(key) {
		return strings.Repeat("*", utf8.RuneCountInString(value))
	} else {
		return value
	}
}

func isNeedMask(name string) bool {
	if name == "" {
		return false
	}
	return slices.Contains(maskItems(), name)
}

func maskItems() []string {

	return []string{"cardNumber", "cardExpire", "birthday", "tel", "firstKanaName",
		"lastKanaName", "mailAddr", "merchantMailAddr", "cancelMailAddr", "name1", "name2", "kana", "kana1", "kana2",
		"telNo", "address1", "address2", "address3", "post1", "post2", "customerNo", "pan", "settleAmount", "exchangeRate",
		"paymentDate", "paymentStatus", "centerTxnId", "shipName", "shipStreet1", "shipStreet2", "shipCity", "shipState",
		"shipCountry", "shipPostalCode", "shipPhone", "reqFirstKanaName", "reqLastKanaName", "reqTel", "reqBirthday",
		"reqCardNumber", "reqCardExpire", "securityCode", "pin", "jis1SecondTrack", "jis2Track", "mailAddress", "firstName",
		"lastName", "recipient", "address", "zip", "emvData", "company", "emailAddress", "phoneNumber", "streetLine",
		"streetLine2", "city", "postal", "birthDate", "mothersMaidenName", "cardHolderName", "hashedCardNumber", "expireDate",
		"comment", "hiraganaId", "email", "mobileEmail", "fullName", "companyName", "departmentName", "mobile", "zipCode",
		"fullKanaName", "customerName", "customerZip", "customerAddress1", "customerAddress2", "customerAddress3", "department",
		"detailName", "shippingName", "shippingPhone", "shippingAddress1", "shippingAddress2", "shippingAddress3", "shippingPrefecture",
		"shippingPostalCode", "buyerId", "buyerEmail", "auId", "identityCode", "billingToken", "oneTimeKey", "authorizationCode",
		"oneTimeCode", "qrCode", "taxiNumber", "taxiDriverName", "taxiDriverMobile", "cardholderName", "cardholderEmail",
		"cardholderHomePhoneCountry", "cardholderHomePhoneNumber", "cardholderMobilePhoneCountry", "cardholderMobilePhoneNumber",
		"cardholderWorkPhoneCountry", "cardholderWorkPhoneNumber", "billingAddressCity", "billingAddressCountry", "billingAddressLine1",
		"billingAddressLine2", "billingAddressLine3", "billingPostalCode", "billingAddressState", "shippingAddressCity",
		"shippingAddressCountry", "shippingAddressLine1", "shippingAddressLine2", "shippingAddressLine3", "shippingPostalCode",
		"shippingAddressState", "customerIp", "deliveryEmailAddress", "holder", "houseNumber1", "street1", "street2", "city",
		"state", "country", "postcode", "address", "fax", "phone", "postboxNumber", "postCode", "street", "suite", "merchantCustomerId",
		"givenName", "middleName", "surname", "workPhone", "ip", "identificationDocId", "recipientAddress", "recipientApartment",
		"recipientCity", "recipientCountry", "recipientEmail", "recipientFirstName", "recipientLastName", "recipientMiddleName",
		"recipientPhone", "recipientPostcode", "recipientSalutation", "recipientState", "recipientStreet", "dob", "frequentFlyerNumber",
		"name", "number"}

}
