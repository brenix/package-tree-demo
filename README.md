# package-tree

## Table of Contents

- [package-tree](#package-tree)
  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
    - [Messages](#messages)
    - [Commands](#commands)
    - [Examples](#examples)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Testing](#testing)
    - [Building](#building)
      - [Static binary](#static-binary)
      - [Docker](#docker)
  - [Usage](#usage)
    - [Starting the server](#starting-the-server)
    - [Client example](#client-example)
  - [Design Rationale](#design-rationale)
    - [Goals](#goals)
    - [Choices made](#choices-made)
  - [Benchmarks](#benchmarks)
  - [References](#references)

## Overview

`package-tree` provides a simple server to track packages and their
dependencies. It's primary function is to start a TCP listener which handles
incoming requests to add, query, or remove a package and its dependencies from
an index. Each incoming request is run concurrently and is stored in the index
in a consistent manner.

### Messages

Messages sent by clients should adhere to the following format:

```
<command>|<package>|<dependencies>
```

**Where:**

- `<command>` (**mandatory**) - Describes what action to take
- `<package>` (**mandatory**) - Name of the package
- `<dependencies>` (*optional*) - Comma-delimited list of packages that need to
  be present before `<package>` is installed
- The message should end with a `\n` character

### Commands

The following commands are available:

| Command  | Function                                          |
| -------- | ------------------------------------------------- |
| `INDEX`  | Adds a package and it's dependencies to the index |
| `REMOVE` | Removes a package from the index                  |
| `QUERY`  | Validates if a package is currently indexed       |

For `INDEX` commands:

- The server returns `OK\n` if the package can be indexed
- The server returns `OK\n` if the package dependencies are updated
- The server returns `FAIL\n` if the package cannot be indexed because some of
  it's dependencies aren't indexed yet and need to be installed first

For `REMOVE` commands:

- The server returns `OK\n` if the package can be removed from the index
- The server returns `FAIL\n` if the package could not be removed from the index
  because another package depends on it
- The server returns `OK\n` if the package is not currently indexed

For `QUERY` commands:

- The server returns `OK\n` if the package is currently indexed
- The server returns `FAIL\n` if the package is not currently indexed

If a command is not recognized, or if there was a problem parsing the incoming
message, the server returns `ERROR\n`

### Examples

Here are some example messages:

```
INDEX|cloog|gmp,isl,pkg-config\n
INDEX|ceylon|\n
REMOVE|cloog|\n
QUERY|cloog|\n
```

## Getting Started

### Prerequisites

To build and run the server, you will need to have `go` installed. See
[this](https://golang.org/doc/install) for installation and setup instructions

### Testing

To run unit tests and compute test coverage, run the following:

```bash
go test -cover
```

### Building

#### Static binary

To build a static binary (preferred), run the following:

```bash
go build -a -ldflags '-w -extldflags "-static"'
```

> By default, the output binary will be `package-tree-demo` unless `-o
> package-tree` is specified

#### Docker

A `Dockerfile` is also included, which can be used to build a container image.

To build a container using Docker, run the following:

```bash
docker build -t package-tree .
```

## Usage

```bash
$ ./package-tree -h
Usage of ./package-tree:
  -port int
        Port to listen for incoming connections (default 8080)
```

### Starting the server

To start the server listening the default port (8080), run the following:

```bash
./package-tree
```

> Upon successful start, the server should print the message: `Listening on
> 0.0.0.0:8080`

To start the server listening an alternative port, pass the `-port` argument
like so:

```bash
./package-tree -port 3337
```

### Client example

To test commands, you can use a tool such as `nc`.

Example `INDEX` command:

```bash
echo "INDEX|foo|\n" | nc -q0 localhost 8080
```

Example `QUERY` command:

```bash
echo "QUERY|foo|\n" | nc -q0 localhost 8080
```

Example `REMOVE` command:

```bash
echo "REMOVE|foo|\n" | nc -q0 localhost 8080
```

## Design Rationale

### Goals

- **Fast and efficient** - The server must be able to handle many requests from
  multiple systems concurrently without locking up
- **Simple** - The code should be simple and easy to understand as well as reach
  a common audience
- **Consistent** - The data should remain consistent at all times, otherwise you
  can run into conflicts or security issues
- **Portable** - The server should be easy to build and run - who likes
  headaches anyways?

### Choices made

- **Language: Go** - Go satisfies several of the goals. It's built-in support
  for concurrency allows for scaling up to thousands of requests per second.
  Building a binary is simple which results in a portable application.
- **Concurrency: goroutines** - This was an obvious choice if you are familiar
  with Go. I chose not to use channels as I am not too familiar with them. While
  they look to be useful, they can also make the code confusing quickly.
- **Indexing: maps and sync.Mutex** - Using a map combined with `sync.Mutext`
  was the most performant over `sync.Map`. This allows the data to be indexed
  consistently and avoids race conditions.


## Benchmarks

The following are average benchmarks taken using a test-suite which adds,
removes, and queries 3000+ packages and dependencies:

 **12-Core CPU**:

| Concurrency | Elapsed Time |
| ----------- | ------------ |
| 1           | 5121ms       |
| 2           | 2811ms       |
| 4           | 1806ms       |
| 10          | 1778ms       |

**2-core CPU**:
| Concurrency | Elapsed Time |
| ----------- | ------------ |
| 1           | 13909ms      |
| 2           | 7138ms       |
| 4           | 4863ms       |
| 10          | 6875ms       |

## References

These following links provided useful information in building this tool (much thanks!):

- https://opensource.com/article/18/5/building-concurrent-tcp-server-go
- https://gobyexample.com/mutexes
- https://yourbasic.org/golang/regexp-cheat-sheet/
- https://blog.golang.org/go-maps-in-action
- https://github.com/cweill/gotests