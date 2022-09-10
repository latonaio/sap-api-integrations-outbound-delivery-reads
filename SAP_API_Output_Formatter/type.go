package sap_api_output_formatter

type OutboundDelivery struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Product       string `json:"Product"`
	APISchema     string `json:"api_schema"`
	DeliveryDocument      string `json:"delivery_document"`
	Deleted       string `json:"deleted"`
}

type Header struct {
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
	ToHeaderPartner               string `json:"to_DeliveryDocumentPartner"`
	ToItem                        string `json:"to_DeliveryDocumentItem"`
}

type HeaderPartner struct {
	AddressID              string `json:"AddressID"`
	Customer               string `json:"Customer"`
	PartnerFunction        string `json:"PartnerFunction"`
	SDDocument             string `json:"SDDocument"`
	SDDocumentItem         string `json:"SDDocumentItem"`
	Supplier               string `json:"Supplier"`
	ToPartnerAddress       string `json:"to_Address"`
}

type PartnerAddress struct {
	DeliveryDocument       string `json:"DeliveryDocument"`
	PartnerFunction        string `json:"PartnerFunction"`
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
}

type Item struct {
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
	ToItemDocumentFlow             string `json:"to_DocumentFlow"`
}

type ToHeaderPartner struct {
	AddressID              string `json:"AddressID"`
	Customer               string `json:"Customer"`
	PartnerFunction        string `json:"PartnerFunction"`
	SDDocument             string `json:"SDDocument"`
	SDDocumentItem         string `json:"SDDocumentItem"`
	Supplier               string `json:"Supplier"`
	ToPartnerAddress       string `json:"to_Address"`
}

type ToPartnerAddress struct {
	DeliveryDocument       string `json:"DeliveryDocument"`
	PartnerFunction        string `json:"PartnerFunction"`
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
}

type ToItem struct {
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
	ToItemDocumentFlow             string `json:"to_DocumentFlow"`
}

type ToItemDocumentFlow struct {
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
}
