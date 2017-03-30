# Отчеты для НВП Болид "Орион ПРО"(MS SQL Database).

Сперва устанавливаем [Go](https://golang.org/).
Далее:
```bash
$ go get github.com/jezman/gorion
$ cd $GOPATH/src/github.com/jezman/gorion
$ vim comfig.json
```
```json
{
    "server": "127.0.0.1",
    "database": "db",
    "user": "user",
    "password": "passwd"
}
```
```bash
$ go build
$ ./gorion -h

ИМЯ:
   gorion - создает отчеты для системы контроля доступом НВП Болид "Орион ПРО"
ИСПОЛЬЗОВАНИЕ:
   gorion [глобальные параметры] команда

КОМАНДЫ:
   hours, h       приход и уход сотрудников + их отработанное время
   listdoors, ld  список всех дверей с id
   summary, s     общий отчет
   help, h        Shows a list of commands or help for one command

ГЛОБАЛЬНЫЕ ПАРАМЕТРЫ:
   --door value, -d value      id двери. для просмотра списка всех дверей введите: gorion ld
   --employee value, -e value  фамилия сотрудника
   --first value, -f value     первая дата (default: "30.03.2017")
   --last value, -l value      последняя дата (default: "31.03.2017")
   --help, -h                  show help
```
