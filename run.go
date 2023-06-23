package cgotest

/*

#include <stdio.h>
#include <string.h>
#include <unistd.h>

void logText(char* data){
    printf("[CGO]: %s\n", data);
    fflush( stdout );
}

*/
import "C"

func logText(data string) {
	cstr := C.CString(data)
	C.logText((*C.char)(cstr))
}
