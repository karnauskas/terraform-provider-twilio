package twiliogo

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type AvaliablePhoneNumberWrapper struct {
	URI  string          `json:"uri"`
	Data json.RawMessage `json:"available_phone_numbers"`
}

type AvaliablePhoneNumber struct {
	FriendlyName        string      `json:"friendly_name"`
	PhoneNumber         string      `json:"phone_number"`
	IsoCountry          string      `json:"iso_country"`
	Capabilities        Capabilites `json:"capabilities"`
	AddressRequirements string      `json:"api_version"`
}

type TollFreeAvaliablePhoneNumber AvaliablePhoneNumber
type MobileAvaliablePhoneNumber AvaliablePhoneNumber

type LocalAvaliablePhoneNumber struct {
	*AvaliablePhoneNumber
	Beta bool `json:"beta"`
}

type USACanadaLocalAvaliablePhoneNumber struct {
	*LocalAvaliablePhoneNumber
	Lata       string `json:"lata"`
	RateCenter string `json:"rate_center"`
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
	Region     string `json:"region"`
	PostalCode string `json:"postal_code"`
}

func GetUSACanadaLocalAvailablePhoneNumbers(client Client, isoCountryCode string, optionals ...Optional) (*[]USACanadaLocalAvaliablePhoneNumber, error) {
	var avaliablePhoneNumbers *[]USACanadaLocalAvaliablePhoneNumber
	avaliablePhoneNumbers = new([]USACanadaLocalAvaliablePhoneNumber)
	if isoCountryCode != "US" && isoCountryCode != "CA" {
		return nil, fmt.Errorf("For this function, isoCountryCode must be 'US' or 'CA', for other countries use GetLocalAvailablePhoneNumbers")
	}

	params := url.Values{}
	for _, optional := range optionals {
		param, value := optional.GetParam()
		params.Set(param, value)
	}

	res, err := client.get(params, "/AvailablePhoneNumbers/"+isoCountryCode+"/Local.json")
	if err != nil {
		return nil, err
	}

	var wrapper AvaliablePhoneNumberWrapper
	if err = json.Unmarshal(res, &wrapper); err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(wrapper.Data), avaliablePhoneNumbers)
	return avaliablePhoneNumbers, err
}

func GetLocalAvailablePhoneNumbers(client Client, isoCountryCode string, optionals ...Optional) (*[]LocalAvaliablePhoneNumber, error) {
	var avaliablePhoneNumbers *[]LocalAvaliablePhoneNumber
	avaliablePhoneNumbers = new([]LocalAvaliablePhoneNumber)

	params := url.Values{}
	for _, optional := range optionals {
		param, value := optional.GetParam()
		params.Set(param, value)
	}

	res, err := client.get(params, "/AvailablePhoneNumbers/"+isoCountryCode+"/Local.json")
	if err != nil {
		return nil, err
	}

	var wrapper AvaliablePhoneNumberWrapper
	if err = json.Unmarshal(res, &wrapper); err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(wrapper.Data), avaliablePhoneNumbers)
	return avaliablePhoneNumbers, err
}

func GetTollFreeAvaliablePhoneNumber(client Client, isoCountryCode string, optionals ...Optional) (*[]TollFreeAvaliablePhoneNumber, error) {
	var avaliablePhoneNumbers *[]TollFreeAvaliablePhoneNumber
	avaliablePhoneNumbers = new([]TollFreeAvaliablePhoneNumber)

	params := url.Values{}
	for _, optional := range optionals {
		param, value := optional.GetParam()
		params.Set(param, value)
	}

	res, err := client.get(params, "/AvailablePhoneNumbers/"+isoCountryCode+"/TollFree.json")
	if err != nil {
		return nil, err
	}

	var wrapper AvaliablePhoneNumberWrapper
	if err = json.Unmarshal(res, &wrapper); err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(wrapper.Data), avaliablePhoneNumbers)
	return avaliablePhoneNumbers, err
}

func GetMobileAvaliablePhoneNumber(client Client, isoCountryCode string, optionals ...Optional) (*[]MobileAvaliablePhoneNumber, error) {
	var avaliablePhoneNumbers *[]MobileAvaliablePhoneNumber
	avaliablePhoneNumbers = new([]MobileAvaliablePhoneNumber)

	params := url.Values{}
	for _, optional := range optionals {
		param, value := optional.GetParam()
		params.Set(param, value)
	}

	res, err := client.get(params, "/AvailablePhoneNumbers/"+isoCountryCode+"/Mobile.json")
	if err != nil {
		return nil, err
	}

	var wrapper AvaliablePhoneNumberWrapper
	if err = json.Unmarshal(res, &wrapper); err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(wrapper.Data), avaliablePhoneNumbers)
	return avaliablePhoneNumbers, err
}

// type AvailablePhoneNumberTypes struct {
// 	Local    bool
// 	TollFree bool
// 	Mobile   bool
// }
//
// type AvaliablePhoneNumberCountryList struct {
// 	Client          Client
// 	Start           int                   `json:"start"`
// 	End             int                   `json:"end"`
// 	Page            int                   `json:"page"`
// 	PageSize        int                   `json:"page_size"`
// 	Uri             string                `json:"uri"`
// 	FirstPageUri    string                `json:"first_page_uri"`
// 	PreviousPageUri string                `json"previous_page_uri"`
// 	NextPageUri     string                `json:"next_page_uri"`
// 	Countries       []IncomingPhoneNumber `json:"countries"`
// }
//
// type AvailablePhoneNumberCountries struct {
// 	Country         string `json:"country"`
// 	CountryCode     string `json:"country_code"`
// 	Beta            bool   `json:"beta"`
// 	AvaliableTypes  AvailablePhoneNumberTypes
// 	SubresourceURIs map[string]interface{} `json:"subresource_uris"`
// }
//
// func GetAvailablePhoneNumberCountries(client Client, beta Optional) ([]AvailablePhoneNumberCountries, err) {
// 	var list AvaliablePhoneNumberCountryList
// 	var avaliablePhoneNumberCountries []AvailablePhoneNumberCountries
//
// 	params := url.Values{}
// 	if beta != nil {
// 		param, value := beta.GetParam()
// 		params.Set(param, value)
// 	}
// 	res, err := client.get(params, "/AvailablePhoneNumbers.json")
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = json.Unmarshal(res, &list)
// }
