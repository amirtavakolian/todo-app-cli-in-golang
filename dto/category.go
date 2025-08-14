package dto

type Category struct {
	Id        int       `json:"id,omitempty"`
	Name      string    `json:"name"`
	Status    bool      `json:"status"`
}
