package sysfuncs

import (
	"machine"
	"tinygo.org/x/tinyfs/littlefs"
	"tinygo.org/x/tinyfs"
)

// this file is mostly a wrapper of the tinyfs and littlefs packages




type Filesystem struct {
	internal *littlefs.LFS
}

func (fs *Filesystem) Open (filePath string) File {
	f, err := fs.internal.Open(filePath)
	if (err != nil) {
		println("failed")
	}
	file := File{internal: f}
	return file
}




type File struct {
	internal tinyfs.File
	isVirt bool
}



func (f *File) Read ()  {
	f.internal.Read()
}


func (f *File) Close () error {
	err := f.internal.Close()
}