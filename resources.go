// Code generated by go-bindata.
// sources:
// transport.tmpl
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _transportTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x54\xcb\x6e\xdb\x3a\x10\x5d\x8b\x5f\x31\x57\x8b\x0b\xb2\x88\xa9\x7d\x8a\x2c\xd2\x5a\x2d\x0c\x04\x76\x91\x18\xe8\xa6\x8b\xa8\x32\xad\xa8\x8e\x48\x95\xa2\xfc\x80\xa1\x7f\xef\xf0\xe5\x38\x81\x9c\x08\x30\xec\x19\xcd\xe3\x9c\x39\x33\xce\x32\xf8\x9e\xcf\xf3\xfb\xdb\x65\x3e\x85\x6f\xb3\xbb\x1c\x26\x13\x98\x2e\x60\xbe\x58\x42\x3e\x9d\x2d\xe1\xc7\x5d\x7e\xfb\x90\x93\xb6\x28\x37\x45\x25\xa0\x29\x6a\x49\x48\xdd\xb4\x4a\x1b\xa0\x24\x49\x85\x2c\xd5\xaa\x96\x55\xf6\xa7\x53\x32\x45\x87\x14\x26\x7b\x32\xa6\x4d\x09\x1a\x55\x6d\x9e\xfa\xdf\xbc\x54\x4d\x56\xa9\xc9\xa6\x36\x99\xfd\x08\xb9\x6a\x55\x2d\x8d\x0d\xaf\xd4\x73\x21\x2b\xae\x74\x95\xed\x33\x9b\x5b\x2a\x69\xc4\x1e\xdf\x31\x42\x10\xdd\xbd\xf8\xdb\x8b\xce\x10\x73\x68\x05\x1c\x8f\xc0\xf3\x90\x3c\x2f\x1a\x01\xc3\x10\xde\x43\x67\x74\x5f\x1a\x38\x12\x1b\x13\x9c\x33\xd9\xf6\x06\x63\x5c\x9e\x33\x96\xb6\x0a\x3a\x1e\x2d\xda\xeb\x74\x24\x36\x7d\x24\x43\x68\xdc\xb5\x4a\x76\xe2\x9d\xce\x3e\xe0\xa5\x35\x80\x2f\xe8\xfd\x8b\xde\x9c\xb5\xf7\xd6\x68\xff\x37\xe1\x88\x20\xc9\xb5\x06\x7c\xb0\x32\x8e\x16\xce\x9e\x90\x28\xb4\xbe\x52\x4d\x6d\x44\xd3\x9a\xc3\x09\x72\x04\x48\x3b\x46\xd6\xbd\x2c\x51\xad\x8d\x18\x41\xfe\x12\xb7\x2d\xc7\x98\x3d\x08\xbd\xad\x4b\xc1\x20\x0a\x75\x0a\x40\x92\x89\x16\xa6\xd7\x12\x6c\x03\x5a\x9a\x3d\x04\xc1\xf8\x57\xff\x7d\x05\x3a\x48\x82\xf1\x42\xaf\x8b\x52\x1c\x07\x06\xf4\xcc\xba\x02\xc4\xaf\x34\x73\x23\xc3\x68\xb8\xbe\x89\x49\x9c\x5e\xd6\x98\x61\xb4\x72\x53\x72\x05\x6c\x16\x12\xe0\x23\x09\x14\xab\xf1\xb7\xe2\x0e\x03\x23\x49\x52\xaf\x5d\xee\x7f\x37\x20\xeb\x67\x07\xc0\x42\x70\x8c\xde\xd1\xf8\xf8\x6b\x54\x2a\x87\x83\xe7\x96\x0c\x65\x68\x61\x49\x6c\x31\x7c\x50\x14\x3e\xac\x9a\xa6\xb1\xd8\x10\xa4\x9d\x0a\x3c\x33\x71\xba\x06\xa7\xee\xca\xf9\x2e\xcf\x8b\x6a\xf8\x64\x2f\x31\x4e\xe1\xa2\x08\xdb\x42\x9f\x44\xbb\x5c\x8e\xc4\xd9\xe1\xdc\xed\x16\xf2\xb9\xd8\x79\x58\x9a\x6a\xfe\x45\xad\x0e\x8c\x7b\x9b\xfe\x1f\xaa\xb1\xcf\xaf\x87\x9d\xc4\xe5\x41\xd3\xf5\xb7\x04\xa3\x2f\xe4\x78\xe2\x71\xa1\x03\xeb\x70\x8a\x8e\xb6\xfb\xc7\x11\xd1\x47\x77\x10\x38\x7a\xfb\xa7\xc6\xab\xd0\x76\x09\xc3\x75\xbe\xda\x42\x47\xf9\x6c\x89\x23\x0f\xdf\x48\xd3\x1d\xe3\xfe\x27\x8d\xf9\xcc\x42\xf9\x17\x00\x00\xff\xff\x62\xd7\x02\xdc\x26\x05\x00\x00")

func transportTmplBytes() ([]byte, error) {
	return bindataRead(
		_transportTmpl,
		"transport.tmpl",
	)
}

func transportTmpl() (*asset, error) {
	bytes, err := transportTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "transport.tmpl", size: 1318, mode: os.FileMode(420), modTime: time.Unix(1460004350, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"transport.tmpl": transportTmpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"transport.tmpl": &bintree{transportTmpl, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

