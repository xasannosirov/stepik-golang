package main

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
