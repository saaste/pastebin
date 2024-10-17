package documents

import (
	"cmp"
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"gopkg.in/yaml.v3"
)

const yamlFile = "documents.yaml"

type Storage struct{}

var ErrNotFound error = errors.New("document not found")

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) List() (*DocumentList, error) {
	yamlFile, err := os.ReadFile(yamlFile)
	if err != nil {
		return nil, err
	}

	docList := &DocumentList{}

	err = yaml.Unmarshal(yamlFile, docList)
	if err != nil {
		return nil, err
	}

	return docList, nil
}

func (s *Storage) Create(doc *Document) error {
	if doc.Name == "" {
		return fmt.Errorf("name is required")
	}

	if doc.Syntax == "" {
		return fmt.Errorf("syntax is required")
	}

	if doc.IsPublic && doc.PublicPath == "" {
		return fmt.Errorf("public path is required if document is public")
	}

	docList, err := s.List()
	if err != nil {
		return fmt.Errorf("failed to fetch document list: %w", err)
	}

	if !s.isUniquePublicPath(doc, docList) {
		return fmt.Errorf(fmt.Sprintf("Public path %s is already in use", doc.PublicPath))
	}

	now := time.Now().UTC()

	doc.Id = uuid.NewString()
	doc.Created = now
	doc.Updated = now
	doc.PublicPath = slug.Make(doc.PublicPath)
	docList.Documents = append(docList.Documents, *doc)

	err = s.saveYamlFile(docList)
	if err != nil {
		return fmt.Errorf("failed to save the document list: %w", err)
	}

	return nil
}

func (s *Storage) Update(id string, doc *Document, content string) error {
	if doc.Name == "" {
		return fmt.Errorf("name is required")
	}

	if doc.Syntax == "" {
		return fmt.Errorf("syntax is required")
	}

	if doc.IsPublic && doc.PublicPath == "" {
		return fmt.Errorf("public path is required if document is public")
	}

	docList, err := s.List()
	if err != nil {
		return fmt.Errorf("failed to fetch the documents: %w", err)
	}

	if !s.isUniquePublicPath(doc, docList) {
		return fmt.Errorf(fmt.Sprintf("Public path %s is already in use", doc.PublicPath))
	}

	updated := false
	for idx, oldDoc := range docList.Documents {
		if oldDoc.Id != id {
			continue
		}

		doc.Id = id
		doc.Created = oldDoc.Created
		doc.Updated = time.Now().UTC()
		docList.Documents[idx] = *doc
		updated = true
		break
	}

	if !updated {
		return errors.New("failed to update the document: not found")
	}

	return s.saveYamlFile(docList)
}

func (s *Storage) Delete(id string) error {
	docList, err := s.List()
	if err != nil {
		return fmt.Errorf("failed to fetch the documents: %w", err)
	}

	newDocList := make([]Document, 0)
	for _, doc := range docList.Documents {
		if doc.Id != id {
			newDocList = append(newDocList, doc)
		}
	}

	docList.Documents = newDocList
	s.saveYamlFile(docList)
	return nil
}

func (s *Storage) GetById(id string) (*Document, error) {
	docList, err := s.List()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch the documents: %w", err)
	}

	for _, doc := range docList.Documents {
		if doc.Id == id {
			return &doc, nil
		}
	}

	return nil, ErrNotFound
}

func (s *Storage) GetByPublicPath(path string) (*Document, error) {
	docList, err := s.List()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch the documents: %w", err)
	}

	for _, doc := range docList.Documents {
		if doc.IsPublic && doc.PublicPath == path {
			return &doc, nil
		}
	}

	return nil, ErrNotFound
}

func (s *Storage) saveYamlFile(docs *DocumentList) error {
	slices.SortFunc(docs.Documents, s.compareDocuments)
	data, err := yaml.Marshal(docs)
	if err != nil {
		return fmt.Errorf("failed to unmarshal the file file: %w", err)
	}
	err = os.WriteFile(yamlFile, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to save the document list: %w", err)
	}
	return nil
}

func (s *Storage) compareDocuments(a, b Document) int {
	return cmp.Compare(strings.ToLower(a.Name), strings.ToLower(b.Name))
}

func (s *Storage) isUniquePublicPath(doc *Document, docList *DocumentList) bool {
	if !doc.IsPublic || doc.PublicPath == "" {
		return true
	}

	for _, d := range docList.Documents {
		if d.PublicPath == doc.PublicPath && d.Id != doc.Id {
			return false
		}
	}
	return true
}
