module github.com/goplus/ispx

go 1.16

require (
	github.com/google/go-github/v41 v41.0.0
	github.com/goplus/igop v0.7.0
	github.com/goplus/reflectx v0.8.10
	github.com/goplus/spx v1.0.0-rc5
	golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be
)

replace (
	github.com/hajimehoshi/oto => github.com/hajimehoshi/oto v1.0.1
	github.com/srwiley/oksvg => github.com/qiniu/oksvg v0.2.0-no-charset
	golang.org/x/image => golang.org/x/image v0.0.0-20210628002857-a66eb6448b8d
	golang.org/x/mobile => golang.org/x/mobile v0.0.0-20210902104108-5d9a33257ab5
	golang.org/x/mod => golang.org/x/mod v0.5.1
)

replace github.com/goplus/gop => github.com/goplus/gop v1.1.0-alpha1.0.20220603113329-36a13b025cc0
