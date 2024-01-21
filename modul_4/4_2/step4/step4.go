package main
import (
	"fmt"
	"log"
	"net"
	"strings"
	// "time"
)

func main(){
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for i := 0; i < 3; i++ {
		res := make([]byte, 1024)
		n, err := conn.Read(res)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(strings.ToUpper(string(res[:n])))
	}
}