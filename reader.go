package frbl

import (
	"bufio"
	"io"
	"os"
)

type FileReader interface {
	ReadLine() error
	Content() chan string
	Close()
}

type file struct {
	path    string
	offset  int64
	file    *os.File
	content chan string
}

func NewFileReader(path string) FileReader {
	return &file{path: path, offset: offsetGet(path), content: make(chan string)}
}

func (f *file) open() error {
	var err error
	f.file, err = os.Open(f.path)
	return err
}

func (f *file) isRotated() bool {
	if end, err := f.file.Seek(0, io.SeekEnd); err != nil {
		return false
	} else {
		return f.offset > end
	}
}

func (f *file) setOffset() error {
	var err error
	if f.offset, err = f.file.Seek(0, io.SeekCurrent); err == nil {
		return offsetUpdate(f.path, f.offset)
	}
	return err
}

func (f *file) seek() error {
	if f.offset > 0 && f.isRotated() {
		f.offset = 0
	}
	_, err := f.file.Seek(f.offset, io.SeekStart)
	return err

}

func (f *file) ReadLine() error {
	if err := f.open(); err != nil {
		return err
	}
	defer f.file.Close()
	if err := f.seek(); err != nil {
		return err
	}
	r := bufio.NewReader(f.file)
	for {
		data, err := r.ReadString('\n')
		f.content <- data
		if err == io.EOF {
			return f.setOffset()
		}
	}
}

func (f file) Content() chan string {
	return f.content
}

func (f *file) Close() {
	close(f.content)
}
