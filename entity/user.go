package entity

import (
	"gameberry/config"
	"sort"
)

type User struct {
	Cuisines    []CuisineTracking `json:"cuisines"`
	CostBracket []CostTracking    `json:"cost_bracket"`
}

func (u *User) GetPrimaryCuisine() *CuisineTracking {
	var userPrimaryCuisine *CuisineTracking
	maxOrder := -1
	for _, cuisine := range u.Cuisines {
		currentCuisine := cuisine
		if currentCuisine.NoOfOrders > maxOrder {
			maxOrder = currentCuisine.NoOfOrders
			userPrimaryCuisine = &currentCuisine
		}
	}
	return userPrimaryCuisine
}

func (u *User) GetSecondaryCuisines() []CuisineTracking {
	userPrimaryCuisine := u.GetPrimaryCuisine()
	userSecondaryCuisines := []CuisineTracking{}
	count := 0
	for _, cuisine := range u.Cuisines {
		if cuisine.CuisineType == userPrimaryCuisine.CuisineType || count > config.TotalSecondaryCuisines {
			continue
		}
		userSecondaryCuisines = append(userSecondaryCuisines, cuisine)
		count++
	}
	sort.SliceStable(userSecondaryCuisines, func(i, j int) bool {
		return userSecondaryCuisines[i].NoOfOrders > userSecondaryCuisines[j].NoOfOrders
	})

	return userSecondaryCuisines
}

func (u *User) GetPrimaryCostBracket() *CostTracking {
	var userPrimaryCostBracket *CostTracking
	maxOrder := -1

	for _, costBracket := range u.CostBracket {
		currentCostBracket := costBracket
		if currentCostBracket.NoOfOrders > maxOrder {
			maxOrder = currentCostBracket.NoOfOrders
			userPrimaryCostBracket = &currentCostBracket
		}
	}
	return userPrimaryCostBracket
}

func (u *User) GetSecondaryCostBracket() []CostTracking {
	userPrimaryCostBracket := u.GetPrimaryCostBracket()
	userSecondaryCostBracket := []CostTracking{}

	count := 0
	for _, costBracket := range u.CostBracket {
		if costBracket.CostType == userPrimaryCostBracket.CostType || count > config.TotalSecondaryCostBrackets {
			continue
		}
		userSecondaryCostBracket = append(userSecondaryCostBracket, costBracket)
		count++
	}
	sort.SliceStable(userSecondaryCostBracket, func(i, j int) bool {
		return userSecondaryCostBracket[i].NoOfOrders > userSecondaryCostBracket[j].NoOfOrders
	})

	return userSecondaryCostBracket
}
