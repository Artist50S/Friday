package main

import "fmt"

func main() {
	month := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30}
	n := 6
	thirteen := 0
	allday := 0
	for _, y := range month {
		fmt.Println("y=", y)
		fmt.Println("curday=", (n+allday+13)%7)
		if (n+allday+13)%7 == 5 {
			fmt.Println("Ура!!! Пятницу 13")
			thirteen++
		}
		allday += y
		fmt.Println("allday=", allday)
	}
	fmt.Println(thirteen)
}
