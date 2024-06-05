package internal

import (
	"fmt"
	"sort"
	"strings"
)

type IndexContent struct {
	Hash string
	Path string
}

func Contains(indexContents []IndexContent, path string) bool {
	for _, entry := range indexContents {
		if entry.Path == path {
			return true
		}
	}
	return false
}

func ParseIndexFile(indexFile string) []IndexContent {
	var indexContents []IndexContent
	lines := strings.Split(indexFile, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		indexContents = append(indexContents, IndexContent{Hash: parts[0], Path: parts[1]})
	}
	return indexContents
}

func ParseIndexContents(indexEntries []IndexContent) *treeNode {
	root := newTreeNode(make(map[string]node))
	for _, entry := range indexEntries {
		parts := strings.Split(entry.Path, "/")
		current := root
		for _, dirName := range parts[:len(parts)-1] {
			if _, ok := current.Children[dirName]; !ok {
				current.Children[dirName] = newTreeNode(make(map[string]node))
			}
			current = current.Children[dirName].(*treeNode)
		}
		current.Children[parts[len(parts)-1]] = newBlobNode(entry.Hash, parts[len(parts)-1])
	}
	return root
}

func CreateTreeObject(tree treeNode) string {
	var treeContent []string
	for name, item := range tree.Children {
		switch node := item.(type) {
		case *blobNode:
			treeContent = append(treeContent, fmt.Sprintf("%s %s %s", node.Type, node.Hash, node.Path))
		case *treeNode:
			subtreeHash := CreateTreeObject(*node)
			treeContent = append(treeContent, fmt.Sprintf("%s %s %s", node.Type, subtreeHash, name))
		}
	}
	sort.Strings(treeContent)
	treeData := strings.Join(treeContent, "\n")
	treeHash := CalcHash(treeData)
	CreateFile(*NewFile(".bbgit/objects/"+treeHash, treeData))
	return treeHash
}

type node interface {
	isNode()
}

type blobNode struct {
	Type string
	Hash string
	Path string
}

func (f *blobNode) isNode() {}

func newBlobNode(hash, path string) *blobNode {
	return &blobNode{"blob", hash, path}
}

type treeNode struct {
	Type     string
	Children map[string]node
}

func (t *treeNode) isNode() {}

func newTreeNode(children map[string]node) *treeNode {
	return &treeNode{"tree", children}
}
