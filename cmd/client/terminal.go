package client

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
)

func NewTerminal() (*Terminal, error) {
	if !terminal.IsTerminal(int(os.Stdout.Fd())) {
		return nil, errors.New("canno't access terminal")
	}

	return &Terminal{
		w:       os.Stdout,
		runtime: runtime.GOOS,
		m:       new(sync.RWMutex),
	}, nil
}

type Terminal struct {
	w       io.Writer
	runtime string
	m       *sync.RWMutex
}

func (t *Terminal) SetLoading(prefix string) func() {
	stopChan := make(chan bool)
	stop := func() {
		stopChan <- true
		close(stopChan)
	}
	go func() {
		w := wow.New(t.w, spin.Get(spin.Dots), prefix)
		w.Start()

		<-stopChan
		w.Stop()
		return
	}()

	return stop
}

func (t *Terminal) WaitInput(prefix string) (string, error) {
	fmt.Print(prefix)
	line, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.Trim(line, "\n"), nil
}

func (t *Terminal) Overwrite(s string) {
	t.Clear()
	t.Write(s)
}

func (t *Terminal) WriteF(s string, args ...interface{}) {
	t.Write(fmt.Sprintf(s, args...))
}

func (t *Terminal) Write(s string) {
	t.m.Lock()
	defer t.m.Unlock()
	if _, err := fmt.Fprint(t.w, s); err != nil {
		log.Print(err)
	}
}

func (t *Terminal) WriteLn(s string) {
	t.Write(fmt.Sprintf("%s\n", s))
}

func (t *Terminal) Clear() {
	switch t.runtime {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = t.w
		_ = cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = t.w
		_ = cmd.Run()
	}
}
