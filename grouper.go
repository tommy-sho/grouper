package main

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/imports"
)

type value struct {
	path string
	name string
}

var options = &imports.Options{
	TabWidth:  8,
	TabIndent: true,
	Comments:  true,
	Fragment:  true,
}

type Env struct {
	Paths []string

	// indicate override file or not
	Write bool

	// beginning with this string after 3rd-party packages.
	LocalPrefix string
}

var env Env

type argType int

const (
	fromStdin argType = iota

	singleArg

	multiArg
)

func grouperMain(in Env) error {
	env = in
	paths := env.Paths
	imports.LocalPrefix = env.LocalPrefix

	if len(paths) == 0 {
		if err := processFile("standard", os.Stdin, os.Stdout, fromStdin); err != nil {
			return err
		}

		return nil
	}

	argType := singleArg
	if len(paths) > 1 {
		argType = multiArg
	}

	for _, path := range paths {
		switch dir, err := os.Stat(path); {
		case err != nil:
			return err
		case dir.IsDir():
			if err := walkDir(path); err != nil {
				return err
			}
		default:
			return processFile(path, nil, os.Stdout, argType)
		}
	}

	return nil
}

func visitFile(path string, file os.FileInfo, err error) error {
	if err == nil && isGoFile(file) {
		err = processFile(path, nil, os.Stdout, multiArg)
	}
	if err != nil {
		return err
	}

	return nil
}

func isGoFile(file os.FileInfo) bool {
	if file.IsDir() {
		return false
	}

	name := file.Name()
	return !strings.HasPrefix(name, ".") && strings.HasSuffix(name, ".go")
}

func walkDir(path string) error {
	return filepath.Walk(path, visitFile)
}

func processFile(filePath string, reader io.Reader, writer io.Writer, arg argType) error {
	if reader == nil {
		f, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer f.Close()
		reader = f
	}

	src, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	target := filePath

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filePath, src, 0)
	if err != nil {
		return err
	}
	var paths []value
	ast.Inspect(f, func(n ast.Node) bool {
		if v, ok := n.(*ast.ImportSpec); ok {
			p := value{
				path: v.Path.Value,
			}
			if v.Name != nil {
				p.name = v.Name.Name
			}
			paths = append(paths, p)
		}
		return true
	})

	for _, path := range paths {
		t, _ := strconv.Unquote(path.path)
		if path.name != "" {
			astutil.DeleteNamedImport(fset, f, path.name, t)
			continue
		}

		astutil.DeleteImport(fset, f, t)
	}

	for _, path := range paths {
		t, _ := strconv.Unquote(path.path)
		astutil.AddNamedImport(fset, f, path.name, t)
	}

	var buf bytes.Buffer
	pp := &printer.Config{Tabwidth: 8, Mode: printer.UseSpaces | printer.TabIndent}
	pp.Fprint(&buf, fset, f)

	res := process(buf.Bytes(), target, options)
	if !bytes.Equal(res, src) {
		if env.Write {
			if arg == fromStdin {
				return errors.New("can't use -w option on stdin")
			}
			err = ioutil.WriteFile(target, res, 0)
			if err != nil {
				return fmt.Errorf("failed to writer result with -w: %w", err)
			}
		}
	}

	if !env.Write {
		_, err = writer.Write(res)
	}

	return err
}

func process(data []byte, name string, opt *imports.Options) []byte {
	b, _ := imports.Process(name, data, opt)
	return b
}
