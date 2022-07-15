package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

func (student Student) getAge() int {
	return student.age
}

//Must use pointer to make the 'this' student pass by reference
func (student *Student) setAge(age int) {
	student.age = age
}

func (student Student) getAVGGrade() float32 {
	result := 0
	for _, grade := range student.grades {
		result += grade
	}

	return float32(result) / float32(len(student.grades))
}

func (student Student) getMaxGrade() (index, value int) {
	index = -1
	value = -1
	if 0 == len(student.grades) {
		return
	}

	index = 0
	value = student.grades[index]

	for i, v := range student.grades {
		if value < v {
			value = v
			index = i
		}
	}

	return
}

func main() {
	var Tuấn = Student{"Tuấn", []int{1, 2, 3, 5}, 17}
	fmt.Println(Tuấn, Tuấn.getAge())
	Tuấn.age = 18
	Tuấn.setAge(123)
	fmt.Println(Tuấn, Tuấn.getAge(), Tuấn.getAVGGrade())
	fmt.Println(Tuấn.getMaxGrade())
}
