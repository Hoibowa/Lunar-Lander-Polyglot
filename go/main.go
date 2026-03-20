package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 物理状態を構造体で管理
type Lander struct {
	Altitude float64
	Velocity float64
	Fuel     int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		l := Lander{Altitude: 100.0, Velocity: 0.0, Fuel: 50}
		gravity := 1.62
		dt := 1.0

		fmt.Println("\n" + strings.Repeat("=", 40))
		fmt.Println("--- LUNAR LANDER: GO MISSION START ---")
		fmt.Println(strings.Repeat("=", 40))

		for l.Altitude > 0 {
			fmt.Printf("高度: %6.2fm | 速度: %6.2fm/s | 燃料: %3d\n", l.Altitude, l.Velocity, l.Fuel)

			var thrust int
			for {
				fmt.Printf("噴射(0-10): ")
				if !scanner.Scan() {
					return
				}
				input := strings.TrimSpace(scanner.Text())
				if input == "" {
					thrust = 0
					break
				}

				val, err := strconv.Atoi(input)
				if err == nil && val >= 0 && val <= 10 {
					if val <= l.Fuel {
						thrust = val
						break
					}
					fmt.Printf("燃料不足です (残り: %d)\n", l.Fuel)
				} else {
					fmt.Println("0-10の数値を入力してください。")
				}
			}

			l.Fuel -= thrust
			accel := gravity - (float64(thrust) * 0.4)
			l.Velocity += accel * dt
			l.Altitude -= l.Velocity * dt
		}

		fmt.Printf("\n--- タッチダウン！ 最終速度: %.2f m/s ---\n", l.Velocity)
		if l.Velocity <= 3.0 {
			fmt.Println("【成功】ソフトランディング！")
		} else if l.Velocity <= 10.0 {
			fmt.Println("【警告】機体中破。")
		} else {
			fmt.Println("【失敗】月面にクレーターが誕生。")
		}

		fmt.Print("\nもう一度挑戦しますか？ (y/n): ")
		if !scanner.Scan() {
			break
		}
		if !strings.HasPrefix(strings.ToLower(scanner.Text()), "y") {
			break
		}
	}
}
