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

var _wski18nResourcesEn_usAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x59\x51\x6f\xe3\xb8\x11\x7e\xcf\xaf\x18\xec\x4b\x5a\x20\x70\x77\xaf\x28\x50\x2c\x50\x14\x41\x93\xb6\xe9\xdd\x26\x41\xb2\x7b\x8b\xc3\xde\x42\x61\xc4\x91\xcc\x33\x45\x0a\x24\x65\xaf\xcf\xf0\x7f\x2f\x86\x14\x6d\x39\x31\x25\xd9\x9b\xe0\xf2\xa4\x58\xc3\xef\x9b\x19\x0e\x67\x86\xa3\x2f\x27\x00\xab\x13\x00\x80\x37\x82\xbf\x79\x0f\x6f\x2a\x5b\x66\xb5\xc1\x42\x7c\xcb\xd0\x18\x6d\xde\x9c\x85\xb7\xce\x30\x65\x25\x73\x42\x2b\x12\xbb\xf4\xef\x4e\x00\xd6\x67\x3d\x08\x0b\x66\x94\x50\x65\x02\xe3\x73\xfb\x76\x08\xc5\x36\x79\x8e\xd6\x26\x50\xee\xdb\xb7\x43\x28\x42\x15\x3a\x01\x71\x45\xaf\x92\xeb\x7f\xb3\x5a\x65\x95\xb0\x56\xa8\x32\xcb\x2b\x9e\xcd\x70\x99\x00\xfa\xdf\xfd\xcd\x35\x08\x55\x37\x0e\x38\x73\x0c\x3e\x84\x55\x70\x9a\x57\xfc\x14\x68\x5d\x92\x85\x80\x0b\xc9\xca\x4c\xb1\x0a\x6d\xcd\x72\x4c\x70\x6c\xdf\x0f\x63\xb1\xc6\x4d\x7b\xd4\xa5\xd7\xda\x88\xdf\xfd\x0f\xf0\xf0\xe3\xe5\x2f\x0f\x63\x40\x6b\x91\x4d\xb5\x75\x09\xd0\xc5\x54\xd8\x19\x9c\xdf\x5e\xc1\xc3\x7f\x6f\xee\x3f\x8e\x45\x9c\xa3\xb1\x84\x30\x08\xfa\xf3\xe5\xdd\xfd\xd5\xcd\xf5\x18\xdc\x19\x2e\xb3\x42\xc8\x94\x27\x6b\xe6\xa6\xa0\x0b\x70\x53\x84\xc9\x0c\x97\xe0\x65\x87\x61\x73\x34\x6e\x34\x2e\x09\x0f\x00\xd7\x46\x57\xb5\xcb\x38\xd6\x52\xa7\xb6\xea\x42\xc3\x52\x37\x60\x90\x49\xb9\x84\x05\x53\x0e\x9c\x86\xb0\x04\xdc\x54\xd8\x7f\xc2\x9f\x96\x7f\xb9\xfe\xf3\x7b\x18\xe4\x69\xd4\x11\x4c\x71\xd1\x81\x5c\x14\x61\xe9\xf8\xfb\x55\xdd\x4a\x64\x16\xa1\x36\x7a\x2e\x38\x02\x53\x40\x2b\x50\x39\x91\x87\xa0\x74\x7a\x86\x6a\x0c\x51\x2d\x7a\x62\xf2\x19\x11\x6d\x0d\xc9\xd3\x61\x82\x42\x1b\xb8\xa9\x51\x7d\xa6\x20\x1b\xc1\x35\x74\x42\x9f\x9b\x05\x9b\x25\xf0\x85\x63\xc1\x1a\xe9\x60\xce\x64\x83\x20\x2c\x94\x0d\x5a\xf7\xb5\x8f\xb7\x62\x4a\x14\x68\x5d\xa6\xb4\xcb\x0a\xdd\x28\x9e\x60\xfe\xd0\x0a\xfa\x80\x03\xa5\x1d\x78\x69\x60\x0e\x7c\x50\x7e\x59\xad\x26\xf4\xb0\x5e\x7f\x9d\xfc\xaa\xd2\x84\x8d\xcf\x75\x1b\xda\xde\x78\xf9\xe4\x33\x5c\x07\xd9\xfb\x33\x2c\xa9\x50\xb9\x43\x88\x06\x42\x73\x3f\x55\x5c\x34\x48\x66\x1a\xe5\x44\x85\x94\xcb\x2b\xe6\xf2\x69\x82\xe5\x2e\x88\x79\x9e\x76\x09\x51\xd9\x1a\x73\x51\x08\xe4\x20\x14\x44\x8d\x81\x6b\xb4\xde\xd1\x1e\x11\x16\xc2\x4d\x81\xe5\x3e\x74\xad\x6e\x4c\x8e\x61\x2b\xf0\x9b\x43\x45\xf9\xcd\xa3\xe2\x37\x17\x95\x6f\x65\xe9\xd7\xf0\x38\xb4\x35\xd1\x88\x7c\xca\x54\x89\xa9\x40\x88\x36\xb4\x52\x74\x82\x9f\x98\xf3\xc8\x2c\x72\xa0\x13\x36\xc5\x7e\x8d\xbf\x4b\xcd\x46\xd9\xa6\xae\xb5\x71\x83\xaa\x8e\x72\xb7\x08\xce\xde\x60\x7a\xe5\x3a\x16\x8c\x57\x30\x48\x65\x52\x54\xc2\x65\xa2\x54\xda\x24\x35\xbc\x52\x73\x26\x05\x8f\x1c\x7e\x89\x67\xf2\x4f\xa4\xec\x13\x15\x5b\xb8\x5e\xfe\x5c\xab\x42\x94\x9b\xbe\xa2\x3f\x51\x7e\x24\x0b\x77\x13\x23\xd5\xab\xd6\x1b\x01\xaa\x39\x94\xb1\x37\x63\x12\x23\x95\x5b\x12\xf9\x3e\x9e\xa1\x6c\x49\x4c\xdb\xf4\x78\x14\x55\x6b\x4a\x5f\x8b\xf7\xd4\x9e\xd5\x6a\x42\x8f\xeb\xf5\x19\x14\x46\x57\xf4\x7f\x88\xfe\xf5\x7a\x14\x63\xd8\xae\x21\x46\x12\x8b\x3b\x65\xd1\x1d\xc7\xb5\x71\xce\x10\xdb\x8e\x17\x57\xab\xc9\xe6\xff\x83\xad\x44\x63\xb2\x12\x5d\x3c\xc5\xa9\xd6\xfb\xdf\x4c\xc8\x90\x5c\x4a\x74\xfe\x18\x6e\x0f\x66\x5c\x1a\x88\x37\xe5\x15\x2c\x9a\xb9\xc8\xf1\x3d\xe9\x82\xc6\x0c\x28\xd2\xa8\x8a\x19\x3b\x65\x52\x66\x52\xe7\x4c\xa6\x0a\x43\x14\xeb\x10\x91\xb3\x02\xb9\x5f\x19\xea\xad\x1d\xcb\xa6\xd0\x2d\xb4\x99\x1d\xc5\x27\x94\x43\xa3\xb0\xbf\x18\x6d\x6b\x56\xb8\xdf\x20\x4f\xe6\x9f\x8b\x8d\x28\xe4\xba\xaa\x25\x92\x7f\xdb\x4b\x51\xd1\x48\xb9\x1c\x4b\x54\xf8\xfd\x1a\x66\xe1\x82\xb7\xa7\x30\xb0\x11\xd9\x86\x0b\xee\x1a\x05\x0f\x0b\x3b\x6b\x1b\xc2\x58\x7e\x1f\x28\x0e\x0c\x56\x7a\x8e\x50\x33\xe3\x84\xef\x1f\xc3\x3b\xe4\xc0\xac\x45\xd7\xef\xfe\x8e\xa6\x39\x53\x39\xca\xb4\xb2\x37\x3f\x4e\xe0\x5f\x41\x86\x5a\x82\xb1\xdd\x86\x3a\xc0\xeb\x9f\x3a\xc2\xc7\xf8\x7d\x87\xac\xd7\xf3\x3b\x4c\xbd\xbe\x1f\xcd\x77\xa0\xff\x46\xb7\x50\x3b\x24\x15\x53\xac\x44\x7e\x80\x71\xba\x00\x8e\xc1\x8f\x54\xca\x9c\x40\xdb\x6b\x30\xf0\xc6\x90\x7e\x2d\x53\x77\x9f\x5f\x2f\x0c\x17\xcc\xa8\xcc\x5f\x38\xa9\xe1\xaf\xd1\x90\x9a\x3d\x69\x97\x3a\x81\x19\x2e\xa9\x0f\xa0\x54\xbf\x60\x16\x0c\x3a\x23\x70\x4e\xfd\x09\x25\x04\x0f\x36\xd9\x82\xd1\x0f\xbe\x59\x94\x12\xac\xd6\x0a\x1e\x91\x34\x34\x98\x33\x72\x4d\x1d\x6e\x0f\x5c\x7b\xbf\x34\x16\x41\x38\xd0\x8d\xb3\x74\x97\xd0\x05\x7c\x34\x6c\x2e\x2c\x3c\x36\x42\xf2\x11\xa6\x50\x9d\xda\xa2\x67\x06\x6b\xc9\xf2\xe4\x7e\x45\x8b\xb4\xe4\x1d\xa3\x44\xe8\x13\x57\xab\x09\x35\x87\x6e\x59\xe3\x7a\x1d\xfa\xc4\x84\x11\x67\xd1\x0a\x52\xdf\xb5\x98\x0a\x17\x3b\x98\xd6\x21\xdb\x2d\xf0\x4f\x8b\x50\x6c\x22\x2a\xa6\x38\x73\xda\x2c\x7b\xa6\x19\xa4\xf9\x46\xce\x33\x74\x76\x46\x58\x68\xb1\x26\xfd\x65\x2f\x5e\x0e\x7c\xd9\xa5\x00\xf8\x0d\xf3\xbe\x26\x29\x52\xf8\xea\xcb\x3d\x25\xa7\xc7\xb6\x2d\xec\x04\xbf\x77\x97\x7f\x1f\x6f\x2f\x4f\x2e\x0f\xae\xad\xe1\x5e\xa8\xea\x80\x54\x3b\xd7\x3a\xff\x36\xde\xe2\x86\xb2\x29\x39\x8f\x63\x8d\x8a\xa3\xca\x93\x57\xfe\x28\x0a\x5b\xd1\xb0\x61\x41\x87\xc9\xa4\x87\x27\x1e\xbc\x51\x4c\x9f\xb6\xc2\xc7\x70\x6d\x97\x3c\xab\x9f\xc9\xf9\xe0\xc5\x5e\x1a\x98\x32\x0b\x8f\x88\x6a\x27\xa1\x6f\xf2\xc4\x50\x9d\xda\xa3\x05\x65\xc1\xc6\xa4\xba\xdc\x8b\x27\x49\x70\xaf\x4e\x7f\x5c\xdd\x8d\xf6\x3c\xaf\x90\x2f\xe3\xd7\x88\x3b\xde\xb3\xcf\xca\x67\xda\xb7\xcf\x4b\xcc\xe1\xde\xed\xd3\x6a\x53\xe7\x74\xa3\x78\xd6\x16\xb0\xcc\x17\xb0\xf4\x89\x92\xe8\x28\xc8\x37\xe9\xa1\xab\x49\x9b\xfe\x7d\xa1\xa0\x7d\x6b\xcb\x04\x9d\xff\xbc\x31\x86\xcc\x88\x15\xaf\x4d\x40\x61\xe8\x11\x9e\x09\x81\x59\xbf\xd7\x64\xed\xe8\xda\x4d\xd9\x2d\x37\xc8\x1c\xf6\xeb\xee\x47\xfb\xe0\x25\x77\x2c\xf0\xb3\x0d\xff\x4d\x00\x2a\xb4\x96\x95\x9d\x26\x1e\x98\xe2\xed\xbb\x5c\xf3\xf0\x82\x1e\x46\xdc\x33\x82\x3f\xc7\xa8\xc4\x9f\x39\xf5\x35\x54\xf2\x7a\x6c\xb3\xe7\x60\xca\xdc\xbb\xc3\xbd\x59\xac\xa5\xe8\x24\xce\x11\xd9\xf2\x68\x9a\x78\xf0\x06\x8e\xf3\x5e\xfc\xef\x48\x92\x4f\x8c\x7c\x49\xfe\x91\xc9\x84\x82\xab\x40\xe4\x99\x50\x73\x3d\x4b\x25\x8f\xed\x1d\x36\x88\xf9\x53\x48\xcb\x60\x31\x45\xb5\x8d\x39\x67\x44\x59\xa2\x69\x5f\xbd\x6c\xdc\xdd\x5f\xde\x9e\xdf\x9d\x7f\xbc\xb9\x4b\xe8\xf8\x8f\xbd\x7f\x70\xa9\xa8\xcb\x24\x0f\x49\xbd\xf0\x83\x79\x90\x42\x21\x30\x13\x66\xbf\x73\x34\x61\x82\xb6\x7f\xfd\xce\xdf\x7e\xc5\x3e\x9f\xdf\x5d\x5f\x5d\xff\xe7\x3d\x50\xa7\x73\x1a\x46\x07\xa7\xf0\xcb\xf9\x87\x9f\x36\x5d\x61\xeb\x96\xb0\xdd\xd4\x66\x6d\x7b\xc0\x09\xdc\x6e\x7b\xc0\x53\x72\xdc\x69\x6c\xfa\x28\x81\x71\xb4\xb9\x11\x8f\x61\xc2\xd7\x8e\xfb\xc2\x70\xcb\x97\xab\xc4\x37\xc3\x3f\x50\xa3\x7e\x17\x7d\xd8\xb4\x9e\x85\x40\xc9\xe1\x96\xe5\x33\x56\x22\xfc\x1c\x3e\x6c\x41\xd5\x58\x47\x3d\xb2\x0d\x83\x81\x01\xf3\x0e\x44\xeb\x57\xed\xe9\xe2\x38\x45\x65\xf3\x60\x6a\xb7\xf0\x6c\x1b\x8c\xf6\x8b\xdc\x08\x65\xbf\x1b\xff\x38\xcf\xfe\x24\x72\x54\x16\x5f\xc8\xb3\x09\xb4\x71\x9e\x8d\x8b\x5f\xcb\xb3\x47\xe3\xf7\xab\x1f\x61\xfd\x2c\xdb\x3f\x5e\x5d\xac\xd7\x91\x85\x41\x98\x7b\x6b\x85\x23\x54\x3d\x08\x6b\x7f\xd2\xc6\x4a\x9b\xe5\xbd\xf8\xdd\xdf\x72\xfd\x70\xdd\xee\x5c\x7f\xec\x54\x37\x92\xd3\xe6\x30\xe5\xa7\x6c\x74\xce\x1f\xd1\x2d\xa8\x44\xbc\xfb\xe1\xef\x3e\xf5\xfe\xed\xdd\x0f\x69\x6d\x5f\x94\x62\xaf\x11\x4e\x54\xa8\x1b\x77\x14\xfc\xdb\xb7\x1e\xfe\xaf\x6f\xe9\x2f\x6d\xc4\x8b\x52\xec\x35\x42\xea\xf2\x58\x1f\x05\xfc\x77\x3d\xea\xbf\x10\xf8\x50\x64\x7b\x64\x88\x5f\x69\x42\x47\x01\x9b\xef\x0b\x53\xa6\x4a\xf6\x28\x91\xd2\xbe\xd3\xa0\xf4\xe2\xac\x3b\x95\x79\xc4\xee\x47\x9c\xc1\xc8\x7f\x41\xae\x7e\xb3\x58\x5d\xcb\xf8\x05\x68\xcf\x74\x61\xcc\x30\x26\xde\x28\x3a\xe3\x97\x01\xfb\x5e\x87\x94\x0c\x3d\xf9\x7a\xf2\xff\x00\x00\x00\xff\xff\xec\xe4\x79\x95\x81\x24\x00\x00")

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

	info := bindataFileInfo{name: "wski18n/resources/en_US.all.json", size: 9345, mode: os.FileMode(420), modTime: time.Unix(1515698093, 0)}
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

