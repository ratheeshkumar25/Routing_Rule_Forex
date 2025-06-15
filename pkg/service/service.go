package service

import (
	models "github.com/ratheeshkumar25/routing_featureMT5VsHS/pkg/model"
)

func EvaluvateRule(rule models.RoutingRule, request map[string]interface{}) bool {
	//action about the rules
	if !rule.Enabled {
		return false
	}

	//check the request type
	if requestType, ok := request["request_type"].(int32); ok {
		if !contains([]models.RuleRequest{rule.RequestType}, models.RuleRequest(requestType)) {
			return false
		}
	} else {
		return false // Invalid type, fail the evaluation
	}

	// Check order types
	if orderType, ok := request["order_type"].(int32); ok {
		if !contains([]models.RuleOrder{rule.OrderType}, models.RuleOrder(orderType)) {
			return false
		}
	} else {
		return false // Invalid type, fail the evaluation
	}

	// Evaluate conditions
	for _, condition := range rule.Conditions {
		if !evaluateCondition(condition, request) {
			return false
		}
	}

	return true

}

func evaluateCondition(condition models.RuleCondition, request map[string]interface{}) bool {
	requestValue, exists := request[string(condition.Type)]
	if !exists {
		return false
	}
	switch condition.Operator {
	case "=":
		return requestValue == condition.ConditionValue
	case "!=":
		return requestValue != condition.ConditionValue
	case ">":
		return requestValue.(float64) > condition.ConditionValue.(float64)
	case ">=":
		return requestValue.(float64) >= condition.ConditionValue.(float64)
	case "<":
		return requestValue.(float64) < condition.ConditionValue.(float64)
	case "<=":
		return requestValue.(float64) <= condition.ConditionValue.(float64)
	default:
		return false
	}
}

func contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// func contains(slice []models.RuleRequest, item models.RuleRequest) bool {
// 	for _, v := range slice {
// 		if v == item {
// 			return true
// 		}
// 	}
// 	return false
// }
