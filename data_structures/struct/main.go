package main

import  (
	"fmt"
	"sort"
)

type dude struct {
    Name string
    Age  int
}

// ByAge implements sort.Interface based on the Age field.
type ByAge []dude

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
    dudes := []dude{
        {"Andrea", 36},
        {"Cooler", 33}, 
        {"spiderman", 25},
        {"thor", 31},
        {"Picard", 56},
        {"Spock", 32},
    }
    sort.Sort(ByAge(dudes))
    fmt.Println(dudes) 
    
   item:=dude{"batman",42}
   dudes=append(dudes,item) 
   fmt.Println(dudes)
}
