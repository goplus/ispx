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
	"github.com/goplus/spx"

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

type ProxyGamer struct {
	spx.Gamer
	onEndMainEntry func()
}

func (p *ProxyGamer) MainEntry() {
	if me, ok := p.Gamer.(interface{ MainEntry() }); ok {
		me.MainEntry()
	}
	if p.onEndMainEntry != nil {
		p.onEndMainEntry()
	}
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
	mode |= igop.DisableRecover
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

	err = gopbuild.RegisterPackagePatch(ctx, "github.com/goplus/spx", `
package spx

import "github.com/goplus/spx"

// GetWidget returns the widget instance with given name. It panics if not found.
func Gopt_Game_Gopx_GetWidget[T any](sg spx.ShapeGetter, name string) *T {
	widget := spx.GetWidget_(sg, name)
	if result, ok := widget.(interface{}).(*T); ok {
		return result
	} else {
		panic("GetWidget: type mismatch")
	}
}
`)
	if err != nil {
		log.Panicln(err)
	}

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

		type Gamer interface {
			initGame(sprites []spx.Spriter) *spx.Game
		}
		gameRun := func(game spx.Gamer, resource interface{}, gameConf ...*spx.Config) {
			assert := root + "/" + resource.(string)
			fs, err := github.NewDir(client, assert)
			if err != nil {
				log.Panicln(err)
			}
			spx.Gopt_Game_Run(game, fs, gameConf...)
		}
		igop.RegisterExternal("github.com/goplus/spx.Gopt_Game_Main", func(game spx.Gamer, sprites ...spx.Spriter) {
			p := &ProxyGamer{}
			p.Gamer = game
			p.onEndMainEntry = func() {
				if me, ok := game.(interface{ IsRunning() bool }); ok {
					if !me.IsRunning() {
						gameRun(game.(spx.Gamer), "assets")
					}
				}
			}
			spx.Gopt_Game_Main(p, sprites...)
		})
		igop.RegisterExternal("github.com/goplus/spx.Gopt_Game_Run", gameRun)
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
