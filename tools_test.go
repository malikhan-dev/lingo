package lingo

import (
	"fmt"
	"testing"
)

func TestCopyTo(t *testing.T) {
	var dest *[]ComplexObjectToSearch

	dest = DeepCopy(items)

	fmt.Println("=============================================")
	fmt.Println("printing values:")

	fmt.Println(items)

	fmt.Println(dest)

	fmt.Println("=============================================")

	fmt.Println("modifying source slice")

	dest = RemoveFirstByPredicate(*dest, func(search ComplexObjectToSearch) bool {
		return search.Name == "Jane"
	})

	fmt.Println("=============================================")
	fmt.Println("printing values:")

	fmt.Println(dest)

	fmt.Println(items)
}

func TestToSlice(t *testing.T) {

	var m map[int]ComplexObjectToSearch

	m = make(map[int]ComplexObjectToSearch, 2)

	m[1] = ComplexObjectToSearch{
		Name: "joe",
		Age:  20,
		Id:   1,
		Flag: false,
	}
	m[2] = ComplexObjectToSearch{
		Name: "john",
		Age:  27,
		Id:   2,
		Flag: true,
	}

	var dest *[]ComplexObjectToSearch

	dest = ToSlice(m)

	fmt.Println(*dest)

	fmt.Println(m)
}
