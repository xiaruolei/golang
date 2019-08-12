package main

import (
	"github.com/gpmgo/gopm/modules/log"
	"imooc/learngo/errhandling/filelistingserver/filelisting"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(hander appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := hander(writer, request)
		if err != nil {
			log.Error("Error handling request: %s", err.Error())
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/list/", errWrapper(filelisting.HandleFileListing))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
