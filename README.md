
# Struct to file system mapper 
Struct to Filesystem mapper that uses (FUSE) Filesystem in USErspace to map an in-memory data structure to an accessible file system.

## How to use 
**Get the package**
```sh 
 go get github.com/Eslam-Nawara/sfsmapper
```
**Package API:**
|Function | Description |
| :--- | :--- |
| `Mount(data any, mountPoint string) error` | Mounts user defined struct to the given mount point |
| `Umount(mountPoint string) error` | Unmounts the filesystem mounted in the given mount point |

## Example
```go
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Eslam-Nawara/sfsmapper"
)

type SubStruct struct {
	SomeValue      int
	SomeOtherValue string
}

type Fuse struct {
	Text   string
	Intger int
	Sub    SubStruct
}

func main() {
	var err error
	if len(os.Args) != 2 {
		fmt.Println("Mounting point not specified")
		return
	}
	mountPoint := os.Args[1]
	data := &Fuse{
		Text:   "This is a text content",
		Intger: 2222,
		Sub: SubStruct{
			SomeValue:      20,
			SomeOtherValue: "some data",
		},
	}

	err = os.MkdirAll(mountPoint, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		time.Sleep(2 * time.Second)
		data.Text = "This is a changed text content"
	}()

	err = sfsmapper.Mount(data, mountPoint)
	if err != nil {
		fmt.Println(err)
		return
	}
	sfsmapper.Umount(mountPoint)
}
```
