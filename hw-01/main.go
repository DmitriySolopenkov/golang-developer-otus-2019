package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func helloNow() error {
	ntpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка при получении времени с NTP сервера:", err)
		return err

	}
	currentTime := time.Now()
	fmt.Println("Текущее локальное время:", currentTime)
	fmt.Println("Точное время с использованием NTP:", ntpTime)

	return nil
}

func main() {
	helloNow()
}
