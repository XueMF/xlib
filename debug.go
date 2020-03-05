package xlib

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

// PrintVar print variable with log package,
// when the debugging is turned on
func PrintVar(debuging bool, varName string, varValue ...interface{}) {
	if debuging {
		return
	}
	fpcs := make([]uintptr, 1)
	n := runtime.Callers(2, fpcs)
	if n == 0 {
		log.Println("MSG: NO CALLER")
	}
	caller := runtime.FuncForPC(fpcs[0] - 1)
	if caller == nil {
		log.Println("MSG CALLER WAS NIL")
	}
	// Print the file name and line number
	sourceFile, lineNumber := caller.FileLine(fpcs[0] - 1)
	log.Printf("Source file:%v line:%v", sourceFile, lineNumber)

	// Print the name of the function
	callerName := caller.Name()
	fmt.Fprintf(os.Stderr, "From func: %v\n", callerName)
	fmt.Fprintf(os.Stderr, "\x1b[31m%v = \x1b[1m%v \x1b[0m\n", varName, varValue)
}

// Logger print the value of
func Logger(preMsg string, variable interface{}) {
	fmt.Printf(preMsg+":%v\n", variable)
}

func getFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

// PageBreak print ---------- ss ------------
func PageBreak(msg ...interface{}) {
	fmt.Print("--------- ")
	for _, m := range msg {
		fmt.Print(m)
		fmt.Print(" ")
	}
	fmt.Print("---------\n")
}
