package sap_api_output_formatter

type SalesContract struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Product       string `json:"Product"`
	APISchema     string `json:"api_schema"`
	SalesContract string `json:"sales_contract"`
	Deleted       string `json:"deleted"`
}

type Header struct {
	SalesContract                  string      `json:"SalesContract"`
	SalesContractType              string      `json:"SalesContractType"`
	SalesOrganization              string      `json:"SalesOrganization"`
	DistributionChannel            string      `json:"DistributionChannel"`
	OrganizationDivision           string      `json:"OrganizationDivision"`
	SalesGroup                     string      `json:"SalesGroup"`
	SalesOffice                    string      `json:"SalesOffice"`
	SalesDistrict                  string      `json:"SalesDistrict"`
	SoldToParty                    string      `json:"SoldToParty"`
	CreationDate                   string      `json:"CreationDate"`
	LastChangeDate                 string      `json:"LastChangeDate"`
	PurchaseOrderByCustomer        string      `json:"PurchaseOrderByCustomer"`
	CustomerPurchaseOrderDate      string      `json:"CustomerPurchaseOrderDate"`
	SalesContractDate              string      `json:"SalesContractDate"`
	TotalNetAmount                 string      `json:"TotalNetAmount"`
	TransactionCurrency            string      `json:"TransactionCurrency"`
	SDDocumentReason               string      `json:"SDDocumentReason"`
	PricingDate                    string      `json:"PricingDate"`
	IncotermsClassification        string      `json:"IncotermsClassification"`
	CustomerPaymentTerms           string      `json:"CustomerPaymentTerms"`
	PaymentMethod                  string      `json:"PaymentMethod"`
	SalesContractValidityStartDate string      `json:"SalesContractValidityStartDate"`
	SalesContractValidityEndDate   string      `json:"SalesContractValidityEndDate"`
	SalesContractSignedDate        string      `json:"SalesContractSignedDate"`
	ReferenceSDDocument            string      `json:"ReferenceSDDocument"`
	ReferenceSDDocumentCategory    string      `json:"ReferenceSDDocumentCategory"`
	SalesDocApprovalStatus         string      `json:"SalesDocApprovalStatus"`
	SalesContractApprovalReason    string      `json:"SalesContractApprovalReason"`
	OverallSDProcessStatus         string      `json:"OverallSDProcessStatus"`
	TotalCreditCheckStatus         string      `json:"TotalCreditCheckStatus"`
	OverallSDDocumentRejectionSts  string      `json:"OverallSDDocumentRejectionSts"`
}

type Item struct {
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
}
