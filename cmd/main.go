package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tomaspinho/k8slate"
)

const filepathUsage = "A filepath or fileglob expression of file(s) to template"
const defaultFileglob = "*.yaml"

func main() {

	var outputPath string

	flags := flag.NewFlagSet("list", flag.ExitOnError)

	flags.StringVar(&outputPath, "output", "output", filepathUsage)
	flags.StringVar(&outputPath, "o", "output", filepathUsage+" (shorthand of --output)")

	var Usage = func() {
		fmt.Println("Usage of k8slate:")
		flags.PrintDefaults()
		fmt.Println("  [FILE ...] string")
		fmt.Println("    	Path(s) of file(s) to template. Defaults to the expansion of fileglob *.yaml")
	}

	flags.Usage = Usage

	flags.Parse(os.Args[1:])

	files := flags.Args()
	var err error
	if len(files) == 0 {
		files, err = filepath.Glob(defaultFileglob)

		if err != nil {
			fmt.Println("Error expanding file glob: ", err)
			os.Exit(-1)
		}

	}

	//fmt.Println("FILES", files)
	//fmt.Println("OUTPUT", outputPath)

	k8slate.Mkdirp(outputPath)

	docs := k8slate.Read(files)

	// fmt.Printf("DOCS %+v\n", docs)

	for _, doc := range docs {
		templatedDoc := k8slate.Render(doc)

		// fmt.Printf("RESULT %+v\n", templatedDoc)

		fOutputPath := filepath.Join(outputPath, k8slate.MaterializeFileName(doc, templatedDoc))
		k8slate.Write(templatedDoc, fOutputPath)
	}
}
