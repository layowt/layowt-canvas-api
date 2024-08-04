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
	ExperienceLevel       string `json:"experienceLevel"`
}

type Website struct {
	WebsiteId              string `json:"websiteId"`
	UserId                 int    `json:"userId"`
	WebsiteName            string `json:"websiteName"`
	WebsiteUrl             string `json:"websiteUrl"`
	WebsiteLogo            string `json:"websiteLogo"`
	WebsitePrimaryColor    string `json:"websitePrimaryColor"`
	WebsiteSecondaryColor  string `json:"websiteSecondaryColor"`
	WebsiteBackgroundColor string `json:"websiteBackgroundColor"`
	HasBeenPublished       bool   `json:"hasBeenPublished"`
	CreatedAt              string `json:"createdAt"`
	LastUpdated            string `json:"lastUpdated"`
	LastUpdatedUid         string `json:"lastUpdatedUid"`
}
