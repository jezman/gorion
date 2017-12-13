[![Build Status](https://travis-ci.org/jezman/gorion.svg?branch=master)](https://travis-ci.org/jezman/gorion)
[![Go Report Card](https://goreportcard.com/badge/github.com/jezman/gorion)](https://goreportcard.com/report/github.com/jezman/gorion)

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
go get github.com/jezman/gorion &&\
  cd $GOPATH/src/github.com/jezman/gorion &&\
  go install
```
Otherwise, please see [Go install](https://golang.org/doc/install).
## Features

- [List of](#lists)
  * [Company](#company-list)
  * [Doors](#doors-list)
  * [Employees](#employees-list)
- [Get events](#events)
- [Calculate employee worked time](#employees-worked-time)

```
$ gorion --help
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
## Lists
```
$ gorion list -h
Get list of company, doors, employees

Usage:
  gorion list [command]

Available Commands:
  company     Displays a list of companies
  doors       List all doors with ID
  employees   Displays a list of employees

Flags:
  -h, --help   help for list

Use "gorion list [command] --help" for more information about a command.
```
## Company list
```
$ gorion list company
+----+---------+
| #  | Company |
+----+---------+
| 1  | Yandex  |
| 2  | Google  |
| 3  | Nestle  |
... 
| 32 | Apple   |
| 33 | Nissan  |
```
## Doors list
```
$ gorion list doors
+-----+----------------+
| ID  | Door           |
+-----+----------------+
| 1   | Main entrance  |
| 2   | Emergency exit |
| 3   | Service room   |
...
| 66   | Bar           |
| 67   | Rest room     |
```
## Employees list
```
$ gorion list employees
+-----+---------------+----------+
| #   | Company       | Employee |
+-----+---------------+----------+
| 1   | Sergey Brin   | Google   |
| 2   | Steve Wozniak | Apple    |
...
```
## Events
```
$ gorion events -h
Displays a list of events depending on entered flags

Usage:
  gorion events [flags]

Aliases:
  events, e

Examples:
  gorion events
  gorion events --employee=lastname --first=05.08.2017
  gorion e -e lastname -d 32
  gorion e -d 2 -f 12.11.2017 -l 16.11.2107

Flags:
  -d, --door uint         door ID. Use: 'gorion list doors' to get a list of all doors with ID.
  -e, --employee string   employee last name. Use: 'gorion list employees' to get a list of all employees.
  -f, --first string      first date (default "12.12.2017")
  -h, --help              help for events
  -l, --last string       last date. (default "13.12.2017")
```
## Employees worked time
```
$ gorion hours -h
Displays employees worked time

Usage:
  gorion hours [flags]

Aliases:
  hours, h

Examples:
  gorion hours
  gorion hours --employee=lastname --first=05.08.2017 --last=07.08.2017
  gorion h -e lastname
  gorion h -f 12.11.2017 -l 16.11.2107

Flags:
  -e, --employee string   employee last name. Use: 'gorion list employees' to get a list of all employees.
  -f, --first string      first date (default "12.12.2017")
  -h, --help              help for hours
  -l, --last string       last date. (default "13.12.2017")
```
