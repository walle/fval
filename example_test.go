package fval_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/walle/fval"
)

func Example() {
	tmpDir, err := ioutil.TempDir("", "go-arg-test")
	if err != nil {
		fmt.Println("TempDir error")
		return
	}
	defer os.RemoveAll(tmpDir)

	os.Args = make([]string, 3)
	os.Args[0] = "example"
	os.Args[1] = filepath.Join(tmpDir, "input")
	os.Args[2] = filepath.Join(tmpDir, "output")

	fmt.Println(fval.FileExists(os.Args[1]))
	os.Create(os.Args[1])
	fmt.Println(fval.FileExists(os.Args[1]))

	exists, err := fval.DirExistsOrCreate(os.Args[2], 0766)
	fmt.Println(exists)
	fmt.Println(err)

	fmt.Println(fval.Exists(os.Args[1]))
	fmt.Println(fval.Exists(os.Args[2]))

	// Output:
	// false
	// true
	// true
	// <nil>
	// true
	// true
}
