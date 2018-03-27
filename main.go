package directtwo

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

// Hello uses jwt
func Hello() {
	fmt.Println("hello from directtwo")
	jwt.GetSigningMethod("")
}
