package internal

import (
	"fmt"
	"os"
	"path/filepath"
)

type Object struct {
	Name    string
	Hash    string
	IsTree  bool
	Content string
}

func NewObject(name, hash, content string, isTree bool) *Object {
	return &Object{
		Name:    name,
		Hash:    hash,
		IsTree:  isTree,
		Content: content,
	}
}

type ObjectMap struct {
	Objects map[string]Object // key: path, value: object
}

func NewObjectMap() *ObjectMap {
	return &ObjectMap{
		Objects: make(map[string]Object),
	}
}

func (om *ObjectMap) AddObject(entry Entry) error {
	if entry.IsDir {
		tree, err := om.newTreeObject(entry)
		if err != nil {
			return err
		}
		om.Objects[entry.Path] = *tree
	} else {
		blob, err := om.newBlobOject(entry)
		if err != nil {
			return err
		}
		om.Objects[entry.Path] = *blob
	}
	return nil
}

func (om *ObjectMap) Get(path string) (Object, bool) {
	if obj, ok := om.Objects[path]; ok {
		return obj, true
	}
	return Object{}, false
}

func (om *ObjectMap) newTreeObject(entry Entry) (*Object, error) {
	entries, err := os.ReadDir(entry.Path)
	if err != nil {
		return &Object{}, err
	}
	content := ""
	for _, e := range entries {
		obj, ok := om.Get(entry.Path + "/" + e.Name())
		if !ok {
			return &Object{}, fmt.Errorf("object not found: %s", entry.Path+"/"+e.Name())
		}
		var objType string
		if obj.IsTree {
			objType = "tree"
		} else {
			objType = "blob"
		}
		content += fmt.Sprintf("%s %s %s\n", objType, obj.Hash, e.Name())
	}
	hash := CalcHash(content)
	return &Object{
		Name:    filepath.Base(entry.Path),
		Hash:    hash,
		IsTree:  true,
		Content: content,
	}, nil
}

func (om *ObjectMap) newBlobOject(entry Entry) (*Object, error) {
	contentInBytes, err := os.ReadFile(entry.Path)
	if err != nil {
		return &Object{}, err
	}
	fileContent := string(contentInBytes)
	blobContent := "blob\n" + fileContent
	hash := CalcHash(blobContent)
	return &Object{
		Name:    filepath.Base(entry.Path),
		Hash:    hash,
		IsTree:  false,
		Content: blobContent,
	}, nil
}
