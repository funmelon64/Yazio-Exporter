package main

import (
	"YazioExporter/cmd"
	"log"
	"os"
)

func main() {
	err := cmd.Init().Run(os.Args)
	if err != nil {
		log.Panicf("fail to start app: %v", err)
	}
}
