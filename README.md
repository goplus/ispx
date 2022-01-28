# ispx
The Interpreter  of Go+ spx engine


### install ispx
```
go get github.com/visualfc/ispx
```

### ispx command
```
ispc [flags] dir
  -dumppkg
    	print import packages
  -dumpsrc
    	print source code
  -dumpssa
    	print ssa code information
  -ghtoken string
    	set github.com api token
  -v	print verbose information
```

### install spx
* install ispx
```
$ go get -v github.com/visualfc/ispx
```
### run spx demo
* get spx FlappyCalf demo
```
$ git clone https://github.com/goplus/FlappyCalf
```

* run spx demo
```
$ ispx FlappyCalf
```

* run on demo work directory
```
$ cd FlappyCalf
$ ispx .
```
### run spx remote demo

* GitHub API token

rate limiting <https://docs.github.com/en/rest/overview/resources-in-the-rest-api#rate-limiting>

token <https://github.com/settings/tokens>

```
$ ispx -ghtoken your_github_api_token github.com/goplus/FlappyCalf
```