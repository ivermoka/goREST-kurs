# Golang REST API kurs

### For vg2 IT elever

## Code snippets fra presentasjonen

### Anbefalt metode / best practice

I denne kodebiten er **"mux"** en **"request multiplexer" (forespørselsmottaker)**, som i hovedsak er en ruter som mottar innkommende **HTTP-forespørsler** og sender dem til riktig handler basert på deres URL-stier. Formålet med å bruke "mux" her er å tillate mer **organisert og fleksibel ruting** av HTTP-forespørsler til ulike håndterere basert på URL-stiene.

```
package main

import (
	"net/http"
)

func main() {
 	// Lag en ny request multiplexer
 	// Motta kommende requests og send ut til matchende handlers
 	mux := http.NewServeMux()

	// Registrer endpoints og handlers
 	mux.Handle("/", &homeHandler{})
	mux.Handle("/greeting", &greetingHandler{})

 	// Kjør serveren
 	http.ListenAndServe(":8080", mux)
}

type homeHandler struct{}

// homeHandler sin ServeHTTP method
func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
 	w.Write([]byte("Hello World!"))
}

type greetingHandler struct{}

func (h *greetingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi there!"))
}
```

### Enklere metode / versjon

Vi kan bruke en enklere metode (**uten mux**), men det er ikke anbefalt generelt. Grunnen til dette er at mux lar deg definere flere routes og dele de ut til en spesifik **handler**, som kan håndtere requests forskjellig, basert på forskjellige **paths**. Altså dette er lettere å lage i første omgang, men det vil bli vanskeligere (+ dårligere kode) å håndtere forskjellige typer forespørsler på en enklere og mer strukturert måte.

```
package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)

	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
```
