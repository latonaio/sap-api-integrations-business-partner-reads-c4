package responses

type ToEmployeeUserBusinessRoleAssignment struct {
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
			UserID              string `json:"UserID"`
			BusinessRoleID      string `json:"BusinessRoleID"`
			EntityLastChangedOn string `json:"EntityLastChangedOn"`
			ETag                string `json:"ETag"`
		} `json:"results"`
	} `json:"d"`
}
