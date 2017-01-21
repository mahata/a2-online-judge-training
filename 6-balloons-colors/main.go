package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n') // Discard first line

	for i := 0; ; i++ {
		nxyText, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		nxyText = strings.TrimRight(nxyText, "\n")
		nxy := strings.Split(nxyText, " ")

		N, _ := strconv.Atoi(nxy[0])
		X := nxy[1]
		Y := nxy[2]

		balloonsText, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		balloonsText = strings.TrimRight(balloonsText, "\n")
		balloons := strings.Split(balloonsText, " ")

		if balloons[0] == X && balloons[N-1] == Y {
			fmt.Println("BOTH")
		} else if balloons[0] == X && balloons[N-1] != Y {
			fmt.Println("EASY")
		} else if balloons[0] != X && balloons[N-1] == Y {
			fmt.Println("HARD")
		} else {
			fmt.Println("OKAY")
		}
	}
}
