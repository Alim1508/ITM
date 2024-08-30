package service

import (
	"io/ioutil"
	"strings"
)

type Producer interface {
	Produce() ([]string, error)
}

type FileProducer struct {
	FilePath string
}

func (p *FileProducer) Produce() ([]string, error) {
	data, err := ioutil.ReadFile(p.FilePath)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	return lines, nil
}
