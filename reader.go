package k8slate

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Document - represents a k8slate document read from a file
type Document struct {
	Preamble struct {
		ReadParams interface{}            `yaml:"params"`
		Params     map[string]interface{} `yaml:"-"`
	}

	Template string
}

func typeCastMap(m1 map[interface{}]interface{}) (m2 map[string]interface{}) {
	m2 = make(map[string]interface{})
	for k, v := range m1 {
		key := k.(string)
		m2[key] = v
	}
	return
}

// Read - Reads files from their paths and returns Documents
func Read(files []string) (documents []Document) {

	for _, fp := range files {
		f, err := ioutil.ReadFile(fp)

		if err != nil {
			fmt.Println("There was an error reading the file", err)
			os.Exit(-1)
		}

		yamlDocumentsInFile := bytes.SplitN(f, []byte("---\n"), -1)
		//fmt.Printf("%q\n", yamlDocumentsInFile)

		if (len(yamlDocumentsInFile) % 2) != 0 {
			fmt.Println("File ", fp, " has an odd number of documents. File must consist of pairs of preamble and template documents, in order.")
			os.Exit(-1)
		}

		for i := 0; i < len(yamlDocumentsInFile); i += 2 {

			doc := Document{}
			err = yaml.Unmarshal(yamlDocumentsInFile[i], &doc.Preamble)
			doc.Template = string(yamlDocumentsInFile[i+1])

			if err != nil {
				fmt.Println("There was an error unmarshaling yaml", err)
				os.Exit(-1)
			}

			//fmt.Printf("%+v\n", doc)

			// Perform type conversions to handle lists of maps or single map
			switch p := doc.Preamble.ReadParams.(type) {
			case []interface{}:
				for _, params := range p {

					// We cannot derive a map[string]inteface{} from interface{} directly
					paramsMap, _ := params.(map[interface{}]interface{})

					tParams := typeCastMap(paramsMap)

					document := Document{}
					document.Preamble.Params = tParams
					document.Template = doc.Template

					documents = append(documents, document)
				}
			case interface{}:
				// We cannot derive a map[string]inteface{} from interface{} directly
				tParams := p.(map[interface{}]interface{})

				doc.Preamble.Params = typeCastMap(tParams)

				documents = append(documents, doc)
			default:
				fmt.Printf("I don't know how to deal with type %T %+v!\n", p, p)
				os.Exit(-1)
			}

		}

	}

	return
}
