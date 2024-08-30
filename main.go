package main

import (
	"flag"
	"log"
	"github.com/Alim/ITM/service/service"
)

func main() {
	inputFile := flag.String("input", "input.txt", "Path to the input file")
	outputFile := flag.String("output", "output.txt", "Path to the output file")
	flag.Parse()

	producer := &service.FileProducer{FilePath: *inputFile}
	presenter := &service.FilePresenter{FilePath: *outputFile}

	svc := service.NewService(producer, presenter)

	if err := svc.Run(); err != nil {
		log.Fatalf("Service run failed: %v", err)
	}
}
