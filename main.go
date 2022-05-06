package main

import (
	sap_api_caller "sap-api-integrations-business-partner-reads-c4/SAP_API_Caller"
	sap_api_input_reader "sap-api-integrations-business-partner-reads-c4/SAP_API_Input_Reader"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs//SDC_Business_Partner_Employee_Collection_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/sap/c4c/odata/v1/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"PartnerCollection", "ContactCollection", "BusinessUserCollection",
			"IndividualCustomerCollection", "CorporateAccount", "ServiceAgentCollection", "CompetitorCollection",
			"EmployeeCollection",
		}
	}

	caller.AsyncGetBP(
		inoutSDC.PartnerCollection.PartnerID,
		inoutSDC.ContactCollection.ContactID,
		inoutSDC.BusinessUserCollection.UserID,
		inoutSDC.ContactCollection.ContactIsContactPersonFor.CorporateAccount.IndividualCustomerCollection.CustomerID,
		inoutSDC.ContactCollection.ContactIsContactPersonFor.CorporateAccount.AccountID,
		inoutSDC.ServiceAgentCollection.ServiceAgentID,
		inoutSDC.CompetitorCollection.CompetitorID,
		inoutSDC.EmployeeCollection.EmployeeID,
		accepter,
	)
}
