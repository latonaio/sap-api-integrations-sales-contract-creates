package sap_api_input_reader

import (
	"sap-api-integrations-sales-contract-creates/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.SalesContract
	return &requests.Header{
		SalesContract:                  data.SalesContract,
		SalesContractType:              data.SalesContractType,
		SalesOrganization:              data.SalesOrganization,
		DistributionChannel:            data.DistributionChannel,
		OrganizationDivision:           data.OrganizationDivision,
		SalesGroup:                     data.SalesGroup,
		SalesOffice:                    data.SalesOffice,
		SalesDistrict:                  data.SalesDistrict,
		SoldToParty:                    data.SoldToParty,
		CreationDate:                   data.CreationDate,
		LastChangeDate:                 data.LastChangeDate,
		PurchaseOrderByCustomer:        data.PurchaseOrderByCustomer,
		CustomerPurchaseOrderDate:      data.CustomerPurchaseOrderDate,
		SalesContractDate:              data.SalesContractDate,
		TotalNetAmount:                 data.TotalNetAmount,
		TransactionCurrency:            data.TransactionCurrency,
		SDDocumentReason:               data.SDDocumentReason,
		PricingDate:                    data.PricingDate,
		IncotermsClassification:        data.IncotermsClassification,
		CustomerPaymentTerms:           data.CustomerPaymentTerms,
		PaymentMethod:                  data.PaymentMethod,
		SalesContractValidityStartDate: data.SalesContractValidityStartDate,
		SalesContractValidityEndDate:   data.SalesContractValidityEndDate,
		SalesContractSignedDate:        data.SalesContractSignedDate,
		ReferenceSDDocument:            data.ReferenceSDDocument,
		ReferenceSDDocumentCategory:    data.ReferenceSDDocumentCategory,
		SalesDocApprovalStatus:         data.SalesDocApprovalStatus,
		SalesContractApprovalReason:    data.SalesContractApprovalReason,
		OverallSDProcessStatus:         data.OverallSDProcessStatus,
		TotalCreditCheckStatus:         data.TotalCreditCheckStatus,
		OverallSDDocumentRejectionSts:  data.OverallSDDocumentRejectionSts,
	}
}

func (sdc *SDC) ConvertToItem() *requests.Item {
	dataSalesContract := sdc.SalesContract
	data := sdc.SalesContract.SalesContractItem
	return &requests.Item{
		SalesContract:                  dataSalesContract.SalesContract,
		SalesContractItem:              data.SalesContractItem,
		SalesContractItemCategory:      data.SalesContractItemCategory,
		SalesContractItemText:          data.SalesContractItemText,
		PurchaseOrderByCustomer:        data.PurchaseOrderByCustomer,
		Material:                       data.Material,
		MaterialByCustomer:             data.MaterialByCustomer,
		PricingDate:                    data.PricingDate,
		RequestedQuantity:              data.RequestedQuantity,
		RequestedQuantityUnit:          data.RequestedQuantityUnit,
		ItemGrossWeight:                data.ItemGrossWeight,
		ItemNetWeight:                  data.ItemNetWeight,
		ItemWeightUnit:                 data.ItemWeightUnit,
		ItemVolume:                     data.ItemVolume,
		ItemVolumeUnit:                 data.ItemVolumeUnit,
		TransactionCurrency:            data.TransactionCurrency,
		NetAmount:                      data.NetAmount,
		MaterialGroup:                  data.MaterialGroup,
		MaterialPricingGroup:           data.MaterialPricingGroup,
		Batch:                          data.Batch,
		ProductionPlant:                data.ProductionPlant,
		StorageLocation:                data.StorageLocation,
		ShippingPoint:                  data.ShippingPoint,
		IncotermsClassification:        data.IncotermsClassification,
		CustomerPaymentTerms:           data.CustomerPaymentTerms,
		SalesDocumentRjcnReason:        data.SalesDocumentRjcnReason,
		ItemBillingBlockReason:         data.ItemBillingBlockReason,
		WBSElement:                     data.WBSElement,
		ProfitCenter:                   data.ProfitCenter,
		ReferenceSDDocument:            data.ReferenceSDDocument,
		ReferenceSDDocumentItem:        data.ReferenceSDDocumentItem,
		SDProcessStatus:                data.SDProcessStatus,
		SalesContractValidityStartDate: data.SalesContractValidityStartDate,
		SalesContractValidityEndDate:   data.SalesContractValidityEndDate,
		SalesContractSignedDate:        data.SalesContractSignedDate,
	}
}
