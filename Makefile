VERSION := v0.0.1

test:
	@ go vet ./...
	@ go run honnef.co/go/tools/cmd/staticcheck@latest ./...
	@ go test -race -count=10 -failfast ./...

release: test
	@ go mod tidy
	@ test -z "`git status --porcelain | grep -vE 'History\.md'`" || (echo "uncommitted changes detected." && false)
	@ git commit -am "Release v$(VERSION)"
	@ git tag "v$(VERSION)"
	@ git push origin main "v$(VERSION)"
	@ go run github.com/cli/cli/v2/cmd/gh@5023b61 release create --generate-notes "v$(BUD_VERSION)"
