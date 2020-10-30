package models

type Organization struct {
	Id         int      `json:"_id"`
	Url        string   `json:"url"`
	ExternalId string   `json:"external_id"`
	CreatedAt  string   `json:"created_at"`
	Tags       []string `json:"tags"`

	Name          string   `json:"name"`
	DomainNames   []string `json:"domain_names"`
	Details       string   `json:"details"`
	SharedTickets bool     `json:"shared_tickets"`
}
