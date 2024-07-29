package main

import (
	"fmt"
	"math/rand"
	"time"
)

var pl = fmt.Println
var sf = fmt.Scanf

type Detail struct {
	Name    string
	Age     int
	Country string
	Code    int
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var name1 string
	var age int
	var country string
	var response string
	var files []Detail

	for {
		pl("Please Fill Detail")
		pl("What is your Name? ")
		sf("%s", &name1)
		randCode := 100 + rand.Intn(900)
		pl(name1, " your alloted code is ", randCode)
		pl("please Enter your age: ")
		sf("%d", &age)
		pl("please Enter your Country: ")
		sf("%s", &country)

		detail := Detail{
			Name:    name1,
			Age:     age,
			Country: country,
			Code:    randCode,
		}

		files = append(files, detail)
		pl("Here is Your Details sir..")

		pl("Do you like to continue it ? (y/n)")
		sf("%s", &response)
		if response == "n" {
			break
		}
		pl("Thank you so much")

	}
	pl("Here are your all details sir")
	pl(files)

}
