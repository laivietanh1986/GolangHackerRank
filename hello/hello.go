package Hello

import "fmt"
import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Word From Go Website")
	})
	http.ListenAndServe(":8080", nil)
}
