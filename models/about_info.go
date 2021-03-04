package models

type Location struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

type AboutInfo struct {
	Name     string            `json:"name"`
	Location Location          `json:"location"`
	Links    map[string]string `json:"links"`
	Details  string            `json:"details"`
}

func GetAboutInfo() AboutInfo {
	return AboutInfo{
		Name:     "Richard Robinson",
		Location: Location{City: "Toronto", Country: "Canada"},
		Links: map[string]string{
			"twitter":  "twitter.com/rr_codes",
			"github":   "github.com/rr_codes",
			"linkedin": "",
		},
		Details: "I'm a full-stack software engineer in Toronto, and currently in my fourth year at University",
	}
}
