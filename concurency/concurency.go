package main

func AddDataChan(c chan int, data int){
	c <- data
}

func main(){
	c := make(chan int, 2)
	AddDataChan(c, 2)
}