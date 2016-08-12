// Code generated by go-bindata.
// sources:
// templates/init/django/.dockerignore
// templates/init/django/Dockerfile
// templates/init/django/docker-compose.yml
// templates/init/rails/.dockerignore
// templates/init/rails/Dockerfile
// templates/init/rails/docker-compose.yml
// templates/init/ruby/.dockerignore
// templates/init/ruby/Dockerfile
// templates/init/ruby/docker-compose.yml
// templates/init/sinatra/.dockerignore
// templates/init/sinatra/Dockerfile
// templates/init/sinatra/docker-compose.yml
// templates/init/unknown/.dockerignore
// templates/init/unknown/Dockerfile
// templates/init/unknown/docker-compose.yml
// templates/templates.go
// DO NOT EDIT!

package templates

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

var _initDjangoDockerignore = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\x4b\xce\xcf\x2b\xcb\xaf\xe0\xd2\x4b\xcd\x2b\xe3\xd2\x4b\xcf\x2c\x01\x13\x99\xe9\x79\xf9\x45\xa9\x5c\x80\x00\x00\x00\xff\xff\x57\x31\x5f\xce\x1d\x00\x00\x00")

func initDjangoDockerignoreBytes() ([]byte, error) {
	return bindataRead(
		_initDjangoDockerignore,
		"init/django/.dockerignore",
	)
}

func initDjangoDockerignore() (*asset, error) {
	bytes, err := initDjangoDockerignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/django/.dockerignore", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initDjangoDockerfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x91\xcd\xae\xd3\x30\x10\x85\xf7\x79\x8a\x91\x60\xdb\x66\xd1\x27\x40\x25\x2c\x40\xb4\x51\x28\x48\x5d\x21\xe3\x4c\x52\x17\xc7\x63\xfc\x03\x8d\x50\xdf\x9d\xb1\xd3\xd0\xe6\xde\xbb\xb8\xbb\xcc\xf1\xcc\x99\x6f\x4e\x3e\x34\xfb\xcf\x20\xc9\xfc\xa6\x4b\xd9\x9e\x85\xe9\xa9\x28\xde\x80\x43\xab\x85\x44\xc0\x8b\x18\xac\x46\x61\x2d\x08\xd3\xce\xa5\x75\x74\x46\x19\x20\x10\x04\xa1\x34\x39\x08\x27\x04\x35\x88\x1e\x93\x36\x52\x74\x70\xeb\x61\xaf\xba\xd9\x7f\xac\xb6\x07\x50\x1e\x84\xf6\x04\xd1\x63\x0b\x3f\x46\xe8\xa3\x51\x92\x9c\x01\x65\xf2\xfc\x02\x02\xde\x93\xfc\x89\xae\x53\x1a\x8b\x6a\xf7\x0d\xde\xd5\xf5\x03\x4c\x96\x66\xdf\x25\x54\xa2\x17\x06\x70\xb0\x61\x84\x2f\xd5\xb6\xa9\x0e\xdf\x3f\x55\x47\x68\xa3\x53\xa6\x87\x41\x18\xa6\x5c\xdb\x91\xd7\x0d\x5c\xb4\x1e\xfe\x28\xad\xf9\x60\x1f\x75\x48\x28\x69\xd8\x39\x72\x79\xc7\x83\x41\x47\x39\x19\x49\x3c\x4b\x46\x8f\x99\x39\xf1\x79\x30\x88\x2d\xdf\xd4\x71\x10\x56\x59\x36\xf1\x41\x68\x5d\x6c\xf7\xf5\x91\x8d\x7f\x45\xe5\x70\x40\x13\xfc\x3a\x5c\x02\x94\xcc\x5f\x3e\x55\x8b\xe6\xeb\x2e\xcd\x6e\xe6\x61\x58\xad\xa2\xed\x9d\x68\x31\xc9\x2f\x3c\xbb\x67\xce\xaf\xa0\x93\xa4\x35\x67\xc4\x16\x41\xc9\x89\xef\xed\x5f\x8e\xf6\x5a\x4e\xd2\x04\xb7\x90\xe6\xae\x5b\xda\xd7\xb9\x65\xae\xa7\xf7\x7b\xac\xf9\xf9\x7f\x39\x81\x8f\xe1\x44\x66\xb3\xc8\xfe\x81\x83\x4f\x35\xa4\x8c\x8d\xf7\x0b\x12\x3c\xff\x91\x00\xd4\xe5\xef\xf4\xcb\xf3\x9e\x75\xf6\x2f\xfe\x05\x00\x00\xff\xff\x22\xcb\xe6\x65\xb5\x02\x00\x00")

func initDjangoDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_initDjangoDockerfile,
		"init/django/Dockerfile",
	)
}

func initDjangoDockerfile() (*asset, error) {
	bytes, err := initDjangoDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/django/Dockerfile", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initDjangoDockerComposeYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x8d\xc1\xaa\xc2\x30\x10\x45\xf7\xfd\x8a\xfc\xc0\xcb\x8b\x36\x82\x04\xba\x92\xae\xdc\xa9\x1b\x57\x92\xb4\x43\x09\xa6\x99\x92\x4c\x6b\xfd\x7b\x13\x68\xbb\x10\xdc\xcd\xbd\xf7\x70\xe6\x05\x46\x15\x8c\x99\xd1\xba\x56\x31\x9e\x4e\xf0\x93\x0d\xe8\x7b\xf0\x94\x17\xc6\xfe\xd8\xb5\x3e\x5d\xea\xdb\xe3\x5c\xdf\x53\xe1\xb4\x01\x17\xd7\xa9\x41\x3f\xe1\xcc\x07\x0c\xc4\xa5\x2c\xf9\x10\x90\xb0\x41\x57\x91\x8b\xbf\x91\xf9\x5d\x51\x18\x21\xdb\xac\x7f\x6e\xb2\x56\x93\x36\x3a\xe6\x3e\xd3\x5b\x7f\x14\x4a\x0a\x21\x96\x94\x1c\x39\xee\x8a\x15\xcf\x98\xed\x75\x07\x6a\xf9\xf5\x3f\x60\xa4\x2e\x40\xfc\x16\x1d\x64\xb9\x2f\x3e\x01\x00\x00\xff\xff\x25\x21\x30\xfe\xf3\x00\x00\x00")

func initDjangoDockerComposeYmlBytes() ([]byte, error) {
	return bindataRead(
		_initDjangoDockerComposeYml,
		"init/django/docker-compose.yml",
	)
}

func initDjangoDockerComposeYml() (*asset, error) {
	bytes, err := initDjangoDockerComposeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/django/docker-compose.yml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initRailsDockerignore = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd7\x4b\x2a\xcd\x4b\xc9\x49\xe5\xd2\xd7\x4b\xce\xcf\x2b\xcb\xaf\x00\x32\x52\xf3\xca\x80\x64\x7a\x66\x09\x97\x7e\x4a\x92\xbe\x96\x5e\x71\x61\x4e\x66\x49\xaa\x31\x2a\x4f\x37\x2b\xbf\xb4\x28\x2f\x31\x87\x4b\x3f\x27\x3f\x5d\x5f\x8b\x4b\xbf\x24\xb7\x80\x0b\x10\x00\x00\xff\xff\xa0\x04\x95\x56\x4e\x00\x00\x00")

func initRailsDockerignoreBytes() ([]byte, error) {
	return bindataRead(
		_initRailsDockerignore,
		"init/rails/.dockerignore",
	)
}

func initRailsDockerignore() (*asset, error) {
	bytes, err := initRailsDockerignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/rails/.dockerignore", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initRailsDockerfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x50\x3d\x6f\x83\x30\x10\xdd\xfd\x2b\x4e\xea\x1e\xf6\xae\x95\xda\xa9\x4d\x85\xd4\xa1\x5b\x1c\x73\xa6\x2e\xc6\x67\xf9\x23\x0a\xff\xbe\x87\x0d\x25\x44\x61\xe2\x7d\x70\xef\x3d\x5e\xdb\xe3\x3b\x28\x72\x17\xba\x36\x41\x1a\x1b\x85\x78\x62\xec\x27\x20\x67\x27\x48\x3f\x08\xda\x58\x8c\xe0\x10\x3b\xec\x40\x53\x80\x73\x76\x9d\x45\x30\x2e\x26\x69\x2d\xfb\xb3\x53\x34\x8e\xe8\x52\xf1\x5f\xd0\x75\x14\x1a\x25\x15\x03\x6b\x1c\x3b\x35\x4c\x94\xe1\xb4\x7c\xe8\xa5\x1a\x64\x8f\xa7\x99\x0c\xd0\xe3\x18\xc5\xcb\xf1\xf3\x1b\xde\x70\x9c\xb3\xa0\x3c\x8d\xf4\xbe\x59\x98\x9d\x7c\xb0\xa4\x86\x9d\x5c\x18\xae\x51\x5c\xbb\xf4\xe2\xba\x65\x44\xfb\xf5\x71\xdf\x7f\x1d\xfc\x9b\x63\x7a\x3c\x58\xc6\x88\x29\x3e\xfb\x80\xbc\xd3\xff\x17\x6a\xe5\x80\x4b\xe1\x12\xb4\xe2\xaa\xf2\x4f\xd5\xa6\xdf\xb6\x54\x5c\x35\x9f\xcf\xd6\xa8\x4d\xab\xb8\x6a\x33\xae\x81\x55\xdb\xf0\x6d\x79\xbc\xa2\x82\xc0\x81\x0f\xca\xad\x83\xe6\x2d\x01\x79\x14\xe9\xf2\xce\x97\x6a\xc4\xa1\x5c\x16\x7f\x01\x00\x00\xff\xff\x3d\xf9\xb5\xe2\xfc\x01\x00\x00")

func initRailsDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_initRailsDockerfile,
		"init/rails/Dockerfile",
	)
}

func initRailsDockerfile() (*asset, error) {
	bytes, err := initRailsDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/rails/Dockerfile", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initRailsDockerComposeYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x4f\x4d\xb2\xe2\x52\x50\x48\x2a\xcd\xcc\x49\xb1\x52\xd0\x03\x32\x73\x12\x93\x52\x73\x8a\x41\x82\x0a\x0a\xba\x0a\xc9\xf9\x79\x65\xf9\x15\x7a\x05\xf9\x45\x25\x7a\x26\x26\xc6\x7a\x05\x45\xf9\x25\xf9\xc9\xf9\x39\xb6\x25\x39\xc5\xb8\x95\x54\x54\xda\x96\x14\x95\xa6\x02\x15\x80\x44\xe1\x86\x59\x18\x58\x99\x18\x18\x18\x40\x79\x40\xb5\x20\xae\x21\x17\x20\x00\x00\xff\xff\xc0\xe1\x22\xef\x84\x00\x00\x00")

func initRailsDockerComposeYmlBytes() ([]byte, error) {
	return bindataRead(
		_initRailsDockerComposeYml,
		"init/rails/docker-compose.yml",
	)
}

func initRailsDockerComposeYml() (*asset, error) {
	bytes, err := initRailsDockerComposeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/rails/docker-compose.yml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initRubyDockerignore = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd7\x4b\x2a\xcd\x4b\xc9\x49\xe5\xd2\xd7\x4b\xce\xcf\x2b\xcb\xaf\x00\x32\x52\xf3\xca\x80\x64\x7a\x66\x09\x17\x20\x00\x00\xff\xff\xc9\x68\x92\x70\x1e\x00\x00\x00")

func initRubyDockerignoreBytes() ([]byte, error) {
	return bindataRead(
		_initRubyDockerignore,
		"init/ruby/.dockerignore",
	)
}

func initRubyDockerignore() (*asset, error) {
	bytes, err := initRubyDockerignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/ruby/.dockerignore", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initRubyDockerfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x8f\xcd\x4e\xc0\x20\x10\x84\xef\x3c\xc5\x26\xde\xcb\x43\x98\xe8\x49\x6b\x9a\x78\xf0\x56\x0a\x4b\x6d\x0a\xbb\x84\x9f\x46\xde\x5e\x8a\x35\xb1\x72\x62\x67\xbf\x61\x86\xa7\x69\x7c\x01\xcd\x74\xf0\x97\x8c\x65\xa9\x42\x3c\xb4\x31\x54\x60\x72\x15\xf2\x27\x82\xdd\x1c\x26\x20\x44\x83\x06\x2c\x47\x58\x0a\x19\x87\xb0\x51\xca\xca\xb9\xc6\x17\xd2\xec\x3d\x52\xee\xfc\x81\x64\x38\x4a\xad\x74\x1b\xdc\x46\x8d\xb4\x50\xb9\xc0\x7c\x19\x83\xd2\xbb\x5a\x71\x3e\xc5\x08\x2b\xfa\x24\x1e\xc7\xb7\x0f\x78\x46\x7f\x66\x41\x3f\x52\x85\x20\x2f\xe5\xb6\x1e\x1c\xeb\xfd\xb6\xee\x4a\xab\xd1\xa9\x5b\x7a\xa7\xfe\x2a\x62\x7a\x7f\xfd\xdf\xff\xf7\xc3\x67\xf7\x88\x29\x03\xdb\x7e\x6f\xde\x9f\xe0\xa1\xbf\x23\xbe\x03\x00\x00\xff\xff\x8b\xae\xa0\xae\x2a\x01\x00\x00")

func initRubyDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_initRubyDockerfile,
		"init/ruby/Dockerfile",
	)
}

func initRubyDockerfile() (*asset, error) {
	bytes, err := initRubyDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/ruby/Dockerfile", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initRubyDockerComposeYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x8c\xc1\x0a\x83\x30\x0c\x86\xef\x3e\x45\xf0\x6e\xe9\x98\x87\x21\xf8\x30\x5a\x03\xca\xa2\x29\x69\x3a\xed\xdb\x2f\x85\x6d\xb7\xdd\xfa\xf5\xff\xf2\x9d\x38\x0f\x0d\xc0\x9c\x37\x5a\x06\x70\xf6\x0c\xbc\xef\xd3\x61\x80\x61\x65\x68\x71\xd9\x14\x16\x0e\x4f\x94\xce\xa6\xc8\x09\x5d\xd9\x09\xce\x4d\x57\x28\x9c\x05\x92\x4e\xa2\x39\x7e\x0f\x5b\x6b\xd0\x34\x23\xa5\x1a\x06\xe8\x6c\x38\x5e\x7c\xb9\xc8\xa2\xae\xef\xef\x2e\x0a\x2b\x07\xa6\x51\x29\xfd\x57\xae\x32\xaa\x64\x34\xa1\xfe\xfe\x62\x0f\x3f\xf4\xde\xfb\x0f\x99\x5b\xf1\xd6\xbc\x03\x00\x00\xff\xff\xae\x01\x4e\xf5\xc8\x00\x00\x00")

func initRubyDockerComposeYmlBytes() ([]byte, error) {
	return bindataRead(
		_initRubyDockerComposeYml,
		"init/ruby/docker-compose.yml",
	)
}

func initRubyDockerComposeYml() (*asset, error) {
	bytes, err := initRubyDockerComposeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/ruby/docker-compose.yml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initSinatraDockerignore = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd7\x4b\x2a\xcd\x4b\xc9\x49\xe5\xd2\xd7\x4b\xce\xcf\x2b\xcb\xaf\x00\x32\x52\xf3\xca\x80\x64\x7a\x66\x09\x17\x20\x00\x00\xff\xff\xc9\x68\x92\x70\x1e\x00\x00\x00")

func initSinatraDockerignoreBytes() ([]byte, error) {
	return bindataRead(
		_initSinatraDockerignore,
		"init/sinatra/.dockerignore",
	)
}

func initSinatraDockerignore() (*asset, error) {
	bytes, err := initSinatraDockerignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/sinatra/.dockerignore", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initSinatraDockerfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x8f\xc1\x6e\x84\x20\x10\x86\xef\x3c\xc5\x24\xbd\xcb\x43\x34\x69\x4f\xad\x8d\x49\x0f\xbd\x49\x61\x70\x89\x30\x43\x00\xcd\xfa\xf6\x8b\xac\x9b\xac\xeb\x49\xbe\xf9\x86\xff\xe7\x63\xe8\xbf\x40\x33\xad\x7c\x95\xd9\x91\x2a\x49\x09\xf1\x56\x49\xdc\x80\xc9\x6f\x50\x2e\x08\xd6\x79\xcc\x40\x88\x06\x0d\x58\x4e\xf0\xbf\x90\xf1\x08\x8e\x72\x51\xde\x57\x7f\x21\xcd\x21\x20\x95\xe6\xaf\x48\x86\x93\xd4\x4a\xd7\x83\x77\x54\x4d\x0b\x1b\x2f\x30\x1e\x8b\x51\xe9\x59\x4d\x38\xee\x30\xc1\x84\x21\x8b\xf7\xfe\xe7\x0f\x3e\x31\xec\x59\xd0\x3e\xa9\x62\x94\x07\x39\x8d\x3b\xcf\x7a\x3e\x8d\x1b\xa9\x35\x9a\x75\x4a\x6f\xd6\x33\x11\xc3\xef\xf7\x6b\xff\xc7\x83\xf7\xee\x09\x73\x01\xb6\xed\xbf\xee\xde\x83\xbb\x76\x8f\xb8\x05\x00\x00\xff\xff\x9c\x51\x49\xbe\x2d\x01\x00\x00")

func initSinatraDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_initSinatraDockerfile,
		"init/sinatra/Dockerfile",
	)
}

func initSinatraDockerfile() (*asset, error) {
	bytes, err := initSinatraDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/sinatra/Dockerfile", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initSinatraDockerComposeYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x4f\x4d\xb2\xe2\x52\x50\x48\x2a\xcd\xcc\x49\xb1\x52\xd0\x03\x32\x73\x12\x93\x52\x73\x8a\x41\x82\x0a\x0a\xba\x0a\xc9\xf9\x79\x65\xf9\x15\x7a\x05\xf9\x45\x25\x7a\x26\x26\xc6\x7a\x05\x45\xf9\x25\xf9\xc9\xf9\x39\xb6\x25\x39\xc5\xb8\x95\x54\x54\xda\x96\x14\x95\xa6\x02\x15\x80\x44\xe1\x86\x59\x18\x58\x99\x18\x18\x18\x40\x79\x40\xb5\x20\xae\x21\x17\x20\x00\x00\xff\xff\xc0\xe1\x22\xef\x84\x00\x00\x00")

func initSinatraDockerComposeYmlBytes() ([]byte, error) {
	return bindataRead(
		_initSinatraDockerComposeYml,
		"init/sinatra/docker-compose.yml",
	)
}

func initSinatraDockerComposeYml() (*asset, error) {
	bytes, err := initSinatraDockerComposeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/sinatra/docker-compose.yml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initUnknownDockerignore = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\x4b\xcd\x2b\xe3\xd2\x4b\xcf\x2c\xe1\x02\x04\x00\x00\xff\xff\x9c\x10\x28\x7b\x0a\x00\x00\x00")

func initUnknownDockerignoreBytes() ([]byte, error) {
	return bindataRead(
		_initUnknownDockerignore,
		"init/unknown/.dockerignore",
	)
}

func initUnknownDockerignore() (*asset, error) {
	bytes, err := initUnknownDockerignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/unknown/.dockerignore", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initUnknownDockerfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x0b\xf2\xf7\x55\x28\x4d\x2a\xcd\x2b\x29\xb5\x32\x34\xd3\x33\x30\xe1\xe2\x72\xf6\x0f\x88\x54\xd0\x53\xd0\x4f\x2c\x28\xe0\x02\x04\x00\x00\xff\xff\xf1\xa3\x65\xfc\x1f\x00\x00\x00")

func initUnknownDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_initUnknownDockerfile,
		"init/unknown/Dockerfile",
	)
}

func initUnknownDockerfile() (*asset, error) {
	bytes, err := initUnknownDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/unknown/Dockerfile", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initUnknownDockerComposeYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xca\x4d\xcc\xcc\xb3\xe2\x52\x50\x48\x2a\xcd\xcc\x49\xb1\x52\xd0\xe3\x02\x04\x00\x00\xff\xff\xa6\xe1\xc1\x85\x11\x00\x00\x00")

func initUnknownDockerComposeYmlBytes() ([]byte, error) {
	return bindataRead(
		_initUnknownDockerComposeYml,
		"init/unknown/docker-compose.yml",
	)
}

func initUnknownDockerComposeYml() (*asset, error) {
	bytes, err := initUnknownDockerComposeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/unknown/docker-compose.yml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x9b\x59\x6f\x1b\xd7\x92\xc7\x9f\xc5\x4f\xd1\x11\x90\x40\x1a\x38\x52\xef\x8b\x80\xbc\xc4\xce\x00\x79\x98\x04\xc8\xf2\x30\x18\x0f\x82\x5e\x4e\x6b\x38\x91\x48\x5f\x92\x4a\xac\x18\xf9\xee\xf7\xfc\xaa\xaa\x25\x5a\x6c\x52\x32\x25\x21\xba\x8b\x81\x36\xc9\xee\xb3\x54\x9d\x53\xf5\xaf\x7f\xd5\x69\x9d\x9e\x06\xaf\xe7\x9d\x0b\xce\xdd\xcc\x2d\xea\x95\xeb\x82\xe6\x3a\x38\x9f\x7f\xd9\x4c\x67\x5d\xbd\xaa\x4f\x26\xbe\xc1\x72\x7e\xb5\x68\xdd\xf2\x8c\xef\x2b\x77\xf9\xee\xc2\xb7\x5b\x9e\x4e\x67\xd3\xd5\x69\xf7\xff\xf5\xec\x7c\x7e\x7a\xd2\xcd\xdb\x5f\xdd\x62\x7a\x3e\x9b\x2f\xdc\xf6\x66\x6f\xa4\x55\x3f\xbd\xd8\xd1\x46\x47\xfa\xb2\x9d\x5f\xbe\x9b\x2f\xdd\xc9\xf5\xe5\xc5\x48\xdb\x45\x3d\xbd\x58\xde\x3b\xab\xb6\xda\x39\xa9\x36\x79\xd8\x9c\x57\xcd\xf5\xfd\x53\xd2\x68\xf7\x8c\xb4\x78\xd0\x84\xcb\xe9\xac\x5e\x2d\xea\x7b\xe7\x1c\xda\xed\x9c\x76\x68\xf4\xa0\x99\xaf\x66\xbf\xce\xe6\xbf\xcf\xee\x9d\x79\x68\xb7\x73\xe6\xa1\xd1\x7d\x33\xdf\x7c\x3b\x39\x9f\xf3\xe4\xcd\xf7\xc1\x77\xdf\xff\x14\x7c\xf3\xe6\xdb\x9f\x3e\x9b\x4c\xde\xd5\xed\xaf\xf5\xb9\xbb\x6d\x3f\x99\x4c\xfd\x40\x8b\x55\x70\x34\x39\x38\x6c\xae\xfd\x9d\x43\xff\x85\xd1\x17\x6e\xb9\x3c\x3d\xff\x63\xfa\x8e\x1b\xfd\xe5\x8a\x8f\xe9\x5c\xff\x3f\x9d\xce\xaf\x56\xd3\x0b\x7e\xcc\xa5\xc3\xbb\x7a\xf5\x7f\xa7\x48\xce\x17\x6e\x2c\x57\x8b\xe9\xec\x5c\x9e\xad\xa6\x97\xee\x70\x72\x3c\x99\xf4\x57\xb3\x36\x30\x8f\xf8\xc1\xd5\xdd\x11\x5f\x82\xff\xf9\x5f\xa6\x7d\x15\xcc\xea\x4b\x17\x68\xb7\xe3\xe0\x68\xb8\xeb\x16\x8b\xf9\xe2\x38\xf8\x30\x39\x38\xff\x43\x7e\x05\x67\x5f\x05\x48\x75\xf2\x9d\xfb\x9d\x41\xdc\xe2\x48\xc4\xe6\xf7\xd7\x57\x7d\xef\x7f\x33\xec\xf1\xf1\xe4\x60\xda\x4b\x87\xcf\xbe\x0a\x66\xd3\x0b\x86\x38\x58\xb8\xd5\xd5\x62\xc6\xcf\x57\x81\x57\xe9\xe4\x1b\x46\xef\x8f\x0e\x19\x28\xf8\xfc\x6f\x67\xc1\xe7\xbf\x1d\xaa\x24\x32\x97\x1f\xe3\xcf\xc9\xe4\xe0\xb7\x7a\x11\x34\x57\x7d\xa0\xf3\xe8\x24\x93\x83\x5f\x54\x9c\xaf\x82\xe9\xfc\xe4\xf5\xfc\xdd\xf5\xd1\x17\xbe\xcd\x2b\x2f\x9b\xef\xd5\x5e\x7c\x33\x48\x7a\xf2\xfa\xc2\xef\xd3\x91\x57\xff\x89\xe4\x61\x18\x1d\x7f\xcb\x40\xbe\xa1\xca\x6d\x37\xbd\x58\x27\x5f\x23\xfa\xd1\xf1\x2b\x5a\x4c\xfc\xb3\xd5\xf5\x3b\x17\xd4\xcb\xa5\x5b\xb1\xe4\x57\xed\x8a\x51\x44\x3f\xdb\x0f\x3f\xcd\xac\x9f\x07\xc1\x7c\x79\xf2\x9f\x7e\x5b\xbf\xf5\x3f\x6e\xfa\xd9\x16\x0e\xf7\xd7\x46\x90\x3d\xf4\xff\x74\x1b\x27\x07\xcb\xe9\x1f\xf2\x7b\x3a\x5b\xe5\xe9\xe4\xe0\x12\x88\x0c\x6e\x06\xfd\x2f\xff\x53\x6e\xfe\xe4\x2d\x24\xc0\x4c\x4e\xf8\xc6\x3c\x62\x2a\x47\xfd\xf4\xee\x5c\xc7\xc1\x77\x7e\x8a\xa3\x63\x9b\x81\x39\x4d\xcb\x7e\x7a\xc2\xec\xbe\xf3\xf6\xbe\x3f\x7a\x71\x7c\x5f\x91\xe6\xe3\xae\x08\xba\xb3\x2b\xb2\xfa\xae\x6b\x92\x7f\x3c\x00\xaa\xdd\x37\x00\xca\xf9\x31\x6e\x14\xdd\x18\xc1\xb4\xdf\x3e\xc8\xb7\xcb\x37\xd3\x85\x1f\xa2\x99\xcf\x2f\xd6\x7b\xd7\x17\xcb\x7b\x34\xbf\x5e\xaa\xe2\x1e\x5f\xea\xd6\x7d\xf8\x73\xad\xb7\x99\x04\x56\xfe\x0b\x50\xf3\x46\x22\xc8\x9b\x35\xcc\xf2\x46\xae\x56\x71\x74\xf8\xf6\x7d\xd4\xbf\x7d\x5f\x36\x6f\xdf\x87\xa5\xbf\x42\xbb\xaa\xb7\xef\x73\xe7\xef\xdb\xbd\xde\xb7\xe9\xe2\xb7\xef\x53\xdf\xae\xf5\xf7\x5b\xff\x3b\xe6\xbb\xbf\x6a\xff\xdd\x85\x6b\xcf\x3b\x7d\xe6\x92\xb5\x7b\xb4\x6f\xfd\x58\x91\x9f\xcf\xdf\xaf\xfc\xf8\xce\x5f\x85\xbf\x7a\x7f\xa5\x99\x1f\xc7\x7f\x66\xbe\x4d\x19\xae\xc9\x61\x73\x73\x65\xc5\xdb\xf7\x89\xef\x9f\xf5\x2a\x43\xd4\xad\xb7\x3b\x1c\xf0\x68\x5c\x63\xf3\x97\x31\x1c\x1a\xbc\x6a\x0d\xc7\xbc\x03\x6e\x59\xb9\x57\xfe\xd1\xe1\xd6\x10\x7f\xe8\x1f\x1f\xdf\x98\xfb\xf8\x08\x08\xf1\x1f\xe2\xa9\xeb\x42\x88\xab\xde\xe0\xe1\x4e\x1d\xee\xc3\x9d\x1b\xb8\x10\x87\xf7\xa3\xdd\x31\x9e\x0f\xb8\xd5\x59\xb0\x4b\x8b\x00\xf7\x39\x0b\xc2\x57\x01\x6e\x70\xb6\xee\x25\x47\xe1\xb1\xdc\xc5\xb4\xcf\xd4\xf4\x7f\x9e\x4d\xdf\x1f\xf9\xb6\xe1\xb1\x07\xb3\x9a\x19\xbf\x10\xfd\x3e\x88\x52\x67\x81\xe9\x86\x38\x67\xf2\xff\x9f\x37\x8b\x5e\xbf\xda\x69\xad\x04\xa0\xbd\x6c\xb5\xf4\x76\x54\x45\x6a\x8b\xb5\x7f\xd6\x79\x9b\x4b\xfc\xb3\xc8\x5f\xa5\xb7\xb5\xbe\x50\xdb\x2b\x6b\x6d\x97\x63\xbf\x7e\xdc\x3c\xf7\x9f\xfe\x77\xec\x9f\xa7\xfe\x5e\x9c\xa9\xdd\xf2\xbd\x49\xbd\xed\xf1\xcc\xcf\x93\xfa\x2b\xc3\xce\x23\xb5\xf3\xd4\xb7\xc9\xbc\xad\x47\xbe\x5f\xeb\xaf\xdc\xdf\xeb\xb1\x77\xff\x59\xfa\x76\x19\xe3\x7b\xb9\x2a\xff\xbd\x89\x54\x9e\xce\xdf\x73\xcc\xe7\xe5\x6b\xfc\xdc\x4d\xa9\x9f\xad\xef\xd7\x47\xfa\x89\x9f\xe4\xbe\x5f\xea\xdb\x24\x5c\x5e\x86\x7e\xf0\x27\xdf\xbf\xad\x74\x9e\x3a\x57\x3f\xeb\xfc\xef\xca\xa9\x8e\xf8\x17\x3e\x85\xbc\xcc\x81\x5f\xa5\x7e\xde\xba\xd1\xe7\xa9\x1f\xab\x0d\x75\x3d\x23\xdf\xa6\x46\x4f\x3f\x4e\x8e\x8e\x9d\xae\x31\x72\xe2\x6b\xb5\x6f\x5f\x70\xa5\xda\x26\xaa\x74\x7e\xd6\x33\xf4\xf7\xea\x48\x65\x4b\x2a\xed\xc7\xfa\x71\x3f\xc9\x74\x5f\x22\x3f\x46\xc5\x1e\xe4\xba\x4e\x8c\x53\xa0\x7f\xa3\xf3\x81\x21\x4d\xad\xf2\x17\xbd\xca\xd2\xf8\xb6\x61\xa1\x6b\x47\xff\x12\xdd\x73\x1d\x97\x3d\x62\x8d\x43\xdf\x3f\xe9\x55\x26\x87\x0e\x89\xee\x51\xe5\xe7\xa8\x0c\x6f\x72\xf6\x3b\xb6\xfd\x88\xf5\xea\x4c\x1e\xee\x95\x95\xda\x48\xe6\x7f\x47\xb5\xae\x47\x5e\xab\x8d\x84\x9d\xb6\xed\x18\x23\xd3\xfd\x64\xaf\xab\xdc\x6c\xa5\x57\x1b\xc9\x58\x03\xdb\xff\x30\x57\xdd\x9a\x50\x75\x43\xee\xb8\xd7\x31\xd0\x89\x3d\x09\x9d\xf6\x45\xf6\x8c\xbd\xc0\x66\x06\xf9\x13\xdd\xcf\x12\x1b\x8c\x6c\x6f\x72\xc5\x46\x6c\x14\x7b\xed\x4c\x36\xee\x61\x97\xac\x4f\xef\x74\xaf\xeb\x4e\x31\x15\x9b\xc6\x5f\xd8\x37\xec\x95\x67\xb9\xbf\xdf\x95\xba\x4f\x45\xac\x3e\x80\xbd\x96\x89\xce\x85\x1c\x3c\x63\x7f\x53\x7f\x25\xad\xda\x15\xeb\x5b\xf6\x6a\x8f\x3c\xc7\x3e\xe9\x8b\x4f\xb1\xbf\xd8\x0b\xfa\x74\xec\x6b\xa4\x76\x91\x21\x73\xa5\x7b\x4e\x7b\xc6\xcf\xcd\x6f\xf2\x56\xd7\x97\x35\x45\x1f\x7c\x84\x7d\x27\x0e\xb8\x4c\xd7\x0f\x9f\x8b\x6c\x8f\x92\x5a\x75\x65\xef\xaa\x54\x7d\x83\x38\x80\x4f\xb0\x7e\xec\x19\xbe\x84\x3d\xc5\x4e\xfd\x1e\x4c\x70\x66\xcf\x99\xad\x0b\x7b\xe4\x3a\xf5\x43\x64\x21\x9e\xe0\x43\xec\x0f\xba\xe2\x7f\x79\x61\x36\x8f\x1d\x86\x6a\x27\xb5\xd9\xb2\x3c\x63\xbd\x73\xd5\x87\xbe\xd8\x8f\xeb\x75\x5c\x64\x2a\x9d\xda\x69\x56\xab\xdf\x12\x03\xb1\xd9\xc6\xf7\xad\xcc\xe7\xc5\xde\xf0\xd7\x5a\xf7\xb2\xa9\xd4\x4e\xb9\x5f\x17\x86\x4f\x8d\xfa\x41\x6f\x31\x92\xf5\x61\xed\xcb\x48\xf7\xc2\x45\xea\xc3\xd8\x61\x83\x9f\x96\x6a\x03\xf2\x9c\xfd\xec\x55\x66\x64\x07\x0f\x59\x63\xb1\x69\xfc\x3d\x56\x7d\xc1\x4a\xd6\x1f\xdc\x64\xef\x58\x7b\x74\xe9\x53\x8d\xed\x7d\xa2\x78\x82\x0d\xa1\x13\xeb\xc4\x1c\x61\xb6\x19\x9f\xe3\x58\xfb\xc8\x9a\x63\xeb\x99\xf9\xdb\xee\xf8\x0c\xc6\x3f\x3e\x3a\x33\xca\x46\x6c\xbe\x7d\xb4\x3b\x30\xd3\x62\x9f\xb0\xbc\x26\xfa\x73\x04\xe5\x75\xf1\x5f\x4a\x44\x7e\xad\x69\xea\x7f\x5f\x5e\xec\x15\x97\xc1\x7d\xec\xb0\x05\xf3\xbd\xed\xb7\xf1\x6d\x5c\x4e\x2d\x2e\xf7\x9d\xc6\x65\xfc\x9e\xf8\x84\x4d\x31\x36\x38\x52\x0e\xbe\x54\x2b\xc6\x4b\x6c\x6f\x15\x53\xa3\x46\x39\x22\xf7\xc1\x44\x62\x1d\x32\x80\x9d\xe0\x16\xf7\xc1\xee\xbc\xd1\x39\xf0\x2f\xf0\x25\xb7\xb8\x8b\x0c\x8c\x05\x76\x34\xe6\x2b\x85\xf9\x2b\xf6\x2e\x71\x2f\x33\x6e\x51\x69\x0c\x42\x0e\x70\x0e\x1c\xc1\x4f\xf0\xf7\xde\x30\x03\x7c\x26\x26\x31\x8f\xdc\x4b\x8d\x1f\xe4\xea\x43\xe0\x30\x3e\x22\x18\x46\xdb\x5a\xf1\x1d\x3e\x21\x58\xdf\x6b\x2c\xc0\xcf\xd1\x47\xb8\x73\xa1\x78\x81\xbe\xe0\x4f\x62\x38\x00\x1e\x12\x2b\xc3\x56\xf1\xa9\x36\xde\x02\x6e\xa0\x57\x65\xf1\x88\x3e\xb2\x46\x91\xae\x69\x63\xfe\x4f\x3b\x64\xc8\x2c\xde\xc0\x6b\x3a\xc3\x1f\x70\x87\x7d\x6c\x62\xd5\x75\x88\xe3\x60\x32\x6b\x93\x18\xbf\x02\xcf\x68\x1b\xb1\xf6\xfe\x59\x58\xeb\x38\xc4\x4a\xd1\xbd\x55\xfc\x72\x4e\xf7\x97\xb5\x84\xc7\x54\xa5\x62\x27\x18\x83\x0e\x12\x7b\x2b\x8d\x0d\xe8\x47\x1c\x03\xc3\xc0\x75\x62\x02\x76\x01\xfe\xc2\xff\xf3\x54\xb1\x33\xb6\x58\x10\x46\x23\xb8\x94\xa9\x1c\xd8\x19\xeb\x0e\xa6\x3d\x20\x6f\xb8\xb5\xf4\xc7\xa3\xd3\xed\x58\x1b\x18\xb5\x59\xfe\xd9\x8d\x55\xb7\x43\xed\x83\x58\x1b\x4a\x3d\x07\x6e\x8d\xa9\xf4\x57\xe2\xd7\x0f\x14\x33\x9f\x24\xfd\xc5\xf5\xa0\x43\x71\xad\x30\x21\xa9\xad\x51\x25\x42\xe7\x7a\x9b\xb1\x34\x99\xb1\x92\x58\x29\x30\x66\xd8\xd6\x4a\xcb\x31\xe3\xa2\xd6\x50\xce\xdc\x15\xd0\xe3\xd4\x8d\x04\xce\x9c\xba\x51\xe6\x94\x5e\x41\x69\x80\x1a\xda\x33\x37\x50\x0a\x04\x21\x97\x84\xfd\x42\xe7\xc5\xbd\x81\x0e\xa8\x9f\xb8\x48\x64\xb4\xd2\x68\x33\x34\x5d\x28\xe8\x40\x5d\x1a\x7d\x46\xbf\xd8\x28\x8b\xa4\xe3\x06\x8f\x77\x5d\xab\x36\xfa\x5f\x65\x0a\x0d\xc8\xb4\xc5\xb5\x36\x36\x61\x3f\xaf\xda\x18\xe6\xd6\xa1\x46\x8a\xdf\x9b\xae\xb4\xd1\xff\xa1\x5e\xb4\x4d\xfe\x27\x75\xa0\x51\x15\x5e\x88\xeb\xec\x9d\x8b\xc3\xf5\xc9\xe1\x92\x4e\x73\xd9\x72\x2d\x17\xef\x3a\x8d\xc5\x18\x2b\xc6\x43\x0c\x8c\x8c\x3f\x13\xdb\x31\x2c\x38\xb2\xc4\xbb\x4e\xf3\xd6\xce\xf2\x4d\xe2\x0d\x31\xa5\xb0\x7c\x49\x78\x7f\xae\x1c\x9d\xf8\x11\x27\x1a\x83\x90\x01\x07\x2a\x2d\x67\x8a\x2d\xf7\xc5\x89\x1c\x39\x50\x67\x71\xbe\x57\x19\x71\xb2\x21\x0e\x27\x8d\x3a\x4f\x61\xb9\x1d\xf1\x3e\x31\xee\x4f\xae\x18\x59\x2e\x5d\x90\x5b\x91\x5b\xb4\xea\x54\xe4\xc6\xc8\x21\x75\x84\x52\x9d\x4c\xf2\xe9\x5a\x73\x11\x38\x84\xb3\x35\x60\x0e\x67\x35\x06\xe2\x3d\x0e\x87\x4e\x45\xae\x79\x0a\xbc\x88\xf5\x42\xbf\x38\xd7\x1c\x8e\x7c\x8b\x3a\x40\x63\xf9\x22\xb9\x37\x79\x07\x7c\x43\xc0\xc5\xea\x64\xd4\x18\x88\xab\xc4\x64\xe4\x45\xf7\xc8\xf2\xd8\xdc\xd6\x8f\xf9\xc9\x13\x24\xaf\x4d\x6f\xf9\x46\xd1\x5a\x4e\x9d\x69\x5f\x40\x8a\xbc\x83\xfc\x04\xfe\x03\xd7\x90\x1a\x46\xa2\x9c\xa1\xcd\x2c\x67\x88\x14\xdc\x58\x1b\x72\x67\xd6\x96\x5c\x1c\x59\x01\x8c\x64\xc8\x45\x52\xdd\x17\x72\x44\x38\x00\xba\x32\x67\x63\xf1\xbe\x34\xfd\xe1\x42\x92\x03\x3b\xe5\x27\x99\xe5\x69\xf4\x67\xff\xd8\x93\xdc\xf6\x94\x35\x81\x37\x00\x92\xf0\x27\xe1\x23\x9d\x72\x86\xaa\xd1\x9a\x01\x40\x8b\x0c\xec\x03\xf6\x22\xf9\x6b\xa5\xfc\x81\xfe\xac\x7f\xd3\x2a\xa7\xc1\x76\x22\xe3\x24\x00\xbc\xd4\x50\x1a\x5d\xa3\xc8\x78\x26\x32\x00\xc0\xd4\x18\x52\xcb\xf3\x6b\xab\x0d\x90\xef\x53\x1b\x40\x66\xd6\xb8\xea\x35\x28\x74\xd8\x77\xa9\x36\x5d\x5b\xed\x87\x39\x92\x44\xf5\x24\x87\x23\xaf\xed\x8c\x3b\xe6\x9d\xe6\xdb\xc8\x45\x0e\x0b\x0f\x96\xda\x0d\x7c\xa7\xd7\x40\x42\x7e\x8b\x6f\xc1\xc9\xa4\x7e\x63\x39\x3a\x7b\x4f\x2d\x08\xf9\x08\x4a\x95\xe9\xd3\x0e\x7e\x64\x7c\x8a\xb5\x1a\xe3\x4f\x89\xd5\x13\xe0\xaa\x6e\xa8\xa3\x44\xf7\x81\xfc\xfe\x69\xdd\x9d\x41\xee\x02\xfc\xae\xa4\xee\x4e\xd7\x3d\xb0\xfd\xb9\x52\xba\x4d\xd9\x5f\x08\xac\x3f\x32\xa1\x1b\x18\x07\x86\x4c\x82\xe0\x8c\xdd\x00\xf8\x52\x74\x32\xa6\x44\x82\x02\x53\xa1\xf8\xc9\x73\x80\x03\x10\x82\x0d\x01\x76\x38\x13\x89\x45\x66\xbf\x49\x10\x00\x59\x49\xf8\x2c\x99\x68\xec\xb3\x1d\x0a\x22\x56\x34\x91\x24\x27\x53\xc7\x0e\xb3\xdb\x03\x85\xd8\xee\x01\x9a\x5c\x38\xfb\xd0\x26\xb5\x76\xb1\x7d\x0e\x63\x92\xa0\xe0\xb4\xdc\xe7\x3b\xa0\x86\x73\x56\x56\xe4\xe3\x02\xc4\x61\x63\x91\x31\x1f\x9c\x9d\x42\x47\x34\x00\x84\x81\x68\x69\x05\x30\xc0\x0b\xb0\xa9\xec\xfb\x70\x01\xf8\xe8\x20\xc5\xe5\x4c\x83\x06\xce\x1f\x5b\xa2\x17\x8f\x30\x2e\x9c\x9e\xb1\x29\xb6\x10\xb0\x28\x2e\xde\xcf\xb8\x1e\x9b\xcb\x8c\x0e\x75\xd7\x31\x1f\x92\xc9\x8c\x0e\xb4\x87\x9b\x3e\x6f\x1e\xb3\x5d\x9f\xbf\xd4\x69\xaf\x9a\xeb\x7f\xa8\x2c\x66\xab\x0d\x57\x5a\x74\x21\xb3\x29\xac\xb8\xbf\xcd\x86\xef\xe8\xbc\xa7\xf9\xde\x19\x65\xcd\x72\x37\x5e\x5e\x19\xb1\xd9\x3b\xbd\x1f\x6c\xae\xe3\xb2\x3f\xad\xa5\x8e\xc8\xff\x32\x6c\x74\xef\x74\x41\x8e\x80\x7b\xb3\x4c\xa7\x88\x37\x1c\x35\x81\x76\xce\x4a\xf5\x20\x73\x6c\xc7\x67\x58\x26\xe5\x3e\x29\x25\x95\x6a\xc9\xd0\xdb\xaa\x56\x6a\xde\xdb\x71\x04\xd1\x03\xab\x86\x4e\x49\x44\x69\x74\xcc\xca\x8e\x90\x18\x2b\xb3\xf2\x14\x74\x8d\x68\x05\xf5\x87\xda\x43\xe7\xc9\x8b\x73\x43\x76\x68\x30\x54\xbc\x30\x3a\x84\xbc\x94\x37\xa5\xd4\x57\x28\xc5\xcf\xed\x58\x3b\xb5\x72\x3b\x74\x51\x8e\xb0\x53\x2d\x6b\x15\x16\x35\xa0\x66\xa4\x0a\x44\x3b\x52\x22\x68\xb9\x44\x2d\x8b\x26\xd0\x3a\x8e\x95\xe4\x28\xb2\xd0\x68\x82\xfc\x1c\xd3\x90\xd6\x40\x99\xa1\xc6\x78\x22\x17\x25\x31\xa2\x9d\x94\x0b\x2d\x2d\x91\xd4\x26\xd6\xf2\x9b\x1c\x7b\x45\xea\xb1\x49\xa9\x75\x86\xda\xe4\x17\xda\x6e\x47\x81\xd0\x74\x39\xf6\x0c\x6d\xcc\x50\xf5\x95\x79\xed\x78\x9f\x35\xe4\xa8\x03\xca\x0d\x85\x66\x5f\xe4\x08\x10\xea\x5a\x6b\x1d\x01\x0f\xe7\x98\x89\x7a\x02\x74\x95\x3d\x49\x8d\x72\xe6\xf9\x6d\x29\x13\xf9\x4a\x8b\x84\xd4\x2b\x9c\x1d\x99\xd1\x9f\x35\x71\x56\xba\xe5\x9e\xb3\xfa\x08\xf4\x99\x52\xac\xbc\x3e\xd0\x28\x02\xc9\x11\x8a\x53\x84\x63\xef\xb8\xc7\x5c\xbd\x1d\x53\x09\x12\x51\x8a\x4d\x2c\x5d\x2c\x6e\xd3\x00\x58\x09\xb4\x9b\xfa\x0b\xe9\xaa\x1c\x15\xda\x2b\x0e\xd0\x65\xa9\x8d\x24\x9a\xba\x84\xc9\x26\xba\x61\xe3\x72\x8c\x33\x44\xf3\x7a\x3b\x5d\xfe\xc8\x5b\x1e\x8b\x6d\x77\xc8\xf2\xc7\x6f\xdc\xed\x82\xb5\x4f\xa2\xca\x63\x22\x3f\x3d\xa4\xbd\x18\xa2\x7c\xa3\xee\x53\x1c\x7c\xb4\x7a\xf0\x81\xe7\x0e\x45\x10\x92\x68\xd0\xc4\x59\x01\x5b\x78\x69\x68\x87\x97\x96\x18\xe3\xe1\x70\xc2\xbe\xd4\x3e\x78\x5b\x68\x71\x97\x64\x10\xab\x05\x89\xa4\x20\xdf\xa9\xf5\xe2\x3d\x78\x12\x48\x87\xd7\x49\xa1\xa5\xd6\xe4\x58\x2c\x35\xd6\xe4\x1a\xef\xc7\x23\x29\x8c\xb4\x96\xa8\x53\x3d\x94\x39\xec\x20\x5b\x0e\xdd\x5b\x4d\x3e\x91\x93\xc4\x58\x50\xd0\x0e\xb8\x41\x39\xe2\xb9\x24\xdd\x96\x70\x92\x50\x72\xc8\x4d\x6e\x20\x07\xb6\x4e\x79\x32\x87\x94\xa1\x1d\x36\xd0\x96\xef\x3c\x23\x7f\xe0\x60\x82\x62\x8b\x1c\xdc\x67\x76\x60\xe3\x54\x47\x39\xb0\x76\x2a\x2b\x5e\x0e\xba\x93\x43\x50\xf8\xc0\x1b\x87\x22\x45\x68\x2f\x02\x50\x74\x42\x37\xa9\x98\xb6\x86\x5e\xa5\x8e\x85\x57\xca\x81\x6d\xaf\x45\x15\x39\xa8\x68\xf4\xf0\x41\x0e\xda\x23\x5d\xd3\xde\x64\xa2\x7d\x62\x49\xbe\x20\xa6\x25\xce\xbd\x1d\xba\x83\x30\xa1\x1d\xe6\x53\x94\xe8\xac\x90\x10\xda\x81\x79\x66\x45\x11\x92\xfe\xa6\x1d\x47\x8d\xda\x0e\x2f\xe4\x00\x29\xb3\x75\xba\x8f\x13\x3d\x9a\xd6\x8f\x8c\x74\x07\x41\x1e\x44\xea\x47\x86\xf9\x74\x3c\x79\x66\x4a\xbf\x4d\x99\xbf\x12\x5d\x7e\xd4\x97\x8f\xff\xd5\x48\xfd\x88\xda\xfb\xd9\xef\xc8\x40\xb7\xe6\x3b\xfa\x9a\xf8\xa6\xf1\x8e\x8c\xf1\x50\xdb\xdd\xae\xc7\x93\x9a\xee\x16\x45\x5e\x8c\xe1\x3e\x9e\xe9\x47\xf6\x2c\x5d\x63\xfa\xf9\x1d\xa6\x9f\xea\xa1\xfc\xc0\xf4\xc1\x60\x62\x1e\xf1\x45\x0a\x9f\x9d\x9a\x3a\x98\xdb\xd8\x77\x62\x13\xa6\x58\xda\x8b\x4a\xa9\xe1\x6e\x3b\xc4\xc2\xdc\xd8\x9a\x15\x87\x6b\x63\x98\x52\xa7\x72\x5a\x0f\x42\x0e\x64\x75\xf6\x52\x1f\xb1\x05\x06\x48\xcd\x46\x8a\xb5\x56\xb4\x96\x83\x8a\x4a\x0f\xd3\xa5\x26\x66\xb1\x4d\x0a\xf1\x76\x4a\xd7\xd9\xcb\xad\xad\xbd\xd4\x97\x59\x51\x5f\x0e\x06\x72\x7b\x19\xcb\x0e\x02\x64\x9d\x72\x8d\xcb\xb0\xfd\xe1\xc5\x02\x39\xb0\x0f\x75\x5c\xfa\x51\xf4\x27\x8b\x20\x66\x49\xf6\x52\x28\xd3\x8d\x2d\x23\xe1\x93\x35\x45\x3e\x79\xf1\xa8\xd7\xba\x5c\x69\x07\xf9\x9d\x1d\xd6\x77\xf6\x12\x58\x51\x68\x76\x42\x3c\x96\x53\x3f\xe3\x12\x89\xb9\x34\xeb\xc5\x5c\x70\x93\xcc\x0e\x65\x88\xb3\xb0\x71\xda\xb4\xc3\x0b\x6f\xf6\x82\x41\x6f\x4c\xdc\x19\xf3\x97\x97\x18\x23\xcd\x48\xe8\x2f\x31\xdc\x29\x7b\xcf\xed\xe5\xae\xd0\x5e\x42\x63\x4f\x13\x7b\xd9\x51\x8a\xf6\x4e\xef\x21\x13\xed\xa5\xa0\x5e\xdb\x0b\x6e\xb5\xb1\xf5\x5c\xf7\x25\x8a\x94\xdd\x3b\x63\xfa\xf2\xd2\x47\x7f\xfb\xb2\x1b\xf7\x90\x99\x3d\xa2\xde\x48\xec\x95\xac\xa2\xd3\x35\xe6\x99\xbc\x34\x99\x68\x86\x53\xd8\x0b\x70\xd4\x02\xc7\x5e\x7e\x82\xb3\xb0\x2e\x83\xed\x0c\x05\xf8\xfb\xe0\x6f\x7f\xe2\xbf\x31\xcc\x26\xf4\xed\xa2\xff\x1b\xdd\xf7\x42\xbd\xe7\x4a\x02\xc6\x34\x78\x31\x80\xf7\xef\x92\xf9\x3f\x6b\xc9\x7c\xcb\x36\x3f\x81\x83\x8e\x11\xec\xed\x7f\x7b\x76\x8f\xbb\x7e\x3a\xcd\xde\xad\xd8\xb3\xb8\xee\x4b\x23\xdb\x3f\xeb\xdf\xdb\x3d\xdd\x9f\xc1\xec\xf8\x33\x17\xcc\x2d\xb4\x57\x15\xc7\x22\x05\x0c\x27\xb6\x57\x03\xc5\x55\xc7\xcd\x71\x44\xe4\xfd\x4c\x71\x64\xa0\x5b\x33\x1c\xfd\xab\xc6\x4d\x0b\x1c\x19\xe3\xa1\xd6\xb7\x5d\x8f\x27\xb5\xbc\x2d\x8a\xbc\x18\xa3\xdb\x9b\x28\x43\xf0\x20\x63\xf2\x57\x0c\x85\xbe\x19\x10\xdb\xdb\x03\x43\xb4\xc0\x14\x63\x3b\xa5\x27\x42\x40\x90\xe5\xaf\x5e\xac\xd0\xe4\x86\xb7\x5d\x62\x2b\xe8\xf4\x3a\x87\xa0\x73\xa8\x6f\x91\xf0\x99\x9a\x09\x33\x3e\x25\xcf\x6d\x66\x0c\x01\xe5\xaf\x15\x24\x9a\xb4\xaa\xcb\x43\xcc\x78\x7f\xc2\xb3\x31\xcc\xa6\x09\xef\x22\x3c\x1b\xdd\xf7\xb2\xde\xe7\x22\x3c\x63\x1a\xbc\x18\xc3\x7d\x24\xe1\xa1\x94\x90\x1a\xa1\x69\xed\x45\xa0\x4f\x21\x3d\x20\xec\x36\x33\x84\x48\xc8\x0b\xd0\x91\xd6\x3b\xa3\xe8\x61\x66\x78\x5f\x70\xff\x7b\x00\x00\x00\xff\xff\x37\xc7\x62\x8d\x00\x40\x00\x00")

func templatesGoBytes() ([]byte, error) {
	return bindataRead(
		_templatesGo,
		"templates.go",
	)
}

func templatesGo() (*asset, error) {
	bytes, err := templatesGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates.go", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"init/django/.dockerignore": initDjangoDockerignore,
	"init/django/Dockerfile": initDjangoDockerfile,
	"init/django/docker-compose.yml": initDjangoDockerComposeYml,
	"init/rails/.dockerignore": initRailsDockerignore,
	"init/rails/Dockerfile": initRailsDockerfile,
	"init/rails/docker-compose.yml": initRailsDockerComposeYml,
	"init/ruby/.dockerignore": initRubyDockerignore,
	"init/ruby/Dockerfile": initRubyDockerfile,
	"init/ruby/docker-compose.yml": initRubyDockerComposeYml,
	"init/sinatra/.dockerignore": initSinatraDockerignore,
	"init/sinatra/Dockerfile": initSinatraDockerfile,
	"init/sinatra/docker-compose.yml": initSinatraDockerComposeYml,
	"init/unknown/.dockerignore": initUnknownDockerignore,
	"init/unknown/Dockerfile": initUnknownDockerfile,
	"init/unknown/docker-compose.yml": initUnknownDockerComposeYml,
	"templates.go": templatesGo,
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
	"init": &bintree{nil, map[string]*bintree{
		"django": &bintree{nil, map[string]*bintree{
			".dockerignore": &bintree{initDjangoDockerignore, map[string]*bintree{}},
			"Dockerfile": &bintree{initDjangoDockerfile, map[string]*bintree{}},
			"docker-compose.yml": &bintree{initDjangoDockerComposeYml, map[string]*bintree{}},
		}},
		"rails": &bintree{nil, map[string]*bintree{
			".dockerignore": &bintree{initRailsDockerignore, map[string]*bintree{}},
			"Dockerfile": &bintree{initRailsDockerfile, map[string]*bintree{}},
			"docker-compose.yml": &bintree{initRailsDockerComposeYml, map[string]*bintree{}},
		}},
		"ruby": &bintree{nil, map[string]*bintree{
			".dockerignore": &bintree{initRubyDockerignore, map[string]*bintree{}},
			"Dockerfile": &bintree{initRubyDockerfile, map[string]*bintree{}},
			"docker-compose.yml": &bintree{initRubyDockerComposeYml, map[string]*bintree{}},
		}},
		"sinatra": &bintree{nil, map[string]*bintree{
			".dockerignore": &bintree{initSinatraDockerignore, map[string]*bintree{}},
			"Dockerfile": &bintree{initSinatraDockerfile, map[string]*bintree{}},
			"docker-compose.yml": &bintree{initSinatraDockerComposeYml, map[string]*bintree{}},
		}},
		"unknown": &bintree{nil, map[string]*bintree{
			".dockerignore": &bintree{initUnknownDockerignore, map[string]*bintree{}},
			"Dockerfile": &bintree{initUnknownDockerfile, map[string]*bintree{}},
			"docker-compose.yml": &bintree{initUnknownDockerComposeYml, map[string]*bintree{}},
		}},
	}},
	"templates.go": &bintree{templatesGo, map[string]*bintree{}},
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
