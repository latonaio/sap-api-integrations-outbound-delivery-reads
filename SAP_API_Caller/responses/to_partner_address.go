package responses

type ToPartnerAddress struct {
	D struct {
		Metadata struct {
			ID   string `json:"id"`
			URI  string `json:"uri"`
			Type string `json:"type"`
		} `json:"__metadata"`
			AddressID              string `json:"AddressID"`
			BusinessPartnerName1   string `json:"BusinessPartnerName1"`
			Country                string `json:"Country"`
			Region                 string `json:"Region"`
			StreetName             string `json:"StreetName"`
			CityName               string `json:"CityName"`
			PostalCode             string `json:"PostalCode"`
			CorrespondenceLanguage string `json:"CorrespondenceLanguage"`
			FaxNumber              string `json:"FaxNumber"`
			PhoneNumber            string `json:"PhoneNumber"`
	} `json:"d"`
}
