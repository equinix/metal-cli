package infrastructure

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

type TokenExchanger struct {
	metalAPIKey string
	client      *http.Client
}

var _ oauth2.TokenSource = (*TokenExchanger)(nil)

func NewTokenExchanger(metalAPIKey string, client *http.Client) *TokenExchanger {
	if client == nil {
		client = http.DefaultClient
	}
	return &TokenExchanger{
		metalAPIKey: metalAPIKey,
		client:      client,
	}
}

func (m TokenExchanger) Token() (*oauth2.Token, error) {
	tokenExchangeURL := "https://iam.metalctrl.io/api-keys/exchange"
	tokenExchangeRequest, err := http.NewRequest("POST", tokenExchangeURL, nil)
	if err != nil {
		return nil, err
	}
	tokenExchangeRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %v", m.metalAPIKey))

	resp, err := m.client.Do(tokenExchangeRequest)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("token exchange request failed with status %v, body %v", resp.StatusCode, string(body[:]))
	}

	token := oauth2.Token{}
	err = json.Unmarshal(body, &token)
	if err != nil {
		return nil, err
	}

	// TODO: expires_in does not appear to be returned. Only access_token
	expiresIn := token.Extra("expires_in")
	if expiresIn != nil {
		expiresInSeconds := expiresIn.(int)
		token.Expiry = time.Now().Add(time.Second * time.Duration(expiresInSeconds))
	}

	return &token, nil
}
