package structs

type HTTPError struct {
Code int
Errors []interface{}
Message *string
}
