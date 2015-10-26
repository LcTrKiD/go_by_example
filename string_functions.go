// The standard library's `strings` package provides many
// useful string-related functionx. Here are some examples
// to give you a sense of the package.

package main

import x "strings"
import "fmt"

// We alias `fmt.Println` to a shorter name as we'll use
// it a lot below.
var p = fmt.Println

func string_functions() {

	// Here's a sample of the functions available in
	// `strings`. Note that these are all functions from
	// package, not methods on the string object itself.
	// This means that we need pass the string in question
	// as the first argument to the function.
        p("<string_functions>")
        p("<---------------->")
	p("Contains:  ", x.Contains("test", "es"))
	p("Count:     ", x.Count("test", "t"))
	p("HasPrefix: ", x.HasPrefix("test", "te"))
	p("HasSuffix: ", x.HasSuffix("test", "st"))
	p("Index:     ", x.Index("test", "e"))
	p("Join:      ", x.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", x.Repeat("a", 5))
	p("Replace:   ", x.Replace("foo", "o", "0", -1))
	p("Replace:   ", x.Replace("foo", "o", "0", 1))
	p("Split:     ", x.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", x.ToLower("TEST"))
	p("ToUpper:   ", x.ToUpper("test"))
	p()

	// You can find more functions in the [`strings`](http://golang.org/pkg/strings/)
	// package docx.

	// Not part of `strings` but worth mentioning here are
	// the mechanisms for getting the length of a string
	// and getting a character by index.
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
}
