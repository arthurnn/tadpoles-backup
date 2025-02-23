// Code generated for package bindata by go-bindata DO NOT EDIT. (@generated)
// sources:
// utils/dist/reset-tadpoles-password.html
package bindata

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _utilsDistResetTadpolesPasswordHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x57\x6f\x6f\xe3\x38\xce\x7f\x3f\x9f\x42\xcd\x40\x0f\x52\x60\xe2\xb8\x72\xd3\x3e\x93\x49\x0b\x34\xc9\x64\x6f\x81\xc3\xdd\x62\x67\x81\xc3\xbd\x2a\x14\x8b\xb1\xd5\xca\x92\x57\xa2\x93\x06\x87\xfb\xee\x07\xca\xce\xdf\x36\xb7\xb7\x46\x9b\xd8\x34\x45\xf2\x47\xfe\x44\x31\x93\xab\xf9\xdf\x67\xbf\xfd\xf3\x97\xef\xac\xc4\xca\x3c\x7e\x9a\xd0\x17\x33\xd2\x16\x0f\x3d\xb0\x3d\x12\x80\x54\x8f\x9f\x18\x63\x6c\x82\x1a\x0d\x3c\x7a\x08\x80\x03\x94\xaa\x76\x06\xc2\xa0\x96\x21\x6c\x9c\x57\x93\x61\xfb\xba\x55\x0d\xb8\xdd\xdd\xd3\x45\x56\xbf\xb0\xa5\x53\xdb\x2f\x4c\xe9\x35\xfb\xd7\xfe\x0d\x5d\x95\xf4\x85\xb6\x63\x96\x7e\x3b\x11\xd7\x52\x29\x6d\x8b\x77\xf2\x95\xb3\x38\x66\xd6\xf9\x4a\x1a\x76\x73\x57\xbf\x0d\xc5\x6d\xfd\xc6\xfe\x02\x66\x0d\xa8\x73\xc9\xfe\x06\x0d\x7c\x39\x3c\x7f\x61\x41\xda\x30\x08\xe0\xf5\xea\xd4\x52\xee\x8c\xf3\x63\xf6\x39\xcb\xb2\xc3\x8b\x7f\x7f\xda\xdf\x7e\xae\xa4\xb6\xef\x82\x7d\x1b\x6c\xb4\xc2\x72\xcc\xee\xb2\xb4\x7e\xfb\xf6\x21\x94\x0c\x2a\x26\x1b\x74\x1f\x9a\x4d\x96\xce\xbd\x56\xd2\xbf\x1a\xc0\x33\xeb\x4a\x87\xda\xc8\xed\x98\x69\x6b\xb4\x85\xc1\xd2\xb8\xfc\xf5\x42\x5a\x92\x11\x54\xec\x06\xaa\x8f\x31\xad\xe2\x75\xfa\x6e\x29\xf3\xd7\xc2\xbb\xc6\xaa\x31\xfb\x7c\x73\xfb\xff\x77\x70\x7f\xa6\xe0\xbc\x02\x3f\xf0\x52\xe9\x26\x8c\xd9\xed\x39\x3c\x84\x37\x1c\x28\xc8\x9d\x97\xa8\x9d\xa5\x2a\x58\xf8\x10\xa3\x3c\x03\xb6\x0b\xeb\xdc\xeb\x71\x5a\x94\xae\x2e\xac\xfa\xfa\xf5\xeb\xf1\x12\xfa\x9c\x0c\x3b\x8e\x4d\x86\x2d\x45\x27\x44\xaf\xc7\x4f\x13\xe2\x97\x56\x0f\x3d\xaa\x5d\xaf\xa3\x63\x79\x73\x99\xb6\xe5\x4d\xa7\xb4\xf4\x8f\x6d\x30\x13\x67\x0e\xd4\x9d\x18\xfd\x78\x12\xd3\xa4\x7e\x9c\x04\xf4\xce\x16\x8f\xbf\x82\x54\x0c\x4b\x60\xab\xc6\x18\xa6\x6d\x40\xdf\xe4\x94\x18\x66\x74\x40\xb6\xd2\x3e\xe0\x15\x05\x1a\xb5\x27\xc3\xfa\xc8\xec\xf0\xd8\xee\x47\x4e\x4e\x9e\xe9\x9a\x7b\x59\x30\x2c\x75\x60\xcb\x06\xd1\x59\x86\x8e\x6d\x5d\xe3\xd9\x8e\x4e\x81\x2d\xa5\x27\x69\x90\x6b\x60\x1a\x99\x0c\x4c\xb2\x23\xb2\x8d\xff\xd0\xc7\x44\xb2\xdc\xc8\x10\x1e\x7a\x47\xcb\x7a\xac\xf4\xb0\x7a\xe8\xbd\xc8\xb5\x0c\xb9\xd7\x35\x8e\xfb\xab\xc6\x46\xa4\xfd\x6b\x7e\x3f\xe5\x42\x34\x01\xb8\x48\x03\x7a\x9d\x23\x17\x82\x67\xd3\x9d\x06\x17\x69\x2e\x8d\x21\xee\x45\xe5\x23\x79\xe8\x03\x49\xd6\xd2\x73\x91\x22\xcf\xe6\x5c\xdc\x92\x28\x9b\x7a\xc0\xc6\x93\x0a\x26\x06\x6c\x81\x25\xcf\x66\x3c\x9b\xa7\x3c\x5b\xf4\x73\x67\x83\x33\x90\x80\xf7\xce\xf7\xff\xef\x73\xf6\xf5\xdb\xcc\x35\x46\x71\x91\x5a\x87\x5c\xa4\x2b\x6d\xe9\x01\x0c\x54\x60\x49\xb0\xd1\x58\x92\x3b\x30\x90\xa3\x23\x67\x5c\x88\xb8\x90\x8b\x29\x70\x31\xed\xee\x45\x12\x6f\xae\xb9\x98\xa5\xd7\x3c\x7b\xea\x23\x1f\x4d\x53\x3e\x9a\x27\xb9\xd1\x31\x7c\x31\xbb\xb9\xe6\xf7\xf3\xab\x13\xf8\xbb\x80\xb4\x5d\xb9\x3e\x17\xe2\x07\x4a\x8f\xda\x16\x5c\xa4\x17\x58\x47\xc1\xc4\x44\x72\x21\xc8\x28\x17\x62\xb3\xd9\x24\x3b\xbd\x24\x77\x15\x17\xe2\x8a\x32\x92\xcd\x37\xda\x2a\xb7\x49\x8c\xcb\xe3\xa6\x4b\x4a\x17\xd0\xca\x0a\xb8\xb8\x8b\x7f\x82\x8b\x45\xe9\x2a\x78\x76\xfe\x79\xe3\xfc\xeb\xe5\x95\xb5\xc4\xf2\xb0\xb2\x2f\x0d\x78\xa4\x80\xff\xa1\x8d\x89\xc1\x2a\xed\x81\xea\x97\xa2\xe3\x22\x2d\x11\xeb\xc0\xb3\x27\x2e\x16\x5c\x2c\x4e\x83\x3b\xf5\x38\x9a\x59\xfa\x7f\x5a\x21\xf8\x8d\x8c\xf8\x62\xc6\xb8\x48\x8f\x78\xc4\x45\x2a\x0b\xa9\x6d\xd2\xa1\x3e\x0f\xcf\x43\x6d\x64\x0e\x14\xd1\xa9\xeb\xf7\xb9\x39\x07\x7c\x4d\xa4\x39\xe5\x11\x17\xa2\xdd\x27\x7c\x34\xc5\x6d\x0d\x3c\x9b\x87\x66\x59\x69\xe4\xa3\x79\x17\x00\x74\xaa\xb1\xea\x4a\xaf\x09\xf3\x4d\x82\xf5\x80\x3a\x8a\xb6\xc5\x80\xfa\x1d\xcf\x9e\x72\x67\x51\x6a\x1b\xc8\xe6\xaf\xf0\x7b\x03\x21\x62\xe1\x22\x3d\x2a\x68\x2c\x35\x19\xde\x73\xc8\x9d\x99\x4f\x62\xc2\xdb\xcf\xb6\xb3\x77\xf7\x44\x1c\x72\x9b\x83\x45\xf0\xa0\x06\xe4\x10\x2c\x06\x3e\x9a\x2a\x89\x72\xb0\x24\x46\x93\x31\xb1\xd6\x41\x2f\x0d\xc4\xd4\xa4\x50\x49\x6d\xba\x80\x12\x1d\x7e\xa2\x47\x22\xa9\xe0\xa3\x79\x17\x45\x36\xd5\xab\x7e\xca\xb3\xd9\x7e\x23\xc5\xda\x93\x04\x3a\xc9\xf5\x7e\xb3\x9d\x13\x79\x41\x47\xc5\x01\x5a\xea\xf7\xd8\xbb\x4c\x8a\xb4\x4b\x31\x51\x38\x0d\xa5\xdb\x68\x5b\x24\xc9\xae\xc0\x98\x90\x28\x6e\x1b\x97\x94\x5a\x41\xbf\x0b\x28\x1c\x72\x82\x75\x9b\x8a\x41\x29\xcd\xea\xc3\x24\x90\xb0\x76\xda\xa2\x24\xe0\xef\x32\xd2\x12\x2d\xe6\x23\x2f\x9d\x0b\xf0\x8b\xf4\x71\xdb\x1f\x65\x21\x62\xee\x9c\xea\xaa\x48\x1c\x96\xe0\x07\xc6\x15\xda\x0e\x5a\x04\xc7\x9e\x2f\xfb\x68\x1d\xfc\xd6\x51\xf1\xa2\x8b\x95\xf3\x15\x17\x29\xc1\x23\x10\xde\x99\x01\x1d\xbb\xf5\xa9\x2c\x44\x0e\xfd\x31\x9e\x85\xf3\x85\x3b\xc1\x73\x68\x9a\xf6\x40\x76\x72\x4a\x28\x0c\xac\xf0\x40\x20\x91\x6a\x5b\x37\x78\xd8\x03\x91\xd2\xa3\x79\x52\x37\xc6\x44\xdd\x58\xab\x6c\x6a\xdf\x35\xda\xd3\x3e\x4b\x74\x90\xda\x80\xda\x75\x87\x5d\x97\x25\xd2\xed\xdd\x44\x5b\x4f\x7d\x9b\x38\xfb\x0a\xdb\xda\x43\x08\x87\xa3\x02\x3e\x6a\x96\x07\xc5\xab\x8e\x35\x27\x1a\x70\xc2\xa3\xfb\xf9\xb1\x82\x71\x05\x59\x98\xb9\xaa\x36\x80\xa0\xfe\xa7\x7e\x9b\x86\x26\xcf\x21\x04\x3a\xb2\xb7\x57\x6d\xeb\xb8\x9f\xf3\xfb\x79\x34\xaf\x57\x7d\xe5\xf2\x86\xce\x8d\xa4\x00\xfc\xde\x1e\x21\xd3\xed\xcf\x2a\xf6\x93\x43\x2f\x7b\x7e\x6e\x0d\x3e\xe7\xb7\xa3\xdb\xbb\xe5\x28\x1a\x3a\x3a\xe7\xb2\x29\x98\x00\xfb\x3a\x05\x9e\xcd\xf7\x86\x73\x0f\x12\xa1\xb3\x4d\x76\x8f\xce\x82\x6c\x1a\x12\xa9\xd4\xf7\x35\x58\xfc\xab\x0e\x08\x16\x3c\xcf\x16\xef\x85\xb4\xce\x38\xa9\x62\xff\x9f\xed\x3c\x73\x31\xbb\xba\xa1\x12\x84\xc4\x83\x54\xdb\x1f\x28\x71\xd7\xee\x43\xe2\x6c\x14\x06\x12\xe6\xa5\xb4\x05\x11\x62\xb7\x94\x32\x1b\x12\xdd\xb1\xf0\xbf\x43\x8d\xaa\xc1\xe7\xad\xee\x69\xb3\xce\x95\x7d\x09\x49\x6e\x5c\xa3\x56\x46\x7a\xe8\x1a\xb6\x7c\x91\x6f\x5c\x2c\x8c\x5e\x06\x2e\x16\x2f\xbf\x37\xe0\xb7\x5c\x2c\xb2\x64\x94\xa4\x7b\x41\x52\x69\x9b\xbc\x84\xd6\xc3\x3e\x5f\x34\xd8\x25\xb2\xae\xc1\xaa\x59\xa9\x8d\xea\x87\xae\x64\xd7\xfd\xeb\xde\xe5\xc1\x4e\x9e\x8d\x54\x7f\x6e\xfc\x5a\x38\x63\xdc\x26\x4e\x77\x47\x83\x5d\x60\x12\x69\x48\x6a\x07\xa2\x08\x7c\x3c\x1c\x16\x1a\xcb\x66\x49\x40\x87\x06\x5c\xee\xd6\x03\x05\xeb\xe1\x3e\x24\xca\x6e\x53\x0f\x97\xc6\x2d\x87\x95\x0c\x08\x7e\x98\xb4\x4b\x86\x3f\x39\x57\x18\x78\xca\x73\xd7\x58\xfc\xa1\x0b\xfb\xb3\x4d\x2a\xd5\x63\x28\x7d\x01\xf8\xd0\x7b\x5e\x1a\x69\x5f\x7b\x8f\x67\xb6\x08\xdc\x9f\xc4\x33\x07\xda\x23\x11\xcf\xf1\xcf\x8f\xdd\x28\x7b\x31\x8b\xdd\x7b\xb6\x29\xc1\xd2\xb8\xc9\xa4\x07\xa6\x9c\x85\xe4\x42\x00\x93\x21\x8d\xcf\x93\xa1\xd2\x6b\xfa\xea\xa6\xf2\x61\xfb\xfb\xf2\x3f\x01\x00\x00\xff\xff\x0f\x70\x7c\x4d\x70\x0e\x00\x00")

func utilsDistResetTadpolesPasswordHtmlBytes() ([]byte, error) {
	return bindataRead(
		_utilsDistResetTadpolesPasswordHtml,
		"utils/dist/reset-tadpoles-password.html",
	)
}

func utilsDistResetTadpolesPasswordHtml() (*asset, error) {
	htmlBytes, err := utilsDistResetTadpolesPasswordHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "utils/dist/reset-tadpoles-password.html", size: 3696, mode: os.FileMode(420), modTime: time.Unix(1587248355, 0)}
	a := &asset{bytes: htmlBytes, info: info}
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
	"utils/dist/reset-tadpoles-password.html": utilsDistResetTadpolesPasswordHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
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
	"utils": &bintree{nil, map[string]*bintree{
		"dist": &bintree{nil, map[string]*bintree{
			"reset-tadpoles-password.html": &bintree{utilsDistResetTadpolesPasswordHtml, map[string]*bintree{}},
		}},
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
