// Copyright 2014 The Convention Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package conventions_test is used for creating Examples for documentation.
// The use of a separate package enables us to import our own package so that
// our examples look exactly like the code people will write in their packages.
//
// Examples have no documentation, however the end of every example ends with
// an output comment that tells the documentation what the expected output of
// executing that example will be.
package conventions_test

import (
	"fmt"

	"github.com/jzelinskie/conventions"
)

func ExampleNew() {
	example := conventions.NewExample(
		1,
		"conventions",
		"this is getting quite meta",
	)
	fmt.Println(example.ID)
	// Output:
	// 1
}
