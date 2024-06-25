package mypackage

import (
    "bufio"
    "os"
)

type FileProducer struct {
    filePath string
}

func NewFileProducer(filePath string) *FileProducer {
    return &FileProducer{filePath: filePath}
}

func (fp *FileProducer) Produce() ([]string, error) {
    file, err := os.Open(fp.filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}
