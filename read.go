package frbl

import (
	"bufio"
	"io"
	"os"
)

type FileRead struct {
	Path    string
	Offset  int64
	File    *os.File
	Content chan string
}

func NewFileRead(path string) *FileRead {
	content := make(chan string)
	return &FileRead{Path: path, Offset: OffsetGet(path), Content: content}
}

func (fr *FileRead) OpenFile() error {
	if file, err := os.Open(fr.Path); err != nil {
		return err
	} else {
		fr.File = file
		return nil
	}
}

func (fr *FileRead) IsFileRotated() error {
	if fr.Offset == 0 {
		return nil
	}
	if end, err := fr.File.Seek(0, io.SeekEnd); err != nil {
		return err
	} else {
		if fr.Offset > end {
			fr.Offset = 0
		}
		return nil
	}
}

func (fr *FileRead) SefOffset() error {
	if offset, err := fr.File.Seek(0, io.SeekCurrent); err != nil {
		return err
	} else {
		fr.Offset = offset
		return OffsetUpdate(fr.Path, offset)
	}
}

func (fr *FileRead) Seek() error {
	_, err := fr.File.Seek(fr.Offset, io.SeekStart)
	return err

}

func (fr *FileRead) ReadLine() error {
	if err := fr.OpenFile(); err != nil {
		return err
	}
	defer fr.File.Close()
	if err := fr.IsFileRotated(); err != nil {
		return err
	}
	if err := fr.Seek(); err != nil {
		return err
	}
	r := bufio.NewReader(fr.File)
	for {
		data, err := r.ReadBytes('\n')
		fr.Content <- string(data)
		if err == io.EOF {
			return fr.SefOffset()
		}
	}
}
