# quickquery

## install
```
$ git clone https://github.com/ritarock/quickquery.git
$ cd quickquery
$ make install
```

## Usage
```
NAME:
   quickquery - SQL-like query for csv

USAGE:
   [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

## Sample
### SELECT
```
$ quickquery "select * from sample.csv"
1|user1|name2|
2|user2|name2|
```

### INSERT
```
$ quickquery "insert into sample.csv values (3,user3,name3)"
$ cat sample.csv
id,user,name
1,user1,name2
2,user2,name2
3,user3,name3

$ quickquery "insert into sample.csv (id,user) values (4,user3)"
~/dev/quickquery on  main [MU] ❯ cat sample.csv
id,user,name
1,user1,name2
2,user2,name2
3,user3,name3
4,user3,""
```

### UPDATE
```
$ quickquery "update sample.csv set user = user3rd where id = 3"
$ cat sample.csv
id,user,name
1,user1,name2
2,user2,name2
3,user3rd,name3
4,user3,""


$ quickquery "update sample.csv set user = user3rd,name = name4   where id = 4"
$ cat sample.csv
id,user,name
1,user1,name2
2,user2,name2
3,user3rd,name3
4,user3rd,name4
```

### DELETE
```
$ quickquery "delete from sample.csv where user = user3rd"
$ cat sample.csv
id,user,name
1,user1,name2
2,user2,name2
```
