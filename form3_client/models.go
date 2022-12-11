//go:build darwin || linux || windows
// +build darwin linux windows

package form3_client

type AccountData struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	ID             string             `json:"id,omitempty"`
	OrganisationID string             `json:"organisation_id,omitempty"`
	Type           string             `json:"type,omitempty"`
	Version        *int64             `json:"version,omitempty"`
}

type AccountAttributes struct {
	AccountClassification   string   `json:"account_classification,omitempty"`
	AccountMatchingOptOut   bool     `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	AlternativeNames        []string `json:"alternative_names,omitempty"`
	BankID                  string   `json:"bank_id,omitempty"`
	BankIDCode              string   `json:"bank_id_code,omitempty"`
	BaseCurrency            string   `json:"base_currency,omitempty"`
	Bic                     string   `json:"bic,omitempty"`
	Country                 string   `json:"country,omitempty"`
	Iban                    string   `json:"iban,omitempty"`
	JointAccount            bool     `json:"joint_account,omitempty"`
	Name                    []string `json:"name,omitempty"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	Status                  string   `json:"status,omitempty"`
	Switched                bool     `json:"switched,omitempty"`
	CustomerID              string   `json:"customer_id,omitempty"`
}

type AccountJSON struct {
	Data  AccountData `json:"data,omitempty"`
	Links Link        `json:"links,omitempty"`
}

type Link struct {
	First string `json:"first,omitempty"`
	Last  string `json:"last,omitempty"`
	Self  string `json:"self,omitempty"`
}

type ResponseJSON struct {
	Data         []AccountData `json:"data,omitempty"`
	Status       string        `json:"status,omitempty"`
	StatusCode   int           `json:"status_code,omitempty"`
	ErrorCode    int           `json:"error_code,omitempty"`
	ErrorMessage string        `json:"error,omitempty"`
}
