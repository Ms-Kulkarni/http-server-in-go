package main

import(
"fmt"
"log"
"net/http"
)

func hellHandler(w http.ResponseWrite, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	} 
	if r.Method != "GET"{
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")


}

func main(){
	fileServer := http.FileServer(http.Dir("./static")) 
	http.HandleFunc("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandlerFunc("/hello", helloHandler)

	fmt.Prinf("strting server at port 8000\n")
	if err=:= http:ListenAndServer(":8000", nil); err|=nil{
		log.Fatal(err)
	}
} 
