package recommendation

import (
	"gameberry/entity"
)

type Service interface {
	GetRestaurantRecommendations(user entity.User, availableRestaurants []entity.Restaurant) []string
}
