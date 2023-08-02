package date

import (
    "fmt"
    "net/http"
    "time"
    "encoding/json"
)

type ResDate struct {
    Date string `json:"date"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    d := ResDate{ time.Now().Format(time.RFC850) }
    bytes, _ := json.Marshal(d)
    fmt.Fprintf(w, string(bytes))
}