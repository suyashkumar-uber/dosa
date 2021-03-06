# DOSA - Declarative Object Storage Abstraction

[![GoDoc][doc-img]][doc]
[![Coverage Status][cov-img]][cov]
[![Build Status][ci-img]][ci]

## Abstract

:warning: DOSA is _BETA_ software. It is not recommended for production use.
We will announce when it's ready.

[DOSA](https://github.com/uber-go/dosa/wiki) is a storage framework that
provides a _delcarative object storage abstraction_ for applications in Golang
and (eventually) Java. DOSA is designed to relieve common headaches developers
face while building stateful, database-dependent services.

If you'd like to start by writing a small DOSA-enabled program, check out
[the getting started guide](https://github.com/uber-go/dosa/wiki/Getting-Started-Guide).

## Overview

DOSA is a storage library that supports:

 * methods to store and retrieve go structs
 * struct annotations to describe queries against data
 * tools to create and/or migrate database schemas
 * implementations that serialize requests to remote stateless servers

## Annotations

This project is released under the [MIT License](LICENSE.txt).

[doc-img]: https://godoc.org/github.com/uber/dosa-go?status.svg
[doc]: https://godoc.org/github.com/uber/dosa-go
[ci-img]: https://travis-ci.com/uber-go/dosa.svg?token=zQquuxnrcfs8yizJ2Dcp&branch=master
[ci]: https://travis-ci.com/uber/dosa-go
[cov-img]: https://coveralls.io/repos/uber/dosa-go/badge.svg?branch=master&service=github
[cov]: https://coveralls.io/github/uber/dosa-go?branch=master
