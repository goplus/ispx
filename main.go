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
	"github.com/goplus/spx"

	_ "github.com/goplus/gossa/pkg/fmt"
	_ "github.com/goplus/gossa/pkg/math"
	_ "github.com/goplus/ispx/pkg/github.com/goplus/spx"

	_ "github.com/goplus/reflectx/icall/icall8192"
)

var (
	flagDumpSrc bool
	flagDumpPkg bool
	flagDumpSSA bool
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "ispc [-dumpsrc|-dumppkg|-dumpssa] dir\n")
		flag.PrintDefaults()
	}
	flag.BoolVar(&flagDumpSrc, "dumpsrc", false, "print source code")
	flag.BoolVar(&flagDumpPkg, "dumppkg", false, "print import packages")
	flag.BoolVar(&flagDumpSSA, "dumpssa", false, "print ssa code information")
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
	data, err := gopbuild.BuildDir(ctx, path)
	if err != nil {
		log.Panicln(err)
	}
	if flagDumpSrc {
		fmt.Println(string(data))
	}
	if !filepath.IsAbs(path) {
		dir, _ := os.Getwd()
		path = filepath.Join(dir, path)
	}
	gossa.RegisterExternal("github.com/goplus/spx.Gopt_Game_Run", func(game spx.Gamer, resource interface{}, gameConf ...*spx.Config) {
		os.Chdir(path)
		spx.Gopt_Game_Run(game, resource, gameConf...)
	})
	_, err = ctx.RunFile("main.go", data, nil)
	if err != nil {
		log.Panicln(err)
	}
}
