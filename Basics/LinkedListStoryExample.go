package main

import (
    "fmt"
    "bufio"
    "os"

)

type storyPage struct{
    text string
    nextPage *storyPage
}

func playStory(page *storyPage){
    for page != nil{
    fmt.Println(page.text)
        page = page.nextPage
    }
}

func (p *storyPage) playStory2(){
    for p != nil{
    fmt.Println(p.text)
        p = p.nextPage
    }
}

func (p *storyPage) addToEnd(text string){
    pageToAdd :=  &storyPage{text, nil}
    for p.nextPage != nil {
        p = p.nextPage
    }
    p.nextPage = pageToAdd
}

func (p *storyPage) addAfter(text string){
    newPage := &storyPage{text,p.nextPage}
    p.nextPage = newPage
}

func ()

func main() {
   // scanner := bufio.NewScanner(os.Stdin)
    
    page1 := storyPage{"Yo me llamo D", nil}
    page2 := storyPage{"Mi nombre es Damani",nil}
    page3 := storyPage{"Lucero en la cama", nil}
    page4 := storyPage{"Sabes que la comi", nil}
    
    page1.nextPage = &page2
    page2.nextPage = &page3
    page3.nextPage = &page4
    
    p1 := storyPage{"Tu ere malo y masochista",nil}
    p1.addToEnd("Jangueando con to lo artiti")
    p1.addToEnd("Yo que soy malocorita")
    p1.addToEnd("Bucando una menorcita")
    
    p1.nextPage.nextPage.addAfter("I am Ceky")
    
    //playStory(&page1)
    p1.playStory2()
    
    
}
