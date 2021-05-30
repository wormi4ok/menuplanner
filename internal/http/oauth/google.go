package oauth

import (
	"context"
	"encoding/json"
	"io"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const userInfoEndpoint = "https://openidconnect.googleapis.com/v1/userinfo"

type Google struct {
	ClientID     string
	ClientSecret string
}

type googleResponse struct {
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Locale        string `json:"locale"`
}

func (config *Google) UserInfo(ctx context.Context, authCode string) (*googleResponse, error) {
	var userInfo googleResponse

	conf := &oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		RedirectURL:  "postmessage",
		Endpoint:     google.Endpoint,
	}

	token, err := conf.Exchange(ctx, authCode)
	if err != nil {
		return nil, err
	}

	authClient := conf.Client(ctx, token)

	res, err := authClient.Get(userInfoEndpoint)
	if err != nil {
		return nil, err
	}

	if body, err := io.ReadAll(res.Body); err != nil {
		return nil, err
	} else {
		_ = res.Body.Close()
		if err = json.Unmarshal(body, &userInfo); err != nil {
			return nil, err
		}
	}
	return &userInfo, nil
}
