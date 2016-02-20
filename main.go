package main
import (
	"net/http"
        "io"
        "net/url"
)
    
func login(w http.ResponseWriter, r *http.Request) {
   s := r.URL.String()
   u, _ := url.Parse(s) 
   m := u.Query()
   i := m["code"]   
   code := i[0]   
   io.WriteString(w, code)
}

func main() {
	http.HandleFunc("/login/", login)
	http.ListenAndServe(":1337", nil)
}
