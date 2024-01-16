package ifs

import (
	"fmt"
	"strconv"
	"syscall/js"
	"time"
)

func getFilesStartingWith(dirname string) (ret []string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("getFilesStartingWith panic: %v", r)
		}
	}()

	// Get JavaScript global object
	jsGlobal := js.Global()

	// Get the JavaScript function we want to call
	jsFunc := jsGlobal.Get("getFilesStartingWith")
	// Check if function is defined
	if jsFunc.Type() == js.TypeUndefined {
		err = fmt.Errorf("getFilesStartingWith function is not defined")
		return
	}
	// Call a JavaScript function, passing dirname as argument
	promise := jsFunc.Invoke(dirname)
	// Prepare a channel for receiving results
	done := make(chan []string)
	var files []string

	// Define success callback function
	onSuccess := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// Get the results returned by JavaScript
		jsResult := args[0]
		length := jsResult.Length()
		files = make([]string, length)
		for i := 0; i < length; i++ {
			files[i] = jsResult.Index(i).String()
		}

		// Send results through channel
		done <- files
		return nil
	})

	// Define failure callback function
	onError := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		err = fmt.Errorf("error calling getFilesStartingWith")
		done <- nil
		return nil
	})

	// Bind callback function to Promise
	promise.Call("then", onSuccess)
	promise.Call("catch", onError)
	// Wait for Promise to resolve
	ret = <-done

	// Clean up callbacks
	onSuccess.Release()
	onError.Release()
	if ret == nil {
		err = fmt.Errorf("error reading directory from IndexedDB, result is empty")
	}
	return
}

func readFileFromIndexedDB(filename string) (ret []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("readFileFromIndexedDB panic: %v", r)
		}
	}()

	// Get JavaScript global object
	jsGlobal := js.Global()

	// Get the JavaScript function we want to call
	jsFunc := jsGlobal.Get("readFileFromIndexedDB")

	// Call a JavaScript function, passing filename as argument
	promise := jsFunc.Invoke(filename)

	// Prepare a channel for receiving results
	done := make(chan []byte)

	// Define success callback function
	onSuccess := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// Handle the returned ArrayBuffer
		jsArrayBuffer := args[0]
		length := jsArrayBuffer.Get("byteLength").Int()
		fileContent := make([]byte, length)
		js.CopyBytesToGo(fileContent, jsArrayBuffer)

		done <- fileContent
		return nil
	})

	// Define failure callback function
	onError := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		err = fmt.Errorf("error reading file from IndexedDB")
		done <- nil
		return nil
	})

	// Bind callback function to Promise
	promise.Call("then", onSuccess)
	promise.Call("catch", onError)

	// Wait for Promise to resolve
	ret = <-done
	// Clean up callbacks
	onSuccess.Release()
	onError.Release()

	if ret == nil {
		err = fmt.Errorf("error reading file from IndexedDB, result is empty")
	}
	return
}

type FileProperties struct {
	Size         int64
	LastModified time.Time
}

func getFileProperties(filename string) (ret FileProperties, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("getFileProperties panic: %v", r)
		}
	}()

	// Get JavaScript global object
	jsGlobal := js.Global()

	// Get the JavaScript function we want to call
	jsFunc := jsGlobal.Get("getFileProperties")

	// Call a JavaScript function, passing filename as argument
	promise := jsFunc.Invoke(filename)

	// Prepare a channel for receiving results
	done := make(chan FileProperties)
	var fileProperties FileProperties

	// Define success callback function
	onSuccess := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// Handle the returned object
		jsResult := args[0]
		size, _ := strconv.ParseInt(jsResult.Get("size").String(), 10, 64)
		lastModifiedMillis, _ := strconv.ParseInt(jsResult.Get("lastModified").String(), 10, 64)
		lastModified := time.Unix(0, lastModifiedMillis*int64(time.Millisecond))

		fileProperties = FileProperties{
			Size:         size,
			LastModified: lastModified,
		}

		done <- fileProperties
		return nil
	})

	// Define failure callback function
	onError := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		err = fmt.Errorf("error calling getFileProperties")
		done <- FileProperties{}
		return nil
	})

	// Bind callback function to Promise
	promise.Call("then", onSuccess)
	promise.Call("catch", onError)

	// Wait for Promise to resolve
	ret = <-done

	// Clean up callbacks
	onSuccess.Release()
	onError.Release()

	if ret.Size == 0 && ret.LastModified.IsZero() {
		err = fmt.Errorf("error getting file properties, result is empty")
	}
	return
}
