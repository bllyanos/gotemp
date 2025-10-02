package gotemp

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path"
)

type Gotemp struct {
	basePath string
	pages    map[string]*template.Template
}

func (tc *Gotemp) RenderPage(w io.Writer, layout, page string, data any) error {
	pageTemplate := tc.pages[page]
	if pageTemplate == nil {
		return fmt.Errorf("page template not found: %s", page)
	}
	return pageTemplate.ExecuteTemplate(w, layout, data)
}

func New(basePath string) (*Gotemp, error) {
	gotemp := Gotemp{basePath: basePath}
	err := gotemp.loadPages()
	if err != nil {
		return nil, err
	}
	return &gotemp, nil
}

func (tc *Gotemp) loadPages() error {
	root, err := tc.loadRoot()
	if err != nil {
		return fmt.Errorf("failed to load root template: %w", err)
	}

	partials, err := tc.loadPartials(root)
	if err != nil {
		return fmt.Errorf("failed to load partials: %w", err)
	}

	layouts, err := tc.loadLayouts(partials)
	if err != nil {
		return fmt.Errorf("failed to load layouts: %w", err)
	}

	pages := make(map[string]*template.Template)
	pagesPath := path.Join(tc.basePath, "pages")

	entries, err := os.ReadDir(pagesPath)
	if err != nil {
		return fmt.Errorf("could not read the pages directory: %w", err)
	}

	for _, entry := range entries {
		dirName := entry.Name()
		dirPath := path.Join(pagesPath, dirName)
		files, err := os.ReadDir(dirPath)
		if err != nil {
			return fmt.Errorf("could not read the subpages directory %s: %w", dirPath, err)
		}

		for _, file := range files {
			if !file.IsDir() {
				fileName := file.Name()
				name := path.Join(pagesPath, dirName, fileName)
				layout, err := clone(layouts)
				if err != nil {
					return fmt.Errorf("failed to clone layout template: %w", err)
				}
				pageKey := path.Join(dirName, fileName)
				pages[pageKey], err = layout.ParseFiles(name)
				if err != nil {
					return fmt.Errorf("failed to parse page template %s: %w", name, err)
				}
			}
		}
	}

	tc.pages = pages
	return nil
}

func (tc *Gotemp) loadRoot() (*template.Template, error) {
	template, err := template.ParseGlob(path.Join(tc.basePath, "root.html"))
	if err != nil {
		return nil, err
	}
	return template, nil
}

func (tc *Gotemp) loadPartials(root *template.Template) (*template.Template, error) {
	clonedRoot, err := clone(root)
	if err != nil {
		return nil, fmt.Errorf("failed to clone root template: %w", err)
	}
	template, err := clonedRoot.ParseGlob(path.Join(tc.basePath, "partials", "*.html"))
	if err != nil {
		return nil, err
	}
	return template, nil
}

func (tc *Gotemp) loadLayouts(partials *template.Template) (*template.Template, error) {
	clonedPartials, err := clone(partials)
	if err != nil {
		return nil, fmt.Errorf("failed to clone partials template: %w", err)
	}
	template, err := clonedPartials.ParseGlob(path.Join(tc.basePath, "layouts", "*.html"))
	if err != nil {
		return nil, err
	}
	return template, nil
}

func clone(temp *template.Template) (*template.Template, error) {
	cloned, err := temp.Clone()
	if err != nil {
		return nil, fmt.Errorf("failed to clone template: %w", err)
	}
	return cloned, nil
}
