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
