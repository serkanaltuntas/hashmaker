# hashmaker

Creates MD5 hash values for all files inside a directory recursively.

## Build:
```bash
cd <into the code>
go build
```

## Usage:
```bash
./hashmaker -path "." -file "output.txt"
```

## Cross Platform Builds:
```
GOOS=target-OS GOARCH=target-architecture go build
```

### Example Windows Build:
```
GOOS=windows GOARCH=amd64 go build
```
