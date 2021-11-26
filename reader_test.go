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
			// close(f.Content())
			// break
		}
		time.Sleep(time.Second)
	}
}

func TestRead(t *testing.T) {
	f1 := NewFileReader("a.txt")
	defer f1.Close()
	// f2 := NewFileRead("b.txt")
	go Print(f1)
	// go Print(f2)
	// go Read(f1)
	Read(f1)
}
