package domain

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	Id        string
	Name      string
	Email     string
	ApiKey    string
	Balance   float64
	mu        sync.RWMutex
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAccount(name string, email string) *Account {
	account := &Account{
		Id:        uuid.New().String(),
		Name:      name,
		Email:     email,
		ApiKey:    generateApiKey(),
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return account
}

func generateApiKey() string {
	bytes := make([]byte, 16)

	rand.Read(bytes)

	return hex.EncodeToString(bytes)
}

func (a *Account) AddToBalance(amount float64) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.Balance += amount
	a.UpdatedAt = time.Now()
}
