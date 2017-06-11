# go-hello-world-plus

Build `go-hello-world-plus-M.m.P-I.x86_64.rpm`
and   `go-hello-world-plus_M.m.P-I_amd64.deb`
where "M.m.P-I" is Major.minor.Patch-Iteration.

## Usage

An extension of 
[go-hello-world](https://github.com/docktermj/go-hello-world)
that adds the following features:

1. Commandline help
1. Logging
1. Sub-command structure

### Invocation

```console
go-hello-world-plus
```

## Development

### Dependencies

#### Set environment variables

```console
export GOPATH="${HOME}/gocode"
export PATH="${PATH}:/usr/local/go/bin:${GOPATH}/bin"
export PROJECT_DIR=${GOPATH}/src/github.com/docktermj
```

#### Download project

```console
mkdir -p ${PROJECT_DIR}
cd ${PROJECT_DIR}
git clone git@github.com:docktermj/go-hello-world-plus.git
```

#### Download dependencies

```console
go get github.com/docopt/docopt-go
```

### Build

#### Local build

```console
go install github.com/docktermj/go-hello-world-plus
```

The results will be in the `${GOPATH}/bin` directory.

#### Docker build

```console
cd ${PROJECT_DIR}/go-hello-world-plus
sudo make build
```

The results will be in the `.../target` directory.

### Install

#### RPM-based

Example distributions: openSUSE, Fedora, CentOS, Mandrake

##### RPM Install

Example:

```console
sudo rpm -ivh go-hello-world-plus-M.m.P-I.x86_64.rpm
```

##### RPM Update

Example: 

```console
sudo rpm -Uvh go-hello-world-plus-M.m.P-I.x86_64.rpm
```

#### Debian

Example distributions: Ubuntu

##### Debian Install / Update

Example:

```console
sudo dpkg -i go-hello-world-plus_M.m.P-I_amd64.deb
```

### Cleanup

```console
sudo make clean
```
