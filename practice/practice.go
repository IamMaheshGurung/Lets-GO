
package practice




import (
    "fmt"
    "errors"
    "math/rand"
)

func Rose(name string) (string, error) {
  
           if name == "" {
             return "", errors.New("I will not talk to you")
         }
    


    message := fmt.Sprintf(randomFormat(), name)

    return message, nil
        
}


func Hellos(names []string)(map[string]string, error) {
    messages := make(map[string]string)


    for _, name := range names{
        message, err := Rose(name)
        if err != nil {
            return nil, err
        }
        messages[name] = message
    } 
    return messages, nil
}

func randomFormat() string {
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you %v",
        "Wow its you, %v",
    }

    return formats[rand.Intn(len(formats))]
}
