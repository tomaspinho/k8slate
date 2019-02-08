package k8slate

import (
	"fmt"
	"os"
	"strings"

	"github.com/flosch/pongo2"
	"gopkg.in/yaml.v2"
)

// RenderedDocument - represents a k8slate document that was already rendered from is template
type RenderedDocument struct {
	Name   string
	Result string
}

type k8sNamedResource struct {
	Kind     string
	Metadata struct {
		Name string
	}
}

// Render takes a Document as input and renders it using the variables and the template
func Render(doc Document) (rdoc RenderedDocument) {

	res := pongo2.RenderTemplateString(doc.Template, doc.Preamble.Params)

	// if err != nil {
	// 	fmt.Println("There was an error reading template for file", doc, err)
	// }

	rdoc.Result = res

	return
}

// MaterializeFileName takes a Document and RenderedDocument as inputs and returns the final name for a file as per the README
func MaterializeFileName(doc Document, rdoc RenderedDocument) (name string) {

	final := k8sNamedResource{}

	err := yaml.Unmarshal([]byte(rdoc.Result), &final)

	if err != nil {
		fmt.Println("There was an error reading resulting Kubernetes Resource file. Cannot derive its name from it.")
		os.Exit(-1)
	}

	if final.Kind == "" || final.Metadata.Name == "" {
		fmt.Println("There is either no Kind or Metadata.name properties in the resulting YAML resource. Quitting...")
		os.Exit(-1)
	}

	name = strings.ToLower(final.Metadata.Name) + "-" + strings.ToLower(final.Kind) + ".yaml"

	return
}
