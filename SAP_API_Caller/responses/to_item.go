package responses

type ToItem struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			DeliveryDocument               string `json:"DeliveryDocument"`
			DeliveryDocumentItem           string `json:"DeliveryDocumentItem"`
			BaseUnit                       string `json:"BaseUnit"`
			ActualDeliveryQuantity         string `json:"ActualDeliveryQuantity"`
			Batch                          string `json:"Batch"`
			BatchBySupplier                string `json:"BatchBySupplier"`
			CostCenter                     string `json:"CostCenter"`
			CreationDate                   string `json:"CreationDate"`
			CreationTime                   string `json:"CreationTime"`
			DeliveryDocumentItemCategory   string `json:"DeliveryDocumentItemCategory"`
			DeliveryDocumentItemText       string `json:"DeliveryDocumentItemText"`
			DeliveryQuantityUnit           string `json:"DeliveryQuantityUnit"`
			DistributionChannel            string `json:"DistributionChannel"`
			Division                       string `json:"Division"`
			GLAccount                      string `json:"GLAccount"`
			GoodsMovementStatus            string `json:"GoodsMovementStatus"`
			GoodsMovementType              string `json:"GoodsMovementType"`
			InternationalArticleNumber     string `json:"InternationalArticleNumber"`
			InventorySpecialStockType      string `json:"InventorySpecialStockType"`
			IsCompletelyDelivered          bool   `json:"IsCompletelyDelivered"`
			ItemBillingBlockReason         string `json:"ItemBillingBlockReason"`
			ItemDeliveryIncompletionStatus string `json:"ItemDeliveryIncompletionStatus"`
			ItemGdsMvtIncompletionSts      string `json:"ItemGdsMvtIncompletionSts"`
			ItemGrossWeight                string `json:"ItemGrossWeight"`
			ItemNetWeight                  string `json:"ItemNetWeight"`
			ItemWeightUnit                 string `json:"ItemWeightUnit"`
			ItemIsBillingRelevant          string `json:"ItemIsBillingRelevant"`
			ItemPackingIncompletionStatus  string `json:"ItemPackingIncompletionStatus"`
			ItemPickingIncompletionStatus  string `json:"ItemPickingIncompletionStatus"`
			ItemVolume                     string `json:"ItemVolume"`
			LastChangeDate                 string `json:"LastChangeDate"`
			LoadingGroup                   string `json:"LoadingGroup"`
			Material                       string `json:"Material"`
			MaterialByCustomer             string `json:"MaterialByCustomer"`
			MaterialFreightGroup           string `json:"MaterialFreightGroup"`
			NumberOfSerialNumbers          int    `json:"NumberOfSerialNumbers"`
			OrderID                        string `json:"OrderID"`
			OrderItem                      string `json:"OrderItem"`
			OriginalDeliveryQuantity       string `json:"OriginalDeliveryQuantity"`
			PackingStatus                  string `json:"PackingStatus"`
			PartialDeliveryIsAllowed       string `json:"PartialDeliveryIsAllowed"`
			PickingConfirmationStatus      string `json:"PickingConfirmationStatus"`
			PickingStatus                  string `json:"PickingStatus"`
			Plant                          string `json:"Plant"`
			ProductAvailabilityDate        string `json:"ProductAvailabilityDate"`
			ProductAvailabilityTime        string `json:"ProductAvailabilityTime"`
			ProfitCenter                   string `json:"ProfitCenter"`
			StorageLocation                string `json:"StorageLocation"`
			TransportationGroup            string `json:"TransportationGroup"`
			ToItemDocumentFlow struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_DocumentFlow"`
		} `json:"results"`
	} `json:"d"`
}
