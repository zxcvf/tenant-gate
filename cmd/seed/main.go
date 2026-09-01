package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"tenant-gate/pkg/postgres"
	"time"

	"github.com/bwmarrin/snowflake"

	"golang.org/x/crypto/bcrypt"
)

var (
	// Sample tenant and user data
	sampleTenantName  = "Sample Tenant"
	sampleTenantEmail = "sample@example.com"
	sampleUserName    = "Sample User"
	sampleUserEmail   = "sample@example.com"
	sampleUserPhone   = "+811234567890"
	samplePassword    = "password"
	_tenanetSql       = `INSERT INTO tenants (id, tenant_name, email, created_by, created_at, updated_at) VALUES ('%s', '%s', '%s', 'system', NOW(), NOW()) ON CONFLICT (id) DO NOTHING;`
	_userSql          = `INSERT INTO users (id, username, email, phone, password_hash, created_at, updated_at) VALUES ('%s', '%s', '%s', '%s', '%s', NOW(), NOW()) ON CONFLICT (id) DO NOTHING;`
	_tenantUserSql    = `INSERT INTO tenants_users (id, tenant_id, user_id, role_code) VALUES (1, '%s', '%s', 1) ON CONFLICT (id) DO NOTHING;`
	bcryptCost        = bcrypt.DefaultCost
)

func main() {
	// is a function that generates a sample tenant for testing purposes.
	// Initialize the database connection
	databaseURL, ok := os.LookupEnv("POSTGRESQL_URL")
	log.Printf("Migrate: database url: %s", databaseURL)
	if !ok || len(databaseURL) == 0 {
		log.Fatalf("migrate: environment variable not declared: POSTGRESQL_URL")
	}

	pg, err := postgres.New(databaseURL,
		postgres.MaxPoolSize(1),
	)
	if err != nil {
		log.Fatalf("app - Run - postgres.New: %v", err)
	}
	defer pg.Pool.Close()

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	node, _ := snowflake.NewNode(1)
	tenant_id := node.Generate()
	user_id := node.Generate()

	password_hash, err := bcrypt.GenerateFromPassword([]byte("password"), bcryptCost)
	if err != nil {
		log.Fatalf("Failed to generate password hash: %v", err)
	}

	log.Println("Generated sample tenant ID:", tenant_id)
	_, err = pg.Pool.Exec(ctx, fmt.Sprintf(_tenanetSql, tenant_id, sampleTenantName, sampleTenantEmail))
	if err != nil {
		log.Fatalf("Failed to generate sample tenant: %v", err)
	}

	log.Println("Generated sample user ID:", user_id)
	_, err = pg.Pool.Exec(ctx, fmt.Sprintf(_userSql, user_id, sampleUserName, sampleUserEmail, sampleUserPhone, password_hash))
	if err != nil {
		log.Fatalf("Failed to generate sample user: %v", err)
	}

	log.Println("Generated tenant user relation")
	_, err = pg.Pool.Exec(ctx, fmt.Sprintf(_tenantUserSql, tenant_id, user_id))
	if err != nil {
		log.Fatalf("Failed to generate sample tenant user relation: %v", err)
	}

	log.Println("Sample tenant and user generated successfully.")
}
