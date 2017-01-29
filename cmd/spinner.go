package cmd

import (
	"github.com/tj/go-spin"
	"fmt"
	"time"
)

type WaitAnimation struct {
	Spinner *spin.Spinner
	stop	bool
}

func Spinner() *WaitAnimation {
	s := &WaitAnimation{}
	s.Spinner = spin.New()
	go func(s *WaitAnimation) {
		for s.stop == false {
			fmt.Printf("\r  \033[36m\033[m %s ", s.Spinner.Next())
			time.Sleep(100 * time.Millisecond)
		}
	}(s)
	return s
}

func (a *WaitAnimation) Stop() {
	a.stop = true
}
