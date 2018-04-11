// Code generated by go-bindata.
// sources:
// clarch/templates/entity.tpl
// clarch/templates/handler/http/hander_test.tpl
// clarch/templates/handler/http/http.tpl
// clarch/templates/project/config.tpl
// clarch/templates/project/handler.tpl
// clarch/templates/project/middleware/middleware.tpl
// clarch/templates/repository/repository.tpl
// clarch/templates/usecase/usecase.tpl
// DO NOT EDIT!

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

var _clarchTemplatesEntityTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\xcc\x31\x8e\x83\x40\x0c\x85\xe1\x1a\x9f\xc2\x72\xb9\x05\x07\xa0\x5b\xb1\xcd\x76\x5b\x6c\x6a\xc6\x62\x2c\xe4\xc0\xc0\x64\x30\x52\x10\xe2\xee\x51\x18\xa2\x54\xa9\xff\xef\xbd\xc8\x6d\xcf\x9d\xe0\xb6\x95\x7f\x7d\xb7\xef\x00\x1a\xe2\x94\x0c\xc9\x34\x08\x01\xd8\x1a\x8f\x5a\x73\x90\xe1\x20\x38\x5b\x5a\x5a\xc3\x0d\x8a\x5f\x8f\x3a\x1a\xba\xeb\x3c\x8d\x15\xa9\x27\xec\xa6\x14\x2a\x8a\x49\x03\xa7\xb5\xe9\x65\x25\x07\x50\xd4\x49\xd8\xc4\x7f\x1b\x3e\x5f\xcb\x7f\x0d\x82\xaf\x55\x9b\x5b\xc3\x46\x0e\x8a\x4b\xf4\x1f\xe5\x92\xdb\x29\x7f\x64\x90\x2c\xbf\xde\xf4\x94\x3e\xb7\x86\x0d\x69\xbe\x0d\x15\xe9\xe8\xe5\x4e\x0e\x76\x78\x04\x00\x00\xff\xff\x23\x0b\x46\x5a\xf1\x00\x00\x00")

func clarchTemplatesEntityTplBytes() ([]byte, error) {
	return bindataRead(
		_clarchTemplatesEntityTpl,
		"clarch/templates/entity.tpl",
	)
}

func clarchTemplatesEntityTpl() (*asset, error) {
	bytes, err := clarchTemplatesEntityTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "clarch/templates/entity.tpl", size: 241, mode: os.FileMode(420), modTime: time.Unix(1523361096, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clarchTemplatesHandlerHttpHander_testTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x52\x4f\x6f\x9c\x3e\x10\x3d\xe3\x4f\x61\x71\xf8\x09\x22\x7e\x26\xe7\x48\x39\xa4\x64\xa3\x74\x5b\x25\xd1\x26\x7b\xaa\xaa\xca\x98\x59\x96\x05\x6c\x77\x3c\x34\x1b\x21\xbe\x7b\x65\xf6\x8f\x36\x74\x5b\xa9\xbd\x80\x99\x99\xe7\x79\xef\xf1\xac\x54\xb5\x2c\x81\xaf\x89\x2c\x63\x55\x6b\x0d\x12\x8f\x58\x10\x6a\xa0\xd4\x17\xc3\x93\xf3\xf8\x20\x70\xe4\x8b\xfe\x5d\xe9\x32\x64\x2c\x08\xcb\x8a\xd6\x5d\x2e\x94\x69\xd3\xbe\x17\x59\x87\x08\x9a\x96\x0e\x70\x18\x4e\x0a\x0b\xb0\x66\x18\xd2\x8d\xc9\xc3\x7f\xc0\xa4\x9d\x03\x25\x1d\xa4\xad\x51\xb5\x9b\xdc\x90\x6f\x95\x29\x40\xa5\x2b\x59\x03\x4e\x7a\x65\xa5\xff\x2f\x8d\xae\x94\x3f\x4d\x7a\xf5\x65\xdd\xe5\x9d\x4e\xad\x9d\x34\x1c\x21\x90\x5a\x63\x3a\xca\x5c\xbd\xa5\xd2\x39\x40\x0a\x59\xcc\xd8\xaa\xd3\x8a\xbf\x80\xa3\xbb\x4a\x17\x37\x4d\x13\x11\xbf\xd8\xbb\x21\x5e\x62\xde\xb3\xe0\x87\x44\xde\xf2\xbe\x17\x4f\x75\x39\x0c\xc2\xcb\x91\x2d\x34\xe3\x17\x0b\x00\x91\x5f\x5d\xf3\x91\xab\xb8\x93\x35\xdc\x4a\x92\xd1\x7f\x6d\xcc\x82\xdd\x16\xf1\x60\x66\x88\x06\x23\x4a\x38\x20\xc6\x8c\x05\x5e\xf4\x32\x93\x0e\x3c\x52\xc3\x6b\x34\xba\x20\xe6\x26\x5f\xee\x6c\x89\x77\x33\x73\x93\x7f\xae\x1c\xf9\xa9\x56\xd6\x10\x7d\xf9\x7a\x71\x9e\x46\xc2\x2f\x27\x90\x6b\x2e\xad\x05\x5d\x44\x27\xc5\x84\x8f\xb4\x8e\xdb\xc5\xa3\x8e\xc2\xbd\xec\x30\x16\x0b\xa0\x0e\xf5\x7b\x80\xae\x1a\x4f\x18\xc1\x79\x12\x87\xc4\x88\x07\x78\x5d\x80\x32\x58\x00\x46\x31\x0b\x54\xc2\xbf\xf9\x7e\x59\x69\x91\x21\x48\x02\x6f\x68\x66\x34\xc1\x96\x22\x04\x17\xb3\x60\x2d\x75\xd1\xc0\xe8\xd5\x3d\x91\x9d\x9b\xfc\x7e\x57\xe9\x59\x10\x1c\x54\x2d\xb3\x2b\x7e\xa4\x97\xb0\x60\x38\xe2\xc4\xe1\xf7\x28\xcf\xc7\x5a\xf1\x84\x95\xa6\x46\xef\x6f\xdf\x5b\x3d\xfb\xde\xc9\xc6\x1b\xed\x99\x8a\x67\x92\xd4\xb9\xc7\x4f\x09\x47\x70\x22\x33\x05\xbc\x53\x7f\x33\x62\x66\x5b\x0b\x8a\x24\x55\x46\xbb\x88\x62\x36\x4c\x22\xf1\xe1\xed\x63\x71\x3e\x13\x9b\xbf\xcd\xc4\xe6\xf7\x99\x38\xdd\xfa\x4c\x06\xe1\x97\x95\xa7\x13\x4b\x5b\x48\xfa\xf3\xc8\x2d\x34\x70\x76\xe4\x67\x00\x00\x00\xff\xff\x45\x1d\x9a\xe0\x24\x04\x00\x00")

func clarchTemplatesHandlerHttpHander_testTplBytes() ([]byte, error) {
	return bindataRead(
		_clarchTemplatesHandlerHttpHander_testTpl,
		"clarch/templates/handler/http/hander_test.tpl",
	)
}

func clarchTemplatesHandlerHttpHander_testTpl() (*asset, error) {
	bytes, err := clarchTemplatesHandlerHttpHander_testTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "clarch/templates/handler/http/hander_test.tpl", size: 1060, mode: os.FileMode(420), modTime: time.Unix(1523351651, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clarchTemplatesHandlerHttpHttpTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x95\x4f\x73\xda\x3c\x10\xc6\xcf\xd6\xa7\xd0\xeb\x43\xc6\xce\xf8\x95\x27\x57\x3a\x39\x94\x3f\x05\xda\x4e\x60\x42\x68\xcf\xaa\xb4\x31\x6a\x8c\xec\xca\x6b\xda\x0c\xe3\xef\xde\x91\x65\x43\x48\x4d\xa6\x94\x76\x26\xb9\x04\xa4\xdd\x47\xfb\x7b\x76\x91\x72\x2e\x1e\x78\x02\x74\x85\x98\x13\xa2\xd6\x79\x66\x90\x06\xc4\xf3\x35\x60\x6c\x17\x7d\x42\x3c\x3f\x51\xb8\x2a\xbf\x30\x91\xad\xe3\xed\x96\x0d\x4a\x63\x40\xe3\xb2\x00\x53\x55\x4f\x16\x6e\x21\xcf\xdc\xc2\xfc\x21\xa9\x2a\xff\x4f\x13\xe3\xb2\x00\xc1\x0b\x38\x59\x20\x37\xd9\x57\x10\xf8\x2c\x2f\x51\xfa\xff\x24\xd3\x4a\xd8\x4f\x3e\x09\x09\x21\xf8\x98\x03\x9d\x20\xe6\x56\x82\xaf\x21\xad\xcf\x9d\x70\x2d\x53\x30\xb4\x40\x53\x0a\xa4\x5b\xe2\x35\x82\xac\xd9\x99\xbb\xaf\xc4\x6b\x2b\x5d\x0e\x68\x53\x2b\x3b\x50\x5a\xba\x45\x52\x11\x72\x5f\x6a\x41\x03\x5c\xa9\x82\x5e\x1e\x3b\x31\xa4\xef\x94\x96\x6f\xd3\x34\x10\xf4\x32\x51\x9a\x0d\x32\x8d\xf0\x03\x43\x5b\x43\xaa\x0a\x8c\x28\x18\x43\x7b\xd7\xd4\xea\xb0\xfd\xe9\xac\xcd\x0b\x89\xa7\xee\xeb\xa0\xff\xae\xa9\x56\xa9\x4d\xf4\x04\x7b\xbf\x98\xdd\x04\xb6\x8b\x6c\x81\x1c\xcb\xa2\xcf\xe5\x2d\x7c\x2b\xa1\x51\x0c\x89\xe7\x19\xc0\xd2\x68\xe2\x55\x84\x74\x24\xcc\x3e\x44\xd4\x16\x10\x9e\x86\xd2\x7f\x9c\xca\x2e\x16\x25\x2d\x84\x60\x63\xc0\xa9\xc6\xc0\x57\xd2\xb7\x95\x23\xac\x5f\x24\xac\xe5\x94\xfc\x1b\x90\x47\x18\x6d\x09\x27\x30\x2e\x30\x33\xd0\x09\x88\xb0\xb6\x14\x17\x2d\xc1\xe1\x5c\x6c\xab\x1d\x42\x8b\xfa\x89\xa7\x4a\x72\x84\xbe\xd2\xb2\xae\x4c\x44\xf4\xa2\x2e\xe7\xcd\xd9\xfd\x34\x50\x1c\xf5\xd5\x21\xb8\x93\xfe\x9d\xaf\x06\x8a\x13\x6c\x5d\xe6\xd6\x89\x0e\x5f\x69\xf3\xe7\xe6\xa7\x46\x99\x73\x53\xc0\x74\x18\x88\x70\xbf\xfb\xdb\x73\xb4\xcb\x38\x0b\xbb\x55\xb1\x56\xbf\x92\xae\x36\x0e\x2a\xd9\x8e\xf4\x2b\xe9\xec\x10\x52\xe8\xec\xac\xd7\xdd\xd2\xe7\x7e\x3e\x41\x6c\xa4\x94\x3c\xdb\x4a\x4f\xb0\xcf\x46\x21\x18\xf7\x6f\x02\x5c\x82\x39\xe4\xdc\x23\xde\xc0\xf7\x43\x34\xc4\xbc\xc1\x0b\x12\x87\x35\xd2\x89\xd2\x10\xd1\x52\xbc\xf8\x2a\xd4\xd8\xab\xfa\x96\x38\x66\x98\x85\xd9\x23\xf7\x68\x29\xa2\xba\xe0\xcd\x95\x4d\x4b\xd8\xd8\x64\x65\x1e\xf8\x31\xcf\x55\xbc\xb9\xb2\x77\xe8\xe6\x8a\x8d\x47\x77\x81\xbf\x7b\x44\x0b\x3f\xa2\xab\xf6\x7d\xe8\x0c\x88\x7b\x4a\xee\x82\xec\x4f\xc3\x45\xcd\x67\x8b\x5f\x75\xea\xdb\xa2\xd9\x5f\x76\xab\xb8\xd1\x73\x31\xc3\xd1\xc7\xd1\xdd\xa8\x33\xcc\xb5\xcf\xfa\xfa\x33\x00\x00\xff\xff\x9a\x8a\x9e\x5e\x7b\x08\x00\x00")

func clarchTemplatesHandlerHttpHttpTplBytes() ([]byte, error) {
	return bindataRead(
		_clarchTemplatesHandlerHttpHttpTpl,
		"clarch/templates/handler/http/http.tpl",
	)
}

func clarchTemplatesHandlerHttpHttpTpl() (*asset, error) {
	bytes, err := clarchTemplatesHandlerHttpHttpTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "clarch/templates/handler/http/http.tpl", size: 2171, mode: os.FileMode(420), modTime: time.Unix(1523361375, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clarchTemplatesProjectConfigTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\xd1\x6b\xdb\x3e\x10\x7e\x96\xfe\x0a\xfd\xf4\x50\xec\x52\x9c\xf4\x57\xd8\x82\x47\x1e\x42\xdb\xb1\x40\xb7\x19\xb6\x3d\x8d\x41\x15\xe7\xe2\x68\x71\x24\xf5\x24\xb7\xb4\x25\xff\xfb\x38\xc9\xc9\x9c\x31\xd6\x1a\x42\xac\xbb\xef\xfb\xee\xf4\xdd\x61\xa7\xea\x8d\x6a\x40\x38\xb4\x3f\xa1\x0e\x9c\xeb\xad\xb3\x18\x44\xc6\x99\x6c\x6d\x23\x39\x67\xb2\xd1\x61\xdd\x2d\x8a\xda\x6e\x47\x9b\xf1\xa6\x5b\x74\x66\xe4\x9c\xfc\x23\x01\xad\x87\xc7\xb5\x6e\xd6\xc1\x3e\x00\x8e\xc0\xdc\xd7\xd6\xac\x74\x13\x71\xb6\xe8\x16\x80\x85\xc5\x66\xf4\xa4\xdc\x5f\x42\xf4\xab\x2d\x82\xe4\x39\xe7\xf7\x0a\xc5\x65\x24\x8b\x53\x12\xe1\x3c\x3c\x3a\x10\xb3\x6a\x2e\x7c\xc0\xae\x0e\xe2\x99\xb3\x8a\xba\xf4\x01\xb5\x69\xc4\xed\x12\x56\xaa\x6b\x43\x29\x27\x93\xc9\x44\xde\xf2\x5d\x4f\xf9\xf8\xe8\xef\xda\x01\xe9\x83\xf5\x41\xd0\xb3\x27\x22\xdc\x75\x1a\x61\x59\xca\x80\x1d\x48\x71\x10\x3a\xff\xff\x6d\x31\x2e\xc6\xc5\xb9\xbc\xe5\xec\x9b\x07\x7c\x1d\x0d\xad\x0d\xc4\xa8\x94\xf7\x0f\x16\x97\x2f\x33\x5c\x8f\x24\xd6\x95\x0a\x6a\xa1\x3c\xbc\xa2\x0e\xd4\x1b\x6b\x00\x63\x2d\x72\x62\xd8\xdd\x01\x75\x71\x31\x7e\x33\x70\xe3\xc6\x36\x0d\x60\x6f\xc7\xf3\x3e\x4a\x0e\x0f\x2c\x9a\x55\x73\xce\x7a\xe4\xe9\x93\x72\x45\x9a\x04\x67\xd1\x4b\xd2\x5a\x75\xa6\x16\x59\x58\x6b\x9f\xc6\x93\x8b\x0a\xb5\x09\x59\x4e\x74\xe7\x8a\x78\x6a\x4d\x44\xe4\x07\xc2\xdc\xe8\x90\xa4\x12\x90\x86\xac\x9c\x16\xb1\x9e\x5e\x09\x40\x14\xe5\x54\x1c\xf6\xa6\xa8\xd0\xd6\xe0\x7d\x26\x95\xd3\xf2\x4c\x9c\x28\xa7\xf3\x77\x11\xf6\xdf\x54\x18\xdd\x92\x08\x6b\x6d\x53\xbc\x57\x41\xb5\x19\x20\xe6\x9c\xed\x78\x52\xde\xc6\xc1\xa7\x96\xff\xad\x1e\x91\xa4\x1f\x5f\x5e\x57\xa1\x5f\xce\xa9\x38\x21\x35\x42\xcd\xaa\x79\x29\x04\x5d\xe8\x8c\xb3\x64\x55\x99\x9a\xa0\x73\xb2\xb3\x14\x64\x44\x96\x0f\x6d\x8d\x25\x58\x1c\x41\x39\x15\xbf\xe3\x31\xcc\x6e\xe0\x1e\xda\x52\xec\x1f\x4a\x7f\x82\x87\x59\xb0\x5b\x5d\xc7\xdc\x2c\x64\x14\xbc\x82\x45\xd7\xc4\x40\x7e\x16\x89\x57\xf4\x6e\xdd\x16\x4c\x20\x3a\x2d\x4f\x4a\x5c\x9b\xda\x2e\xb5\x69\x7a\x51\x59\x5b\xe3\x6d\x0b\x72\x90\x05\x4c\x2d\x94\xfb\x7a\x03\xb1\x23\x40\xd6\x17\xfb\xa2\xb6\xae\x8d\x9a\x27\xc4\xd8\x1f\x87\x17\x61\x34\x7d\xad\xd2\x5d\xce\xc7\xe3\x44\x64\x5f\xd7\x80\xa0\x56\x81\xbc\x39\x44\x77\xe9\xef\x73\x17\x5c\x17\x2a\x15\xd6\xbe\x14\xdf\x7f\xa4\xdd\xee\xd5\xa4\x0f\x4b\xdb\x05\x79\x44\xb8\x46\xb4\xf8\x12\x0b\x10\x8f\x58\xbb\xbd\xfd\xc5\xd1\xd5\xfa\x53\xb4\x54\xc4\xb9\xd0\xc7\xa9\xb8\x54\x4e\x07\xd5\x5e\xda\xd6\x62\xcc\xf5\x24\x12\x41\x08\x1d\x9a\xb4\x11\x9c\xb1\x5d\x74\x67\xc7\x77\xfc\x57\x00\x00\x00\xff\xff\x16\x0a\xa5\x0c\x62\x05\x00\x00")

func clarchTemplatesProjectConfigTplBytes() ([]byte, error) {
	return bindataRead(
		_clarchTemplatesProjectConfigTpl,
		"clarch/templates/project/config.tpl",
	)
}

func clarchTemplatesProjectConfigTpl() (*asset, error) {
	bytes, err := clarchTemplatesProjectConfigTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "clarch/templates/project/config.tpl", size: 1378, mode: os.FileMode(420), modTime: time.Unix(1523355320, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clarchTemplatesProjectHandlerTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8f\x41\x4f\x83\x40\x10\x85\xcf\xcc\xaf\x18\x39\x18\x68\x2a\xbd\x6b\x7a\xd1\x8b\xf1\xa0\x26\x4d\xbc\x9a\x75\x19\x96\xb1\x74\x96\x0c\x43\xa3\x21\xfc\x77\x03\xb5\x46\x6e\xbd\xbd\x7d\x79\xdf\xbe\x37\xad\xf3\x7b\x17\x08\x5b\x8d\x9f\xe4\x0d\x80\x0f\x6d\x54\xc3\x0c\x92\x34\xb0\xd5\xfd\x47\xe1\xe3\x61\xe3\x3a\xd7\xed\xe9\xc8\xbe\xde\x84\x78\x74\x0d\x97\xce\xa2\xa6\xcb\x50\x60\xb9\x09\x51\xd8\x4f\x2a\x85\x1c\xc0\xbe\x5b\xc2\x47\x27\x65\x43\xfa\x7a\x6a\xc0\xce\xb4\xf7\x36\x8c\x00\x55\x2f\x1e\x33\xab\xb9\xc3\xd5\x32\x94\xe3\xdb\xa9\x83\xee\x59\xca\xa7\xdd\xcb\x73\xe6\x71\x15\x58\x8a\x87\x28\x46\x5f\xb6\xc6\x88\x2c\x46\x5a\x39\x4f\xc3\x98\x23\xa9\x46\xc5\x01\x12\xae\x26\x8d\xb7\x5b\xf4\xc5\x1f\x7c\x1d\xf3\xbb\xd9\xbe\xda\xa2\x70\x33\xe5\x12\x25\xeb\x55\x26\x17\x92\x11\x66\xf0\x7d\x7d\x66\xff\x1d\x59\x9c\xa7\xec\xe6\xe1\xd9\x05\x5f\xfd\xbe\x85\x1b\x18\xe1\x27\x00\x00\xff\xff\xa3\x28\x87\x2a\x62\x01\x00\x00")

func clarchTemplatesProjectHandlerTplBytes() ([]byte, error) {
	return bindataRead(
		_clarchTemplatesProjectHandlerTpl,
		"clarch/templates/project/handler.tpl",
	)
}

func clarchTemplatesProjectHandlerTpl() (*asset, error) {
	bytes, err := clarchTemplatesProjectHandlerTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "clarch/templates/project/handler.tpl", size: 354, mode: os.FileMode(420), modTime: time.Unix(1523351651, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clarchTemplatesProjectMiddlewareMiddlewareTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\xdf\x6f\xdb\x36\x10\x7e\x96\xfe\x0a\x8e\x40\x07\x2a\x90\xe9\x01\x7b\x1b\xd0\x01\x6e\xda\xa2\x06\xe2\x2c\x8b\x1b\xec\xa1\x28\x0a\x86\x3a\x49\x9c\x25\x52\x3d\x9d\x92\xb8\x41\xfe\xf7\x81\xa4\x22\xcb\x4d\xda\xec\x41\x30\x7d\x77\xdf\xf7\xdd\x2f\xb2\x53\x7a\xa7\x2a\x60\xad\x29\x8a\x06\x6e\x15\x42\x9a\x9a\xb6\x73\x48\x4c\xa4\x09\x2f\x5b\xe2\x69\xc2\xc9\xb4\xc0\xd3\x34\x51\x03\xd5\x9b\x10\xc9\x78\x65\xa8\x1e\xae\xa5\x76\xed\xf2\xf7\x45\x5f\xab\x1d\x2c\x11\xf4\xce\x59\xc0\x45\xd1\xd6\x8b\x1e\xf0\x06\x70\xe9\x21\xcb\x5a\xd9\xa2\x01\x5c\x1e\x54\x3c\xeb\x8c\xa1\x32\x76\x51\x39\x6b\xb4\x3f\x05\x9f\x93\xc3\x35\xa0\x74\x58\x2d\xbf\xa9\xee\x19\x93\xff\xb4\xf3\x4c\x59\x9a\xd2\xbe\x7b\xac\x81\xf5\x84\x83\x26\x76\x3f\x4f\x57\x6e\x26\xe5\xd5\x40\xf5\x87\x98\x4f\xfa\x30\x22\x0f\xde\x75\xdb\x35\xcc\x58\x02\x2c\x95\x06\x4f\x72\xe6\xaa\xca\xd8\x4a\x68\x67\x4b\x76\xf2\x4d\x75\xf2\xd4\xd9\xd2\x54\x19\xab\x8c\x95\x23\xd3\xfb\xc1\xea\x34\xf1\xd4\x60\xc9\x68\x45\x20\x9e\xfa\x1f\xd2\xb4\x1c\xac\x66\xe7\x70\x7b\x10\x14\xea\x90\x0f\x7b\x29\xe1\xec\xfb\x4c\xef\xd3\x04\x81\x06\xb4\xec\xd7\x58\xfd\xfd\x8c\xee\x61\x52\x14\x54\x9b\x9e\x9d\xc4\x90\x8c\xfd\xef\x92\x3c\x7f\xe3\xaa\x0a\x30\x67\x5f\xd8\x1f\xaf\x99\x47\xc8\x37\x83\x69\x0a\x91\x4d\xd2\x5e\x42\x68\x76\xe2\xc1\xa7\xce\x12\xdc\x51\xe6\x91\x49\x4f\x0a\xc9\xc3\xfc\xfe\xc8\x73\x77\x2b\x32\x79\xf5\xf1\xd4\x43\x13\x2d\xb7\x40\x82\x47\x76\x9e\xb3\x78\x88\x9e\x73\xb8\x23\x91\xa5\x69\x92\x34\x8a\xc0\xea\xfd\xc4\xb1\x35\x56\x83\x08\xbc\xc1\xdf\x29\xaa\x43\x5e\xf2\x12\xbe\x0e\xd0\x93\xbc\xba\x3c\x93\x17\x8a\xea\x34\x49\x5a\xa0\xda\x15\xc7\xee\x4d\xb0\x79\xe6\xa0\x27\xd7\xb6\x74\x22\x4d\x92\xa4\x6c\x49\x6e\x3b\x34\x96\x4a\xc1\x5f\xf5\xec\x55\xcf\xf3\x27\xb8\x9c\x79\xc1\x2c\xf7\x00\xdf\xb6\x2d\xa1\xef\x23\xef\x10\x4a\x73\xc7\x73\xc6\x57\x6f\xcf\x17\xab\x8b\x35\x7f\x26\x46\x51\xcd\x7f\x44\x60\xba\xa0\x76\xda\x18\xb0\xb4\xbe\x10\xd9\xd3\x90\xa1\x07\x5c\xa8\x0a\x2c\x1d\x25\x76\xd5\x03\xae\xbc\x75\x8e\x59\x5b\x12\xbc\x27\x45\x43\xac\xe2\x1f\x34\x04\x28\xb7\xc1\x32\x0f\xfc\x68\x5a\x10\xf1\x7a\xe7\x2c\xb6\xf5\x89\x70\x6c\x23\xcf\x59\x3c\x1c\x02\xde\x0e\xa8\xc8\x38\x2b\xf8\x38\x26\x3f\xc6\x78\x9a\x82\xde\x1b\x68\x8a\xc8\xd4\xbf\xd9\xc7\xc3\x46\x75\x82\xd7\xa0\x8a\x30\xf8\x43\x29\x1f\x82\xe9\x65\xe8\xd7\x01\x70\x7f\xdc\x84\xcb\x33\xf9\xb7\xb7\x8e\xb5\x85\xdd\x30\x25\x6b\xc0\x0a\x2d\xdf\x21\x3a\xec\x33\xf6\x27\xfb\x2d\xac\x65\x52\x3a\x64\x5f\x72\x06\x88\x9b\xbe\xf2\xfb\x81\xca\x56\xc0\x1e\x23\xc7\x1f\x11\x97\x78\x5a\x95\x60\x15\x11\xe4\xf7\x34\x79\x48\xc3\x77\xb8\x66\x3f\x49\x7a\x07\x7b\xff\x30\x19\x5b\xe5\xac\x65\xad\xea\x3e\xc5\x7f\x9f\x3f\x7d\x8e\x87\x8c\x8d\xef\x99\x0c\x14\xb3\x9b\xed\x5b\xfd\xd7\xf5\xbf\xa0\xc9\xb3\xe4\x23\xcd\x8c\x1d\xfb\x5a\x35\x80\xa2\xcd\xb2\x29\x95\x1f\x07\xfd\x5c\x3c\x0a\x4d\xe1\x8f\x8f\xc0\xfc\xa6\x1b\x6b\x01\xbf\x03\xbc\xb3\xda\xf9\xd9\xf9\x9e\x3a\x0c\x7d\xf3\x4d\xde\xe5\xec\x46\x35\x03\xf4\x87\x26\xb7\xb1\xa9\x80\xe8\x6d\x81\x4b\xae\x8a\x62\x85\xa8\xf6\x62\x97\x4f\xbc\xc1\x70\x94\x87\x78\x46\x3d\x44\x3d\x23\x3e\xcd\xf8\xe6\xa0\x3c\x66\x12\xdd\xc9\xa8\xdc\x75\x60\xc7\x79\x89\x9b\x30\xd5\x30\xd6\xe4\xb1\x62\x6b\x9a\x30\xea\x2c\xf8\x4c\xe9\x35\xd8\x2f\xaf\xbd\x7d\x64\x1a\x03\x01\x71\xb6\x13\x47\x70\xbf\x20\xff\x05\x00\x00\xff\xff\xce\xa9\x7b\xae\x68\x07\x00\x00")

func clarchTemplatesProjectMiddlewareMiddlewareTplBytes() ([]byte, error) {
	return bindataRead(
		_clarchTemplatesProjectMiddlewareMiddlewareTpl,
		"clarch/templates/project/middleware/middleware.tpl",
	)
}

func clarchTemplatesProjectMiddlewareMiddlewareTpl() (*asset, error) {
	bytes, err := clarchTemplatesProjectMiddlewareMiddlewareTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "clarch/templates/project/middleware/middleware.tpl", size: 1896, mode: os.FileMode(420), modTime: time.Unix(1523355479, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clarchTemplatesRepositoryRepositoryTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x94\x41\x8f\xd3\x30\x10\x85\xcf\xf1\xaf\x18\x7a\x88\x9c\x55\x94\xe5\x0c\xea\x81\x6c\x40\xda\x0b\x42\x8b\xf6\x84\x38\x84\x64\xda\x9a\x26\x76\x34\x71\x40\x25\xf2\x7f\x47\xe3\xa4\xa5\x0b\x0d\x6c\x52\x89\x1e\xdd\xf7\xde\xf8\x1b\xcf\xa4\xc9\x8b\x7d\xbe\x45\x20\x6c\x4c\xab\xac\xa1\x83\x10\xaa\x6e\x0c\x59\x90\x22\x58\x6d\x95\xdd\x75\x5f\x92\xc2\xd4\xb7\x7d\x9f\xdc\x75\x44\xa8\xed\x63\x8b\xe4\xdc\xd9\xc1\x03\x36\x66\x38\xf8\xb0\xdf\x3a\xb7\x7a\x6a\xfc\xaa\xf4\x8f\x5d\x77\xbb\x35\x54\xaf\x44\x24\x84\x3d\x34\x08\x6c\xce\x6b\xac\xbc\xe1\xe1\x54\x1c\x94\xb6\x48\x9b\xbc\x40\xe8\x45\xf0\x4e\xe9\xf2\x4d\x55\xc9\x08\xe4\xa7\xcf\x37\xc7\xf8\xe4\x89\x37\x06\x24\x32\x14\x0d\xea\xf4\x70\x5f\x4a\x55\x72\x4c\x04\xf2\x9f\x96\x8f\xd6\x10\x4e\xc9\x9e\x13\xf0\xd8\x94\xb9\x45\xa9\xb4\x8d\xe1\x8a\x98\x0c\x2b\xe4\x98\xf1\xe2\xfe\x58\xb8\x5f\xad\xfa\xbd\x4b\xad\xa5\xae\xb0\xdc\xa2\x2c\x85\x1b\xee\x6c\x92\xa5\x6c\xd8\x74\xba\x80\xf7\xf8\x7d\xa2\xbd\xf2\x4c\x1e\x4d\xbe\x41\x2f\x02\x42\xdb\x91\x86\x90\xb5\x13\xb2\x3e\x4b\xdd\xa9\xa4\xb4\x3b\xd5\x0e\xd1\x13\xf2\x08\xe6\xbc\x26\xdf\x41\x59\xac\xe1\xd5\x1a\xea\x7c\x8f\x7f\x73\xbc\x8c\x44\xa0\x36\x6c\x64\x35\x5f\x24\xc9\xd2\x84\xab\xc9\x90\x33\xa2\xe4\x2d\x67\xbe\xf6\x8a\x17\x6b\xd0\xaa\xe2\xf8\x23\xa3\x56\x95\xaf\x2a\x02\x77\xe2\x66\x5b\xcc\xff\xcc\x06\x9c\x35\x80\xe7\x94\xe1\x65\x69\xef\x2e\xc3\x51\x6b\x07\xba\x18\x54\xb9\x80\x30\x5c\x86\x38\x2c\xcc\xb7\xbc\xea\x70\xf9\xbc\x7b\xec\x3f\x98\xee\x08\x79\x95\x42\x1f\xbe\x80\xc8\xfb\xe6\x13\x1d\x37\xd8\x3f\x59\x0c\xd7\xa3\xc1\xf8\xf3\x49\xc9\x7d\x09\x6b\x50\xe5\x25\xe0\xb1\xf2\xff\x06\xbe\xf4\xad\x59\x3c\x88\x63\xd8\x73\x27\x71\x24\x10\x67\x58\xc2\x89\x9f\x01\x00\x00\xff\xff\x4c\x95\x58\x89\x84\x06\x00\x00")

func clarchTemplatesRepositoryRepositoryTplBytes() ([]byte, error) {
	return bindataRead(
		_clarchTemplatesRepositoryRepositoryTpl,
		"clarch/templates/repository/repository.tpl",
	)
}

func clarchTemplatesRepositoryRepositoryTpl() (*asset, error) {
	bytes, err := clarchTemplatesRepositoryRepositoryTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "clarch/templates/repository/repository.tpl", size: 1668, mode: os.FileMode(420), modTime: time.Unix(1523429493, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clarchTemplatesUsecaseUsecaseTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x94\x4d\x6f\x1a\x31\x10\x86\xcf\xf6\xaf\x98\xee\xa1\xb2\x2b\xb4\xb9\x47\xe2\xd0\x0f\x55\xea\x05\x55\xa9\x38\x55\x55\xb5\xdd\x1d\x88\x15\x63\xaf\xc6\x76\x53\xb4\xda\xff\x5e\xd9\xfb\x01\x24\x10\xd8\xa0\x70\x63\x98\x79\xe7\x9d\x67\xc6\xd4\x45\xf9\x50\xac\x11\x82\xc3\xb2\x70\xc8\xb9\xda\xd4\x96\x3c\x08\xce\x32\x24\xb2\xe4\x32\xce\xb2\xb5\xf2\xf7\xe1\x4f\x5e\xda\xcd\x4d\xd3\xe4\x9f\x03\x11\x1a\xbf\x74\x48\x6d\xbb\x17\xb8\xc3\xda\x76\x81\xef\x0f\xeb\xb6\x7d\x75\xe1\x0d\x61\x6d\x9d\xf2\x96\xb6\x19\x97\x9c\xfb\x6d\x8d\x10\xb3\x8b\x0d\xea\x94\xb1\xec\xdc\x82\x32\x1e\x69\x55\x94\x08\x0d\x67\x5f\x95\xa9\x3e\x6a\x2d\x24\x88\x9f\xbf\x3e\x0c\x62\xf9\x41\xe1\x0c\xd2\x50\xb2\xcb\xfe\xb4\xfd\x56\x09\x55\x45\x19\x09\xe2\x6c\xc9\x0f\x6f\x09\x4f\xa5\x5d\x22\xb0\xac\xab\xc2\xa3\x50\xc6\xcf\xe0\x0a\x99\x2f\xa8\x31\xca\xf4\xc6\x53\x98\xb7\x3b\x4e\x07\x88\x9c\xa7\x50\xfa\xc8\x67\xf8\x25\xd2\x86\x1d\xe2\xc3\x2e\x77\x63\x3c\x2a\xae\x82\x29\x61\x81\x8f\xc7\xe0\x0b\xba\x40\x44\x1e\xdf\x5b\xc3\x19\xa1\x0f\x64\xe0\xfd\x13\xbf\x0d\x67\x07\x46\x6f\x81\x66\x9c\xb5\xa3\x19\xe1\xef\x95\xdb\xc1\xeb\xab\x24\x4c\x59\x7e\xd7\xde\xa5\xaf\x70\x3b\x87\x28\x99\xef\x37\xcd\x47\x35\xce\xd4\x2a\xa5\xbd\x9b\x83\x51\x3a\x56\x0e\xce\x8d\xd2\x49\x21\xba\x1b\x62\x49\xd5\x28\x7d\x91\xdd\x49\xd7\x77\x99\xe7\x5e\xf2\xcd\x6c\x77\x2f\xe0\x6f\xa1\x03\xbe\xfe\x80\xf7\xb6\xff\x7c\x8a\xbd\x0e\xf2\xac\x9d\xe1\x3d\x25\x86\x33\xb8\xde\x17\xf4\x1f\xfc\xa7\x9c\xc7\x6a\x12\xed\xae\x04\xe6\x27\x89\x5b\x72\xf9\x02\x1f\x45\xb6\xb0\x1e\x56\x36\x98\x2a\x93\x69\x0b\x93\x76\xf5\xdc\xca\x48\xa1\x27\x70\x9e\xdb\xb1\x3f\x90\xd8\x77\x1c\xfb\xf7\xd5\x43\xbf\x30\x2f\x67\x27\xa9\x8e\xc6\x5e\xba\xe0\x27\x40\xba\xbb\xfd\x1f\x00\x00\xff\xff\xc7\x80\x1e\x9b\xc8\x06\x00\x00")

func clarchTemplatesUsecaseUsecaseTplBytes() ([]byte, error) {
	return bindataRead(
		_clarchTemplatesUsecaseUsecaseTpl,
		"clarch/templates/usecase/usecase.tpl",
	)
}

func clarchTemplatesUsecaseUsecaseTpl() (*asset, error) {
	bytes, err := clarchTemplatesUsecaseUsecaseTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "clarch/templates/usecase/usecase.tpl", size: 1736, mode: os.FileMode(420), modTime: time.Unix(1523361400, 0)}
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
	"clarch/templates/entity.tpl": clarchTemplatesEntityTpl,
	"clarch/templates/handler/http/hander_test.tpl": clarchTemplatesHandlerHttpHander_testTpl,
	"clarch/templates/handler/http/http.tpl": clarchTemplatesHandlerHttpHttpTpl,
	"clarch/templates/project/config.tpl": clarchTemplatesProjectConfigTpl,
	"clarch/templates/project/handler.tpl": clarchTemplatesProjectHandlerTpl,
	"clarch/templates/project/middleware/middleware.tpl": clarchTemplatesProjectMiddlewareMiddlewareTpl,
	"clarch/templates/repository/repository.tpl": clarchTemplatesRepositoryRepositoryTpl,
	"clarch/templates/usecase/usecase.tpl": clarchTemplatesUsecaseUsecaseTpl,
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
	"clarch": &bintree{nil, map[string]*bintree{
		"templates": &bintree{nil, map[string]*bintree{
			"entity.tpl": &bintree{clarchTemplatesEntityTpl, map[string]*bintree{}},
			"handler": &bintree{nil, map[string]*bintree{
				"http": &bintree{nil, map[string]*bintree{
					"hander_test.tpl": &bintree{clarchTemplatesHandlerHttpHander_testTpl, map[string]*bintree{}},
					"http.tpl": &bintree{clarchTemplatesHandlerHttpHttpTpl, map[string]*bintree{}},
				}},
			}},
			"project": &bintree{nil, map[string]*bintree{
				"config.tpl": &bintree{clarchTemplatesProjectConfigTpl, map[string]*bintree{}},
				"handler.tpl": &bintree{clarchTemplatesProjectHandlerTpl, map[string]*bintree{}},
				"middleware": &bintree{nil, map[string]*bintree{
					"middleware.tpl": &bintree{clarchTemplatesProjectMiddlewareMiddlewareTpl, map[string]*bintree{}},
				}},
			}},
			"repository": &bintree{nil, map[string]*bintree{
				"repository.tpl": &bintree{clarchTemplatesRepositoryRepositoryTpl, map[string]*bintree{}},
			}},
			"usecase": &bintree{nil, map[string]*bintree{
				"usecase.tpl": &bintree{clarchTemplatesUsecaseUsecaseTpl, map[string]*bintree{}},
			}},
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

