package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

//CheckEmail Check Email Status
func CheckEmail(ctx context.Context, email string) (bool, error) {

	client := http.Client{
		Timeout: time.Second * 10,
	}

	url := "https://isitarealemail.com/api/email/validate"

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		fmt.Println(err)
		return false, err
	}

	q := req.URL.Query()
	q.Add("email", email)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return false, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return false, err
	}

	var respStruct struct {
		Status string `json:"status"`
	}

	err = json.Unmarshal(body, &respStruct)

	if err != nil {
		fmt.Println(err)
		return false, err
	}

	if strings.ToLower(respStruct.Status) != "valid" {
		return false, &gqlerror.Error{
			Message: "Email Unavailable",
			Extensions: map[string]interface{}{
				"code": "Unavailable",
			},
		}
	}

	return true, nil
}
