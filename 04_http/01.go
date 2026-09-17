//normal server

// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintln(w, "Welcome to the server")
// 	})
// 	fmt.Println("Listening on :8080")
// 	http.ListenAndServe(":8080", nil)
// }

//FOR PRINTING LOG ERROR

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintln(w, "welcome to the store")
// 	})
// 	log.Println("listening on :8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))

// }

//For printing default serve max

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// type productHandler struct{}

// func (productHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Featured product: Notebook")
// }

// func main() {
// 	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintln(w, "Hello customer")
// 	})
// 	http.Handle("/Featured", productHandler{})

// 	log.Println("listening on :8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))

// }

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"strings"
// )

// func main() {
// 	http.HandleFunc("/greet/", func(w http.ResponseWriter, r *http.Request) {
// 		name := strings.TrimPrefix(r.URL.Path, "/greet/")
// 		if name == "" {
// 			name = "customer"
// 		}
// 		fmt.Fprintf(w, "Hello, %s. You sent a %s request.\n", name, r.Method)

//		})
//		log.Println("listening on :8080")
//		log.Fatal(http.ListenAndServe(":8080", nil))
//	}
//
//

//Query
// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
// 		query := r.URL.Query().Get("q")
// 		category := r.URL.Query().Get("category")

// 		if query == "" {
// 			fmt.Fprintln(w, "no search term given")
// 			return
// 		}
// 		fmt.Fprintf(w, "searching for %q in category %q\n", query, category)
// 	})

// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// package main

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
// 		body, err := io.ReadAll(r.Body)
// 		if err != nil {
// 			http.Error(w, "could not read body", http.StatusBadRequest)
// 			return
// 		}
// 		defer r.Body.Close()
// 		fmt.Fprintf(w, "you sent %d bytes: %s\n", len(body), body)
// 	})
// 	log.Println("listening on server")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/product/featured", func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
// 		w.Header().Set("X-Store-Region", "in")
// 		w.WriteHeader(http.StatusOK)
// 		fmt.Fprintln(w, "Featured today: leather notebook, $19.99")
// 	})
// 	http.HandleFunc("/products/missing", func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprintln(w, "product not found")
// 	})
// 	log.Fatal(http.ListenAndServe(":8080", nil))

// }

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		category := r.URL.Query().Get("category")

		if query == "" {
			fmt.Fprintln(w, "no search term given")
			return
		}
		fmt.Fprintf(w, "searching for %q in category %q\n", query, category)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
