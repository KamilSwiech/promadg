package parse

import (
	"testing"
)

func TestRulesPageToMD(t *testing.T) {
	var rp = RulesPage{
		Status: "poc",
	}
	if err := RulesPageToMD(rp); err != nil {
		t.Error(err)
	}
}
