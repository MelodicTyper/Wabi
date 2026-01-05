package sysfuncs

import (
	//"machine"

	"os"

	"tinygo.org/x/tinyfs"
	"tinygo.org/x/tinyfs/littlefs"
)

// this file is mostly a wrapper of the tinyfs and littlefs packages currently

type Filesystem struct {
	Internal *littlefs.LFS
}

func (fs *Filesystem) Open(filePath string) File {
	f, err := fs.Internal.Open(filePath)
	if err != nil {
		panic("Failed to open file: " + filePath + " " + err.Error())
	}
	file := File{internal: f, isVirt: false}
	return file
}

func (fs *Filesystem) OpenFile(filePath string) File {
	f, err := fs.Internal.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC)
	if err != nil {
		panic("Failed to open file")
	}
	file := File{internal: f, isVirt: false}
	return file
}

func (fs *Filesystem) Mkdir(path string) error {
	err := fs.Internal.Mkdir(path, 0777)
	return err
}

func (fs *Filesystem) Remove(path string) error {
	return fs.Internal.Remove(path)
}

func (fs *Filesystem) Rename(oldPath string, newPath string) error {
	return fs.Internal.Rename(oldPath, newPath)
}

func (fs *Filesystem) Size() (n int, err error) {
	n, err = fs.Internal.Size()
	return n, err
}

type File struct {
	internal tinyfs.File
	isVirt   bool
}

func (f *File) Read(buf []byte) (n int, err error) {
	i, err := f.internal.Read(buf)
	return i, err
}

func (f *File) Close() error {
	err := f.internal.Close()
	if err != nil {
		print(err.Error())

	}
	return err
}

func (f *File) Write(buf []byte) (n int, err error) {
	n, err = f.internal.Write(buf)
	return n, err
}
