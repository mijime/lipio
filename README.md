lipio
===

```
lipio serv ws://:8080 | gunzip > path/to/file
lipio serv ws://:8081 | lipio pipe ws://localhost:8080
cat path/to/file | gzip | lipio pipe ws://localhost:8080
```

```
lipio file path/to/file.txt | gzip | lipio file path/to/file.txt.gz
```

```
lipio s3 example/path/to/file | lipio lambda example | lipio ssh example.com/var/log/example.log
```
