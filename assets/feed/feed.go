// Code generated by go-bindata.
// sources:
// templates/feed/atom_1.0.xml
// templates/feed/rss_2.0.xml
// DO NOT EDIT!

package feed

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

var _templatesFeedAtom_10Xml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x54\x5d\x6b\xdb\x48\x14\x7d\xde\xfc\x8a\x8b\x36\xb0\xbb\x6c\xa3\x91\xdd\x3a\x25\x42\x36\xb4\x29\x81\x40\x3f\x1e\x9a\x52\xe8\x4b\x18\x6b\xae\xa2\xa1\xf3\x21\x66\xae\xb0\x1d\x57\xff\xbd\xcc\x48\x8e\x65\x27\x90\xd0\x3c\xc5\x33\xe7\xdc\x8f\x73\xce\x68\xbb\x05\x81\x95\x34\x08\x49\x85\x28\x6e\x39\x59\x7d\x3b\xc9\x12\xe8\xba\x93\x22\x9c\xc0\x5a\x2b\xe3\xe7\x49\x4d\xd4\xe4\x8c\xad\x56\xab\x74\xf5\x3a\xb5\xee\x8e\x4d\xb3\x6c\xc6\xde\x91\xd5\x49\x8f\xc9\xeb\xa7\x50\x93\x8b\x8b\x0b\xb6\xae\x49\xab\x64\x71\x02\x50\x48\xb1\x08\x20\x9f\x33\x26\xa4\xa7\x74\x55\x5b\x6f\x4d\x25\x9d\xa7\x88\xdf\x6e\x21\xbd\xd9\x34\x08\x5d\x57\x30\x29\x22\x87\x24\x29\x5c\x7c\xaf\xed\x3f\x1e\xbe\x18\xb8\x0a\x60\x18\x01\x21\x54\x72\x72\xd9\x92\xb4\xc6\x17\xac\xc7\x07\xa6\x92\xe6\x27\xd4\x0e\xab\x7e\xb4\x97\x74\x65\x09\x38\x54\xf3\x84\x2b\x42\x67\x38\x61\x02\xb1\xde\x3c\x39\x1c\x20\x01\xda\x34\x38\x4f\x08\xd7\xc4\xe2\x7e\xec\x8f\x5b\x06\xd5\xd3\xb5\x56\x43\x6b\x8f\xaa\x7a\xa6\x2b\x6f\x1a\x25\x4b\x1e\x16\x8e\xec\xff\xd7\x0f\x03\xb4\x8d\xe0\x84\x62\x11\x1a\xbc\x6f\xa5\x12\x1f\x38\x61\x7a\x65\x9d\xe6\x04\xc9\x34\xcb\xce\xcf\xb2\xc9\x59\x36\xbd\x99\xcc\xf2\xec\x4d\x9e\xcd\x7e\x64\x6f\xf3\x2c\x5a\x5e\xb0\x1d\x39\x14\xe2\x2d\xd5\xd6\x85\x7f\x01\x0a\xc3\xf5\xb1\x05\x05\x8b\x87\x01\xca\xf6\xd8\xc2\xb7\xcb\x38\xfa\x48\x9f\xe4\xd8\x3c\xe9\x81\xc3\x1d\xbf\x47\x22\x74\x60\x2b\xe0\x4a\x01\xd5\x08\x8d\xe2\x25\xfa\xb4\x60\xbb\x2a\x7d\x7b\x08\x76\x3b\x6e\xee\x10\x4e\xe5\x2b\x38\x95\x04\xf9\x1c\xd2\x6b\x42\xed\x43\x00\xe2\xdf\x80\x2c\xd0\x90\xdb\x0c\x63\xbf\x20\x6e\xa7\x0f\x36\x84\x1f\x92\xd2\xcf\x5c\xe3\xa5\xd5\x8d\x43\xef\x51\x40\xd7\xfd\x3d\x5c\x7c\xad\xf9\x74\x76\x7e\x70\x15\x33\xfa\x57\xdf\x2b\x3a\x7f\x1c\x8a\x17\x46\xe1\xf9\x21\x1e\xc5\x92\x0d\x2b\xf6\x32\x8d\x68\x71\xac\x91\x78\x45\xd3\x2e\x95\xf4\x75\x9f\x89\x80\xfa\xc8\x3d\x7d\xb2\x42\x56\x72\x58\x62\x8f\xe8\x19\xa3\x0c\xed\xf0\xdf\xe2\xd1\x71\x48\x00\x8a\xd2\x1a\x42\x43\x07\x76\xdf\xd4\xd2\x1f\xbc\x4a\x08\x28\x2e\x8d\x87\xa1\xe4\xa5\x6d\x0d\xc1\x2f\xa8\x5b\xcd\x8d\xbc\xc7\xdb\xd2\x6a\xcd\x83\x97\x47\x51\x71\x58\x5a\x27\x3c\x70\x23\x42\x6c\x76\x56\xc8\xfb\xb1\x3c\xa3\x3a\xcb\x0d\x61\xcc\x44\xb9\xbf\x0e\xdc\x11\xf1\x11\xbc\xeb\xa0\x35\x7b\x7c\x0a\xd7\x04\x2b\xee\xa1\x74\x18\x16\x85\xa7\x65\x8b\x65\xff\x0d\xa9\x15\x9c\x38\x48\x0a\xef\xf6\xbf\x48\x54\xdc\x13\x0c\x32\x81\x35\xf0\x94\x8e\x69\xc1\x06\xe9\xfa\x27\xf4\x10\xdb\xed\x16\xd0\x88\xf8\x05\x66\xe1\x13\xbc\xd8\x1f\xfc\x0e\x00\x00\xff\xff\x9b\x6f\xcf\x71\xaf\x05\x00\x00")

func templatesFeedAtom_10XmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesFeedAtom_10Xml,
		"templates/feed/atom_1.0.xml",
	)
}

func templatesFeedAtom_10Xml() (*asset, error) {
	bytes, err := templatesFeedAtom_10XmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/feed/atom_1.0.xml", size: 1455, mode: os.FileMode(436), modTime: time.Unix(1534173322, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesFeedRss_20Xml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x53\x5d\x4f\xdb\x30\x14\x7d\x1e\xbf\xe2\x28\x43\xda\x26\xa1\x38\x74\x03\x89\xca\xf4\x61\x4c\x48\x4c\x63\x7b\x80\x69\x8f\xc8\xc4\xb7\x8d\xb5\xc4\xae\x7c\x6f\x56\x20\xcb\x7f\x9f\x9c\x96\x36\xb0\x49\xec\xad\xee\x3d\x1f\xce\xb9\xc7\x5d\x07\x4b\x73\xe7\x09\xd9\x9c\xc8\xde\x44\xe6\x9b\x49\x91\xa1\xef\xf7\x74\x64\xc6\x5d\x53\x7b\x9e\x56\xa7\x59\x25\xb2\x9c\x2a\xb5\x5a\xad\xf2\xd5\xfb\x3c\xc4\x85\x3a\x3c\x39\x39\x51\x77\x95\x34\x75\x86\x5f\x14\xd9\x05\x7f\x9a\x4d\xf2\x22\x9b\xed\x01\xba\xac\x8c\xf7\x54\xa7\xdf\x80\x16\x27\x35\xcd\x7e\x54\xe1\x0d\xe3\x9b\xc7\xb9\x8b\x2c\xe8\x3a\xe4\xd7\xf7\x4b\x42\xdf\xc3\x3a\x96\xe8\x6e\x5b\x71\xc1\xb3\x56\x6b\xfc\x9a\x5b\x3b\xff\x73\x96\xec\x79\xaa\x54\xc2\xe5\xab\x2a\x70\xf0\xf3\x24\x32\xdc\x64\x24\xa4\xb4\x1a\xf0\x1b\xaa\x61\xf9\xd8\xba\xda\x7e\x32\x42\xb3\x04\xdb\x9e\xf2\xf3\x10\x1b\x23\xc8\x2e\x83\x3f\x40\x31\xc1\x67\xe3\x31\x29\x8a\x63\x1c\x1e\x4d\x8b\x0f\xd3\xe2\x08\x97\x57\xd7\x29\x08\xad\x9e\xca\xac\xa5\x2d\x71\x19\xdd\x32\xdd\xf7\xf9\x77\x39\x86\xc1\xc2\x3c\x90\x08\x51\x44\x98\xc3\xd4\x35\xa4\x22\x2c\x6b\x53\x12\xe7\x5a\x8d\xd9\x83\x1e\x52\x1a\xd1\xf8\x05\x61\xdf\x1d\x60\xdf\x09\xa6\xa7\xc8\x2f\x84\x1a\x4e\xf9\x00\xd8\xe0\xb4\x13\x6a\x1e\x49\x7a\xd1\x3a\xfb\x62\x38\xfb\xdb\x74\xd2\xc1\x49\xfe\xd5\x34\x74\x16\x9a\x65\x24\x66\xb2\xe8\xfb\xd7\x9b\xc1\x55\x65\x26\x47\xc7\x4f\x46\x5a\x0d\x1e\x6b\xbf\x57\x8f\xbe\xff\xb5\x94\x97\x7d\xc7\xdb\xda\xd6\x64\x04\x1e\x20\xa3\x32\x3c\x0b\xfe\xba\x72\xfc\xa4\x3a\x28\x83\x17\xe3\x3c\x63\x23\x72\x16\x5a\x2f\xf8\x8d\xaa\x6d\x8c\x77\x0f\x74\x53\x86\xa6\x31\x29\xd1\x67\x4b\x8b\x54\x86\x68\x19\xc6\xdb\xb4\xc0\xc7\x3c\xdc\xc3\xf8\xc2\x23\x9d\xdb\x7b\xa1\x61\x33\xe5\x6e\x9c\xb8\x23\xe2\x5f\xf0\xbe\x47\xeb\x77\xf8\x1c\x17\x82\x95\x61\x94\x91\x8c\xd0\x96\xfb\xc5\xb0\x5c\x06\xeb\xe6\x6e\xc8\x68\x90\x7d\x9b\xea\x63\x8d\x18\x38\x61\xaa\xe7\xef\x06\x62\x2a\x26\xda\xa5\x1d\xd8\xc1\x8f\x05\xbe\x0f\xff\xa2\xef\xff\xd5\x36\xad\x76\x25\xea\x3a\x90\x4f\x3e\xe9\xd5\xaa\xed\xb3\xd5\x2a\x32\xcf\x76\xc3\x3f\x01\x00\x00\xff\xff\x3a\x0d\x82\xc3\x29\x04\x00\x00")

func templatesFeedRss_20XmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesFeedRss_20Xml,
		"templates/feed/rss_2.0.xml",
	)
}

func templatesFeedRss_20Xml() (*asset, error) {
	bytes, err := templatesFeedRss_20XmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/feed/rss_2.0.xml", size: 1065, mode: os.FileMode(436), modTime: time.Unix(1534173010, 0)}
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
	"templates/feed/atom_1.0.xml": templatesFeedAtom_10Xml,
	"templates/feed/rss_2.0.xml": templatesFeedRss_20Xml,
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
	"templates": &bintree{nil, map[string]*bintree{
		"feed": &bintree{nil, map[string]*bintree{
			"atom_1.0.xml": &bintree{templatesFeedAtom_10Xml, map[string]*bintree{}},
			"rss_2.0.xml": &bintree{templatesFeedRss_20Xml, map[string]*bintree{}},
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

