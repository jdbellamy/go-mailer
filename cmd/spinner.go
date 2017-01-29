package cmd

import (
	"github.com/tj/go-spin"
	"fmt"
	"time"
)

type WaitAnimation struct {
	Spinner *spin.Spinner
}

func NewSpinner() *WaitAnimation {
	return &WaitAnimation{
		Spinner: spin.New(),
	}
}

func (s *WaitAnimation) Spin() {
	for {
		fmt.Printf("\r  \033[36msending\033[m %s ", s.Spinner.Next())
		time.Sleep(100 * time.Millisecond)
	}
}

func (s *WaitAnimation) Stop() {
	s.Spinner.Reset()
}
