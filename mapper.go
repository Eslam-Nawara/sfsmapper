package sfsmapper

import (
	"fmt"
	"reflect"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"github.com/Eslam-Nawara/sfsmapper/internal/fsnode"
	"github.com/fatih/structs"
)

const errNotPermitted = "Operation not permitted"

// File system Struct
type FS struct {
	dataMp map[string]any
	Struct any
}

// Mounts a fuse connection to a mounting point and starts a server to serve the connection requests
func Mount(data any, mountPoint string) error {
	con, err := fuse.Mount(mountPoint)
	if err != nil {
		return err
	}

	defer con.Close()

	err = fs.Serve(con, NewFS(data))
	if err != nil {
		return err
	}
	return nil
}

func Umount(mountPoint string) error {
	err := fuse.Unmount(mountPoint)
	if err != nil {
		return err
	}
	return nil
}

// Creates a new file system initiated with the data argument
func NewFS(data any) *FS {
	return &FS{
		Struct: data,
		dataMp: structs.Map(data),
	}
}

// Initialize the root directory
func (fs *FS) Root() (fs.Node, error) {
	dir := fsnode.NewDir()
	dir.Entries = createEntries(fs.dataMp, []string{}, fs.Struct)
	return dir, nil
}

// Creates a map of entries that a directory contains
func createEntries(structMap any, path []string, Struct any) map[string]any {
	entries := map[string]any{}
	for key, val := range structMap.(map[string]any) {
		if reflect.TypeOf(val).Kind() == reflect.Map {
			path = append(path, key)
			dir := fsnode.NewDir()
			dir.Entries = createEntries(val, path, Struct)
			entries[key] = dir
		} else {
			entries[key] = fsnode.NewFile(len([]byte(fmt.Sprintln(reflect.ValueOf(val)))), path, Struct, key)
		}
	}
	return entries
}
