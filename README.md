# quickquery

## install
```
$ git clone https://github.com/ritarock/quickquery.git
$ cd quickquery
$ make install
```

## Usage
```
SQL-like query for csv

Usage:
  quickquery [flags]

Flags:
  -h, --help   help for quickquery
```

## Sample
### SELECT
```
$ quickquery 'select * from sample.csv'
1|user1|name2|
2|user2|name2|
```

### INSERT
```
$ quickquery 'insert into sample.csv values (3,user3,name3)'
$ cat sample.csv
id,user,name
1,user1,name2
2,user2,name2
3,user3,name3

$ quickquery 'insert into sample.csv (id,user) values (3,user4)'
$ cat sample.csv
id,user,name
1,user1,name2
2,user2,name2
3,user3,name3
3,user4,
```

### UPDATE
```
$ quickquery 'update sample.csv set user = user3rd where id = 3'
$ cat sample.csv
id,user,name
1,user1,name2
2,user2,name2
3,user3rd,name3
3,user3rd,""

$ quickquery 'update sample.csv set user = user_3rd,name = name_3rd where id = 3'
$ cat sample.csv
id,user,name
1,user1,name2
2,user2,name2
3,user_3rd,name_3rd
3,user_3rd,name_3rd
```

### DELETE
```
$ quickquery 'delete from sample.csv where user = user_3rd'
$ cat sample.csv
id,user,name
1,user1,name2
2,user2,name2
```

## development
Run `cobra-cli` on docker

```bash
$ docker-compose run cobra-cli <command>
```
