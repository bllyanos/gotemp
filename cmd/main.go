package main

import (
	"fmt"
	"os"

	"github.com/bllyanos/gotemp"
)

func main() {
	tee, err := gotemp.New("examples")
	if err != nil {
		fmt.Printf("Error creating template engine: %v\n", err)
		return
	}
	
	err = tee.RenderPage(os.Stdout, "app_layout", "home/index.html", nil)
	if err != nil {
		fmt.Printf("Error rendering page: %v\n", err)
		return
	}
}
