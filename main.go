package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"time"
)

var (
	user32         = syscall.NewLazyDLL("user32.dll")
	mouseEventProc = user32.NewProc("mouse_event")
	clickSpeed     = 0.5
	closeChan      = make(chan int)
)

const (
	MOUSEEVENTF_LEFTDOWN = 0x0002
	MOUSEEVENTF_LEFTUP   = 0x0004
)

func main() {
	fmt.Print("Введите скорость кликов в секунду: ")
	fmt.Scanln(&clickSpeed)

	delay := time.Duration(clickSpeed*1000) * time.Millisecond

	clearConsole()
	fmt.Println("Старт автокликера\nСкорость:", clickSpeed, "клик(ов)/сек")

	ticker := time.NewTicker(delay)
	defer ticker.Stop()

	go func(closeChan chan int) {
		defer close(closeChan)

		for range ticker.C {
			select {
			case <-closeChan:
				break
			default:
				mouseEventProc.Call(
					uintptr(MOUSEEVENTF_LEFTDOWN),
					0, 0, 0, 0,
				)
				time.Sleep(5 * time.Millisecond)
				mouseEventProc.Call(
					uintptr(MOUSEEVENTF_LEFTUP),
					0, 0, 0, 0,
				)
			}
		}
	}(closeChan)

	fmt.Scanln()
	closeChan <- 0
}

func clearConsole() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
