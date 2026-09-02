package jwt

import (
	"errors"
	"fmt"
	"time"

	jwtlib "github.com/golang-jwt/jwt/v5"
)

var (
	// ErrUnexpectedSigningMethod is returned when the JWT signing method is not expected.
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")

	_defaultDuration      = 24 * time.Hour
	_defaultSigningMethod = jwtlib.SigningMethodHS256
)

// Manager handles JWT token generation and parsing.
type Manager struct {
	secret   string
	duration time.Duration
}

// New -.
func New(secret string, duration time.Duration) *Manager {
	if duration == 0 {
		duration = _defaultDuration
	}

	return &Manager{
		secret:   secret,
		duration: duration,
	}
}

// GenerateToken creates a new JWT token for the given user ID.
func (m *Manager) GenerateToken(c *UserClaims) (string, error) {
	c.ExpiresAt = jwtlib.NewNumericDate(time.Now().Add(m.duration))
	token := jwtlib.NewWithClaims(_defaultSigningMethod, c)
	tokenString, err := token.SignedString([]byte(m.secret))
	if err != nil {
		return "", fmt.Errorf("jwt - GenerateToken - token.SignedString: %w", err)
	}

	return tokenString, nil
}

// ParseToken validates a JWT token and returns the user ID.
func (m *Manager) ParseToken(tokenString string) (*UserClaims, error) {
	u := &UserClaims{}
	token, err := jwtlib.ParseWithClaims(tokenString, u, func(token *jwtlib.Token) (any, error) {
		if _, ok := token.Method.(*jwtlib.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%w: %v", ErrUnexpectedSigningMethod, token.Header["alg"])
		}
		return []byte(m.secret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("jwt - ParseToken - jwtlib.Parse: %w", err)
	}

	_, err = token.Claims.GetSubject()
	if err != nil {
		return nil, fmt.Errorf("jwt - ParseToken - GetSubject: %w", err)
	}

	return u, nil
}
