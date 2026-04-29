package errors


import (
	"errors"
	"fmt"
)

var ErrorNotFound = errors.New("not found")
var ErrNotFound = errors.New("topilmadi")

type Validation struct {
	Field   string
	Message string
}

func (v *Validation) Error() string {
	return v.Field + ": " + v.Message
}

var users = map[int]string{
	1: "John Doe",
	2: "Jane Doe",
	3: "Jack Doe",
	4: "Jill Doe",
}

func GetUser(id int) (string, error) {

	if user, ok := users[id]; ok {
		return user, nil
	}
	return "", ErrorNotFound
}

func main2() {
	fmt.Println("Error handling in Go")

	user, err := GetUser(3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("User:", user)
	}

	// Xatoni o'raymiz (wrap)
	err = fmt.Errorf("bazada muammo: %w", ErrNotFound)

	// Oddiy == bilan tekshirish ish bermasligi mumkin
	// Lekin errors.Is zanjir ichini ko'ra oladi
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Xato turi: Ma'lumot topilmadi")
	}
}
