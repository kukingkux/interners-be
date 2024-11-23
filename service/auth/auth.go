package auth

import (
	"net/http"

	"github.com/kukingkux/interners-be/types"
)

type User struct {
	User types.User
}

func main() {
	// sessionManager := scs.New()
}

func oAuthCallbackHandler(w http.ResponseWriter, r *http.Request) {

}
