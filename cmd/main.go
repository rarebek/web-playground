package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("started at: %s", time.Now().Format(time.DateTime))
}
