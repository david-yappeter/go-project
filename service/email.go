package service

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/davidyap2002/user-go/mailing"
	"github.com/davidyap2002/user-go/template"
)

//EmailVerification Email Verification Handler
func EmailVerification(w http.ResponseWriter, r *http.Request) {
	emailHash := r.URL.Query().Get("hash")

	if emailHash == "" {
		http.Error(w, "BAD REQUEST", http.StatusBadRequest)
	}

	_, err := UserUpdateEmailVerification(context.Background(), emailHash)

	if err != nil {
		http.Error(w, "BAD REQUEST", http.StatusBadRequest)
	}

	http.Redirect(w, r, "https://david-user-go.herokuapp.com", http.StatusTemporaryRedirect)
}

//SendEmailVerificationEmail Send Email Verification Email
func SendEmailVerificationEmail(ctx context.Context, email string, hash string) (string, error) {
	var tempBuffer bytes.Buffer

	template.AuthentificationEmail("https://david-user-go.herokuapp.com/verify/email?hash="+hash, &tempBuffer)

	go func() {
		mailing.SendEmail([]string{email}, "Email Verification", tempBuffer.String())
		fmt.Println("Successfully Send Email")
	}()

	return "Waiting GoRoutines", nil
}
