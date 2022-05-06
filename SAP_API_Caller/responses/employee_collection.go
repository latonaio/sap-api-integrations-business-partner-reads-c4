package responses

type EmployeeCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                             string `json:"ObjectID"`
			EmployeeID                           string `json:"EmployeeID"`
			UUID                                 string `json:"UUID"`
			UserID                               string `json:"UserID"`
			IdentityUUID                         string `json:"IdentityUUID"`
			GlobalUserID                         string `json:"GlobalUserID"`
			BusinessPartnerID                    string `json:"BusinessPartnerID"`
			EmployeeValidityStartDate            string `json:"EmployeeValidityStartDate"`
			EmployeeValidityEndDate              string `json:"EmployeeValidityEndDate"`
			BusinessPartnerFormattedName         string `json:"BusinessPartnerFormattedName"`
			TitleCode                            string `json:"TitleCode"`
			TitleCodeText                        string `json:"TitleCodeText"`
			AcademicTitleCode                    string `json:"AcademicTitleCode"`
			AcademicTitleCodeText                string `json:"AcademicTitleCodeText"`
			FirstName                            string `json:"FirstName"`
			MiddleName                           string `json:"MiddleName"`
			LastName                             string `json:"LastName"`
			SecondLastName                       string `json:"SecondLastName"`
			NickName                             string `json:"NickName"`
			GenderCode                           string `json:"GenderCode"`
			GenderCodeText                       string `json:"GenderCodeText"`
			MaritalStatusCode                    string `json:"MaritalStatusCode"`
			MaritalStatusCodeText                string `json:"MaritalStatusCodeText"`
			LanguageCode                         string `json:"LanguageCode"`
			LanguageCodeText                     string `json:"LanguageCodeText"`
			NationalityCountryCode               string `json:"NationalityCountryCode"`
			NationalityCountryCodeText           string `json:"NationalityCountryCodeText"`
			BirthName                            string `json:"BirthName"`
			BirthDate                            string `json:"BirthDate"`
			BirthPlace                           string `json:"BirthPlace"`
			FormattedPostalAddressDescription    string `json:"FormattedPostalAddressDescription"`
			CountryCode                          string `json:"CountryCode"`
			CountryCodeText                      string `json:"CountryCodeText"`
			RegionCode                           string `json:"RegionCode"`
			RegionCodeText                       string `json:"RegionCodeText"`
			CareOfName                           string `json:"CareOfName"`
			AddressLine1                         string `json:"AddressLine1"`
			AddressLine2                         string `json:"AddressLine2"`
			HouseNumber                          string `json:"HouseNumber"`
			Street                               string `json:"Street"`
			AddressLine4                         string `json:"AddressLine4"`
			AddressLine5                         string `json:"AddressLine5"`
			District                             string `json:"District"`
			City                                 string `json:"City"`
			DifferentCity                        string `json:"DifferentCity"`
			PostalCode                           string `json:"PostalCode"`
			County                               string `json:"County"`
			CompanyPostalCode                    string `json:"CompanyPostalCode"`
			POBox                                string `json:"POBox"`
			POBoxPostalCode                      string `json:"POBoxPostalCode"`
			POBoxCountryCode                     string `json:"POBoxCountryCode"`
			POBoxCountryCodeText                 string `json:"POBoxCountryCodeText"`
			POBoxRegionCode                      string `json:"POBoxRegionCode"`
			POBoxRegionCodeText                  string `json:"POBoxRegionCodeText"`
			POBoxCity                            string `json:"POBoxCity"`
			TimeZoneCode                         string `json:"TimeZoneCode"`
			TimeZoneCodeText                     string `json:"TimeZoneCodeText"`
			TaxJurisdictionCode                  string `json:"TaxJurisdictionCode"`
			TaxJurisdictionCodeText              string `json:"TaxJurisdictionCodeText"`
			Building                             string `json:"Building"`
			Floor                                string `json:"Floor"`
			Room                                 string `json:"Room"`
			InhouseMail                          string `json:"InhouseMail"`
			OfficePhoneNumber                    string `json:"OfficePhoneNumber"`
			NormalisedOfficePhoneNumber          string `json:"NormalisedOfficePhoneNumber"`
			MobilePhoneNumber                    string `json:"MobilePhoneNumber"`
			NormalisedMobilePhoneNumber          string `json:"NormalisedMobilePhoneNumber"`
			FaxNumber                            string `json:"FaxNumber"`
			Email                                string `json:"Email"`
			UserValidityStartDate                string `json:"UserValidityStartDate"`
			UserValidityEndDate                  string `json:"UserValidityEndDate"`
			UserPasswordPolicyCode               string `json:"UserPasswordPolicyCode"`
			UserPasswordPolicyCodeText           string `json:"UserPasswordPolicyCodeText"`
			UserLockedIndicator                  bool   `json:"UserLockedIndicator"`
			UserCountedIndicator                 bool   `json:"UserCountedIndicator"`
			CreatedOn                            string `json:"CreatedOn"`
			CreatedBy                            string `json:"CreatedBy"`
			CreatedByIdentityUUID                string `json:"CreatedByIdentityUUID"`
			ChangedOn                            string `json:"ChangedOn"`
			ChangedBy                            string `json:"ChangedBy"`
			ChangedByIdentityUUID                string `json:"ChangedByIdentityUUID"`
			EntityLastChangedOn                  string `json:"EntityLastChangedOn"`
			ETag                                 string `json:"ETag"`
			EmployeeOrganisationalUnitAssignment struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"EmployeeOrganisationalUnitAssignment"`
			EmployeeSalesResponsibility struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"EmployeeSalesResponsibility"`
			EmployeeUserBusinessRoleAssignment struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"EmployeeUserBusinessRoleAssignment"`
			EmployeeWorkingHours struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"EmployeeWorkingHours"`
		} `json:"results"`
	} `json:"d"`
}
