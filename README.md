Reports view for access control system NVP Bolid "Orion Pro"
======
# Installing

Set environment variable **BOLID_DSN**:
```bash
export BOLID_DSN="server=127.0.0.1;user id=username;password=passwd;database=base"
```
Install Gorion:
```bash
go install github.com/jezman/gorion
```

# Features

- [List of](#lists)
  * [Company](#company-list)
  * [Doors](#doors-list)
  * [Employees](#employees-list)
- [Get events](#events)
- [Calculate employee worked time](#employees-worked-time)

```
~ gorion --help
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
# Lists
```
~ gorion list -h
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
~ gorion list company
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
~ gorion list doors
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
~ gorion list employees
+-----+---------------+----------+
| #   | Company       | Employee |
+-----+---------------+----------+
| 1   | Sergey Brin   | Google   |
| 2   | Steve Wozniak | Apple    |
...
```
## Events
```
~ gorion events -h
Displays a list of events depending on entered flags

Usage:
  gorion events [flags]

Flags:
  -d, --door uint         door ID. Use: 'gorion list doors' to get a list of all doors with ID.
  -e, --employee string   employee last name. Use: 'gorion list employees' to get a list of all employees.
  -f, --first string      first date (default "12.12.2017")
  -h, --help              help for events
  -l, --last string       last date. (default "13.12.2017")
```
## Employees worked time
```
~ gorion hours -h
Displays employees worked time

Usage:
  gorion hours [flags]

Flags:
  -d, --door uint         door ID. Use: 'gorion list doors' to get a list of all doors with ID.
  -e, --employee string   employee last name. Use: 'gorion list employees' to get a list of all employees.
  -f, --first string      first date (default "12.12.2017")
  -h, --help              help for hours
  -l, --last string       last date. (default "13.12.2017")
```