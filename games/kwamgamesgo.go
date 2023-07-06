package main

import "fmt"

// todo insert and delete page

type storypage struct {
	c    string
	next *storypage
}

func playstory(p *storypage) {

	if p == nil {
		return
	}

	fmt.Println(p.c)
	playstory(p.next)
}

func main() {

	f := storypage{"yoyo", nil}
	e := storypage{"e11", nil}
	t := storypage{"t1", nil}

	f.next = &e
	e.next = &t

	playstory(&f)

}
