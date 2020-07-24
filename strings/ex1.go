// just another strings example in golang
package main
 
import (
    "fmt"
    "strings"
    "bytes"
    "regexp"
)
 
func main() {
    // concatenate string
    s :="Test "
    s += "and other"
    fmt.Println(s)  
    // second way to concatenate
    v := fmt.Sprintf("Result = %s\n", s)
    fmt.Println(v)

    // using builder to concatenate
    var sb strings.Builder
    sb.WriteString("Other ")
    sb.WriteString("test of strings\n")
    fmt.Println(sb.String())

    // using buffer to concatenate
    var x bytes.Buffer
    x.WriteString("Second ")
    x.WriteString("other test\n") 
    fmt.Println(x.String()) 

    // replace string in quote
    z := strings.Replace(s, "Test", "Production", -1) // -1 replace all ocurrences
    fmt.Println(z)

    // parser from scratch
    var str = "<example>quote to parser here!</example> \n"
    var tmp = ""
    var parse = false
    var result = ""

    for i := 0; i < len(str); i++ {
	tmp+=string(str[i])  // if use buffer or builder can be better...
	if (parse) {
		if(str[i] == '<' && str[i+1] == '/' && str[i+2] == 'e'){ 
			parse=false
			break
		}
		if (parse) {
			result+=string(str[i])
		}
	}

	if (strings.Contains(tmp,"</example>")) {
		parse=false
		break
	}

	if (strings.Contains(tmp,"<example>")) {
		parse=true
		tmp=""
	}
    }
    fmt.Printf("Parsed value is: \"%s\"\n",result)

    // test regex on golang
    m := regexp.MustCompile(`Bingo`)
    fmt.Println(m.FindString("just another test"))
    fmt.Println(m.FindString("Ops Bingo here"))
    fmt.Println(m.FindString("I donmt know this"))
    fmt.Println(m.FindString("ops Bingo here again"))
    m2 := regexp.MustCompile(`nintendo.?`) 
    fmt.Println("The game with nintendo here! ")
    r := m2.FindStringIndex("The game with nintendo here! ") 
    fmt.Printf("Found:regex with index value: %d\n", r) 
    
}
