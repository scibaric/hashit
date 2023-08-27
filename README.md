# hashit

CLI tool for generating bcrypt hashed passwords.

```bash
hashit mypassword
```

Above command should produce:
```
$myhashedpassword
```

## Build
Build the program with standard `go build` command and place it your `~/.local/bin`.

```bash
go build -o hashit main.go
cp hashit ~/.local/bin
```
