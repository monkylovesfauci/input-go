package input

import (
	"bufio"
	"fmt"
	"os"
)

func Print(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg)
	scanner.Scan()
	input := scanner.Text()
	return input
}
