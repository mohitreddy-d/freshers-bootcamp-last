//package
//
//import (
//	"errors"
//	"fmt"
//	"github.com/golang-jwt/jwt/v5"
//)
//
//func mainm() {
//	str := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4iLCJzdWIiOiIxMzQ1Njc4OTAiLCJuYW1lIjoiSm9obiBEb2UiLCJpYXQiOjE1MTYyMzkwMjJ9.A6mm79Jahn5dWvgeI19q0nghfWfmdE9Y-H7zyZYXcj8"
//
//	token, err := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) { return []byte("mohitt"), nil })
//	switch {
//	case token.Valid:
//		fmt.Println(token.Claims)
//	case errors.Is(err, jwt.ErrTokenMalformed):
//		fmt.Println("That's not even a token")
//	case errors.Is(err, jwt.ErrTokenSignatureInvalid): // Invalid signature
//		fmt.Println("Invalid signature")
//	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet): // Token is either expired or not active yet
//		fmt.Println("Timing is everything")
//	default:
//		fmt.Println("Couldn't handle this token:", err)
//	}
//}
