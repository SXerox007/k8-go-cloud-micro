package samples

import (
	"fmt"
	"time"
)

/*
Question: Given an array of prices and an array of positions,
calculate the market value for each position by multiplying the price with the amount.
If a matching price is found for a position based on the product key and date,
display the position details along with the price value and market value.
If no matching price is found, display the position details with a message indicating that the market value is not found.
Finally, calculate the total market value of all positions.
*/

type Price struct {
	Date       time.Time
	ProductKey string
	Value      float64
}

type Position struct {
	Date       time.Time
	Amount     float64
	ProductKey string
}

func mainPosition() {
	prices := []Price{
		{Date: time.Now(), ProductKey: "A1", Value: 10},
		{Date: time.Now(), ProductKey: "A2", Value: 15},
		{Date: time.Now(), ProductKey: "A3", Value: 20},
		{Date: time.Now(), ProductKey: "A4", Value: 25},
	}

	positions := []Position{
		{Date: time.Now(), ProductKey: "A1", Amount: 8},
		{Date: time.Now(), ProductKey: "A2", Amount: 6},
		{Date: time.Now(), ProductKey: "A3", Amount: 4},
		{Date: time.Now(), ProductKey: "A4", Amount: 2},
	}

	totalMarketValue := 0.0

	for _, pos := range positions {
		var matchingPrice *Price

		for _, price := range prices {
			if price.ProductKey == pos.ProductKey && price.Date.Truncate(24*time.Hour).Equal(pos.Date.Truncate(24*time.Hour)) {
				matchingPrice = &price
				break
			}
		}

		if matchingPrice != nil {
			positionMarketValue := matchingPrice.Value * pos.Amount
			totalMarketValue += positionMarketValue

			fmt.Printf("Product Key: %s, Date: %s, Amount: %.2f, Price Value: %.2f, Market Value: %.2f\n", pos.ProductKey, pos.Date, pos.Amount, matchingPrice.Value, positionMarketValue)
		} else {
			fmt.Printf("Product Key: %s, Date: %s, Market Value: Not Found\n", pos.ProductKey, pos.Date)
		}
	}

	fmt.Printf("Total Market Value: %.2f\n", totalMarketValue)
}
