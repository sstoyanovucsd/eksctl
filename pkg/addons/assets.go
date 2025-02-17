// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/efa-device-plugin.yaml (3.084kB)
// assets/neuron-device-plugin.yaml (3.618kB)
// assets/nvidia-device-plugin.yaml (2.361kB)
// assets/vpc-admission-webhook-config.yaml (524B)
// assets/vpc-admission-webhook-csr.yaml (234B)
// assets/vpc-admission-webhook-dep.yaml (1.675kB)
// assets/vpc-admission-webhook.yaml (231B)
// assets/vpc-resource-controller-dep.yaml (1.679kB)
// assets/vpc-resource-controller.yaml (565B)

package addons

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _efaDevicePluginYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x56\x4d\x8f\xdb\x36\x10\xbd\xfb\x57\x0c\x1c\xec\xad\x92\xba\xd9\x5d\x20\x50\x4f\xc6\xc6\x2d\x82\x6e\xb6\x41\x9d\xb4\x87\xa2\x87\x31\x39\xb2\x08\x53\x24\xcb\x19\xd9\xeb\xfe\xfa\x82\xb2\xac\x58\xce\x3a\xeb\x5e\x72\x8a\x2f\x16\xc8\xf7\xe6\xcd\x0c\x1f\x3f\xb2\x2c\x9b\x60\x30\x7f\x50\x64\xe3\x5d\x09\x18\x02\x17\x9b\xeb\xc9\xda\x38\x5d\xc2\x5b\xa4\xc6\xbb\x05\xc9\xa4\x21\x41\x8d\x82\xe5\x04\xc0\x61\x43\x25\xe0\x96\x33\xaa\x30\x5b\xbf\xe1\x4c\xd3\xc6\x28\xca\x82\x6d\x57\xc6\x65\xba\x63\x31\x49\x8f\xe5\x80\x8a\x4a\x58\xb7\x4b\xca\x78\xc7\x42\xcd\x84\x03\xa9\x14\x8a\xc9\x92\x12\x1f\xd3\x37\x40\x83\xa2\xea\x07\x5c\x92\xe5\xfd\xc0\x41\xeb\xbc\xd8\x04\xa0\x0d\x1a\x85\x16\x12\x51\x68\xb5\xdb\x13\x65\x17\xa8\x84\xdf\xbd\xb5\xc6\xad\x3e\x75\x80\x09\x80\x50\x13\x2c\x0a\xf5\x6a\x47\x25\xa5\xdf\x2b\xf8\x58\x1b\x06\x74\xce\x0b\x8a\xf1\x0e\x0c\x83\xa6\x10\x49\xa1\x90\xce\xe1\x57\x0a\x02\x35\x45\x82\xca\x47\x58\xa2\x5a\x6f\x31\x6a\x50\xbe\x09\x28\x66\x69\xac\x91\xdd\x10\x6b\x41\x04\xb5\x48\xe0\xb2\x28\x52\xe5\xd1\x91\x10\xe7\xc6\x17\xda\x2b\x2e\x04\x79\xcd\x05\xea\xc6\x38\xc3\x42\x31\x53\xb6\x4d\xff\xc5\xaa\xc5\x88\x4e\x88\x74\xc6\xaa\x26\xdd\xa6\x0a\x32\x15\x8d\x18\x85\x36\x43\xad\xbd\xcb\x82\xd7\x5c\xf4\x52\x9f\xf3\x1d\x9a\x06\xd0\x53\x29\xe6\x68\x43\x8d\xf9\x38\x83\x21\x5a\xf0\xba\x84\xe9\xb4\xa7\xd9\x51\xe7\x5f\x5c\xe7\x0e\x77\x58\xc8\xee\x9b\x62\x9a\x9e\x29\xe5\x5b\x27\x25\x68\xaa\xb0\xb5\xd2\xcf\x8a\xb7\x14\x4f\xf3\xcc\x60\x4d\xbb\x12\xee\xfb\x84\x66\xa9\x3a\xfe\xcd\xd9\xdd\x80\x00\xf0\x21\xf1\x7c\x2c\x61\xfe\x64\x58\xf8\x94\x8c\x5b\xce\xb1\xc1\x7f\xbd\xcb\x95\x6f\x0a\xaa\xf0\x12\x32\x00\x55\x15\x29\x29\xe1\xd1\x2f\xfa\x6e\x0d\x8b\xf7\x1e\xe3\x1a\x24\xb9\x21\x78\x0d\xc8\x80\x70\xe8\x19\xa0\xd6\x99\x77\x3f\xc1\xb6\x26\x07\xe4\x70\x69\x49\xff\x00\x52\xd3\x29\x64\x88\x36\x2c\x06\x44\x4a\x3d\x22\x4e\x1f\xbe\x8d\x8a\xb8\xb3\xd2\x09\x31\x89\x32\xb0\x07\xa9\x51\x52\xe4\x1d\x28\xfc\x1c\x6e\x49\x89\xde\xc7\xd4\x80\x95\x50\x04\x84\x0a\x8d\x6d\x23\xe5\xdf\xde\x81\x21\x1a\x1f\x8d\xec\xee\x2d\x32\x3f\x76\xa6\x99\xee\xf7\x79\xe6\xbc\xa6\x81\x7a\xf0\x19\x56\x95\x71\x46\x76\x47\x4e\xf3\x9a\x66\x5f\x8c\xa6\x22\xe6\x3f\xcf\x80\xdb\x10\x7c\x14\xd2\x60\x1c\x0b\x3a\x45\x5c\x0e\x95\xa5\x5a\xf2\x13\x0b\xcc\xfe\x5c\xcc\xef\x5f\x17\x69\xab\xb3\x14\x9f\x98\xe2\x2f\xad\xd1\x94\xac\x91\xd7\xd2\xd8\x57\xc9\xcf\x87\x58\x59\x3a\x2b\x8e\x7d\x11\xe9\x9f\xd6\x44\xd2\x6f\xdb\x68\xdc\x6a\x31\xb4\xe0\xdd\xca\xf9\x61\x78\xfe\x44\xaa\x4d\x66\x3e\x4e\x77\x5f\xc8\xa2\x3f\xd3\x3e\x52\x6c\x78\x3c\x9d\x3c\xdb\x1d\x72\xf3\xa7\x10\x89\x79\xbc\x19\x8e\x51\x9d\xb3\xa7\x4b\x92\xd3\xbd\x3b\x4a\x7b\xfa\x0c\xf7\xd8\xf3\xef\xdc\xb3\x80\x0d\xda\x96\x9e\x15\xde\x8b\xab\x3b\x97\x5f\xbf\x79\xb2\x18\x57\xf4\x55\x50\x3a\x42\xed\x59\xc4\xea\x56\xbf\x04\x31\x37\xe4\xf2\xd7\xb7\x2f\x48\x75\xa8\x17\x02\xb9\xea\xfa\xe5\x40\xcd\x9d\xbe\x40\xae\xb9\xbb\x00\x14\x6e\x2e\x09\x15\x2f\x12\x8c\x17\x09\xde\xea\x73\xa0\xff\x67\xaa\xe4\xd1\xef\xa6\xfa\x6e\xaa\x33\xa6\xaa\x3d\xcb\x23\xc9\xd6\xc7\x75\x09\x12\xdb\xc3\xb8\xf2\x4e\xd0\x38\x8a\xa3\xeb\xdb\x34\xb8\x4a\xe7\xfd\x15\xe7\x7a\x1d\x73\x52\x31\xbf\xe2\xfc\x8a\x0b\x4a\x97\xcb\xb9\xa7\x43\xb9\xf9\x31\xbf\xc9\x6f\x8e\x9d\x76\xc9\x63\x63\xff\x63\x52\x6d\x77\xdd\x78\x27\xf4\x24\x63\xc7\xa1\xb5\x7e\xfb\x21\x9a\x8d\xb1\xb4\xa2\x39\x2b\xb4\xdd\x9b\xa3\x84\x0a\x2d\x8f\x1b\xa1\x30\x60\xf7\x6a\x33\x5f\xfa\x56\x47\x1f\x4a\xf8\x6b\x3a\x7b\x78\x98\xfe\x7d\x34\xb7\xf1\xb6\x6d\xe8\x7d\x7a\xe0\x9c\x70\xb2\xbe\x84\x73\x69\xa7\x5f\x93\x78\x1f\x50\xea\x12\x8a\x0d\xc6\xc2\x9a\x65\x77\x35\x5b\x92\x62\xc4\x3b\xdc\x47\x7b\xb9\x51\xc7\xbf\xae\x92\x56\xaf\x13\x18\x29\x87\x8b\x24\xff\x0b\x00\x00\xff\xff\xf5\xa9\x92\x60\x0c\x0c\x00\x00")

func efaDevicePluginYamlBytes() ([]byte, error) {
	return bindataRead(
		_efaDevicePluginYaml,
		"efa-device-plugin.yaml",
	)
}

func efaDevicePluginYaml() (*asset, error) {
	bytes, err := efaDevicePluginYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "efa-device-plugin.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf, 0x88, 0xf3, 0x33, 0x57, 0xdb, 0x84, 0xb8, 0x9e, 0x87, 0x8d, 0x16, 0xec, 0x91, 0x23, 0x4, 0xcc, 0x1d, 0x9f, 0x58, 0x2a, 0xa8, 0xee, 0x6b, 0xa7, 0xe9, 0xd6, 0x14, 0xbf, 0x16, 0x57, 0xeb}}
	return a, nil
}

var _neuronDevicePluginYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x56\xc1\x8e\xdb\x36\x10\xbd\xfb\x2b\x06\x0e\x72\xab\xa4\x6e\x5a\x14\x81\x7a\xda\x6c\x9c\x20\x68\xb2\x59\xc4\x4d\x2f\x45\x51\x8c\xc9\x91\xcc\x9a\x22\x59\x72\xe8\xb5\xf3\xf5\x05\x25\x59\x2b\xed\x2a\xee\x6e\xd1\x9e\xaa\x93\x64\xce\xbc\x19\x3e\xbe\x79\x74\x96\x65\x8b\x9d\x32\xb2\x84\x2b\x1d\x03\x93\xff\x64\x35\x2d\xd0\xa9\x5f\xc8\x07\x65\x4d\x09\x7e\x83\x22\xc7\xc8\x5b\xeb\xd5\x17\x64\x65\x4d\xbe\x7b\x19\x72\x65\x8b\xfd\xc5\xa2\x21\x46\x89\x8c\xe5\x02\xc0\x60\x43\x25\x18\x8a\xde\x9a\x4c\xd2\x5e\x09\xca\x9c\x8e\xb5\x32\x0b\x1f\x35\x85\x14\x93\x01\x3a\xf5\xd6\xdb\xe8\xda\xcf\xf4\x64\xb0\x5c\xb6\xaf\x9e\x82\x8d\x5e\xd0\x68\xc5\x58\x49\xa1\xfd\xda\x93\xdf\x8c\x16\x6a\xe2\xe1\x5d\xab\x70\xf7\x71\x8b\x2c\xb6\x4f\x2f\x44\x7b\x32\x3c\x5b\x49\x78\x42\xa6\xe1\xd3\xfd\x33\x7c\x67\xe5\x2c\x7a\x74\x72\x0e\xfd\xbf\xd9\x64\xcb\x66\x11\x18\x39\xce\x36\x33\xad\xde\xb7\x96\x04\x32\x96\xc3\xfe\xa2\xd7\xcb\x9a\x7c\x3a\xe2\x4b\x21\x6c\x34\xfc\x58\x25\x74\x8b\xc1\xa1\xa0\x12\x76\x71\x43\x59\x38\x06\xa6\x66\x31\x2b\xc4\x57\xca\x48\x65\xea\x7f\x5d\x8f\x5f\xed\xc2\x5b\x4d\x9f\xa8\x4a\xe9\x27\x6a\xcf\x54\x5c\x00\x3c\x9c\x9d\xb3\x85\x43\xdc\xfc\x41\x82\xfb\x59\x98\x25\x32\xb1\x7f\xb6\xf5\xb3\x14\x3e\x83\x2d\xb3\x0b\x65\x51\xa4\x15\x6f\x88\xa9\xe5\x46\x5a\x11\x0a\x61\x8d\x20\xc7\xa1\xa0\x03\x93\x91\xd9\x5d\x48\x21\x6c\xe3\x22\x53\x16\xd8\x7a\xac\x29\x33\xc4\xc5\xa4\x70\x28\x26\xa7\x80\xce\x85\x62\xd0\xc2\x6b\xa4\xc6\x9a\x35\x3d\x56\x06\x99\x6c\x13\x42\x2b\xf0\xaf\xec\x26\x38\x12\x09\x25\x90\x26\xc1\xd6\x77\x3a\x6d\x92\x46\xdf\xe3\x86\xf4\x20\xdc\x73\x65\x92\xce\x3b\x21\xaf\xd9\x23\x53\x7d\xec\xb2\xf8\xe8\xa8\x84\x4f\x56\x6b\x65\xea\xcf\xa7\x21\x64\x6a\x9c\x46\xa6\xbe\xd4\x68\x2b\xe9\x41\x63\x2c\xb7\xc7\x3f\x94\x06\x08\x62\x4b\x32\x6a\xf2\x39\x6a\xb7\xc5\x7c\xca\xba\xf0\x8a\x95\x40\x9d\x39\x2b\xcb\xd3\x64\x02\xe8\x49\xff\x7f\xbf\x03\x80\x13\x19\xed\xfb\x44\x2f\x67\x64\x02\xc0\x56\x93\xbf\xdf\x72\x06\x3b\x3a\x96\x70\xd5\xf7\x76\x29\xa5\x35\xe1\xa3\xd1\xc7\x21\x02\xc0\xba\x94\x67\x7d\x09\xab\x83\x0a\xbd\x33\x8e\x92\xf1\x36\xe4\xd8\xe0\x17\x6b\x72\x61\x9b\xa2\x6b\xe1\x31\xf9\x00\x54\x55\x24\xb8\x84\x6b\xbb\xee\xb9\xeb\x17\x9f\xc1\x07\xf4\x3b\xe0\xad\x0a\xc9\x2f\x01\x03\x20\x9c\x18\x04\x94\x32\xb3\xe6\x47\xb8\xdd\x92\x01\x32\xb8\xd1\x24\xbf\x01\xde\xd2\xfd\x90\x01\x6d\x38\x9a\xe4\x85\xe4\xf7\x14\xee\x4c\x11\x2a\xeb\xef\x27\xb6\x26\x0d\xc1\x02\x6f\x91\x13\xf2\x11\x04\xde\xc1\x6d\x28\xa5\xf7\x98\x12\xb0\x62\xf2\x80\x50\xa1\xd2\xd1\x53\x3e\xc4\xad\x89\xce\xcd\x20\x63\xd8\x85\x02\x65\xa3\x8c\x4a\x9e\x91\x89\xce\x3b\x8a\x3a\xa2\x47\xc3\x44\x32\xeb\xab\x28\x53\x67\x83\x82\x30\x1d\x53\xd2\x51\x28\xfa\x52\xce\x2b\xeb\x15\x1f\xaf\x34\x86\x70\xdd\x4a\x68\xd9\xcd\x4e\x96\x9c\x7e\x48\x3d\xa9\x0e\xab\x4a\x19\xc5\xc7\x91\xee\xac\xa4\xcb\x07\xbf\xa6\xbb\xe3\xcf\xa8\x3c\xc9\xd7\xd1\x2b\x53\xaf\x87\x6e\xde\xd5\xc6\x0e\x3f\xaf\x0e\x24\x62\x92\xd6\x38\xb3\xc3\x5c\xf7\x23\xfb\x33\xf9\x26\x4c\x97\x93\x82\xda\x19\x5e\x1d\x9c\xa7\x10\xa6\xd2\x1c\x47\xb5\x3a\x5b\x6e\x88\xef\x0f\x95\x32\x81\xd1\x08\xca\xd2\x10\x2f\x67\x72\xc7\xf2\x7b\x67\x66\x03\xf6\xa8\x23\xcd\x16\xee\x8a\x2b\x53\x5d\xe4\x07\x8d\xbe\xa6\xf3\x31\x2f\x1e\x13\xf4\xc3\xa3\x90\xbe\x9f\x8d\x7a\x1a\x5d\x89\xfd\xff\x25\x5d\xc2\x1a\x46\x65\xc8\x4f\x8c\x4e\x35\x58\xa7\xb1\x78\x1e\x72\xb9\xf3\x39\x09\x9f\x3f\x0f\xf9\xf3\x50\xcc\x79\x66\x79\x91\x7f\x97\xbf\xc8\xbf\x1d\x73\xd4\x02\xdc\x44\xad\x6f\xac\x56\xe2\x58\xc2\xa5\xbe\xc5\xe3\xd8\xce\x3a\xef\xde\xbd\x0c\xd9\xac\x7f\x0b\xf6\x63\xef\x33\xfb\x29\x89\x59\x9f\xff\xd3\xe7\x57\xab\xab\x8f\xd7\x6f\xde\xbd\x5d\xcc\x30\x5f\x42\x41\x2c\x46\x6e\xd2\xbe\x6a\xe2\x5c\x58\x53\xcd\x02\x5e\x7f\x7c\xbd\xfa\xfd\xfa\xf2\xc3\x6a\x0e\xef\x8d\xb7\xcd\xc3\xc3\xac\x14\x69\xd9\xff\xf9\x99\x5d\xbb\x41\xde\x96\xed\x55\x94\x27\x9d\x25\xcb\x19\x85\x06\x12\xb1\xb5\x23\x6b\x98\x0e\x3c\x45\x41\xad\xed\xed\x8d\x57\x7b\xa5\xa9\xa6\x55\x10\xa8\xdb\x6b\xa9\x84\x0a\x75\x98\x1e\xb5\x40\x87\x1b\xa5\x15\xab\x87\x92\x93\xde\xba\x12\x7e\x5d\x5e\xbe\x7f\xbf\xfc\x6d\xb4\xb6\xb7\x3a\x36\xf4\x21\x5d\x88\x61\x9e\xe1\xb9\xeb\xf1\xf4\x34\x29\xaf\xdb\x5d\xb1\x47\x5f\x68\xb5\x39\x31\x7c\xef\x4f\xd0\x2c\xb4\x32\x15\x66\x0d\xba\x33\xa8\x3e\x9e\x6a\x76\x9d\x4e\x54\x7a\xbe\xc1\xad\x0d\x1d\xca\x04\xde\x3d\xa9\xdb\x33\x9d\x9e\x85\x4f\x6d\xff\x15\x00\x00\xff\xff\xd2\xc9\xf3\xad\x22\x0e\x00\x00")

func neuronDevicePluginYamlBytes() ([]byte, error) {
	return bindataRead(
		_neuronDevicePluginYaml,
		"neuron-device-plugin.yaml",
	)
}

func neuronDevicePluginYaml() (*asset, error) {
	bytes, err := neuronDevicePluginYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "neuron-device-plugin.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1d, 0x76, 0xf5, 0x81, 0xb9, 0x6f, 0x3e, 0x8, 0x73, 0x5c, 0xd3, 0xcd, 0xca, 0x64, 0x21, 0xed, 0x91, 0x66, 0x61, 0xfe, 0xec, 0x6e, 0xdf, 0x4c, 0xf4, 0x18, 0x25, 0x62, 0x19, 0xd7, 0x40, 0xf}}
	return a, nil
}

var _nvidiaDevicePluginYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x55\x41\x8f\xdb\x36\x13\xbd\xeb\x57\x3c\xd8\x97\x04\x58\x49\x9b\x5c\xbe\xaf\x0a\x7a\x70\x77\xb7\xa8\x91\x8d\x1d\xac\x37\x09\x82\xa2\x87\x31\x39\x96\x08\x53\x24\x4b\x52\xf6\xea\xdf\x17\x94\xb5\x5e\x2b\x05\x9a\x9c\x0a\x54\x17\x1b\xd2\x9b\x99\x37\x6f\xde\x90\x73\xdc\x58\xd7\x7b\x55\x37\x11\xaf\xc4\x6b\xbc\xbd\x7e\xf3\xd3\x15\x56\x9f\x97\xb7\xcb\x05\x6e\xd6\x0f\x1f\xd7\x0f\x8b\xc7\xe5\x7a\x55\x00\x0b\xad\x31\x00\x03\x3c\x07\xf6\x07\x96\x45\x36\xcf\xe6\xb8\x57\x82\x4d\x60\x89\xce\x48\xf6\x88\x0d\x63\xe1\x48\x34\xfc\xfc\xe5\x0a\x9f\xd9\x07\x65\x0d\xde\x16\xd7\x78\x95\x00\xb3\xf1\xd3\xec\xf5\xbb\x6c\x8e\xde\x76\x68\xa9\x87\xb1\x11\x5d\x60\xc4\x46\x05\xec\x94\x66\xf0\x93\x60\x17\xa1\x0c\x84\x6d\x9d\x56\x64\x04\xe3\xa8\x62\x33\x94\x19\x93\x14\xd9\x1c\x5f\xc7\x14\x76\x1b\x49\x19\x10\x84\x75\x3d\xec\xee\x12\x07\x8a\x03\xe1\xf4\x34\x31\xba\xaa\x2c\x8f\xc7\x63\x41\x03\xd9\xc2\xfa\xba\xd4\x27\x60\x28\xef\x97\x37\x77\xab\xcd\x5d\xfe\xb6\xb8\x1e\x42\x3e\x19\xcd\x21\x35\xfe\x67\xa7\x3c\x4b\x6c\x7b\x90\x73\x5a\x09\xda\x6a\x86\xa6\x23\xac\x07\xd5\x9e\x59\x22\xda\xc4\xf7\xe8\x55\x54\xa6\xbe\x42\xb0\xbb\x78\x24\xcf\xd9\x1c\x52\x85\xe8\xd5\xb6\x8b\x13\xb1\x9e\xd9\xa9\x30\x01\x58\x03\x32\x98\x2d\x36\x58\x6e\x66\xf8\x65\xb1\x59\x6e\xae\xb2\x39\xbe\x2c\x1f\x7f\x5b\x7f\x7a\xc4\x97\xc5\xc3\xc3\x62\xf5\xb8\xbc\xdb\x60\xfd\x80\x9b\xf5\xea\x76\x99\x06\xb5\xc1\xfa\x57\x2c\x56\x5f\xf1\x7e\xb9\xba\xbd\x02\xab\xd8\xb0\x07\x3f\x39\x9f\xf8\x5b\x0f\x95\x64\x1c\x46\x87\x0d\xf3\x84\xc0\xce\x9e\x08\x05\xc7\x42\xed\x94\x80\x26\x53\x77\x54\x33\x6a\x7b\x60\x6f\x94\xa9\xe1\xd8\xb7\x2a\xa4\x61\x06\x90\x91\xd9\x1c\x5a\xb5\x2a\x52\x1c\xde\xfc\xad\xa9\x22\xcb\xc8\xa9\x71\xfc\x55\xd2\x2c\x94\x87\x37\xd9\x5e\x19\x59\xe1\x96\xb8\xb5\x66\xc3\x31\x6b\x39\x92\xa4\x48\x55\x06\x18\x6a\xb9\x82\x39\x28\xa9\x28\x97\x7c\x50\x82\x73\xa7\xbb\x5a\x99\x5c\x0e\x01\x81\xe3\x08\x0b\x8e\x04\x57\xd8\x77\x5b\xce\x43\x1f\x22\xb7\x59\xe2\x9e\xb2\x04\xd6\x2c\xa2\xf5\xe9\x3f\xd0\x52\x14\xcd\x3d\x6d\x59\x87\xd3\x8b\x7f\x2e\x13\x32\xa0\x73\x92\x22\x6f\xa2\xa7\xc8\x75\x7f\x8a\x8a\xbd\xe3\x0a\x0f\x56\x6b\x65\xea\x4f\x03\x20\x03\x22\xb7\x4e\x53\xe4\xb1\xd4\x45\x2b\xe9\x99\xe3\x31\xb9\x99\x8c\xb1\x27\x95\x86\x39\xb3\xf3\x2c\x28\xb2\x2c\xf0\x3e\x19\xbc\x61\x7f\xd2\x7f\x4b\x62\x7f\x24\x2f\x07\xbf\x53\x54\x5b\xa5\x55\xec\xcf\xb9\xd2\xc8\x92\x75\x43\x55\x96\xa9\x6d\x6f\x38\x72\x28\x94\x2d\xa5\x15\xa1\x8c\x14\xf6\xa1\x24\xd9\x2a\xa3\x42\x64\x9f\x0b\xdd\xa5\xdf\xb2\xee\xc8\x93\x89\xcc\x32\x0f\xa2\x61\xd9\xa5\x0e\x72\x91\x3c\x2a\x48\xe7\x24\xa5\x35\xb9\xb3\x32\x94\x63\xa9\x17\xbe\x67\xc5\x80\x31\x94\x7d\x41\xda\x35\x54\x4c\x19\x9c\xb3\x39\x2b\x2b\xcc\x66\x63\x98\x9e\xc8\xfe\x7d\xe1\x81\xe7\x19\x0e\x92\x5b\xcd\x7e\xca\x63\x54\xf4\xe5\xcb\x7f\x43\xd1\x1c\x7b\xee\x2b\xdc\x8c\x88\x45\x02\x84\xb5\xd1\xfd\x59\x19\xeb\x52\x43\xd6\x57\xb8\x7b\x52\x21\x86\x69\xe0\x49\xb1\x42\xd8\xb6\xac\x5d\xf7\xbd\x20\x80\x77\x3b\x16\xb1\xc2\xca\x6e\xc6\xb1\x9d\x7b\xfe\x40\x7e\x7f\x3a\x64\x9d\x95\xa0\x90\x8e\xcb\x91\x16\x48\xca\xdc\x9a\x77\x38\x36\x6c\xc0\x26\x9d\x6f\xf2\x6a\x58\xe9\x6f\x20\xe7\x6c\x67\x57\x3c\xdf\x0b\xc3\x05\x61\x3b\x2f\x38\x0c\x13\xf8\x26\x30\x15\x0d\x08\x16\xb1\xa1\x98\x32\xf7\x10\xf4\x92\x6e\xcb\x29\x7c\xcc\x29\x41\xbb\xc8\x1e\x84\x1d\x29\xdd\x79\x2e\xfe\xfd\xc1\x39\xaf\xac\x57\xb1\xbf\xd1\x14\xc2\x6a\x70\xef\xec\x74\xda\xe4\xc6\x4a\x3e\x87\x3e\x1b\x5e\x58\x93\xee\x20\xf6\x67\xc7\xe6\x50\x2d\xd5\x67\xd7\x97\xfb\xff\x87\xa9\xf3\xab\xc3\x75\xf1\xbf\xe2\xfa\x47\x76\x44\x44\x7f\x86\x91\xaf\x43\x85\xdf\x67\x79\x9e\xd4\xc9\xad\xc9\x95\x51\x31\x67\xef\xad\xff\x79\x47\x3a\xf0\xec\x8f\x97\xe5\x65\xd1\x0d\x6d\x58\x13\xf9\x29\xbe\x2c\x24\x40\x5a\xdb\xe3\x47\xaf\x0e\x4a\x73\xcd\x77\x41\x90\x1e\x16\xab\xc2\x90\xe4\x02\x29\xc8\xd1\xb0\x40\x8a\xc3\x65\x06\x40\x7a\xeb\x12\x97\xc5\xfd\xfd\x45\xd1\x83\xd5\x5d\xcb\x1f\x6c\x67\xe2\x04\x9f\x8f\x2d\x4e\x7a\x9b\xe4\x6b\x53\xcc\x47\x8a\x4d\x85\xf2\x40\xbe\xd4\x6a\x3b\x8c\x59\x73\x2c\x27\x51\xcf\x8e\x3f\x95\xba\xa8\xf2\xbd\x1a\x8d\x0d\xa7\x02\x93\xba\xee\x87\x4a\x66\x7f\x05\x00\x00\xff\xff\x31\x6e\xa1\x29\x39\x09\x00\x00")

func nvidiaDevicePluginYamlBytes() ([]byte, error) {
	return bindataRead(
		_nvidiaDevicePluginYaml,
		"nvidia-device-plugin.yaml",
	)
}

func nvidiaDevicePluginYaml() (*asset, error) {
	bytes, err := nvidiaDevicePluginYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "nvidia-device-plugin.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc4, 0x9d, 0xa, 0xb7, 0x22, 0x67, 0xb3, 0x61, 0xf5, 0x37, 0x17, 0x90, 0x53, 0x59, 0x85, 0x25, 0xc7, 0x1, 0x1d, 0x96, 0x1, 0x51, 0xd5, 0x17, 0xea, 0x87, 0x44, 0x28, 0xc8, 0x3a, 0x53, 0xfc}}
	return a, nil
}

var _vpcAdmissionWebhookConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x91\x4f\x6b\xf3\x30\x0c\xc6\xef\xf9\x14\x22\xf7\xa4\xf4\xf6\xe2\xdb\x4b\x29\x63\x87\xc1\x18\x63\x3b\x8c\x1d\x14\x47\x4d\x45\x62\xcb\x58\x76\x4a\xf7\xe9\x47\xfe\xb4\xac\xb0\xd5\x17\xdb\x7a\xa4\xdf\x63\xc9\x18\xf8\x8d\xa2\xb2\x78\x03\xd8\x3a\xd6\xe9\x18\xa9\x63\x4d\x11\x13\x8b\xaf\xfb\x7f\x5a\xb3\x6c\xc6\x6d\x43\x09\xb7\x45\xcf\xbe\x35\xf0\x94\x13\x26\xf6\xdd\x3b\x35\x47\x91\x7e\x27\xfe\xc0\x5d\x5e\x2a\x0a\x47\x09\x5b\x4c\x68\x0a\x00\x8f\x8e\x0c\x8c\xc1\x56\x57\x7a\x75\x5a\x8a\x2a\x7b\xe8\xd6\x0c\x0d\x68\xc9\x40\x9f\x1b\xaa\xf4\xac\x89\x5c\x01\x30\x60\x43\x83\x4e\x10\x00\x0c\xe1\x0f\x4a\xb1\xee\x73\x62\x75\xcf\xaf\x46\x87\x5f\xe2\xf1\xa4\xb5\x15\x37\x63\xed\xc0\xe4\xd3\xf2\xfa\xc5\x08\x40\x29\x8e\x6c\xe9\x72\xbd\xdb\xc2\x4d\xce\xaf\x4d\x2c\x2b\x60\x3a\x1a\x28\x37\x6e\x1a\x1b\x95\x73\x3c\xe6\x81\xf4\xe2\x52\x81\x04\x5a\xc6\xa7\x06\x3e\xa0\xdc\xbd\xec\xff\xbf\xee\x4b\xf8\xbc\x32\x30\xf0\x43\x94\x1c\x26\xbd\x2c\x6f\xe2\xeb\x0f\xce\xca\xb8\xfd\xa1\x45\x52\xc9\xd1\xd2\xac\x04\x69\x75\xd5\x0e\xc8\x43\x8e\xf4\x2c\x03\xdb\xb3\x81\xc7\xce\x4b\xa4\xe2\x3b\x00\x00\xff\xff\x49\xee\x9e\x02\x0c\x02\x00\x00")

func vpcAdmissionWebhookConfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_vpcAdmissionWebhookConfigYaml,
		"vpc-admission-webhook-config.yaml",
	)
}

func vpcAdmissionWebhookConfigYaml() (*asset, error) {
	bytes, err := vpcAdmissionWebhookConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vpc-admission-webhook-config.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x7d, 0x11, 0x23, 0x17, 0x2f, 0xdc, 0x4e, 0x18, 0xaa, 0xc8, 0x66, 0xee, 0xf3, 0xc1, 0x85, 0x63, 0xb1, 0xe3, 0x53, 0x57, 0x80, 0x96, 0xe6, 0x54, 0x26, 0x46, 0x6b, 0x3f, 0x17, 0x20, 0x31, 0x8}}
	return a, nil
}

var _vpcAdmissionWebhookCsrYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8d\x4b\x4e\xc4\x30\x10\x44\xf7\x3e\x45\x5f\x20\x41\xb3\x43\xde\x72\x03\x90\xd8\x77\xec\xc2\x69\x19\x7f\x70\xb7\x83\xe6\xf6\x28\x13\xa4\xd9\x95\xaa\x9e\x5e\x71\x97\x4f\x0c\x95\x56\x3d\x05\x0c\x93\x2f\x09\x6c\xd0\x35\xbf\xea\x2a\xed\xe5\xb8\x6d\x30\xbe\xb9\x2c\x35\x7a\x7a\x7b\x12\x1f\x92\xaa\xd4\xf4\x8e\x9f\x09\x35\x57\x60\x1c\xd9\xd8\x3b\xa2\xca\x05\x9e\x8e\x1e\x16\x8e\x45\xf4\x94\x2f\xbf\xd8\xf6\xd6\xf2\x9a\xe7\x86\x45\xef\x6a\x28\x4e\x3b\xc2\xc9\xa7\xd1\x66\xd7\x33\x2d\x74\x4d\x9e\xa7\xed\xa8\xf6\x78\x8a\x8e\x68\x2a\x27\xfc\x23\x51\x92\x18\x7f\x93\x4a\xaa\x6c\x73\xe0\xd1\x66\xdc\x09\x35\x48\xdf\x31\x0a\xaa\x5d\x36\x8c\x03\x83\x4e\x9b\xfb\x0b\x00\x00\xff\xff\xf1\x7d\x42\x97\xea\x00\x00\x00")

func vpcAdmissionWebhookCsrYamlBytes() ([]byte, error) {
	return bindataRead(
		_vpcAdmissionWebhookCsrYaml,
		"vpc-admission-webhook-csr.yaml",
	)
}

func vpcAdmissionWebhookCsrYaml() (*asset, error) {
	bytes, err := vpcAdmissionWebhookCsrYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vpc-admission-webhook-csr.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x56, 0xf4, 0x54, 0x57, 0xf8, 0xbd, 0x8a, 0x1a, 0x67, 0xb2, 0x29, 0x18, 0xe2, 0x44, 0x8a, 0xc2, 0x7f, 0x44, 0x26, 0x4c, 0x52, 0xe0, 0x70, 0xe0, 0xc8, 0x5a, 0x74, 0x76, 0x2, 0xcd, 0x3e, 0xa}}
	return a, nil
}

var _vpcAdmissionWebhookDepYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x54\xc1\x6e\xdb\x3a\x10\xbc\xfb\x2b\xf6\x92\x97\x93\xa5\x97\x20\xe8\x81\x40\x0a\x04\x4d\x0a\x04\x6d\x93\xa0\x29\x7a\x5f\x53\x13\x99\x10\x45\xb2\xe4\x4a\x8e\xfe\xbe\xa0\xad\xa4\xb6\x1c\x37\x39\x14\xad\x0e\x3e\xec\xec\xcc\x2c\x87\x5e\x72\x30\xdf\x11\x93\xf1\x4e\x11\x87\x90\xca\xfe\x64\xd6\x18\x57\x29\xba\x44\xb0\x7e\x68\xe1\x64\xd6\x42\xb8\x62\x61\x35\x23\x72\xdc\x42\x51\x1f\xf4\x9c\xab\xd6\xa4\xcc\x9c\xaf\xb0\x58\x7a\xdf\x8c\x68\x0a\xac\xa1\xa8\xe9\x16\x98\xa7\x21\x09\xda\x19\x91\xe5\x05\x6c\xca\x02\x94\x7d\x0e\x29\xa4\x00\x9d\x9b\x22\x82\x35\x9a\x93\xa2\x93\x19\x51\x92\xc8\x82\x7a\xd8\xd0\x65\x08\x50\xf4\x15\x3a\x82\x05\x19\x86\x85\x16\x1f\x37\x70\xcb\xa2\x97\x9f\xb7\xec\x7e\x6b\x48\x24\x68\x83\x65\xc1\xc8\xde\x3a\x6a\xfe\xec\x8e\xd0\x2b\x52\x44\x4f\xf3\xe7\x4f\x7b\x27\x6c\x1c\xe2\x16\x7d\xfe\x4a\x7e\xcf\x36\xb1\xde\x62\x6d\x98\x73\xb1\xe9\x03\xa2\x7c\x34\x16\xe7\x25\x44\x97\x23\xaf\xd4\x88\x92\xd6\xbf\x45\x58\xa7\x3d\xa5\x7d\xc2\x70\x88\xd5\x60\x78\x89\x74\x7b\xbf\x8e\xf0\x7e\x8c\xf6\xb6\x47\x8c\xa6\xc2\xf9\xca\xb8\xca\xaf\xd2\xb4\x9d\x6d\xf2\xd6\xd7\xe2\x93\x54\x88\x71\x0a\xf7\xe7\x67\x93\xd2\xe9\xfb\xff\x4e\xb6\x4a\xa6\xe5\x1a\x8a\x8e\x8f\x52\x51\x35\xb1\x80\x8e\xc5\x51\x2a\x8e\x52\x89\x26\x95\x2f\x86\xa5\xfa\xff\x8b\xd3\xe2\xec\x78\x2a\x72\xd7\x59\x7b\xe7\xad\xd1\x83\xa2\x0b\xbb\xe2\x61\x7b\xd6\xde\xdb\xae\xc5\x17\xdf\x39\xd9\x8b\x77\x73\x31\xa3\xfa\x7c\x1d\xce\x4e\x07\x51\x9b\x79\x77\x2c\x4b\x45\xfb\x41\x4e\x7a\x23\xb8\xba\x75\x76\x50\x24\xb1\xc3\x08\x2e\x7d\x92\x1b\xc8\xca\xc7\x66\xa7\xce\x0f\x0f\xc6\x19\x19\x7e\x8d\xe4\x7c\x85\x8b\xbd\x6a\x96\xfd\xd1\x99\x88\xea\xb2\x8b\xc6\xd5\xf7\x7a\x89\xaa\xb3\xc6\xd5\xd7\xb5\xf3\xcf\xe5\xab\x47\xe8\x4e\xf2\x4a\xef\x0c\x95\x35\x9f\xee\xf3\x1b\x62\xbb\x97\xc0\x7a\x73\xae\x1e\x43\xc4\x3a\xe9\x09\x9e\x3b\x1a\x0c\x8a\x16\x10\x2e\xf2\x76\x47\x07\x41\x2a\x8c\x2f\xfd\xf4\xf8\x44\x3e\x20\x72\x5e\x4a\xba\x76\x7b\x60\xcf\xb6\xc3\x9e\x7e\x76\xb0\xc6\x75\x8f\x6f\xf6\xe5\xa8\x97\x7f\xca\x99\xdb\xea\xdd\xf4\x5f\xfa\xc6\x44\xfe\x41\x18\x7f\x27\x87\xcd\xbe\xbc\xf0\x7e\x1d\x5a\x93\x94\x5f\x65\xd9\x95\xdd\xd4\x6e\x0e\xbf\x7b\xa3\xca\xcf\x00\x00\x00\xff\xff\x55\x30\xb6\xdb\x8b\x06\x00\x00")

func vpcAdmissionWebhookDepYamlBytes() ([]byte, error) {
	return bindataRead(
		_vpcAdmissionWebhookDepYaml,
		"vpc-admission-webhook-dep.yaml",
	)
}

func vpcAdmissionWebhookDepYaml() (*asset, error) {
	bytes, err := vpcAdmissionWebhookDepYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vpc-admission-webhook-dep.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x45, 0x84, 0xcf, 0x10, 0xfa, 0xe, 0x24, 0xeb, 0xff, 0xf4, 0xbb, 0xae, 0xcc, 0xc4, 0xc, 0x4f, 0xdb, 0xf8, 0x5f, 0x62, 0x5c, 0x88, 0x2c, 0xee, 0xbd, 0x4e, 0x6b, 0x47, 0xeb, 0x28, 0xaa, 0x0}}
	return a, nil
}

var _vpcAdmissionWebhookYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8c\x31\xae\xc2\x40\x0c\x05\xfb\x3d\x85\x2f\xe0\xe2\xeb\xa7\xf2\x29\x90\x90\xe8\x9d\xcd\x13\xac\x92\xcd\x5a\x6b\x13\xc4\xed\x51\x22\x0a\x1a\x44\x67\xf9\xcd\x0c\x33\x27\xb5\x72\x41\xf7\xd2\x56\xa1\xed\x2f\xcd\x65\x9d\x84\xce\xe8\x5b\xc9\x48\x15\xa1\x93\x86\x4a\x22\x5a\xb5\x42\x68\xb3\xcc\x3a\xd5\xe2\xbb\xc1\x0f\x8c\xb7\xd6\xe6\xf7\xea\xa6\x19\x42\xf3\x7d\x04\xfb\xd3\x03\x35\x11\x2d\x3a\x62\xf1\x3d\x40\xa4\x66\xdf\x0a\x6e\xc8\x3b\x64\xad\xc7\x41\xf3\x71\x0a\x0d\xc3\xff\xe1\x86\xf6\x2b\xe2\xf4\xf1\x73\x2c\xc8\xd1\xfa\xcf\xf6\x2b\x00\x00\xff\xff\xbc\xa7\x78\x3e\xe7\x00\x00\x00")

func vpcAdmissionWebhookYamlBytes() ([]byte, error) {
	return bindataRead(
		_vpcAdmissionWebhookYaml,
		"vpc-admission-webhook.yaml",
	)
}

func vpcAdmissionWebhookYaml() (*asset, error) {
	bytes, err := vpcAdmissionWebhookYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vpc-admission-webhook.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc0, 0x14, 0x3e, 0x68, 0x1c, 0x57, 0x26, 0x7e, 0x3b, 0xab, 0xf1, 0x55, 0x21, 0x61, 0xc3, 0xe1, 0xaf, 0x46, 0xb6, 0xf7, 0xdd, 0x11, 0x29, 0x41, 0x64, 0x57, 0xc8, 0xd4, 0xea, 0x97, 0xba, 0xff}}
	return a, nil
}

var _vpcResourceControllerDepYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x94\xcf\x6f\xdb\x3a\x0c\xc7\xef\xfe\x2b\x78\x29\x7a\xb2\x93\xf4\xf5\xb5\x78\x02\xde\xa1\x58\x8b\xad\xc0\x30\x04\x68\xb0\x3b\x23\x31\xb1\x60\x59\xd2\x28\xda\x89\xf7\xd7\x0f\x5a\x9b\x1f\x6e\xba\xb5\x18\x36\xcc\x27\xe3\xfb\x91\xf8\x25\x09\x91\x65\x59\x16\x18\xed\x67\xe2\x64\x83\x57\x80\x31\xa6\x49\x3f\x2b\x1a\xeb\x8d\x82\x5b\x8a\x2e\x0c\x2d\x79\x29\x5a\x12\x34\x28\xa8\x0a\x00\x8f\x2d\x29\xe8\xa3\x2e\x99\x52\xe8\x58\x53\xa9\x83\x17\x0e\xce\x11\x3f\xf1\x14\x51\x93\x82\xa6\x5b\x52\x99\x86\x24\xd4\x16\x29\x92\xce\xd7\x99\xa2\xb3\x1a\x93\x82\x59\x01\x90\xc8\x91\x96\xc0\x99\x00\xb4\x28\xba\xfe\x88\x4b\x72\xe9\x51\x80\x9c\xd2\xcf\xcc\xf2\x27\x96\x58\xc1\x12\x75\x43\xde\xec\x34\x46\xdd\x28\x48\x82\x4b\x47\x05\x80\x50\x1b\x1d\x0a\x3d\xf9\x1c\x95\x93\x3f\x37\xb2\x7c\x93\xe9\xcb\xb6\xa7\xc6\x00\xbb\xc2\xbf\xff\x13\xf7\x56\xd3\x8d\xd6\xa1\xf3\xf2\x9a\x47\x16\xd0\x7a\xe2\x7d\x6a\x25\xe8\xd0\xb6\xe8\xcd\x21\xd7\x12\x26\xaf\x65\x8a\xbc\x4e\xc7\x17\xca\x24\x86\x98\xa5\x66\x4a\x75\x70\xe6\x7f\xeb\x57\x61\xcf\x6d\x8b\x6b\x52\x70\x7e\x96\x2a\xd3\x70\x45\x9a\xab\xb3\x54\x9d\xa5\x09\x35\x69\xb2\xb1\xde\x84\x4d\x2a\x7f\x60\xa9\xfa\x69\x75\x51\x5d\x9e\x8f\x83\xcd\x3b\xe7\xe6\xc1\x59\x3d\x28\xb8\x71\x1b\x1c\xd2\x9e\x3b\xdb\x93\xa7\x94\xe6\x1c\x96\x74\xc8\x11\x60\x85\xd6\x75\x4c\x8b\x5d\x8e\x0a\xfe\x3d\xa2\xb5\x48\x7c\x4f\x72\x7c\x01\xa0\x0e\x49\x14\xcc\x2e\xae\xab\x69\x35\xad\x66\x23\x16\x51\x6a\x05\x93\x9a\xd0\x49\xfd\x75\x8c\x02\x8b\x82\xab\xd9\xf5\xf5\x7f\x23\x3d\xe9\x9a\xf2\x4b\xff\xb0\x58\xcc\x8f\x80\xf5\x56\x2c\xba\x5b\x72\x38\x3c\x90\x0e\xde\x24\x05\xff\x4c\x8f\x4e\x44\x62\x1b\xcc\xcb\x4c\x6c\x4b\xa1\x93\x3d\x3c\x14\xf5\xda\x58\xed\x1e\x90\xee\xd8\xca\xf0\x2e\x78\xa1\xed\xa8\x01\x91\x6d\x6f\x1d\xad\xc9\x28\x10\xee\xa8\x38\x74\xe5\x13\xc9\x26\x70\x33\xd2\x71\xb5\xca\xa5\x0c\x87\x10\x3e\x18\xba\x39\x51\xf3\xc4\x7e\xe9\x2c\x93\xb9\xed\xd8\xfa\xf5\x83\xae\xc9\x74\xce\xfa\xf5\xfd\xda\x87\xbd\x7c\xb7\x25\xdd\x49\xde\x22\xa3\x26\xe6\x98\x0f\x4f\x33\xbe\x20\x6e\xd3\x18\x97\x8f\x23\x7f\xb7\x8d\x4c\x29\xef\xa0\x67\x3c\x9f\x68\x68\x50\xb0\x24\xc1\x2a\xaf\x13\xf6\x24\x94\x2a\x1b\x26\x21\x3d\x3b\x0a\x10\x22\x31\xe6\x6d\x02\xf7\xfe\x04\xf6\xe8\x3a\x3a\x89\x9f\x1d\x9c\xf5\xdd\xf6\xcd\xbe\xc8\xba\xfe\x5d\xce\xd8\x9a\xab\xcb\x5f\xeb\xc8\x5f\x68\xc6\x9f\xee\xc3\xb7\x00\x00\x00\xff\xff\x75\xe0\xac\xa8\x8f\x06\x00\x00")

func vpcResourceControllerDepYamlBytes() ([]byte, error) {
	return bindataRead(
		_vpcResourceControllerDepYaml,
		"vpc-resource-controller-dep.yaml",
	)
}

func vpcResourceControllerDepYaml() (*asset, error) {
	bytes, err := vpcResourceControllerDepYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vpc-resource-controller-dep.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x89, 0x65, 0xcc, 0x32, 0x3, 0x76, 0x1c, 0x2a, 0x22, 0x9f, 0xd0, 0xc4, 0xcc, 0xb1, 0xba, 0x48, 0x57, 0x6, 0x58, 0x7, 0x2, 0x8, 0xd3, 0x14, 0x93, 0x22, 0xb5, 0xd8, 0xb8, 0x7b, 0x4f, 0xde}}
	return a, nil
}

var _vpcResourceControllerYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x90\x31\x4f\x33\x31\x0c\x86\xf7\xfc\x8a\xa8\x7b\x5a\x7d\xdb\xa7\xdb\x80\x81\xbd\x48\xec\x3e\x9f\xdb\x9a\xe6\xe2\xc8\x76\x0e\xc1\xaf\x47\x77\xd7\xc2\x80\x44\x85\x98\xf2\xc4\xf6\x6b\x4b\x4f\x4a\x29\x40\xe5\x67\x52\x63\x29\x5d\xd4\x1e\x70\x0b\xcd\x4f\xa2\xfc\x0e\xce\x52\xb6\xe7\xff\xb6\x65\xd9\x4d\xff\xc2\x99\xcb\xd0\xc5\x87\xdc\xcc\x49\xf7\x92\x29\x8c\xe4\x30\x80\x43\x17\x62\x2c\x30\x52\x17\xa7\x8a\x49\xc9\xa4\x29\x52\x42\x29\xae\x92\x33\x69\xd0\x96\xc9\xba\x90\x22\x54\x7e\x54\x69\xd5\xe6\x4c\x8a\x9b\x4d\x88\xf1\x1a\xb8\xd4\x8a\x0c\x64\x5f\xb4\x33\x07\x6f\x6b\xa1\xca\xb0\x02\x4a\x39\xf0\x71\x84\x3a\x7f\x27\xd2\xfe\x92\x6d\x75\x00\xa7\x05\x8f\xe4\xcb\x9b\xd9\x56\x78\x05\xc7\xd3\xba\xe6\x93\x50\x69\x9e\xff\x9b\x87\x7b\x2e\x03\x97\xe3\x6f\x74\x48\xa6\x3d\x1d\xe6\xc1\xab\x90\x1f\x8e\x86\x18\xbf\xbb\xbf\x75\xc2\x5a\xff\x42\xe8\x8b\xf4\x35\xfd\x44\x3a\x31\xd2\x1d\xa2\xb4\xe2\x37\x17\xac\x7d\xab\x80\xd4\xc5\x73\xeb\x29\xd9\x9b\x39\x8d\xe1\x23\x00\x00\xff\xff\x83\x2e\x8d\x64\x35\x02\x00\x00")

func vpcResourceControllerYamlBytes() ([]byte, error) {
	return bindataRead(
		_vpcResourceControllerYaml,
		"vpc-resource-controller.yaml",
	)
}

func vpcResourceControllerYaml() (*asset, error) {
	bytes, err := vpcResourceControllerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vpc-resource-controller.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa8, 0xff, 0xd2, 0xc, 0x46, 0xfb, 0xe4, 0x4e, 0xb, 0x26, 0x4a, 0xa0, 0x7d, 0x73, 0x5e, 0xaf, 0x6d, 0xab, 0xfd, 0xe5, 0xc2, 0x4f, 0xed, 0xae, 0xb6, 0xd7, 0x17, 0x38, 0x68, 0xb0, 0xe8, 0x2d}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"efa-device-plugin.yaml":            efaDevicePluginYaml,
	"neuron-device-plugin.yaml":         neuronDevicePluginYaml,
	"nvidia-device-plugin.yaml":         nvidiaDevicePluginYaml,
	"vpc-admission-webhook-config.yaml": vpcAdmissionWebhookConfigYaml,
	"vpc-admission-webhook-csr.yaml":    vpcAdmissionWebhookCsrYaml,
	"vpc-admission-webhook-dep.yaml":    vpcAdmissionWebhookDepYaml,
	"vpc-admission-webhook.yaml":        vpcAdmissionWebhookYaml,
	"vpc-resource-controller-dep.yaml":  vpcResourceControllerDepYaml,
	"vpc-resource-controller.yaml":      vpcResourceControllerYaml,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"efa-device-plugin.yaml": {efaDevicePluginYaml, map[string]*bintree{}},
	"neuron-device-plugin.yaml": {neuronDevicePluginYaml, map[string]*bintree{}},
	"nvidia-device-plugin.yaml": {nvidiaDevicePluginYaml, map[string]*bintree{}},
	"vpc-admission-webhook-config.yaml": {vpcAdmissionWebhookConfigYaml, map[string]*bintree{}},
	"vpc-admission-webhook-csr.yaml": {vpcAdmissionWebhookCsrYaml, map[string]*bintree{}},
	"vpc-admission-webhook-dep.yaml": {vpcAdmissionWebhookDepYaml, map[string]*bintree{}},
	"vpc-admission-webhook.yaml": {vpcAdmissionWebhookYaml, map[string]*bintree{}},
	"vpc-resource-controller-dep.yaml": {vpcResourceControllerDepYaml, map[string]*bintree{}},
	"vpc-resource-controller.yaml": {vpcResourceControllerYaml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
