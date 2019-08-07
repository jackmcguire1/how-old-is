clean:
	@echo Cleaning dist folder
	@rm -rf dist
	@mkdir -p dist

build: clean
	@echo running lint
	- golangci-lint run
	@echo building lambda handlers
	@for dir in `ls lambda/handlers/`; do \
		GOARCH=amd64 GOOS=linux go build -gcflags='-N -l' -o dist/$$dir/main lambda/handlers/$$dir/main.go; \
	done

validate:
	sam validate

package: validate build
	sam package \
		--template-file template.yml \
		--s3-bucket {{bucket-name}} \
		--output-template-file packaged.yml \
		--region us-east-1 \

deploy:
	sam deploy \
	--template-file ./packaged.yml \
	--stack-name how-old-is \
	--capabilities CAPABILITY_IAM \
	--debug

describe-staging:
		@aws cloudformation describe-stacks \
			--region us-east-1 \
			--stack-name how-old-is \
