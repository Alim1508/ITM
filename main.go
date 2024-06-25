package main

import (
    "flag"
    "log"
    "mypackage"
)

func main() {
    inputFilePath := flag.String("input", "", "Path to the input file")
    outputFilePath := flag.String("output", "output.txt", "Path to the output file")
    flag.Parse()

    if *inputFilePath == "" {
        log.Fatal("Input file path is required")
    }

    producer := mypackage.NewFileProducer(*inputFilePath)
    presenter := mypackage.NewFilePresenter(*outputFilePath)

    service := mypackage.NewService(producer, presenter)

    if err := service.Run(); err != nil {
        log.Fatalf("Service run failed: %v", err)
    }
}
