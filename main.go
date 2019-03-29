package main

import (
	"fmt"
	"playground/faninfanout"
	"playground/pipeline"
)

func main() {
	fmt.Println("Runnning pipeline")
	pipeline.Execute()
	fmt.Println("End of runnning pipeline")

	fmt.Println("Runnning faninfanout")
	faninfanout.Execute()
	fmt.Println("End of runnning faninfanout")
}
