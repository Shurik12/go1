# go1

### BUILD
```bash
go mod init github.com/Shurik1266/go1
# Add third-party packages and dependencies
go mod tidy
go install .
go clean -modcache
```

### RUN
```bash
go run main.go
# or
go run . 
```
