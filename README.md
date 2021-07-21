tool for continuous reading file by line while file is still writing or is rotated 

```
f := NewFileRead("/path/to/file")
go func(f *FileRead) {
	for m := range f.Content {
		fmt.Print(m)
	}
}(f)
for {
	if err := f.ReadLine(); err != nil {
		fmt.Println(err)
		close(f.Content)
		break
	}
	time.Sleep(time.Second)
}
```