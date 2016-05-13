lipio
===

```
lipio serv --addr :8081 | gunzip
lipio serv --addr :8080 | sed s/hello/world/g | lipio post http://localhost:8081
echo test | gzip | lipio post http://localhost:8080
```

```
lipio file path/to/file.txt | gzip | lipio file path/to/file.txt.gz
```

```
lipio s3 example/path/to/file | lipio lambda example | lipio ssh example.com/var/log/example.log
```
