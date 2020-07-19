
package main

//import "strconv"
import "sort"
import "strings"
import "fmt"

type notebook struct {
 name string
 cell string
 age  int
}

func main() {

    // Create map following struct of notebook
    contacts := make(map[int]notebook)

    // insert data following struct
    contacts[0] = notebook{ "spiderman", "55 909495976", 34}
    contacts[1] = notebook{ "hulk", "55 9088887876", 49}
    contacts[2] = notebook{ "thor", "55 9022334444", 38}
    contacts[3] = notebook{ "ironman", "55 90112334", 47}
    contacts[4] = notebook{ "vision", "55 902333414", 38}
    contacts[5] = notebook{ "hawkie", "55 9032423", 40}

    // inter data show each element
    fmt.Println(contacts)
    show_contacts(contacts)

    // delete element hawkie
    contacts=delete_per_name(contacts,"hawkie")
    show_contacts(contacts)

    // find Vision and replace to swamp thing
    contacts=change_name(contacts, "vision","Swamp thing")
    show_contacts(contacts)

    // Sort By Age
    contacts=sort_contacts_by_age(contacts)
    show_contacts(contacts)
     
}




func show_contacts( input map[int]notebook) {
	fmt.Println("\nShow data: \n")
    	for position, notebook :=  range input {
		fmt.Printf("Values position %d and values name %s age %d \n",position, notebook.name, notebook.age)
	}
}



func delete_per_name( input map[int]notebook, name string) map[int]notebook {
    	for position, note :=  range input {
                if (strings.Contains(note.name,name) == true) {
			delete(input,position)
		}
	
	}
	return input
}

func change_name( input map[int]notebook, name string, replace string) map[int]notebook {
    	for position, note :=  range input {
                if (strings.Contains(note.name,name) == true) {
			input[position]=notebook{ replace,note.cell,note.age }
		}
	
	}
	return input
}

type ByAge map[int]notebook
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func sort_contacts_by_age( input map[int]notebook) map[int]notebook {
 	//fmt.Printf("Age test %d", input[0].age )
	sort.Sort(ByAge(input))
	return input
}
