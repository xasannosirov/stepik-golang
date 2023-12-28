package concurency

func DddDataChan(c chan int, data int){
	c <- data
}