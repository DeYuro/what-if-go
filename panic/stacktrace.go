package main

import "runtime/debug"

func main()  {
	//main.showTraceString(0x4740a6, 0x24)
	//.../github.com/deyuro/what-if-go/panic/stacktrace.go:10 +0x39
	//main.main()
	///.../deyuro/what-if-go/panic/stacktrace.go:5 +0x36
	showTraceString("Show it and more for hex")
	//0x4740a6 - header value
	//HEX 0x18 - bytes in string (DEC=24)
	println("==========================================")
	//main.showTraceString(0x488de7, 0x12)
	//.../github.com/deyuro/what-if-go/panic/stacktrace.go:27 +0x5b
	//main.main()
	//.../github.com/deyuro/what-if-go/panic/stacktrace.go:14 +0x72
	showTraceString("кириллица")
	//0x488de7 - header value
	//HEX 0x12 - bytes in string (DEC=18)
}

//Use compiler directive to prevent inline function and main.showTraceString(...) representation
//go:noinline
func showTraceString(s string)  {
	defer func() {
		if r := recover(); r != nil {
			debug.PrintStack()
		}
	}()
	panic("String panic")
}


