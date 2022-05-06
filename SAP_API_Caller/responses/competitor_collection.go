package responses

type CompetitorCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                          string `json:"ObjectID"`
			CompetitorID                      string `json:"CompetitorID"`
			CompetitorUUID                    string `json:"CompetitorUUID"`
			StatusCode                        string `json:"StatusCode"`
			StatusCodeText                    string `json:"StatusCodeText"`
			ClassificationCode                string `json:"ClassificationCode"`
			ClassificationCodeText            string `json:"ClassificationCodeText"`
			BusinessPartnerFormattedName      string `json:"BusinessPartnerFormattedName"`
			Name                              string `json:"Name"`
			AdditionalName                    string `json:"AdditionalName"`
			FormattedPostalAddressDescription string `json:"FormattedPostalAddressDescription"`
			CountryCode                       string `json:"CountryCode"`
			CountryCodeText                   string `json:"CountryCodeText"`
			RegionCode                        string `json:"RegionCode"`
			RegionCodeText                    string `json:"RegionCodeText"`
			CareOfName                        string `json:"CareOfName"`
			AddressLine1                      string `json:"AddressLine1"`
			AddressLine2                      string `json:"AddressLine2"`
			HouseNumber                       string `json:"HouseNumber"`
			Street                            string `json:"Street"`
			AddressLine4                      string `json:"AddressLine4"`
			AddressLine5                      string `json:"AddressLine5"`
			City                              string `json:"City"`
			AdditionalCityName                string `json:"AdditionalCityName"`
			District                          string `json:"District"`
			County                            string `json:"County"`
			CompanyPostalCode                 string `json:"CompanyPostalCode"`
			StreetPostalCode                  string `json:"StreetPostalCode"`
			POBoxPostalCode                   string `json:"POBoxPostalCode"`
			POBox                             string `json:"POBox"`
			POBoxDeviatingCountryCode         string `json:"POBoxDeviatingCountryCode"`
			POBoxDeviatingCountryCodeText     string `json:"POBoxDeviatingCountryCodeText"`
			POBoxDeviatingCity                string `json:"POBoxDeviatingCity"`
			TimeZoneCode                      string `json:"TimeZoneCode"`
			TimeZoneCodeText                  string `json:"TimeZoneCodeText"`
			TaxJurisdictionCode               string `json:"TaxJurisdictionCode"`
			TaxJurisdictionCodeText           string `json:"TaxJurisdictionCodeText"`
			POBoxDeviatingStateCode           string `json:"POBoxDeviatingStateCode"`
			POBoxDeviatingStateCodeText       string `json:"POBoxDeviatingStateCodeText"`
			Phone                             string `json:"Phone"`
			Fax                               string `json:"Fax"`
			Email                             string `json:"Email"`
			WebSite                           string `json:"WebSite"`
			LanguageCode                      string `json:"LanguageCode"`
			LanguageCodeText                  string `json:"LanguageCodeText"`
			BestReachedByCode                 string `json:"BestReachedByCode"`
			BestReachedByCodeText             string `json:"BestReachedByCodeText"`
			NormalisedPhone                   string `json:"NormalisedPhone"`
			CreatedOn                         string `json:"CreatedOn"`
			CreatedBy                         string `json:"CreatedBy"`
			CreatedByIdentityUUID             string `json:"CreatedByIdentityUUID"`
			ChangedOn                         string `json:"ChangedOn"`
			ChangedBy                         string `json:"ChangedBy"`
			ChangedByIdentityUUID             string `json:"ChangedByIdentityUUID"`
			EntityLastChangedOn               string `json:"EntityLastChangedOn"`
			ETag                              string `json:"ETag"`
		} `json:"results"`
	} `json:"d"`
}
