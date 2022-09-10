package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-outbound-delivery-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
			DeliveryDocument:              data.DeliveryDocument,
		DeliveryDocumentType:          data.DeliveryDocumentType,
		DocumentDate:                  data.DocumentDate,
		ActualGoodsMovementDate:       data.ActualGoodsMovementDate,
		ActualDeliveryRoute:           data.ActualDeliveryRoute,
        Shippinglocationtimezone:      data.Shippinglocationtimezone,
        Receivinglocationtimezone:     data.Receivinglocationtimezone,
		ActualGoodsMovementTime:       data.ActualGoodsMovementTime,
		BillingDocumentDate:           data.BillingDocumentDate,
		CompleteDeliveryIsDefined:     data.CompleteDeliveryIsDefined,
		ConfirmationTime:              data.ConfirmationTime,
		CreationDate:                  data.CreationDate,
		CreationTime:                  data.CreationTime,
		CustomerGroup:                 data.CustomerGroup,
		DeliveryBlockReason:           data.DeliveryBlockReason,
		DeliveryDate:                  data.DeliveryDate,
		DeliveryDocumentBySupplier:    data.DeliveryDocumentBySupplier,
		DeliveryIsInPlant:             data.DeliveryIsInPlant,
		DeliveryPriority:              data.DeliveryPriority,
		DeliveryTime:                  data.DeliveryTime,
		GoodsIssueOrReceiptSlipNumber: data.GoodsIssueOrReceiptSlipNumber,
		GoodsIssueTime:                data.GoodsIssueTime,
		HeaderBillingBlockReason:      data.HeaderBillingBlockReason,
		HeaderGrossWeight:             data.HeaderGrossWeight,
		HeaderNetWeight:               data.HeaderNetWeight,
		HeaderVolume:                  data.HeaderVolume,
		HeaderVolumeUnit:              data.HeaderVolumeUnit,
		HeaderWeightUnit:              data.HeaderWeightUnit,
		IncotermsClassification:       data.IncotermsClassification,
		IsExportDelivery:              data.IsExportDelivery,
		LastChangeDate:                data.LastChangeDate,
		LoadingDate:                   data.LoadingDate,
		LoadingPoint:                  data.LoadingPoint,
		LoadingTime:                   data.LoadingTime,
		MeansOfTransport:              data.MeansOfTransport,
		OrderCombinationIsAllowed:     data.OrderCombinationIsAllowed,
		OrderID:                       data.OrderID,
		OverallDelivConfStatus:        data.OverallDelivConfStatus,
		OverallDelivReltdBillgStatus:  data.OverallDelivReltdBillgStatus,
		OverallGoodsMovementStatus:    data.OverallGoodsMovementStatus,
		OverallPackingStatus:          data.OverallPackingStatus,
		OverallPickingConfStatus:      data.OverallPickingConfStatus,
		OverallPickingStatus:          data.OverallPickingStatus,
		PickingDate:                   data.PickingDate,
		PickingTime:                   data.PickingTime,
		PlannedGoodsIssueDate:         data.PlannedGoodsIssueDate,
		ReceivingPlant:                data.ReceivingPlant,
		ShippingCondition:             data.ShippingCondition,
		ShippingPoint:                 data.ShippingPoint,
		ShippingType:                  data.ShippingType,
		ShipToParty:                   data.ShipToParty,
		SoldToParty:                   data.SoldToParty,
		Supplier:                      data.Supplier,
		TransportationGroup:           data.TransportationGroup,
		TransportationPlanningDate:    data.TransportationPlanningDate,
		TransportationPlanningTime:    data.TransportationPlanningTime,
		ToHeaderPartner:               data.ToHeaderPartner.Deferred.URI,
		ToItem:                        data.ToItem.Deferred.URI,
		})
	}

	return header, nil
}

func ConvertToHeaderPartner(raw []byte, l *logger.Logger) ([]HeaderPartner, error) {
	pm := &responses.HeaderPartner{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to HeaderPartner. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	headerPartner := make([]HeaderPartner, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		headerPartner = append(headerPartner, HeaderPartner{
			AddressID:              data.AddressID,
			Customer:               data.Customer,
			PartnerFunction:        data.PartnerFunction,
			SDDocument:             data.SDDocument,
			SDDocumentItem:         data.SDDocumentItem,
			Supplier:               data.Supplier,
			ToPartnerAddress:       data.ToPartnerAddress.Deferred.URI,
		})
	}

	return headerPartner, nil
}

func ConvertToPartnerAddress(raw []byte, l *logger.Logger) (*PartnerAddress, error) {
	pm := &responses.PartnerAddress{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to PartnerAddress. unmarshal error: %w", err)
	}
	if pm.D == nil {
		return nil, xerrors.New("Metadata data is not exist")
	}
	data := pm.D

	return &PartnerAddress{
        AddressID:                     data.AddressID, 
        BusinessPartnerName1:          data.BusinessPartnerName1,
        Country:                       data.Country,
		Region:                        data.Region,
		StreetName:                    data.StreetName,
		CityName:                      data.CityName,
		PostalCode:                    data.PostalCode,
		CorrespondenceLanguage:        data.CorrespondenceLanguage,
		FaxNumber:                     data.FaxNumber,
		PhoneNumber:                   data.PhoneNumber,
	}, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) ([]Item, error) {
	pm := &responses.Item{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	item := make([]Item, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		item = append(item, Item{		
		DeliveryDocument:              data.DeliveryDocument,
        DeliveryDocumentItem:          data.DeliveryDocumentItem,
		BaseUnit:                      data.BaseUnit,
		ActualDeliveryQuantity:        data.ActualDeliveryQuantity,
		Batch:                         data.Batch,
		BatchBySupplier:               data.BatchBySupplier,
		CostCenter:                    data.CostCenter,
		CreationDate:                  data.CreationDate,
		CreationTime:                  data.CreationTime,
		DeliveryDocumentItemCategory:  data.DeliveryDocumentItemCategory,
		DeliveryDocumentItemText:      data.DeliveryDocumentItemText,
		DeliveryQuantityUnit:          data.DeliveryQuantityUnit,
		DistributionChannel:           data.DistributionChannel,
		Division:                      data.Division,
		GLAccount:                     data.GLAccount,
		GoodsMovementStatus:           data.GoodsMovementStatus,
		GoodsMovementType:             data.GoodsMovementType,
		InternationalArticleNumber:    data.InternationalArticleNumber,
		InventorySpecialStockType:     data.InventorySpecialStockType,
		IsCompletelyDelivered:         data.IsCompletelyDelivered,
		ItemBillingBlockReason:        data.ItemBillingBlockReason,
		ItemDeliveryIncompletionStatus: data.ItemDeliveryIncompletionStatus,
		ItemGdsMvtIncompletionSts:     data.ItemGdsMvtIncompletionSts,
		ItemGrossWeight:               data.ItemGrossWeight,
		ItemNetWeight:                 data.ItemNetWeight,
		ItemWeightUnit:                data.ItemWeightUnit,
		ItemIsBillingRelevant:         data.ItemIsBillingRelevant,
		ItemPackingIncompletionStatus: data.ItemPackingIncompletionStatus,
		ItemPickingIncompletionStatus: data.ItemPickingIncompletionStatus,
		ItemVolume:                    data.ItemVolume,
		LastChangeDate:                data.LastChangeDate,
		LoadingGroup:                  data.LoadingGroup,
		Material:                      data.Material,
		MaterialByCustomer:            data.MaterialByCustomer,
		MaterialFreightGroup:          data.MaterialFreightGroup,
		NumberOfSerialNumbers:         data.NumberOfSerialNumbers,
		OrderID:                       data.OrderID,
		OrderItem:                     data.OrderItem,
		OriginalDeliveryQuantity:      data.OriginalDeliveryQuantity,
		PackingStatus:                 data.PackingStatus,
		PartialDeliveryIsAllowed:      data.PartialDeliveryIsAllowed,
		PickingConfirmationStatus:     data.PickingConfirmationStatus,
		PickingStatus:                 data.PickingStatus,
		Plant:                         data.Plant,
		ProductAvailabilityDate:       data.ProductAvailabilityDate,
		ProductAvailabilityTime:       data.ProductAvailabilityTime,
		ProfitCenter:                  data.ProfitCenter,
		StorageLocation:               data.StorageLocation,
		TransportationGroup:           data.TransportationGroup,
		ToItemDocumentFlow:            data.ToItemDocumentFlow.Deferred.URI,
		})
	}

	return item, nil
}

func ConvertToToHeaderPartner(raw []byte, l *logger.Logger) ([]ToHeaderPartner, error) {
	pm := &responses.ToHeaderPartner{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToHeaderPartner. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toHeaderPartner := make([]ToHeaderPartner, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toHeaderPartner = append(toHeaderPartner, ToHeaderPartner{
	AddressID:              data.AddressID,
	Customer:               data.Customer,
	PartnerFunction:        data.PartnerFunction,
	SDDocument:             data.SDDocument,
	SDDocumentItem:         data.SDDocumentItem,
	Supplier:               data.Supplier,
	ToPartnerAddress:       data.ToPartnerAddress.Deferred.URI,
		})
	}

	return toHeaderPartner, nil
}

func ConvertToToPartnerAddress(raw []byte, l *logger.Logger) (*ToPartnerAddress, error) {
	pm := &responses.ToPartnerAddress{}

	err := json.Unmarshal(raw, &pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToPartnerAddress. unmarshal error: %w", err)
	}
	
	return &ToPartnerAddress{
        AddressID:                     pm.D.AddressID, 
        BusinessPartnerName1:          pm.D.BusinessPartnerName1,
        Country:                       pm.D.Country,
		Region:                        pm.D.Region,
		StreetName:                    pm.D.StreetName,
		CityName:                      pm.D.CityName,
		PostalCode:                    pm.D.PostalCode,
		CorrespondenceLanguage:        pm.D.CorrespondenceLanguage,
		FaxNumber:                     pm.D.FaxNumber,
		PhoneNumber:                   pm.D.PhoneNumber,
	}, nil
}

func ConvertToToItem(raw []byte, l *logger.Logger) ([]ToItem, error) {
	pm := &responses.ToItem{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItem. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toItem := make([]ToItem, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItem = append(toItem, ToItem{
		DeliveryDocument:              data.DeliveryDocument,
        DeliveryDocumentItem:          data.DeliveryDocumentItem,
		BaseUnit:                      data.BaseUnit,
		ActualDeliveryQuantity:        data.ActualDeliveryQuantity,
		Batch:                         data.Batch,
		BatchBySupplier:               data.BatchBySupplier,
		CostCenter:                    data.CostCenter,
		CreationDate:                  data.CreationDate,
		CreationTime:                  data.CreationTime,
		DeliveryDocumentItemCategory:  data.DeliveryDocumentItemCategory,
		DeliveryDocumentItemText:      data.DeliveryDocumentItemText,
		DeliveryQuantityUnit:          data.DeliveryQuantityUnit,
		DistributionChannel:           data.DistributionChannel,
		Division:                      data.Division,
		GLAccount:                     data.GLAccount,
		GoodsMovementStatus:           data.GoodsMovementStatus,
		GoodsMovementType:             data.GoodsMovementType,
		InternationalArticleNumber:    data.InternationalArticleNumber,
		InventorySpecialStockType:     data.InventorySpecialStockType,
		IsCompletelyDelivered:         data.IsCompletelyDelivered,
		ItemBillingBlockReason:        data.ItemBillingBlockReason,
		ItemDeliveryIncompletionStatus: data.ItemDeliveryIncompletionStatus,
		ItemGdsMvtIncompletionSts:     data.ItemGdsMvtIncompletionSts,
		ItemGrossWeight:               data.ItemGrossWeight,
		ItemNetWeight:                 data.ItemNetWeight,
		ItemWeightUnit:                data.ItemWeightUnit,
		ItemIsBillingRelevant:         data.ItemIsBillingRelevant,
		ItemPackingIncompletionStatus: data.ItemPackingIncompletionStatus,
		ItemPickingIncompletionStatus: data.ItemPickingIncompletionStatus,
		ItemVolume:                    data.ItemVolume,
		LastChangeDate:                data.LastChangeDate,
		LoadingGroup:                  data.LoadingGroup,
		Material:                      data.Material,
		MaterialByCustomer:            data.MaterialByCustomer,
		MaterialFreightGroup:          data.MaterialFreightGroup,
		NumberOfSerialNumbers:         data.NumberOfSerialNumbers,
		OrderID:                       data.OrderID,
		OrderItem:                     data.OrderItem,
		OriginalDeliveryQuantity:      data.OriginalDeliveryQuantity,
		PackingStatus:                 data.PackingStatus,
		PartialDeliveryIsAllowed:      data.PartialDeliveryIsAllowed,
		PickingConfirmationStatus:     data.PickingConfirmationStatus,
		PickingStatus:                 data.PickingStatus,
		Plant:                         data.Plant,
		ProductAvailabilityDate:       data.ProductAvailabilityDate,
		ProductAvailabilityTime:       data.ProductAvailabilityTime,
		ProfitCenter:                  data.ProfitCenter,
		StorageLocation:               data.StorageLocation,
		TransportationGroup:           data.TransportationGroup,
		ToItemDocumentFlow:            data.ToItemDocumentFlow.Deferred.URI,
		})
	}

	return toItem, nil
}


func ConvertToToItemDocumentFlow(raw []byte, l *logger.Logger) ([]ToItemDocumentFlow, error) {
	pm := &responses.ToItemDocumentFlow{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ItemDocumentFlow. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	toItemDocumentFlow := make([]ToItemDocumentFlow, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItemDocumentFlow = append(toItemDocumentFlow, ToItemDocumentFlow{
	Deliveryversion:                data.Deliveryversion,
	PrecedingDocument:              data.PrecedingDocument,
	PrecedingDocumentCategory:      data.PrecedingDocumentCategory,
	PrecedingDocumentItem:          data.PrecedingDocumentItem,
	Subsequentdocument:             data.Subsequentdocument,
	QuantityInBaseUnit:             data.QuantityInBaseUnit,
	SubsequentDocumentItem:         data.SubsequentDocumentItem,
	SDFulfillmentCalculationRule:   data.SDFulfillmentCalculationRule,
	SubsequentDocumentCategory:     data.SubsequentDocumentCategory,
	TransferOrderInWrhsMgmtIsConfd: data.TransferOrderInWrhsMgmtIsConfd,
		})
	}

	return toItemDocumentFlow, nil
}