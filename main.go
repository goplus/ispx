package main

//go:generate qexp -outdir pkg github.com/goplus/spx

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/goplus/gossa"
	"github.com/goplus/gossa/gopbuild"
	"github.com/goplus/ispx/github"
	"github.com/goplus/spx"

	_ "github.com/goplus/gossa/pkg/fmt"
	_ "github.com/goplus/gossa/pkg/math"
	_ "github.com/goplus/ispx/pkg/github.com/goplus/spx"

	_ "github.com/goplus/reflectx/icall/icall8192"
)

var (
	flagDumpSrc     bool
	flagDumpPkg     bool
	flagDumpSSA     bool
	flagGithubToken string
	flagVerbose     bool
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "ispc [flags] dir\n")
		flag.PrintDefaults()
	}
	flag.BoolVar(&flagDumpSrc, "dumpsrc", false, "print source code")
	flag.BoolVar(&flagDumpPkg, "dumppkg", false, "print import packages")
	flag.BoolVar(&flagDumpSSA, "dumpssa", false, "print ssa code information")
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
	var mode gossa.Mode
	if flagDumpPkg {
		mode |= gossa.EnableDumpPackage
	}
	if flagDumpSSA {
		mode |= gossa.EnableTracing
	}
	ctx := gossa.NewContext(mode)
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
		gossa.RegisterExternal("github.com/goplus/spx.Gopt_Game_Run", func(game spx.Gamer, resource interface{}, gameConf ...*spx.Config) {
			assert := root + "/" + resource.(string)
			fs, err := github.NewDir(client, assert)
			if err != nil {
				log.Panicln(err)
			}
			spx.Gopt_Game_Run(game, fs, gameConf...)
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
