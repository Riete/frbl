# File Read By Line With File In Writing Or File Rotated
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