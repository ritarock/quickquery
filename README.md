# quickquery

## install
```
$ git clone https://github.com/ritarock/quickquery.git
$ cd quickquery
$ make install
```

## Usage
```
$ quickquery -h
quickquery can search from csv like sql

Usage:
  quickquery [flags]

Flags:
  -h, --help   help for quickquery
```

## Example
```
$ cat sample.csv
id,team_id,name,note
1,1,name1,note1
2,1,name2,note2
3,2,name3,note3
4,3,name4,note4
5,4,name5,note5
6,1,name6,note5
7,2,name7,note6

$ quickquery "select * from sample.csv"
| id | team_id | name  | note  |
+----+---------+-------+-------+
| 1  | 1       | name1 | note1 |
| 2  | 1       | name2 | note2 |
| 3  | 2       | name3 | note3 |
| 4  | 3       | name4 | note4 |
| 5  | 4       | name5 | note5 |
| 6  | 1       | name6 | note5 |
| 7  | 2       | name7 | note6 |
(7 rows)
```
