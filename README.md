# GO BBS

This is BBS using Go Lang

using [goose](https://bitbucket.org/liamstask/goose) for migration


## Preview

![Preview](images/gobbs.gif "Preview")

##### Functions
- Create User
- Login.
- Post, Update, Delete the text.



### How To set Up 

##### set up mysql
first, use should create user `gobbsuser`

```
CREATE USER 'gobbsuser'@'localhost' IDENTIFIED BY 'gobbs';
GRANT ALL PRIVILEGES ON gobbs.* TO 'gobbsuser'@'localhost';
```

type below to apply all available migrations

```
goose up
```

to run go server

```
go run main.go
```

and will appear on the console `Starting Goji on [::]:8000`

So you can use Go BBS to access below URL

```
localhost:8000
```





