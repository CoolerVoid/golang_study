
package main

import "sort"
import "fmt"

func main() {

    // To create an empty map, you can use `make`:
    x := make(map[string]int)

    // Set key/value pairs using  `name[key] = val`
    x["A"] = 1337
    x["B"] = 4242
    x["C"] = 1234
    x["D"] = 100000
    x["E"] = 200000

    fmt.Println("map values:", x)

    // Get a value for a key with `name[key]`.
    v1 := x["A"]
    fmt.Println("v1: ", v1)

    // The builtin `len` returns the number of key/value
    fmt.Println("len:", len(x))

    // sort x map by high value
    fmt.Println("\nSort by high value\n")
    keys := make([]string, 0, len(x))
    for key := range x {
        keys = append(keys, key)
    }
    sort.Slice(keys, func(i, j int) bool { return x[keys[i]] > x[keys[j]] })

    for _, key := range keys {
        fmt.Printf("%s, %d\n", key, x[key])
    }
    // sort by string
    fmt.Println("\nPrint sort by letter \n")
    keys2 := make([]string, 0, len(x))
    for k := range x {
        keys2 = append(keys2, k)
    }
    sort.Strings(keys2)

    for _, k := range keys2 {
         fmt.Println(k, x[k])
    }

    // The builtin `delete` removes key/value from map
    delete(x, "B")
    fmt.Println("map before delete value:", x)
    // test value is empty
    _, prs := x["B"]
    fmt.Println("prs:", prs)

    // other example
    fib := map[string]int{"z": 1, "y": 2}
    fmt.Println("map:", fib)

    


}
