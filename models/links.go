package models

type Links struct {
	ID         int    `json:"id,omitempty"`
	Alias      string `json:"alias"`
	Url        string `json:"url"`
	Created_at string `json:"createdAt,omitempty"`
	Updated_at string `json:"updatedAt,omitempty"`
}
