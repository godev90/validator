package errors

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"gopkg.in/yaml.v3"
)

type rawYAMLErrors struct {
	HttpStatus int                               `yaml:"http_status"`
	Errors     map[string]map[LanguageTag]string `yaml:"errors"`
}

func loadYamlFile(filename string) error {
	log.Print("load built-in error file: ", filename)

	_, src, _, ok := runtime.Caller(0)
	if !ok {
		panic("Cannot get caller info")
	}

	data, err := os.ReadFile(fmt.Sprintf("%s/%s", filepath.Dir(src), filename))

	if err != nil {
		panic(err)
	}

	return collectBuildinErrors(data)
}

func collectBuildinErrors(data []byte) error {
	var raw rawYAMLErrors
	if err := yaml.Unmarshal(data, &raw); err != nil {
		log.Panic(err)
	}

	log.Print("collecting error...")

	for key, val := range raw.Errors {
		code, _ := strconv.Atoi(val["code"])
		log.Printf("registering: %s", key)
		errorLangPack[key] = ErrAttr{
			Code: ErrCode(code),
			Messages: []LangPackage{
				{Tag: English, Message: val[English]},
				{Tag: Bahasa, Message: val[Bahasa]},
			},
		}
	}

	log.Print("built-in error load completed.")

	return nil
}
