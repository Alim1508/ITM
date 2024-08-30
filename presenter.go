package service

import (
	"io/ioutil"
	"strings"
)

type Presenter interface {
	Present([]string) error
}

type FilePresenter struct {
	FilePath string
}

func (p *FilePresenter) Present(lines []string) error {
	data := strings.Join(lines, "\n")
	return ioutil.WriteFile(p.FilePath, []byte(data), 0644)
}
