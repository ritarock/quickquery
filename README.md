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

Supported:
  SELECT, FROM, WHERE, AND, ORDER BY, LIMIT

Unsupported:
  OR, IN, GROUP BY, etc...

Usage:
  quickquery [flags]

Flags:
  -h, --help   help for quickquery
```

## Example
```
$ cat ./sample.csv
id,team_id,name,note
1,1,name1,note1
2,1,name2,note2
3,2,name3,note3
4,3,name4,note4
5,4,name5,note5
6,1,name6,note5
7,2,name7,note6

$ qq "select * from ./sample.csv where team_id < 2"
id  team_id  name   note
--  -------  ----   ----
1   1        name1  note1
2   1        name2  note2
6   1        name6  note5

$ qq "select id, team_id, name from ./sample.csv where id <= 5 and team_id >= 2 order by id desc limit 2"
id  team_id  name
--  -------  ----
5   4        name5
4   3        name4
```
