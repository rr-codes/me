package models

type Location struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

type AboutInfo struct {
	Name      string            `json:"name"`
	Location  Location          `json:"location"`
	Details   string            `json:"details"`
	Education []string          `json:"education"`
	Links     map[string]string `json:"links"`
	Email     string            `json:"email"`
}

func GetAboutInfo() AboutInfo {
	return AboutInfo{
		Name:     "Richard Robinson",
		Location: Location{City: "Cupertino", Country: "United States"},
		Details:  "I'm an incoming WebKit Platform Engineer at Apple",
		Education: []string{
			"Bachelor of Engineering (B.Eng.), Spec. Hons. Software Engineering",
		},
		Links: map[string]string{
			"twitter":  "https://twitter.com/rr_codes",
			"github":   "https://github.com/rr-codes",
			"linkedin": "https://linkedin.com/in/rr-codes",
		},
		Email: "richard@rr.codes",
	}
}
