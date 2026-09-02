package jwt_test

import (
	"tenant-gate/pkg/jwt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJWT_GenerateAndParse(t *testing.T) {
	t.Parallel()

	j := jwt.New("test-secret", time.Hour)

	c := jwt.NewUserClaims(
		12983719827, "user@example.com", "John Doe", 1, "TenantName", 1,
	)

	token, err := j.GenerateToken(c)
	require.NoError(t, err)
	assert.NotEmpty(t, token)

	c2, err := j.ParseToken(token)

	require.NoError(t, err)
	assert.Equal(t, *c, *c2)
}

func TestJWT_ParseToken_Invalid(t *testing.T) {
	t.Parallel()

	j := jwt.New("test-secret", time.Hour)

	_, err := j.ParseToken("invalid-token")
	require.Error(t, err)
}

func TestJWT_ParseToken_WrongSecret(t *testing.T) {
	t.Parallel()

	j1 := jwt.New("secret-1", time.Hour)
	j2 := jwt.New("secret-2", time.Hour)

	token, err := j1.GenerateToken(jwt.NewUserClaims(
		12983719827, "user@example.com", "John Doe", 1, "TenantName", 1,
	))
	require.NoError(t, err)

	_, err = j2.ParseToken(token)
	require.Error(t, err)
}

func TestJWT_ParseToken_Expired(t *testing.T) {
	t.Parallel()

	j := jwt.New("test-secret", -time.Hour)

	token, err := j.GenerateToken(jwt.NewUserClaims(
		12983719827, "user@example.com", "John Doe", 1, "TenantName", 1,
	))
	require.NoError(t, err)

	_, err = j.ParseToken(token)
	require.Error(t, err)
}
