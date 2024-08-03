package types

type User struct {
	Uid                   int    `json:"uid"`
	Email                 string `json:"email"`
	HasAuthenticatedEmail bool   `json:"hasAuthenticatedEmail"`
	FirstName             string `json:"firstName"`
	LastName              string `json:"lastName"`
	DisplayName           string `json:"displayName"`
	CreatedAt             string `json:"createdAt"`
	UpdatedAt             string `json:"updatedAt"`
	LastLogin             string `json:"lastLogin"`
	ExperienceLevel       string `json:"expierenceLevel"`
}
