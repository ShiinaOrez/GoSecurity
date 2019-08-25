package main

import (
    "fmt"
    "github.com/ShiinaOrez/GoSecurity/security"
)

func main() {
    password := "Muxi4ever"
    passwordHash := security.GeneratePasswordHash(password)
    fmt.Println(passwordHash)
    fmt.Println(security.CheckPasswordHash(password, passwordHash))
}
