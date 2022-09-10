package responses

type ToItemDocumentFlow struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			Deliveryversion                string `json:"Deliveryversion"`
			PrecedingDocument              string `json:"PrecedingDocument"`
			PrecedingDocumentCategory      string `json:"PrecedingDocumentCategory"`
			PrecedingDocumentItem          string `json:"PrecedingDocumentItem"`
			Subsequentdocument             string `json:"Subsequentdocument"`
			QuantityInBaseUnit             string `json:"QuantityInBaseUnit"`
			SubsequentDocumentItem         string `json:"SubsequentDocumentItem"`
			SDFulfillmentCalculationRule   string `json:"SDFulfillmentCalculationRule"`
			SubsequentDocumentCategory     string `json:"SubsequentDocumentCategory"`
			TransferOrderInWrhsMgmtIsConfd bool   `json:"TransferOrderInWrhsMgmtIsConfd"`
		} `json:"results"`
	} `json:"d"`
}
