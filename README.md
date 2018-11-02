# Gorion

[![Build Status](https://travis-ci.org/jezman/gorion.svg?branch=master)](https://travis-ci.org/jezman/gorion) [![codecov](https://codecov.io/gh/jezman/gorion/branch/master/graph/badge.svg)](https://codecov.io/gh/jezman/gorion) [![Go Report Card](https://goreportcard.com/badge/github.com/jezman/gorion)](https://goreportcard.com/report/github.com/jezman/gorion)


Reports view for access control system NVP Bolid "Orion Pro"

## Installing

Set environment variable **BOLID_DSN**:

```bash
export BOLID_DSN="server=127.0.0.1;user id=username;password=passwd;database=base"
```

Install Gorion
If you have [Go](https://golang.org/) installed:

```bash
go get github.com/jezman/gorion && go install github.com/jezman/gorion
```

Otherwise, please see [Go install](https://golang.org/doc/install).

## Features

- Lists of
  - companies
    - all
    - by name
  - employees
    - all
    - by company name
  - doors
  - events types
- Events
  - all
  - only denied
  - by door
  - by employee last name
  - date range
- Employees worked time
  - all
  - by company
  - by employee last name
  - date range

### For more info run `gorion` with `--help` flag

```bash
$ gorion --help
 _____            _
|  __ \          (_)
| |  \/ ___  _ __ _  ___  _ __  
| | __ / _ \| '__| |/ _ \| '_ \
| |_\ \ (_) | |  | | (_) | | | |
 \____/\___/|_|  |_|\___/|_| |_|
https://github.com/jezman/gorion
Reports view for access control system NVP Bolid 'Orion Pro'

Usage:
  gorion [command]

Available Commands:
  events      Displays a list of events depending on entered flags
  help        Help about any command
  hours       Displays employees worked time
  list        Get list of company, doors, employees

Flags:
  -h, --help   help for gorion

Use "gorion [command] --help" for more information about a command.
```

## License

MIT © 2018 jezman
