package candidates

type GetInformationResponse struct {
	Meta struct {
		Version           string `json:"version"`
		AndroidMinVersion string `json:"androidMinVersion"`
		IosMinVersion     string `json:"iosMinVersion"`
		IsLogged          bool   `json:"isLogged"`
		Language          string `json:"language"`
		Timestamp         int64  `json:"timestamp"`
		Customer          string `json:"customer"`
	} `json:"meta"`
	Data struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			CreationDate string `json:"creationDate"`
			UpdateDate   string `json:"updateDate"`
			Civility     int    `json:"civility"`
			LastName     string `json:"lastName"`
			FirstName    string `json:"firstName"`
			Title        string `json:"title"`
			TypeOf       int    `json:"typeOf"`
			State        int    `json:"state"`
			StateReason  struct {
				TypeOf int    `json:"typeOf"`
				Detail string `json:"detail"`
			} `json:"stateReason"`
			Email1      string `json:"email1"`
			Email2      string `json:"email2"`
			Email3      string `json:"email3"`
			Phone1      string `json:"phone1"`
			Phone2      string `json:"phone2"`
			Phone3      string `json:"phone3"`
			Fax         string `json:"fax"`
			Address     string `json:"address"`
			Postcode    string `json:"postcode"`
			Town        string `json:"town"`
			Country     string `json:"country"`
			SubDivision string `json:"subDivision"`
			Source      struct {
				TypeOf int    `json:"typeOf"`
				Detail string `json:"detail"`
			} `json:"source"`
			DateOfBirth                string `json:"dateOfBirth"`
			MobilityAreas              []any  `json:"mobilityAreas"`
			GlobalEvaluation           string `json:"globalEvaluation"`
			Evaluations                []any  `json:"evaluations"`
			Availability               int    `json:"availability"`
			Thumbnail                  string `json:"thumbnail"`
			NumberOfActivePositionings int    `json:"numberOfActivePositionings"`
			InformationComments        string `json:"informationComments"`
			SocialNetworks             []struct {
				Network string `json:"network"`
				URL     string `json:"url"`
			} `json:"socialNetworks"`
		} `json:"attributes"`
		Relationships struct {
			MainManager struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"mainManager"`
			HrManager struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"hrManager"`
			CreatedBy struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"createdBy"`
			Agency struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"agency"`
			Pole struct {
				Data any `json:"data"`
			} `json:"pole"`
			Resumes struct {
				Data []struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"resumes"`
			Resource struct {
				Data any `json:"data"`
			} `json:"resource"`
		} `json:"relationships"`
	} `json:"data"`
	Included []struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			LastName  string `json:"lastName"`
			FirstName string `json:"firstName"`
			Thumbnail string `json:"thumbnail"`
		} `json:"attributes,omitempty"`
		Attributes0 struct {
			Name string `json:"name"`
		} `json:"attributes,omitempty"`
		Attributes1 struct {
			Name string `json:"name"`
		} `json:"attributes,omitempty"`
	} `json:"included"`
}
