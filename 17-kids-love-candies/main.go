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
		knText, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		knText = strings.TrimRight(knText, "\n")
		kn := strings.Split(knText, " ")
		K, _ := strconv.Atoi(kn[1])

		candiesText, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		candiesText = strings.TrimRight(candiesText, "\n")
		candies := strings.Split(candiesText, " ")

		kidsCount := 0
		for j := 0; j < len(candies); j++ {
			cookieCount, _ := strconv.Atoi(candies[j])
			kidsCount += (cookieCount / K)
		}

		fmt.Println(kidsCount)
	}
}
