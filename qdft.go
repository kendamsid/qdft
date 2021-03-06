package main

import (
    "log"
    "net/http"
	"strconv"
	"io"
	"os"
)

func transfer(w http.ResponseWriter, r *http.Request) {
	res, err := os.Open(os.Args[1])
	if err != nil {
		// fix this
		os.Exit(2)
	}
	defer res.Close()
	resStat, err := res.Stat()
	if err != nil {
		// fix this
		os.Exit(3)
	}
	
	w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(os.Args[1]));
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", strconv.FormatInt(resStat.Size(), 10))
	
	io.Copy(w, res)
	
	log.Print("Grabbed from ", r.RemoteAddr)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Invalid argument (format: qdft <filename>)")
		os.Exit(1)
	}

	http.HandleFunc("/", transfer)
	
	log.Print("Listening at port 369")
	log.Fatal(http.ListenAndServe(":369", nil))
}