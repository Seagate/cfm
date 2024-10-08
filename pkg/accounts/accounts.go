// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package accounts

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/slices"

	"k8s.io/klog/v2"
)

const (
	RoleAdministrator string = "Administrator"
	RoleOperator      string = "Operator"
	RoleReadOnly      string = "ReadOnly"
	DefaultUsername   string = "admin"
	DefaultPassword   string = "admin12345"
	MinPasswordLength int64  = 8
)

var RoleTypes []string = []string{RoleAdministrator, RoleOperator, RoleReadOnly}

// Accounts structure
type AccountSingle struct {
	Id                      string `json:"id"`
	Username                string `json:"username"`
	Password                string `json:"password"`
	Role                    string `json:"role"`
	Enabled                 bool   `json:"enabled"`
	Locked                  bool   `json:"locked"`
	PasswordChangedRequired bool   `json:"passwordChangeRequired"`
}

type AccountsData struct {
	Accounts []*AccountSingle `json:"accounts"`
}

// Return our singleton accounts handler
func AccountsHandler() *accountsHandler {
	storageFileName := "cxl-host-accounts.json"
	pwd, _ := os.Getwd()
	if pwd == "/" { // executable run as systemd service
		storageFileName = "/etc/cxl-host-accounts.json"
	}
	once.Do(func() {
		a = &accountsHandler{
			storage: storageFileName,
			accounts: AccountsData{
				Accounts: make([]*AccountSingle, 0),
			},
		}
	})
	return a
}

// The Singleton instance of the Accounts Handler
var (
	a    *accountsHandler
	once sync.Once
)

// Accounts control for storing
type accountsHandler struct {
	storage  string       // The persistent storage used to store accounts
	accounts AccountsData // All accounts
	logger   klog.Logger  // The current logger
}

// IsPasswordValid: Return true if the password meets the minimum requirements
func IsPasswordValid(password string) bool {
	valid := false

	if int64(len(password)) >= MinPasswordLength {
		valid = true
	}

	return valid
}

// GeneratePasswordHash: Generate a secure storage form for the password
func GeneratePasswordHash(password string) (string, error) {
	hash := ""

	if !IsPasswordValid(password) {
		return "", fmt.Errorf("Password does not meet minimum requirements: MinPasswordLength: %d", MinPasswordLength)
	}

	// Hashing the password with the default cost of 10
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Unable to generate hash from password: length: %d, error: %s", len(password), err)
	} else {
		hash = string(hashed)
	}

	return hash, nil
}

// InitLogger: Initialize the audit handler logger
func (a *accountsHandler) InitLogger(logger klog.Logger) (err error) {
	a.logger = logger
	return nil
}

// Clear: Remove all account entries
func (a *accountsHandler) Clear() (err error) {
	a.accounts.Accounts = []*AccountSingle{}
	return nil
}

// Default: Set up a single default account
func (a *accountsHandler) Default() {
	a.logger.V(4).Info("accounts default")
	_, err := a.AddAccount(DefaultUsername, DefaultPassword, RoleAdministrator)
	if err != nil {
		a.logger.Error(err, "accounts default")
	}
}

// GetPasswordHash: Return the password hash for a username
func (a *accountsHandler) GetPasswordHash(username string) string {
	hash := ""
	for _, account := range a.accounts.Accounts {
		if account.Username == username {
			hash = account.Password
		}
	}
	return hash
}

// ArePasswordsEqual: Check to see if the password string matches the hashed password for a user
func (a *accountsHandler) ArePasswordsEqual(username, password string) bool {
	equal := false
	// Comparing the password with the hash
	err := bcrypt.CompareHashAndPassword([]byte(a.GetPasswordHash(username)), []byte(password))
	if err == nil {
		equal = true
	}
	return equal
}

// ValidAccount: Check to see this is a valid username and password
func (a *accountsHandler) ValidAccount(username, password string) (bool, error) {
	// Compare the hash and the password
	hash := a.GetPasswordHash(username)
	if hash != "" {
		err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
		if err == nil {
			return true, nil
		} else {
			return false, fmt.Errorf("stored password does not match entered password")
		}
	} else {
		return false, fmt.Errorf("unable to locate hash for username (%s)", username)
	}

}

// GetAccount: returns a single account
func GetAccount(acountId string) *AccountSingle {
	var account *AccountSingle = nil

	// Determine is username already exists in the global table
	for _, a := range a.accounts.Accounts {
		if a.Username == acountId {
			account = a
			break
		}
	}

	return account
}

// AddAccount: Add a new account entry and return the account id
func (a *accountsHandler) AddAccount(username, password, role string) (*AccountSingle, error) {
	a.logger.V(4).Info("add account", "username", username, "role", role)

	hash, err := GeneratePasswordHash(password)

	if err != nil {
		return nil, fmt.Errorf("Unable to generate password hash for account (%s), error: %s", username, err)
	}

	// Determine is username already exists in the table
	for _, account := range a.accounts.Accounts {
		if account.Username == username {
			return nil, fmt.Errorf("Unable to add account, already exists for (%s)", username)
		}
	}

	// If account was not found, add a new entry
	uuid := uuid.New().String()
	id := uuid[24:]

	account := &AccountSingle{
		Id:                      id,
		Username:                username,
		Password:                hash,
		Enabled:                 true,
		Locked:                  false,
		PasswordChangedRequired: false,
	}

	// Validate the role value before storing
	if slices.Contains(RoleTypes, role) {
		account.Role = role
	} else {
		account.Role = RoleReadOnly
		a.logger.Error(fmt.Errorf("role value (%s) is not acceptable, defaulting to (%s)", role, account.Role), "")
	}

	a.accounts.Accounts = append(a.accounts.Accounts, account)

	// Store the updated data
	err = a.Store()
	if err != nil {
		return nil, fmt.Errorf("Unable to store account data to (%s) error: %s", a.storage, err)
	}

	return account, nil
}

// UpdateAccount: Update account entry
func (a *accountsHandler) UpdateAccount(username, password, role string) (string, error) {
	a.logger.V(4).Info("update account", "username", username, "role", role)

	id := ""
	hash, err := GeneratePasswordHash(password)

	if err != nil {
		return "", fmt.Errorf("Unable to generate password hash for account (%s), error: %s", username, err)
	}

	// Determine is username already exists in the table
	found := false
	for _, account := range a.accounts.Accounts {
		if account.Username == username {
			a.logger.V(4).Info("updating found account", "username", username)
			found = true
			// Update this entry
			id = account.Id
			account.Password = hash
			if role != "" {
				if slices.Contains(RoleTypes, role) {
					account.Role = role
				} else {
					account.Role = RoleReadOnly
					a.logger.Error(fmt.Errorf("role value (%s) is not acceptable, defaulting to (%s)", role, account.Role), "")
				}
			}
		}
	}

	if !found {
		return "", fmt.Errorf("Unable to find account (%s)", username)
	}

	// Store the updated data
	err = a.Store()
	if err != nil {
		return id, fmt.Errorf("Unable to store account data to (%s) error: %s", a.storage, err)
	}

	a.logger.V(4).Info("update account SUCCESS")

	return id, nil
}

// DeleteAccount: Delete account entry
func (a *accountsHandler) DeleteAccount(username string) error {
	a.logger.V(4).Info("delete account", "username", username)

	if username == DefaultUsername {
		return fmt.Errorf("Not allowed to delete the default username (%s)", username)
	}

	// Find the username in the table
	found := false
	for index, account := range a.accounts.Accounts {
		if account.Username == username {
			// Delete this entry
			a.logger.V(5).Info("accounts before", "account", a.accounts.Accounts)
			a.accounts.Accounts[index] = a.accounts.Accounts[len(a.accounts.Accounts)-1]
			a.accounts.Accounts = a.accounts.Accounts[:len(a.accounts.Accounts)-1]
			a.logger.V(5).Info("accounts after", "account", a.accounts.Accounts)
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("Unable to find account (%s)", username)
	}

	// Store the updated data
	err := a.Store()
	if err != nil {
		return fmt.Errorf("Unable to store account data to (%s) error: %s", a.storage, err)
	}

	return nil
}

// Store: Save all accounts to a JSON file, needs to be done after every account update
func (a *accountsHandler) Store() (err error) {
	a.logger.V(4).Info("Store accounts", "storage", a.storage)

	// Convert object into JSON format
	js, err := json.MarshalIndent(a.accounts, "", " ")
	if err != nil {
		return fmt.Errorf("unable to translate accounts to JSON, error: %v", err)
	}

	// Write report to file
	// Set permissions so that owner can read/write (6), group can read (first 4), all others can read (second 4)
	err = os.WriteFile(a.storage, js, 0o644)
	if err != nil {
		return fmt.Errorf("unable to write accounts to storage, error: %v", err)
	}

	return nil
}

// Restore: Read all accounts from a JSON file
func (a *accountsHandler) Restore() (err error) {
	a.logger.V(4).Info("Restore accounts", "storage", a.storage)

	a.Clear()
	file, err := os.ReadFile(a.storage)
	if err != nil {
		a.logger.Error(err, "unable to restore accounts", "storage", a.storage)
		a.Default()
		return fmt.Errorf("unable to restore accounts from (%s), using defaults, error: %v", a.storage, err)
	}

	err = json.Unmarshal([]byte(file), &a.accounts)
	if err != nil {
		a.logger.Error(err, "unable to unmarshal accounts json file")
		return fmt.Errorf("unable to unmarshal accounts json file, error: %v", err)
	}

	return nil
}

// GetAccountUsernames: return a slice of active account usernames
func GetAccountUsernames() []string {
	usernames := []string{}

	for _, account := range a.accounts.Accounts {
		if account.Username != "" {
			usernames = append(usernames, account.Username)
		}
	}

	return usernames
}

// GetRoles: return a slice of all available roles
func GetRoles() []string {
	return RoleTypes
}
