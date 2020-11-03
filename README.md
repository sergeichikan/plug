# plug

Build plugin:

```go build -buildmode=plugin -o p/main.so p/main.go```

Build core:

```go build main.go```

Run:

```./main```

or

```./main -plugin=./p/main.so -cycle=100000000 -length=100000```
