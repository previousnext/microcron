MicroCron
=========

[![CircleCI](https://circleci.com/gh/previousnext/microcron.svg?style=svg)](https://circleci.com/gh/previousnext/microcron)

**Maintainer**: Nick Schuch

A cron for running repetitive tasks in the foreground of a container.

```
cron-drush:
  image: busybox
  entrypoint: ['microcron', '60s', 'echo 1']
```

## Development

### Getting started

For steps on getting started with Go:

https://golang.org/doc/install

To get a checkout of the project run the following commands:

```bash
# Make sure the parent directories exist.
mkdir -p $GOPATH/src/github.com/previousnext

# Checkout the codebase.
git clone git@github.com:previousnext/microcron $GOPATH/src/github.com/previousnext/microcron

# Change into the project to run workflow commands.
cd $GOPATH/src/github.com/previousnext/microcron
```

### Documentation

See `/docs`

### Resources

* [Dave Cheney - Reproducible Builds](https://www.youtube.com/watch?v=c3dW80eO88I)
* [Bryan Cantril - Debugging under fire](https://www.youtube.com/watch?v=30jNsCVLpAE&t=2675s)
* [Sam Boyer - The New Era of Go Package Management](https://www.youtube.com/watch?v=5LtMb090AZI)
* [Kelsey Hightower - From development to production](https://www.youtube.com/watch?v=XL9CQobFB8I&t=787s)

### Tools

```bash
# Dependency management
go get -u github.com/golang/dep/cmd/dep

# Testing
go get -u github.com/golang/lint/golint

# Release management.
go get -u github.com/tcnksm/ghr

# Build
go get -u github.com/mitchellh/gox
```

### Workflow

**Testing**

```bash
make lint test
```

**Building**

```bash
make build
```

**Releasing**

```bash
make release
```
