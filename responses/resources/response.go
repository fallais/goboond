package resources

import "encoding/json"

// ResourcesPaginatedResponse is the paginated response.
type ListResourcesResponse struct {
	Meta Meta   `json:"meta"`
	Data []Data `json:"data"`
}
type Totals struct {
	Rows int `json:"rows"`
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
type StateReason struct {
	TypeOf int    `json:"typeOf"`
	Detail string `json:"detail"`
}
type Attributes struct {
	Civility                      int         `json:"civility"`
	FirstName                     string      `json:"firstName"`
	LastName                      string      `json:"lastName"`
	CreationDate                  string      `json:"creationDate"`
	Reference                     string      `json:"reference"`
	TypeOf                        int         `json:"typeOf"`
	State                         int         `json:"state"`
	StateReason                   StateReason `json:"stateReason"`
	IsVisible                     bool        `json:"isVisible"`
	Thumbnail                     string      `json:"thumbnail"`
	Skills                        string      `json:"skills"`
	MobilityAreas                 []string    `json:"mobilityAreas"`
	Title                         string      `json:"title"`
	Availability                  json.Number `json:"availability"`
	AverageDailyPriceExcludingTax int         `json:"averageDailyPriceExcludingTax"`
	Email1                        string      `json:"email1"`
	Email2                        string      `json:"email2"`
	Email3                        string      `json:"email3"`
	Phone1                        string      `json:"phone1"`
	Phone2                        string      `json:"phone2"`
	Currency                      int         `json:"currency"`
	ExchangeRate                  int         `json:"exchangeRate"`
	CurrencyAgency                int         `json:"currencyAgency"`
	ExchangeRateAgency            int         `json:"exchangeRateAgency"`
	NumberOfResumes               int         `json:"numberOfResumes"`
	NumberOfActivePositionings    int         `json:"numberOfActivePositionings"`
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
type Data struct {
	ID            string        `json:"id"`
	Type          string        `json:"type"`
	Attributes    Attributes    `json:"attributes"`
	Relationships Relationships `json:"relationships"`
}

type GetResourceResponse struct {
	Meta     MetaRR       `json:"meta"`
	Data     DataRR       `json:"data"`
	Included []IncludedRR `json:"included"`
}
type MetaRR struct {
	Version           string `json:"version"`
	AndroidMinVersion string `json:"androidMinVersion"`
	IosMinVersion     string `json:"iosMinVersion"`
	IsLogged          bool   `json:"isLogged"`
	Language          string `json:"language"`
	Timestamp         int64  `json:"timestamp"`
	Customer          string `json:"customer"`
}
type AttributesRR struct {
	CreationDate          string `json:"creationDate"`
	UpdateDate            string `json:"updateDate"`
	Civility              int    `json:"civility"`
	LastName              string `json:"lastName"`
	FirstName             string `json:"firstName"`
	Thumbnail             string `json:"thumbnail"`
	TypeOf                int    `json:"typeOf"`
	Level                 string `json:"level"`
	Title                 string `json:"title"`
	DateOfBirth           string `json:"dateOfBirth"`
	NumberOfResumes       int    `json:"numberOfResumes"`
	SeniorityDate         string `json:"seniorityDate"`
	OriginalSeniorityDate string `json:"originalSeniorityDate"`
	ForceSeniorityDate    bool   `json:"forceSeniorityDate"`
	ValiditySeniorityDate string `json:"validitySeniorityDate"`
}
type DataRR2 struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}
type MainManagerRR struct {
	Data DataRR2 `json:"data"`
}
type HrManagerRR struct {
	Data DataRR2 `json:"data"`
}
type AgencyRR struct {
	Data DataRR2 `json:"data"`
}
type PoleRR struct {
	Data any `json:"data"`
}
type RelationshipsRR struct {
	MainManager MainManager `json:"mainManager"`
	HrManager   HrManagerRR `json:"hrManager"`
	Agency      Agency      `json:"agency"`
	Pole        Pole        `json:"pole"`
}
type DataRR struct {
	ID            string          `json:"id"`
	Type          string          `json:"type"`
	Attributes    AttributesRR    `json:"attributes"`
	Relationships RelationshipsRR `json:"relationships"`
}
type AttributesRR2 struct {
	Calendar string `json:"calendar"`
}
type IncludedRR struct {
	ID         string        `json:"id"`
	Type       string        `json:"type"`
	Attributes AttributesRR2 `json:"attributes"`
}
