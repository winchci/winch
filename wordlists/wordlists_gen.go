// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package wordlists

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2020, 6, 2, 8, 55, 52, 273141424, time.UTC),
		},
		"/adjectives.txt": &vfsgen۰CompressedFileInfo{
			name:             "adjectives.txt",
			modTime:          time.Date(2019, 11, 27, 10, 4, 8, 580910182, time.UTC),
			uncompressedSize: 10489,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x5a\x59\x9a\xad\xac\x0e\x7d\x67\x2e\x35\x28\xd4\xa8\xa9\x4d\x77\x02\xe8\xb1\x46\x7f\xbf\xb5\x62\x9d\xfb\x3f\xec\x24\xf4\x01\x42\x3a\x77\x5c\x62\xd9\x6a\x91\x2d\xc4\x25\x49\x88\x4b\xaf\x69\x0e\x09\x71\x8d\x9b\x64\x5d\x43\x5c\x57\x69\xc3\x1b\xd7\x35\x45\xcd\xe8\xbb\xae\x35\xb7\xa4\xfd\xf4\xc2\xb4\xc8\x31\xa7\x96\x23\xc4\x55\x37\x0e\xb4\xba\xc4\x41\x6a\xe8\x85\xe6\x31\x63\x0a\x71\x93\x36\x42\xdc\xb2\x9a\xcf\x0a\x0a\xd3\x6c\x35\x49\x5f\xa5\xa0\xb1\xfe\xb6\x55\x6f\xba\x62\x59\x9d\x90\x32\xa6\xd5\xd9\x43\xdc\x77\x59\x87\xd6\xc2\xa5\x77\x8b\xba\x85\x78\xa0\xd3\x71\x58\xbc\xe2\x20\x2f\xc7\x61\xd2\x3b\x97\x3f\x34\x11\x8e\x38\xd8\xab\x16\xfd\xf1\x3e\x26\xe2\xcb\xa5\x68\xdc\x1e\x30\x9b\x92\xd8\x08\x31\xa9\x14\x1f\x94\x38\x53\x1a\x36\xb5\x73\x6b\x39\xfa\x1c\x79\xd1\xa1\x64\x2b\x37\xcc\x94\x67\x47\xff\x3c\x3b\x9b\xcb\x7a\xfa\x4e\xca\xaa\xdc\x61\x39\x24\x61\x7c\x39\xec\x01\x9c\xef\x59\x16\xcd\xbe\x52\x29\x3c\xac\x32\xf4\xcf\x94\x10\xcb\x5f\x9f\xbd\x35\x93\x53\x8a\xef\xa8\x35\xab\xcd\x94\x07\x80\x33\xb5\x95\x3c\x19\x4e\xc2\x6a\xf6\xc3\xb7\xf1\xb2\xda\xcf\xc8\xcd\xf5\x3e\xc9\x4a\x1f\xb5\x68\xf7\x2b\x1b\x67\x12\x76\x1a\x23\xae\x64\x64\x0c\x29\x7e\x6d\x63\xd8\xef\x0d\xce\x3e\xc4\x80\xc7\x89\xd6\x95\x54\x35\xfd\xc1\x88\x39\x7e\x97\xbc\xa2\xe9\xea\xfc\x5e\x62\xf1\x90\x10\xef\x88\x71\xb7\xf4\x9a\x81\xf7\x99\x42\xbc\x3f\x77\xb4\x2d\x2c\x71\x79\xb4\x9f\x61\x89\xeb\x27\x2c\x11\x15\xc7\xf1\x84\x05\x23\x96\x68\x26\x25\x2c\xb1\xeb\x1a\x16\x89\x73\x28\x86\x2e\x92\x78\x4a\x8b\xa4\x7a\x11\x17\xd9\x75\xd5\x88\xa6\x5b\x07\xb7\xb0\xe8\x81\xdf\xd7\x29\xd1\xd8\x59\xeb\x26\x87\xc5\x8d\x77\xbd\xe8\x90\xaf\x4e\xce\x17\x1d\x43\x2c\x2c\x29\x96\x8d\xf0\x03\x68\x38\x98\x25\x49\x44\x49\xd9\xa2\xbd\x73\xf5\x34\xfd\xd8\x96\x7a\xcc\x1e\x96\xaa\xc9\x4b\x69\x0b\x4b\x2d\x4f\x58\xaa\x0f\xae\xbd\xa3\x30\xcb\xea\xe8\x65\xbe\xde\x58\xd4\xe2\x25\x61\x31\x89\x1f\xe7\xc7\x54\x76\xc0\xe3\x1c\x40\x29\x69\x2c\xa4\xfa\x27\x2c\x56\x3f\x38\x06\xab\xe5\x07\x5d\xa7\x42\xba\x96\xb9\x2c\xe9\x09\xcb\x4c\x1f\xc0\xdc\x00\xeb\xc3\x61\xd3\x36\x29\x3c\xec\x65\x1a\x3b\xf5\xe1\x5c\x4e\xf0\x34\xb1\x63\xe0\x1f\x8a\xef\x1a\xd3\x3a\x93\x3f\x99\x35\xa6\x1c\xd6\x58\x36\xdd\xc2\x1a\x4d\x76\x13\x71\x62\x26\xe2\x24\xbd\x83\xf0\xce\xd3\xc5\x7e\xc5\x4d\x17\x52\x92\x64\x31\x5e\xcf\x7a\xbe\xef\x68\x3d\x25\x36\x40\x31\xce\x02\xe2\x09\xeb\x89\x2d\xaf\xa7\xa6\x84\xc2\x5c\x96\x27\xac\x6a\xe0\xc4\xc2\x9a\x62\xc7\xa5\xaf\x49\x62\x21\x44\x9d\x5c\x02\x54\xbb\x38\xdc\x80\xe6\xf6\x84\x35\xcd\x97\xb1\x34\x73\x67\x19\x5b\x44\x87\x1a\x0d\xdd\x71\x3b\x6b\x4d\xd5\x39\x00\xe1\xfd\x6b\xaa\xbd\x47\x54\xe5\xbd\x9a\x6b\xba\xb5\xe6\x5c\x0b\x50\x03\x17\xae\x63\x50\x12\xbc\x0a\x52\x49\xde\xaa\x24\x7f\x1d\xeb\xea\x9b\xae\xb9\x39\x67\xb5\xac\x62\xe5\xa5\xcc\xfb\x97\x7d\xbe\x6d\xdd\x5f\x08\x28\xdd\xc4\x7c\x85\xd2\x47\xe4\x02\xe5\x5d\xa8\x40\xe1\x81\x01\x30\x58\x3f\x1c\x5a\x49\x37\x8c\xc1\xab\x5c\x6b\xb5\x4d\xcb\xbb\xb8\x95\x07\xd0\x66\xc3\xf0\x3e\x70\xb4\x75\xe2\x11\xfa\x6a\xd3\x86\x53\x16\xf7\xf1\x00\xfd\x00\x4a\xcc\x8e\x7c\x4a\x13\x69\x28\x6b\x56\x2e\x6d\xda\x1b\xe0\xd0\x95\xc5\x97\x15\xab\xf7\x46\x3c\x05\xb5\xef\xb3\x58\xe7\xb6\x61\xd9\x99\x86\x5e\xce\xd6\x4c\x83\x4a\x67\x9d\x79\x11\xa3\x5c\xae\x94\xcb\x75\xda\x05\x88\xdd\x3f\x78\x66\xc6\x15\xb6\x98\xa9\xcb\xb7\x98\x5b\xd8\xa0\x2d\xa9\xf3\xb7\xd8\x9a\x58\xd8\x5c\xf8\xb6\x68\x1f\x80\xe4\x85\x9f\x1f\x27\x24\x6e\x04\xe9\x01\xda\xa5\xbc\xb5\x46\x20\x7d\x84\x4d\x68\x66\x36\x59\x35\x63\x31\x59\x95\x0a\x75\x13\x69\x61\x93\x5d\x4a\x77\x61\x72\xda\x9b\x76\xbe\xc7\x8d\x7a\x46\x5e\xaa\xe8\x90\x5f\xc2\x7b\xa5\xf8\x80\x6d\x49\xb2\xba\x28\x6d\xd0\xf3\xbc\x69\x50\xc7\x39\x20\x7e\x20\xed\xad\xcc\x78\x6a\xe4\xb0\x74\x74\x2f\x83\x2c\x35\x29\xdb\x3b\x01\x48\x5f\xb1\xaf\xa6\xed\x5d\xa9\x0b\xb5\xda\x26\x23\x6a\x72\x42\xf0\xda\x48\x5e\x95\x6d\xba\xef\x62\x1c\xaa\xfb\xae\xb8\x85\xb0\x29\x2c\x60\x0a\x9b\x26\x3d\xbc\x29\xe3\xd7\x38\x87\xe6\x1b\xca\x10\x94\xc9\x8a\x46\x1b\x4f\xd8\xb4\xc7\x3e\xfc\x06\xb4\xbb\x28\x6f\xda\x77\x3d\x78\xa9\x9b\xf6\xc3\x35\x12\xa9\x4e\x35\xb2\x69\x3f\x6b\xe1\x69\x6b\x4f\xf5\xe1\x92\x3d\x3b\xa2\x90\x03\x6b\x59\x9d\xa8\xbe\x19\xfd\xf9\x79\xc2\x56\x9b\x00\xfa\x3c\x75\xf2\x14\xea\x5d\x5c\x37\x6e\x16\x17\x00\x48\xef\x66\xd1\x4d\xce\x66\x12\x0d\xe5\x5a\x1b\xd0\x13\x36\x98\xcf\x6d\x26\x00\xd7\xbb\x12\x0f\xb1\x20\x11\x62\x27\xd1\xc8\x9b\xc4\xfe\x10\x7c\x1d\x15\xab\xc9\xda\x07\x27\x94\x4d\xb1\xac\x6c\xd3\x5f\xb5\xa4\xb8\x54\x3e\x52\x49\x91\xf6\xf4\x35\x41\x92\x36\xe1\x8c\xb8\x71\x63\xbd\x1c\xd8\x9e\x24\xc9\xb8\x4c\x43\x5b\xc2\xb5\x41\xb0\x25\xc3\xa4\xc5\x8e\xd3\x92\xbc\xa0\x85\x86\x5f\x70\x73\x18\x95\xeb\xfb\xdc\x25\xb7\xf1\x04\x29\xeb\x19\x0b\x17\x72\x8a\x5c\x16\xb1\x83\xf6\x5a\x0a\x65\x4a\x0a\x3b\x54\xcb\xb8\x23\x29\xc6\xc7\x03\xcd\x61\x12\xa4\x5c\x14\x35\xf9\x83\x23\x01\x1c\xd5\x60\x2a\xa5\x77\x74\x21\x35\x44\xe0\x1c\xc8\x38\x9d\xcd\xd9\x60\xd7\xd7\x00\x85\x0b\xff\xa8\x90\xe2\xd6\xc1\xc0\xa5\x29\xc8\xdf\x98\xc8\xd8\xdf\x55\x52\x22\xf3\x7f\x57\x75\xa9\x27\xe5\x6d\xea\x23\xfe\x4a\x6e\x89\x87\xf1\xf7\xa4\x1f\x81\xc6\xca\x3d\xfc\x6d\xef\x23\x03\x65\x2a\x74\xf5\x48\x63\xca\x61\xb1\x50\x63\x81\xac\x97\x4b\xfd\x1e\x97\x99\x50\xb9\x47\x37\xbe\x7b\xd4\x32\x00\x0d\x60\x9c\xb8\xf0\x3d\x7e\x24\xec\x31\x75\xc0\xac\x49\x23\x1a\xb3\x0f\x83\x4d\xde\x63\x19\x7e\x99\x3b\x9b\xec\x6b\x4f\x93\x93\xd9\x57\xdd\x77\xe0\x78\x47\xf4\xeb\x98\x9b\xbf\x88\x79\xc7\xc9\x3b\xdf\xe3\xf5\x3a\xaa\xa4\xa0\x0c\x76\x89\x34\x30\xc0\x54\x21\xbb\x68\x1f\x4f\xd8\x75\xfd\xa0\x9f\xa6\x71\xa2\x54\x40\x17\xbf\xfb\x5d\x2d\x03\x60\x11\xc0\x13\x7e\xc8\xae\x83\x07\xb7\xeb\x5f\x74\x49\xf1\xf3\x00\xe6\xc5\x0d\xfc\x9e\x62\x3f\x59\x43\xfa\xf6\x3e\xb7\x2f\x99\x74\xfd\x08\x55\xe4\x9e\x14\xd6\x70\x4f\xda\x9a\x0f\xab\x37\x4c\xef\x9e\xe6\xbe\x13\x29\x06\xba\x5f\xb7\x85\xbd\xae\x34\x4f\x7b\x05\x07\xb5\xa6\x33\xda\xf6\x90\x82\x87\xb6\x57\x5b\xe9\x06\xec\xd5\x3e\xec\x66\x78\xd0\x7b\xb5\x1e\xe1\x9f\xc0\x7e\x9e\xfe\x4a\x41\x4e\x9a\xcd\xdd\xe2\x61\x5c\xdb\xa2\x26\xc0\xf2\x01\x84\x9e\xa4\x6f\xb1\x9b\xfc\x99\xc2\x0e\x82\x45\x20\x01\x50\xde\xbb\xfd\x13\xed\x5f\x92\x5b\x32\x3d\x94\x55\xc9\x3b\x5d\xd5\x25\xc1\xa8\x3c\x76\xab\x9c\xaa\xf2\xd4\xad\xfe\x80\x2f\x9b\x07\xf8\xb4\xa9\x54\xbf\x3b\xd4\xc2\x3e\xf3\xe2\xa2\x33\xcb\xfa\xbe\xb9\x7d\x96\xf2\x84\x7d\xc2\x6f\xdb\x27\xa6\x3b\xa2\xe1\x35\xcf\x58\xc2\x11\x3b\xe5\xf0\xc0\xf3\x8b\xc9\xf1\x5b\x31\x92\x00\x4d\xdc\xeb\x41\x43\x71\xe8\xb6\x3d\xe1\xd0\x1d\xe2\x7a\x28\x26\xd1\x35\x1c\x7a\x61\xc9\x23\xc5\x5c\x7d\xec\xeb\x68\x1e\x70\x77\xc2\x91\x24\x66\x2f\x0a\x0f\xfa\x48\xda\xdf\x7d\x1f\x89\x6e\xaa\x93\xb5\xe6\x07\xc8\x8d\xc8\x91\xe8\x69\x1e\x69\xe6\x70\xd4\xb4\x49\x09\x47\xad\x1b\xc1\x57\x89\x6e\x7a\x8f\x6a\xee\x03\x1c\x16\xfd\x16\x41\xe8\x5b\x53\x36\x87\x0a\xd7\x0a\x14\x9d\xb0\x03\x1a\xef\xed\x7a\xa1\x5e\xe2\x00\x14\x6c\xcd\xe4\x88\xef\xfa\xa6\x99\x00\xb5\xda\x1a\x79\xc4\x75\xc0\xa6\x1c\x56\xb1\x33\xab\x43\x3a\x42\x99\xc3\xea\x5c\xcf\x87\xb8\x6c\xde\xe1\xf6\x11\xf5\x4e\xbf\x44\x09\x87\xd1\x21\x3c\x6c\x7a\xd0\x70\x18\x3d\xdc\x63\x6a\x1a\x40\x29\x51\x43\x1f\x33\xe7\x27\x9c\x51\x0d\x30\xed\x01\xaf\x27\xc7\x4d\x48\x70\x24\x08\x34\xb6\x06\x68\x1b\x40\xc6\xae\x80\xf9\x62\x40\xd4\xc2\xbd\x9c\xd1\xfa\x19\xce\x08\xf1\x39\xdf\xdd\x9f\x71\xba\xe2\x3d\x25\xf2\xfd\x32\xa4\xd8\x25\x0d\xa7\x58\x71\x49\x49\x4e\x00\xee\xac\x4c\x8d\xc3\x25\x35\x5f\x46\x37\xdc\xcd\xa9\x1b\x2f\xe2\xd4\xe3\x24\xf8\x82\x5f\x9b\xc2\xa9\xe9\x3d\xd0\xd3\x7d\xd6\xb3\xa6\x54\xef\x70\xd6\xec\xee\x08\x89\x27\xbc\x76\xf5\xac\xe5\xd5\x3e\xa4\x64\x0b\x67\x6d\xce\x6f\x35\x53\x6f\xe8\xed\xd5\xc8\x67\x1d\xe1\x9c\x87\x84\x13\x72\x4f\x04\x8d\xe8\xdb\x9a\x39\xbf\xb8\x96\x83\x0c\x4c\x06\xa6\xe7\x34\x3e\x98\x73\xf6\xcf\x13\x74\x25\x78\x82\x6e\x12\x93\x43\x8f\x2d\x75\x63\x38\xc8\x4a\xad\x5e\x93\x24\xe8\x56\x13\x23\x2c\x3d\xc0\x6a\x19\x41\x53\xc2\xef\x6b\xa7\xed\x04\xa5\x05\x5a\xc4\x0b\x82\x57\x8a\xc8\x67\xb8\x3f\xac\x09\x9a\xc9\x8f\x44\x73\x3c\xb4\xc0\x80\xbc\x14\xdd\x20\xcd\x39\x32\x72\x71\x12\xef\x03\x53\xe4\x2c\x9b\xbe\x95\x74\xab\x34\x37\x04\xc2\x6c\x7b\x5d\x7a\x2c\x99\x9b\xac\x2b\x8f\x07\xa4\xed\x70\x79\x48\x8d\x69\xcb\x6f\x3d\x34\x9f\xe6\x56\xc1\x16\x09\xa3\xfb\x02\xaa\x77\x7d\x3b\x31\x44\x5e\x7d\x7e\x66\x1c\x6a\xf9\x1d\xff\x9b\x80\x00\x59\xff\xcd\x3a\x4d\x82\x16\x46\x18\x7e\x89\x6f\x61\xe8\xff\x0b\x8c\x31\x40\x96\xee\x2a\x92\x1b\x40\x2c\xb1\xbd\xbd\xe0\x48\xbe\x54\xa5\x09\xd6\xf2\x5f\x23\xaa\xe5\xb5\x77\x20\xca\x50\xf6\xdc\xe3\x98\x7e\xfc\x65\x17\xd3\x6a\x20\xdc\x95\xf5\xcb\xe0\x22\xa5\xae\x3e\x5f\x97\xd5\x79\xed\xb8\x5b\xce\xd5\xf5\x28\xba\xeb\x1a\xbd\x03\x55\x14\xa9\x61\xd3\x33\x05\x5a\xfa\x5c\xe8\xe6\x39\xcb\x03\x8e\xce\xe1\xbd\xc6\xff\x91\x6b\x5d\xd0\xf0\xcb\x21\x82\xa4\xff\x55\xe2\x96\x7f\xbb\x98\x34\xdd\x82\x5a\x2d\x6b\x8a\x20\x4c\x7a\x43\xf0\xc4\xfd\x9b\xe9\x70\x61\x46\xf4\xff\x84\xef\x08\xd5\xf2\x1d\x8f\x83\x28\x7f\xb5\xb8\x7e\x48\xce\x32\x9e\xf0\x2d\x91\x46\xe3\x5b\x3d\x10\xfe\xae\x70\x1f\xbe\x2b\x4c\xca\x77\xbd\xc0\xf4\x77\x7d\x20\xfa\xdf\xf5\x61\xc7\xb9\x68\xc2\x7e\xbf\xe7\xf6\xfa\xf3\xdf\x13\x6f\xe1\x7b\xe6\xa5\x02\x36\xd0\x05\xc7\xf9\x3d\x2f\x29\x38\xeb\x4f\x4c\xa2\x5b\xed\x6b\x6d\xba\x86\x0f\xbc\xa8\x8f\x3c\xe1\xa3\x65\x23\xf8\xcd\x4c\x80\x4e\x4f\xf8\xa4\x39\x7e\x9e\xf0\x29\x15\x8a\xef\x53\xea\x18\x44\xd4\x8e\xc0\x49\xb6\xc3\x33\x55\x28\x95\xf0\xa9\xf5\xf3\x84\x14\xb3\x84\x14\x0b\x49\x3b\x40\xfb\x59\xf2\x59\xa4\x78\x41\x80\x93\xe7\x5c\x12\xa2\xbd\x24\x91\x11\x47\x92\xb8\xb3\x54\x42\x92\x7d\x04\x7f\x7f\x49\x0e\x1d\xcc\x42\x05\xfa\x96\x0e\x7f\x59\x4d\xea\xe9\x8a\xa4\x1f\xa8\xa2\xa4\x59\xbd\x3a\x37\x02\xce\xab\x05\xd1\x56\x62\x2c\x92\xf4\x0f\x7c\x0c\x18\x2f\x0e\xbb\x1c\x70\xec\x85\x86\x1a\xc7\x49\x35\x9d\x6a\x71\x80\xb6\x8a\x79\x6a\x39\xbe\x10\xd4\x84\x54\x61\x96\x52\x6d\x08\x95\x31\xa6\x8f\x80\xb8\x3f\xa4\x7a\x39\x3b\xd5\xa7\xac\x34\xaf\x50\x9a\x1e\x72\xa4\x09\x9d\x95\x18\x75\xb2\x65\x66\x65\x96\x22\xf1\xc2\xa8\x5e\xbc\xf8\x77\xba\x9e\xc9\x71\xc3\x4f\xbe\x66\x0b\x39\xba\x9c\x43\x60\x73\xfc\x16\xea\xba\x1c\xbf\xab\x85\x1c\x73\xae\xe3\x0c\x39\x9a\xa9\x60\x88\x5d\x92\x7c\x02\x7f\xee\x99\x76\x37\x64\x0f\x3b\xb2\xc4\xf4\x00\x16\x80\xee\xf4\x00\xdc\xa8\x36\x80\xeb\xca\xfe\x9b\xce\x1c\xb2\xc8\x27\x64\xa1\xfe\xcf\x92\xea\x86\x95\x25\xbf\x0a\x3f\x4b\x89\x2b\xf6\x93\xc5\x0c\x93\xc0\xfc\x67\x19\x31\x25\xf4\xd3\xb4\x01\x7c\x9e\x90\x21\x59\x30\x1d\x59\x8b\xbe\x0c\x69\x01\xff\x8a\x67\x90\xb5\xcb\x3b\x23\xa8\xc4\x9a\x63\xf2\x90\x33\xfd\xd6\x4c\xf7\x33\xd7\x4d\xac\x10\xf5\x11\x72\x55\xc2\xf2\x1e\x5e\xae\x65\x9c\x18\x5b\xcb\xcc\x1e\xbf\x82\x51\xc2\xa1\x3b\x8f\xa7\xbe\x2e\xb3\x87\x37\xce\x52\x9d\x65\x44\xbf\x8e\x3c\xe1\x2d\xe5\xb9\xef\x70\x19\xf2\x4c\x43\x99\xa2\x61\xa1\x6c\xb1\x48\xc8\xd3\xb0\xa1\x09\xb7\x37\x4f\xf2\x36\x21\x7a\xf9\x81\xf3\xca\xbb\x2b\x11\x07\x5f\xa2\x59\xbd\x43\xa1\x09\xa7\xf7\x13\x53\x28\x71\x1e\x27\xcb\xd3\xf5\x34\x45\xb4\xc0\x9f\x29\xb2\x4a\xef\xb0\x2c\x85\x8e\x4d\x91\xc3\x8d\x4b\x91\x03\x41\x9d\x6c\xa4\xf4\xa0\xa6\x29\xa2\xc7\xf9\x26\xf8\x8a\xd8\xc5\x75\xe5\x0e\x45\xfe\x8e\x50\x74\x95\x50\x14\xc6\xbf\x28\xcd\x6c\x51\x78\x1d\xa5\xae\x63\x52\xad\x95\xaa\x1d\xe5\xd2\x47\x6d\xa1\xb8\xb6\x2d\xd5\x6d\x74\x61\xd4\x0e\x78\xc3\x93\x46\x3f\xf8\x05\xa5\x7a\xf2\xb7\xcc\xbc\x84\x32\x87\xbd\xa9\xe6\x32\xa1\x20\xea\x22\x1b\x13\x12\x75\x91\x2e\xa1\x2e\x7c\x3b\x75\xf1\xa0\xaf\xae\x6b\xec\xae\x43\xeb\xb6\xe1\xb7\xc4\x94\x42\xdd\xf7\x05\x7b\xaf\xfb\x6f\x82\xa3\xee\x6f\x02\xb5\x6a\x7a\x42\x4d\x1b\x7e\x5f\x7b\xec\xa7\xdb\xca\x0a\xef\xa6\x36\x29\xa1\xb6\xc1\xb4\x09\xb1\xdb\xff\xda\x26\xad\x4f\x35\x0f\x86\x3d\x21\x65\x20\x8e\x58\xd0\x01\x4e\x3b\x99\x30\x06\x05\xd5\x0a\xd4\x2e\xa5\x22\xd4\x39\x3c\xfa\xae\x73\x24\xf8\xa0\xfd\x24\xf9\xbc\x75\xbf\xc9\xab\x3a\x07\x4c\x0a\x75\x57\xbd\x30\xd9\x25\xf6\x66\xc5\x40\x6e\x53\x88\xbf\xeb\xf3\xd6\x24\x6f\x6c\x31\x45\x3f\xe1\x16\x1d\x0c\x7b\x02\x0c\x6f\x4a\x92\x40\x30\x61\xfc\xeb\x22\xfc\x27\xe7\xd7\x10\xf9\x01\xa0\x9b\xbc\x7e\x73\x93\xd6\xc0\xfd\xaf\xc7\x00\x3c\xe1\xc2\xc0\xf0\xf2\xa5\x36\x81\xac\x36\xb1\xce\x93\x87\x43\xe1\xe1\x7d\x93\xce\x86\xde\x7f\x8f\xae\x09\x2e\xb1\x9d\xb5\x00\x3e\x9d\xc2\xd9\x54\x8c\xef\xbb\xa9\xe7\x2e\x5a\x8a\x5a\x1c\x52\x34\xdb\x9b\x84\x68\x29\x3e\xde\x2e\xb1\xc3\x3c\x91\x00\x2b\xc0\x9c\x01\x9a\x0e\xb0\x9f\xa1\xc1\xca\xa1\x11\x98\x6f\xb0\x55\xe6\x6d\x5a\x7d\xd3\x10\xaf\x97\x43\xe4\x9c\xd4\x6a\xa1\xe1\x7e\x23\xb0\x8d\xf4\x84\x56\x39\x57\xf7\x9c\xd7\x3f\x2f\xa8\xbd\x52\xdc\x10\x62\x92\x29\x10\xbe\xcc\x3f\xf7\xa8\x99\xb8\xf5\x84\x5b\xc4\x13\xa1\xe5\x3f\x7e\xeb\x78\x18\x26\xd7\x5b\xd6\x55\x1e\xa2\x4f\x22\xce\x10\x2a\x60\x01\x84\x99\x23\x71\xf1\xae\x4c\x7f\x00\x5f\x1f\xab\x59\xdd\x5e\x6f\xa4\x59\xdd\x5f\x37\x18\xe4\xec\xc4\x4d\x0c\x68\x6e\xa1\xd9\x64\x82\xad\x21\x12\x9c\xe0\x72\x96\xe3\xad\x78\x02\x5d\xb5\x46\xbd\xd3\xf0\xf8\xb6\xd0\xa6\x87\x36\xc4\x38\xe4\x3f\x93\xe9\x87\x3f\x33\x26\xd7\x77\x7f\x66\x34\x93\x44\xbb\x06\x7a\xf0\x59\xfc\x99\xcc\x37\xfd\x99\xd8\xf2\xeb\x20\xfe\x99\xba\x7e\x1c\x7e\xbd\xb9\xb7\x3f\x53\x05\x93\xe1\xb2\xfe\xa5\x68\xfe\x4c\x85\x4c\xfd\x99\xea\x59\x93\x3f\x53\x7f\x7e\x78\xa4\x16\x37\x46\x9c\xe6\xce\x8e\x45\xf8\x49\x16\x4d\x82\xc5\x7e\x06\x8b\x77\x30\x89\xdb\x03\x98\x08\x5c\xf2\x4c\x62\x7f\x99\x30\x4f\x88\x9a\xac\x1f\xde\x97\xc9\x3a\x62\x39\x78\xe9\x26\x3b\xb4\x20\xb6\x69\x74\x10\x4c\x7e\x1b\x92\xbe\xa3\x93\xca\x85\xa5\x25\x47\xfb\xbc\x75\xb9\x5a\xe7\x73\x01\x39\x50\xd3\x60\x1b\xb8\x4c\x9b\x89\x4a\x07\x91\x3f\x3f\xfe\xc1\x97\x93\x75\x78\xf7\xff\xbb\x75\x26\x57\x4d\x97\x2f\x7d\x47\xe3\xcb\x37\x5d\xcf\xe0\x49\x06\xcf\x04\x40\x15\x63\x0a\x6d\x12\xac\x46\xa6\x96\xac\x2e\xb3\x8f\x60\xb5\x3f\xc1\xea\xeb\x19\x5a\x1d\x43\x4a\xb0\x3a\x8f\x33\x30\xfa\x0c\x56\x6f\x1c\x0c\x7d\x05\xc4\x9b\x78\xd6\x46\x5b\x04\x91\x08\x36\xcb\x06\x47\xcb\x98\x17\x30\x5a\x12\xa3\xed\xe9\x71\x0b\x3d\xee\x12\x7a\x4c\x2c\x66\x90\x88\x33\x3b\x8c\x55\x8f\xb6\xfa\x0b\xed\xd1\xb6\x5a\x48\x0c\xed\x14\x8e\xbe\xc2\x17\xe8\x6b\xb4\x55\x88\xbc\xce\x50\x27\x7c\xa0\x7d\x3d\x6b\x62\xa2\xb2\x33\xd9\x0c\x27\x24\xf4\xb5\x5a\xc1\x09\xf5\xd5\x22\x9d\x5c\x10\x77\x79\x42\x97\xb5\x96\xed\x45\x5f\x4c\x1f\x39\xcd\x39\x65\x35\x19\xa1\x4b\xda\xbf\x7e\xbf\xfd\xb1\xc0\xfb\x2b\xde\x02\x95\x4b\x49\x7b\xed\x77\x17\x68\xc7\x21\xa1\x8b\x49\x21\xe2\x8b\xec\x62\xb8\x45\x65\xd5\xc5\x34\x08\xb1\x84\x7e\x46\x78\xad\xfd\x8c\x5b\xbd\x1d\x13\xd2\x83\xe1\xa7\x47\xb2\x0e\x82\x22\xd6\xcf\x68\x2d\xf4\x13\x11\x1a\x8d\x69\x3f\x15\x5b\x39\x2b\xfd\x74\x62\xaf\xae\x1b\x67\xaa\x36\x1c\xba\x4b\xd8\x4f\x5f\xc6\x10\x52\xe2\x6d\x76\xbc\xa3\xae\xb4\x43\x9d\xce\x4f\x67\x86\xa8\x6b\xba\xc4\x02\x14\x2e\xa4\xb6\x33\xd5\xed\xc8\xdf\x41\x57\x3f\x55\x2d\x07\x1b\xde\x37\xdd\x3f\x92\x84\x67\xf1\x51\xdc\x7e\x4f\xfc\x26\xd2\xdd\x29\xee\x49\x33\x01\x6b\xdc\x1e\x74\x6e\x35\x51\x53\xf4\x0c\x33\xdb\x73\x04\xd7\xb9\x1e\x07\xaa\x2a\x3c\xc6\x9e\xe7\x11\x7a\x61\x36\xa2\x97\xf7\xeb\x45\x2f\x12\xc1\x71\x81\x6b\xec\x15\xcc\x64\xf7\xba\xfa\x1b\xeb\x75\x1f\xa1\xfb\x3c\x35\xe9\x16\x7a\x85\x4f\x1b\xa8\x61\x7a\x6d\x27\xf7\xc2\xb0\xae\x57\x5c\x47\x85\x13\xc4\x7d\xd5\x99\x5e\xcc\x09\xa7\x85\xde\xa2\x7d\x7c\x99\xc6\x44\x03\x1e\x9f\xcb\x18\x5e\x61\xf4\x2f\x70\xbd\xd1\x1d\xea\xed\x14\xff\x28\xd3\x1b\x02\x9e\xde\x74\xdf\x89\x8c\xae\x7f\x6f\xea\x59\x92\xde\x92\xf0\xab\x61\x6f\xd5\xad\x0b\x08\xef\x82\xd3\x81\x1e\x14\x20\xdf\xea\x9f\xa9\xc7\x81\xfb\x71\xc5\xdc\x47\xc4\x50\x38\x84\x42\xec\x75\x90\x61\x16\x8d\x02\x3f\xa2\x7d\x08\x31\xe1\x10\x49\x84\x2d\x60\xf3\x98\x74\xe8\xbe\x03\xe6\xdf\xaf\x99\x8c\x84\xd0\x50\x0d\x57\x35\x2c\xfa\xf5\x0d\x8b\xe5\xc0\x0a\xa6\x2b\x8b\x4c\x61\x90\x70\xb1\x1b\xd0\x27\x58\xd9\x2a\x8b\xd3\x83\xdf\x3e\x66\x29\xde\x61\x36\x29\xdb\x5b\xd7\xc8\xfb\x34\x9c\xd7\x78\x98\x2c\xed\x73\xd9\x26\x66\x98\x4b\x56\x0f\x09\xfe\x1b\x13\xf7\xb9\x20\x20\xea\x73\x99\xb6\xc4\x12\xfa\x64\x6e\xa8\xcf\x83\x8f\x96\x0a\xa7\x4f\xd8\x28\xc2\xc5\xd1\xeb\xa3\x91\x46\xa8\xd9\x67\x83\x51\xf6\xc9\x4d\xbe\xf6\x4a\x17\xb2\x4f\x83\x79\x24\x85\x5b\x73\xd6\x2f\x49\x78\xd1\x37\x03\x8f\x7e\x0b\xb4\xc2\x8d\x3a\x7f\x7f\xb7\x42\xc4\x9e\xdc\xe2\x38\xf9\x71\x61\xc4\xf4\x71\x4f\x18\x81\x45\x18\xd0\x71\x23\x16\xfc\xdc\x17\x1e\x90\xee\x41\x3f\x7b\xc4\xf7\x73\xeb\x88\x73\x84\x21\x7e\x5a\x43\x84\x89\xa5\x21\xb9\x0d\x27\xca\x26\x06\xd4\x25\x0c\x46\xf4\x43\xde\x64\x15\x09\x08\xe1\x10\xce\x78\xc6\xf2\x81\x5c\x8d\x33\x8e\x30\x4e\x38\xb7\xe3\xc4\x23\x1f\xa7\x16\x00\xdb\x08\xbd\xb3\xf6\x30\x4e\x7e\xfd\x04\xa2\x8a\x1f\x67\xe5\x18\x14\x86\xcf\x04\x43\xc8\x3f\x16\x8c\xd3\xe8\xa6\x8f\x73\x82\x23\x32\xab\xdb\x13\x06\xc5\x03\x7a\x30\xa1\x40\x6d\x3c\xa0\x98\x06\xed\xd4\xa8\x56\xc2\xa8\x50\x0b\xc3\xd7\xb0\x78\x80\x65\x73\xb9\x1d\x16\xa7\x7f\x91\x1a\x30\xaf\x54\xb6\xc3\x24\xbf\x92\x32\x4c\x7f\xcd\x2a\x04\xef\xf3\x00\xed\x7c\x89\xc3\x34\x03\x30\xf7\x80\xb8\x6a\x49\x1c\x3b\x05\xc0\x03\x7a\x12\x6f\x68\x30\xdc\x0e\x0d\x9b\xfe\xb1\x63\x30\x5b\x3a\xa6\x2d\xee\x8a\x8f\x5b\x4b\x98\x78\x60\x08\xa6\x18\xce\xcf\xf2\x9f\x3f\xf3\xcc\xe2\xff\xc9\x98\xe5\xbf\x1f\xbe\x59\xca\xb5\x90\xf8\xfd\x44\xcd\x13\xea\xfe\xb7\x99\x59\xf8\x19\xc9\xa9\x4b\xd0\xf1\xdf\x77\x0c\x90\x03\xb0\xa6\xcd\x8b\xff\x3e\x02\xcc\xe2\xe9\xd7\x59\x7e\x53\xa8\xb3\xe8\x5e\x2d\x03\xff\x4b\xa9\xcd\xc2\x3f\xbb\xcc\xa2\xbe\xd4\x07\x62\x03\xcc\x6c\xc7\x2c\x6f\x12\x63\x16\x4f\x2d\xcc\xe2\xe1\xfd\x2c\xbf\x51\xde\x2c\xff\x1c\xe1\x59\xfe\xef\xe9\xcc\x42\xe7\x60\x16\x9b\x38\x8f\xf2\x6b\xf1\x66\xe9\xb8\x6e\xaf\x1b\xf4\x91\x66\xe9\xb3\x1c\x61\x16\x4a\xc3\x2c\xaf\x24\xcc\x32\x18\xe6\x03\x93\x43\x7e\x26\x01\x9a\x5c\xf6\x96\xb4\x42\x13\xcf\x72\xab\x24\x8e\xbc\x35\xf1\x62\x41\xf9\xa7\x9c\x59\x6e\x53\xba\x20\xb3\x31\xfa\x9a\xcd\xfd\x98\xd9\xba\x8c\xe0\x7a\x60\x76\xbf\x08\xce\xef\x0e\xd4\x7c\xbf\x46\xcf\xa1\x9e\x4e\xe5\x5f\x1b\xc2\x15\x99\x80\xbb\xe2\x31\x25\x5c\x88\x0f\xae\x08\xcb\x70\xc5\x34\x39\xc7\x45\x4f\xf0\x8a\xa6\x6f\xb1\x8f\x70\x49\xba\x64\x3c\xe1\xe2\x47\x0c\x1c\xf3\x25\xe5\xe0\x32\x97\xe0\xf9\x79\x57\x5d\x98\xae\xbd\x5e\xd5\x71\xe9\x3a\xde\x6f\x0d\x97\x1e\x9e\x08\xbb\xf4\xf0\x4f\x18\x97\xa6\xf4\x86\xf1\x97\x7a\x16\x92\x18\xc8\xe8\x4e\x13\x7b\xbb\xfb\x75\x17\x3f\x41\x5f\x7a\xc5\xdf\xf9\x2f\x70\x5a\xff\xa5\x67\xee\x58\xc2\x0d\xdb\xf8\x11\xe0\x4c\xf0\x9b\x83\xba\xa3\x35\x47\x4f\xb8\xe1\xef\x81\xfb\x1b\x2e\xd1\x4b\x20\xf0\xa3\x2f\x4c\x1a\x9d\xae\x27\xdc\x12\x3f\x00\x94\xbe\x9b\x9f\x8b\x6f\x59\x16\xf4\x12\xc1\x0f\x21\xc6\x4d\x1b\x7f\x23\xc6\x1f\xc4\x86\x66\xbf\xdb\x5b\x52\xfa\xda\xea\xca\x1c\x07\x87\xa5\xf4\x75\x58\xad\xf9\xb7\xf0\x2f\x97\xcd\x52\xd2\xe1\x04\xbf\x46\x90\xaa\xfb\xee\xc4\xa8\x5f\x5b\x75\xf2\x86\x4a\xb9\x65\x84\xfb\x84\x67\x7b\x9f\x9a\x3d\x08\xbc\x4f\xb5\x74\x6b\xd9\x40\xf5\x46\xed\x7a\x9f\x35\x09\xa0\x7f\x64\xb9\x95\xce\xd2\xad\x58\x40\x37\xf9\x92\x87\x45\x1a\xd6\x5b\x13\x68\x97\xc2\x5b\xf9\xad\x16\xd3\x79\x11\xae\xea\xed\x9e\xf3\xad\x38\x0b\xed\x98\x03\xf1\xd7\x5d\xf9\xbf\xa2\xbb\xfa\xc1\x56\x3c\x7f\xa7\x2a\x4c\xd4\x5d\xeb\x0f\x5a\x61\xec\xee\x6a\x69\x63\x5f\x6c\xa2\x7a\x2e\x8c\x98\x8e\xc9\x8d\x18\x80\x70\x04\xea\x2d\xca\x31\xa9\xfb\xd4\x24\xe1\x55\x66\xb7\x45\xd7\x62\xb7\x89\xff\x79\x0b\x2f\x85\x7f\x26\xb9\x69\x7f\x6f\x7b\xc2\x13\x6f\x5a\xde\xc7\x3f\xe1\x3f\x9e\x1c\x73\x84\xd7\xfc\x54\xbc\xdc\xa7\xbe\x0a\xf1\xe1\x87\xa1\x9f\x58\x9e\xf0\xf3\x66\x7b\x7f\x68\x58\x7e\xf4\xf8\x89\xc7\xff\x02\x00\x00\xff\xff\x49\xf0\xba\x68\xf9\x28\x00\x00"),
		},
		"/animals.txt": &vfsgen۰CompressedFileInfo{
			name:             "animals.txt",
			modTime:          time.Date(2019, 11, 27, 10, 4, 8, 581479822, time.UTC),
			uncompressedSize: 3870,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x57\x4b\x96\xf5\x2c\xab\xee\x33\x4b\xa2\x44\xa9\x20\xf8\xa2\xee\x54\x6a\xf4\x67\x61\xf6\x77\xd6\xdf\x08\x3c\x2a\xde\x90\x5b\x10\x3d\x7f\xd0\x2f\xc0\x03\xc5\x94\x00\xe5\xc0\xe9\x36\x06\xa0\x14\x8c\xb6\x70\xc1\x69\x0e\x28\x1d\x13\x02\x36\xa3\x23\x58\xaf\x7c\x30\x2a\xa0\x62\x32\xcd\x08\xa8\xd4\xf6\x12\x5a\x48\x4e\x1e\xf5\x45\xa3\xee\xf5\xb5\x98\xda\x15\x52\xdc\x50\x00\xd5\x24\x44\x67\x7c\x84\x93\x7c\x03\xb1\x4e\x80\x1d\xa7\x0d\x5c\xbe\x06\xe0\x6e\x57\xce\x80\x8e\xa9\xea\x06\xa9\x22\x59\x9f\xe4\xcf\x2f\xa0\x37\xcc\x2c\x62\x80\x3e\xab\x5b\xb7\x90\x98\x6c\x19\xd3\x7c\x64\x00\x8e\x0e\x07\x1e\x66\x0a\x07\xe6\x42\x0e\x07\x6a\xe6\x64\x36\xe1\x40\x57\x4c\x42\x01\x1c\xd3\xca\x08\x07\x0e\x16\x1e\x57\x80\x01\x07\x4e\x38\x08\x4b\x88\x10\x7a\x90\x4f\xac\x40\xf9\x58\x05\x0e\x8a\x6e\x9a\x31\xca\x72\x3c\x70\xb0\xce\xe5\xa6\x05\x0e\xf6\x0c\x07\x8f\xd8\x95\x3f\x28\x1f\x82\x43\x30\x5d\x6f\xbf\x98\xe5\x6a\x4b\x37\xbc\xb7\xae\x0e\x59\xf4\x0e\x1a\xc2\x61\x47\x8a\x9d\x4d\x8b\x05\xb5\x23\xd8\x6f\xec\xbc\xce\x13\xc5\x60\x6f\xbf\x44\x36\xc9\xf6\x36\x4e\xdf\xa0\x1d\x42\xfb\x68\x6b\xd4\x03\xe3\x58\x6b\x4e\xf2\x53\x1e\x48\xc8\x0d\x15\x12\x36\x12\x48\xa8\xe8\xd1\xd7\x9f\x83\x1c\x21\x85\x8a\x31\xfa\x3d\xb3\xbe\x80\x0f\x5b\xc1\x95\x3f\xb6\x45\xc6\xb0\xfb\x9d\x35\xe3\xc3\x66\x4b\x37\x22\xef\x2c\x82\x1e\x38\x34\x92\xf0\xf3\x40\x22\x9d\xdc\x29\x13\x24\xea\x15\x25\x5e\x27\xd5\xd8\x9d\x4c\x21\x55\xa2\x89\x15\x52\xe5\x74\x61\x26\x7a\x11\xc5\x08\xd7\x85\x75\x61\xa0\xc6\x1d\xf5\xef\x1d\xd5\x54\x63\x97\x80\xbd\x2d\xbd\x02\xcc\xbd\x94\x3f\x03\x85\x07\x24\x4e\x98\x11\x92\x60\x83\x24\x76\xeb\x56\x6f\x32\x9c\x0c\xc9\x8e\xb8\x84\xa5\x0b\xa7\xd9\x06\xe4\xa3\xa3\x72\xa8\xc3\xd2\xe5\x86\x29\x84\x33\x24\x23\xc1\x84\x3a\xa3\x29\xc2\x14\x6c\x95\x98\xa4\x5b\x44\xc7\x74\x4e\xe1\x19\xc9\x3a\xed\x8b\x99\x87\xd2\x6c\x4e\xd3\x66\x6b\xcf\x5c\x25\x34\x62\x37\x24\x7b\x6c\xc6\x22\x4f\x5f\x90\x1c\x8f\x20\x4a\x41\x9f\xf7\x84\x1e\x57\x9f\x90\xdc\x92\x65\x0e\x0d\x7a\xcc\xf3\x35\x26\x26\x8a\x57\x5b\xa1\xd8\x57\x78\xcd\xdb\xbc\x41\xc6\x54\x47\x0d\x43\xca\x28\x0d\x67\x38\x64\xc6\x36\x48\xe2\xb5\x33\x91\x43\xae\xe1\x69\x99\x71\x5a\x83\xcc\x61\x51\x99\xd5\x5c\x79\x6c\x10\x8e\x06\x99\xbb\x58\xb6\xb4\x06\x64\x3b\xc8\xc3\x48\xb2\x65\x83\x30\xad\x6c\x65\x6f\x9a\x4d\x7a\xe5\x18\xd0\x8b\x1e\xc8\xe6\x41\x3e\x04\xd9\xb1\x98\x7e\xd9\xde\xd8\xad\x51\x0e\x33\xc9\x2b\x5d\x90\x57\x09\x9f\x78\xfd\x88\xc2\x4f\xf7\xe1\x09\xfd\xe6\x02\x94\x2a\x67\xc5\xcd\xd5\x32\x79\x1b\x40\x24\x40\xa5\x00\x15\xa7\x09\x74\xa1\x4c\xcc\x24\x13\x81\x04\x35\x03\x49\x98\x93\x4e\x20\xb9\x80\xda\x02\xf2\xc6\x4a\x70\xa2\x24\x53\x38\xc9\x63\xde\x19\xe6\x02\x27\x3b\x85\x68\xf0\x38\xdc\xbe\xcb\x29\xd8\xb6\x2e\x4e\xc1\x57\x95\xa7\x10\xc2\x29\xe1\x95\xe4\xb0\x05\x6d\x0c\x84\xd3\x6e\x81\xd3\x7e\x61\x7b\xd7\xe9\x8b\x67\x44\x85\x0d\xb6\xd4\x2a\x0b\x0a\x0a\x16\x1b\x10\xaf\x5d\x70\xcc\x37\x0e\x15\xfc\x30\x0a\x14\xfc\x23\x11\x82\x42\xe9\x32\x28\xe4\x07\x0b\x14\x3e\x0e\x53\x28\xec\x78\x9e\x04\x45\x71\x42\xd1\x05\xc5\x02\x98\xe4\xf7\xf0\x2f\x1a\x01\x6c\x10\x14\xeb\x95\x1c\x8a\xf9\xf6\x82\xe2\x38\x46\xb5\xde\xa3\xcf\xe9\x79\x63\x4a\xf1\x60\xd5\xca\x46\xef\x98\xad\x98\xed\xeb\x80\x12\x51\xa3\x62\x1b\x11\x73\x2b\x3a\x41\xc5\x99\x2a\xcd\xbd\x4d\xc5\xfb\x82\x4a\xb9\x50\xcc\xaf\xa4\x50\xc9\x6d\x53\x67\x2d\x50\xb9\x77\x7b\x69\xb7\x89\x6d\x0d\xa8\xe6\x4a\x33\xd8\x20\x78\x4f\x50\x63\xbb\x50\x4e\xb5\x5b\xc8\xdb\x6b\x30\x75\x85\x59\xd5\xd5\x42\xf3\x3b\xd6\xd5\x35\xae\x07\xea\x43\x8a\x50\x1f\xc7\x5f\xe0\x83\x07\x70\xaa\xb3\x3e\xaf\x89\x72\x59\xa8\x08\xdc\xb0\x18\x70\xeb\x28\x08\xac\x83\xd2\xfc\xb2\x08\x4c\x34\x80\xf5\x43\x3e\xe9\x70\x9c\xd1\x1a\xfb\x01\x7e\xb0\x2c\x74\xf8\xc1\x0f\x09\x2b\xc2\x0f\x3e\xf0\x43\x22\xaf\xc7\xfd\x90\x47\xb0\xfd\x31\x7a\xe0\x67\x69\xb2\xa0\x14\x91\xf5\xc2\x0b\xbb\xc1\x85\x5a\xd0\x2d\x80\xe3\xb5\x04\x2e\x9c\x4f\xe6\x0c\x17\xeb\x85\x3f\xb6\xe0\xe2\x9b\xe1\xb2\x38\xd4\x65\x76\xe1\xb1\xdc\x11\xae\x78\x1d\x10\xcc\x4f\x2c\x16\x96\xd1\xcc\x7b\x05\x89\x14\xe7\x6f\x40\x12\xf4\x0f\x82\x50\x09\xf5\x81\xd0\x56\x4a\xf0\xe5\x20\x64\x1d\x3d\x83\x70\x22\x10\x6e\x9d\x26\x08\x9b\x82\xf0\xdf\xee\x17\x6c\x08\x62\xc7\x7e\x43\x09\xb7\x9d\x20\xa6\xdf\xb5\x22\xd5\x89\x39\x5f\x14\xf3\xcc\x79\x80\xec\xf7\x97\x47\x7f\xa1\x61\xc2\x7f\x8b\x36\xbf\x83\x5e\xe4\x24\xd0\xb0\x14\x9b\xd0\x50\xb0\xad\x19\xc3\xed\xc0\xa0\x91\xad\x83\xd9\xac\xd0\x50\x71\x52\x0c\x6a\xde\x77\x6c\xa8\x91\x0e\x82\x4f\xdc\x94\xf3\xcb\x06\x34\xf4\x66\x83\xe6\x0b\x36\x1b\xab\xf3\x5e\x6e\x4c\x3e\xcf\xcd\x2d\x9b\x42\x23\xcc\x76\x4b\x14\x0a\x8d\x50\xb6\x3f\x36\x22\xbf\x70\x42\xa3\x82\xf2\x95\xf2\x82\x3a\xc8\xa1\x85\x62\x1a\xe7\x12\xb4\xc4\x93\x43\x63\x91\x37\xd3\x34\xd6\x2b\x88\xda\x0d\xcd\x10\x9a\xa5\x84\x83\x35\xc0\xf5\x9f\xe5\xb5\x08\x8a\x41\x3c\xb6\x30\x91\x35\x2e\x68\x91\x6f\xc3\xcd\xbe\x06\xdb\x4c\x6d\x3a\xb5\xdd\xb1\x75\xdd\xbe\xc3\xe3\xdf\xe2\x69\xf0\x2a\x65\xeb\xb6\xad\xdc\x57\xef\x0f\xb4\xe5\x4e\x12\xd7\x5e\xe3\xda\xab\xaf\x31\x48\x46\xf0\x49\xc2\x79\x80\xe2\x81\x2e\x74\x81\xa2\xe6\x05\x8a\x7e\x57\x14\x50\x5c\x93\x65\x0d\x50\x52\x02\xa5\x31\x41\xe9\x9e\xa0\x5c\xea\x64\x2d\x28\x04\xba\x5a\x44\x1f\x5d\x73\x7b\x6d\x00\x67\x04\x7d\x5a\xaf\x60\x89\xc4\x26\x58\x9a\xd6\xd7\x00\xbb\xb0\x33\x98\x62\x14\x3d\xd6\x6d\x8c\xd5\xc0\x1c\xb5\xac\x89\x0a\xe6\x09\xc1\x9c\x43\x13\xe6\xd6\x49\xb3\x49\xf4\x3c\xbf\x60\x3b\xb9\x55\xb0\x28\x1c\x20\x22\xa0\xfd\x82\xfd\x76\x0a\x6b\x01\x7b\xb6\x32\x3a\x46\xd9\xd7\xa3\xb4\x13\xd6\x00\xb3\xee\x6e\xc7\x6d\x7a\x1d\xdd\x6d\x42\x27\x8c\xd4\x1a\x7c\xc7\xd2\x4e\xf1\x0a\x34\x08\x3a\x09\x27\x54\xe8\xa4\x65\xc5\x0a\x34\xc3\x14\x3b\x97\xf8\xa2\x42\xe8\x7c\x21\x74\x56\x8d\xa7\x1d\xd0\xd9\x51\x2b\x42\x17\xd4\x2b\xd2\x7e\x17\x9c\xcf\x41\xaf\x7d\xec\x46\x5c\xbc\x8b\x45\x95\xd6\x23\x6f\xdf\x56\xa0\x5b\x23\x47\x8d\xe4\xd8\xad\x75\x54\x83\x6e\x96\x85\xa0\x9b\xa7\xd5\x23\x73\x74\xf3\x6e\x1c\x87\x7a\xf5\xd4\x9d\xdb\x0e\x25\xdd\xff\x73\xac\xee\x36\xed\xcf\x62\x95\x75\x9e\xe4\xdb\x8f\x03\xc6\xd1\x57\x81\xbe\x1a\x42\x5f\x7d\x93\x77\xf0\x99\xd5\x14\xfe\x2d\x2c\x05\x83\xb1\xc0\xbf\x45\xf3\x0f\x83\xdb\x75\x45\xa7\x89\x80\xe3\x71\xf0\x04\xc7\x94\xc2\x7b\xc3\x6c\x7c\x97\x51\x43\xf1\x22\x70\x7c\xc0\x89\x75\x27\x74\xa7\x3e\xa3\x3c\x70\x9a\xce\x14\x17\xf5\x4a\x08\x1e\xd9\xf3\xa5\x89\xdc\x06\xbc\xb5\x8c\x1b\x66\x5f\xaa\x21\x66\x07\x2b\xb8\x65\xd2\x09\x6e\xb6\x1f\xd1\x6d\xce\x9b\x58\x36\x5c\x9a\xb7\xef\x0d\xe4\xb7\x82\x1f\x11\x0e\x70\x67\xc4\x81\xd2\x4c\x61\xa0\xe6\xce\x91\x54\x46\x42\x11\xeb\x30\x92\x79\x8f\xe8\x34\x08\xd3\x4a\xab\x1d\x31\x48\xf8\x26\x84\x41\x28\x9b\xbc\x12\xfe\x89\xe6\xfe\x25\x18\x95\xa8\xbf\x34\xaa\x8c\x51\x9d\xee\xa0\xdc\x3a\x0c\x8e\x6d\x0b\x0c\x96\xeb\x3d\x11\xcb\xe7\xab\xf1\x11\xd1\x37\xe8\x4e\x7b\xe3\x8a\x2a\x70\x48\x78\xe2\x90\x55\x60\x68\x68\xf9\xd5\xdb\x50\x7c\x85\xb6\x29\xde\x30\x3a\xef\xab\x74\xd3\x42\x30\x7a\xa4\xb4\xb9\xc5\xff\x2d\xce\x9b\x86\xf3\xc2\x98\xf8\xdd\x6b\xa2\x4b\x04\xe6\x11\x0e\x18\xcf\x30\xa6\xe9\x4e\x6d\x63\x5a\x5c\x62\xbd\xe1\x7c\xdc\xa8\x30\x6e\x3e\x27\x8c\xdb\xfc\xcd\xd9\x13\x73\x0f\xef\x9a\xd8\xd0\x59\x61\xe2\xeb\x8b\x13\x3b\x07\x75\xd4\xb9\x04\x03\x75\x8b\x61\x1f\x1c\xc3\x51\xcc\x4c\x0a\xae\x41\x7c\x77\x56\xe7\x0e\x93\xd3\x05\x93\xf7\x22\x86\x19\xa6\xf9\xdc\x56\x3b\x6d\x85\x27\x4d\x27\xda\x75\xca\x74\x16\x3b\xf6\x2a\x6e\x6b\xc2\x5c\x18\xfb\xc1\x5c\x1a\xc4\x23\xbc\xcd\xe5\x51\xa8\xaf\x76\x38\x89\xe0\x8e\x8b\x4b\xcb\x92\x6d\xf7\x4b\xd9\xb1\x31\xea\x80\xe5\x51\x9a\xc1\xff\x26\xd8\x0f\x2b\x95\x50\xaa\x29\x7c\xb6\x39\x7c\x96\xcc\xe5\x04\x37\x4a\xf8\xf7\xd8\x47\xbd\x31\x16\x7e\x82\xc7\xdf\xdc\x1d\xb5\x9f\x15\xb8\xe3\xa7\xec\x8e\x7f\x86\x48\x90\x37\xe1\x20\x81\x9b\xe8\xc3\x02\x11\x0a\x09\xee\x1a\x95\xdd\x1d\x85\x06\x4d\xb8\x59\x72\xfc\xd2\x8c\x09\xb7\xc9\x19\xe4\x43\x1e\x6e\x7b\xdb\x8e\x87\xb7\x59\x4e\x35\x6a\xcd\x40\xdf\x30\xb5\x0d\xe7\x76\x52\xf8\x25\xb5\x3e\xe0\xc1\x0b\x1e\x12\xb1\xfb\x07\x77\xad\xfd\x17\xd7\x81\x3f\xb3\xff\x0f\x28\x7f\xdf\xca\xea\xcf\xdc\xed\xff\x02\x00\x00\xff\xff\xba\xcc\xac\x7a\x1e\x0f\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/adjectives.txt"].(os.FileInfo),
		fs["/animals.txt"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
