package main

import "C"
import spam "github.com/irevenko/spamchk"

//export spamchk_is_string_spam
func spamchk_is_string_spam(str *C.char) C.int {
	is_spam := spam.IsStringSpam(C.GoString(str))
	if is_spam {
		return 1
	} else {
		return 0
	}
}

//export spamchk_is_text_file_spam
func spamchk_is_text_file_spam(file_name *C.char) C.int {
	is_spam := spam.IsStringSpam(C.GoString(file_name))
	if is_spam {
		return 1
	} else {
		return 0
	}
}


func main() {}
