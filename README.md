# Validator

**validator** is a lightweight, thread-safe, and extensible struct validation library for Go. It uses struct tags to define validation rules and supports custom validators out of the box.

---

## ✨ Features

- ✅ Tag-based struct validation
- 🔒 Thread-safe validator registration
- 🧩 Support for custom validation rules
- 🪶 Lightweight with no external dependencies
- ⚙️ Built-in rules for common use cases (e.g., required, email, min, max)

---

## 📦 Installation

```bash
go get github.com/godev90/validator
```

##  📘 Basic Example

Here's a complete working example using common rules like `required`, `email`, and `min`.

```go
package main

import (
    "fmt"
    "github.com/godev90/validator"
)

type User struct {
    Name  string `validate:"required"`
    Email string `validate:"required,email"`
    Age   int    `validate:"min=18"`
}

func main() {
    user := User{
        Name:  "",
	Email: "invalid-email",
	Age:   16,
    }

    err := validator.ValidateStruct(user)
    if err != nil {
	fmt.Println("Validation failed:")
	fmt.Println(err)
	return
    }

    fmt.Println("Validation passed!")
}
```

## 🔧 Advanced Example (Custom Validator)

Register your own validation `even`.

```go
package main

import (
    "fmt"
    "github.com/godev90/validator"
)

func init() {
    validator.RegisterValidator("even", func(value any, param string) error {
	num, ok := value.(int)
	if !ok {
	    return fmt.Errorf("invalid type for even check")
        }

        if num%2 != 0 {
	    return fmt.Errorf("must be an even number")
	}

	return nil
    })
}

type Transaction struct {
    Amount int    `validate:"required,even"`
    Code   string `validate:"required,min=4,max=10"`
}

func main() {
    tx := Transaction{
        Amount: 7,
        Code:   "AB",
    }

    err := validator.ValidateStruct(tx)
    if err != nil {
        fmt.Println("Validation failed:")
        fmt.Println(err)
        return
    }

    fmt.Println("Validation passed!")
}
```
