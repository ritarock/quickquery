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
$ cat sample.csv 
id, name, user
1, name1, user1
2, name2, user2
3, name3, user3
4, name4, user4
5, name5, user5

$ qq "select * from ./sample.csv where id >= 2 and id <= 3 order by id desc"
id, name, user
3, name3, user3
2, name2, user2

$ qq "select user, name from ./sample.csv where id >= 2 and id <= 3 order by id"
user, name
user2, name2
user3, name3
```
