package main

import "fmt"

func dayOfTheWeek(number int) string {
	switch number {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Invalid day of the week"
	}
}

func dayOfTheWeek_2(number int) string {
	switch {
	case number == 1:
		return "Sunday"
	case number == 2:
		return "Monday"
	case number == 3:
		return "Tuesday"
	case number == 4:
		return "Wednesday"
	case number == 5:
		return "Thursday"
	case number == 6:
		return "Friday"
	case number == 7:
		return "Saturday"
	default:
		return "Invalid day of the week"
	}
}

func main() {
	fmt.Println("Switch")

	day := dayOfTheWeek(5)
	day_2 := dayOfTheWeek_2(5)

	fmt.Println(day)
	fmt.Println(day_2)

	fmt.Println(day == day_2)

}
