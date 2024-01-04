// export by github.com/goplus/igop/cmd/qexp

package spx

import (
	q "github.com/goplus/spx"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "spx",
		Path: "github.com/goplus/spx",
		Deps: map[string]string{
			"encoding/json":                       "json",
			"errors":                              "errors",
			"flag":                                "flag",
			"fmt":                                 "fmt",
			"github.com/goplus/spx/fs":            "fs",
			"github.com/goplus/spx/fs/asset":      "asset",
			"github.com/goplus/spx/fs/zip":        "zip",
			"github.com/goplus/spx/internal/anim": "anim",
			"github.com/goplus/spx/internal/audiorecord": "audiorecord",
			"github.com/goplus/spx/internal/camera":      "camera",
			"github.com/goplus/spx/internal/coroutine":   "coroutine",
			"github.com/goplus/spx/internal/effect":      "effect",
			"github.com/goplus/spx/internal/gdi":         "gdi",
			"github.com/goplus/spx/internal/gdi/clrutil": "clrutil",
			"github.com/goplus/spx/internal/gdi/font":    "font",
			"github.com/goplus/spx/internal/math32":      "math32",
			"github.com/goplus/spx/internal/tools":       "tools",
			"github.com/hajimehoshi/ebiten/v2":           "ebiten",
			"github.com/hajimehoshi/ebiten/v2/audio":     "audio",
			"github.com/pkg/errors":                      "errors",
			"github.com/qiniu/audio":                     "audio",
			"github.com/qiniu/audio/convert":             "convert",
			"github.com/qiniu/audio/mp3":                 "mp3",
			"github.com/qiniu/audio/wav":                 "wav",
			"github.com/qiniu/audio/wav/adpcm":           "adpcm",
			"golang.org/x/image/colornames":              "colornames",
			"golang.org/x/image/font":                    "font",
			"image":                                      "image",
			"image/color":                                "color",
			"image/jpeg":                                 "jpeg",
			"image/png":                                  "png",
			"io":                                         "io",
			"log":                                        "log",
			"math":                                       "math",
			"math/rand":                                  "rand",
			"os":                                         "os",
			"path":                                       "path",
			"reflect":                                    "reflect",
			"strconv":                                    "strconv",
			"strings":                                    "strings",
			"sync":                                       "sync",
			"sync/atomic":                                "atomic",
			"syscall":                                    "syscall",
			"time":                                       "time",
			"unsafe":                                     "unsafe",
		},
		Interfaces: map[string]reflect.Type{
			"Gamer": reflect.TypeOf((*q.Gamer)(nil)).Elem(),
			"Shape": reflect.TypeOf((*q.Shape)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Camera":        reflect.TypeOf((*q.Camera)(nil)).Elem(),
			"Config":        reflect.TypeOf((*q.Config)(nil)).Elem(),
			"EffectKind":    reflect.TypeOf((*q.EffectKind)(nil)).Elem(),
			"Game":          reflect.TypeOf((*q.Game)(nil)).Elem(),
			"List":          reflect.TypeOf((*q.List)(nil)).Elem(),
			"MovingInfo":    reflect.TypeOf((*q.MovingInfo)(nil)).Elem(),
			"RotationStyle": reflect.TypeOf((*q.RotationStyle)(nil)).Elem(),
			"Sound":         reflect.TypeOf((*q.Sound)(nil)).Elem(),
			"Sprite":        reflect.TypeOf((*q.Sprite)(nil)).Elem(),
			"StopKind":      reflect.TypeOf((*q.StopKind)(nil)).Elem(),
			"TurningInfo":   reflect.TypeOf((*q.TurningInfo)(nil)).Elem(),
			"Value":         reflect.TypeOf((*q.Value)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{
			"Color":   reflect.TypeOf((*q.Color)(nil)).Elem(),
			"Key":     reflect.TypeOf((*q.Key)(nil)).Elem(),
			"Spriter": reflect.TypeOf((*q.Spriter)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Exit__0":              reflect.ValueOf(q.Exit__0),
			"Exit__1":              reflect.ValueOf(q.Exit__1),
			"Gopt_Game_Main":       reflect.ValueOf(q.Gopt_Game_Main),
			"Gopt_Game_Reload":     reflect.ValueOf(q.Gopt_Game_Reload),
			"Gopt_Game_Run":        reflect.ValueOf(q.Gopt_Game_Run),
			"Gopt_Sprite_Clone__0": reflect.ValueOf(q.Gopt_Sprite_Clone__0),
			"Gopt_Sprite_Clone__1": reflect.ValueOf(q.Gopt_Sprite_Clone__1),
			"Iround":               reflect.ValueOf(q.Iround),
			"RGB":                  reflect.ValueOf(q.RGB),
			"RGBA":                 reflect.ValueOf(q.RGBA),
			"Rand__0":              reflect.ValueOf(q.Rand__0),
			"Rand__1":              reflect.ValueOf(q.Rand__1),
			"Sched":                reflect.ValueOf(q.Sched),
			"SchedNow":             reflect.ValueOf(q.SchedNow),
			"SetDebug":             reflect.ValueOf(q.SetDebug),
		},
		TypedConsts: map[string]igop.TypedConst{
			"AllOtherScripts":      {reflect.TypeOf(q.AllOtherScripts), constant.MakeInt64(int64(q.AllOtherScripts))},
			"AllSprites":           {reflect.TypeOf(q.AllSprites), constant.MakeInt64(int64(q.AllSprites))},
			"BrightnessEffect":     {reflect.TypeOf(q.BrightnessEffect), constant.MakeInt64(int64(q.BrightnessEffect))},
			"ColorEffect":          {reflect.TypeOf(q.ColorEffect), constant.MakeInt64(int64(q.ColorEffect))},
			"Down":                 {reflect.TypeOf(q.Down), constant.MakeInt64(int64(q.Down))},
			"Edge":                 {reflect.TypeOf(q.Edge), constant.MakeInt64(int64(q.Edge))},
			"EdgeBottom":           {reflect.TypeOf(q.EdgeBottom), constant.MakeInt64(int64(q.EdgeBottom))},
			"EdgeLeft":             {reflect.TypeOf(q.EdgeLeft), constant.MakeInt64(int64(q.EdgeLeft))},
			"EdgeRight":            {reflect.TypeOf(q.EdgeRight), constant.MakeInt64(int64(q.EdgeRight))},
			"EdgeTop":              {reflect.TypeOf(q.EdgeTop), constant.MakeInt64(int64(q.EdgeTop))},
			"GhostEffect":          {reflect.TypeOf(q.GhostEffect), constant.MakeInt64(int64(q.GhostEffect))},
			"Key0":                 {reflect.TypeOf(q.Key0), constant.MakeInt64(int64(q.Key0))},
			"Key1":                 {reflect.TypeOf(q.Key1), constant.MakeInt64(int64(q.Key1))},
			"Key2":                 {reflect.TypeOf(q.Key2), constant.MakeInt64(int64(q.Key2))},
			"Key3":                 {reflect.TypeOf(q.Key3), constant.MakeInt64(int64(q.Key3))},
			"Key4":                 {reflect.TypeOf(q.Key4), constant.MakeInt64(int64(q.Key4))},
			"Key5":                 {reflect.TypeOf(q.Key5), constant.MakeInt64(int64(q.Key5))},
			"Key6":                 {reflect.TypeOf(q.Key6), constant.MakeInt64(int64(q.Key6))},
			"Key7":                 {reflect.TypeOf(q.Key7), constant.MakeInt64(int64(q.Key7))},
			"Key8":                 {reflect.TypeOf(q.Key8), constant.MakeInt64(int64(q.Key8))},
			"Key9":                 {reflect.TypeOf(q.Key9), constant.MakeInt64(int64(q.Key9))},
			"KeyA":                 {reflect.TypeOf(q.KeyA), constant.MakeInt64(int64(q.KeyA))},
			"KeyAlt":               {reflect.TypeOf(q.KeyAlt), constant.MakeInt64(int64(q.KeyAlt))},
			"KeyAny":               {reflect.TypeOf(q.KeyAny), constant.MakeInt64(int64(q.KeyAny))},
			"KeyApostrophe":        {reflect.TypeOf(q.KeyApostrophe), constant.MakeInt64(int64(q.KeyApostrophe))},
			"KeyB":                 {reflect.TypeOf(q.KeyB), constant.MakeInt64(int64(q.KeyB))},
			"KeyBackslash":         {reflect.TypeOf(q.KeyBackslash), constant.MakeInt64(int64(q.KeyBackslash))},
			"KeyBackspace":         {reflect.TypeOf(q.KeyBackspace), constant.MakeInt64(int64(q.KeyBackspace))},
			"KeyC":                 {reflect.TypeOf(q.KeyC), constant.MakeInt64(int64(q.KeyC))},
			"KeyCapsLock":          {reflect.TypeOf(q.KeyCapsLock), constant.MakeInt64(int64(q.KeyCapsLock))},
			"KeyComma":             {reflect.TypeOf(q.KeyComma), constant.MakeInt64(int64(q.KeyComma))},
			"KeyControl":           {reflect.TypeOf(q.KeyControl), constant.MakeInt64(int64(q.KeyControl))},
			"KeyD":                 {reflect.TypeOf(q.KeyD), constant.MakeInt64(int64(q.KeyD))},
			"KeyDelete":            {reflect.TypeOf(q.KeyDelete), constant.MakeInt64(int64(q.KeyDelete))},
			"KeyDown":              {reflect.TypeOf(q.KeyDown), constant.MakeInt64(int64(q.KeyDown))},
			"KeyE":                 {reflect.TypeOf(q.KeyE), constant.MakeInt64(int64(q.KeyE))},
			"KeyEnd":               {reflect.TypeOf(q.KeyEnd), constant.MakeInt64(int64(q.KeyEnd))},
			"KeyEnter":             {reflect.TypeOf(q.KeyEnter), constant.MakeInt64(int64(q.KeyEnter))},
			"KeyEqual":             {reflect.TypeOf(q.KeyEqual), constant.MakeInt64(int64(q.KeyEqual))},
			"KeyEscape":            {reflect.TypeOf(q.KeyEscape), constant.MakeInt64(int64(q.KeyEscape))},
			"KeyF":                 {reflect.TypeOf(q.KeyF), constant.MakeInt64(int64(q.KeyF))},
			"KeyF1":                {reflect.TypeOf(q.KeyF1), constant.MakeInt64(int64(q.KeyF1))},
			"KeyF10":               {reflect.TypeOf(q.KeyF10), constant.MakeInt64(int64(q.KeyF10))},
			"KeyF11":               {reflect.TypeOf(q.KeyF11), constant.MakeInt64(int64(q.KeyF11))},
			"KeyF12":               {reflect.TypeOf(q.KeyF12), constant.MakeInt64(int64(q.KeyF12))},
			"KeyF2":                {reflect.TypeOf(q.KeyF2), constant.MakeInt64(int64(q.KeyF2))},
			"KeyF3":                {reflect.TypeOf(q.KeyF3), constant.MakeInt64(int64(q.KeyF3))},
			"KeyF4":                {reflect.TypeOf(q.KeyF4), constant.MakeInt64(int64(q.KeyF4))},
			"KeyF5":                {reflect.TypeOf(q.KeyF5), constant.MakeInt64(int64(q.KeyF5))},
			"KeyF6":                {reflect.TypeOf(q.KeyF6), constant.MakeInt64(int64(q.KeyF6))},
			"KeyF7":                {reflect.TypeOf(q.KeyF7), constant.MakeInt64(int64(q.KeyF7))},
			"KeyF8":                {reflect.TypeOf(q.KeyF8), constant.MakeInt64(int64(q.KeyF8))},
			"KeyF9":                {reflect.TypeOf(q.KeyF9), constant.MakeInt64(int64(q.KeyF9))},
			"KeyG":                 {reflect.TypeOf(q.KeyG), constant.MakeInt64(int64(q.KeyG))},
			"KeyGraveAccent":       {reflect.TypeOf(q.KeyGraveAccent), constant.MakeInt64(int64(q.KeyGraveAccent))},
			"KeyH":                 {reflect.TypeOf(q.KeyH), constant.MakeInt64(int64(q.KeyH))},
			"KeyHome":              {reflect.TypeOf(q.KeyHome), constant.MakeInt64(int64(q.KeyHome))},
			"KeyI":                 {reflect.TypeOf(q.KeyI), constant.MakeInt64(int64(q.KeyI))},
			"KeyInsert":            {reflect.TypeOf(q.KeyInsert), constant.MakeInt64(int64(q.KeyInsert))},
			"KeyJ":                 {reflect.TypeOf(q.KeyJ), constant.MakeInt64(int64(q.KeyJ))},
			"KeyK":                 {reflect.TypeOf(q.KeyK), constant.MakeInt64(int64(q.KeyK))},
			"KeyKP0":               {reflect.TypeOf(q.KeyKP0), constant.MakeInt64(int64(q.KeyKP0))},
			"KeyKP1":               {reflect.TypeOf(q.KeyKP1), constant.MakeInt64(int64(q.KeyKP1))},
			"KeyKP2":               {reflect.TypeOf(q.KeyKP2), constant.MakeInt64(int64(q.KeyKP2))},
			"KeyKP3":               {reflect.TypeOf(q.KeyKP3), constant.MakeInt64(int64(q.KeyKP3))},
			"KeyKP4":               {reflect.TypeOf(q.KeyKP4), constant.MakeInt64(int64(q.KeyKP4))},
			"KeyKP5":               {reflect.TypeOf(q.KeyKP5), constant.MakeInt64(int64(q.KeyKP5))},
			"KeyKP6":               {reflect.TypeOf(q.KeyKP6), constant.MakeInt64(int64(q.KeyKP6))},
			"KeyKP7":               {reflect.TypeOf(q.KeyKP7), constant.MakeInt64(int64(q.KeyKP7))},
			"KeyKP8":               {reflect.TypeOf(q.KeyKP8), constant.MakeInt64(int64(q.KeyKP8))},
			"KeyKP9":               {reflect.TypeOf(q.KeyKP9), constant.MakeInt64(int64(q.KeyKP9))},
			"KeyKPDecimal":         {reflect.TypeOf(q.KeyKPDecimal), constant.MakeInt64(int64(q.KeyKPDecimal))},
			"KeyKPDivide":          {reflect.TypeOf(q.KeyKPDivide), constant.MakeInt64(int64(q.KeyKPDivide))},
			"KeyKPEnter":           {reflect.TypeOf(q.KeyKPEnter), constant.MakeInt64(int64(q.KeyKPEnter))},
			"KeyKPEqual":           {reflect.TypeOf(q.KeyKPEqual), constant.MakeInt64(int64(q.KeyKPEqual))},
			"KeyKPMultiply":        {reflect.TypeOf(q.KeyKPMultiply), constant.MakeInt64(int64(q.KeyKPMultiply))},
			"KeyKPSubtract":        {reflect.TypeOf(q.KeyKPSubtract), constant.MakeInt64(int64(q.KeyKPSubtract))},
			"KeyL":                 {reflect.TypeOf(q.KeyL), constant.MakeInt64(int64(q.KeyL))},
			"KeyLeft":              {reflect.TypeOf(q.KeyLeft), constant.MakeInt64(int64(q.KeyLeft))},
			"KeyLeftBracket":       {reflect.TypeOf(q.KeyLeftBracket), constant.MakeInt64(int64(q.KeyLeftBracket))},
			"KeyM":                 {reflect.TypeOf(q.KeyM), constant.MakeInt64(int64(q.KeyM))},
			"KeyMax":               {reflect.TypeOf(q.KeyMax), constant.MakeInt64(int64(q.KeyMax))},
			"KeyMenu":              {reflect.TypeOf(q.KeyMenu), constant.MakeInt64(int64(q.KeyMenu))},
			"KeyMinus":             {reflect.TypeOf(q.KeyMinus), constant.MakeInt64(int64(q.KeyMinus))},
			"KeyN":                 {reflect.TypeOf(q.KeyN), constant.MakeInt64(int64(q.KeyN))},
			"KeyNumLock":           {reflect.TypeOf(q.KeyNumLock), constant.MakeInt64(int64(q.KeyNumLock))},
			"KeyO":                 {reflect.TypeOf(q.KeyO), constant.MakeInt64(int64(q.KeyO))},
			"KeyP":                 {reflect.TypeOf(q.KeyP), constant.MakeInt64(int64(q.KeyP))},
			"KeyPageDown":          {reflect.TypeOf(q.KeyPageDown), constant.MakeInt64(int64(q.KeyPageDown))},
			"KeyPageUp":            {reflect.TypeOf(q.KeyPageUp), constant.MakeInt64(int64(q.KeyPageUp))},
			"KeyPause":             {reflect.TypeOf(q.KeyPause), constant.MakeInt64(int64(q.KeyPause))},
			"KeyPeriod":            {reflect.TypeOf(q.KeyPeriod), constant.MakeInt64(int64(q.KeyPeriod))},
			"KeyPrintScreen":       {reflect.TypeOf(q.KeyPrintScreen), constant.MakeInt64(int64(q.KeyPrintScreen))},
			"KeyQ":                 {reflect.TypeOf(q.KeyQ), constant.MakeInt64(int64(q.KeyQ))},
			"KeyR":                 {reflect.TypeOf(q.KeyR), constant.MakeInt64(int64(q.KeyR))},
			"KeyRight":             {reflect.TypeOf(q.KeyRight), constant.MakeInt64(int64(q.KeyRight))},
			"KeyRightBracket":      {reflect.TypeOf(q.KeyRightBracket), constant.MakeInt64(int64(q.KeyRightBracket))},
			"KeyS":                 {reflect.TypeOf(q.KeyS), constant.MakeInt64(int64(q.KeyS))},
			"KeyScrollLock":        {reflect.TypeOf(q.KeyScrollLock), constant.MakeInt64(int64(q.KeyScrollLock))},
			"KeySemicolon":         {reflect.TypeOf(q.KeySemicolon), constant.MakeInt64(int64(q.KeySemicolon))},
			"KeyShift":             {reflect.TypeOf(q.KeyShift), constant.MakeInt64(int64(q.KeyShift))},
			"KeySlash":             {reflect.TypeOf(q.KeySlash), constant.MakeInt64(int64(q.KeySlash))},
			"KeySpace":             {reflect.TypeOf(q.KeySpace), constant.MakeInt64(int64(q.KeySpace))},
			"KeyT":                 {reflect.TypeOf(q.KeyT), constant.MakeInt64(int64(q.KeyT))},
			"KeyTab":               {reflect.TypeOf(q.KeyTab), constant.MakeInt64(int64(q.KeyTab))},
			"KeyU":                 {reflect.TypeOf(q.KeyU), constant.MakeInt64(int64(q.KeyU))},
			"KeyUp":                {reflect.TypeOf(q.KeyUp), constant.MakeInt64(int64(q.KeyUp))},
			"KeyV":                 {reflect.TypeOf(q.KeyV), constant.MakeInt64(int64(q.KeyV))},
			"KeyW":                 {reflect.TypeOf(q.KeyW), constant.MakeInt64(int64(q.KeyW))},
			"KeyX":                 {reflect.TypeOf(q.KeyX), constant.MakeInt64(int64(q.KeyX))},
			"KeyY":                 {reflect.TypeOf(q.KeyY), constant.MakeInt64(int64(q.KeyY))},
			"KeyZ":                 {reflect.TypeOf(q.KeyZ), constant.MakeInt64(int64(q.KeyZ))},
			"Left":                 {reflect.TypeOf(q.Left), constant.MakeInt64(int64(q.Left))},
			"Mouse":                {reflect.TypeOf(q.Mouse), constant.MakeInt64(int64(q.Mouse))},
			"Next":                 {reflect.TypeOf(q.Next), constant.MakeInt64(int64(q.Next))},
			"OtherScriptsInSprite": {reflect.TypeOf(q.OtherScriptsInSprite), constant.MakeInt64(int64(q.OtherScriptsInSprite))},
			"Prev":                 {reflect.TypeOf(q.Prev), constant.MakeInt64(int64(q.Prev))},
			"Right":                {reflect.TypeOf(q.Right), constant.MakeInt64(int64(q.Right))},
			"ThisScript":           {reflect.TypeOf(q.ThisScript), constant.MakeInt64(int64(q.ThisScript))},
			"ThisSprite":           {reflect.TypeOf(q.ThisSprite), constant.MakeInt64(int64(q.ThisSprite))},
			"Up":                   {reflect.TypeOf(q.Up), constant.MakeInt64(int64(q.Up))},
		},
		UntypedConsts: map[string]igop.UntypedConst{
			"All":          {"untyped int", constant.MakeInt64(int64(q.All))},
			"DbgFlagAll":   {"untyped int", constant.MakeInt64(int64(q.DbgFlagAll))},
			"DbgFlagEvent": {"untyped int", constant.MakeInt64(int64(q.DbgFlagEvent))},
			"DbgFlagInstr": {"untyped int", constant.MakeInt64(int64(q.DbgFlagInstr))},
			"DbgFlagLoad":  {"untyped int", constant.MakeInt64(int64(q.DbgFlagLoad))},
			"GopPackage":   {"untyped bool", constant.MakeBool(bool(q.GopPackage))},
			"Gop_sched":    {"untyped string", constant.MakeString(string(q.Gop_sched))},
			"Invalid":      {"untyped int", constant.MakeInt64(int64(q.Invalid))},
			"Last":         {"untyped int", constant.MakeInt64(int64(q.Last))},
			"LeftRight":    {"untyped int", constant.MakeInt64(int64(q.LeftRight))},
			"None":         {"untyped int", constant.MakeInt64(int64(q.None))},
			"Normal":       {"untyped int", constant.MakeInt64(int64(q.Normal))},
			"Random":       {"untyped int", constant.MakeInt64(int64(q.Random))},
		},
	})
}
