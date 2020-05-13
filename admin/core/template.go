package core

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"html/template"

	"github.com/gin-contrib/multitemplate"
)

func LoadTemplates(templatesDir string,funcMap template.FuncMap) multitemplate.Renderer {
	templatesDir = strings.Trim(templatesDir, "/")
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/*.html")
	if err != nil {
		panic(err.Error())
	}

	for _, layout := range layouts {
		filename := filepath.Base(layout)
		dir := strings.ReplaceAll(filename, filepath.Ext(filename), "")

		includes, err := filepath.Glob(fmt.Sprintf("%s/%s/**/*.html", templatesDir, dir))
		if err != nil {
			panic(err.Error())
		}

		baseDir := fmt.Sprintf("%s/%s/", templatesDir, dir)
		for _, include := range includes {
			files := []string{layout, include}
			r.AddFromFilesFuncs(strings.Replace(filepath.ToSlash(include), baseDir, "", 1),funcMap, files...)

			log.Println(strings.Replace(filepath.ToSlash(include), baseDir, "", 1))
		}
	}

	return r
}
