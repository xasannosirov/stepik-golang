package main

import "fmt"

// inoutStream channelda bir xil ma'lumotlar ketma-ket kelsa
// shulardan bittasini outputStream channelga yozadi
func removeDuplicates(inputStream chan string, outputStream chan string) {
	defer close(outputStream)
	var temp string
	for {
		value, ok := <-inputStream
		if !ok {
			break
		}
		if value != temp {
			outputStream <- value
			temp = value
		}
	}
}

func main() {
	inputStream := make(chan string, 20)
	outputStream := make(chan string, 15)

	inputStream <- "A"
	inputStream <- "B"
	inputStream <- "C"
	inputStream <- "D"
	inputStream <- "D"
	inputStream <- "E"
	inputStream <- "E"
	inputStream <- "F"
	close(inputStream)

	removeDuplicates(inputStream, outputStream)

	fmt.Println(<-outputStream) // A
	fmt.Println(<-outputStream) // B
	fmt.Println(<-outputStream) // C
	fmt.Println(<-outputStream) // D
	fmt.Println(<-outputStream) // E
	fmt.Println(<-outputStream) // F
	fmt.Println(<-outputStream) //
}
