package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("an error occurred")
	wrappedErr := fmt.Errorf("additional context: %w", err)
	fmt.Println(wrappedErr)
}

/*sample output
additional context: an error occurred
*/
