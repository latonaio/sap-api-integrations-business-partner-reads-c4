package responses

type ToEmployeeOrganisationalUnitAssignment struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID            string `json:"ObjectID"`
			ParentObjectID      string `json:"ParentObjectID"`
			EmployeeID          string `json:"EmployeeID"`
			OrgUnitID           string `json:"OrgUnitID"`
			RoleCode            string `json:"RoleCode"`
			RoleCodeText        string `json:"RoleCodeText"`
			StartDate           string `json:"StartDate"`
			EndDate             string `json:"EndDate"`
			JobID               string `json:"JobID"`
			EntityLastChangedOn string `json:"EntityLastChangedOn"`
			ETag                string `json:"ETag"`
		} `json:"results"`
	} `json:"d"`
}
