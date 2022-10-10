package sub

import "fmt"

type Statistics struct {
	Transmitted int
	Received    int
	Loss        int
}

func (s *Statistics) ShowResult() {
	fmt.Println()
	fmt.Println("--- statistics ---")
	fmt.Printf("transmitted=%d::received=%d::loss=%d\n", s.Transmitted, s.Received, s.Loss)
}

func NewStatistics() *Statistics {
	return &Statistics{
		Transmitted: 0,
		Received:    0,
		Loss:        0,
	}
}
