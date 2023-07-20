package main

import "fmt"

type storyPage struct {
	text   string
	nextpg *storyPage
}

func (s *storyPage) playStory() {
	for s != nil {
		fmt.Println(s.text)
		s = s.nextpg
	}

}

func (p *storyPage) addPage(text string) {
	newPage := &storyPage{text, nil}

	for p.nextpg != nil {
		p = p.nextpg
	}

	p.nextpg = newPage
}

func main() {
	p1 := storyPage{"page1", nil}
	p1.addPage("page2")
	p1.addPage("page3")

	p1.playStory()

}
