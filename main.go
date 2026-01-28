package main

import (
	"fmt"
	"syscall"
	"time"
)

var (
	user32         = syscall.NewLazyDLL("user32.dll")
	mouseEventProc = user32.NewProc("mouse_event")
	clickSpeed     float32
)

const (
	MOUSEEVENTF_LEFTDOWN = 0x0002
	MOUSEEVENTF_LEFTUP   = 0x0004
)

func main() {
	fmt.Print("Введите скорость кликов в секунду: ")

	_, err := fmt.Scanln(&clickSpeed)
	if err != nil {
		clickSpeed = 0.5
	}

	delay := time.Duration(clickSpeed*1000) * time.Millisecond

	fmt.Println("🖱️  Автокликер (Ctrl+C для остановки)\n⏱️  Скорость:", clickSpeed, "кликов/сек")

	ticker := time.NewTicker(delay)
	defer ticker.Stop()

	for range ticker.C {
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
