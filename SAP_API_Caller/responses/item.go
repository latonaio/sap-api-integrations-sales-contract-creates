package responses

type Item struct {
	D struct {
			SalesContract                  string      `json:"SalesContract"`
			SalesContractItem              string      `json:"SalesContractItem"`
			SalesContractItemCategory      string      `json:"SalesContractItemCategory"`
			SalesContractItemText          string      `json:"SalesContractItemText"`
			PurchaseOrderByCustomer        string      `json:"PurchaseOrderByCustomer"`
			Material                       string      `json:"Material"`
			MaterialByCustomer             string      `json:"MaterialByCustomer"`
			PricingDate                    string      `json:"PricingDate"`
			RequestedQuantity              string      `json:"RequestedQuantity"`
			RequestedQuantityUnit          string      `json:"RequestedQuantityUnit"`
			ItemGrossWeight                string      `json:"ItemGrossWeight"`
			ItemNetWeight                  string      `json:"ItemNetWeight"`
			ItemWeightUnit                 string      `json:"ItemWeightUnit"`
			ItemVolume                     string      `json:"ItemVolume"`
			ItemVolumeUnit                 string      `json:"ItemVolumeUnit"`
			TransactionCurrency            string      `json:"TransactionCurrency"`
			NetAmount                      string      `json:"NetAmount"`
			MaterialGroup                  string      `json:"MaterialGroup"`
			MaterialPricingGroup           string      `json:"MaterialPricingGroup"`
			Batch                          string      `json:"Batch"`
			ProductionPlant                string      `json:"ProductionPlant"`
			StorageLocation                string      `json:"StorageLocation"`
			ShippingPoint                  string      `json:"ShippingPoint"`
			IncotermsClassification        string      `json:"IncotermsClassification"`
			CustomerPaymentTerms           string      `json:"CustomerPaymentTerms"`
			SalesDocumentRjcnReason        string      `json:"SalesDocumentRjcnReason"`
			ItemBillingBlockReason         string      `json:"ItemBillingBlockReason"`
			WBSElement                     string      `json:"WBSElement"`
			ProfitCenter                   string      `json:"ProfitCenter"`
			ReferenceSDDocument            string      `json:"ReferenceSDDocument"`
			ReferenceSDDocumentItem        string      `json:"ReferenceSDDocumentItem"`
			SDProcessStatus                string      `json:"SDProcessStatus"`
			SalesContractValidityStartDate string      `json:"SalesContractValidityStartDate"`
			SalesContractValidityEndDate   string      `json:"SalesContractValidityEndDate"`
			SalesContractSignedDate        string      `json:"SalesContractSignedDate"`
	} `json:"d"`
}
