package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type jsonInput struct {
	Students []struct {
		Rating []float64
	}
}

type jsonOutput struct {
	Average float64
}

func main() {
	var s jsonInput
	data, _ := io.ReadAll(os.Stdin)
	if err := json.Unmarshal(data, &s); err != nil {
		fmt.Println(err)
		return
	}
	average := 0.0
	for i := 0; i < len(s.Students); i++ {
		average += float64(len(s.Students[i].Rating))
	}
	s_out := jsonOutput{Average: average / float64(len(s.Students))}
	data_output, _ := json.MarshalIndent(s_out, "", "    ")
	fmt.Printf("%s", data_output)
}
