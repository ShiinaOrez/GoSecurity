package main

import (
    "fmt"
    "testing"
    "github.com/ShiinaOrez/GoSecurity/security"
)

func test(t *testing.T) {
    password := "muxi304"
    passwordHash := security.GeneratePasswordHash(password)
    fmt.Println(passwordHash)
    fmt.Println(security.CheckPasswordHash(password, passwordHash))
}
