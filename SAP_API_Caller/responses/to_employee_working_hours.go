package responses

type ToEmployeeWorkingHours struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                   string `json:"ObjectID"`
			ParentObjectID             string `json:"ParentObjectID"`
			EmployeeID                 string `json:"EmployeeID"`
			WorkingHoursType           string `json:"WorkingHoursType"`
			WorkingHoursTypeText       string `json:"WorkingHoursTypeText"`
			StartDate                  string `json:"StartDate"`
			EndDate                    string `json:"EndDate"`
			TimeZoneCode               string `json:"TimeZoneCode"`
			TimeZoneCodeText           string `json:"TimeZoneCodeText"`
			WorkingDayCalendarCode     string `json:"WorkingDayCalendarCode"`
			WorkingDayCalendarCodeText string `json:"WorkingDayCalendarCodeText"`
			ETag                       string `json:"ETag"`
		} `json:"results"`
	} `json:"d"`
}
