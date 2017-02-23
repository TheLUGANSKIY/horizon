// Code generated by go-bindata.
// sources:
// latest.sql
// migrations/1_initial_schema.sql
// migrations/2_index_participants_by_toid.sql
// migrations/3_use_sequence_in_history_accounts.sql
// migrations/4_add_protocol_version.sql
// DO NOT EDIT!

package schema

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

var _latestSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x5a\x6d\x6f\xdb\x38\x12\xfe\x9e\x5f\x41\xec\x17\x27\x80\x1d\xe4\xa5\x49\x5b\x07\x59\xc0\x4d\xb4\x57\xe3\x5c\x65\x1b\xcb\xd7\x2d\x16\x0b\x82\x96\xc6\x32\xaf\x12\xa9\x92\x54\x6a\xef\xe1\xfe\xfb\x41\x6f\xb6\x2c\x8b\x7a\xb1\x95\xde\x47\x9b\xa3\x99\xe7\x99\x19\x0e\x67\x28\x0d\x06\x27\x83\x01\xfa\x9d\x4b\xe5\x0a\x98\x7e\x9e\x20\x87\x28\x32\x27\x12\x90\x13\xfa\xc1\xc9\x60\x70\x12\xad\x3f\x86\x7e\x00\x0e\x5a\x08\xee\x6f\x05\x5e\x40\x48\xca\x19\x7a\x7f\x7e\x7b\x7e\x99\x93\x9a\xaf\x51\xe0\xe2\xe8\xf1\x82\xc8\xc9\xd4\xb0\x90\x54\x44\x81\x0f\x4c\x61\x45\x7d\xe0\xa1\x42\xf7\xe8\xe2\x2e\x5e\xf2\xb8\xfd\x6d\xff\x5f\xea\x78\x80\x29\xc3\x4a\x10\x26\x89\xad\x28\x67\x58\x82\x8c\xf4\xee\x0b\xdb\x1e\x8d\x54\x03\xb3\xb9\x43\x99\x8b\xee\x51\x6f\x66\xfd\xf6\xae\x77\x97\xd9\x66\x0e\x11\x0e\xb6\x39\x5b\x70\xe1\x53\xe6\x62\xa9\x04\x65\xae\x44\xf7\x88\xb3\x54\xc7\x12\xec\x6f\x78\x11\xb2\xc4\xd6\x9c\x3b\x14\xa2\xf5\x05\xf1\x24\xec\x98\xf1\x29\xc3\x3e\x48\x49\xdc\x58\xe0\x07\x11\x8c\x32\x37\x11\x11\xfc\x07\x96\x60\x87\x82\xaa\x75\xa4\x7c\xb1\xb8\x4b\x1d\x00\x44\xd8\x4b\x1c\x10\xb5\x44\xf7\x28\x08\xe7\x1e\xb5\xfb\x91\xc7\x6c\xa2\x88\xc7\xdd\x54\xcc\x81\x05\x09\x3d\x85\x15\x99\x7b\x20\x03\x62\x43\x44\xa6\x57\x58\xfd\x41\xd5\x12\x73\xea\xe4\xf0\x9d\x24\x21\x35\x89\x0f\x43\xe4\x72\x11\x60\x9f\xba\x82\x44\x5c\xe4\x1d\xb2\xd6\x01\x0c\x91\x35\xfa\x30\x31\xee\xd0\xd4\x5e\x82\x4f\x86\x29\x88\x3b\xf4\xf4\x83\x81\x18\xa2\x41\x1c\xf6\x87\x67\x63\x64\x19\x89\x68\x51\x0f\x3a\x3d\x41\x08\x21\xea\x20\x05\x2b\x85\xcc\x27\x0b\x99\xb3\xc9\xa4\x1f\xff\x4b\x82\xc0\xa3\xe0\x60\xa2\x50\x14\x1f\xa9\x88\x1f\xa0\x08\x68\xfc\x13\xfd\xcd\x19\x9c\x9c\xdd\x9d\xec\x02\x5d\x52\xa9\xb8\x58\x63\x62\xdb\x3c\x64\x4a\x62\xea\x60\x09\xdf\x33\xc0\x53\xe3\xf3\xcc\x30\x1f\x1a\x62\xce\xa4\x75\x5a\x63\x98\x53\x6b\xf4\x6c\xa1\x2f\x63\xeb\x23\xba\x8c\xff\x18\x9b\x0f\xcf\xc6\x27\xc3\xb4\xd0\x87\xaf\xe9\x5f\xe6\x13\xfa\x34\x36\xff\x35\x9a\xcc\x8c\xcd\xef\xd1\x1f\xdb\xdf\x0f\xa3\x87\x8f\x06\xba\xac\x23\x73\xb0\xdb\x8b\x8a\xb6\x7e\x9f\x53\x97\x32\x85\x1e\x8d\xdf\x46\xb3\x89\x85\x18\xac\xd4\x0b\xf1\x4e\x7b\x1a\xc6\xbd\xe1\x50\x80\x6b\x7b\x44\xca\xb3\x62\xb8\x1c\x47\x80\x94\xc8\x5e\x12\x41\x6c\x05\x02\xbd\x10\xb1\xa6\xcc\x3d\xbd\x7d\x73\xa6\x0f\x14\x2c\x16\x60\x77\x40\x2d\xd5\x93\x32\x2b\xc0\xc7\x5b\xa6\xbb\xa0\x33\x39\x1e\x40\x92\x92\x5a\xc9\x5f\xb8\x70\x40\xfc\x82\x28\x53\xe0\x82\x28\xac\xaa\x75\x00\x9a\x25\x07\x14\xa1\x9e\x44\xff\x96\x9c\xcd\xf5\x7e\xf0\xc0\x71\x41\x1c\xef\x87\x54\x4f\xea\x07\x09\xdf\x43\x60\xb6\x0e\x5b\x22\x8c\x97\x44\x2e\xcb\xe3\x56\x90\x0f\x04\xbc\x50\x1e\x4a\x5c\xfb\x60\xea\x96\x5c\xad\x8d\x03\xb1\xc1\x91\x25\xdc\x45\xc1\xc2\x36\x10\xcd\xe4\x6d\x8f\xcb\xb2\x1a\x11\x55\xf3\x4d\x99\x28\x3e\x23\x80\xa8\xda\x87\x12\xd9\x30\x70\x1a\xcb\x6e\x52\x27\xfd\xe9\x07\x5c\x28\x10\x38\x3b\xbd\x8a\x5c\x2e\x8b\x49\xc4\x15\xf1\xb0\xcd\x29\x93\xe5\x39\xb8\x00\xc0\x01\xe7\x5e\xf9\x6a\x74\x98\xe2\x05\xe8\x62\x1d\x2f\x0b\x90\x20\x5e\x74\x22\x3e\x59\x61\xb5\xc2\x12\x14\x96\xf4\x6f\x9d\x54\x20\xb8\xe2\x36\xf7\xb4\xbc\xb6\x31\xd2\xa7\xfb\x36\xce\x01\x11\x8a\xda\x34\x20\x5d\x14\xb8\x72\xb5\xdb\x72\x57\xce\xa8\x79\x15\xa8\xaf\x2b\x6d\x29\x77\x7b\x40\x55\xda\xf8\x59\xc7\x55\x2b\xa2\xe8\xe9\x8b\x69\x3c\xa2\x0f\x5f\x6b\x18\x8f\x26\x96\xf1\xdc\x92\xf0\x46\x77\x8d\xf8\x39\x75\x6a\xb9\x74\x98\x9b\xfb\xc7\x6f\xa1\x0e\xe4\xaa\xa6\x4e\x26\x6e\x8e\xec\x84\x4a\x7c\x32\x1d\x79\x30\x25\x7f\x49\x1e\x0a\x1b\xb2\xec\xd6\x1c\x09\xd9\x36\xef\xf5\x86\xc3\x3d\x89\x06\xfb\x20\x4f\xaf\xd3\xcd\xaf\x53\xdc\x74\xfb\x37\xf1\xfb\x31\x05\x40\x87\xaf\xdb\x12\x50\x63\xe5\x67\x15\x81\x96\x64\x8f\x2c\x03\x35\xd6\xf6\x0b\x81\xee\x81\x8a\x52\x90\x7b\xa4\xd3\x5c\xcd\xf2\x33\x0f\xa9\x71\x47\x96\x36\x62\x35\x7d\x5e\xd3\x6a\x51\xbd\xf1\x4b\x65\xb7\xa6\xf5\x2d\x0b\xd1\x6e\x3d\x5d\xbb\xf7\x7f\x69\xd8\xd4\x0a\x03\x7b\x01\x8f\x07\x50\x36\x8f\xaa\x55\xd4\x3e\x85\x9e\xd2\x2c\xfa\xa0\x88\x66\x29\xf2\x82\x6e\x59\x52\x97\x11\x15\x0a\x28\x1b\x9d\xde\xdf\x9e\xfd\xf9\xd7\xb6\xe2\xfe\xe7\xbf\x65\x35\xf7\xcf\xbf\x8a\x7d\x1c\xf8\x1c\xc7\x95\x7f\xbf\x3e\x6f\x74\x31\xce\xa0\xb2\x82\x6f\x75\xed\xab\x49\x99\x51\x1f\xf0\x9c\x87\xcc\x91\x51\xe4\xde\x09\xc2\xdc\x8a\x99\x5c\xd3\x9c\x51\x27\xdb\x4d\x29\xb6\x46\x25\x20\xd9\x4e\x4f\xe6\xa4\xee\x70\x47\x89\xfc\xc3\xd3\x64\xf6\xc9\x8c\x42\x3e\x35\x2c\xfd\xbc\x5b\xd9\x4f\xe4\xa7\xdf\xb6\x75\xaf\x3b\x9a\x5a\x0b\xad\x88\xd6\x54\xcc\x72\xaa\x8f\x44\x11\xb4\xe0\xa2\xc1\xad\x10\x7a\x1c\x59\xa3\x1a\x8a\x63\x73\x6a\x3c\x5b\x68\x6c\x5a\x4f\x7b\x37\x43\xf1\x41\x33\x45\xa7\xbd\x4b\x4c\x19\x55\x94\x78\x58\xc6\xba\xce\xe5\x77\xaf\xd7\x47\xbd\xab\x8b\xcb\xb7\x83\x8b\xab\xc1\xd5\x35\xba\xbc\x1e\xde\xbc\x19\xde\x5c\x9e\x5f\xdc\xbc\xbf\x79\xfb\x6e\x70\x71\xdb\x3b\xbb\x6b\xa6\xfd\x0a\x53\xe6\xc0\x6a\xd7\x05\xf3\x35\x56\x9c\x3a\x95\x96\x6e\xaf\x2f\xae\xda\x18\xba\xc6\xa1\x84\x4d\xb1\xc4\x94\xe1\xe2\x15\x4b\xb5\xb9\x9b\x37\xb7\x6f\xdb\xd8\x7b\x83\x89\xe3\xe0\xe2\xac\x56\x69\xe3\xed\xd5\xfb\x8b\xdb\xd4\x86\x26\xe2\x95\x37\x52\x4d\x42\x7e\xd0\x6d\x5d\x94\xc9\x35\x7a\xa7\xc6\xc4\x78\xb0\x72\xd7\x9f\xe7\x12\xaa\x6f\xb2\xfa\xe8\xb2\x9f\xdc\x75\xd6\xd3\x2d\xbb\xa4\x6a\xc3\x56\xa3\xb6\xec\xce\xa7\x03\xb5\x0d\x66\xeb\xc3\x43\xd5\x6e\xb8\xeb\x22\x70\xd5\x25\xb9\x4d\x18\x35\xc3\x5c\x07\x2e\x6f\x34\xd3\x1c\xee\xf4\xb6\xcd\x74\x17\x6e\xaf\x3b\x20\xda\x38\x5e\xdb\x3a\xb7\x77\x49\xb1\xe0\x15\x7e\xe3\xe0\x1b\xac\x33\x13\x0f\x4f\xe6\xd4\x7a\x1e\x8d\xcd\xd6\x67\x6d\x41\x6b\xdc\xed\x8c\x1e\x1f\x73\x1a\x4b\x0d\xa3\xdf\x9f\xc7\x9f\x46\xcf\x5f\xd1\x3f\x8d\xaf\xe8\x94\x3a\x6d\x5b\xa1\x9a\x94\xef\x86\x5b\xb5\x91\x32\xaa\x0d\x60\x35\x66\xae\xed\x5e\x6a\xf3\xae\x5b\xf6\x3a\x33\x55\xfc\x2b\xa1\xd5\x7a\x60\xbe\x39\x83\x32\x16\x63\xf3\xd1\xf8\xa3\xd9\xf4\x18\x8b\xe6\x54\xa0\x27\xb3\x7c\x96\x9c\x4d\xc7\xe6\x3f\xd0\x5c\x09\x00\x74\x9a\x0a\xf7\xf7\x86\xb5\x32\x70\xd1\xcc\x79\x0c\xb2\x78\x66\x6d\x04\xab\x38\xe9\x96\xa1\x49\xce\xc6\x63\xf0\x24\x1a\x9a\x21\x2a\x8c\xd1\xfd\xfd\x89\xb9\x34\xa1\x31\x44\xbd\x62\xbc\x7e\x00\xd2\x99\x39\xfe\x3c\xcb\x00\x17\xd4\xe5\x61\x67\xaf\xb6\x76\x10\x97\xdd\x59\xf7\xb3\xb7\x54\x3a\xb0\xdb\x19\xe4\x48\x98\xd4\x69\x0c\x70\x7b\x53\xd6\x2f\xbd\x68\xaf\x01\xcd\x03\x1c\x74\x85\x3b\xd5\x95\x87\xae\x29\xc4\x07\x31\x29\x27\xa0\x56\xdd\x11\x48\x75\x69\x72\xfa\x40\x0a\xbb\xd7\x9e\xfb\x24\x78\x10\x65\xe5\x92\x1f\xc4\x21\x05\xbf\xd5\x71\xa8\xf3\xab\x1d\xbd\x79\x23\x39\x5f\x77\xe1\xeb\x5d\x75\x79\xc8\xd9\xeb\xd5\x1d\x8c\xe5\x88\xf2\x7e\xed\x0a\xd6\x9e\xce\x66\xe5\xad\x0c\xa0\x4a\x42\xa2\x8e\x09\xeb\x56\xc7\xe1\x29\x59\x93\x7e\xc9\x78\xbe\x37\xc4\x71\x86\xd3\x0f\x0e\x8e\xf5\x6a\xad\x81\x3c\xb5\xcd\x07\x14\xbb\xc7\x6c\x22\xd8\x02\xfb\xf1\xc9\x50\xa5\xbb\x1e\x71\xad\xa3\xd3\x92\x1e\xe9\x53\xeb\x00\x0e\x4e\x91\x4a\xad\xb5\x67\x48\x24\x54\x03\x34\xdd\x90\x91\xca\xcd\xb7\x00\x1d\xa1\x2d\x53\x5d\x5b\x0b\x36\x92\xcd\x71\x77\x9d\x0c\x3b\xaa\x0f\x29\x5e\x7a\x75\x85\x4f\x1a\xba\x77\xf4\xde\x47\x13\xb5\xf0\x0b\x0f\x34\x27\x93\xfb\x86\xe5\xd5\xfc\x9f\xff\x4e\xa6\x8e\x49\x4e\xb6\x39\x89\xb2\x2f\x72\x5e\x8d\x4d\xe9\xe7\x3f\x75\xb4\xca\x1e\x6a\xce\x2f\xeb\xc8\x5f\x8d\xd3\xe6\xf5\x55\x1d\x0f\xed\xe8\xb4\xab\x7a\x7b\xc7\xf4\x1a\x5b\xbb\xa8\xbd\xb4\x9b\x6a\xbb\xc1\x77\x95\xee\x9e\xc7\x1d\xed\xf0\x2a\x13\x4d\x38\xb4\x6a\x12\x0a\xc6\xba\x3b\xbe\xf6\x15\x37\xc2\x5e\x7f\x88\xe5\x3b\xb7\xd7\x48\x9b\x7d\xfd\x07\xf7\x8d\x4a\x10\x07\x36\x07\x79\x36\xae\xe2\x39\xe7\xdf\x0e\xf6\x72\x85\xce\xda\x16\xe1\xf4\x34\xfb\x8e\x65\xf0\xeb\xaf\xa8\x27\xb9\xe7\x60\x22\x25\xa8\x38\x3e\xbd\xe1\x50\xc1\x4a\x9d\x9d\xf5\x91\x5e\xd0\xe6\x4e\x33\x41\x2a\x65\x08\x42\x2f\x3a\xe7\xa1\xbb\x54\x8d\xcc\xef\x88\x56\x03\xd8\x11\x2d\x40\x38\x43\x5f\x3e\x1a\xcf\x46\x92\x64\xe8\x1e\x5d\x5f\xe7\x02\xa6\xfb\x1a\x1f\xd9\xdc\x0f\x3c\x50\x10\x47\xe2\x7f\x01\x00\x00\xff\xff\x1e\xd8\xf0\xbd\xba\x2f\x00\x00")

func latestSqlBytes() ([]byte, error) {
	return bindataRead(
		_latestSql,
		"latest.sql",
	)
}

func latestSql() (*asset, error) {
	bytes, err := latestSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "latest.sql", size: 12218, mode: os.FileMode(420), modTime: time.Unix(1487879691, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations1_initial_schemaSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x5a\x5f\x6f\xdb\xc8\x11\x7f\xf7\xa7\x18\xdc\x8b\x6c\xd4\x6a\x2f\xb8\xe2\x70\x95\xe1\x03\x14\x99\x69\x84\xca\x54\x22\x51\x4d\x82\xc3\x61\xb1\x22\x47\xd4\xd6\xe4\x2e\xb3\xbb\x74\xa4\x2b\xfa\xdd\x0b\x52\x24\xc5\xff\xa4\x1c\xc9\xf7\x28\xee\xec\xcc\xfc\x66\x66\x7f\x33\x5c\x6a\x38\x84\xbf\xf8\xcc\x95\x54\x23\xac\x82\xab\xe1\xf0\x6a\x38\x84\x0f\x42\x69\x57\xe2\xf2\xe3\x0c\x1c\xaa\xe9\x9a\x2a\x04\x27\xf4\xe3\xe5\xab\xa5\x61\x81\xd2\x54\xa3\x8f\x5c\x13\xcd\x7c\x14\xa1\x86\x7b\xf8\xf1\x2e\x5e\xf2\x84\xfd\x54\x7d\x6a\x7b\x2c\x92\x46\x6e\x0b\x87\x71\x17\xee\x61\xb0\xb2\xde\xfd\x32\xb8\x4b\xd5\x71\x87\x4a\x87\xd8\x82\x6f\x84\xf4\x19\x77\x89\xd2\x92\x71\x57\xc1\x3d\x08\x9e\xe8\xd8\xa2\xfd\x44\x36\x21\xb7\x35\x13\x9c\xac\x85\xc3\x30\x5a\xdf\x50\x4f\x61\xc1\x8c\xcf\x38\xf1\x51\x29\xea\xc6\x02\xdf\xa8\xe4\x8c\xbb\x77\x57\x09\x3c\x93\xfa\x38\x82\xc0\x0b\x5c\xf5\xd5\xbb\x03\x6b\x1f\xe0\x08\x8c\xcf\x96\x61\x2e\xa7\x73\xf3\x0e\x96\xf6\x16\x7d\x3a\x82\xe1\x1d\xcc\xbf\x71\x94\x23\x18\xc6\xc8\x27\x0b\x63\x6c\x19\x47\x49\x98\xbe\x03\x73\x6e\x81\xf1\x79\xba\xb4\x96\xa9\x42\xf8\x34\xb5\xde\xc3\x72\xf2\xde\x78\x1c\x43\xe0\x12\x9b\x6a\xea\x89\xc8\x7a\xc1\xfc\x51\x4b\xc9\x91\xc9\xfc\xf1\xd1\x30\xad\x16\x37\x0e\x02\x30\x37\xab\x4a\x60\xba\x84\xc1\x87\xd9\xdf\x02\x37\x4a\x5e\x20\x85\x8d\x4e\x28\xa9\x07\x1e\xe5\x6e\x48\x5d\x1c\x94\xfd\xd8\x2a\x2d\x24\x9e\x2f\x0a\x07\x7d\xc5\x20\x84\x6b\x8f\xd9\xcd\x01\x28\xba\xf0\x32\xfc\x89\xd9\x08\x7e\x54\xb2\xa0\xf7\x01\xc2\x46\x48\x88\x9e\x47\x15\xa7\x50\x2b\x10\x1b\xb8\x7e\xc2\xfd\x2d\x3c\x53\x2f\xc4\x1b\x08\x28\x93\x2a\x0e\x49\x5c\x86\x48\xa5\xbd\x25\x01\xd5\x5b\xb8\x4f\xbc\xbe\x2d\xa6\x30\x12\x73\x70\x43\x43\x4f\x13\x4d\xd7\x1e\xaa\x80\xda\x18\x95\xf3\xa0\xb4\xfa\x8d\xe9\x2d\x11\xcc\xc9\x55\x68\x31\xee\x2c\xf2\x6c\x4f\xa8\x6d\x8b\x90\x6b\x95\xc2\xb7\xc6\x6f\x67\xc6\x11\x7c\x12\xbb\x2c\x02\x77\x60\x65\x66\x47\xf9\x7c\xc4\xfb\x2a\x5a\xe1\xfa\x0a\x00\x80\x39\xb0\x66\x2e\xe3\x3a\xce\x94\xb9\x9a\xcd\x6e\xe3\xe7\xd4\x71\x24\x2a\x05\xf6\x96\x4a\x6a\x6b\x94\xf0\x4c\xe5\x9e\x71\xf7\xfa\xe7\xbf\xdf\x5c\xdd\x54\x6a\x25\xd1\x8e\x9b\x0d\xda\xe7\x76\x39\x51\x9a\x78\x5c\x02\x42\x9a\x10\xa4\x72\x22\x40\x49\x63\x5e\x68\x92\xfc\x41\x48\x07\xe5\x0f\xc0\xb8\x46\x17\x65\x69\x35\xae\x97\xfa\x25\x07\x35\x65\x9e\x82\xff\x28\xc1\xd7\xcd\x41\xf1\xd0\x71\x51\x9e\x39\x28\x89\xd2\x24\x28\x0a\xbf\x86\xc8\xed\x26\x47\x0f\xc2\x64\x4b\xd5\xb6\x3e\xa3\x25\xf9\x40\xe2\x33\x13\xa1\x22\x9d\x1b\x93\x18\x49\xca\x15\x3d\xb0\x6f\x9c\x95\xcc\x8f\x07\xe3\xdd\x78\x35\xb3\xe0\xc7\x92\x85\x63\x56\xfa\xc9\xdb\x9e\x50\xe8\x10\xaa\x21\xea\x20\x4a\x53\x3f\x80\xe8\x20\x45\xbd\x24\x7a\x02\x7f\x08\x8e\xe5\x3d\x12\xa9\xee\xdc\x74\x90\x0d\x03\xa7\xb7\x6c\x56\x47\xc9\x4f\x3f\x10\x52\xa3\x24\xcf\x28\x15\x13\xbc\x82\xe5\x4d\xb9\xa2\x84\xa6\x1e\xb1\x05\xe3\xaa\xbe\x20\x37\x88\x24\x10\xc2\xab\x5f\x8d\x9a\x2e\xd9\x60\x53\xae\xe3\x65\x89\x0a\xe5\x73\x93\x88\x4f\x77\x44\xef\x88\x42\x4d\x14\xfb\xa3\x2a\xd5\x5c\xca\xc7\xb4\x05\x54\x6a\x66\xb3\x80\x9e\x9d\xa1\xea\x6d\x1c\xf9\xaa\x1e\x53\xff\xe3\xde\x4d\x20\xa7\xe2\x27\xcc\x21\x0a\xbf\xa6\x61\x58\x1a\x1f\x57\x86\x39\x69\x89\x44\x1e\x7c\x2a\xdd\xcf\x46\x8c\x60\x69\x8d\x17\xd6\xa1\x91\xbe\x89\x1f\x4c\xcd\xc9\xc2\x88\x5b\xdf\xdb\x2f\xc9\x23\x73\x0e\x8f\x53\xf3\xdf\xe3\xd9\xca\xc8\x7e\x8f\x3f\x1f\x7f\x4f\xc6\x93\xf7\x06\xbc\x39\x0b\x50\x98\x7f\x32\x8d\x07\x78\xfb\xa5\x03\xf1\x78\x66\x19\x8b\x13\x01\x67\xba\x3b\xc4\xff\xca\x9c\x4e\x2c\x97\x2a\xd4\xae\x66\x9a\xa7\xc7\xc6\x86\x1b\x04\x1e\xb3\x0f\xb8\xe2\x7e\xf4\x9d\xed\xe8\xf0\x48\x89\x50\xda\x98\x96\x7a\x03\xf7\xa7\x3c\x35\x18\x8c\x46\x15\x89\x1e\x87\x22\x0f\xef\x72\xb4\xd0\x64\x25\x8e\x7d\x03\x2d\xd4\xed\xad\x4f\xc0\xf7\x90\x42\x93\x67\xe7\xa5\x85\x0e\x2b\xaf\x45\x0c\x27\x82\xfd\x4e\x6a\xe8\xb0\x56\x25\x87\xa6\x0d\x2d\xf4\x90\xdb\x72\xb9\x92\x4d\x29\x22\xef\x5f\xef\x71\x2c\x99\xc2\x3a\x86\xbc\xbe\x0c\xd2\x4e\x06\xb5\xb2\x47\xd3\xcd\xf3\x0a\x6d\x6c\xcd\x4d\xb3\xde\x9f\x32\xad\xe9\x1d\x41\xfe\x8c\x9e\x08\x10\x34\xee\x2a\x54\xbd\x8b\x66\xa7\xd0\xd3\x0d\x8b\x3e\x46\xaf\x90\xb5\x4b\x51\x14\x9a\x96\x15\x73\x39\xd5\xa1\xc4\xba\x37\xaa\x7f\xfc\x7c\xf3\xdb\xef\x47\x16\xfe\xef\xff\xea\x78\xf8\xb7\xdf\xcb\x43\x1c\xfa\x82\xc4\xdd\xa0\xca\xd9\x99\x2e\x2e\x38\xb6\xb2\xfa\x51\x57\x55\x4d\x82\x8c\xf9\x48\xd6\x22\xe4\x8e\x8a\x32\xf7\x8b\xa4\xdc\xc5\x98\x0c\xf3\x87\x89\x39\xe9\xd1\x49\x6c\xf7\x3a\xef\x87\xe3\x32\x37\x67\x5d\xdd\x1d\x0e\xf2\x93\xf9\x6c\xf5\x68\x46\x29\x8d\x5e\xa8\x53\x94\x1c\x77\xfa\x99\x7a\xd7\x83\x5e\x03\xc5\x60\x34\x92\xe8\xda\x1e\x55\xaa\xc2\xe8\x67\x43\xd1\xd8\xac\x4e\xc2\xd1\xc1\x7e\x6d\x48\x3a\x42\x11\x3c\xe1\xfe\x78\xad\x62\x2e\xad\xc5\x78\x6a\xb6\xa0\xad\x12\xde\x89\x09\x8c\x4b\x69\xfc\xf0\x90\xb3\xd6\xc7\x47\xf8\xb0\x98\x3e\x8e\x17\x5f\xe0\x5f\xc6\x17\xb8\x66\xce\xe9\x3d\xf8\x82\x48\x9b\x6c\xb6\x61\x6d\xf5\xb3\x13\xed\x3a\x1b\x50\x52\x48\x53\xf3\xc1\xf8\xfc\x82\x46\x15\xef\xcb\xe9\x83\xb9\x59\xdf\xb6\x56\xcb\xa9\xf9\x4f\x58\x6b\x89\x08\xd7\x89\xf0\x6d\xa5\x2f\xd4\x79\x1a\xb5\xb7\xb3\xb9\x19\xf7\xca\x5e\x3e\x96\x3b\x6c\x9d\x6b\x87\x86\x7a\x36\xe7\x0e\xea\xfa\xb9\x57\xea\xe5\xb7\xd5\xb6\x5d\x5b\xe3\x04\xc9\x7a\x7f\x58\xff\x5e\xb7\x57\xe6\xf4\xe3\x2a\xf5\xbe\xa4\x3b\x8f\x21\xbd\x76\x2b\xb8\x5f\xf7\x9a\x7d\x9b\xde\xa0\x35\x79\x7e\xa4\xd5\x73\xfa\xcc\x9c\xde\xde\x1e\xa7\xfa\xdb\xda\x8b\x82\x0e\x04\x22\x20\xc1\x45\x40\x24\x8a\xf3\x38\x1a\xfa\xdf\x8b\x60\x55\xd1\x64\x37\x7a\xeb\xfd\xd9\x01\x15\x75\xe7\x31\xa5\x77\x95\x05\x10\xf5\xee\xe5\x4f\xef\x45\x7c\xac\x18\xe8\x77\x6c\x6b\xbc\x65\xdc\xc1\x1d\x29\xdf\xab\x13\xc1\x49\x72\x79\x7e\x56\xd7\x3b\xad\xe5\x71\x64\x97\xfc\x45\xf6\x3e\x08\x9e\x00\xe4\xcc\xe1\x6f\x33\xd4\xed\x7e\x67\x0a\x12\x0a\x88\xf4\x45\x73\xf1\x79\xe8\xbd\xd5\x44\x27\x01\x45\x42\x1d\x5e\x27\x87\x23\x52\x99\x5d\x72\x5f\xc2\xf5\x3a\x3b\x9d\x87\x34\x93\xec\x0f\xe2\xa2\x35\x53\xb0\xf3\x12\x8a\x69\x56\x57\xba\xc5\xbf\x70\x0a\x2a\x1f\x0d\x3a\xb1\x94\x36\xf4\x47\x96\xfb\x86\xf3\x3a\x99\xc9\x7f\x34\xea\x82\x95\x93\xed\x8f\xa8\xee\xf3\xd4\xeb\x40\xab\xfd\x30\xd6\x85\xb1\x6e\x53\x7f\xb0\xe9\xa4\xf8\x3a\x00\xb3\x8b\x9e\x2e\x50\x8d\x93\x7f\x51\xf5\xf1\x8e\xfc\xe2\xdc\x50\x36\x55\x3b\x55\x9d\xca\x10\x45\xa5\xc5\x7b\xe4\x4b\x50\x44\x9b\xbd\x3e\x80\x8a\x3b\x4e\x03\x77\xa1\x9e\x59\xb5\xd2\x0b\x48\x5d\xe7\x8c\x87\x66\xbd\xbb\xd0\x34\x9e\x28\x6e\x18\x08\x5f\x38\x8f\x57\x13\xd2\x9c\x8f\xfc\xf8\x79\xf1\xe3\x52\x35\xf6\xe2\x49\x58\x4b\xea\x60\x36\x1b\xa5\xef\x92\x64\x2d\xc4\xd3\x79\x0a\xaa\xc5\x40\xe7\x08\x76\x7d\x9d\x7e\x17\x1b\xfe\xfa\x2b\x0c\x94\xf0\x1c\x42\x95\x42\x1d\x97\xe2\x60\x34\xd2\xb8\xd3\x37\x37\xb7\xd0\x2c\x68\x0b\xa7\x9f\x20\x53\x2a\x44\xd9\x2c\xba\x16\xa1\xbb\xd5\xbd\xcc\x17\x44\xdb\x1d\x28\x88\x96\x5c\xb8\x81\x4f\xef\x8d\x85\x71\x38\x4f\x70\x0f\x3f\xfd\x94\xcb\x5e\xd3\xbf\xf9\xc0\x16\x7e\xe0\xa1\xc6\x38\x13\xf9\x3f\x02\x3e\x88\x6f\xfc\xca\x91\x22\x80\xf8\x3f\x4e\xf5\xe5\x62\x53\x65\x53\x07\xef\x3a\x04\x8b\x07\xaa\x6d\x53\x8e\x23\x7a\x89\xf5\xd7\x9c\xb6\xb6\x36\x99\xb4\xaa\xda\x64\xb2\x37\x96\x4c\xe8\xff\x01\x00\x00\xff\xff\x5d\xb2\x1f\x7d\x3f\x29\x00\x00")

func migrations1_initial_schemaSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations1_initial_schemaSql,
		"migrations/1_initial_schema.sql",
	)
}

func migrations1_initial_schemaSql() (*asset, error) {
	bytes, err := migrations1_initial_schemaSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/1_initial_schema.sql", size: 10559, mode: os.FileMode(420), modTime: time.Unix(1461452707, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations2_index_participants_by_toidSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x8f\xb1\xca\xc2\x50\x0c\x46\xf7\x3c\x45\xc6\xff\x47\xfa\x04\x9d\xc4\x16\xe9\xd2\x4a\xb5\xe0\x76\x49\xdb\x8b\xcd\xe0\xcd\x25\x37\x20\x7d\x7b\x41\x07\x5b\xbb\xb8\x86\x8f\x73\x72\xb2\x0c\x77\x77\xbe\x29\x99\xc7\x2e\x02\x1c\xda\x72\x7f\x29\xb1\xaa\x8b\xf2\x8a\x93\x44\xd7\xcf\x6e\x12\x1e\xb1\xa9\x71\xe2\x64\xa2\xb3\x93\xe8\x95\x8c\x25\xb8\x48\x6a\x3c\x70\xa4\x60\x09\xbb\x73\x55\x1f\xb1\x37\xf5\x1e\xff\xb6\x5b\x1e\xff\xf3\x2f\xbc\xbd\xf1\xb6\xc6\x9b\x52\x48\x34\xfc\x28\x58\xae\x5f\x0a\x58\x26\x15\xf2\x08\x00\x45\xdb\x9c\xb6\x49\xf9\xea\xfe\xf9\x25\x87\x67\x00\x00\x00\xff\xff\x33\xec\x54\x7a\x15\x01\x00\x00")

func migrations2_index_participants_by_toidSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations2_index_participants_by_toidSql,
		"migrations/2_index_participants_by_toid.sql",
	)
}

func migrations2_index_participants_by_toidSql() (*asset, error) {
	bytes, err := migrations2_index_participants_by_toidSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/2_index_participants_by_toid.sql", size: 277, mode: os.FileMode(420), modTime: time.Unix(1461452707, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations3_use_sequence_in_history_accountsSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x91\x4d\x6b\xb3\x40\x14\x85\xf7\xf3\x2b\xce\x2e\xca\xfb\x66\x91\x6d\x5c\x4d\xc6\x1b\x22\x8c\x63\x3b\x5e\xdb\x64\x25\xa2\x43\x3a\x90\x6a\xeb\xd8\xaf\x7f\x5f\x48\xd3\x0f\x08\x6d\xa1\xcb\x73\x78\xe0\x39\xdc\x3b\x9f\xe3\xdf\xad\xdf\x8f\xcd\xe4\x50\xdd\x09\x65\x49\x32\xa1\xa4\xcb\x8a\x8c\x22\xdc\xf8\x30\x0d\xe3\x4b\xdd\xb4\xed\xf0\xd0\x4f\xa1\xf6\x5d\x1d\xdc\xbd\x00\x80\x92\xa5\x65\x5c\x67\xbc\xc1\xe2\x58\x64\x46\x59\xca\xc9\x30\x56\xbb\x53\x65\x0a\xe4\x99\xb9\x92\xba\xa2\x8f\x2c\xb7\x9f\x59\x49\xb5\x21\x2c\x12\x51\x92\x26\xc5\x08\x6e\x7a\x6c\x0e\xd1\xec\x1b\xef\xec\x3f\xa2\x13\x99\xcb\x6d\xe4\xbb\x18\x6b\x5b\xe4\x67\x33\xe3\x38\x11\x52\x33\x59\xb0\x5c\x69\x42\x61\xf4\xee\x0c\xc2\x1b\xa1\x0a\x5d\xe5\x06\xbe\x43\x49\x8c\x94\xd6\xb2\xd2\x8c\xde\x3d\xff\xbc\x64\xb9\x1c\xdd\xbe\x3d\x34\x21\xc4\x89\x10\x5f\xcf\x98\x0e\x4f\xfd\x1f\xec\xa9\x2d\x2e\xde\xf5\x89\x38\xa6\xdf\xde\x90\x88\xd7\x00\x00\x00\xff\xff\x55\xe2\xdd\x2c\xbf\x01\x00\x00")

func migrations3_use_sequence_in_history_accountsSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations3_use_sequence_in_history_accountsSql,
		"migrations/3_use_sequence_in_history_accounts.sql",
	)
}

func migrations3_use_sequence_in_history_accountsSql() (*asset, error) {
	bytes, err := migrations3_use_sequence_in_history_accountsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/3_use_sequence_in_history_accounts.sql", size: 447, mode: os.FileMode(420), modTime: time.Unix(1481047100, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations4_add_protocol_versionSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\xcd\xb1\x0a\xc2\x30\x10\x06\xe0\x3d\x4f\xf1\xef\x52\x70\xef\x14\x4d\x9d\xce\x44\x4a\x32\x38\x15\xd1\xa3\x06\x6a\xae\x5c\x82\xe2\xdb\xbb\xba\x88\x4f\xf0\x75\x1d\x36\x8f\x3c\xeb\xa5\x31\xd2\x6a\x2c\xc5\x61\x44\xb4\x3b\x1a\x10\x3c\x9d\x71\xcf\xb5\x89\xbe\xa7\x85\x6f\x33\x6b\x85\x01\xac\x73\xd8\x07\x4a\x47\x8f\x55\xa5\xc9\x55\x96\xe9\xc9\x5a\xb3\x14\xe4\xd2\x78\x66\x85\x1b\x0e\x36\x51\xc4\x16\x3e\x44\xf8\x44\xd4\x1b\xf3\x6d\x39\x79\x95\xff\x9a\x1b\xc3\xe9\x97\xd5\x9b\x4f\x00\x00\x00\xff\xff\x83\xbb\x30\x2e\xbc\x00\x00\x00")

func migrations4_add_protocol_versionSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations4_add_protocol_versionSql,
		"migrations/4_add_protocol_version.sql",
	)
}

func migrations4_add_protocol_versionSql() (*asset, error) {
	bytes, err := migrations4_add_protocol_versionSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/4_add_protocol_version.sql", size: 188, mode: os.FileMode(420), modTime: time.Unix(1487224181, 0)}
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
	"latest.sql": latestSql,
	"migrations/1_initial_schema.sql": migrations1_initial_schemaSql,
	"migrations/2_index_participants_by_toid.sql": migrations2_index_participants_by_toidSql,
	"migrations/3_use_sequence_in_history_accounts.sql": migrations3_use_sequence_in_history_accountsSql,
	"migrations/4_add_protocol_version.sql": migrations4_add_protocol_versionSql,
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
	"latest.sql": &bintree{latestSql, map[string]*bintree{}},
	"migrations": &bintree{nil, map[string]*bintree{
		"1_initial_schema.sql": &bintree{migrations1_initial_schemaSql, map[string]*bintree{}},
		"2_index_participants_by_toid.sql": &bintree{migrations2_index_participants_by_toidSql, map[string]*bintree{}},
		"3_use_sequence_in_history_accounts.sql": &bintree{migrations3_use_sequence_in_history_accountsSql, map[string]*bintree{}},
		"4_add_protocol_version.sql": &bintree{migrations4_add_protocol_versionSql, map[string]*bintree{}},
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

