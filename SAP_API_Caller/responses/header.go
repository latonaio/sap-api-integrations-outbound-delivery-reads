package responses

type Header struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			DeliveryDocument              string `json:"DeliveryDocument"`
			DeliveryDocumentType          string `json:"DeliveryDocumentType"`
			DocumentDate                  string `json:"DocumentDate"`
			ActualGoodsMovementDate       string `json:"ActualGoodsMovementDate"`
			ActualDeliveryRoute           string `json:"ActualDeliveryRoute"`
            Shippinglocationtimezone      string `json:"Shippinglocationtimezone"`
            Receivinglocationtimezone     string `json:"Receivinglocationtimezone"`
			ActualGoodsMovementTime       string `json:"ActualGoodsMovementTime"`
			BillingDocumentDate           string `json:"BillingDocumentDate"`
			CompleteDeliveryIsDefined     bool   `json:"CompleteDeliveryIsDefined"`
			ConfirmationTime              string `json:"ConfirmationTime"`
			CreationDate                  string `json:"CreationDate"`
			CreationTime                  string `json:"CreationTime"`
			CustomerGroup                 string `json:"CustomerGroup"`
			DeliveryBlockReason           string `json:"DeliveryBlockReason"`
			DeliveryDate                  string `json:"DeliveryDate"`
			DeliveryDocumentBySupplier    string `json:"DeliveryDocumentBySupplier"`
			DeliveryIsInPlant             bool   `json:"DeliveryIsInPlant"`
			DeliveryPriority              string `json:"DeliveryPriority"`
			DeliveryTime                  string `json:"DeliveryTime"`
			GoodsIssueOrReceiptSlipNumber string `json:"GoodsIssueOrReceiptSlipNumber"`
			GoodsIssueTime                string `json:"GoodsIssueTime"`
			HeaderBillingBlockReason      string `json:"HeaderBillingBlockReason"`
			HeaderGrossWeight             string `json:"HeaderGrossWeight"`
			HeaderNetWeight               string `json:"HeaderNetWeight"`
			HeaderVolume                  string `json:"HeaderVolume"`
			HeaderVolumeUnit              string `json:"HeaderVolumeUnit"`
			HeaderWeightUnit              string `json:"HeaderWeightUnit"`
			IncotermsClassification       string `json:"IncotermsClassification"`
			IsExportDelivery              string `json:"IsExportDelivery"`
			LastChangeDate                string `json:"LastChangeDate"`
			LoadingDate                   string `json:"LoadingDate"`
			LoadingPoint                  string `json:"LoadingPoint"`
			LoadingTime                   string `json:"LoadingTime"`
			MeansOfTransport              string `json:"MeansOfTransport"`
			OrderCombinationIsAllowed     bool   `json:"OrderCombinationIsAllowed"`
			OrderID                       string `json:"OrderID"`
			OverallDelivConfStatus        string `json:"OverallDelivConfStatus"`
			OverallDelivReltdBillgStatus  string `json:"OverallDelivReltdBillgStatus"`
			OverallGoodsMovementStatus    string `json:"OverallGoodsMovementStatus"`
			OverallPackingStatus          string `json:"OverallPackingStatus"`
			OverallPickingConfStatus      string `json:"OverallPickingConfStatus"`
			OverallPickingStatus          string `json:"OverallPickingStatus"`
			PickingDate                   string `json:"PickingDate"`
			PickingTime                   string `json:"PickingTime"`
			PlannedGoodsIssueDate         string `json:"PlannedGoodsIssueDate"`
			ReceivingPlant                string `json:"ReceivingPlant"`
			ShippingCondition             string `json:"ShippingCondition"`
			ShippingPoint                 string `json:"ShippingPoint"`
			ShippingType                  string `json:"ShippingType"`
			ShipToParty                   string `json:"ShipToParty"`
			SoldToParty                   string `json:"SoldToParty"`
			Supplier                      string `json:"Supplier"`
			TransportationGroup           string `json:"TransportationGroup"`
			TransportationPlanningDate    string `json:"TransportationPlanningDate"`
			TransportationPlanningTime    string `json:"TransportationPlanningTime"`
			ToHeaderPartner struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_DeliveryDocumentPartner"`
			ToItem struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_DeliveryDocumentItem"`
		} `json:"results"`
	} `json:"d"`
}
