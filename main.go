package main
import (
	"net/http"
        "io"
)
    
func login(w http.ResponseWriter, r *http.Request) {
  
   io.WriteString(w, r.URL.String())
}

func main() {
	http.HandleFunc("/login/", login)
	http.ListenAndServe(":1337", nil)
}
