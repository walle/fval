package fval_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/walle/fval"
)

func TestFileExists(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "go-arg-test")
	if err != nil {
		t.Errorf("TempDir error")
		return
	}
	defer os.RemoveAll(tmpDir)

	path := filepath.Join(tmpDir, "test")
	os.Create(path)

	if !fval.FileExists(path) {
		t.Errorf("Did not indicate that existing file %s exists", path)
	}

	os.Remove(path)

	if fval.FileExists(path) {
		t.Errorf("Did indicate that non existing file %s exists", path)
	}

	os.Mkdir(path, 0766)

	if fval.FileExists(path) {
		t.Errorf("Did indicate that directory %s is a file", path)
	}
}

func TestDirExists(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "go-arg-test")
	if err != nil {
		t.Errorf("TempDir error")
		return
	}
	defer os.RemoveAll(tmpDir)

	path := filepath.Join(tmpDir, "test")
	os.Mkdir(path, 0766)

	if !fval.DirExists(path) {
		t.Errorf("Did not indicate that existing directory %s exists", path)
	}

	os.Remove(path)

	if fval.DirExists(path) {
		t.Errorf("Did indicate that non existing directory %s exists", path)
	}

	os.Create(path)

	if fval.DirExists(path) {
		t.Errorf("Did indicate that file %s is a directory", path)
	}
}

func TestFileOrDirExists(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "go-arg-test")
	if err != nil {
		t.Errorf("TempDir error")
		return
	}
	defer os.RemoveAll(tmpDir)

	fpath := filepath.Join(tmpDir, "test.txt")
	dpath := filepath.Join(tmpDir, "test")
	os.Create(fpath)
	os.Mkdir(dpath, 0766)

	if !fval.Exists(fpath) {
		t.Errorf("Did not indicate that existing file %s exists", fpath)
	}
	if !fval.Exists(dpath) {
		t.Errorf("Did not indicate that existing directory %s exists", dpath)
	}

	os.Remove(fpath)
	os.Remove(dpath)

	if fval.FileExists(fpath) {
		t.Errorf("Did indicate that non existing file %s exists", fpath)
	}
	if fval.FileExists(dpath) {
		t.Errorf("Did indicate that non existing directory %s exists", dpath)
	}
}

func TestDirExistsOrCreate(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "go-arg-test")
	if err != nil {
		t.Errorf("TempDir error")
		return
	}
	defer os.RemoveAll(tmpDir)

	path := filepath.Join(tmpDir, "test1", "test2", "test3")

	ok, err := fval.DirExistsOrCreate(path, 0766)
	if err != nil {
		t.Errorf("Error creating dir: %s", err)
	}
	if !ok {
		t.Errorf("Directory was not created")
	}
	ok, err = fval.DirExistsOrCreate(path, 0766)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if !ok {
		t.Errorf("Directory did not exist")
	}

	os.Remove(path)

	if fval.DirExists(path) {
		t.Errorf("Directory %s should not exist", path)
	}

	ok, err = fval.DirExistsOrCreate(path, 0766)
	if err != nil {
		t.Errorf("Error creating dir: %s", err)
	}
	if !ok {
		t.Errorf("Directory was not created")
	}
	if !fval.DirExists(path) {
		t.Errorf("Directory %s was not created", path)
	}
}

func TestDirExistsOrCreateError(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "go-arg-test")
	if err != nil {
		t.Errorf("TempDir error")
		return
	}
	defer os.RemoveAll(tmpDir)

	path := filepath.Join(tmpDir, "test")
	os.Mkdir(path, 0400)
	path = filepath.Join(path, "test2")
	ok, err := fval.DirExistsOrCreate(path, 0766)

	if ok {
		t.Errorf("Directory should not have been created")
	}
	if err == nil {
		t.Errorf("Error should not have been nil")
	}
}
