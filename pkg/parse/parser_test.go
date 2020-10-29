package parse

import (
	"testing"
)

func TestRulesPageToMD(t *testing.T) {
	var rp = RulesPage{
		Status: "poc",
	}
	RulesPageToMD(rp)
}
