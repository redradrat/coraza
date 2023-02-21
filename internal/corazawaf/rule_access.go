package corazawaf

import "github.com/corazawaf/coraza/v3/types/variables"

func hasRulesForRequestBody(rg RuleGroup) bool {
	for _, rule := range rg.rules {
		for _, v := range rule.variables {
			if v.Variable == variables.RequestBody {
				return true
			}
		}
	}

	return false
}
