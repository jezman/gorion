# Reports for Orion Pro(MS SQL Database).

Install and edit config file.
```baah
$ go get github.com/jezman/orion./orion
$ cd $GOPATH/src/github.com/jezman/orion
$ vim comfig.json
```
config.json
```json
{
    "server": "127.0.0.1",
    "database": "db",
    "user": "user",
    "password": "passwd"
}
```
```bash
$ orion -h

NAME:
   OrionCLI - generates a reports for Bolid access control system "Orion Pro"

USAGE:
   orion [global options] command

COMMANDS:
     generate, g    generate a report
     listdoors, ld  list all doors with indexes
     help, h        Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --door value, -d value   door index. For show all doors indexes use: orion ld
   --first value, -f value  first date of a report (default: "23.03.2017 00:00")
   --last value, -l value   last date of a report (default: "23.03.2017 23:59")
   --user value, -u value   user last name
   --help, -h               show help
```
