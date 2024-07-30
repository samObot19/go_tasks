package main

import (
	"fmt"
	grading "GradeCalculator/calculate"
)



func main() {
	var studentName string
	var numberOfSubjects int
	var courseAndGradePair = make(map[string]string)

	fmt.Println("Enter your name:")
	fmt.Scan(&studentName)
	fmt.Println("Enter the number of subjects:")
	fmt.Scan(&numberOfSubjects)
	fmt.Println("Enter the subject and the score, separated by a space:")

	for i := 0; i < numberOfSubjects; i++ {
		var subject string
		var score float64
		fmt.Scan(&subject, &score)
		err, grade := grading.Grade(score)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		courseAndGradePair[subject] = grade
	}

	fmt.Println("\nGrade Report:")
	fmt.Println("+--------------+-------+")
	fmt.Println("| Subject      | Grade |")
	fmt.Println("+--------------+-------+")
	for subject, grade := range courseAndGradePair {
		fmt.Printf("| %-12s | %-6s |\n", subject, grade)
	}
	fmt.Println("+--------------+-------+")
}

