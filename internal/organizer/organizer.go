package organizer

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type FileCategory struct {
	Name       string
	Extensions []string
}

var DefaultCategories = []FileCategory{
	{
		Name:       "Documents",
		Extensions: []string{".pdf", ".doc", ".docx", ".txt", ".rtf", ".odt"},
	},
	{
		Name:       "Images",
		Extensions: []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".svg"},
	},
	{
		Name:       "Videos",
		Extensions: []string{".mp4", ".avi", ".mkv", ".mov", ".wmv"},
	},
	{
		Name:       "Audio",
		Extensions: []string{".mp3", ".wav", ".flac", ".m4a", ".aac"},
	},
	{
		Name:       "Archives",
		Extensions: []string{".zip", ".rar", ".7z", ".tar", ".gz"},
	},
}

type Organizer struct {
	SourceDir  string
	Categories []FileCategory
}

func New(sourceDir string) *Organizer {
	return &Organizer{
		SourceDir:  sourceDir,
		Categories: DefaultCategories,
	}
}

func (o *Organizer) getCategoryForFile(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	for _, category := range o.Categories {
		for _, extension := range category.Extensions {
			if ext == extension {
				return category.Name
			}
		}
	}
	return "Others"
}

func (o *Organizer) Organize() error {
	return filepath.Walk(o.SourceDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Dir(path) != o.SourceDir {
			return nil
		}

		category := o.getCategoryForFile(info.Name())
		categoryPath := filepath.Join(o.SourceDir, category)

		if err := os.MkdirAll(categoryPath, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %v", categoryPath, err)
		}

		newPath := filepath.Join(categoryPath, info.Name()) 
		if err := os.Rename(path, newPath); err != nil {
			return fmt.Errorf("failed to move file %s: %v", info.Name(), err)
		}
		return nil
	})
}
