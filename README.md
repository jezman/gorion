# Reports for Orion Pro(MS SQL Database).

Install and edit config file.
```baah
$ go install github.com/jezman/orion./orion
$ vim $GOPATH/src/github.com/jezman/orion/config.json
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
Usage of orion:
  -door string
    	List all doors with indexes.
  -end string
    	End of time range. Format: "31.12.2017 23:59" (default "21.03.2017 23:59")
  -list
    	List all doors.
  -name string
    	User last name.
  -start string
    	Start of time rande. Format: "31.12.2017 23:59" (default "21.03.2017 00:00")
```
