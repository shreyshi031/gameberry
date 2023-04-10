package entity

type Cuisine string

const (
	SouthIndian Cuisine = "SouthIndian"
	NorthIndian Cuisine = "NorthIndian"
	Chinese     Cuisine = "Chinese"
)

type CuisineTracking struct {
	NoOfOrders  int     `json:"no_of_orders"`
	CuisineType Cuisine `json:"cuisine_type"`
}
