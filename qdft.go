package main

import (
    "log"
    "net/http"
	"strconv"
	"io"
	"os"
)

func transfer(w http.ResponseWriter, r *http.Request) {
	res, err := os.Open(os.Arg[1])
	if err != nil {
		return 0, err
	}
	defer res.Close()
	
	w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(os.args[1]));
	w.Header().Set("Content-Type", "application/octet-stream")
	
	io.Copy(w, res)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("No arguments.")
		os.Exit(1)
	}

	http.HandleFunc("/", transfer)
	
	log.Fatal(http.ListenAndServe(":369", nil))
}