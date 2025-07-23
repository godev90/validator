package faults

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

type YamlPackage struct {
	Errors   map[string]map[LanguageTag]string `yaml:"errors"`
	Packages map[string]ErrAttr
}

func NewYamlPackage() YamlPackage {
	return YamlPackage{
		Errors:   make(map[string]map[LanguageTag]string),
		Packages: make(map[string]ErrAttr),
	}
}

func (raw YamlPackage) LoadYaml(filename string) error {
	log.Printf("validator: Attempting to load YAML file %q", filename)

	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("validator: Failed to read %q (%w)", filename, err)
	}

	return raw.collectErrors(data)
}

func (raw *YamlPackage) LoadBytes(yamlBytes []byte) error {
	return raw.collectErrors(yamlBytes)
}

func (yml *YamlPackage) collectErrors(data []byte) error {

	if err := yaml.Unmarshal(data, yml); err != nil {
		return err
	}

	log.Println("validator: Collecting errors from yaml.")

	for key, val := range yml.Errors {
		code, _ := strconv.Atoi(val["code"])
		log.Printf("validator: registering: %s", key)
		yml.Packages[key] = ErrAttr{
			Code: ErrCode(code),
			Messages: []LangPackage{
				{Tag: English, Message: val[English]},
				{Tag: Bahasa, Message: val[Bahasa]},
			},
		}
	}

	log.Println("validator: Package successfully loaded.")

	return nil
}

func (yml YamlPackage) NewError(key string) Error {
	errmsg := fmt.Sprintf("validator errorx: %s.", key)
	err := Error{
		code:          http.StatusInternalServerError,
		err:           errors.New(errmsg),
		localMessages: make(map[LanguageTag]string),
	}

	if langPack, found := yml.Packages[key]; found {
		err.code = langPack.Code

		for _, msg := range langPack.Messages {
			err.localMessages[msg.Tag] = msg.Message
		}
	}

	return err
}
