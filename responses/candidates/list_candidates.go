package candidates

type ListCandidatesResponse struct {
	Meta Meta    `json:"meta"`
	Data []Data1 `json:"data"`
}

type Totals struct {
	Rows  int   `json:"rows"`
	Types []any `json:"types"`
}

type Meta struct {
	Totals            Totals `json:"totals"`
	Solr              bool   `json:"solr"`
	Version           string `json:"version"`
	AndroidMinVersion string `json:"androidMinVersion"`
	IosMinVersion     string `json:"iosMinVersion"`
	IsLogged          bool   `json:"isLogged"`
	Language          string `json:"language"`
	Timestamp         int64  `json:"timestamp"`
	Customer          string `json:"customer"`
}

type Source struct {
	TypeOf int    `json:"typeOf"`
	Detail string `json:"detail"`
}
type SocialNetworks struct {
	Network string `json:"network"`
	URL     string `json:"url"`
}
type Attributes struct {
	CreationDate               string           `json:"creationDate"`
	UpdateDate                 string           `json:"updateDate"`
	Civility                   int              `json:"civility"`
	FirstName                  string           `json:"firstName"`
	LastName                   string           `json:"lastName"`
	State                      int              `json:"state"`
	TypeOf                     int              `json:"typeOf"`
	IsVisible                  bool             `json:"isVisible"`
	Thumbnail                  string           `json:"thumbnail"`
	Availability               int              `json:"availability"`
	Skills                     string           `json:"skills"`
	Diplomas                   []any            `json:"diplomas"`
	MobilityAreas              []string         `json:"mobilityAreas"`
	ActivityAreas              []any            `json:"activityAreas"`
	GlobalEvaluation           string           `json:"globalEvaluation"`
	Languages                  []any            `json:"languages"`
	ExpertiseAreas             []any            `json:"expertiseAreas"`
	Experience                 int              `json:"experience"`
	References                 []any            `json:"references"`
	Evaluations                []any            `json:"evaluations"`
	Tools                      []any            `json:"tools"`
	Title                      string           `json:"title"`
	Email1                     string           `json:"email1"`
	Email2                     string           `json:"email2"`
	Email3                     string           `json:"email3"`
	Phone1                     string           `json:"phone1"`
	Phone2                     string           `json:"phone2"`
	Town                       string           `json:"town"`
	Country                    string           `json:"country"`
	Source                     Source           `json:"source"`
	NumberOfResumes            int              `json:"numberOfResumes"`
	NumberOfActivePositionings int              `json:"numberOfActivePositionings"`
	SocialNetworks             []SocialNetworks `json:"socialNetworks"`
	CanShowTechnicalData       bool             `json:"canShowTechnicalData"`
	CanShowActions             bool             `json:"canShowActions"`
}
type Data2 struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}
type MainManager struct {
	Data Data2 `json:"data"`
}
type Agency struct {
	Data Data2 `json:"data"`
}
type Pole struct {
	Data any `json:"data"`
}
type Relationships struct {
	MainManager MainManager `json:"mainManager"`
	Agency      Agency      `json:"agency"`
	Pole        Pole        `json:"pole"`
}
type Data1 struct {
	ID            string        `json:"id"`
	Type          string        `json:"type"`
	Attributes    Attributes    `json:"attributes"`
	Relationships Relationships `json:"relationships"`
}
