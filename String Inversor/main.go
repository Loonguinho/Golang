package main

import(
    "fmt"
    "bufio"
    "os"
    "strings"
);

func invertStr(array []string) []string{
    length := len(array)

    for i :=0; i< length / 2; i++ {

        array[i], array[length-1-i] = array[length-1-i], array[i]
    
    }
     
    return array
}


func main() {
    var invertedStr []string
    reader := bufio.NewReader(os.Stdin)
    
    fmt.Println("What's your string?")
    input, err := reader.ReadString('\n');
    
    if err != nil{   
        fmt.Println("Error:", err)
    }

    fmt.Println("Your string says:", input)
    input = strings.TrimSpace(input)
    invertedStr = strings.Split(input, "")

    result := invertStr(invertedStr)
    fmt.Println("Your inverted string:", result)

}

