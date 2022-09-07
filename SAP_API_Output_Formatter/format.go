package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-sales-contract-creates/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) (*Header, error) {
	pm := &responses.Header{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	header := &Header{
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

	return header, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) (*Item, error) {
	pm := &responses.Item{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	item := &Item{
		SalesContract:                  data.SalesContract,
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

	return item, nil
}
