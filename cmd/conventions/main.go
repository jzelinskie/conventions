// Copyright 2014 The Convention Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package main implements the main entry point for a conventions application.
//
// "Go install" requires a main package to compile a binary to place in $GOBIN.
// For projects that house a library, but also want to expose a binary, the
// "cmd/binaryname/main.go" file structure is recommended. This can appear
// verbose for those that want a binary with the same name as the main package
// (i.e. "go install github.com/conventions/cmd/conventions").
package main

import (
	"fmt"

	"github.com/jzelinskie/conventions"
)

// Second to last in a file is an init func. This func enables you to run code
// at package import time. Clever usages of init include Google App Engine and
// the database/sql package.
func init() {
}

// Last in a file should be a main func. This func should only be present in a
// package named main to generate a binary file.
func main() {
	// Break long function calls into its multi-line form. Functions with lots of
	// parameters are often a code smell.
	example := conventions.NewExample(
		1,
		"conventions",
		"this is getting quite meta",
	)

	// If there is no desire to expose any variables outside of the scope of an
	// if-statement, this two-clause form is preferable.
	if _, err := example.Method(); err != nil {
		// Panics are used only in situations that inform the programmer that they
		// are doing something wrong. They should not be used like exceptions --
		// error is a builtin type for a reason. To make life easier, panic strings
		// should be prefixed with "pkgname: ".
		panic("conventions: example method should always return a nil error.")
	}

	// When a zero-value is desired, use the var keyword. Avoid using := inside
	// if statements as variables can accidentally become shadowed.
	var multiplier int
	if example.ID > 10 {
		multiplier = 10
	} else {
		multiplier = example.ID
	}
	fmt.Printf("Multiplier: %v\n", multiplier)
}
