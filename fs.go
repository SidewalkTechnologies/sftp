package sftp

import (
	"io/fs"
	"os"
)

type VirtualFileHandle interface {
	Close() error
	Stat() (fs.FileInfo, error)
	Name() string
	Readdir(length int) ([]fs.FileInfo, error)
	Truncate(toSize int64) error
	Chmod(mode os.FileMode) error
	Chown(uid, gid int) error
	ReadAt(buffer []byte, l int64) (int, error)
	WriteAt(data []byte, l int64) (int, error)
}

type FS interface {
	Stat(path string) (os.FileInfo, error)
	Lstat(path string) (os.FileInfo, error)
	Mkdir(path string, mask fs.FileMode) error
	Remove(path string) error
	Rename(path, to string) error
	Symlink(path, to string) error
	Readlink(path string) (string, error)
	OpenFile(path string, flags int, mode os.FileMode) (VirtualFileHandle, error)
	Write(path string, data []byte) (int, error)
	WriteAt(path string, data []byte, offset int64) (int, error)
}
