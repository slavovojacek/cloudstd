/*
Copyright © 2022 Slavo Vojacek <public@slavovojacek.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package proto

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/slavovojacek/cloudstd/logger"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type templateData struct {
	Package            string
	PackagePath        string
	ResourceName       string
	ResourcePluralName string
	ParentName         string
	ParentPluralName   string
}

var caser = cases.Title(language.Und)

var funcs = template.FuncMap{
	"ToTitle": caser.String,
}

func newResourceTemplate(styleguide string) *template.Template {
	return newTemplate(styleguide+"/resource.tmpl", "resource.proto")
}

func newRequestResponseTemplate(styleguide string) *template.Template {
	return newTemplate(styleguide+"/request_response.tmpl", "request_response.proto")
}

func newServiceTemplate(styleguide string) *template.Template {
	return newTemplate(styleguide+"/service.tmpl", "service.proto")
}

func newTemplate(filename, name string) *template.Template {
	text, err := content.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return template.Must(
		template.
			New(name).
			Funcs(funcs).
			Parse(string(text)),
	)
}

func newTemplateData(pkg string, resource, parentResource *resourceName) *templateData {
	pkgSegments := strings.Split(pkg, ".")
	pkgPath := strings.Join(pkgSegments, "/")

	return &templateData{
		Package:            strings.ToLower(pkg),
		PackagePath:        strings.ToLower(pkgPath),
		ResourceName:       strings.ToLower(resource.singular),
		ResourcePluralName: strings.ToLower(resource.plural),
		ParentName:         strings.ToLower(parentResource.singular),
		ParentPluralName:   strings.ToLower(parentResource.plural),
	}
}

func writeTemplate(tmpl *template.Template, data *templateData) error {
	path := fmt.Sprintf("%s/%s", data.PackagePath, tmpl.Name())

	logger.Default().Infof("writing file %s", path) // TODO: abort or warn if already exists?

	wr, err := os.Create(path)
	if err != nil {
		return err
	}
	defer wr.Close()

	return tmpl.Execute(wr, data)
}
