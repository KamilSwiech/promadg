package main

type AlertGroup struct {
	Name     string
	File     string
	Rules    []Alert
	Interval float64
}

type Alert struct {
	Name        string
	Labels      map[string]string
	Annotations map[string]string
	Query       string
	Duration    float64
}
