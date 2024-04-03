package main

import "fmt"

func main() {

	var courseName []string
	courseName = []string{"Jave", "Python"}
	fmt.Println(courseName)

	//add data append
	courseName = append(courseName, "C", "C#", "HTML", "CSS", "JavaScript")
	fmt.Println(courseName)

	//delete data
	courseWeb := courseName[4:7]
	fmt.Println(courseWeb)

	courseWeb = courseName[:4]
	fmt.Println(courseWeb)
}
