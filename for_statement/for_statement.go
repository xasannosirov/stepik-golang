package forstatement

import "fmt"

func ThisFor(n int){
	for i := 0; i < n; i++ {
		fmt.Println(i)
	} 
}