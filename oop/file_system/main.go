package main

import (
	"fmt"
	"sort"
	"strings"
)

type FileSystem struct {
	Root *Node
}

type Node struct {
	IsFile   bool
	Name     string
	Children map[string]*Node
	Content  string
}

func (n *Node) Ls() []string {
	keys := make([]string, 0, len(n.Children))
	for k := range n.Children {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys
}

func NewFileSystem() *FileSystem {
	return &FileSystem{
		Root: &Node{
			IsFile:   false,
			Name:     "",
			Children: map[string]*Node{},
			Content:  "",
		},
	}
}

func (fs *FileSystem) Mkdir(path string) error {
	pathComponents := strings.Split(path, "/")[1:]

	curNode := fs.Root
	for _, pathComponent := range pathComponents {
		if childNode, ok := curNode.Children[pathComponent]; !ok {
			curNode.Children[pathComponent] = &Node{
				IsFile:   false,
				Name:     pathComponent,
				Children: map[string]*Node{},
				Content:  "",
			}
			curNode = curNode.Children[pathComponent]
		} else if childNode.IsFile {
			return fmt.Errorf("file with the same name exists")
		} else {
			curNode = childNode
		}
	}

	return nil
}

func (fs *FileSystem) Ls(path string) ([]string, error) {
	node, _, err := fs.Seek(path)
	if err != nil {
		return nil, err
	}

	if node.IsFile {
		return []string{node.Name}, err
	}

	return node.Ls(), nil
}

func (fs *FileSystem) Seek(path string) (*Node, int, error) {
	if path == "/" {
		return fs.Root, 0, nil
	}

	pathComponents := strings.Split(path, "/")[1:]
	curNode := fs.Root
	traversalIndex := 0
	for _, pathComponent := range pathComponents {
		if childNode, ok := curNode.Children[pathComponent]; !ok {
			return curNode, traversalIndex, fmt.Errorf("path doesn't exist %s", pathComponent)
		} else {
			traversalIndex = traversalIndex + 1
			curNode = childNode
		}
	}

	return curNode, traversalIndex, nil
}

func (fs *FileSystem) AddContentToFile(path string, content string) error {
	pathComponents := strings.Split(path, "/")[1:]
	if pathComponents[len(pathComponents)-1] == "" {
		return fmt.Errorf("not a file")
	}

	node, traversalIndex, err := fs.Seek(path)
	if err != nil && !(traversalIndex == len(pathComponents)-2 && strings.HasPrefix(err.Error(), "path doesn't exist")) {
		return err
	}

	if err != nil && traversalIndex == len(pathComponents)-2 && strings.HasPrefix(err.Error(), "path doesn't exist") {
		node.Children[pathComponents[len(pathComponents)-1]] = &Node{
			IsFile:   true,
			Name:     pathComponents[len(pathComponents)-1],
			Content:  content,
			Children: map[string]*Node{},
		}

		return nil
	}

	if !node.IsFile {
		return fmt.Errorf("can not add content to dir")
	}

	node.Content += content

	return nil
}

func (fs *FileSystem) readContentFromFile(path string) (string, error) {
	node, _, err := fs.Seek(path)
	if err != nil {
		return "", err
	}

	if !node.IsFile {
		return "", fmt.Errorf("can not read content from dir")
	}

	return node.Content, nil
}

func main() {

}
