# How-Old-Is
> This repository contains the Alexa skill to find out how long something or someone has been alive for.

[git]:    https://git-scm.com/
[golang]: https://golang.org/
[modules]: https://github.com/golang/go/wiki/Modules
[golint]: https://github.com/golangci/golangci-lint
[aws-cli]: https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html
[aws-cli-config]: https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html
[aws-sam-cli]: https://github.com/awslabs/aws-sam-cli

## SETUP
> How to configure your Alexa Skill

> Please Configure the Makefile with your own available S3 bucket

1. Create a new Alexa skill with a name of your choice 

2. Set Alexa skill invocation method as 'how old is x'

3. Set built-in invent invocations to their relevant phrases i.e. 'help', 'stop', 'cancel', etc.

4. Set a random phrase for the built-in fallback intent, i.e. 'bumbaclart'

5. Create new Intent named 'AgeIntent'

6. Create new intent slot named 'name', with <b>SLOT TYPE</b> 'AMAZON.FirstName'

7. Add invocation phrase for the 'AgeIntent' intent with phrase 'how old is {name}'

8. Configure slot values for the 'AMAZON.FirstName' <b>SLOT TYPE</b> i.e. 'Bob'

9. Edit the 'internal/dom/age.go' file to feature your name and DoB in RFC3339 format

10. Package and deploy how-old-is lambda

11. Configure Alexa skill endpoint lambda ARN:<br>
Once the <b>'how-old-is'</b> lambda has been deployed, <br>
retrieve the generated lambda ARN using the AWS console or<br>
one of the describe stack methods found above.<br>
input the lambda <b>ARN</b> as the default endpoint of your Alexa skill,<br>
within your Alexa development console!

12. Begin testing your Alexa skill by querying for 'how old is x'

13. Query Alexa 'how old is Bob'

14. Query Alexa 'bumbaclart' or your fallback invocation phrase!

15. Tell Alexa to 'stop'

16. <b>Testing complete!</b>

## Development

To develop `how-old-is` or interact with its source code in any meaningful way, be
sure you have the following installed:

### Prerequisites

- [Git][git]
- [Go 1.12][golang]+
- [golangCI-Lint][golint]
- [AWS CLI][aws-cli]
- [AWS SAM CLI][aws-sam-cli]

>You will need to activate [Modules][modules] for your version of [GO][golang], 

> by setting the `GO111MODULE=on` environment variable set

### [golangCI-Lint][golint]
```shell
curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $GOPATH/bin latest
```

### [AWS CLI Configuration][aws-cli-config]
> Make sure you configure the AWS CLI
- AWS Access Key ID
- AWS Secret Access Key
- Default region 'us-east-1'
```shell
aws configure
```


## Build
> How to build the Alexa skill lambda

> **Note***:- Executables are placed in 'dist/handlers/*/main' 

```shell
make build
```

**OR**

- **Lint the project FIRST**
> **Note***:-This may throw build **errors**, which **may** not break anything!

```Shell 
golangci-lint run
```

- **Build the lambda**
```Shell 
GOARCH=amd64 GOOS=linux go build -gcflags='-N -l' -o dist/how-old-is/main  ./lambda/handlers/how-old-is/main.go
```


## Package & Deploy
> This section describes how to validate, compile and deploy the latest Alexa skill.

### Validate
> Validate the template.yml before packaging!

```shell
make validate
```

**OR**

```shell
sam validate
```

### Package
> How to compile template.yml after updating the back-end stack definition.

```shell
make package
```
- runs clean
- runs build
- runs sam validate
- runs sam package

**OR**

```shell 
sam package --template-file template.yml --s3-bucket {{bucket-name}} --output-template-file packaged.yml
```

**OR**

```shell 
aws cloudformation package --template-file template.yml --output-template-file packaged.yml --s3-bucket {{bucket-name}}
```

### Deploy
> How to deploy the complied packaged.yml to update the serverless stack.

```shell
make deploy
```

**OR**

```shell
sam deploy --template-file ./packaged.yaml --stack-name how-old-is --capabilities CAPABILITY_IAM
```

**OR**

```shell 
aws cloudformation deploy --template-file ./packaged.yml --stack-name how-old-is --capabilities CAPABILITY_IAM
```

### Describe Stack
> Get information about the back-end infrastructure

```shell
make describe
```

**OR**

```shell 
aws cloudformation describe-stacks  --region us-east-1 --stack-name how-old-is
```

## Contributors

This project exists thanks to **all** the people who contribute.

## Donations
All donations are appreciated!

[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](http://paypal.me/crazyjack12)