package main

import "fmt"

type storyPage struct {
	text   string
	nextpg *storyPage
}

func (s *storyPage) playStort() {
	for s != nil {
		fmt.Println(s.text)
		s = s.nextpg
	}

}

func addPage(story *storyPage, text string) {
	newPage := &storyPage{text, nil}

	for story.nextpg != nil {
		story = story.nextpg
	}

	story.nextpg = newPage
}

func main() {
	p1 := storyPage{"page1", nil}
	p2 := storyPage{"page2", nil}
	p3 := storyPage{"page3", nil}

	p1.nextpg = &p2
	p2.nextpg = &p3

	addPage(&p3, "page4")
	p1.playStort()

}
