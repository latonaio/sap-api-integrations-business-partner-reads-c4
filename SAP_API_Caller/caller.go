package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-business-partner-reads-c4/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

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

func (c *SAPAPICaller) PartnerCollection(partnerID string) {
	partnerCollectionData, err := c.callBPSrvAPIRequirementPartnerCollection("PartnerCollection", partnerID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(partnerCollectionData)

	partnerAddressData, err := c.callParterAddress(partnerCollectionData[0].ToPartnerAddress)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(partnerAddressData)

	partnerHasPartnerContactData, err := c.callParterHasPartnerContact(partnerCollectionData[0].ToPartnerHasPartnerContact)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(partnerHasPartnerContactData)

}

func (c *SAPAPICaller) callBPSrvAPIRequirementPartnerCollection(api, partnerID string) ([]sap_api_output_formatter.PartnerCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithPartnerCollection(req, partnerID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToPartnerCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callParterAddress(url string) ([]sap_api_output_formatter.PartnerAddress, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToPartnerAddress(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callParterHasPartnerContact(url string) ([]sap_api_output_formatter.PartnerHasPartnerContact, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToPartnerHasPartnerContact(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ContactCollection(contactID string) {
	contactCollectionData, err := c.callBPSrvAPIRequirementContactCollection("ContactCollection", contactID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(contactCollectionData)

	contactIsContactPersonForData, err := c.callContactIsContactPersonFor(contactCollectionData[0].ToContactIsContactPersonFor)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(contactIsContactPersonForData)

	corporateAccountData, err := c.callCorporateAccount(contactCollectionData[0].ToCorporateAccount)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(corporateAccountData)

}

func (c *SAPAPICaller) callBPSrvAPIRequirementContactCollection(api, contactID string) ([]sap_api_output_formatter.ContactCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithContactCollection(req, contactID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToContactCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callContactIsContactPersonFor(url string) ([]sap_api_output_formatter.ContactIsContactPersonFor, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToContactIsContactPersonFor(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callCorporateAccount(url string) (*sap_api_output_formatter.ToCorporateAccount, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToCorporateAccount(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) BusinessUserCollection(userID string) {
	businessUserCollectionData, err := c.callBPSrvAPIRequirementBusinessUserCollection("BusinessUserCollection", userID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(businessUserCollectionData)

	businessRoleAssignmentData, err := c.callBusinessRoleAssignment(businessUserCollectionData[0].ToBusinessUserBusinessRoleAssignment)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(businessRoleAssignmentData)

	employeeBasicDataData, err := c.callEmployeeBasicData(businessUserCollectionData[0].ToEmployeeBasicData)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(employeeBasicDataData)

}

func (c *SAPAPICaller) callBPSrvAPIRequirementBusinessUserCollection(api, userID string) ([]sap_api_output_formatter.BusinessUserCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithBusinessUserCollection(req, userID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToBusinessUserCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callBusinessRoleAssignment(url string) ([]sap_api_output_formatter.BusinessUserBusinessRoleAssignment, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToBusinessUserBusinessRoleAssignment(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callEmployeeBasicData(url string) (*sap_api_output_formatter.EmployeeBasicData, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToEmployeeBasicData(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) IndividualCustomerCollection(customerID string) {
	individualCustomerCollectionData, err := c.callBPSrvAPIRequirementIndividualCustomerCollection("IndividualCustomerCollection", customerID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(individualCustomerCollectionData)

	individualCustomerAddressData, err := c.callIndividualCustomerAddress(individualCustomerCollectionData[0].ToIndividualCustomerAddress)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(individualCustomerAddressData)

}

func (c *SAPAPICaller) callBPSrvAPIRequirementIndividualCustomerCollection(api, customerID string) ([]sap_api_output_formatter.IndividualCustomerCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithIndividualCustomerCollection(req, customerID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToIndividualCustomerCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callIndividualCustomerAddress(url string) ([]sap_api_output_formatter.IndividualCustomerAddress, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToIndividualCustomerAddress(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) CorporateAccount(accountID string) {
	data, err := c.callBPSrvAPIRequirementCorporateAccount("CorporateAccountCollection", accountID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBPSrvAPIRequirementCorporateAccount(api, accountID string) ([]sap_api_output_formatter.CorporateAccount, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithCorporateAccount(req, accountID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToCorporateAccount(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ServiceAgentCollection(serviceAgentID string) {
	data, err := c.callBPSrvAPIRequirementServiceAgentCollection("ServiceAgentCollection", serviceAgentID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBPSrvAPIRequirementServiceAgentCollection(api, serviceAgentID string) ([]sap_api_output_formatter.ServiceAgentCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithServiceAgentCollection(req, serviceAgentID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToServiceAgentCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) CompetitorCollection(competitorID string) {
	data, err := c.callBPSrvAPIRequirementCompetitorCollection("CompetitorCollection", competitorID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBPSrvAPIRequirementCompetitorCollection(api, competitorID string) ([]sap_api_output_formatter.CompetitorCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithCompetitorCollection(req, competitorID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToCompetitorCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) EmployeeCollection(employeeID string) {
	employeeCollectionData, err := c.callBPSrvAPIRequirementEmployeeCollection("EmployeeCollection", employeeID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(employeeCollectionData)

	employeeOrganisationalUnitAssignmentData, err := c.callEmployeeOrganisationalUnitAssignment(employeeCollectionData[0].ToEmployeeOrganisationalUnitAssignment)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(employeeOrganisationalUnitAssignmentData)

	employeeSalesResponsibilityData, err := c.callEmployeeSalesResponsibility(employeeCollectionData[0].ToEmployeeSalesResponsibility)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(employeeSalesResponsibilityData)

	employeeUserBusinessRoleAssignmentData, err := c.callEmployeeUserBusinessRoleAssignment(employeeCollectionData[0].ToEmployeeUserBusinessRoleAssignment)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(employeeUserBusinessRoleAssignmentData)

	employeeWorkingHoursData, err := c.callEmployeeWorkingHours(employeeCollectionData[0].ToEmployeeWorkingHours)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(employeeWorkingHoursData)

}

func (c *SAPAPICaller) callBPSrvAPIRequirementEmployeeCollection(api, employeeID string) ([]sap_api_output_formatter.EmployeeCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithEmployeeCollection(req, employeeID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToEmployeeCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callEmployeeOrganisationalUnitAssignment(url string) ([]sap_api_output_formatter.EmployeeOrganisationalUnitAssignment, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToEmployeeOrganisationalUnitAssignment(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callEmployeeSalesResponsibility(url string) ([]sap_api_output_formatter.EmployeeSalesResponsibility, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToEmployeeSalesResponsibility(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callEmployeeUserBusinessRoleAssignment(url string) ([]sap_api_output_formatter.EmployeeUserBusinessRoleAssignment, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToEmployeeUserBusinessRoleAssignment(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callEmployeeWorkingHours(url string) ([]sap_api_output_formatter.EmployeeWorkingHours, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToEmployeeWorkingHours(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithPartnerCollection(req *http.Request, partnerID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("PartnerID eq '%s'", partnerID))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithContactCollection(req *http.Request, contactID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("ContactID eq '%s'", contactID))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithBusinessUserCollection(req *http.Request, userID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("UserID eq '%s'", userID))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithIndividualCustomerCollection(req *http.Request, customerID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("CustomerID eq '%s'", customerID))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithCorporateAccount(req *http.Request, accountID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("AccountID eq '%s'", accountID))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithServiceAgentCollection(req *http.Request, serviceAgentID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("ServiceAgentID eq '%s'", serviceAgentID))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithCompetitorCollection(req *http.Request, competitorID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("CompetitorID eq '%s'", competitorID))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithEmployeeCollection(req *http.Request, employeeID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("EmployeeID eq '%s'", employeeID))
	req.URL.RawQuery = params.Encode()
}
