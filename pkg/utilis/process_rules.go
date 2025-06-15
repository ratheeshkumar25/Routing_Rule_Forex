package utilis

import (
	"sync"

	models "github.com/ratheeshkumar25/routing_featureMT5VsHS/pkg/model"
	"github.com/ratheeshkumar25/routing_featureMT5VsHS/pkg/service"
)

func ProcessRuleConcurrently(rules []models.RoutingRule, request map[string]interface{}) []string {

	results := make(chan string, len(rules))

	var wg sync.WaitGroup

	for _, rule := range rules {
		wg.Add(1)
		go func(rule models.RoutingRule) {
			defer wg.Done()
			if service.EvaluvateRule(rule, request) {
				results <- "Matched Rule" + rule.RuleName
			} else {
				results <- "No Match" + rule.RuleName
			}

		}(rule)
	}
	wg.Wait()
	close(results)

	response := []string{}

	for result := range results {
		response = append(response, result)
	}
	return response
}
