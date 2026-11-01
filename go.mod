// Module declaration — this is your project's unique identity
// All internal imports start from this path:
//   "github.com/kumarharsh13/notes-app/note"
//   "github.com/kumarharsh13/notes-app/todo"
// Matches your GitHub repo URL so anyone can `go get` it
module github.com/kumarharsh13/notes-app

// Minimum Go version required to build this project
// Anyone with Go < 1.27.1 gets a clear error instead of a confusing failure
go 1.27.1

// No `require` block here because you have zero external dependencies
// Everything you used (fmt, os, encoding/json, strings, time, errors, bufio)
// is part of Go's standard library — no download needed
//
// When you add an external package like Gin:
//   go get github.com/gin-gonic/gin
//
// This file auto-updates to:
//   require github.com/gin-gonic/gin v1.9.1
//
// And a go.sum file appears alongside — checksums to verify
// downloaded packages haven't been tampered with
