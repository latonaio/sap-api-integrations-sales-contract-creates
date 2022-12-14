# sap-api-integrations-sales-contract-creates  
sap-api-integrations-sales-contract-creates は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で 販売基本契約データ を登録するマイクロサービスです。  
sap-api-integrations-sales-contract-creates には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-sales-contract-creates は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/OP_API_SALES_CONTRACT/overview  

## 動作環境  
sap-api-integrations-sales-contract-creates は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）   
・ AION のリソース （推奨)   
・ OS: LinuxOS （必須）   
・ CPU: ARM/AMD/Intel（いずれか必須）  

## クラウド環境での利用
sap-api-integrations-sales-contract-creates は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。


## 本レポジトリ が 対応する API サービス
sap-api-integrations-sales-contract-creates が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_SALES_CONTRACT/overview  
* APIサービス名(=baseURL): API_SALES_CONTRACT_SRV

## 本レポジトリ に 含まれる API名
sap-api-integrations-sales-contract-creates には、次の API をコールするためのリソースが含まれています。  

* A_SalesContract（販売基本契約 - ヘッダ）
* A_SalesContractItem（販売基本契約 - 明細）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に登録したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて登録することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header" が指定されています。    
  
```
"api_schema": "SAPSalesContractCreates",
"accepter": ["Header"],
"sales_contract": "",
"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
"api_schema": "SAPSalesContractCreates",
"accepter": ["All"],
"sales_contract": "",
"deleted": false
```
## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncPostSalesContract(
	header            *requests.Header,
	item              *requests.Item,
	accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(header)
				wg.Done()
			}()
		case "Item":
			func() {
				c.Item(item)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

## Output  
本マイクロサービスでは、[golang-logging-library](https://github.com/latonaio/golang-logging-library) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 販売基本契約 の ヘッダデータ が登録された結果の JSON の例です。  
以下の項目のうち、"SalesContract" ～ "to_Item" は、/SAP_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
	"cursor": "/home/ampamman/go/src/sap-api-integrations-sales-contract-creates/SAP_API_Caller/caller.go#L113",
	"function": "sap-api-integrations-sales-contract-creates/SAP_API_Caller.(*SAPAPICaller).Header",
	"level": "INFO",
	"message": {
		"SalesContract": "40000000",
		"SalesContractType": "CQ",
		"SalesOrganization": "1710",
		"DistributionChannel": "10",
		"OrganizationDivision": "00",
		"SalesGroup": "",
		"SalesOffice": "",
		"SalesDistrict": "",
		"SoldToParty": "17100001",
		"CreationDate": "2017-06-23T09:00:00+09:00",
		"LastChangeDate": "",
		"PurchaseOrderByCustomer": "123",
		"CustomerPurchaseOrderDate": "",
		"SalesContractDate": "2017-06-23T09:00:00+09:00",
		"TotalNetAmount": "1755.00",
		"TransactionCurrency": "USD",
		"SDDocumentReason": "",
		"PricingDate": "2017-06-23T09:00:00+09:00",
		"IncotermsClassification": "EXW",
		"CustomerPaymentTerms": "0004",
		"PaymentMethod": "",
		"SalesContractValidityStartDate": "2017-06-23T09:00:00+09:00",
		"SalesContractValidityEndDate": "2017-07-31T09:00:00+09:00",
		"SalesContractSignedDate": "",
		"ReferenceSDDocument": "",
		"ReferenceSDDocumentCategory": "",
		"SalesDocApprovalStatus": "",
		"SalesContractApprovalReason": "",
		"OverallSDProcessStatus": "A",
		"TotalCreditCheckStatus": "",
		"OverallSDDocumentRejectionSts": "A",
		"to_Partner": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_SALES_CONTRACT_SRV/A_SalesContract('40000000')/to_Partner",
		"to_Item": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_SALES_CONTRACT_SRV/A_SalesContract('40000000')/to_Item"
	},
	"time": "2022-09-06T17:47:48+09:00"
```
