package repeatfunc

import "fmt"

func RepeatLine(line string, time int) {
	for i := 0; i < time; i++ {
		fmt.Println(line)
	}
}
