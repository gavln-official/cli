package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Gvinaxu/cli/util"
	"github.com/valyala/fasthttp"
)

const (
	treeUrl = "http://127.0.0.1:18080/api/1/file/list"
)

var (
	dirCount  = 0
	fileCount = 0
)

type FileReq struct {
}

func (f *FileReq) Tree(path string) ([]*File, error) {
	args := &fasthttp.Args{}
	args.Add("path", path)
	head := map[string]interface{}{
		"access_token": token.AccessToken,
	}
	fmt.Println(path)
	code, body, err := util.DoTimeout(args, "POST", treeUrl, head)
	if err != nil {
		return nil, err
	}
	if code != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("response code is %d", code))
	}
	type R struct {
		Code int     `json:"code"`
		Data []*File `json:"data"`
	}
	r := &R{}
	err = json.Unmarshal(body, r)
	if err != nil {
		return nil, err
	}
	if path == "" {
		path = "/"
	}

	fmt.Println(path)
	f.printDirectory(r.Data, 0)
	fmt.Println()
	fmt.Printf("%d directories, %d files \n", dirCount, fileCount)
	return r.Data, nil
}

func (f *FileReq) printDirectory(files []*File, depth int) {
	for _, file := range files {
		if file.Dir {
			dirCount++
			f.printListing(file.Name, depth, true)
			f.printDirectory(file.Child, depth+1)
		} else {
			fileCount++
			f.printListing(file.Name, depth+1, false)
		}
	}

}

func (f *FileReq) printListing(entry string, depth int, dir bool) {
	indent := strings.Repeat("|   ", depth)
	fmt.Printf("%s|-- %s\n", indent, entry)
	// output color
}

type File struct {
	FecEnable bool         `json:"fec_enable"`
	Blocks    []*FileBlock `json:"blocks"`
	Path      string       `json:"path"`
	Name      string       `json:"name"`
	Size      int64        `json:"size"`
	Time      int64        `json:"time"`
	Dir       bool         `json:"dir"`
	Expires   int64        `json:"expires"`
	Child     []*File      `json:"child"`
}

type FileBlock struct {
	Key string `json:"key"`
	Cid string `json:"cid"`
}
