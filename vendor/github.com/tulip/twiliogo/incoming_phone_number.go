package twiliogo

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type Capabilites struct {
	Voice bool `json:"voice"`
	SMS   bool `json:"SMS"`
	MMS   bool `json:"MMS"`
}

type IncomingPhoneNumber struct {
	Sid                  string      `json:"sid"`
	AccountSid           string      `json:"account_sid"`
	FriendlyName         string      `json:"friendly_name"`
	PhoneNumber          string      `json:"phone_number"`
	VoiceUrl             string      `json:"voice_url"`
	VoiceMethod          string      `json:"voice_method"`
	VoiceFallbackUrl     string      `json:"voice_fallback_url"`
	VoiceFallbackMethod  string      `json:"voice_fallback_method"`
	StatusCallback       string      `json:"status_callback"`
	StatusCallbackMethod string      `json:"status_callback_method"`
	VoiceCallerIdLookup  bool        `json:"voice_caller_id_lookup"`
	VoiceApplicationId   string      `json:"voice_application_id"`
	DateCreated          string      `json:"date_created"`
	DateUpdated          string      `json:"date_updated"`
	SmsUrl               string      `json:"sms_url"`
	SmsMethod            string      `json:"sms_method"`
	SmsFallbackUrl       string      `json:"sms_fallback_url"`
	SmsFallbackMethod    string      `json:"sms_fallback_method"`
	SmsApplicationId     string      `json:"sms_application_id"`
	Capabilities         Capabilites `json:"capabilities"`
	ApiVersion           string      `json:"api_version"`
	Uri                  string      `json:"uri"`
}

func GetIncomingPhoneNumber(client Client, sid string) (*IncomingPhoneNumber, error) {
	var incomingPhoneNumber *IncomingPhoneNumber

	res, err := client.get(url.Values{}, "/IncomingPhoneNumbers/"+sid+".json")

	if err != nil {
		return nil, err
	}

	incomingPhoneNumber = new(IncomingPhoneNumber)
	err = json.Unmarshal(res, incomingPhoneNumber)

	return incomingPhoneNumber, err
}

func BuyPhoneNumber(client Client, number Optional) (*IncomingPhoneNumber, error) {
	var incomingPhoneNumber *IncomingPhoneNumber

	if number == nil {
		return nil, Error{"Must input PhoneNumber or AreaCode"}
	}

	params := url.Values{}
	param, value := number.GetParam()
	params.Set(param, value)

	res, err := client.post(params, "/IncomingPhoneNumbers.json")

	if err != nil {
		return incomingPhoneNumber, err
	}

	incomingPhoneNumber = new(IncomingPhoneNumber)
	err = json.Unmarshal(res, incomingPhoneNumber)

	return incomingPhoneNumber, err
}

// UpdateSetPhoneNumberFields updates any non-empty fields (other than sid) in the passed struct
//  because pretty much every field can be updated, we assume we are handed a
//  struct with only the fields that we want to update and the sid nonempty.
//
//  Because we can't set the bool VoiceCallerIdLookup to false or true, we must
//  pass it explicitly as a pointer to true or false, or else leave it nil to leave it alone
func UpdatePhoneNumberFields(client Client, number *IncomingPhoneNumber, voiceCallerIdLookup *bool) (*IncomingPhoneNumber, error) {
	if number == nil || len(number.Sid) == 0 {
		return nil, fmt.Errorf("You must pass a struct with sid set")
	}

	changes := url.Values{}
	if len(number.AccountSid) > 0 {
		changes.Add("AccountSid", number.AccountSid)
	}
	if len(number.FriendlyName) > 0 {
		changes.Add("FriendlyName", number.FriendlyName)
	}
	if len(number.ApiVersion) > 0 {
		changes.Add("ApiVersion", number.ApiVersion)
	}
	if len(number.VoiceUrl) > 0 {
		changes.Add("VoiceUrl", number.VoiceUrl)
	}
	if len(number.VoiceMethod) > 0 {
		changes.Add("VoiceMethod", number.VoiceMethod)
	}
	if len(number.VoiceFallbackUrl) > 0 {
		changes.Add("VoiceFallbackUrl", number.VoiceFallbackUrl)
	}
	if len(number.VoiceFallbackMethod) > 0 {
		changes.Add("VoiceFallbackMethod", number.VoiceFallbackMethod)
	}
	if len(number.StatusCallback) > 0 {
		changes.Add("StatusCallback", number.StatusCallback)
	}
	if len(number.StatusCallbackMethod) > 0 {
		changes.Add("StatusCallbackMethod", number.StatusCallbackMethod)
	}
	if voiceCallerIdLookup != nil {
		changes.Add("VoiceCallerIdLookup", strconv.FormatBool(*voiceCallerIdLookup))
	}
	if len(number.SmsUrl) > 0 {
		changes.Add("SmsUrl", number.SmsUrl)
	}
	if len(number.SmsMethod) > 0 {
		changes.Add("SmsMethod", number.SmsMethod)
	}
	if len(number.SmsFallbackUrl) > 0 {
		changes.Add("SmsFallbackUrl", number.SmsFallbackUrl)
	}
	if len(number.SmsFallbackMethod) > 0 {
		changes.Add("SmsFallbackMethod", number.SmsFallbackMethod)
	}

	res, err := client.post(changes, "/IncomingPhoneNumbers/"+number.Sid+".json")
	if err != nil {
		return nil, err
	}

	var incomingPhoneNumber *IncomingPhoneNumber
	incomingPhoneNumber = new(IncomingPhoneNumber)
	err = json.Unmarshal(res, incomingPhoneNumber)

	return incomingPhoneNumber, err
}

func ReleasePhoneNumber(client Client, sid string) error {
	err := client.delete("/IncomingPhoneNumbers/" + sid)
	return err
}
