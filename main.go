package main // import "github.com/harvey93/pipeline"

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Pipeline")
	port := os.Getenv("PORT")
	fmt.Println(port)

}
