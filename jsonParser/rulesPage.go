package jsonParser

type RulesPage struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}
type Labels struct {
	App      string `json:"app"`
	Severity string `json:"severity"`
}
type Annotations struct {
	Dashboard   string `json:"dashboard"`
	Description string `json:"description"`
	Summary     string `json:"summary"`
}
type Rules struct {
	State       string        `json:"state"`
	Name        string        `json:"name"`
	Query       string        `json:"query"`
	Duration    int           `json:"duration"`
	Labels      Labels        `json:"labels"`
	Annotations Annotations   `json:"annotations"`
	Alerts      []interface{} `json:"alerts"`
	Health      string        `json:"health"`
	Type        string        `json:"type"`
}
type Groups struct {
	Name     string  `json:"name"`
	File     string  `json:"file"`
	Rules    []Rules `json:"rules"`
	Interval int     `json:"interval"`
}
type Data struct {
	Groups []Groups `json:"groups"`
}
