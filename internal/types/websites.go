package types

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
