package cli

import (
	"fmt"

	mydb "server/internal/db"
)

func displayMenu() {
	fmt.Println("--- Weekly ---")
	fmt.Println("1. Get all tasks")
	fmt.Println("2. Get one-day tasks")
	fmt.Println("3. Insert task")
	fmt.Println("4. Toggle day of task")
	fmt.Println("5. Delete task")
	fmt.Println("0. Exit")
	fmt.Println()
}

func menuDisplayer() {
	var optionDisplay int

	fmt.Println("--- Select type of Display ---")
	fmt.Println("1. JSON Display")
	fmt.Println("2. XML Display")
	fmt.Println("3. Default Display")
	fmt.Println()

	fmt.Print("> ")
	fmt.Scan(&optionDisplay)
	fmt.Println()

	switch optionDisplay {
	case 1:
		DisplayerJSON = true
	case 2:
		DisplayerXML = true
	}
}

func menuDaysOfWeek() {
	var optionDayOfWeek int

	fmt.Println("--- Select day of week ---")
	fmt.Println("1. Monday")
	fmt.Println("2. Tuesday")
	fmt.Println("3. Wednesday")
	fmt.Println("4. Thursday")
	fmt.Println("5. Friday")
	fmt.Println("6. Saturday")
	fmt.Println("7. Sunday")
	fmt.Println()

	fmt.Print("> ")
	fmt.Scan(&optionDayOfWeek)
	fmt.Println()

	switch optionDayOfWeek {
	case 1:
		Monday = true
	case 2:
		Tuesday = true
	case 3:
		Wednesday = true
	case 4:
		Thursday = true
	case 5:
		Friday = true
	case 6:
		Saturday = true
	case 7:
		Sunday = true
	}
}

// Crear mapa que contenga variables Monday y respectivos para luego optimizar utilizando el indice
func getDayOfWeek() (dayOfWeek mydb.DayOfWeek, err error) {
	switch {
	case Monday:
		dayOfWeek = mydb.Monday
	case Tuesday:
		dayOfWeek = mydb.Tuesday
	case Wednesday:
		dayOfWeek = mydb.Wednesday
	case Thursday:
		dayOfWeek = mydb.Thursday
	case Friday:
		dayOfWeek = mydb.Friday
	case Saturday:
		dayOfWeek = mydb.Saturday
	case Sunday:
		dayOfWeek = mydb.Sunday
	default:
		return "", fmt.Errorf("error needs at least one flag")
	}

	return dayOfWeek, nil
}
