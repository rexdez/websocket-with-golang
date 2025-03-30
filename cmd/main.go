package cmd

import "net/http"


func main() {
	server := http.Server{
		Addr: ":4867",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	
}