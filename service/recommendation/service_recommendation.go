package recommendation

import (
	"context"
	"gameberry/config"
	"gameberry/entity"
	"gameberry/utils"
)

type RecommendationService struct {
	ctx context.Context
}

func InitRecommendationService(
	ctx context.Context,
) *RecommendationService {
	s := &RecommendationService{
		ctx: ctx,
	}
	return s
}

func (s *RecommendationService) GetRestaurantRecommendations(user entity.User,
	availableRestaurants []entity.Restaurant) []string {
	sortedRestaurants := utils.RemoveDuplicateRestaurants(s.GetSortedRestaurants(user, availableRestaurants))
	restaurantIds := []string{}
	for _, sortedRestaurant := range sortedRestaurants {
		restaurantIds = append(restaurantIds, sortedRestaurant.RestaurantId)
	}
	return utils.LimitMaxIds(restaurantIds, config.MaxTotalRestaurants)
}

func (s *RecommendationService) GetSortedRestaurants(user entity.User,
	availableRestaurants []entity.Restaurant) []entity.Restaurant {

	primaryOrder := Primary
	secondaryOrder := Secondary
	isTrue := true
	filteredRestaurants := []entity.Restaurant{}
	totalLimit := 4

	// Pipeline 1 for filter 1
	{
		filteredRestaurants = s.GetFilteredRestaurants(availableRestaurants, user,
			Filter{CuisineFilter: &primaryOrder, CostFilter: &primaryOrder, IsFeatured: &isTrue})
		if len(filteredRestaurants) == 0 {
			filteredRestaurants = s.CollectAllRestaurants(
				s.GetFilteredRestaurants(availableRestaurants, user,
					Filter{CuisineFilter: &primaryOrder, CostFilter: &secondaryOrder, IsFeatured: &isTrue}),
				s.GetFilteredRestaurants(availableRestaurants, user,
					Filter{CuisineFilter: &secondaryOrder, CostFilter: &primaryOrder, IsFeatured: &isTrue}),
			)
		}
	}

	// Pipeline 2 for other filters (filter 2,3,4)
	{
		filteredRestaurants = s.CollectAllRestaurants(
			filteredRestaurants,
			s.GetFilteredRestaurants(availableRestaurants, user,
				Filter{CuisineFilter: &primaryOrder, CostFilter: &primaryOrder, RatingRange: &[]float64{4, 5}}),
			s.GetFilteredRestaurants(availableRestaurants, user,
				Filter{CuisineFilter: &primaryOrder, CostFilter: &secondaryOrder, RatingRange: &[]float64{4.5, 5}}),
			s.GetFilteredRestaurants(availableRestaurants, user,
				Filter{CuisineFilter: &secondaryOrder, CostFilter: &primaryOrder, RatingRange: &[]float64{4.5, 5}}),
		)
	}

	// Pipeline 3 for other filters (filter 5)
	// we need to exclude all restaurants that we got till now in result
	{
		filteredRestaurants = s.CollectAllRestaurants(
			filteredRestaurants,
			s.GetFilteredRestaurants(availableRestaurants, user,
				Filter{TotalLimit: &totalLimit,
					ExcludedRestaurant: &filteredRestaurants,
					NewOnboardedFilter: &isTrue,
				},
			),
		)
	}

	// Pipeline 3 for other filters (filter 6,7,8,9)
	{
		filteredRestaurants = s.CollectAllRestaurants(
			filteredRestaurants,
			s.GetFilteredRestaurants(availableRestaurants, user,
				Filter{CuisineFilter: &primaryOrder, CostFilter: &primaryOrder, RatingRange: &[]float64{0, 4}}),
			s.GetFilteredRestaurants(availableRestaurants, user,
				Filter{CuisineFilter: &primaryOrder, CostFilter: &secondaryOrder, RatingRange: &[]float64{0, 4.5}}),
			s.GetFilteredRestaurants(availableRestaurants, user,
				Filter{CuisineFilter: &secondaryOrder, CostFilter: &primaryOrder, RatingRange: &[]float64{0, 4.5}}),
			s.GetFilteredRestaurants(availableRestaurants, user, Filter{}),
		)
	}
	return filteredRestaurants
}
