package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"

    "user-api/models" // use your Go module name
)

var users []models.User
var nextID = 1

func main() {
    http.HandleFunc("/users", usersHandler)
    http.HandleFunc("/users/", userHandler)

    fmt.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    switch r.Method {
    case "GET":
        json.NewEncoder(w).Encode(users)
    case "POST":
        var u models.User
        if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }
        u.ID = nextID
        nextID++
        users = append(users, u)
        json.NewEncoder(w).Encode(u)
    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
}

func userHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    idStr := r.URL.Path[len("/users/"):]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    index := -1
    for i, u := range users {
        if u.ID == id {
            index = i
            break
        }
    }
    if index == -1 {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    switch r.Method {
    case "GET":
        json.NewEncoder(w).Encode(users[index])
    case "PUT":
        var updated models.User
        if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }
        updated.ID = id
        users[index] = updated
        json.NewEncoder(w).Encode(updated)
    case "DELETE":
        users = append(users[:index], users[index+1:]...)
        w.WriteHeader(http.StatusNoContent)
    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
}
