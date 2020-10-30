package models

type User struct {
	Id         int      `json:"_id"`
	Url        string   `json:"url"`
	ExternalId string   `json:"external_id"`
	CreatedAt  string   `json:"created_at"`
	Tags       []string `json:"tags"`

	OrganizationId int    `json:"organization_id"`
	Name           string `json:"name"`
	Alias          string `json:"alias"`
	Locale         string `json:"locale"`
	Timezone       string `json:"timezone"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Signature      string `json:"signature"`
	Role           string `json:"role"`
	LastLoginAt    string `json:"last_login_at"`
	Suspended      bool   `json:"suspended"`
	Active         bool   `json:"active"`
	Verified       bool   `json:"verified"`
	Shared         bool   `json:"shared"`
}
