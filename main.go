package main

import (
	"fmt"
)

func grade(score float64) string {

	switch {
		case score >= 90 && score < 101 :
			return "A+"
		case score >= 85 && score < 90:
			return "A"
	    case score >= 80 && score < 85:
			return "A-"
		case score >= 75 && score < 80:
			return "B+"
		case score >= 70 && score < 75:
			return "B"
		case score >= 65 && score < 70:
			return "B-"
	    case score >= 60 && score < 66:
			return "C+"
		case score >= 55 && score < 60:
			return "C"
		case score >= 50 && score < 55:
			return "D"
	    case score < 50 && score >= 0 :
			return "F"	
		default :
			return "invalid"

	    										 
	}
}

func main() {
	var name string
	fmt.Println("Enter your name:")
	fmt.Scanf("%s\n", &name)
	fmt.Println("Hello,", name)

	var numberOfCourses int

	fmt.Println("How many courses did you take? ")
	fmt.Scanf("%d", &numberOfCourses)
	if numberOfCourses <= 0 {
		fmt.Println("You entered the wrong number of courses. It must be a positive number greater than zero" )
		
	}
	var total float64
	dictionary := make(map[string]string)
	for i := 1 ; i <= numberOfCourses ; i++ {
		var course string
		var score float64

		fmt.Println("Name of the course")
		fmt.Scan(&course) 

		fmt.Println("your score")
		fmt.Scanf("%f", &score)
		
		grade := grade(score)
		dictionary[course] = grade

		if grade != "invalid" {
			total += float64(score)
		}
		
	}

	for key , value := range dictionary {

		fmt.Printf("For the course : %v , you scored %v \n" , key , value)
	}
	average := total / float64(numberOfCourses)
	avg_grade := grade(average)
	fmt.Printf("The average of %v scores is %v with average grade of %v" , name , average , avg_grade)

}
