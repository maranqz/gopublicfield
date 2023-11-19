# Go public field linter

[![CI](https://github.com/maranqz/gopublicfield/actions/workflows/ci.yml/badge.svg)](https://github.com/maranqz/gopublicfield/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/maranqz/gopublicfield)](https://goreportcard.com/report/github.com/maranqz/gopublicfield?dummy=unused)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
[![Coverage Status](https://coveralls.io/repos/github/maranqz/gopublicfield/badge.svg?branch=main)](https://coveralls.io/github/maranqz/gopublicfield?branch=main)

The linter blocks the changing of public fields. Unwritable fields help to avoid validation.
The linter is useful when:

* A project is migrated to Domain Model. <!-- TODO link to Fowler -->
* Business logic should not be broken by a direct variable assigning.
* You don't want to use snapshot pattern to protect business logic in new project.

## Use

Installation

    go install github.com/maranqz/gopublicfield/cmd/gopublicfield@latest

### Options

- `-p`, `--pkgs` - list of packages, where struct fields should be unwritable. By default, all fields in all external
  packages should be unwritable except local, [tests](testdata/src/publicfield/pgks).
    - `-op`, `onlyPkgs` - only packages `pkgs` should have unwritable public
      fields, [tests](testdata/src/publicfield/onlyPkgs).

## Example

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
package main

import (
	"bad"
)

func main() {
	u := &bad.User{}

	u.UpdatedAt = time.Now() // Field 'ID' in bad.User can be changes only inside nested package.`
}

```

```go
package bad

import "time"

type User struct {
	UpdatedAt time.Time
}
```

</td><td>

```go
package main

import (
	"good"
)

func main() {
	u := good.NewUser()

	u.Update()
}

```

```go
package user

import "time"

type User struct {
	UpdatedAt time.Time
}

func (u *User) Update() {
	u.UpdatedAt = time.Now()
}

```

</td></tr>
</tbody></table>

## TODO

### Feature, hardly implementable and not  planned

1. Type assertion, type declaration and type underlying, [tests](testdata/src/publicfield/default/type_nested.go.skip).
2. Unreadable fields.

### Problems, hardly fixable and not planed

1. Updating of slice, map items.
2. Updating by pointer to the unwritable field.
   ```go
    //..
    n := nested.Struct{}
    fieldPtr := &n.Int
    (*fieldPtr)++
    //..
   ```
 