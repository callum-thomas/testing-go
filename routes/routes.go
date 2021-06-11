package routes

import (
	"fmt"
)

type SomeStruct struct {
	Field1 int
	Field2 int
}

type Articles []*SomeStruct

func (Articles Articles) String() string {
	return "Hello!"
}

func TestingFunc() {
	fmt.Println("This is from the routes package.")

	testingAppendToSlice()
	testMutatingFunc()
}

func testingAppendToSlice() {
	fmt.Println("Testing appending to a slice.")

	Articles := Articles{
		{Field1: 1, Field2: 2}, {10, 20}, {0, 4}, {},
	}

	art := &SomeStruct{8, 4}
	Articles = append(Articles, art)

	fmt.Printf("Art: %v\n", *art)

	for _, art := range Articles {
		fmt.Printf("Article: %v\n", *art)
	}
}

func testMutatingFunc() {
	fmt.Println("Testing mutating a value passed to a func.")

	mutate := func(a SomeStruct) { a.Field1 = 1 }

	testStruct := SomeStruct{5, 5}
	fmt.Printf("Prior to mutation: %v\n", testStruct)
	mutate(testStruct)
	fmt.Printf("After mutation: %v\n", testStruct)
}
