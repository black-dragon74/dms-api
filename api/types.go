//
// This file stores the structure of the user defined data types
//

package api

type MessMenuModel struct {
	LastUpdatedAt   string   `json:"last_updated_at"`
	LastUpdatedMeal string   `json:"last_updated_meal"`
	Breakfast       []string `json:"breakfast"`
	Lunch           []string `json:"lunch"`
	HighTea         []string `json:"high_tea"`
	Dinner          []string `json:"dinner"`
}
