# quickquery

## install
```
$ git clone https://github.com/ritarock/quickquery.git
$ cd quickquery
$ make install
```

## Usage
```
$ qq -h
quickquery can search from csv like sql

Usage:
  quickquery "select * from ./filepath"

OK:
  SELECT, FROM, WHERE, AND, ORDER BY

NG:
  OR, LIMIT, GROUP BY
```

## Sample
```
$ qq "select * from ./sample.csv where id >= 2 and id <= 3 order by id desc"
id, name, user
3, name3, user3
2, name2, user2
```
