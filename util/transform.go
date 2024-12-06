package util

import (
	"log"
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

var transformManagerInstance Transformers

/*
Perl like search and replace
*/

func init() {
	homeDir, _ := os.UserHomeDir()

	paths := []string{"transforms.yaml", "../transforms.yaml", homeDir + "/.dough/transforms.yaml"}
	tPath := ""
	for _, p := range paths {

		_, err := os.Stat(p)
		if err == nil {
			tPath = p
			break
		}
	}
	if tPath != "" {
		data, err := os.ReadFile(tPath)
		if err != nil {
			log.Fatalf("error read %s %v", tPath, err)
		}
		yaml.Unmarshal(data, &transformManagerInstance)

		transformManagerInstance.BuildRegex()
	}
}

func NewTransformers(tranforms []Transform) *Transformers {
	t := &Transformers{Transforms: tranforms}
	t.BuildRegex()
	return t
}

type Transformers struct {
	Transforms []Transform `yaml:"transforms"`
}



func (t *Transformers) BuildRegex() {
	for i := 0; i < len(t.Transforms); i++ {
		reRawString := t.Transforms[i].Regex
		reParts := strings.Split(reRawString, "/")
		reStr := ""
		// split /^foo//
		if len(reParts) > 3 {
			reStr = reParts[1]
			t.Transforms[i].replaceValue = reParts[2]
		}
		t.Transforms[i].re = regexp.MustCompile(reStr)
	}
}

func GetTransformers() *Transformers {
	return &transformManagerInstance
}

type Transform struct {
	Name         string         `yaml:"name"`
	Field        string         `yaml:"field"`
	Regex        string         `yaml:"regex"`
	re           *regexp.Regexp `yaml:"-"`
	replaceValue string         `yaml:"-"`
}

func (t *Transform) FindReplace(field string) string {
	if t.re == nil {
		return field
	}
	return t.re.ReplaceAllString(field, t.replaceValue)
}
