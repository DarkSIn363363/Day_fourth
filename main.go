package main

import (
	"fmt"
)

func validateTrip(
	distanceKm float64,
	consumptionPer100Km float64,
	fuelPrice float64,
	tolls float64,
	passengers int,
	driverPays bool,
) string {

	if distanceKm <= 0 {
		return "расстояние должно быть больше 0"
	}
	if consumptionPer100Km <= 0 {
		return "расход топлива должен быть больше 0"
	}
	if fuelPrice <= 0 {
		return "цена топлива должна быть больше 0"
	}
	if tolls < 0 {
		return "платные дороги не могут быть отрицательными"
	}
	if passengers < 1 || passengers > 8 {
		return "количество пассажиров должно быть от 1 до 8"
	}
	if passengers == 1 && !driverPays {
		return "поездка не может быть бесплатной"
	}
	return ""
}

func calculateFuelLiters(
	distanceKm float64,
	consumptionPer100Km float64,
) float64 {
	liters := distanceKm * consumptionPer100Km / 100
	return liters
}

func calculateTripCosts(
	fuelLiters float64,
	fuelPrice float64,
	tolls float64,
) (float64, float64) {
	fuelCost := fuelLiters * fuelPrice
	totalCost := fuelCost + tolls
	return fuelCost, totalCost
}

func calculatePaymentPerPerson(
	totalCost float64,
	passengers int,
	driverPays bool,
) (paymentPerPerson float64, payingPassengers int) {

	if driverPays {
		payingPassengers = passengers
		paymentPerPerson = totalCost / float64(payingPassengers)
	} else {
		payingPassengers = passengers - 1
		paymentPerPerson = totalCost / float64(payingPassengers)
	}
	return
}

func main() {

	distanceKm := 350.0
	consumptionPer100Km := 8.5
	fuelPrice := 62.3
	tolls := 500.0
	passengers := 4
	driverPays := true
    
	validateError := validateTrip(distanceKm, consumptionPer100Km, fuelPrice, tolls, passengers, driverPays)
	if  validateError != "" {
		fmt.Println(validateError)
		return
	}
	fuelLiters := calculateFuelLiters(distanceKm, consumptionPer100Km)
	fuelCost, totalCost := calculateTripCosts(fuelLiters, fuelPrice, tolls)
	costPerPerson, _ := calculatePaymentPerPerson(totalCost, passengers, driverPays)
	fmt.Printf("Расстояние: %.2f\n", distanceKm)
	fmt.Printf("Требуемое топливо: %.2f\n", fuelLiters)
	fmt.Printf("Стоимость топлива: %.2f\n", fuelCost)
	fmt.Printf("Платные дороги: %.2f\n", tolls)
	fmt.Printf("Общая стоимость поездки: %.2f\n", totalCost)
	fmt.Printf("Стоимость на одного пассажира: %.2f\n", costPerPerson)
}
