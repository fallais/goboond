package candidates

type ListActionsResponse struct {
	Meta struct {
		Totals struct {
			Rows  int   `json:"rows"`
			Types []any `json:"types"`
		} `json:"totals"`
		Version           string `json:"version"`
		AndroidMinVersion string `json:"androidMinVersion"`
		IosMinVersion     string `json:"iosMinVersion"`
		IsLogged          bool   `json:"isLogged"`
		Language          string `json:"language"`
		Timestamp         int64  `json:"timestamp"`
		Customer          string `json:"customer"`
	} `json:"meta"`
	Data []struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			StartDate      string `json:"startDate"`
			CreationDate   string `json:"creationDate"`
			TypeOf         int    `json:"typeOf"`
			Text           string `json:"text"`
			NumberOfFiles  int    `json:"numberOfFiles"`
			CanReadAction  bool   `json:"canReadAction"`
			CanWriteAction bool   `json:"canWriteAction"`
		} `json:"attributes"`
		Relationships struct {
			MainManager struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"mainManager"`
			DependsOn struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"dependsOn"`
			RelatedActions struct {
				Data []any `json:"data"`
			} `json:"relatedActions"`
		} `json:"relationships"`
	} `json:"data"`
	Included []struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			LastName  string `json:"lastName"`
			FirstName string `json:"firstName"`
		} `json:"attributes"`
	} `json:"included"`
}
