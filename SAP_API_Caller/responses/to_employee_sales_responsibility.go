package responses

type ToEmployeeSalesResponsibility struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                    string `json:"ObjectID"`
			ParentObjectID              string `json:"ParentObjectID"`
			EmployeeID                  string `json:"EmployeeID"`
			SalesOrganisationID         string `json:"SalesOrganisationID"`
			DistributionChannelCode     string `json:"DistributionChannelCode"`
			DistributionChannelCodeText string `json:"DistributionChannelCodeText"`
			DivisionCode                string `json:"DivisionCode"`
			DivisionCodeText            string `json:"DivisionCodeText"`
			MainIndicator               bool   `json:"MainIndicator"`
			ETag                        string `json:"ETag"`
		} `json:"results"`
	} `json:"d"`
}
