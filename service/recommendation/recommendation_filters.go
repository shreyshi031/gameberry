package recommendation

import (
	"gameberry/entity"
	"gameberry/utils"
	"strconv"
)

func (s *RecommendationService) CuisineFilter(currentRestaurant entity.Restaurant, user entity.User,
	sortFilter Filter) bool {
	if sortFilter.CuisineFilter == nil {
		return true
	}

	userCuisineTypes := []entity.Cuisine{}
	for _, userCuisineType := range user.Cuisines {
		userCuisineTypes = append(userCuisineTypes, userCuisineType.CuisineType)
	}

	cuisineOrder := *sortFilter.CuisineFilter
	if cuisineOrder == Primary {
		userPrimaryCuisine := user.GetPrimaryCuisine()
		if userPrimaryCuisine != nil && (*userPrimaryCuisine).CuisineType == currentRestaurant.Cuisine {
			return true
		}
	}
	if cuisineOrder == Secondary {
		userSecondaryCuisines := user.GetSecondaryCuisines()
		if utils.ContainsCuisines(userSecondaryCuisines, currentRestaurant.Cuisine) {
			return true
		}
	}
	return false
}

func (s *RecommendationService) CostFilter(currentRestaurant entity.Restaurant, user entity.User,
	sortFilter Filter) bool {
	if sortFilter.CostFilter == nil {
		return true
	}

	userCostTypes := []int{}
	for _, userCostType := range user.CostBracket {
		userCostTypes = append(userCostTypes, userCostType.CostType)
	}

	costOrder := *sortFilter.CostFilter
	if costOrder == Primary {
		userPrimaryCostBracket := user.GetPrimaryCostBracket()
		if userPrimaryCostBracket != nil && (*userPrimaryCostBracket).CostType == currentRestaurant.CostBracket {
			return true
		}
	}
	if costOrder == Secondary {
		userSecondaryCostBracket := user.GetSecondaryCostBracket()
		if utils.ContainsCostBrackets(userSecondaryCostBracket, currentRestaurant.CostBracket) {
			return true
		}
	}
	return false
}

func (s *RecommendationService) RatingRangeFilter(currentRestaurant entity.Restaurant, user entity.User,
	sortFilter Filter) bool {
	if sortFilter.RatingRange == nil {
		return true
	}
	ratingRanges := *sortFilter.RatingRange
	if (currentRestaurant.Rating > (ratingRanges)[0]) && (currentRestaurant.Rating < (ratingRanges)[1]) {
		return true
	}
	return false
}

func (s *RecommendationService) IsFeaturedFilter(currentRestaurant entity.Restaurant, user entity.User,
	sortFilter Filter) bool {
	if sortFilter.IsFeatured == nil {
		return true
	}
	if *sortFilter.IsFeatured && currentRestaurant.IsRecommended {
		return true
	}
	return false
}

func (s *RecommendationService) ExcludedRestaurantFilter(currentRestaurant entity.Restaurant, user entity.User,
	sortFilter Filter) bool {
	if sortFilter.ExcludedRestaurant == nil {
		return true
	}
	excludedRestaurants := *sortFilter.ExcludedRestaurant
	for _, excludedRestaurant := range excludedRestaurants {
		if excludedRestaurant.RestaurantId == currentRestaurant.RestaurantId {
			return false
		}
	}
	return true
}

func (s *RecommendationService) NewOnboardedFilter(currentRestaurant entity.Restaurant, user entity.User,
	sortFilter Filter) bool {
	if sortFilter.NewOnboardedFilter == nil {
		return true
	}

	if *sortFilter.NewOnboardedFilter && currentRestaurant.OnboardedTime != "" {
		onboardedHours, _ := strconv.Atoi(currentRestaurant.OnboardedTime)
		if onboardedHours < 48 {
			return true
		}
	}
	return false
}
