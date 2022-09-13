package main

import (
	"mathtime"
	"os"
	"time"
)

func main() {
	t := time.Now()
	mathtime.SVGWriter(os.Stdout, t)
}
