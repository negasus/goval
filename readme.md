# Goval

Goval is a tool for generate validation functions for Go structs.

Goval uses struct tags to generate validation functions for the struct fields.

> Note: This tool is in the early stages of development and may contain bugs.
> Also, not all types are supported yet.

## Install

Install the CLI tool

```bash
go install github.com/negasus/goval/cmd/goval@latest
```

And add the package to your project

```bash
go get github.com/negasus/goval
```

## Usage

Add tags to struct fields and run `goval` command from current directory.

```go
type Request struct {
    Name string `goval:"min=3;max=10"`
    Age  int    `goval:"min=18"`
}
```

```bash
goval --type Request
```

I recommend use `go:generate` directive.

```go
//go:generate goval --type Request
```

```bash
go generate ./...
```

Goval will generate a file `goval_request.go` with validation functions for the struct.

## CLI flags

- `-t` or `--type` - type name, you can use multiple types `-t One -t Two`, required
- `-o` or `--output` - output file name (default: `goval_{snake_cased_type}.go`)
- `-d` or `--debug` - debug mode, verbose output, default false
- `-n` or `--tag` - struct tag name, default `goval`

## Rules

Rules are defined in the struct tags with the `goval` key. You can redefine the tag name by the `--tag` CLI flag.


You can use multiple rules separated by `;`.

| Rule | Description             | Support Types                             |
|------|-------------------------|-------------------------------------------|
| min  | minimum value or length | `int`, `float64`, `string`, `any slices`  |
| max  | maximum value or length | `int`, `float64`, `string`, `any slices`  |
| in   | list of valid values    | `int`, `string`, `[]int`, `[]string`      |


For `in` rule you can define a list of valid values separated by `,`.

Also, you can use syntax `in={variable_name}` to use a variable as the list of valid values.

```go
var validKeys = []string{"foo", "bar"}

type Request struct {
	Foo string `goval:"in=foo,bar"`
	Key string `goval:"in={validKeys}"`
}
```

### Custom validation function

You can use a custom validation function.

```go
type Request struct {
    Name string `goval:"@customValidate"`
}

func (r *Request) customValidate() goval.Errors {
    errors := goval.Errors{}

    e := goval.NewError(goval.ErrorTypeInvalid).
        AddValue("field", "Custom")

    errors.Add("Custom", e)

    return errors
}
```

Custom validation function must return `goval.Errors` and have no arguments.

You can use alias `@` for call `Valiadate()` function for embedded structs.

```go
//go:generate goval --type Request --type Meta

type Meta struct {
	ID int `goval:"min=1"`
}

type Request struct {
	Meta `goval:"@"`
	Name string `goval:"min=3;max=10"`
}

```

## Error messages

Type `Error` implements `Stringer` interface and has two methods to generate error messages.

- `String()` - returns a string representation of the error with default language.
- `StringLang(ln string)` - returns a string representation of the error with the specified language.

Now the tool supports only English and Russian languages.

The default language is English.

You can redefine the default language by call `SetDefaultLang("ru")`.

### Custom error messages

You can use a custom error message for the field.

```go
var customMessageID = 1

func main() {
    goval.AddCustomMessage(customMessageID, "en", "My Message {value}")
}

type Request struct {
    Name string `goval:"@customValidate"`
}

func (r *Request) customValidate() goval.Errors {
    errors := goval.Errors{}

    e := goval.NewCustomError(customMessageID).AddValue("value", c.Index)

    errors.Add("Name", e)

    return errors
}
```

Custom Message ID must be unique and greater than 0.
