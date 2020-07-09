package main

import (
	"fmt"
	"strings"
)

func AlertGroupsToMarkdown(alertGroups []AlertGroup) string {
	var str strings.Builder
	for _, v := range alertGroups {
		AlertGroupToMarkdown(v, &str)
	}
	return str.String()
}

func AlertGroupToMarkdown(alertGroup AlertGroup, str *strings.Builder) {
	Heading1(alertGroup.Name, str)
	for k, _ := range alertGroup.Rules {
		rule := alertGroup.Rules[k]
		Heading3(rule.Name, str)
		str.WriteString("Labels: ")
		Map(rule.Labels, str)
		str.WriteString("Annotations: ")
		Map(rule.Annotations, str)
		str.WriteString("Query: ")
		Text(rule.Query, str)
		str.WriteString("Duration [s]: ")
		Text(fmt.Sprintf("%v", rule.Duration), str)
	}
}

func Heading1(h string, str *strings.Builder) {
	h1format := []string{"# ", h, "\n"}
	for _, v := range h1format {
		str.WriteString(v)
	}
}

func Heading2(h string, str *strings.Builder) {
	h2format := []string{"## ", h, "\n"}
	for _, v := range h2format {
		str.WriteString(v)
	}
}

func Heading3(h string, str *strings.Builder) {
	h3format := []string{"### ", h, "\n"}
	for _, v := range h3format {
		str.WriteString(v)
	}
}

func Map(h map[string]string, str *strings.Builder) {
	str.WriteString("  \n")
	for k, v := range h {
		format := fmt.Sprintf("&nbsp;&nbsp;&nbsp;&nbsp; %s: %s  \n", k, v)
		str.WriteString(format)
	}
	str.WriteString("  \n")
}

func Text(h string, str *strings.Builder) {
	str.WriteString(fmt.Sprintf("%s  \n", h))
}
