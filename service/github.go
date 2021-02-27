package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/davidyap2002/user-go/api/githubapi"
	"github.com/davidyap2002/user-go/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func errorHandling(body []byte) error {
	var githubErrorHandling model.GithubErrorHandling

	err := json.Unmarshal(body, &githubErrorHandling)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return &gqlerror.Error{
		Message: githubErrorHandling.Message,
		Extensions: map[string]interface{}{
			"code": "ERROR_RESPONSE",
		},
	}
}

//GithubGetUserRepositories Get Github User Repositories
func GithubGetUserRepositories(ctx context.Context, username string) ([]*model.UserGithubRepository, error) {
	githubClient := githubapi.InitClient(fmt.Sprintf("/users/%s/repos", username))

	body, err := githubClient.Get(ctx, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var repositories []*model.UserGithubRepository

	err = json.Unmarshal(body, &repositories)

	if err != nil {
		fmt.Println(err)

		return nil, errorHandling(body)
	}

	return repositories, nil
}
