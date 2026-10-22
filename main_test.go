package main

import (
	"math"
	"testing"
)

func TestValidateTrip(t *testing.T) {
	tests := []struct {
		name                string
		distanceKm          float64
		consumptionPer100Km float64
		fuelPrice           float64
		tolls               float64
		passengers          int
		driverPays          bool
		want                string
	}{
		{
			name:                "корректная поездка",
			distanceKm:          100,
			consumptionPer100Km: 10.5,
			fuelPrice:           78.65,
			tolls:               300,
			passengers:          4,
			driverPays:          true,
			want:                "",
		},
		{
			name:                "нулевое расстояние",
			distanceKm:          0,
			consumptionPer100Km: 10.5,
			fuelPrice:           78.65,
			tolls:               300,
			passengers:          4,
			driverPays:          true,
			want:                "расстояние должно быть больше 0",
		},
		{
			name:                "отрицательный расход",
			distanceKm:          100,
			consumptionPer100Km: -10.5,
			fuelPrice:           78.65,
			tolls:               300,
			passengers:          4,
			driverPays:          true,
			want:                "расход топлива должен быть больше 0",
		},
		{
			name:                "нулевая цена топлива",
			distanceKm:          100,
			consumptionPer100Km: 10.5,
			fuelPrice:           0,
			tolls:               300,
			passengers:          4,
			driverPays:          true,
			want:                "цена топлива должна быть больше 0",
		},
		{
			name:                "отрицательные платные дороги",
			distanceKm:          100,
			consumptionPer100Km: 10.5,
			fuelPrice:           78.65,
			tolls:               -300,
			passengers:          4,
			driverPays:          true,
			want:                "платные дороги не могут быть отрицательными",
		},
		{
			name:                "ноль пассажиров",
			distanceKm:          100,
			consumptionPer100Km: 10.5,
			fuelPrice:           78.65,
			tolls:               300,
			passengers:          0,
			driverPays:          true,
			want:                "количество пассажиров должно быть от 1 до 8",
		},
		{
			name:                "девять пассажиров",
			distanceKm:          100,
			consumptionPer100Km: 10.5,
			fuelPrice:           78.65,
			tolls:               300,
			passengers:          9,
			driverPays:          true,
			want:                "количество пассажиров должно быть от 1 до 8",
		},
		{
			name:                "водитель не платит при 1 пассажире",
			distanceKm:          100,
			consumptionPer100Km: 10.5,
			fuelPrice:           78.65,
			tolls:               300,
			passengers:          1,
			driverPays:          false,
			want:                "поездка не может быть бесплатной",
		},
		{
			name:                "несколько ошибок",
			distanceKm:          100,
			consumptionPer100Km: -10.5,
			fuelPrice:           78.65,
			tolls:               -300,
			passengers:          4,
			want:                "расход топлива должен быть больше 0",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := validateTrip(test.distanceKm, test.consumptionPer100Km, test.fuelPrice, test.tolls, test.passengers, test.driverPays)

			if got != test.want {
				t.Errorf(
					"validateTrip(%.2f, %.2f, %.2f, %.2f, %d, %t) = %s, want %s",
					test.distanceKm,
					test.consumptionPer100Km,
					test.fuelPrice,
					test.tolls,
					test.passengers,
					test.driverPays,
					got,
					test.want,
				)
			}
		})
	}
}

func TestCalculateFuelLiters(t *testing.T) {
	tests := []struct {
		name                string
		distanceKm          float64
		consumptionPer100Km float64
		want                float64
	}{
		{
			name:                "Дистанция 100км, расход 8,5л",
			distanceKm:          100,
			consumptionPer100Km: 8.5,
			want:                8.5,
		},
		{
			name:                "Дистанция 250км, расход 7,2л",
			distanceKm:          250,
			consumptionPer100Km: 7.2,
			want:                18,
		},
		{
			name:                "Дистанция 350км, расход 10л",
			distanceKm:          350,
			consumptionPer100Km: 10,
			want:                35,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := calculateFuelLiters(test.distanceKm, test.consumptionPer100Km)

			const epsilon = 0.000001

			if math.Abs(got-test.want) > epsilon {
				t.Errorf(
					"calculateFuelLiters(%.2f, %.2f) = %.2f, want %.2f",
					test.distanceKm,
					test.consumptionPer100Km,
					got,
					test.want,
				)
			}
		})
	}
}

func TestCalculateTripCosts(t *testing.T) {
	tests := []struct {
		name       string
		fuelLiters float64
		fuelPrice  float64
		tolls      float64
		wantFuel   float64
		wantTotal  float64
	}{
		{
			name:       "без платных дорог",
			fuelLiters: 50,
			fuelPrice:  60,
			tolls:      0,
			wantFuel:   3000,
			wantTotal:  3000,
		},
		{
			name:       "c платными дорогами",
			fuelLiters: 50,
			fuelPrice:  60,
			tolls:      500,
			wantFuel:   3000,
			wantTotal:  3500,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotFuel, gotTotal := calculateTripCosts(test.fuelLiters, test.fuelPrice, test.tolls)

			const epsilon = 0.000001

			if math.Abs(gotFuel-test.wantFuel) > epsilon || math.Abs(gotTotal-test.wantTotal) > epsilon {
				t.Errorf(
					"calculateTripCosts(%.2f, %.2f, %.2f) = %.2f, %.2f, want %.2f, %.2f",
					test.fuelLiters,
					test.fuelPrice,
					test.tolls,
					gotFuel,
					gotTotal,
					test.wantFuel,
					test.wantTotal,
				)
			}
		})
	}
}

func TestSplitCost(t *testing.T) {
	tests := []struct {
		name       string
		totalCost  float64
		passengers int
		driverPays bool
		wantPayingPerPerson float64
		wantPayingPassengers int
	}{
		{
			name:       "общая стоимость: 3500, пассажиров: 4, водитель платит",
			totalCost:  3500,
			passengers: 4,
			driverPays: true,
			wantPayingPerPerson: 875,
			wantPayingPassengers: 4,
		},
		{
			name:       "общая стоимость: 5000, пассажиров: 5, водитель платит",
			totalCost:  5000,
			passengers: 5,
			driverPays: true,
			wantPayingPerPerson: 1000,
			wantPayingPassengers: 5,
		},
		{
			name:       "общая стоимость: 1000, пассажиров: 1, водитель платит",
			totalCost:  1000,
			passengers: 1,
			driverPays: true,
			wantPayingPerPerson: 1000,
			wantPayingPassengers: 1,
		},
		{
			name:       "общая стоимость: 1000, пассажиров: 2, водитель не платит",
			totalCost:  1000,
			passengers: 2,
			driverPays: false,
			wantPayingPerPerson: 1000,
			wantPayingPassengers: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotPayingPerPerson, gotPayingPassengers := calculatePaymentPerPerson(test.totalCost, test.passengers, test.driverPays)

			const epsilon = 0.000001

			if math.Abs(gotPayingPerPerson-test.wantPayingPerPerson) > epsilon || gotPayingPassengers != test.wantPayingPassengers{
				t.Errorf(
					"calculatePaymentPerPerson(%.2f, %d, %t) = %.2f, %d want %.2f, %d",
					test.totalCost,
					test.passengers,
					test.driverPays,
					gotPayingPerPerson,
					gotPayingPassengers,
					test.wantPayingPerPerson,
					test.wantPayingPassengers,
				)
			}
		})
	}
}
