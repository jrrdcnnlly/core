package auth

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/jrrdcnnlly/core/sessions"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/confidential"
)

type MSALClient struct {
	confidential.Client
	authority    string
	clientId     string
	clientSecret string
	scopes       []string
	redirectURL  *url.URL
}

func NewMSALClient(
	authority string,
	clientId string,
	clientSecret string,
	scopes []string,
	redirectURL *url.URL,
) (*MSALClient, error) {
	cred, err := confidential.NewCredFromSecret(clientSecret)
	if err != nil {
		return nil, fmt.Errorf("NewMSALClient: %w", err)
	}

	client, err := confidential.New(authority, clientId, cred)
	if err != nil {
		return nil, fmt.Errorf("NewMSALClient: %w", err)
	}

	msal := &MSALClient{
		Client:       client,
		authority:    authority,
		clientId:     clientId,
		clientSecret: clientSecret,
		scopes:       scopes,
		redirectURL:  redirectURL,
	}

	return msal, nil
}

func (c *MSALClient) refreshToken(ctx context.Context, s *sessions.Session) error {
	account, err := c.account(ctx, s)
	if err != nil {
		return fmt.Errorf("MSALClient.refreshToken(): failed to retrieve account: %w", err)
	}

	result, err := c.Client.AcquireTokenSilent(ctx, c.scopes, confidential.WithSilentAccount(account))
	if err != nil {
		return fmt.Errorf("MSALClient.refreshToken(): failed to refresh token: %w", err)
	}

	s.Expires = result.ExpiresOn
	s.UserID = result.Account.HomeAccountID
	s.Username = result.Account.PreferredUsername
	return nil
}

func (c *MSALClient) account(ctx context.Context, s *sessions.Session) (confidential.Account, error) {
	if s.UserID == "" {
		return confidential.Account{}, errors.New("session is not authenticated")
	}
	account, err := c.Client.Account(ctx, s.UserID)
	if err != nil {
		return confidential.Account{}, fmt.Errorf("MSALClient.Account(): %w", err)
	}
	return account, nil
}

func (c *MSALClient) authCodeURL(ctx context.Context) (string, error) {
	return c.Client.AuthCodeURL(ctx, c.clientId, c.redirectURL.String(), c.scopes)
}

func (c *MSALClient) acquireToken(ctx context.Context, code string, s *sessions.Session) error {
	result, err := c.Client.AcquireTokenByAuthCode(ctx, code, c.redirectURL.String(), c.scopes)
	if err != nil {
		return fmt.Errorf("MSALClient.acquireToken(): %w", err)
	}

	s.Expires = result.ExpiresOn
	s.UserID = result.Account.HomeAccountID
	s.Username = result.Account.PreferredUsername
	return nil
}
