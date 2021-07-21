package frbl

import (
	"bufio"
	"io"
	"os"
)

type FileRead struct {
	Path    string
	Current int64
	File    *os.File
	Content chan string
}

func NewFileRead(path string) *FileRead {
	content := make(chan string)
	return &FileRead{Path: path, Current: 0, Content: content}
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
	if fr.Current == 0 {
		return nil
	}
	if end, err := fr.File.Seek(0, io.SeekEnd); err != nil {
		return err
	} else {
		if fr.Current > end {
			fr.Current = 0
		}
		return nil
	}
}

func (fr *FileRead) SetCurrent() error {
	if current, err := fr.File.Seek(0, io.SeekCurrent); err != nil {
		return err
	} else {
		fr.Current = current
		return nil
	}
}

func (fr *FileRead) Seek() error {
	_, err := fr.File.Seek(fr.Current, io.SeekStart)
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
		if err != nil {
			if err == io.EOF {
				if err := fr.SetCurrent(); err != nil {
					return err
				}
				return nil
			}
			return err
		}
		return nil
	}
}
