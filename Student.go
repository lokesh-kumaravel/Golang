package main

import "fmt"

func main() {
	name := "Lokesh K"
	age := 22
	subjects := []string{"Tamil", "English", "Maths", "Physics", "Chemistry", "Computer Science"}
	grades := []int{90, 85, 88, 92, 95, 89}
	average := findAverage(grades)
	status := true

	fmt.Println("Student Details:")
	fmt.Printf("Name : %v\n", name)
	fmt.Printf("Age : %v\n", age)
	fmt.Printf("Average : %v\n", average)
	fmt.Printf("Status : %v\n", status)
	fmt.Printf("Student marks on each subject : \n")

	for i := 0; i < len(subjects); i++ {
		fmt.Printf("%s : %d\n", subjects[i], grades[i])
	}
}

func findAverage(grades []int) float64 {
	var sum int
	for _, grade := range grades {
		sum += grade
	}
	average := float64(sum) / float64(len(grades))
	return average
}
