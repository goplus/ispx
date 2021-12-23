# ispx
The Interpreter  of Go+ spx engine


### install ispx
Go 1.16
```
go get github.com/visualfc/ispx
```
Go 1.17
```
GOEXPERIMENT=noregabi go get -v github.com/visualfc/ispx
```


### ispx command
```
ispc [-dumpsrc|-dumppkg|-dumpssa] dir
  -dumppkg
    	print import packages
  -dumpsrc
    	print source code
  -dumpssa
    	print ssa code information
```

### run spx demo
* install ispx
```
$ GOEXPERIMENT=noregabi go get -v github.com/visualfc/ispx
```
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