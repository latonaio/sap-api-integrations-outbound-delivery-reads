package responses

type ToHeaderPartner struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			AddressID                     string `json:"AddressID"`
			Customer                      string `json:"Customer"`
			PartnerFunction               string `json:"PartnerFunction"`
			SDDocument                    string `json:"SDDocument"`
			SDDocumentItem                string `json:"SDDocumentItem"`
			Supplier                      string `json:"Supplier"`
			ToPartnerAddress              struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Address"`
		} `json:"results"`
	} `json:"d"`
}
