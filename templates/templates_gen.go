// Code generated by vfsgen; DO NOT EDIT.

package templates

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
			modTime: time.Date(2020, 6, 1, 21, 58, 25, 146507922, time.UTC),
		},
		"/changelog.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "changelog.tmpl",
			modTime:          time.Date(2019, 11, 27, 22, 59, 0, 487458630, time.UTC),
			uncompressedSize: 533,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x90\xcd\x4e\xc3\x30\x10\x84\xef\x79\x8a\x55\xc3\x21\x41\xaa\x7b\x47\xe2\x00\x05\xb5\x48\xfc\x54\x0d\x70\x41\x1c\x4c\xb2\x4d\x0c\x89\x1d\x6c\x97\x1f\xad\xf6\xdd\x91\xd3\x86\xa6\x17\x50\x4f\x89\xbf\x9d\xf1\x8e\x27\x86\xe9\xfc\xec\x76\x76\x79\x7d\x37\x8b\x22\xa2\x4f\xe5\x2b\x38\xca\x2b\xa9\x4b\xac\x4d\x09\x27\xa7\x20\x60\xcc\x1c\x11\xd9\xc0\x40\x2c\xb1\x46\xe9\xd0\x6d\xa9\x5a\x01\xbe\x83\x78\x44\xeb\x94\xd1\x30\x1a\x75\x3c\x8e\xe1\x41\xdb\x8d\xb2\x88\x88\xb0\x76\xd8\x0f\x9e\x88\x7a\x39\xf3\x73\x42\xb4\x5b\x27\x96\xd8\x1a\xa7\xbc\xb1\xdf\xcc\x93\xdc\x34\xad\xb4\x38\x19\xea\x85\x20\x12\x0b\x8b\x1f\xca\xac\xdd\x2f\x4d\x21\x21\x12\x17\xd2\x23\x73\x1a\xd6\xe9\x62\x3f\x74\x86\xb9\x57\x46\xbb\x3e\xb2\xb8\x72\xe7\x16\xe5\x9b\xd2\x65\x17\x2a\x06\x22\x71\xaf\x7c\x8d\x43\xd7\xd4\x34\x8d\xf2\xc1\x74\x1c\xe6\x37\xe8\x9c\x2c\x51\xf4\xce\xed\xb9\x73\x8c\x01\x75\xd1\xff\xed\x1e\x7b\xd0\xbd\x59\x6e\x5a\x5c\x58\x5c\xa9\x2f\xe6\x21\x5f\xbf\xbc\x62\xee\x99\x21\x09\xdd\x65\x95\xb1\x7e\x2e\x5d\xf5\x6f\x7b\x8d\xf2\xa1\xbc\x8d\x36\x4d\xf7\x62\xfe\xf9\xfd\x09\x00\x00\xff\xff\xe5\xe1\xdd\xda\x15\x02\x00\x00"),
		},
		"/circleci.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "circleci.tmpl",
			modTime:          time.Date(2020, 5, 29, 18, 56, 22, 349177735, time.UTC),
			uncompressedSize: 1248,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\x4f\x8f\xd3\x3e\x10\xbd\xf7\x53\x8c\xa2\xdf\xb5\x8d\x7e\x1c\x73\x5b\x68\xa5\x95\xf8\xb3\xa8\xa5\xe2\x18\xb9\xf6\xac\xe3\x4d\xec\x09\x1e\xbb\x61\x15\x85\xcf\x8e\x12\x77\x4b\x60\x81\x56\x42\x7b\xb3\x3d\x6f\x9e\x67\x9e\xfd\xe6\x88\x9e\x0d\xb9\x02\x5e\x2d\x1e\xe8\xc0\xc5\x02\xe0\x10\x4d\xa3\xc6\x05\x80\x22\x59\xa3\x4f\x6b\x80\x25\x18\x2b\x34\x16\xd0\x19\x27\xab\x5c\x53\x23\x9c\x3e\xc5\x00\x44\x0c\x55\x71\xde\x01\x44\x46\xef\x84\xc5\x02\xfe\x5b\xdf\xbd\x79\xbb\xd9\xde\xee\x5f\x97\xfb\xdd\x66\xfb\xe1\xe6\xfd\x66\x86\x6b\x05\x73\x47\x5e\xfd\x84\xfb\x78\xb3\xdb\x7d\xbe\xdb\xae\x17\x7d\xbf\x04\x73\x0f\xf8\x05\x56\xef\x84\xd3\x51\x68\x84\x4c\x53\x36\x0c\x13\x45\x47\xbe\x36\x4e\x97\xca\x78\x94\x81\xfc\x63\x01\xb9\xa6\x9c\xbd\xcc\xfb\xbe\x33\xa1\xa2\x18\x58\x56\x68\x11\x56\x5b\x6c\x89\xcd\x08\x1a\x86\x89\x17\x1b\xc6\xe7\xe4\x8e\x14\xfe\x85\xfe\x5b\xee\xb1\xa5\x94\xef\xd4\x09\xc7\x01\x5b\xfe\x21\x93\xac\x50\xd6\x14\xc3\xf9\x80\x31\xc4\xb6\xf4\x68\x29\x60\x99\x44\xbd\xd4\xd9\x98\xe6\x91\x03\x79\x2c\xa5\x90\x15\xce\xb5\xad\xf1\x91\xe7\xfb\x11\xac\xb0\x45\xa7\xd0\x49\x83\xbc\x3c\xfe\xbf\xec\xfb\xac\xef\xb3\x61\x48\xd5\x70\xb4\x23\xfb\x8a\xa3\xcd\x60\x3a\xcf\xae\x56\xe1\x25\x4a\x69\x85\xac\x85\xc6\xd5\x03\x93\xfb\xb5\xa0\xb3\xac\xd3\xc5\xd1\xcd\xe9\xd3\x8f\xfa\x14\xbd\x83\x50\x21\x54\xc2\xa9\x06\x67\x71\x49\xd6\x0a\xa7\x4e\x7f\x14\xa4\xb9\x46\x67\x16\xc7\xdf\x77\x56\xbc\xb0\xaa\xff\x76\xf1\x45\x0d\x2f\xb6\x9e\xfc\x17\xaa\x67\x2f\x98\x8d\x2e\x6a\x6b\x9d\x5b\x52\xd9\xd5\x0d\xfd\x99\x6e\xc4\x95\x96\x54\x6c\x90\xb3\x59\x85\xa3\xc1\xee\x1b\xea\xa6\x8c\xd9\x2c\x4a\xd6\x1b\x23\x89\xea\x69\x36\x25\xba\xd9\x84\x7a\x7a\x75\x17\xf0\x6b\x28\x60\x3d\x79\xeb\x36\x1e\xf6\xad\x12\x01\xfd\xe2\x7b\x00\x00\x00\xff\xff\x17\x10\x14\x14\xe0\x04\x00\x00"),
		},
		"/go_action.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "go_action.tmpl",
			modTime:          time.Date(2020, 6, 1, 21, 58, 25, 149729953, time.UTC),
			uncompressedSize: 635,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\x4f\x8f\xda\x30\x10\xc5\xef\xf9\x14\x73\xa8\x44\x7b\x30\x11\x3d\xfa\x04\x94\x14\x2a\x54\xa8\xf8\xa3\x1e\x56\x2b\xe4\x98\x51\x9c\x4d\x62\xb3\x1e\x3b\x08\x21\xbe\xfb\xca\x04\x58\x10\x2c\x27\xcb\x7e\xef\xe7\xf7\x66\xb4\xa8\x90\x43\xdf\xe7\xe5\x3a\x8a\x8c\xe6\x11\xc0\xc6\x93\x0a\x27\x40\x6a\x85\x96\x0a\x89\xc3\x0b\x54\x82\x1c\x5a\x78\x3d\x1a\xca\x72\x65\xf1\xdd\x23\xb9\x27\xc6\xe8\xcd\xa4\xc4\xa3\x08\x20\x0d\xdf\x37\xce\xeb\xbc\x70\xb7\x5e\x13\x33\x9a\x83\x4f\xbd\x76\x9e\x95\xc2\x21\xb9\xa3\x84\xba\x6e\x18\x80\xe1\x9f\xc5\x68\xd9\x5f\x2d\xa6\xe3\x64\xc2\xe1\xdb\x7e\x0f\x84\xd2\xa2\xa3\xf6\x70\xd4\xbc\xc2\xe1\x70\xf2\x0e\xa6\xbf\xc6\xc9\x2c\xd8\x97\xf3\x64\x36\xe9\xfd\x4d\x6e\x89\x7b\xfd\x11\xfb\xaf\x37\x9f\xff\x9f\xce\x06\x5f\xb1\x67\xfd\xcc\x92\xc3\x0d\x35\x75\x19\x78\x0a\xab\x10\xd2\xe5\x46\x53\x2c\x15\xca\xc2\x78\xd7\xad\x7f\x9e\x62\xb6\xb9\x53\xe7\xd1\x00\xc8\xa7\x95\x59\xfb\x32\x30\xce\x7a\xbc\x08\xce\x14\xa8\x9f\x8f\x7b\x97\x25\xa4\xc2\x6e\xdd\x79\x18\xb4\x11\x4e\x71\x88\x95\xa9\x30\xb6\x5e\x6b\xb4\x71\x66\x2e\x6a\x81\x3b\x0e\x75\x87\x65\x86\x55\x66\xcd\x42\xaa\x12\xa4\x7e\xe7\x25\xd2\xf7\x56\x66\xda\xe4\xab\xd6\x8f\xcf\x5d\x01\x58\x24\x67\x2c\xb2\x02\x77\x74\x8d\xde\x34\xdb\xe6\x5a\x2a\x99\xc7\xa7\x86\x2c\x33\xa5\xd0\x59\xa8\xf8\x11\x00\x00\xff\xff\x1b\x4d\x3f\x14\x7b\x02\x00\x00"),
		},
		"/go_dockerfile.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "go_dockerfile.tmpl",
			modTime:          time.Date(2020, 6, 1, 21, 57, 43, 649783937, time.UTC),
			uncompressedSize: 276,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8e\xb1\x4e\xc3\x30\x14\x45\xf7\xf7\x15\x57\x1d\xba\xbd\x3a\x0b\x5f\x40\x8d\x54\x21\xec\xc8\x69\x05\x15\x30\x3c\x6c\x43\xad\xd2\x60\xb9\xc9\x14\xe5\xdf\x11\x91\x48\xc4\xc0\x7a\xcf\xb9\xd2\xb9\x73\xf6\x01\xf2\x99\x53\x1b\xc9\x1d\x0c\x24\x84\xfe\x1a\x0b\xb8\x01\x6f\xc1\x27\x28\xc9\x19\xc3\xb0\x31\x72\x89\xe3\x88\xf5\x1a\x2f\x04\x00\x92\xcf\xe8\x73\x90\x2e\xfe\xdd\x98\xdb\x2f\xf6\xe2\x4f\x11\x7d\xfe\x28\x12\xfe\xe7\x12\x02\xbc\xb0\x8f\xa5\x4b\xef\xc9\x4b\x17\xaf\x8b\x7b\x39\x87\x54\xc0\x79\x0a\x50\x6f\xa9\xa5\x43\xa3\xdd\x52\x42\xb7\xb6\x3e\x62\xf3\x43\xd4\x92\xf7\x2b\x2b\x7a\xb4\xee\x7e\xbb\x73\xd3\x42\xda\xec\xdd\xb1\xb6\x3b\xb3\xc7\xf3\x6a\x76\xe6\xdb\xea\x95\xf4\x53\x6d\x1b\x8d\x9b\xaa\xaa\xe8\x3b\x00\x00\xff\xff\xf6\xab\x4d\x97\x14\x01\x00\x00"),
		},
		"/java_mvn_dockerfile.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "java_mvn_dockerfile.tmpl",
			modTime:          time.Date(2020, 2, 3, 0, 33, 1, 204264222, time.UTC),
			uncompressedSize: 291,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8e\x41\x4b\xc3\x40\x10\x85\xef\xfb\x2b\x1e\x3d\xf4\x36\x4d\x04\xbd\x78\xb5\x11\x8a\x98\x0d\x9b\x16\x2d\xea\x61\xdc\x19\xed\x36\x36\x1d\xd2\xe4\x54\xfa\xdf\xc5\x82\x0d\x1e\xbc\xbe\xef\x7b\xf0\xdd\x07\xff\x88\xbd\x69\xbb\x95\xe6\xf6\xea\x9a\xb6\xd2\x10\x7f\x59\x6a\xd5\x85\x55\x09\x16\x19\x0e\xda\x81\x6a\xd0\x1c\xb4\x41\xc6\x66\x38\x1e\x67\x25\xef\xf4\x74\xc2\x74\x8a\x57\x07\x00\x6c\x0d\x06\x13\xee\xf5\xef\x46\xd4\xee\x29\x72\xdc\x28\x06\xfb\xec\x58\xfe\xe7\x2c\x82\xc8\x14\xb5\xeb\xd3\x47\x8a\xdc\xeb\x61\x74\x77\x8d\xa4\x0e\x64\xe7\x80\xec\x3d\xb5\x6e\x55\x17\x61\x2c\x71\x77\xbe\x5a\x63\xf6\x43\xb2\x31\xef\x57\xce\xdc\x93\x0f\x0f\xf3\x45\x38\x2f\xae\x28\x97\x61\x5d\xf9\x45\xb9\xc4\xcb\xe4\xe2\x5c\x6e\x93\x37\x57\x3c\x57\xbe\x2e\x70\x93\xe7\xb9\xfb\x0e\x00\x00\xff\xff\x0e\xc6\xa6\x0d\x23\x01\x00\x00"),
		},
		"/node_action.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "node_action.tmpl",
			modTime:          time.Date(2020, 6, 1, 21, 58, 44, 975561312, time.UTC),
			uncompressedSize: 690,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\xc1\xce\xda\x30\x10\x84\xef\x79\x8a\x3d\x54\xfa\xdb\x83\x41\x7f\x8f\x3e\x11\x44\x0a\x55\x4b\x82\x92\xa0\x1e\xaa\x2a\x72\xcc\xaa\x0e\x09\x76\x9a\xb5\x83\x10\xa2\xcf\x5e\x85\x04\x0a\x6a\x5a\xf5\x64\x79\xe7\x9b\xdd\xf1\x5a\x8b\x03\x72\x98\xbb\xa2\xda\x79\x9e\xd1\xdc\x03\xa8\x1d\xa9\xee\x04\xc8\x1b\xa1\xa5\x42\xe2\xf0\x15\x0e\x82\x2c\x36\xf0\xed\x0a\x54\x55\xd6\xe0\x0f\x87\x64\xff\x01\x7a\x7b\x93\x13\xf7\x3c\x80\xbc\x6b\xdf\x93\x8f\xf3\xba\x7b\xe3\x34\x31\xa3\x39\xb8\xdc\x69\xeb\x58\x25\x2c\x92\xbd\x4a\xa8\xdb\xde\x03\xb0\xfc\x98\xae\xb6\xf3\x2c\x8d\x3e\x05\x21\x87\x37\xe7\x33\x10\xca\x06\x2d\x4d\x96\xab\xbe\x0a\x97\xcb\xc0\x86\x9b\xf5\x18\x78\x2f\x3f\x90\xd1\x22\xc8\xfc\x6d\xba\xfa\xbf\xc6\x69\xec\x87\x49\xb4\xce\xb6\xf1\xe7\x67\xf4\x41\x18\xa1\x93\x20\x0e\xfd\x75\xf0\x17\xcb\xa0\xfe\xe9\xdb\xf8\x49\xf2\x25\x8a\x17\xe3\xbe\x9b\x7a\xf3\x91\xc5\x9a\xfa\x6d\x31\x70\xd4\xfd\x84\x90\xb6\x30\x9a\xa6\x52\xa1\x2c\x8d\xb3\xb3\xf6\xfd\xb8\x2e\xa4\xc2\x59\xfb\x3a\xcc\x3f\x16\x56\xdd\xd6\x0e\x50\x0b\xab\x38\xfc\x9c\x4e\x74\x7d\xb8\x17\x4b\x3c\x71\x68\x5f\x99\x36\x3b\x64\x3b\xac\x59\x97\x50\x09\x52\x1f\x8a\x0a\xe9\xed\x4b\x2d\x64\x29\xbe\x23\xab\x8c\x2c\x27\x7b\x32\xfa\xe5\xdd\xef\xf7\x01\x34\x48\xd6\x34\xc8\x4a\x3c\xd1\x73\x9f\xa7\x7c\xc7\x42\x4b\x25\x8b\xe9\x90\xf3\x4a\x75\x31\x7f\x05\x00\x00\xff\xff\x3f\x96\x0b\x31\xb2\x02\x00\x00"),
		},
		"/node_npm_dockerfile.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "node_npm_dockerfile.tmpl",
			modTime:          time.Date(2020, 2, 3, 0, 33, 1, 204554105, time.UTC),
			uncompressedSize: 265,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8e\x3f\x4f\xc3\x30\x14\xc4\x77\x7f\x8a\x53\x86\x4e\xbc\x92\x85\x85\x95\x06\xa9\x42\xc4\x91\xd3\x0a\x2a\x60\x78\xb2\x0d\xb5\x48\x9c\x27\xff\x99\xaa\x7e\x77\x54\x18\x10\x03\xd3\xe9\xee\x86\xdf\xef\xde\xe8\x47\xc4\xc5\xf9\x5b\x9e\x24\x44\xaf\xcc\xbe\x07\x3b\x57\xb3\x4f\xa0\x11\xb4\x01\x1d\x71\xcd\x22\x38\x9d\xd6\x3d\xcf\xfe\x7c\xc6\x6a\x85\x57\x05\x00\x2c\x9f\xa8\xe2\xb8\xf8\xbf\x1b\x51\x5c\xc8\xb2\x3d\x7a\x54\xf9\x48\xec\xfe\xff\xd9\x39\x58\x26\xeb\x53\x09\xef\xc1\x72\xf1\x59\xed\xc7\xce\xfc\xe2\xd4\x9d\x1e\x0e\x58\x7f\x4b\xa8\x27\x6d\x1e\x36\x5b\xf3\x53\x2e\xae\x51\x66\x84\x98\x0b\x4f\x13\x88\x24\x2d\xae\xda\x12\x96\xa8\xba\x7e\x67\x0e\x83\xde\xf6\x3b\xbc\x34\x51\xe6\xe6\x0a\x4d\xaa\xf1\x12\xb9\x70\x2a\xcd\x9b\xea\x9e\x07\x3d\x76\xb8\x69\xdb\x56\x7d\x05\x00\x00\xff\xff\xdf\x71\xb5\xfa\x09\x01\x00\x00"),
		},
		"/node_yarn_dockerfile.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "node_yarn_dockerfile.tmpl",
			modTime:          time.Date(2020, 2, 3, 0, 33, 1, 204808815, time.UTC),
			uncompressedSize: 259,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\xce\x31\x4f\xc3\x30\x14\x04\xe0\xdd\xbf\xe2\x94\xa1\x13\xaf\x64\x61\x61\xa5\x41\xaa\x10\x71\xe4\xb4\x82\x0a\x18\x9e\xec\x07\xb5\x48\x9d\x27\xc7\x19\x50\x95\xff\x8e\x80\x01\x31\x30\xde\xdd\x70\xdf\xad\xb3\xf7\x48\x63\x90\x6b\x1e\x34\x26\x31\x6e\xdf\x82\x43\x98\x27\xc9\xa0\x1e\xb4\x01\x1d\x71\xc9\xaa\x38\x9f\xd7\x2d\x9f\x64\x59\xb0\x5a\xe1\xd9\x00\x00\xeb\x3b\x66\x0d\x5c\xe4\x6f\x47\x94\x46\xf2\xec\x8f\x82\x59\xdf\x32\x87\xff\x77\x0e\x01\x9e\xc9\x4b\x2e\xf1\x35\x7a\x2e\x32\x99\x7d\xdf\xb8\xdf\x3b\x73\x63\xbb\x03\xd6\xdf\x08\xf3\x60\xdd\xdd\x66\xeb\x7e\xc2\x97\x35\xe9\x09\x31\x4d\x85\x87\x01\x44\x9a\xc7\x30\xfb\x12\xc7\x64\x9a\x76\xe7\x0e\x9d\xdd\xb6\x3b\x3c\x55\x1f\x9c\x53\x75\x81\x6a\x2a\x9c\x4b\xf5\x62\x9a\xc7\xce\xf6\x0d\xae\xea\xba\x36\x9f\x01\x00\x00\xff\xff\x5e\x30\x37\x40\x03\x01\x00\x00"),
		},
		"/python_dockerfile.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "python_dockerfile.tmpl",
			modTime:          time.Date(2020, 2, 3, 0, 33, 1, 205049011, time.UTC),
			uncompressedSize: 272,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8f\x41\x4b\x03\x31\x10\x85\xef\xf9\x15\x8f\x3d\xf4\xe4\x6c\x17\x44\x10\xaf\x76\x85\x22\x26\x4b\xb6\x45\x8b\x7a\x18\x93\xe8\x86\xda\x74\x48\xb3\x87\x52\xfa\xdf\x85\x16\x2c\x1e\x3c\xce\x7c\xef\xc1\xf7\x1e\xac\x79\x82\xec\xcb\xb0\x4d\x77\xd7\xf5\x2d\xf1\xb7\xc4\x14\x94\x5d\x6a\xb0\xf7\xe3\x2e\x64\x50\x0f\x9a\x81\x06\x4c\x59\x04\x87\x43\xad\x79\x13\x8e\x47\x4c\x26\x78\x53\x00\xc0\xb2\xc6\x28\x9e\x4b\xf8\xfb\x23\x4a\x5b\x72\xec\x86\x80\x51\xbe\x32\xfb\xff\x39\x7b\x0f\xc7\xe4\x42\x2e\xf1\x33\x3a\x2e\x61\x77\xc9\x6e\xd6\x3e\x66\x90\x9c\x04\xa6\x1f\x31\xa9\x65\xdf\xda\x8b\x89\xba\x37\xdd\x0a\xf5\x09\xab\x67\x63\x1f\x67\x73\x7b\x3e\x5a\xbd\xb0\xab\xce\xcc\xf5\x02\xaf\xd5\x79\x65\x75\x85\xea\xb7\x59\xcb\xbe\x7a\x57\xed\x4b\x67\xfa\x16\x37\x4d\xd3\xa8\x9f\x00\x00\x00\xff\xff\xc6\xc3\x81\x36\x10\x01\x00\x00"),
		},
		"/release.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "release.tmpl",
			modTime:          time.Date(2019, 11, 29, 6, 20, 19, 830044147, time.UTC),
			uncompressedSize: 361,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x90\x41\x4b\x43\x31\x0c\xc7\xef\xfb\x14\x81\xe7\x61\x13\x96\xdd\x05\x2f\x7a\xd1\x83\x20\x7b\xde\xc4\x43\xad\x59\x1b\xed\xda\xd1\x54\x54\x42\xbe\xbb\xbc\xea\x70\xbb\x08\x3b\xb5\xfd\x35\xbf\xfc\x43\x54\x3f\xb8\x45\x38\xf3\xd1\xe5\x40\xa9\x04\xb8\xb8\x04\x84\xa5\xd9\x4c\xb5\x4e\x0c\x70\x4d\x89\x9c\x90\x1c\xd3\x91\x7c\xe3\x92\xa5\x33\xde\x00\xde\xca\x55\x25\xf7\xc6\x39\x98\xcd\x86\x61\x00\x55\x7c\xe0\x96\xe8\xd0\xba\x2e\xdb\x2d\xb7\x49\x3a\x9f\xfe\xef\x48\xc4\x05\xc2\xbd\xf9\xfb\xee\xc6\x12\x28\xbf\xec\x6f\x49\xa8\xc7\x9f\xd8\x77\xf4\x65\x47\xf7\x95\x36\xfc\x69\x76\xc8\xdf\x9f\x5f\xc9\x37\x33\x98\x3f\xaa\xe2\x18\x4b\x6d\x37\x4e\xa2\xd9\xd3\x5c\xf5\x6f\x1b\xb8\xa6\x5d\x11\x6e\xa5\x7e\x99\xad\x7c\x0f\x59\xa9\xe2\x4f\xed\x62\x71\x34\xe6\xbf\xe7\x77\x00\x00\x00\xff\xff\x32\x5d\x39\x69\x69\x01\x00\x00"),
		},
		"/scala_sbt_dockerfile.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "scala_sbt_dockerfile.tmpl",
			modTime:          time.Date(2020, 2, 3, 0, 33, 1, 205277781, time.UTC),
			uncompressedSize: 291,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8e\x41\x4b\xc3\x40\x10\x85\xef\xfb\x2b\x1e\x3d\xf4\x36\x4d\x04\xbd\x78\xb5\x11\x8a\x98\x0d\x9b\x16\x2d\xea\x61\xdc\x19\xed\x36\x36\x1d\xd2\xe4\x54\xfa\xdf\xc5\x82\x0d\x1e\xbc\xbe\xef\x7b\xf0\xdd\x07\xff\x88\xbd\x69\xbb\x95\xe6\xf6\xea\x9a\xb6\xd2\x10\x7f\x59\x6a\xd5\x85\x55\x09\x16\x19\x0e\xda\x81\x6a\xd0\x1c\xb4\x41\xc6\x66\x38\x1e\x67\x25\xef\xf4\x74\xc2\x74\x8a\x57\x07\x00\x6c\x0d\x06\x13\xee\xf5\xef\x46\xd4\xee\x29\x72\xdc\x28\x06\xfb\xec\x58\xfe\xe7\x2c\x82\xc8\x14\xb5\xeb\xd3\x47\x8a\xdc\xeb\x61\x74\x77\x8d\xa4\x0e\x64\xe7\x80\xec\x3d\xb5\x6e\x55\x17\x61\x2c\x71\x77\xbe\x5a\x63\xf6\x43\xb2\x31\xef\x57\xce\xdc\x93\x0f\x0f\xf3\x45\x38\x2f\xae\x28\x97\x61\x5d\xf9\x45\xb9\xc4\xcb\xe4\xe2\x5c\x6e\x93\x37\x57\x3c\x57\xbe\x2e\x70\x93\xe7\xb9\xfb\x0e\x00\x00\xff\xff\x0e\xc6\xa6\x0d\x23\x01\x00\x00"),
		},
		"/version_go.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "version_go.tmpl",
			modTime:          time.Date(2019, 12, 12, 10, 11, 44, 556615677, time.UTC),
			uncompressedSize: 554,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\x51\x4b\xc3\x30\x14\x85\x9f\x73\x7f\xc5\x35\x20\x34\x50\xdb\x77\x61\x0f\xb2\x89\xfa\xb2\xca\x36\x7c\xaf\xd9\xed\x16\x5c\xd3\x90\xa4\x8a\x94\xfe\x77\x69\x9a\x6d\x61\x60\xe9\x4b\xce\xf9\xce\xb9\xb9\x29\x4b\x5c\x76\x7b\xc2\x03\x69\xb2\xb5\xa7\x3d\x7e\xfe\xe2\x8f\xd2\xf2\x58\xe0\xaa\xc2\x75\xb5\xc3\xe7\xd5\xdb\xae\x00\x53\xcb\xaf\xfa\x40\xf8\x4d\xd6\xa9\x4e\x03\xa8\xd6\x74\xd6\x63\x06\x8c\x37\xad\xe7\xc0\xb8\xed\xb5\x57\x2d\x71\x10\x00\xb2\xd3\x2e\x98\xeb\xba\x25\x8c\xdf\x02\xf9\x30\x14\x93\x32\x8e\x1c\xd8\x8a\x9c\xb4\xca\x78\xd5\xe9\x68\x25\x4a\x20\x36\x74\xa2\xda\x51\xe8\x98\x89\x44\x09\xc4\xc7\x7c\x9d\xa4\x3e\x2a\xc1\x7d\xb7\x64\xe7\xc0\xd9\xbd\x2a\x13\x20\x00\xca\x12\xb7\xde\x2a\x7d\x40\x4b\xbe\xb7\xda\xa1\x3f\x12\xca\xae\x35\x27\xf2\x97\x75\xd1\x05\x26\x47\xa5\xe5\xa9\xdf\x4f\xb8\xb9\x34\x41\xd3\x6b\x19\x5b\x32\x11\x51\x1c\x80\x39\x7c\x5c\x60\xd3\xfa\x62\x6b\xac\xd2\xbe\xc9\xf8\xbd\xc3\xf0\xf3\x1c\xe3\x6b\x15\x2f\x55\xb5\x4d\x4f\x4f\x9b\xe5\xeb\xf5\x1c\xb7\xc9\x84\x00\xa6\x1a\x4c\x16\xba\x5b\x20\xe7\xd3\x14\x36\x5f\xfc\x76\xd0\xc3\x79\x50\xac\xc8\x93\x70\x8e\x4e\x00\x1b\xe1\x9f\xe8\x4d\xce\x09\x18\x01\xfe\x02\x00\x00\xff\xff\xdc\x0d\x6b\xef\x2a\x02\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/changelog.tmpl"].(os.FileInfo),
		fs["/circleci.tmpl"].(os.FileInfo),
		fs["/go_action.tmpl"].(os.FileInfo),
		fs["/go_dockerfile.tmpl"].(os.FileInfo),
		fs["/java_mvn_dockerfile.tmpl"].(os.FileInfo),
		fs["/node_action.tmpl"].(os.FileInfo),
		fs["/node_npm_dockerfile.tmpl"].(os.FileInfo),
		fs["/node_yarn_dockerfile.tmpl"].(os.FileInfo),
		fs["/python_dockerfile.tmpl"].(os.FileInfo),
		fs["/release.tmpl"].(os.FileInfo),
		fs["/scala_sbt_dockerfile.tmpl"].(os.FileInfo),
		fs["/version_go.tmpl"].(os.FileInfo),
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
