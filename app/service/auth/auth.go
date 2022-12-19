package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Hiroya3/learning-graphql/graph/model"
	"io"
	"net/http"
	"os"
)

type AuthService struct {
}

func NewAuthService() AuthService {
	return AuthService{}
}

func (a *AuthService) GetAuth(ctx context.Context, code string) (*model.AuthPayload, error) {
	clientId := os.Getenv("OAUTH_CLIENT_ID")
	secretKey := os.Getenv("OAUTH_SECRET_KEY")

	// --はじめにaccessTokenを取得する--

	tokenUrl := "https://github.com/login/oauth/access_token"

	// リクエストボディの作成
	body := struct {
		ClientId     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		Code         string `json:"code"`
	}{
		ClientId:     clientId,
		ClientSecret: secretKey,
		Code:         code,
	}

	tokenReqBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	tokenReq, err := http.NewRequest(http.MethodPost, tokenUrl, bytes.NewBuffer(tokenReqBody))
	if err != nil {
		return nil, err
	}

	tokenReq.Header.Set("Content-Type", "application/json")

	tokenClient := &http.Client{}
	tokenRes, err := tokenClient.Do(tokenReq)
	if err != nil {
		return nil, err
	}
	defer tokenRes.Body.Close()
	tokenResBody, err := io.ReadAll(tokenRes.Body)
	if err != nil {
		return nil, err
	}

	// レスポンスボディにUnmarshalする
	token := struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}{}

	err = json.Unmarshal(tokenResBody, &token)
	if err != nil {
		return nil, err
	}

	// --そしてuserを取得する--
	userUrl := "https://api.github.com/user"

	userReq, err := http.NewRequest(http.MethodGet, userUrl, nil)
	if err != nil {
		return nil, err
	}

	userReq.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token.AccessToken))

	userCli := &http.Client{}
	userRes, err := userCli.Do(userReq)
	if err != nil {
		return nil, err
	}
	defer userRes.Body.Close()

	user := struct {
		Login     string `json:"login"`
		AvatarUrl string `json:"avatarUrl"`
		Name      string `json:"name"`
	}{}
	userResBody, err := io.ReadAll(userRes.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(userResBody, &user)
	if err != nil {
		return nil, err
	}

	return &model.AuthPayload{
		Token: token.AccessToken,
		User: &model.User{
			GithubLogin:  user.Login,
			Name:         ptr(user.Name),
			Avatar:       ptr(user.AvatarUrl),
			PostedPhotos: nil,
			InPhotos:     nil,
		},
	}, nil
}

func ptr[T any](p T) *T {
	return &p
}
