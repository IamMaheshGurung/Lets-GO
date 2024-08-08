//rock paper scissor

package main

import (
    "fmt"
    "time"
    "math/rand"
)

//Variable Declearation
var pl = fmt.Println
var sf = fmt.Scanf 
var response string
var user int


var comp int
var result bool




func main(){
     rand.Seed(time.Now().UnixNano())
     pl("Welcome To Rock paper Scissor game")
     pl("********* 🪨📃✂️  ***********")
     for{
         pl("Would you like to start game? (y/n)")
         sf("%s", &response)
         if response == "y" ||response == "Y" {
        
                pl("Press 1 for Rock 2 for paper and 3 for scissor")
                sf("%d", &user)
                switch user {
                case 1 :
                    pl("You have choosen Rock")
                case 2 :
                    pl("You have choosen Paper")
                case 3 :
                    pl("you have choosen Scissor")
                }
                
                comp = rand.Intn(2) + 1
                switch comp {
                case 1 :
                    pl("Computer have choosen Rock")
                case 2 :
                    pl("Computer have choosen Paper")
                case 3 :
                    pl("Computer have choosen Scissor")
                }
                if comp == user {
                    pl("It has been draw you need to try again")
                } else if comp == 3 && user == 1 {
                    
                    
                    
                    
                    
                    
                    pl("You are Winner")
                } else if comp == 1 && user == 3 {
                    pl("Computer Won")
                } else {
                    if comp > user {
                        pl("Computer Won")
                    } else {
                        pl("You are Winnner")
                    }
                }
                
                 
                
                } else if response != "y" {
             break
         }
     }
 


}
