.PHONY: test coverage covreport

# semantic-release-dry-run:
# 	@echo "GH_TOKEN is $$GH_TOKEN"
# 	@npx semantic-release --dry-run

test:
	@go clean -testcache
	@go test -v ./... -cover

coverage:
	@go clean -testcache
	@go test -v ./... -coverprofile=coverage.txt

covreport:
	go tool cover -html coverage.txt -o index.html