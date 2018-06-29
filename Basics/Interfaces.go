package main

import "fmt"


type human interface{
    sayHello() string    
}

type man struct{
    greeting string
}

type woman struct{
    greeting string
}

func (m man) sayHello() string{
    return m.greeting
}

func (w woman) sayHello() string{
    return w.greeting
}


func printGreeting(h human){
    fmt.Println(h.sayHello())
}

func main(){
    tom := man{greeting:"klk"}
    laura := woman{"Heeerrr"}
    
    printGreeting(laura)
    printGreeting(tom)
}
