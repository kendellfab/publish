package domain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type UploadNode struct {
	Name     string       `json:"name"`
	Size     int64        `json:"size"`
	ModTime  time.Time    `json:"modTime"`
	IsDir    bool         `json:"isDir"`
	Children []UploadNode `json:"children"`
}

func (u *UploadNode) String() string {
	jNode, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(jNode)
}

func GetUploadDir(rootPath string) *UploadNode {
	info, err := os.Lstat(rootPath)
	if err != nil {
		fmt.Println(err)
	}
	root := fileInfoToUploadNode(info)
	root.Children = generator(rootPath, info)
	return &root
}

func generator(path string, info os.FileInfo) []UploadNode {

	if !info.IsDir() {
		fmt.Println("Not a directory :(")
		return nil
	}
	children := make([]UploadNode, 0)
	list, listErr := ioutil.ReadDir(path)
	if listErr != nil {
		fmt.Println(listErr)
		return nil
	}
	for _, item := range list {
		if item.IsDir() {
			node := fileInfoToUploadNode(item)
			node.Children = generator(filepath.Join(path, item.Name()), info)
			children = append(children, node)
		} else {
			children = append(children, fileInfoToUploadNode(item))
		}
	}

	return children
}

func fileInfoToUploadNode(info os.FileInfo) UploadNode {
	node := UploadNode{
		Name:    info.Name(),
		Size:    info.Size(),
		ModTime: info.ModTime(),
		IsDir:   info.IsDir(),
	}
	return node
}

func CreateNewDirectory(uploadDir, loc, dir string) bool {
	base := path.Base(uploadDir)
	temp := strings.Replace(loc, base, "", 1)
	newDir := path.Join(temp, dir)
	newDir = path.Join(uploadDir, newDir)
	os.Mkdir(newDir, 0777)
	fmt.Println(newDir)
	return true
}
