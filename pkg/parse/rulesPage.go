package parse

type RulesPage struct {
	Status string `json:"status"`
	Data   struct {
		Groups []struct {
			Name  string `json:"name"`
			File  string `json:"file"`
			Rules []struct {
				State    string `json:"state"`
				Name     string `json:"name"`
				Query    string `json:"query"`
				Duration int    `json:"duration"`
				Labels   struct {
					App      string `json:"app"`
					Severity string `json:"severity"`
				} `json:"labels"`
				Annotations struct {
					Dashboard   string `json:"dashboard"`
					Description string `json:"description"`
					Summary     string `json:"summary"`
				} `json:"annotations"`
				Alerts []interface{} `json:"alerts"`
				Health string        `json:"health"`
				Type   string        `json:"type"`
			} `json:"rules"`
			Interval int `json:"interval"`
		} `json:"groups"`
	} `json:"data"`
}
