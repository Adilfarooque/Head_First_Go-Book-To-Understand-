package callingmethod

import (
	"fmt"
	"strings"
)

func Gorocks() {
	broken := "G# r#cks!"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)
}
