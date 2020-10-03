package parse

type RulesPage struct {
	Status string `json:"status"`
	Data   struct {
		Groups []struct {
			Name  string `json:"name"`
			File  string `json:"file"`
			Rules []struct {
				State       string            `json:"state"`
				Name        string            `json:"name"`
				Query       string            `json:"query"`
				Duration    int               `json:"duration"`
				Labels      map[string]string `json:"labels"`
				Annotations map[string]string `json:"annotations"`
				Alerts      []interface{}     `json:"alerts"`
				Health      string            `json:"health"`
				Type        string            `json:"type"`
			} `json:"rules"`
			Interval int `json:"interval"`
		} `json:"groups"`
	} `json:"data"`
}
