package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	model "github.com/davidyap2002/user-go/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig = &oauth2.Config{
	RedirectURL:  os.Getenv("QL_HOST") + os.Getenv("OAUTH_REDIRECT"),
	ClientID:     os.Getenv("CLIENT_ID"),
	ClientSecret: os.Getenv("CLIENT_SECRET"),
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	},
	Endpoint: google.Endpoint,
}

//GoogleOauth2RedirectLink Google Redirect Oauth2
func GoogleOauth2RedirectLink(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, googleOauthConfig.AuthCodeURL(os.Getenv(("RANDOM_STATE"))), http.StatusTemporaryRedirect)
}

//GoogleOauth2ParseCallback Google Oauth2 Callback
func GoogleOauth2ParseCallback(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	if r.FormValue("state") != os.Getenv("RANDOM_STATE") {
		http.Error(w, fmt.Sprintf("%s", "state is not valid"), http.StatusInternalServerError)
		return
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, r.FormValue("code"))
	if err != nil {
		fmt.Printf("Could Not Get Token, %s", err)
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
		return
	}

	var callbackModel model.GoogleOauth2Callback

	err = json.Unmarshal(body, &callbackModel)

	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
		return
	}

	if callbackModel.VerifiedEmail == false {
		jsonString, err := json.Marshal(gqlerror.Error{
			Message: "Email Unverified",
			Extensions: map[string]interface{}{
				"code": "BAD REQUEST",
			},
		})

		if err != nil {
			http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, fmt.Sprintf("%s", jsonString))
	}

	user, err := UserFindOrCreateByGoogleID(ctx, callbackModel.ID, callbackModel.Name, callbackModel.Email, callbackModel.Locale)

	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
	}

	signedToken, err := CreateTokenGoogle(ctx, *user.GoogleID)

	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
		return
	}

	jsonString, err := json.Marshal(signedToken)

	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%s", jsonString))
}
