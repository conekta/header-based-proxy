PHONY: test
test:
	 go test ./... -coverprofile cover.out -tags musl

codeclimate:
	/bin/bash .codeclimate/codeclimate