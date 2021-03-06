package models

type Location struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

type AboutInfo struct {
	Name     string            `json:"name"`
	Location Location          `json:"location"`
	Details  string            `json:"details"`
	Links    map[string]string `json:"links"`
	Email    string            `json:"email"`
}

func GetAboutInfo() AboutInfo {
	return AboutInfo{
		Name:     "Richard Robinson",
		Location: Location{City: "Toronto", Country: "Canada"},
		Details:  "I'm a full-stack software engineer in Toronto, and currently in my fourth year at University",
		Links: map[string]string{
			"twitter":  "https://twitter.com/rr_codes",
			"github":   "https://github.com/rr-codes",
			"linkedin": "https://linkedin.com/in/rr-codes",
		},
		Email: "richard@rr.codes",
	}
}
