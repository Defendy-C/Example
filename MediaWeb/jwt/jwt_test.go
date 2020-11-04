package jwt

import (
	"fmt"
	"testing"
	"time"
)

func TestJwtGenerated(t *testing.T)  {
	token, err := GenerateVToken(10, "cdw8922323")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(token))
	}

	exp, id, err := ParseVToken(token, "cdw8922323")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(time.Unix(exp, 0), id)
	}
}
