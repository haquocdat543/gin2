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
	migrate -path src/db/migration -database "postgres://develop:effimatebackend@localhost:5432/postgres?sslmode=disable" version
mf:
	migrate -path src/db/migration -database "postgres://develop:effimatebackend@localhost:5432/postgres?sslmode=disable" force 1
mup:
	migrate -path src/db/migration -database "postgres://develop:effimatebackend@localhost:5432/postgres?sslmode=disable" up 1
mdown:
	migrate -path src/db/migration -database "postgres://develop:effimatebackend@localhost:5432/postgres?sslmode=disable" down 1

# Seeding
seed-start:
	go run main.go seed start

# KeyGenerate
private:
	openssl genpkey -algorithm RSA -out private.pem -pkeyopt rsa_keygen_bits:2048
public:
	openssl rsa -in private.pem -pubout -out public.pem
