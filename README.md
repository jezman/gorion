# Gorion

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
  - workers
    - all
    - by company name
  - doors
  - events types
- Events
  - all
  - only denied
  - by door
  - by worker last name
  - date range
- Workers worked time
  - all
  - by company
  - by worker last name
  - date range
- Workers
  - add
  - delete
  - disable card
  - enable card

### For more info run `gorion` with `--help` flag

```text
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
  add         add worker to access control system
  delete      delete workers from access control system
  disable     disable worker card
  enable      enable worker card
  events      Displays a list of events depending on entered flags
  help        Help about any command
  hours       Displays workers worked time
  list        Get list of company, doors, workers
  version     show application version

Flags:
  -h, --help   help for gorion

Use "gorion [command] --help" for more information about a command.
```

## License

MIT © 2018 jezman
