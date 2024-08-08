package main 



import (
    "fmt"
    "log"
    "example.com/practice"
)


func main(){

    log.SetPrefix("Argument: ")
    log.SetFlags(0)
    var names []string
    var userChoice string
    var response string
    for{
    fmt.Println("Enter a name: ")
    fmt.Scanf("%s", &userChoice)
    names = append(names, userChoice)


    fmt.Println("Do you like to add more?")
    fmt.Scanf("%s", &response)

    if response != "yes"{
    break
    }

}

    
    message, err := practice.Hellos(names)

    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println(message)

}
