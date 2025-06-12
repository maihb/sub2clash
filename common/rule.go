package common

import (
	"fmt"
	"strings"

	"github.com/bestnite/sub2clash/model"
)

func PrependRuleProvider(
	sub *model.Subscription, providerName string, group string, provider model.RuleProvider,
) {
	if sub.RuleProvider == nil {
		sub.RuleProvider = make(map[string]model.RuleProvider)
	}
	sub.RuleProvider[providerName] = provider
	PrependRules(
		sub,
		fmt.Sprintf("RULE-SET,%s,%s", providerName, group),
	)
}

func AppenddRuleProvider(
	sub *model.Subscription, providerName string, group string, provider model.RuleProvider,
) {
	if sub.RuleProvider == nil {
		sub.RuleProvider = make(map[string]model.RuleProvider)
	}
	sub.RuleProvider[providerName] = provider
	AppendRules(sub, fmt.Sprintf("RULE-SET,%s,%s", providerName, group))
}

func PrependRules(sub *model.Subscription, rules ...string) {
	if sub.Rule == nil {
		sub.Rule = make([]string, 0)
	}
	sub.Rule = append(rules, sub.Rule...)
}

func AppendRules(sub *model.Subscription, rules ...string) {
	if sub.Rule == nil {
		sub.Rule = make([]string, 0)
	}
	matchRule := sub.Rule[len(sub.Rule)-1]
	if strings.Contains(matchRule, "MATCH") {
		sub.Rule = append(sub.Rule[:len(sub.Rule)-1], rules...)
		sub.Rule = append(sub.Rule, matchRule)
		return
	}
	sub.Rule = append(sub.Rule, rules...)
}
