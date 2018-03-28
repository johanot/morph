// Code generated by go-bindata.
// sources:
// data/eval-machines.nix
// data/options.nix
// DO NOT EDIT!

package assets

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

var _dataEvalMachinesNix = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x56\x4d\x6f\xdb\x38\x10\xbd\xeb\x57\x3c\xb8\x01\x6c\x21\x8d\xd2\xdd\xdb\xda\xf0\xa1\x5b\xec\x17\xd0\x8f\xa0\x2d\x7a\x09\x82\x05\x2d\x8d\x6c\x22\x14\xa9\x92\x94\x13\xaf\xe1\xff\xbe\x18\x52\x12\xe5\x34\x87\xc6\x87\x44\x1a\x91\xc3\x99\x37\x8f\x6f\xe6\x15\xde\x99\xa6\x55\xe4\x49\x1d\xe0\xbc\x95\x6d\x4b\x15\x2a\xf3\xa0\xb1\x27\xeb\xa4\xd1\x30\x35\xb4\x7c\x34\xad\x9b\x83\xf6\x42\x75\xc2\x1b\x9b\x1d\xa1\xc9\x3f\x18\x7b\xff\xc7\x63\x6b\x71\x5a\x66\x99\x22\x9f\x61\xb0\x62\x0d\xd9\xb4\xc6\xfa\xe9\xb2\x55\x06\xb4\xf7\x5b\x07\x00\xeb\xe1\x43\x31\xfc\xe7\x2f\xbc\x42\xc9\x0d\xe2\x0a\xb6\x14\x4a\x6e\x56\x99\xd4\x19\xf0\x20\xfd\x0e\xc3\xaa\xf0\x12\xbe\x65\x96\x4a\x1c\x33\x20\xa6\xd2\x79\x82\xdf\x11\x2a\xaa\xa5\x96\x5e\x1a\xed\x38\x03\x36\x35\xa2\xdc\x49\x4d\xae\xe0\x38\x4d\x45\x0e\xeb\x8c\x4f\x52\xd2\xf9\xaf\xe6\xad\xf7\xd6\x61\xd1\x88\x96\xff\x84\x95\x1f\x45\x43\xcb\xb0\x04\x88\xe9\xc5\xdf\x2b\xfc\x45\x3e\xb8\x2c\x8d\xae\xe5\xb6\xb3\xc2\xf7\x48\xf9\x9d\x74\xc3\x41\xa8\xad\x69\x40\xa2\xdc\x0d\xc9\x4e\x3c\xd0\x63\x6b\xc9\x31\xc0\xaf\x21\xbc\x0f\x3b\xb6\xf8\xb7\x96\x8a\xf8\xdd\xca\x4d\xe7\xc9\xc1\x99\x70\xce\x47\xf9\xf8\xe9\x0b\x1a\x53\x75\x8a\x26\x4e\xdc\xc1\x79\x6a\x50\x0a\x8d\xad\xdc\x13\x1c\x69\x27\x37\x8a\x40\xd6\x1a\x8b\x86\x9c\x13\xdb\x98\x70\xfc\x45\x0f\x0e\x6b\xdc\xe2\xd8\x97\x28\xbe\x0d\x65\xb8\x38\x4e\x92\x3f\xe1\x6e\x85\x13\xaf\xd4\x3b\xb2\xd2\x63\xd1\x2f\xcb\x63\xa4\xfc\xf1\x6e\xd5\x7b\x0f\x45\xe2\xdf\x11\x5a\x34\x84\x35\x26\x9e\x56\x63\x08\x4c\x21\x4a\xfc\x98\x5d\x1c\xbd\xf9\xe2\x2d\x67\x1f\x0a\xde\x0a\xbf\x3b\x5d\x33\xe3\xdc\xb5\x92\x9b\x6b\xe6\xdc\x55\xc4\xb9\xd0\xf2\x71\x16\x6a\xfd\x43\x3e\x13\x5b\xb2\x5e\x5e\x9e\x99\x39\xe5\x7b\x3a\x60\x8d\x59\x45\xad\x32\x87\x2b\xe7\xbb\xba\x9e\xad\xce\x56\x85\x54\x26\xc0\x14\xd7\xa6\x0d\x34\xe2\xd3\x53\xb6\xe9\xf7\x0a\x37\xd6\xec\x65\x45\x10\xcc\x3a\xd1\x29\x8f\x9d\x71\x3e\x80\x20\x74\x85\x78\x58\x43\xda\xc3\x0b\xbb\x25\x0f\xfa\xde\x09\xf5\x8c\x23\x1f\xab\x3d\xd6\x3f\x02\x79\x4e\x5f\x48\x1d\x5f\x4d\x45\xaa\xf8\xc1\x49\x5f\x20\xa9\xb7\x05\x07\xf1\xb1\xaf\xc4\xfd\xa7\x3d\x59\xcb\x41\xfe\xf6\xe6\xcd\xf3\x85\x19\x7e\x29\xdc\x22\x86\xfb\xb7\x71\xfe\x45\x3e\x4e\x67\xef\x67\x90\xd1\xa3\xb7\xe2\xad\xdd\x32\xb6\x89\x56\xf1\x36\xae\x9e\xe3\x0d\x4e\x69\xfb\xf8\x18\x0f\xc8\xb1\x60\xa4\x78\x99\xc3\xc2\x52\x63\xf6\x14\x6f\xf1\xa0\x42\xb7\x98\xf5\x8f\x33\xae\x79\xa8\x8d\x9b\x61\x66\xc9\x99\xce\x96\x14\x9f\xbf\x77\xd2\xd2\x0c\xb3\x40\xe9\x19\xee\xf2\x3c\x5f\x65\x59\x36\x45\xe2\x1f\x5d\x9b\x0f\x81\x55\x1c\x77\x76\x0e\xd3\x68\xc2\x90\x40\xfa\x36\x85\x47\xd8\x6d\xc7\x36\xce\x5d\xd8\xa8\x63\x81\x6d\x3d\x0a\x5d\x27\xab\x68\x0b\x89\x9e\x56\x59\x50\xb5\x9b\x9d\x70\x84\x5f\x96\x83\xfa\x12\x8c\x56\x87\x5e\xe4\xc6\x18\x92\x66\x30\x27\xa4\xae\xcd\xa0\x6e\xa3\x70\xf5\x50\xcc\x93\xf8\x0e\x11\x04\xfc\x83\x9d\x1f\xa2\x55\x6a\x44\x65\x0d\x6f\x83\x76\x8e\x37\xad\x56\xb2\x45\x23\xda\x1e\xf0\x50\xc0\x85\x5e\x62\x3f\x5f\xf2\x91\xd8\x63\x0d\x57\xda\x6e\xf3\x29\xdc\x9e\x6f\xe1\xd6\xef\xe7\xab\x24\x13\x98\xea\xca\xbe\xe8\xaf\x78\x4a\x29\xc7\x84\x7e\x8e\x4a\x4b\xde\x4d\xa9\xd4\x63\xad\xcf\x6c\x2c\x1b\x9f\x49\x11\x43\xb6\xc6\xe8\x35\x0a\x65\x71\xf6\xd9\xd8\x81\x34\x5f\xba\xba\x96\x8f\xcf\xaf\xfe\x16\xbb\xe0\x4f\x2c\xc9\x27\x44\xed\x9f\xf2\xd5\x19\x7a\xef\x65\xb8\x48\xb1\xc7\xdc\xd3\x61\x89\x2d\x79\x06\x30\xc8\xd2\x00\xf1\x19\xad\x47\x63\xef\x3d\xf5\xd7\xa1\x9a\x43\x03\xe5\x7c\x8e\xcf\x10\xe7\xd7\x25\x36\x9d\x54\x15\xca\xbe\xd1\x8f\x42\x72\xd6\xbc\x02\x6d\x52\x95\x7b\x11\x77\xdc\xda\x7b\x16\x25\x96\xd4\x52\x79\xb2\x7d\xc3\xe4\x9a\x2f\x41\x8a\x1a\xe8\xb8\x27\xef\x69\x34\x94\xda\x76\xfa\x9d\x69\x1a\x96\xc2\x59\x63\x6c\xbb\x9b\x8d\x8d\xa2\xb5\x54\x93\x7d\x6f\x4a\xa1\x7e\x0f\x41\xae\xe1\x6d\xc7\xb7\xbe\x5f\x32\x9f\xa7\xde\x75\x5f\x49\x8b\xab\x16\x17\xa6\x4b\xbd\x78\xd2\x3e\x02\x6c\x81\x68\xb1\x8f\x9f\x05\x38\x71\x04\x28\x8d\x2b\x87\x8b\xe3\xd3\x72\x06\xa0\x0a\x6f\x5a\x45\x7b\x52\xa7\x70\xd2\xf5\xc5\x51\x9f\x5e\xbc\xb7\xa8\xec\xfe\x86\x3b\x59\xf2\xc1\xa6\xd1\xcf\x7c\xde\xa3\x34\xcf\xf3\x94\x6b\x5f\xb8\x3f\x3b\x5d\x86\x81\x42\x13\x55\x54\x71\x5b\x28\x85\x2a\x3b\x25\xfa\xe1\x26\xce\x63\x49\x50\x0a\x7c\xe5\xb9\xc3\xed\x4c\xa7\x2a\x04\x36\xd0\x9e\x34\x1e\x76\xa4\xd3\xaa\xe0\x5b\x58\x82\x36\x7c\xa1\x3c\x0e\xe4\x5f\xf3\x84\xf1\x40\xa8\xc5\x3d\x4d\x04\xca\x1b\x6c\x08\x82\x47\x09\x6f\x92\xf0\xf0\xd9\xbd\x5e\x4e\x9a\x94\xa9\x83\xe7\x49\x60\x03\x25\xd3\x90\xe3\x8a\x20\xa8\x5d\xd3\x1c\x7a\xf5\xaf\x23\x2f\xbd\xd4\x3c\xe4\x3d\x9d\xc0\xc4\x92\xa7\xbb\x82\xf9\x14\x2a\x7a\x23\xa4\x85\x40\x2d\x94\xa3\x1c\x8b\x71\xe7\xa4\x01\x8c\xb6\xba\x87\x2f\x1c\x54\x47\x35\x07\x5f\x35\x9e\xa2\x5a\xc7\xa3\x28\x07\xa0\xe2\x11\x9d\x96\xdf\x3b\xc2\x82\x9f\x6b\x25\xbc\x27\x1d\x63\xd8\x92\xff\x1c\x73\x75\x50\xc9\xc9\x68\xe3\x14\x9e\x2a\x6c\x48\x3f\x8c\xc4\xe3\x80\x53\x0f\xb2\x60\xc7\x8d\x23\x0b\x64\x9d\x30\x90\x2e\x95\x3d\xb9\xf1\x3b\xd2\x13\xf2\x2d\x16\x93\x6f\x8b\x04\x67\xb2\xe6\x79\x31\x54\xc8\x58\xdc\xde\xe5\xe3\x6e\x52\x8e\xa6\xae\xd2\x9e\x27\x3b\x06\xfd\xef\x17\xdf\xa2\xc6\x1d\x2e\x2f\xf1\x14\x94\x21\xa1\x80\x0c\xf7\xce\xaf\x66\xac\xed\xcf\x03\x33\x9e\xf3\x42\x30\x06\x9a\x1c\x79\x84\x14\xa7\x19\xd6\x69\xfb\x28\x09\x35\x4f\xa8\x2f\xe2\xcb\x14\xc9\x2c\x01\x87\xdb\xbb\x73\x16\x0d\x99\xba\xc8\xa2\xff\xe4\x20\x37\x4f\x39\x55\x1a\x5d\x0a\xff\x41\xb4\x53\x88\x16\x4f\xd8\x58\xbb\x40\xd3\x53\xf6\x7f\x00\x00\x00\xff\xff\x6e\x06\x9f\xce\x90\x0d\x00\x00")

func dataEvalMachinesNixBytes() ([]byte, error) {
	return bindataRead(
		_dataEvalMachinesNix,
		"data/eval-machines.nix",
	)
}

func dataEvalMachinesNix() (*asset, error) {
	bytes, err := dataEvalMachinesNixBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/eval-machines.nix", size: 3472, mode: os.FileMode(420), modTime: time.Unix(1522255979, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataOptionsNix = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x91\xcd\x8e\x9b\x40\x10\x84\xef\xf3\x14\x25\x5f\xbc\x2b\x59\x3c\x40\x90\x0f\xb9\xe5\x10\x65\xa5\x28\x2f\x30\x40\x01\x23\x03\x3d\x9a\x6e\x62\x5b\x88\x77\x8f\xf8\x89\x93\x5d\xf9\x10\x29\x17\x5a\x14\x55\x5f\x77\xd3\x13\x4a\x19\xea\xd0\x9c\xd0\x85\xe2\x84\x78\x69\xf4\x84\x2c\xcb\x30\x7f\x72\xee\x1a\xac\x5d\xf4\xdc\xb9\x8e\xe6\xdc\x85\xf7\xb7\x68\x41\x06\xfd\x71\x8f\xc4\x19\x76\x8f\xd4\x4c\xc7\xa2\x97\x6a\xec\x88\x97\x69\xcf\x62\x72\x80\x6c\x5e\x9c\xd7\x37\xa0\xa2\x5a\x18\xfc\x22\xe2\x8c\xfe\xb2\xb1\xf6\x8f\x58\x61\x7f\x98\x96\xf2\x5d\xaf\xa8\x65\x0a\x71\x8f\x1d\xbe\xb3\x17\x23\xa2\xb7\xf6\xb0\x59\xe6\xdc\xad\x55\x65\x4c\x25\xff\x03\xfd\x55\x4a\xdf\x7d\x20\xaf\x8f\xf9\x35\x77\x2e\x0c\xce\xfd\xb5\x56\x56\x31\x76\x72\xef\x39\xd8\x63\x43\xf3\xa9\xa1\x7d\x11\xb5\x7f\x9f\x62\xde\x8a\xb2\x4c\x34\x7d\x96\xab\x58\xfb\xb1\x5b\xbb\xcc\xbf\x27\xe7\xcd\xf7\xb1\x5b\x70\x13\xa2\x57\xbd\x4a\xaa\x32\xe3\x6d\x71\x1d\x6a\x91\xc2\xa7\x43\x8e\x87\xfd\x5d\x6b\x6f\x96\xf4\xad\xc6\xfb\x73\x3e\xff\x27\xc7\xe3\x2e\x03\x9f\x97\x18\x0d\xd7\x96\x89\xa0\x2f\x5b\x2c\xa4\x50\x8c\xc6\x3d\x55\x50\xe1\x17\x2e\x4c\x50\x10\xa5\xc4\xc0\x0a\x3f\x83\x87\x6a\xfb\x20\x85\x41\x8d\xbe\x82\xd4\xb0\x36\xc9\xd8\xb4\xb0\x96\xf8\x16\x6e\x28\x3b\xd1\x31\x11\x2f\x17\x32\x86\xa1\x41\x30\xc8\x68\x9b\x75\xb3\xa8\x49\x62\xf6\xba\xd3\x8e\xc7\x0f\xa7\x72\xbf\x02\x00\x00\xff\xff\x5c\xf3\xee\x8e\xd3\x02\x00\x00")

func dataOptionsNixBytes() ([]byte, error) {
	return bindataRead(
		_dataOptionsNix,
		"data/options.nix",
	)
}

func dataOptionsNix() (*asset, error) {
	bytes, err := dataOptionsNixBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/options.nix", size: 723, mode: os.FileMode(420), modTime: time.Unix(1521909568, 0)}
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
	"data/eval-machines.nix": dataEvalMachinesNix,
	"data/options.nix": dataOptionsNix,
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
	"data": &bintree{nil, map[string]*bintree{
		"eval-machines.nix": &bintree{dataEvalMachinesNix, map[string]*bintree{}},
		"options.nix": &bintree{dataOptionsNix, map[string]*bintree{}},
	}},
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

