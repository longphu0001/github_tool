[![Build Status](https://travis-ci.org/YuPengZTE/github_tool.svg?branch=master)](https://travis-ci.org/YuPengZTE/github_tool)
[![codecov](https://codecov.io/gh/YuPengZTE/github_tool/branch/master/graph/badge.svg)](https://codecov.io/gh/YuPengZTE/github_tool)
[![codebeat badge](https://codebeat.co/badges/ee2f2d59-3faa-4f08-8cd8-23f0bc8d095d)](https://codebeat.co/projects/github-com-yupengzte-github_tool-master)
[![License MIT](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://github.com/YuPengZTE/github_tool/blob/master/LICENSE)

[Codecov][1] Go Example
=======================

This repository serves as an **example** on how to use [Codecov Global][4] for Go.

# Travis CI

Add to your `.travis.yml` file.
```yml
language: go

go:
  - 1.8.x
  - tip

before_install:
  - go get -t -v ./...

script:
  - chmod 777 *
  - ./go.test.sh

after_success:
  - bash <(curl -s https://codecov.io/bash) -t uuid-repo-token
```

> - All other CI you can simply run `bash <(curl -s https://codecov.io/bash)`.
> - `-race` is a suggestion, not required. Learn more at https://blog.golang.org/race-detector

## Private Repos
> Set `CODECOV_TOKEN` in your environment variables.

Add to your `.travis.yml` file.
```yml
after_success:
  - bash <(curl -s https://codecov.io/bash) -t uuid-repo-token
```
> Or you can set the environment variable `CODECOV_TOKEN=uuid-repo-token` and remove the `-t` flag

## Caveat: Multiple files
> If you see this `cannot use test profile flag with multiple packages` then use this shell template to execute your tests and store coverage output

```shell
#!/usr/bin/env bash

set -e
echo "" > coverage.txt

for d in $(go list ./... | grep -v vendor); do
    go test -race -coverprofile=profile.out -covermode=atomic $d
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done
```
> Note: use `-covermode=atomic` or `-covermode=count` to show how many times a statement was executed.

Then run this file as your test (ex. `./test.sh`)

> Reference http://stackoverflow.com/a/21142256/2055281

View source and learn more about [Codecov Global Uploader][4]

Need help? Contact support https://github.com/codecov/support

[1]: https://codecov.io/
[4]: https://github.com/codecov/codecov-bash
