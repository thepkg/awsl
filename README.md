awsl (AWS Lambda)
-

[![CircleCI](https://circleci.com/gh/thepkg/awsl.svg?style=svg)](https://circleci.com/gh/thepkg/awsl)
[![Build Status](https://travis-ci.org/thepkg/awsl.svg?branch=master)](https://travis-ci.org/thepkg/awsl)
[![Go Report Card](https://goreportcard.com/badge/github.com/thepkg/awsl)](https://goreportcard.com/report/github.com/thepkg/awsl)
[![Coverage Status](https://coveralls.io/repos/github/thepkg/awsl/badge.svg?branch=master)](https://coveralls.io/github/thepkg/awsl?branch=master)
[![GoDoc](https://godoc.org/github.com/thepkg/awsl?status.svg)](https://godoc.org/github.com/thepkg/awsl)
[![DependaBot](https://img.shields.io/badge/dependabot-enabled-informational)](https://github.com/thepkg/awsl)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat)](http://makeapullrequest.com)
<!-- [![Dependency Review](https://github.com/thepkg/awsl/actions/workflows/dependency-review.yml/badge.svg)](https://github.com/thepkg/awsl/actions/workflows/dependency-review.yml) -->

Package `awsl` provides helpful functions to work with AWS Lambda.

## Installation

````
go get -u github.com/thepkg/awsl
````

## Usage

* **FromDynamoDBMap** - converts data from AWS Lambda Event for DynamoDB to regular map
(from `map[string]events.DynamoDBAttributeValue` to `map[string]interface{}`):

````go
func main() {
	lambda.Start(Handler)
}

func Handler(event events.DynamoDBEvent) {
	for _, record := range event.Records {
		rec := awsl.FromDynamoDBMap(record.Change.NewImage)
		// rec contains map[string]interface{} so it's easier to work with data!
	}
}
````
