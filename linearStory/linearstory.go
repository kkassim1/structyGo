package main

import (
	"bufio"
	"fmt"
	"os"
)

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
	scanner := bufio.NewScanner(os.Stdin)
	p1 := storyPage{"page1", nil}
	scanner.Scan()
	p1.addPage(scanner.Text())
	p1.addPage("page3")

	p1.playStory()

}
