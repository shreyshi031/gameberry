package entity

type Restaurant struct {
	RestaurantId  string  `json:"restaurant_id"`
	Cuisine       Cuisine `json:"cuisine"`
	CostBracket   int     `json:"cost_bracket"`
	Rating        float64 `json:"rating"`
	IsRecommended bool    `json:"is_recommended"`
	OnboardedTime string  `json:"onboarded_time"`
}
