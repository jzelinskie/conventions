// Copyright 2014 The Convention Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Every file starts with a copyright. Gophers tend to use BSD-like licenses.
// "The PackageName Authors" in copyrights along with maintained AUTHORS,
// CONTRIBUTORS, and LICENSE files are idiomatic. Comments are proper English
// sentences, always started with "//", and bound to 80 char line length. Code
// itself is not bound any line length, but long lines are often a code smell.

// Package conventions is an example Go package that is loaded with
// documentation specifically targetting new developers that have done the Go
// Tour and have read the documentation on the "go" binary, but want to know
// more about writing packages and software.
//
// The documentation for an entire package written here; only one of the files
// in a package should have it. Comments are always on the line above the thing
// they are documenting. Package declarations and any exports require comments
// and should be in the form of "Package pkgname ..." and "ExportName ..."
// respectively.
package conventions

// The import and const keywords should always be expanded into their multi-line
// form.
//
// One $GOPATH contains all of the packages on the internet. Whenever something
// is exported, the author must maintain stability of that package's API. For
// git repos, "go get" will look for a "go1" tag or fallback to the master
// branch. If an author is a bad citizen, they may break their API and those
// depending on it will have two options: update their code or fork the package.
// Forking comes at a cost: it changes the name of all of the imports,
// essentially creating a brand new package.
import (
	// Imports from the Go standard libraries are grouped together at the top.
	"errors"
	"fmt"

	// Third party imports are placed below the standard libraries. Setting up
	// text editor hooks to run gofmt and lint on save is a good idea.
	_ "github.com/golang/lint"

	// Local imports are placed below the third party libraries. Importing a
	// package as _ is useful if you require some side-effect of init() from the
	// package.
	_ "github.com/jzelinskie/conventions/subpkg"
)

// Constants always follow the imports.
const (
	enums = iota
	are
	pretty
	cool
)

// Exported variables are usually in multi-line form, however when defining
// huge literals such as the constants used in cryptographic hashing algorithms,
// it is fine to have repeated var statements.
var (
	// ErrStupidMistake is an example of an exported error type.
	ErrStupidMistake = errors.New("stupid mistakes are the story of my life")
)

// Example is an example of how to define new types.
//
// Fields are one-per-line and grouped together however you think is most
// logical. Don't worry too much about writing struct tags for every encoding
// under the sun. Someone can easily create a wrapper type that implements
// the encoding's marshaler interface (i.e. json.Marshaler).
type Example struct {
	ID          int
	Name        string
	Description string

	irrelevant string
}

// NewExample is an example constructor for the Example type.
//
// Constructors are always the first thing to follow the type declaration.
// Constructors always begin with "New" and can simply be called "New" if the
// name of the type being returned is the same as the package (i.e.
// config.New()). Returning a reference literal is better than allocating with
// the new keyword. Always have a trailing comma on the last item in anything
// multi-line.
func NewExample(id int, name, description string) *Example {
	return &Example{
		ID:          id,
		Name:        name,
		Description: description,
	}
}

// Method is an example method for the Example type.
//
// Methods are always what come after constructors. Godoc will order them
// alphabetically, but it is often more useful to order them logically.
// For most receivers, prefer pointers over values.  Exceptions to this
// are small arrays, small structs (or structs that are natural value
// types), maps, functions, and channels.
func (e Example) Method() (bool, error) {
	// The use of := for capturing the result of function calls and expressions
	// is preferrable to the use of var.
	nameLength := len(e.Name)
	result := nameLength < 10

	// Keep lines between logical groups of statements, including the final
	// return.
	return result, nil
}

// Second to last in a file is an init func. This func enables you to run code
// at package import time.
func init() {
}

// Last in a file should be a main func. This func should only be present in a
// package named main to generate a binary file.
func main() {
	// Break long function calls into its multi-line form. Functions with lots of
	// parameters are often a code smell.
	example := NewExample(
		1,
		"conventions",
		"this is getting quite meta",
	)

	_, err := example.Method()
	if err != nil {
		// Panics are used only in situations that inform the programmer that they
		// are doing something wrong. They should not be used like exceptions --
		// error is a builtin type for a reason.
		panic("conventions: example method should always return a nil error.")
	}

	// When a zero-value is desired, use the var keyword. Avoid using := inside
	// if statements as variables can accidently become shadowed.
	var multiplier int
	if example.ID > 10 {
		multiplier = 10
	} else {
		multiplier = example.ID
	}
	fmt.Printf("Multiplier: %v", multiplier)
}
