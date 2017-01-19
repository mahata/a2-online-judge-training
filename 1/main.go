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
		text, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		text = strings.TrimRight(text, "\n")
		s := strings.Split(text, " ")
		if len(s) == 0 {
			break
		}

		leftVal, err := strconv.Atoi(s[0])
		if err != nil {
			fmt.Println(err)
		}

		rightVal, err := strconv.Atoi(s[1])
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("%d\n", leftVal+rightVal)
	}
}
