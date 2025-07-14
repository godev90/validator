package faults

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
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
	log.Println("validator: Load error from file: ", filename)

	_, src, _, ok := runtime.Caller(0)
	if !ok {
		return errors.New("validator: Cannot get caller info")
	}

	data, err := os.ReadFile(fmt.Sprintf("%s/%s", filepath.Dir(src), filename))

	if err != nil {
		return err
	}

	return raw.collectErrors(data)
}

func (yml *YamlPackage) collectErrors(data []byte) error {

	if err := yaml.Unmarshal(data, yml); err != nil {
		return err
	}

	log.Println("validator: Collecting errors.")

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
