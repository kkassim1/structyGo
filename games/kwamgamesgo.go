package main

import "fmt"

// todo insert and delete page

type storypage struct {
	c    string
	next *storypage
}

func (p *storypage) playstory() {

	for p != nil {
		fmt.Println(p.c)
		p = p.next
	}
}

func main() {

	f := storypage{"yoyo", nil}
	e := storypage{"e11", nil}
	t := storypage{"t1", nil}

	f.next = &e
	e.next = &t

	f.playstory()

}
