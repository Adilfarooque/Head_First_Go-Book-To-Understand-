package function

import()

func repeatLine(line string, time int){
	for i := 0; i < time; i++ {
		fmt.Println(line)
	}
}


func main(){
	repeatLine("Except Allah swt :)",3)
}