package postal

import (
    "fmt"
    "net/http"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
//return requested post
    fmt.Println("Getting posts...")
}