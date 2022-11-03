package internal

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"bazil.org/fuse"
	"github.com/fatih/structs"
)

type File struct {
	Type       fuse.DirentType
	Path       []string
	Struct     any
	Name       string
	Attributes fuse.Attr
}

// Create new empty file
func NewFile(contentLength int, path []string, str any, name string) *File {
	return &File{
		Type:   fuse.DT_File,
		Path:   path,
		Struct: str,
		Name:   name,
		Attributes: fuse.Attr{
			Inode: 0,
			Atime: time.Now(),
			Mtime: time.Now(),
			Ctime: time.Now(),
			Size:  uint64(contentLength),
			Mode:  0o444,
		},
	}
}

// Provides the core information for a file
func (file *File) Attr(ctx context.Context, a *fuse.Attr) error {
	file.updateFileContent()
	*a = file.Attributes
	return nil
}

// Returns the content of a file
func (file *File) ReadAll(ctx context.Context) ([]byte, error) {
	return append(file.updateFileContent(), []byte("\n")...), nil
}

func (file *File) GetDirentType() fuse.DirentType {
	return file.Type
}

// Read the file content
func (file *File) updateFileContent() []byte {
	structMap := structs.Map(file.Struct)

	for _, v := range file.Path {
		structMap = structMap[v].(map[string]any)
	}
	content := []byte(fmt.Sprintln(reflect.ValueOf(structMap[file.Name])))
	file.Attributes.Size = uint64(len(content))
	return content
}
