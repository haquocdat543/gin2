DB_USER=$(shell cat .env | grep DB_USER | cut -d= -f 2 | sed "s|\"||g")
DB_PASSWORD=$(shell cat .env | grep DB_PASSWORD | cut -d= -f 2 | sed "s|\"||g")
DB_HOST=$(shell cat .env | grep DB_HOST | cut -d= -f 2 | sed "s|\"||g")
DB_PORT=$(shell cat .env | grep DB_PORT | cut -d= -f 2 | sed "s|\"||g")

DB_URL="postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/postgres?sslmode=disable"

i-ginkgo:
	go install github.com/onsi/ginkgo/v2/ginkgo@latest
	go get github.com/onsi/gomega
i-air:
	go install github.com/air-verse/air@latest
i-reflex:
	go install github.com/cespare/reflex@latest

air:
	air
reflex:
	reflex -r '\.go$$' -s -- go run main.go

ginkgo-b:
	ginkgo bootstrap
ginkgo-w:
	ginkgo watch

# Migration
m-gen:
	migrate create -ext sql -dir src/db/migration -seq create_users_table
mv:
	migrate -path pkg/db/migration -database $(DB_URL) version
mf:
	migrate -path pkg/db/migration -database $(DB_URL) force 1
mup:
	migrate -path pkg/db/migration -database $(DB_URL) up 1
mdown:
	migrate -path pkg/db/migration -database $(DB_URL) down 1

# Seeding
seed-start:
	go run main.go seed start

# KeyGenerate
private:
	openssl genpkey -algorithm RSA -out private.pem -pkeyopt rsa_keygen_bits:2048
public:
	openssl rsa -in private.pem -pubout -out public.pem
update-env-key: private public
	sed -i '' -E "s|^JWT_PRIVATE_KEY_BASE64_ENCODED=[^ ]*|JWT_PRIVATE_KEY_BASE64_ENCODED=$$(cat private.pem | base64)|" .env
	sed -i '' -E "s|^JWT_PUBLIC_KEY_BASE64_ENCODED=[^ ]*|JWT_PUBLIC_KEY_BASE64_ENCODED=$$(cat public.pem | base64)|" .env
	rm private.pem
	rm public.pem

