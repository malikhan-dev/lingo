package Thor

import (
	"fmt"
	"testing"
)

type ComplexObjectToSearch struct {
	Name string
	Age  int
	Id   int
	Flag bool
}

var items []ComplexObjectToSearch

func LoadLargeData() {
	randFlag := false
	for i := 0; i < 50000000; i++ {

		items = append(items, ComplexObjectToSearch{
			Name: "Jane",
			Flag: randFlag,
			Id:   i,
			Age:  i,
		})
		randFlag = !randFlag
	}
}
func init() {
	LoadLargeData()
}

func TestQueryEngine(t *testing.T) {

	result := From(items).Where(func(search ComplexObjectToSearch) bool {
		return search.Name == "Jane" && search.Flag == false
	}).Collect()

	result2 := From(result).Any(func(search ComplexObjectToSearch) bool {
		return (search.Name != "Jane") || (search.Flag != false)
	}).Assert()

	if result2 {
		t.Error("result should be false")
	}

}

func TestGroupByNew(t *testing.T) {

	res :=
		Collect(
			Group[bool, ComplexObjectToSearch](

				From(items).Where(func(search ComplexObjectToSearch) bool {

					return search.Age > 20

				}),
				func(item ComplexObjectToSearch) bool {
					return item.Flag
				}))

	fmt.Println(res.Items[false][1])
	fmt.Println(res.Items[true][1])

	fmt.Println("==========================================================")
	fmt.Println("==========================================================")
	fmt.Println("==========================================================")

}
