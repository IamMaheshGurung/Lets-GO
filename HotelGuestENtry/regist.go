package main



import(
    "fmt"
)


 var pl = fmt.Println
 var sf = fmt.Scanf

type details struct {
    firstname string
    lastname string
    age int
    country string
    roomName string
}
type OnsenRes struct {
    nameofPerson string
    numOfMen int 
    numOfWomen int
    hour int}


type Onsen interface {
    letsBook() string
}

func ( o OnsenRes) letsBook() string {
    return fmt.Sprintf("The family of %s with memeber of %d male and %d female will stay for %d hours", o.nameofPerson, o.numOfMen, o.numOfWomen, o.hour)
}
func OnsenResInfo ( person Onsen) {
    pl(person.letsBook())
}
func (d details) roomInfo() {
    pl(d.firstname, "Sir, you are alloted", d.roomName)
}

var roomnames= [10]string{"kadan", "lalaca", "yutowa", "hyaat", "nepal", "japan", "sakura", "atami", "hakone", "suin"}
var totalrooms = len(roomnames)
var BookingDetails []details
var counter = 0
var response string




//main function starts here

func main(){
     pl("**********Welcome to Tokyo Hotel***********")
     for {
         pl("We have", totalrooms - counter, "rooms available out of", totalrooms )
         pl("Would you like to book a room in our hotels")
         sf("%s", &response) 
        if response == "Yes" || response == "yes" {
            var info details

            
            pl("Thank you for choosing us")
            pl("Dear sir/maam, May  I have your First Name Please")
            sf("%s", &info.firstname)
            pl("May  I have your lastname Please")
            sf("%s", &info.lastname)
            pl("May  I have your age Please")
            sf("%d", &info.age)
            if info.age < 18 {
                panic("you are under age")
            }else if info.age > 100 {
                panic("Sorry sir your age is too high")
            }
            pl("May  I have your Country Name Please")
            sf("%s", &info.country)
            info.roomName = roomnames[counter]
            info.roomInfo()
            counter++

            BookingDetails = append(BookingDetails, info)


            pl("Your Booking has been confirmed")
            pl("First Name :", info.firstname)
            pl("Last Name :", info.lastname)
            pl("From :", info.country)
            pl("__________________________________")

            var osenResponse string

            pl("One Last Qn sir, would you like to Book onsen for that day sir")
            sf("%s", &osenResponse)
            if osenResponse == "yes" {
                var onsenInfo OnsenRes
                onsenInfo.nameofPerson = info.firstname
                
                pl("Number of Gentlemen: ")
                sf("%d", &onsenInfo.numOfMen)
                pl("Number of Ladies:")
                sf("%d", &onsenInfo.numOfWomen)
                onsenInfo.hour = 3

                 OnsenResInfo(onsenInfo)
             }
        }else if response != "yes" {
            pl(BookingDetails)
            pl("Thank you so much sir")
            break
        }
    }

    
}


