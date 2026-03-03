package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"pokemonBE/models"

	"golang.org/x/crypto/bcrypt"
)

type supabaseUser struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
	SaldoUang    int    `json:"saldo_uang"`
}

func RegisterUser(username, password string) (*models.User, error) {
	if _, err := getUserByUsername(username); err == nil {
		return nil, errors.New("username already taken")
	} else if err != nil && !errors.Is(err, errUserNotFound) {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	payload := map[string]interface{}{
		"username":      username,
		"password_hash": string(hash),
		"saldo_uang":    0,
	}

	var created []supabaseUser
	if err := supabaseRequest("POST", "/users", payload, &created); err != nil {
		return nil, err
	}
	if len(created) == 0 {
		return nil, errors.New("failed to create user")
	}

	return &models.User{
		ID:        created[0].ID,
		Username:  created[0].Username,
		SaldoUang: created[0].SaldoUang,
	}, nil
}

func AuthenticateUser(username, password string) (*models.User, error) {
	u, err := getUserByUsername(username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)) != nil {
		return nil, errors.New("invalid credentials")
	}

	return &models.User{
		ID:        u.ID,
		Username:  u.Username,
		SaldoUang: u.SaldoUang,
	}, nil
}

var errUserNotFound = errors.New("user not found")

func getUserByUsername(username string) (*supabaseUser, error) {
	q := url.Values{}
	q.Set("select", "id,username,password_hash,saldo_uang")
	q.Set("username", "eq."+username)

	var users []supabaseUser
	if err := supabaseRequest("GET", "/users?"+q.Encode(), nil, &users); err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errUserNotFound
	}
	return &users[0], nil
}

func supabaseRequest(method, path string, body interface{}, out interface{}) error {
	baseURL := os.Getenv("SUPABASE_URL")
	apiKey := os.Getenv("SUPABASE_API_KEY")
	if baseURL == "" || apiKey == "" {
		return errors.New("SUPABASE_URL atau SUPABASE_API_KEY belum diset")
	}

	u, err := url.Parse(baseURL)
	if err != nil {
		return fmt.Errorf("invalid SUPABASE_URL: %w", err)
	}
	fullPath := "/rest/v1" + path
	if idx := strings.Index(fullPath, "?"); idx >= 0 {
		u.Path = fullPath[:idx]
		u.RawQuery = fullPath[idx+1:]
	} else {
		u.Path = fullPath
	}

	var buf *bytes.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return err
		}
		buf = bytes.NewReader(b)
	} else {
		buf = bytes.NewReader(nil)
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return err
	}

	req.Header.Set("apikey", apiKey)
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	if method == "POST" || method == "PATCH" {
		req.Header.Set("Prefer", "return=representation")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("supabase error: status %d", resp.StatusCode)
	}

	if out == nil {
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(out)
}
