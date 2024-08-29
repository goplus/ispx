package main

//go:generate qexp -outdir pkg github.com/goplus/spx

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/goplus/igop"
	"github.com/goplus/igop/gopbuild"
	"github.com/goplus/ispx/github"
	spxfs "github.com/goplus/spx/fs"

	_ "github.com/goplus/igop/pkg/fmt"
	_ "github.com/goplus/igop/pkg/math"
	_ "github.com/goplus/ispx/pkg/github.com/goplus/spx"
	_ "github.com/goplus/reflectx/icall/icall8192"
)

var (
	flagDumpSrc     bool
	flagTrace       bool
	flagDumpSSA     bool
	flagDumpPKG     bool
	flagGithubToken string
	flagVerbose     bool
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "ispc [flags] dir\n")
		flag.PrintDefaults()
	}
	flag.BoolVar(&flagDumpSrc, "dumpsrc", false, "print source code")
	flag.BoolVar(&flagDumpSSA, "dumpssa", false, "print ssa code information")
	flag.BoolVar(&flagDumpPKG, "dumppkg", false, "print import pkgs")
	flag.BoolVar(&flagTrace, "trace", false, "trace")
	flag.BoolVar(&flagVerbose, "v", false, "print verbose information")
	flag.StringVar(&flagGithubToken, "ghtoken", "", "set github.com api token")
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
		return
	}
	path := args[0]
	var mode igop.Mode
	if flagDumpSSA {
		mode |= igop.EnableDumpInstr
	}
	if flagTrace {
		mode |= igop.EnableTracing
	}
	if flagDumpPKG {
		mode |= igop.EnableDumpImports
	}
	ctx := igop.NewContext(mode)
	var (
		data []byte
		err  error
	)
	if root, ok := github.IsSupport(path); ok {
		if flagVerbose {
			github.Verbose = true
		}
		client := github.NewClient(flagGithubToken)
		fs, err := github.NewFileSystem(client, root)
		if err != nil {
			log.Fatalln("error", err)
		}
		if flagVerbose {
			log.Println("BuildDir", root)
		}
		data, err = gopbuild.BuildFSDir(ctx, fs, root)
		if err != nil {
			log.Panicln(err)
		}
		spxfs.RegisterSchema("", func(file string) (spxfs.Dir, error) {
			assert := root + "/" + file
			fs, err := github.NewDir(client, assert)
			if err != nil {
				log.Panicln("load fs error", err)
			}
			return fs, nil
		})
	} else {
		if flagVerbose {
			log.Println("BuildDir", path)
		}
		data, err = gopbuild.BuildDir(ctx, path)
		if err != nil {
			log.Panicln(err)
		}
		if !filepath.IsAbs(path) {
			dir, _ := os.Getwd()
			path = filepath.Join(dir, path)
		}
		os.Chdir(path)
		if flagVerbose {
			log.Println("Chdir", path)
		}
	}
	if flagDumpSrc {
		fmt.Println(string(data))
	}
	_, err = ctx.RunFile("main.go", data, nil)
	if err != nil {
		log.Panicln(err)
	}
}
