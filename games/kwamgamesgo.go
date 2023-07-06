package main

import "fmt"

// todo insert and delete page

type storypage struct {
	text string
	next *storypage
}

func (p *storypage) playstory() {

	for p != nil {
		fmt.Println(p.text)
		p = p.next
	}
}

func (p *storypage) addtoend(textToBeAddedToPage string) {

	for p.next != nil {
		p = p.next
	}
	p.next = &storypage{textToBeAddedToPage, nil}

}

func main() {

	f := storypage{"page 1", nil}
	f.addtoend("page 2")
	f.addtoend("page 3")

	f.playstory()

}
