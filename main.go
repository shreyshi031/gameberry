package main

import (
	"context"
	"fmt"
	"gameberry/entity"
	"gameberry/service/recommendation"
)

func main() {

	ctx := context.Background()
	recommendationService := recommendation.InitRecommendationService(ctx)

	user := entity.User{
		Cuisines: []entity.CuisineTracking{
			{
				CuisineType: entity.NorthIndian,
				NoOfOrders:  4,
			},
			{
				CuisineType: entity.Chinese,
				NoOfOrders:  3,
			},
			{
				CuisineType: entity.SouthIndian,
				NoOfOrders:  5,
			},
		},
		CostBracket: []entity.CostTracking{
			{
				CostType:   3,
				NoOfOrders: 5,
			},
			{
				CostType:   2,
				NoOfOrders: 2,
			},
			{
				CostType:   5,
				NoOfOrders: 0,
			},
		},
	}

	userMissingInfo := entity.User{
		Cuisines: []entity.CuisineTracking{
			{
				CuisineType: entity.SouthIndian,
			},
		},
		CostBracket: []entity.CostTracking{
			{
				CostType: 3,
			},
		},
	}

	restaurantCase1 := entity.Restaurant{
		RestaurantId:  "1",
		Cuisine:       entity.SouthIndian,
		CostBracket:   3,
		Rating:        4,
		IsRecommended: true,
		OnboardedTime: "",
	}

	restaurantCase2 := entity.Restaurant{
		RestaurantId:  "2",
		Cuisine:       entity.SouthIndian,
		CostBracket:   3,
		Rating:        4.5,
		IsRecommended: false,
		OnboardedTime: "",
	}

	restaurantCase3 := entity.Restaurant{
		RestaurantId:  "3",
		Cuisine:       entity.SouthIndian,
		CostBracket:   2,
		Rating:        4.6,
		IsRecommended: false,
		OnboardedTime: "",
	}

	restaurantCase4 := entity.Restaurant{
		RestaurantId:  "4",
		Cuisine:       entity.NorthIndian,
		CostBracket:   3,
		Rating:        4.6,
		IsRecommended: false,
		OnboardedTime: "",
	}

	restaurantCase5 := entity.Restaurant{
		RestaurantId:  "5",
		Cuisine:       entity.NorthIndian,
		CostBracket:   3,
		Rating:        4.6,
		IsRecommended: false,
		OnboardedTime: "46",
	}

	restaurantCase6 := entity.Restaurant{
		RestaurantId:  "6",
		Cuisine:       entity.SouthIndian,
		CostBracket:   3,
		Rating:        3.2,
		IsRecommended: false,
		OnboardedTime: "50",
	}

	restaurantCase7 := entity.Restaurant{
		RestaurantId:  "7",
		Cuisine:       entity.SouthIndian,
		CostBracket:   2,
		Rating:        4.2,
		IsRecommended: false,
		OnboardedTime: "",
	}

	restaurantCase8 := entity.Restaurant{
		RestaurantId:  "8",
		Cuisine:       entity.NorthIndian,
		CostBracket:   3,
		Rating:        4.3,
		IsRecommended: false,
		OnboardedTime: "",
	}

	restaurantCase9 := entity.Restaurant{
		RestaurantId:  "9",
		Cuisine:       entity.Chinese,
		CostBracket:   1,
		Rating:        3,
		IsRecommended: false,
		OnboardedTime: "",
	}

	runTestCase0(user, []entity.Restaurant{restaurantCase4, restaurantCase8, restaurantCase9, restaurantCase6, restaurantCase2,
		restaurantCase3, restaurantCase5, restaurantCase7, restaurantCase1}, recommendationService)
	runTestCase1(user, []entity.Restaurant{restaurantCase4, restaurantCase2,
		restaurantCase3, restaurantCase1}, recommendationService)
	runTestCase2(user, []entity.Restaurant{restaurantCase8, restaurantCase6,
		restaurantCase7}, recommendationService)

	runTestCase3(user, []entity.Restaurant{restaurantCase7, restaurantCase3,
		restaurantCase5}, recommendationService)
	runTestCase4(user, []entity.Restaurant{restaurantCase6, restaurantCase5}, recommendationService)

	runTestCase5(userMissingInfo, []entity.Restaurant{restaurantCase7, restaurantCase3,
		restaurantCase5}, recommendationService)

}

func runTestCase0(user entity.User, restaurantsList []entity.Restaurant,
	service *recommendation.RecommendationService) {
	ids := service.GetRestaurantRecommendations(user, restaurantsList)
	fmt.Println("\nTest Case -1 ------------------")
	fmt.Println("Result:", ids)
	fmt.Println("Expected: [1,2,3,4,5,6,7,8,9]")
}

func runTestCase1(user entity.User, restaurantsList []entity.Restaurant,
	service *recommendation.RecommendationService) {
	ids := service.GetRestaurantRecommendations(user, restaurantsList)

	fmt.Println("\nTest Case -2 ------------------")
	fmt.Println("Result:", ids)
	fmt.Println("Expected: [1,2,3,4]")
}

// Comparison b/n - 6,7,8 case
func runTestCase2(user entity.User, restaurantsList []entity.Restaurant,
	service *recommendation.RecommendationService) {
	ids := service.GetRestaurantRecommendations(user, restaurantsList)

	fmt.Println("\nTest Case -3 ------------------")
	fmt.Println("Result:", ids)
	fmt.Println("Expected: [6,7,8]")
}

func runTestCase3(user entity.User, restaurantsList []entity.Restaurant,
	service *recommendation.RecommendationService) {
	ids := service.GetRestaurantRecommendations(user, restaurantsList)

	fmt.Println("\nTest Case -4 ------------------")
	fmt.Println("Result:", ids)
	fmt.Println("Expected: [3,5,7]")
}

func runTestCase4(user entity.User, restaurantsList []entity.Restaurant,
	service *recommendation.RecommendationService) {
	ids := service.GetRestaurantRecommendations(user, restaurantsList)
	fmt.Println("\nTest Case -5 ------------------")
	fmt.Println("Result:", ids)
	fmt.Println("Expected: [5,6]")
}

func runTestCase5(user entity.User, restaurantsList []entity.Restaurant,
	service *recommendation.RecommendationService) {
	ids := service.GetRestaurantRecommendations(user, restaurantsList)

	fmt.Println("\nTest Case -4 ------------------")
	fmt.Println("Result:", ids)
	fmt.Println("Expected: [5,7,3]")
}
