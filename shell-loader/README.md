# Shell loader exec

#### 1. Encrypt shell

* Go to [`shell encryption`](./shell-encryption)

```shell
cd shell-encryption
```

* Run the script with the given shell script to encrypt

```shell
go run ./main.go <file.sh>
```

You will find a file `output.sh` that you can copy
and past to [`shell-exec`](./shell-exec) as `file.sh`

##### 2. Embed script in go binary

* Build the binary

```shell
go build -o bin .
```

:bulb: This embed the shell script `file.sh` and create an executable
binary that will execute the script using syscall.

* Execute binary

```shell
./bin
```

