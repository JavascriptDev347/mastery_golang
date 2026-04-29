package errors

import (
    "errors"
    "fmt"
)

// Maxsus xato strukturasi
type MyError struct {
    Code    int
    Message string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("Kod: %d, Xabar: %s", e.Code, e.Message)
}

func main() {
    customErr := &MyError{Code: 404, Message: "Sahifa topilmadi"}
    err := fmt.Errorf("tizimda xato: %w", customErr)

    var targetErr *MyError
    // errors.As xatoni targetErr o'zgaruvchisiga "quyib" beradi
    if errors.As(err, &targetErr) {
        fmt.Println("Xato kodi aniqlandi:", targetErr.Code)
        fmt.Println("Xato xabari:", targetErr.Message)
    }
}