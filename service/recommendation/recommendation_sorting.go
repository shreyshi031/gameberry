package recommendation

import (
	"gameberry/entity"
	"gameberry/utils"
)

type Order string
type SortType string

const (
	Primary   Order = "Primary"
	Secondary Order = "Secondary"
)

type Filter struct {
	SortType           *SortType
	CuisineFilter      *Order
	CostFilter         *Order
	RatingRange        *[]float64
	IsFeatured         *bool
	ExcludedRestaurant *[]entity.Restaurant
	NewOnboardedFilter *bool
	TotalLimit         *int
}

type SortFilterFunction func(restaurant entity.Restaurant, user entity.User, filter Filter) bool

func (s *RecommendationService) GetFilteredRestaurants(restaurants []entity.Restaurant,
	user entity.User, filter Filter) []entity.Restaurant {
	restaurants = s.applyOtherFilter(restaurants, user, filter)
	return utils.LimitMaxRestaurants(restaurants, filter.TotalLimit)
}

func (s *RecommendationService) applyOtherFilter(restaurants []entity.Restaurant,
	user entity.User, filter Filter) []entity.Restaurant {
	sortFilterFunctions := []SortFilterFunction{
		s.CuisineFilter,
		s.CostFilter,
		s.RatingRangeFilter,
		s.IsFeaturedFilter,
		s.ExcludedRestaurantFilter,
		s.NewOnboardedFilter,
	}

	filteredRestaurant := []entity.Restaurant{}
	for _, currentRestaurant := range restaurants {
		isFiltered := true
		for _, sortFilterFunction := range sortFilterFunctions {
			isFiltered = isFiltered && sortFilterFunction(currentRestaurant, user, filter)
		}
		if isFiltered {
			filteredRestaurant = append(filteredRestaurant, currentRestaurant)
		}
	}
	return filteredRestaurant
}

func (s *RecommendationService) CollectAllRestaurants(availableRestaurants ...[]entity.Restaurant) []entity.Restaurant {
	result := []entity.Restaurant{}
	for _, availableRestaurant := range availableRestaurants {
		result = append(result, availableRestaurant...)
	}
	return result
}
