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
