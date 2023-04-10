package utils

import "gameberry/entity"

func RemoveDuplicateRestaurants(restaurants []entity.Restaurant) []entity.Restaurant {
	restaurantIds := map[string]bool{}
	result := []entity.Restaurant{}
	for _, restaurant := range restaurants {
		if !restaurantIds[restaurant.RestaurantId] {
			result = append(result, restaurant)
			restaurantIds[restaurant.RestaurantId] = true
		}
	}
	return result
}

func LimitMaxRestaurants(restaurants []entity.Restaurant, limit *int) []entity.Restaurant {
	if limit == nil {
		return restaurants
	}
	maxSize := len(restaurants)
	if *limit < maxSize {
		maxSize = *limit
	}
	limitRestaurants := []entity.Restaurant{}
	for ind, restaurant := range restaurants {

		if ind < maxSize {
			limitRestaurants = append(limitRestaurants, restaurant)
		} else {
			break
		}
	}
	return limitRestaurants
}

func LimitMaxIds(ids []string, limit int) []string {

	maxSize := len(ids)
	if limit < maxSize {
		maxSize = limit
	}
	limitIds := []string{}
	for ind, id := range ids {

		if ind < maxSize {
			limitIds = append(limitIds, id)
		} else {
			break
		}
	}
	return limitIds
}

func ContainsCuisines(userCuisines []entity.CuisineTracking, cuisine entity.Cuisine) bool {
	for _, userCuisine := range userCuisines {
		if userCuisine.CuisineType == cuisine {
			return true
		}
	}
	return false
}

func ContainsCostBrackets(userCostBrackets []entity.CostTracking, costTracking int) bool {
	for _, userCostBracket := range userCostBrackets {
		if userCostBracket.CostType == costTracking {
			return true
		}
	}
	return false
}
