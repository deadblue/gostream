package multipart

// FileInfo holds file name and file size.
// Caller can use os.FileInfo as FileInfo, or create one by SimpleFileInfo().
type FileInfo interface {
	Name() string
	Size() int64
}

type implFileInfo struct {
	name string
	size int64
}

func (i *implFileInfo) Name() string {
	return i.name
}

func (i *implFileInfo) Size() int64 {
	return i.size
}

// Create a FileInfo by name and size.
func SimpleFileInfo(name string, size int64) FileInfo {
	return &implFileInfo{name, size}
}
