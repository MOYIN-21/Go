package main

import "fmt"

func main() {
	var totalMiles, totalGallons int
	tripCount := 0

	fmt.Println("Miles Per Gallon Calculator")

	for {
		var miles, gallons int

		fmt.Print("Enter miles driven (or -1 to exit): ")
		_, err := fmt.Scan(&miles)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if miles == -1 {
			break
		}

		fmt.Print("Enter gallons used: ")
		_, err = fmt.Scan(&gallons)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		tripCount++
		tripMPG := float64(miles) / float64(gallons)

		fmt.Printf("Trip %d - MPG: %.2f\n", tripCount, tripMPG)

		totalMiles += miles
		totalGallons += gallons

		if tripCount > 1 {
			totalMPG := float64(totalMiles) / float64(totalGallons)
			fmt.Printf("Combined MPG (all trips): %.2f\n", totalMPG)
		}
	}
}