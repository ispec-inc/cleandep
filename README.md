# cleandep

[![](https://pkg.go.dev/badge/github.com/ispec-inc/cleandep.svg)](https://pkg.go.dev/github.com/ispec-inc/cleandep/cmd/cleandep)

`cleandep` detects illegal dependencies.


## Installation

```shell
go install github.com/ispec-inc/cleandep/cmd/cleandep@latest
```

## Usage

### Local

```shell
cleandep ./...
```

### GitHub Actions

```yaml
name: cleandep

on:
  pull_request:

jobs:
  cleandep:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        dir: [go/foo, go/bar, go/baz]
    steps:
      - uses: actions/checkout@v3
      - run: go run github.com/ispec-inc/cleandep/cmd/cleandep@latest ./...
        working-directory: ${{ matrix.dir }}
```
