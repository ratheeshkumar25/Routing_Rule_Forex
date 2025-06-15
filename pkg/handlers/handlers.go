package handlers

import (
	"github.com/gofiber/fiber/v2"
	models "github.com/ratheeshkumar25/routing_featureMT5VsHS/pkg/model"
)

var rules = []models.RoutingRule{}

// AddRule adds a new routing rule
//	@Summary		Add a new routing rule
//	@Description	Create a new routing rule for a dealer or trader
//	@Tags			Rules
//	@Accept			json
//	@Produce		json
//	@Param			rule	body		models.RoutingRule	true	"Routing Rule"
//	@Success		201		{object}	models.RoutingRule
//	@Failure		400		{object}	map[string]string	"Invalid Payload"
//	@Router			/api/v1/rules [post]
func AddRule(c *fiber.Ctx) error {
	rule := new(models.RoutingRule)

	if err := c.BodyParser(rule); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid payload"})
	}

	if rule.DealerID == 0 && rule.TraderID == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Either DealerID or TraderID must be specified"})
	}
	rules = append(rules, *rule)
	return c.Status(201).JSON(rule)
}

// UpdateRule updates an existing routing rule
//	@Summary		Update a routing rule
//	@Description	Update the details of an existing routing rule
//	@Tags			Rules
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Rule ID"
//	@Param			rule	body		models.RoutingRule	true	"Updated Routing Rule"
//	@Success		200		{object}	models.RoutingRule
//	@Failure		400		{object}	map[string]string	"Invalid Rule ID or Payload"
//	@Failure		404		{object}	map[string]string	"Rule Not Found"
//	@Router			/api/v1/rules/{id} [put]
func UpdateRule(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid rule ID"})
	}

	for i, rule := range rules {
		if rule.ID == int32(id) {
			updatedRule := new(models.RoutingRule)
			if err := c.BodyParser(updatedRule); err != nil {
				return c.Status(400).JSON(fiber.Map{"error": "Invalid payload"})
			}

			rules[i] = *updatedRule
			return c.Status(200).JSON(updatedRule)
		}
	}

	return c.Status(404).JSON(fiber.Map{"error": "Rule not found"})
}

// DeleteRule deletes an existing routing rule
//	@Summary		Delete a routing rule
//	@Description	Remove a routing rule by ID
//	@Tags			Rules
//	@Produce		json
//	@Param			id	path		int					true	"Rule ID"
//	@Success		204	{string}	string				"No Content"
//	@Failure		400	{object}	map[string]string	"Invalid Rule ID"
//	@Failure		404	{object}	map[string]string	"Rule Not Found"
//	@Router			/api/v1/rules/{id} [delete]
func DeleteRule(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid rule ID"})
	}

	for i, rule := range rules {
		if rule.ID == int32(id) {
			rules = append(rules[:i], rules[i+1:]...)
			return c.SendStatus(204)
		}
	}

	return c.Status(404).JSON(fiber.Map{"error": "Rule not found"})
}

// GetRules retrieves routing rules
//	@Summary		Get routing rules
//	@Description	Retrieve routing rules filtered by dealer or trader ID
//	@Tags			Rules
//	@Produce		json
//	@Param			dealer_id	query	int	false	"Dealer ID"
//	@Param			trader_id	query	int	false	"Trader ID"
//	@Success		200			{array}	models.RoutingRule
//	@Router			/api/v1/rules [get]
func GetRules(c *fiber.Ctx) error {
	dealerID := c.QueryInt("dealer_id") // Defaults to 0 if not present
	traderID := c.QueryInt("trader_id") // Defaults to 0 if not present

	filteredRules := []models.RoutingRule{}
	for _, rule := range rules {
		if (dealerID != 0 && rule.DealerID == int32(dealerID)) ||
			(traderID != 0 && rule.TraderID == int32(traderID)) {
			filteredRules = append(filteredRules, rule)
		}
	}

	return c.JSON(filteredRules)
}
