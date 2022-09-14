package frbl

import (
	"fmt"
	"testing"
	"time"
)

func Print(f FileReader) {
	for m := range f.Content() {
		fmt.Print(m)
	}
}

func Read(f FileReader) {
	for {
		if err := f.ReadLine(); err != nil {
			fmt.Println(err)
			break
		}
		time.Sleep(time.Second)
	}
}

func TestRead(t *testing.T) {
	f1 := NewFileReader("/tmp/a.txt")
	defer f1.Close()
	go Print(f1)
	Read(f1)
}
