// Code generated by go-bindata.
// sources:
// wski18n/resources/de_DE.all.json
// wski18n/resources/en_US.all.json
// wski18n/resources/es_ES.all.json
// wski18n/resources/fr_FR.all.json
// wski18n/resources/it_IT.all.json
// wski18n/resources/ja_JA.all.json
// wski18n/resources/ko_KR.all.json
// wski18n/resources/pt_BR.all.json
// wski18n/resources/zh_Hans.all.json
// wski18n/resources/zh_Hant.all.json
// DO NOT EDIT!

package wski18n

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

var _wski18nResourcesDe_deAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesDe_deAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesDe_deAllJson,
		"wski18n/resources/de_DE.all.json",
	)
}

func wski18nResourcesDe_deAllJson() (*asset, error) {
	bytes, err := wski18nResourcesDe_deAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/de_DE.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1515697090, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesEn_usAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x5a\x7b\x8f\x1b\xb7\x11\xff\xdf\x9f\x62\x60\x14\x70\x02\x9c\x65\x27\x45\x81\xc2\xc0\xa1\x70\x6b\x37\xb9\x26\xf6\x19\xf7\x48\x10\x38\x87\x35\xb5\x1c\x49\x8c\xb8\xe4\x82\xe4\x4a\x56\x0e\xea\x67\x2f\x86\x5c\xae\x56\xba\xe3\x2e\x25\x27\x68\xfe\x89\x7c\x1c\xce\x6f\x1e\xe4\xbc\xb8\x1f\x9f\x00\xdc\x3f\x01\x00\x78\x2a\xf8\xd3\x57\xf0\xb4\xb2\xf3\xa2\x36\x38\x13\x9f\x0b\x34\x46\x9b\xa7\x67\x61\xd5\x19\xa6\xac\x64\x4e\x68\x45\x64\x6f\xfd\xda\x13\x80\xed\xd9\x00\x07\xa1\x66\x3a\xc1\xe0\x82\x96\xc6\xf6\xdb\xa6\x2c\xd1\xda\x04\x8b\xeb\x76\x75\x8c\xcb\x9a\x19\x25\xd4\x3c\xc1\xe5\xe7\x76\x35\xc9\xa5\xac\x78\xc1\xd1\x96\x85\xd4\x6a\x5e\xd4\xcd\x54\x0a\xbb\x48\x30\xfb\x10\x56\x81\x41\xcd\xca\x25\x9b\x23\x38\x0d\x6e\x81\x60\x70\x2e\xac\x33\x1b\xb0\xe8\x40\x28\xf8\xef\x8b\xc9\xda\x2e\x6b\xa3\x6b\x3b\xc9\x85\x36\x58\x6b\xe3\x12\xc8\x57\x7e\xd1\x82\x56\xc0\xb1\x96\x7a\x83\x1c\x50\x39\xe1\x04\x5a\xf8\x4a\x4c\x70\x72\x06\x1f\x82\x4c\xf6\x0c\x5e\x97\xb4\xcf\x9e\xc1\x8d\x11\xf3\x39\x1a\x7b\x06\x57\x8d\xa4\x15\x74\xe5\xe4\x6b\x60\x16\xd6\x28\x25\xfd\xdf\x60\x89\xca\xf9\x1d\x2b\x8f\x66\x49\x7e\xd2\xc9\xd6\x58\x8a\x99\x40\x0e\x8a\x55\x68\x6b\x56\x62\xbe\x2e\x5a\xa7\x34\x79\x0d\x4e\x6b\x49\x86\x0b\x8a\x9c\x41\xa3\xc2\x2f\x60\x8a\x83\xdd\xa8\x12\x74\x8d\x6a\xbd\x10\x76\x19\xed\x6c\xa1\xb1\x42\xcd\x81\x41\xc5\x94\x98\xa1\x75\x9e\x58\xd7\xc4\x95\xc9\x96\x55\x45\x9a\xcc\x84\xec\xc8\x7f\x79\xfd\xee\xc7\x1c\x99\xed\x42\x1b\x77\x8a\xef\x59\xe7\xf9\x7c\x98\x41\x3f\x7f\x30\x7a\x25\x38\x5a\x60\x60\x9b\xaa\x62\x66\x03\x81\x1e\xf4\x0c\xd6\x0b\xe6\x9e\x59\x98\x22\xf6\x4e\xc1\x97\x79\xab\x15\x69\xd4\x5d\x74\xae\x9d\x86\x05\xca\xba\x85\x86\x8d\x6e\x4c\x96\xa7\xc8\x23\xf9\xb2\x30\xce\x53\xa2\x70\x0e\x4c\x01\xf3\x87\xfb\x0c\x66\x88\xfc\x0c\x5c\x38\xe1\xa0\x0d\x98\x46\x76\x17\x32\x82\x1f\x03\x5b\xf8\x0b\xb5\x19\x42\x7f\xc1\x14\xdc\xdf\x4f\x96\xb8\xd9\x6e\x0f\xa1\xbc\x9e\xd9\x78\x2b\x34\x96\x38\xa7\x8e\x81\x50\xce\x73\x6f\xe9\x40\x35\xd5\x94\xd4\x9c\xc1\xda\x2e\x83\x07\x86\xb1\x66\x92\xcd\x0b\x56\x8b\x62\xa1\x6d\xca\xb9\xc1\x73\xaf\x3f\x5c\xc0\xa7\xef\x2f\xaf\x6f\x3e\x65\x72\x1c\x96\xbd\xc7\xf4\xa7\xb7\x57\xd7\x17\x97\xef\xb3\xf8\x36\x6e\x51\x2c\x31\x65\x7d\x5a\xd6\x46\xfc\xee\xff\x00\x9f\x7e\x78\xfb\x4b\x0e\xd3\x12\x8d\x2b\xbc\x5b\x1e\xe7\x5a\x33\xb7\x20\x93\x92\xa1\x27\x44\x9c\xe1\xc3\xc0\x58\xab\x99\x48\xa5\x9c\xb0\xe8\x59\xc1\x57\x1c\x67\xac\x91\x0e\x84\x85\xbf\x7c\x7f\xf9\xee\xed\x2e\x31\x7c\x9d\x63\x15\x29\xf5\xba\x68\x79\xa4\x12\xa5\x27\x82\x8e\x68\x9c\xeb\x2e\x5a\x0e\xd9\xa5\x8b\xd0\x5d\x58\xcd\x60\x2d\x94\x43\x43\x37\x74\x95\xb2\x79\x90\xb6\x47\x07\xb5\xd1\x55\x9d\x25\xf8\x12\x37\xd9\xee\x5c\xe2\x26\x57\xe8\x60\xe5\x8a\x29\x36\xc7\x54\xf0\x09\x62\xd7\x46\xff\x86\xa5\xdb\xa5\x5e\xa7\x61\x4a\x21\xc0\x2c\x91\x43\xe4\x30\x8e\xd8\x85\xa7\x61\xfb\x1f\x11\x5a\x3c\xdb\x2e\xf4\x27\xf8\xee\xd6\xc7\x79\xb5\xaa\x8e\x48\x68\xd1\xac\xd0\x48\xb4\x36\xda\x26\x83\xb5\x75\x46\x24\x39\x07\x43\x37\x16\x0d\x1d\x69\xa1\x90\x83\x69\x94\x13\x55\x17\x0b\x33\x10\x9c\x9e\xcf\x25\x16\x94\xae\x12\x30\x37\x9e\x02\xbe\xa7\x84\x56\xa1\xb5\x6c\x9e\x7f\x52\x56\x68\xa6\xda\xa6\x8c\xdc\xae\x82\x6e\x5c\xdd\x0c\x99\xc3\x87\x89\xa2\x12\x96\x12\xa6\x0f\x80\xe9\xf8\x77\xb3\x40\x20\x0a\x3a\x78\x65\x08\x82\x74\xc0\x85\x05\xa5\x1d\x04\x56\x8d\x41\x3e\xf9\x75\xc8\x3c\x07\x88\xb5\x18\xc8\x0d\x84\x48\x41\x9c\x48\xbe\x0c\x67\xec\x54\x12\x52\x47\x73\x1a\x54\xab\xca\x50\x47\x72\xa8\xcf\xc7\xfb\xfb\x09\xfd\xde\x6e\xef\xce\x60\x66\x74\x45\x89\xdd\xea\xc6\x94\xb8\xdd\x66\x61\x06\x87\x8d\x61\x12\x59\xf4\x95\x45\x77\x1a\x56\x67\x9e\x31\xb4\x3d\x3b\x92\x8a\xdd\x1f\x4e\xd7\xb3\x16\xf3\x75\xc1\x7c\x33\x56\x38\xbd\x44\x35\xaa\x32\xed\x80\xb0\x03\xfc\x8e\xd3\x94\x6f\x54\xc5\x8c\x5d\x30\x59\x48\x5d\x32\x99\x40\xbc\x8d\x54\x70\x59\xa3\xfa\xd9\x57\x1f\x6d\xc4\xb0\x01\xcf\xef\x86\x15\x93\x0d\xda\x4c\x40\x85\x6e\xad\xcd\xf2\x64\x48\x9f\xdf\x14\x3a\x60\x8e\xd4\x6d\x8c\x1c\xd1\x75\x97\x6a\x8b\x92\xa9\x12\xa5\x4c\xa6\xa2\xcb\x1f\x26\xf0\xaf\x40\x43\x95\xf6\x6e\x67\x2e\xc0\x8c\x89\x34\xf7\x37\xbb\x9c\xcf\x05\x6f\xef\x62\x55\x4b\x74\x08\xb6\x21\x97\xce\x1a\x29\x37\x13\xb8\x6a\x14\x7c\xea\x8a\xd1\xae\x8b\xfb\x44\x69\xc1\x60\xa5\x29\xb3\x33\xe3\x04\x93\x72\xb3\x6b\x57\x98\xb5\xe8\x86\xbd\xd0\x93\x34\xf4\x3e\x85\x75\xcc\x35\xa9\xf2\xe7\xf9\xf3\xe7\xcf\xcf\xcf\xcf\xcf\x7b\xbe\xe8\xe9\x70\xed\xb7\x02\x11\x10\x61\x16\xaa\x9f\x4b\x20\xcf\x31\x51\x34\x0d\x87\x76\x98\x11\x8c\x33\x7c\xc8\x4e\xf7\x75\x7f\x6f\x3e\xc8\xa0\xbf\x6f\x7b\x94\xc3\x1e\xcf\xc6\x1b\xb3\xdf\x1e\xe4\x09\x16\x8c\x65\x51\xe1\x5b\xcd\xf1\x72\xf6\xd6\x77\xa4\x14\x0d\xa9\x6c\xd9\x6e\xef\x60\xa6\x4d\xee\xbd\x39\x00\xeb\x2b\x7a\x14\x5c\xb6\xeb\x42\x17\x5a\xc4\x1b\x33\x32\x25\xeb\xba\xd1\x18\xec\x09\x6f\xc1\xda\x29\x41\xdf\xa4\xdd\x1d\xcc\x47\x4f\x8f\xd5\xde\xc4\x75\x78\x54\x80\xc9\x64\xa0\xe1\x6f\x21\xa2\x41\xfe\x48\x15\x77\x3c\x73\x94\x8c\xd4\x69\x35\x6f\x77\x14\x27\x28\xca\xb1\x46\xc5\x51\x95\xc7\xd8\x73\xb7\xa9\x0f\x74\x1c\xce\xee\x16\x26\x8d\xfa\xe6\x51\x98\x2f\x39\x39\x8f\x4b\x41\xb1\xa7\x31\xa9\xd2\xaf\x17\x49\xf5\x2c\xa1\xfa\xff\x31\x0d\x45\x7d\x8e\x3b\x28\x5f\xe6\xc1\x87\x91\xf4\x8f\xf1\x61\xe6\xd5\x48\x49\x32\xec\xc7\xbd\x88\x7e\xa2\x27\x47\x82\x30\x75\xd6\xa7\xa6\x35\x2f\x51\x48\x32\x5d\xe7\x3e\x24\x0b\xf0\xc6\x90\x27\x5b\xd8\x7e\xaa\xf8\xf3\xce\x5b\xd4\x71\xa6\x1b\xc5\x8b\x56\xde\xe1\x99\xe4\x1b\x22\x4a\xc6\xa6\xf5\x42\x94\x0b\x58\xfb\x21\x3f\xc9\xc5\x43\x69\xea\x16\x08\x65\x63\x0c\x19\x26\x2a\x18\xe7\x1a\x3e\x69\x85\xdf\xc4\x81\x59\xaf\x0b\xd9\x2f\x3b\x7d\x85\x69\xce\x48\x73\xf9\xab\xfa\x20\x91\x59\x3f\xfb\x59\x09\x8e\x5e\x28\xa2\x27\xd9\x7d\xba\xec\x8a\xb9\x57\x30\x8e\x35\xd8\x3a\x3f\xc0\x62\xea\xb0\x95\xf6\x3d\x4a\x06\x50\x3b\x77\x4d\x38\x43\xc3\x46\x37\x60\xd0\x7b\x7e\xcd\x94\xdb\x8d\xd0\xc0\x2d\x84\xfd\x07\x7c\xb5\x79\xf1\xfe\xeb\x0c\x9c\xb1\x8e\xf9\xa1\x4a\xbd\xc6\xef\x63\x1c\x39\xfa\x86\x87\x3a\xaf\x79\x83\xd6\xdd\x65\xe0\x46\x27\x1f\xa5\x61\xf7\x7c\x93\xa1\xe3\xdb\xab\xab\xcb\xab\xeb\x04\xfb\xf3\xc3\xff\x20\x90\xc3\x83\x85\xf3\xf3\x81\xdc\x6e\xcc\x7e\x10\x5b\x2a\xbd\x56\x85\xdb\xd4\x03\x49\x28\x06\x2b\xa2\x22\x8b\xb5\xbb\x26\xb0\x9b\xb6\x83\x56\x72\x03\xb6\xa9\xc3\x43\xdc\x0b\x3f\xe6\x9e\xd8\x8d\x75\x58\xc1\x54\x28\x2e\xd4\xdc\x82\x36\x30\x17\x6e\xd1\x4c\x27\xa5\xae\xba\x37\x91\xe1\x62\x24\x08\x1c\xa3\x1a\x39\xb2\x50\xda\x85\x38\x30\xd0\x65\x3f\x7a\xe7\x85\x3a\x1c\xda\xfa\x20\xe7\x79\xd1\xe2\xfe\xa3\xcc\x98\x58\x6d\x9d\x54\x1a\x64\x2e\x65\x3d\xff\x54\x0c\x9e\x64\x2f\x12\xad\x85\x5b\x80\x7f\x63\x8e\xf3\xb5\x57\xb4\x88\xc6\x6c\xb7\xfe\xf1\x2e\xac\x95\x9a\x87\x05\xfa\x31\xd2\x23\xf7\x44\x0a\xe1\x71\x50\x24\xfe\x20\x38\xfe\x49\x22\xcd\x10\x79\x21\xd4\x4a\x2f\x53\x02\xfd\xdb\x67\x2a\xba\x2c\x81\xcc\x87\x3b\xda\x06\xeb\x85\x7f\xcb\x6b\x25\x8d\x8f\x59\x61\xe9\xcf\x91\x76\x89\x9b\x6e\x32\x57\x31\xc5\x99\xd3\x66\x68\xea\xd8\xd1\xf8\x21\xd6\xc7\x68\xcc\x3b\xba\x26\x2d\x9f\x51\xcc\xae\x87\x1a\x3b\xd7\xef\xfa\xc7\xb3\x77\x72\x99\x03\x3f\x71\xee\x35\x55\xa3\xa0\xfe\x22\x55\xc2\x56\xcc\x95\xa9\x47\xdd\xbd\x7b\x44\x1b\xb8\x87\xe0\x03\xd7\xc9\xaf\xc7\xc6\x8e\x6b\x0c\xe3\x4a\x0f\xe2\xdd\xea\x33\x18\x11\x55\x3d\x26\x7b\xd7\x2e\xac\x46\x35\x86\x95\x68\xa7\x4a\x74\xbc\x98\x14\x29\xb3\x5d\x84\x55\x8a\x3e\xad\x4b\xba\x99\x39\x61\xb5\xbf\x49\x96\xdd\x33\xf1\x9e\x54\xda\x78\xd9\xc3\xd3\xaa\xdf\x13\x7e\xe6\xd8\x39\x8a\x38\x62\xea\xab\x63\x04\x3a\xb0\xab\xbf\x0a\x41\xa2\x67\x16\xc2\xec\x30\x98\x12\x3f\x3b\x54\x36\x0a\x8d\x9f\x5d\x6c\xb7\xbf\x44\x15\x5b\xcc\x31\x55\xb9\xec\xae\xf2\x1c\xc3\x0b\x6d\x9b\x12\x76\xef\x14\xed\x08\x70\x37\x8e\xb2\x68\x56\xa2\xec\x5d\xdf\x51\x41\x1a\x23\x8f\x77\x79\x98\x33\x52\x0a\xdb\x6e\xe1\xf6\xea\x47\xaf\xbc\x9f\x3c\xfa\x33\x48\xff\x22\x9b\x05\x82\xbb\xbc\x0c\x40\x82\x54\x4c\xce\xb4\xa9\x92\xa5\xf6\xbb\xb8\x3e\x24\xc1\x04\x6e\xcc\x06\xd8\x9c\x09\x35\x99\x8c\xc2\xfe\x66\xb5\xea\xa2\x54\x59\xf1\x81\x97\xe1\xff\x5c\x5f\xbe\x07\xa1\xea\xc6\x01\x67\x8e\xc1\xbb\xd6\x1a\xcf\xca\x8a\x3f\xa3\x98\x35\x8c\xc4\x6a\xd1\x01\xad\x71\x5a\x84\xc3\x92\xfa\x06\xe0\x91\x43\x15\x1f\x2c\x18\xac\x71\xda\x7d\x9c\xf0\xfa\xc3\x45\x20\xab\x05\xd1\x94\x4c\x85\xfa\x61\x8a\x21\x55\x22\x6f\xbf\x93\xd8\x6d\x9a\x40\x5b\xd4\x35\x35\x67\xee\xe0\xbb\x02\x3a\x70\xa5\x56\x2b\x34\xee\x00\xde\xe9\x3e\x8f\x31\xc3\xf6\xd5\x3d\x49\xd5\x78\xd8\xfc\x11\xdf\x13\x31\x4b\xe9\x29\xb3\xc8\x41\xab\x7e\xb8\x79\xc8\x6a\xd4\x14\x42\x95\xb2\xe1\x78\x20\x1e\xb3\x7b\x5e\x48\x1a\xe3\xe7\xd7\x57\xef\x2f\xde\x7f\x97\x5f\x87\xc6\x0d\xc7\x55\xa2\x6b\x66\x54\x51\xea\x8a\x32\x68\x61\xd0\x25\xd3\xec\x15\xad\xc5\x99\x61\x59\x71\xaf\xcb\xcc\xa1\x09\x29\xfe\x55\x88\x6d\x14\x38\xee\x86\xfc\xdb\xe2\xf9\x47\x9b\xa3\x83\x47\xff\xa3\x85\xfe\xe8\x92\xa3\xc3\xd2\x8d\x0c\x0d\x3c\x32\x15\x16\x1c\x6b\x83\x25\x79\xba\x30\x58\x4b\x56\x26\x83\x06\x65\x5e\xc2\xd1\x92\xb7\xf5\x84\x7f\x23\x0b\x07\x63\x2f\x52\x05\x99\xd6\x42\x4a\xb0\x5a\x2b\x3a\x4d\x3b\x98\x33\xa8\xdb\x93\x62\x43\x55\xe5\x8b\x61\x5c\xef\xf1\xb4\x0e\x59\xa6\x02\xad\x39\x4e\xa9\x87\xec\x42\x37\x92\x93\x78\x16\xdd\x04\xc2\x14\x78\xbf\x17\x23\x6a\xff\x2b\xcc\x83\xb2\x24\xf2\xf4\x23\xfe\x24\xb9\x02\x02\xa5\xbe\x87\x75\x1a\xc5\x27\xbf\xff\x18\x48\xaa\xd5\x2c\x5b\x0d\x7a\x70\x0c\xd4\xef\x8f\x5e\x8d\x43\x87\xf8\x49\x53\xff\x5b\xa6\x71\xc1\xa4\xa8\x84\x2b\xc4\x5c\x69\x93\x14\x29\x9e\xeb\x36\xb0\xf8\x2d\x5e\x2a\xff\xeb\xb0\x16\x13\x16\x5a\x76\xb9\xe8\xe5\x82\xa9\x39\xb2\x69\xf2\x0b\x94\x1f\x3b\xc4\xae\xf8\xb3\x51\x6f\xb9\x09\xf3\xa6\x8e\xc7\x04\x2e\x08\x9e\x0a\xe8\x8c\xb3\xe0\x25\xb0\x85\xd4\xf3\xc2\x8a\xdf\x53\x02\x48\x3d\xbf\x16\xbf\x23\xd9\x36\x6c\xd8\xd3\x78\x77\x44\x99\xf2\x4f\x94\xd4\x6c\x4c\xd1\xad\x11\x15\xbc\xf4\x4d\xc5\x37\x2f\xb3\x45\xa9\xb0\xd2\x66\x33\x24\x4d\xa0\x38\x55\xa0\x6f\xbe\xfd\xbb\x17\xe9\x6f\xdf\x7c\x9b\x2d\x13\xd5\x5f\xba\x49\x15\x6f\xed\xea\x49\xc2\xbc\x0c\xf6\xf9\xeb\x4b\xfa\x6f\x5c\x1e\x3f\x1e\x28\x6a\xa3\x6b\x34\x4e\x60\x6a\x84\x1b\xc3\x60\x2f\x5e\x85\x81\x9d\x33\x02\xbb\x91\x5d\x98\x35\xec\x98\xc5\xd1\xde\xe3\x31\x31\x86\x44\xae\xfd\x81\xa3\xc8\x28\x1c\xe8\xc6\x59\xc1\xbd\x23\x6e\x0c\x5b\x09\x0b\xd3\x46\x48\x3e\x3c\x9b\xf0\xaa\x84\x70\x60\xe8\xd8\x66\x85\x82\xee\xf4\xef\x05\x04\x75\x10\xd5\x5b\x6b\xfb\x89\xcb\xfd\xfd\xa4\xfd\x6b\x34\x37\x75\x48\x42\xb5\x8d\x2e\xfd\x83\x95\x23\x65\xb3\x17\x35\xf6\x22\xe1\x92\xa5\xc2\x44\x6c\x45\x5a\x2a\x2a\x28\x0e\xba\x92\x47\xca\x94\x64\xe3\x71\x52\xb7\xe1\xa5\x6d\x67\x19\xbe\x51\xc5\xcf\xc2\x26\xbf\xf0\x7b\xd0\xa6\xee\x85\x18\x26\x0d\x32\xbe\x81\xc0\xa2\xab\x9d\x2c\x4a\x2c\x1d\x30\xa5\xdd\x02\x8d\xdf\x36\x2e\x52\x9c\x58\x65\x4d\xa0\x1e\x76\xb5\xb1\x6a\x28\xb5\x72\xcc\x7f\xbb\xa5\xf4\xf8\x14\xec\xcd\xdb\x7f\xde\x7e\x97\x5d\x8c\x79\xea\xe3\x2a\x31\x3e\x0d\xdf\x0e\xae\xd0\x88\x59\xaa\x0a\xfb\xc9\x2f\xb6\x8d\xd4\xc3\x03\xdc\x5a\x77\x38\x52\x13\x50\xf7\x10\x14\x5e\x8c\x47\x1f\xfc\x3c\xd5\xab\x31\xae\xbd\x27\xa8\x41\xbe\xfd\x87\xa8\x3c\xce\x16\x99\x29\x17\xc4\xb7\x9d\xfc\x17\x5c\x18\x2c\x07\xc6\x42\xd7\x71\x47\xf7\x6e\xd0\xed\x78\xf0\xec\xbd\xb3\xe0\xce\x74\x4f\xee\x9e\xfc\x2f\x00\x00\xff\xff\x9e\x0a\x2f\xc8\x99\x32\x00\x00")

func wski18nResourcesEn_usAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesEn_usAllJson,
		"wski18n/resources/en_US.all.json",
	)
}

func wski18nResourcesEn_usAllJson() (*asset, error) {
	bytes, err := wski18nResourcesEn_usAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/en_US.all.json", size: 12953, mode: os.FileMode(420), modTime: time.Unix(1519054009, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesEs_esAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesEs_esAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesEs_esAllJson,
		"wski18n/resources/es_ES.all.json",
	)
}

func wski18nResourcesEs_esAllJson() (*asset, error) {
	bytes, err := wski18nResourcesEs_esAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/es_ES.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1515697090, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesFr_frAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8a\xe6\x52\x50\xa8\xe6\x52\x50\x50\x50\x50\xca\x4c\x51\xb2\x52\x50\x4a\xaa\x2c\x48\x2c\x2e\x56\x48\x4e\x2d\x2a\xc9\x4c\xcb\x4c\x4e\x2c\x49\x55\x48\xce\x48\x4d\xce\xce\xcc\x4b\x57\xd2\x81\x28\x2c\x29\x4a\xcc\x2b\xce\x49\x2c\xc9\xcc\xcf\x03\xe9\x08\xce\xcf\x4d\x55\x40\x12\x53\xc8\xcc\x53\x70\x2b\x4a\xcd\x4b\xce\x50\xe2\x52\x50\xa8\xe5\x8a\xe5\x02\x04\x00\x00\xff\xff\x45\xa4\xe9\x62\x65\x00\x00\x00")

func wski18nResourcesFr_frAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesFr_frAllJson,
		"wski18n/resources/fr_FR.all.json",
	)
}

func wski18nResourcesFr_frAllJson() (*asset, error) {
	bytes, err := wski18nResourcesFr_frAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/fr_FR.all.json", size: 101, mode: os.FileMode(420), modTime: time.Unix(1515697090, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesIt_itAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesIt_itAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesIt_itAllJson,
		"wski18n/resources/it_IT.all.json",
	)
}

func wski18nResourcesIt_itAllJson() (*asset, error) {
	bytes, err := wski18nResourcesIt_itAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/it_IT.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1515697090, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesJa_jaAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesJa_jaAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesJa_jaAllJson,
		"wski18n/resources/ja_JA.all.json",
	)
}

func wski18nResourcesJa_jaAllJson() (*asset, error) {
	bytes, err := wski18nResourcesJa_jaAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/ja_JA.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1515697090, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesKo_krAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesKo_krAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesKo_krAllJson,
		"wski18n/resources/ko_KR.all.json",
	)
}

func wski18nResourcesKo_krAllJson() (*asset, error) {
	bytes, err := wski18nResourcesKo_krAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/ko_KR.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1515697090, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesPt_brAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesPt_brAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesPt_brAllJson,
		"wski18n/resources/pt_BR.all.json",
	)
}

func wski18nResourcesPt_brAllJson() (*asset, error) {
	bytes, err := wski18nResourcesPt_brAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/pt_BR.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1515697090, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesZh_hansAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesZh_hansAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesZh_hansAllJson,
		"wski18n/resources/zh_Hans.all.json",
	)
}

func wski18nResourcesZh_hansAllJson() (*asset, error) {
	bytes, err := wski18nResourcesZh_hansAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/zh_Hans.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1515697090, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesZh_hantAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesZh_hantAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesZh_hantAllJson,
		"wski18n/resources/zh_Hant.all.json",
	)
}

func wski18nResourcesZh_hantAllJson() (*asset, error) {
	bytes, err := wski18nResourcesZh_hantAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/zh_Hant.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1515697090, 0)}
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
	"wski18n/resources/de_DE.all.json": wski18nResourcesDe_deAllJson,
	"wski18n/resources/en_US.all.json": wski18nResourcesEn_usAllJson,
	"wski18n/resources/es_ES.all.json": wski18nResourcesEs_esAllJson,
	"wski18n/resources/fr_FR.all.json": wski18nResourcesFr_frAllJson,
	"wski18n/resources/it_IT.all.json": wski18nResourcesIt_itAllJson,
	"wski18n/resources/ja_JA.all.json": wski18nResourcesJa_jaAllJson,
	"wski18n/resources/ko_KR.all.json": wski18nResourcesKo_krAllJson,
	"wski18n/resources/pt_BR.all.json": wski18nResourcesPt_brAllJson,
	"wski18n/resources/zh_Hans.all.json": wski18nResourcesZh_hansAllJson,
	"wski18n/resources/zh_Hant.all.json": wski18nResourcesZh_hantAllJson,
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
	"wski18n": &bintree{nil, map[string]*bintree{
		"resources": &bintree{nil, map[string]*bintree{
			"de_DE.all.json": &bintree{wski18nResourcesDe_deAllJson, map[string]*bintree{}},
			"en_US.all.json": &bintree{wski18nResourcesEn_usAllJson, map[string]*bintree{}},
			"es_ES.all.json": &bintree{wski18nResourcesEs_esAllJson, map[string]*bintree{}},
			"fr_FR.all.json": &bintree{wski18nResourcesFr_frAllJson, map[string]*bintree{}},
			"it_IT.all.json": &bintree{wski18nResourcesIt_itAllJson, map[string]*bintree{}},
			"ja_JA.all.json": &bintree{wski18nResourcesJa_jaAllJson, map[string]*bintree{}},
			"ko_KR.all.json": &bintree{wski18nResourcesKo_krAllJson, map[string]*bintree{}},
			"pt_BR.all.json": &bintree{wski18nResourcesPt_brAllJson, map[string]*bintree{}},
			"zh_Hans.all.json": &bintree{wski18nResourcesZh_hansAllJson, map[string]*bintree{}},
			"zh_Hant.all.json": &bintree{wski18nResourcesZh_hantAllJson, map[string]*bintree{}},
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

