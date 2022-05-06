# sap-api-integrations-business-partner-reads-c4
sap-api-integrations-business-partner-reads-c4 は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で ビジネスパートナ データを取得するマイクロサービスです。    
sap-api-integrations-business-partner-reads-c4 には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-business-partner-reads-c4 は、オンプレミス版である（＝クラウド版ではない）SAPC4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。   
https://api.sap.com/api/businesspartner/overview   

## 動作環境  
sap-api-integrations-business-partner-reads-c4 は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）    

## クラウド環境での利用
sap-api-integrations-business-partner-reads-c4 は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-business-partner-reads-c4 が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/businesspartner/overview    
* APIサービス名(=baseURL): c4codataapi

## 本レポジトリ に 含まれる API名
sap-api-integrations-business-partner-reads-c4 には、次の API をコールするためのリソースが含まれています。  

* PartnerCollection（ビジネスパートナ - パートナ）※パートナの詳細データを取得するために、ToPartnerAddress、ToPartnerHasPartnerContact、と合わせて利用されます。
* ToPartnerAddress（ビジネスパートナ - アドレス ※To）
* ToPartnerHasPartnerContact（ビジネスパートナ - コンタクト ※To）
* ContactCollection（ビジネスパートナ - コンタクト）※コンタクトの詳細データを取得するために、ToContactIsContactPersonFor、ToCorporateAccountと合わせて利用されます。
* ToContactIsContactPersonFor（ビジネスパートナ - コンタクトパーソン ※To）
* ToCorpotrateAccount（ビジネスパートナ - 会社アカウント ※To）
* BusinessUserCollection（ビジネスパートナ - ビジネスユーザ）※ビジネスユーザの詳細データを取得するために、ToBusinessRoleAssignment、ToEmployeeBasicData、と合わせて取得されます。
* ToBusinessRoleAssignment（ビジネスパートナ - ビジネスロール割当）
* ToEmployeeBasicData（ビジネスパートナ - 従業員基本情報）
* IndividualCustomerCollection（ビジネスパートナ - 個人顧客）※個人顧客の詳細データを取得するために、ToIndividualCustomerAddressと合わせて利用されます。
* ToIndividualCustomerAddress（ビジネスパートナ - 個人顧客住所 ※To）
* CorporateAccountCollection（ビジネスパートナ - 会社住所）
* ServiceAgentCollection（ビジネスパートナ - サービスエージェント）
* CompetitorCollection（ビジネスパートナ - 競合）
* EmployeeCollection（ビジネスパートナ - 従業員）※従業員の詳細データを取得するために、ToEmployeeOrganisationalUnitAssignment、ToEmployeeSalesResponsibility、ToEmployeeUserBusinessRoleAssignment、ToEmployeeUserBusinessRoleAssignment、ToEmployeeWorkingHours、と合わせて利用されます。
* ToEmployeeOrganisationalUnitAssignment（ビジネスパートナ - 従業員組織ユニット割当 ※To）
* ToEmployeeSalesResponsibility（ビジネスパートナ - 従業員販売責任 ※To）
* ToEmployeeUserBusinessRoleAssignment（ビジネスパートナ - 従業員ユーザービジネスロール割当 ※To）
* ToEmployeeWorkingHours（ビジネスパートナ - 従業員労働時間 ※To）

## API への 値入力条件 の 初期値
sap-api-integrations-business-partner-reads-c4 において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.PartnerCollection.PartnerID（パートナID）
* inoutSDC.ContactCollection.ContactID（コンタクトID）
* inoutSDC.BusinessUserCollection.UserID（ユーザーID）
* inoutSDC.ContactCollection.ContactIsContactPersonFor.CorporateAccount.IndividualCustomerCollection.CustomerID（顧客ID）
* inoutSDC.ContactCollection.ContactIsContactPersonFor.CorporateAccount.AccountID（アカウントID）
* inoutSDC.ServiceAgentCollection.ServiceAgentID（サービスエージェントID）
* inoutSDC.CompetitorCollection.CompetitorID（競合ID）
* inoutSDC.EmployeeCollection.EmployeeID（従業員ID）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"PartnerCollection" が指定されています。    
  
```
	"api_schema": "BusinessPartnerCollection",
	"accepter": ["PartnerCollection"],
	"business_partner_code": "1000410",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "BusinessPartnerCollection",
	"accepter": ["All"],
	"business_partner_code": "1000410",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetBP(partnerID, contactID, userID, customerID, accountID, serviceAgentID, competitorID, employeeID string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "PartnerCollection":
			func() {
				c.PartnerCollection(partnerID)
				wg.Done()
			}()
		case "ContactCollection":
			func() {
				c.ContactCollection(contactID)
				wg.Done()
			}()
		case "BusinessUserCollection":
			func() {
				c.BusinessUserCollection(userID)
				wg.Done()
			}()
		case "IndividualCustomerCollection":
			func() {
				c.IndividualCustomerCollection(customerID)
				wg.Done()
			}()
		case "CorporateAccount":
			func() {
				c.CorporateAccount(accountID)
				wg.Done()
			}()
		case "ServiceAgentCollection":
			func() {
				c.ServiceAgentCollection(serviceAgentID)
				wg.Done()
			}()
		case "CompetitorCollection":
			func() {
				c.CompetitorCollection(competitorID)
				wg.Done()
			}()
		case "EmployeeCollection":
			func() {
				c.EmployeeCollection(employeeID)
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
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP ビジネスパートナ の パートナデータ が取得された結果の JSON の例です。  
以下の項目のうち、"ObjectID" ～ "PartnerHasPartnerContact" は、/SAP_API_Output_Formatter/type.go 内 の Type PartnerCollection {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-business-partner-reads-c4/SAP_API_Caller/caller.go#L88",
	"function": "sap-api-integrations-business-partner-reads-c4/SAP_API_Caller.(*SAPAPICaller).PartnerCollection",
	"level": "INFO",
	"message": [
		{
			"ObjectID": "00163E038C701ED298AFFC9A58CB1677",
			"PartnerID": "1000410",
			"PartnerUUID": "00163E03-8C70-1ED2-98AF-FC9A58CB1677",
			"LifeCycleStatusCode": "2",
			"LifeCycleStatusCodeText": "Active",
			"DUNS": "",
			"LegalFormCode": "",
			"LegalFormCodeText": "",
			"PartnerABCClassificationCode": "A",
			"PartnerABCClassificationCodeText": "A-Partner",
			"IndustrialSectorCode": "23",
			"IndustrialSectorCodeText": "Construction",
			"ContactPermissionCode": "",
			"ContactPermissionCodeText": "",
			"BusinessPartnerFormattedName": "Jaxshare",
			"Name": "Jaxshare",
			"AdditionalName": "",
			"AdditionalName2": "",
			"AdditionalName3": "",
			"FormattedPostalAddressDescription": "256 West El Camino Real / Mountian View 94041 / US",
			"CountryCode": "US",
			"CountryCodeText": "United States",
			"StateCode": "CA",
			"StateCodeText": "California",
			"CareOfName": "",
			"AddressLine1": "",
			"AddressLine2": "",
			"HouseNumber": "256",
			"Street": "West El Camino Real",
			"AddressLine4": "",
			"AddressLine5": "",
			"District": "",
			"City": "Mountian View",
			"DifferentCity": "",
			"StreetPostalCode": "94041",
			"County": "",
			"CompanyPostalCode": "",
			"POBoxIndicator": false,
			"POBox": "",
			"POBoxPostalCode": "",
			"POBoxDeviatingCountryCode": "",
			"POBoxDeviatingCountryCodeText": "",
			"POBoxDeviatingStateCode": "",
			"POBoxDeviatingStateCodeText": "",
			"POBoxDeviatingCity": "",
			"TimeZoneCode": "PST",
			"TimeZoneCodeText": "(UTC-08:00) Pacific Time (Los Angeles, Tijuana, Yukon)",
			"Building": "",
			"Floor": "",
			"Room": "",
			"Phone": "+1 650917431",
			"Mobile": "",
			"Fax": "+1 6509278731",
			"Email": "",
			"WebSite": "",
			"LanguageCode": "",
			"LanguageCodeText": "",
			"BestReachedByCode": "",
			"BestReachedByCodeText": "",
			"TotalPoints": "0.00 ",
			"PartnerLevel": "",
			"PartnerLevelText": "",
			"NormalisedPhone": "+1650917431",
			"NormalisedMobile": "",
			"CreatedOn": "2013-01-19T00:07:03+09:00",
			"CreatedBy": "SAP WORKER",
			"CreatedByIdentityUUID": "00163E03-A070-1EE2-88BA-39BD20F290B5",
			"ChangedOn": "2013-01-19T00:27:26+09:00",
			"ChangedBy": "Eddie Smoke",
			"ChangedByIdentityUUID": "00163E03-A070-1EE2-88BA-39BD20F290B5",
			"EntityLastChangedOn": "2013-01-19T00:27:26+09:00",
			"ETag": "2013-01-19T00:27:26+09:00",
			"PartnerAddress": "https://sandbox.api.sap.com/sap/c4c/odata/v1/c4codataapi/PartnerCollection('00163E038C701ED298AFFC9A58CB1677')/PartnerAddress",
			"PartnerHasPartnerContact": "https://sandbox.api.sap.com/sap/c4c/odata/v1/c4codataapi/PartnerCollection('00163E038C701ED298AFFC9A58CB1677')/PartnerHasPartnerContact"
		}
	],
	"time": "2022-05-06T21:32:17+09:00"
}

```