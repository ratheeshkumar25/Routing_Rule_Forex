package models

// Performance Action Enum
type RulePerformanceAction int32

const (
	RulePerformanceAction_All                 RulePerformanceAction = 0
	RulePerformanceAction_DelayInMilliseconds RulePerformanceAction = 1
	RulePerformanceAction_DelayInTicks        RulePerformanceAction = 2
	RulePerformanceAction_Reject              RulePerformanceAction = 3
	RulePerformanceAction_ConfirmByPrice      RulePerformanceAction = 4
	RulePerformanceAction_ConfirmByMarket     RulePerformanceAction = 5
	RulePerformanceAction_CancelOrder         RulePerformanceAction = 6
)

var RulePerformanceAction_name = map[int32]string{
	0: "all",
	1: "delay_in_milliseconds",
	2: "delay_in_ticks",
	3: "reject",
	4: "confirm_by_price",
	5: "confirm_by_market",
	6: "cancel_order",
}

var RulePerformanceAction_value = map[string]int32{
	"all":                   0,
	"delay_in_milliseconds": 1,
	"delay_in_ticks":        2,
	"reject":                3,
	"confirm_by_price":      4,
	"confirm_by_market":     5,
	"cancel_order":          6,
}

// Request Enum
type RuleRequest int32

const (
	RuleRequest_All              RuleRequest = 0
	RuleRequest_Price            RuleRequest = 1
	RuleRequest_MarketExecution  RuleRequest = 2
	RuleRequest_InstantExecution RuleRequest = 3
	RuleRequest_PendingOrders    RuleRequest = 4
	RuleRequest_SLTPModification RuleRequest = 5
	RuleRequest_RequestExecution RuleRequest = 6
	RuleRequest_SLActivation     RuleRequest = 7
)

var RuleRequest_name = map[int32]string{
	0: "price",
	1: "market_execution",
	2: "instant_execution",
	3: "pending_orders",
	4: "sl_tp_modification",
	5: "request_execution",
	6: "sl_activation",
}

var RuleRequest_value = map[string]int32{
	"price":              0,
	"market_execution":   1,
	"instant_execution":  2,
	"pending_orders":     3,
	"sl_tp_modification": 4,
	"request_execution":  5,
	"sl_activation":      6,
}

// Order Enum
type RuleOrder int32

const (
	RuleOrder_All           RuleOrder = 0
	RuleOrder_Buy           RuleOrder = 1
	RuleOrder_Sell          RuleOrder = 2
	RuleOrder_BuyLimit      RuleOrder = 3
	RuleOrder_SellLimit     RuleOrder = 4
	RuleOrder_BuyStop       RuleOrder = 5
	RuleOrder_SellStop      RuleOrder = 6
	RuleOrder_BuyStopLimit  RuleOrder = 7
	RuleOrder_SellStopLimit RuleOrder = 8
)

var RuleOrder_name = map[int32]string{
	0: "all",
	1: "buy",
	2: "sell",
	3: "buy_limit",
	4: "sell_limit",
	5: "buy_stop",
	6: "sell_stop",
	7: "buy_stop_limit",
	8: "sell_stop_limit",
}

var RuleOrder_value = map[string]int32{
	"all":             0,
	"buy":             1,
	"sell":            2,
	"buy_limit":       3,
	"sell_limit":      4,
	"buy_stop":        5,
	"sell_stop":       6,
	"buy_stop_limit":  7,
	"sell_stop_limit": 8,
}

// Conditions Enum
type RuleConditionType string

const (
	RuleConditionType_All        RuleConditionType = "all"
	RuleConditionType_Request    RuleConditionType = "request"
	RuleConditionType_Account    RuleConditionType = "account"
	RuleConditionType_Order      RuleConditionType = "order"
	RuleConditionType_Statistics RuleConditionType = "statistics"
	RuleConditionType_Symbol     RuleConditionType = "symbol"
)

// Operator Enum
type RuleOperator string

const (
	OperatorEqual        RuleOperator = "="
	OperatorNotEqual     RuleOperator = "!="
	OperatorGreaterThan  RuleOperator = ">"
	OperatorGreaterEqual RuleOperator = ">="
	OperatorLessThan     RuleOperator = "<"
	OperatorLessEqual    RuleOperator = "<="
)

// RuleCondition Model
type RuleCondition struct {
	Type           RuleConditionType `json:"type"`
	Operator       RuleOperator      `json:"operator"`
	ConditionValue interface{}       `json:"value"` // Flexible to accommodate multiple types
}

// RoutingRule Model

type RoutingRule struct {
	ID          int32                 `json:"id"`
	RuleName    string                `json:"rule_name"`
	Performance RulePerformanceAction `json:"performance"`
	RequestType RuleRequest           `json:"request_type"`
	OrderType   RuleOrder             `json:"order_type"`
	Conditions  []RuleCondition       `json:"conditions"`
	Enabled     bool                  `json:"enabled"`
	Priority    int32                 `json:"priority"`
	DealerID    int32                 `json:"dealer_id"` // Associated Dealer
	TraderID    int32                 `json:"trader_id"` // Associated Trader
}
