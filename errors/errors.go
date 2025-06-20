package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
)

type (
	ErrCode int
	Error   struct {
		code          ErrCode
		err           error
		localMessages map[LanguageTag]string
	}

	Errors map[string]error

	ErrAttr struct {
		Code     ErrCode
		Messages []LangPackage
	}
)

var errorLangPack = make(map[string]ErrAttr)

func registerBuiltinError(key string, args ...any) error {

	newError := Error{
		code:          http.StatusInternalServerError,
		err:           errors.New(key),
		localMessages: make(map[LanguageTag]string),
	}

	if langPack, found := errorLangPack[key]; found {
		newError.code = langPack.Code

		for _, msg := range langPack.Messages {
			newError.localMessages[msg.Tag] = fmt.Sprintf(msg.Message, args...)
		}

		newError.err = errors.New(newError.localMessages[English])
	} else {
		log.Panic("error package not found: ", key)
	}

	return &newError
}

func New(err error, attr *ErrAttr, args ...any) error {

	newError := Error{
		err:           errors.New(err.Error()),
		localMessages: make(map[LanguageTag]string),
	}

	if langPack, found := errorLangPack[err.Error()]; found {
		newError.code = langPack.Code

		for _, msg := range langPack.Messages {
			newError.localMessages[msg.Tag] = fmt.Sprintf(msg.Message, args...)
		}
	} else {
		newError.code = http.StatusInternalServerError

		if attr != nil {
			if attr.Code != 0 {
				newError.code = attr.Code
			}

			for _, msg := range attr.Messages {
				newError.localMessages[msg.Tag] = msg.Message
			}
		}
	}

	return &newError
}

func Is(err error, target error) bool {
	return errors.Is(err, target)
}

func (err Error) Code() ErrCode {
	return err.code
}

func (err Error) Error() string {
	if err.err != nil {
		if message, ok := err.localMessages[English]; ok {
			return message
		}

		return err.err.Error()
	}

	return fmt.Sprintf("something went wrong (code %d)", err.code)
}

func (err Error) LocalizedError(tag LanguageTag) string {
	if msg, found := err.localMessages[tag]; found {
		return msg
	}

	return err.Error()
}

func (errs Errors) Error() string {
	if len(errs) == 0 {
		return ""
	}

	keys := make([]string, len(errs))
	i := 0
	for key := range errs {
		keys[i] = key
		i++
	}
	sort.Strings(keys)

	var s strings.Builder
	for i, key := range keys {
		if i > 0 {
			s.WriteString("; ")
		}

		if ers, ok := errs[key].(Errors); ok {
			_, _ = fmt.Fprintf(&s, "%v: (%v)", key, ers)
		} else if er, ok := errs[key].(*Error); ok {
			if message, ok := er.localMessages[English]; ok {
				_, _ = fmt.Fprintf(&s, "%v: %v", key, message)
			} else {
				_, _ = fmt.Fprintf(&s, "%v: %v", key, er.Error())
			}
		} else {
			_, _ = fmt.Fprintf(&s, "%v: %v", key, errs[key])
		}
	}
	s.WriteString(".")

	return s.String()
}

func (errs Errors) LocalizedError(tag LanguageTag) map[string]any {
	result := make(map[string]any)

	if len(errs) == 0 {
		return result
	}

	for key, err := range errs {
		if ers, ok := err.(Errors); ok {
			result[key] = ers.LocalizedError(tag)
		} else if er, ok := errs[key].(*Error); ok {
			result[key] = er.LocalizedError(tag)
		} else {
			result[key] = err.Error()
		}
	}

	return result
}

// MarshalJSON returns the error as a JSON string
func (e Error) MarshalJSON() ([]byte, error) {
	// This will output: "some error string"
	return json.Marshal(e.Error())
}
