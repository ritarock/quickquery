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

USE:
  SELECT, FROM, WHERE AND

DON'T USE:
  OR, OREDER BY
```

## Sample
```
$ qq "select * from ./sample.csv where id >= 2 AND name = name3"
id, name, user
3, name3, user3
```
