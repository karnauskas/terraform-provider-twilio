package twilio

import (
	"github.com/davegardnerisme/phonegeocode"
	"github.com/hashicorp/terraform/helper/schema"
	twilioc "github.com/tulip/twiliogo"
)

func resourcePhonenumberImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	client := m.(*twilioMeta).Client

	numberStruct, err := twilioc.GetIncomingPhoneNumber(client, d.Id())
	if err != nil {
		return nil, err
	}

	cc, err := phonegeocode.New().Country(numberStruct.PhoneNumber)
	if err != nil {
		return nil, err
	}
	d.Set("iso_country_code", cc)

	d.Set("api_version", numberStruct.ApiVersion)
	d.Set("name", numberStruct.FriendlyName)
	d.Set("phone_number", numberStruct.PhoneNumber)
	d.Set("sms_fallback_method", numberStruct.SmsFallbackMethod)
	d.Set("sms_fallback_url", numberStruct.SmsFallbackUrl)
	d.Set("sms_method", numberStruct.SmsMethod)
	d.Set("sms_url", numberStruct.SmsUrl)
	d.Set("status_callback", numberStruct.StatusCallback)
	d.Set("status_callback_method", numberStruct.StatusCallbackMethod)
	d.Set("voice_caller_id_lookup", numberStruct.VoiceCallerIdLookup)
	d.Set("voice_fallback_method", numberStruct.VoiceFallbackMethod)
	d.Set("voice_fallback_url", numberStruct.VoiceFallbackUrl)
	d.Set("voice_method", numberStruct.VoiceMethod)
	d.Set("voice_url", numberStruct.VoiceUrl)

	return []*schema.ResourceData{d}, nil
}
