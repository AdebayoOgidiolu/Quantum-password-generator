# Quantum-password-generator
Password generator based on a quantum source

Todos:

1. Use qrand
1. Use typeable character set
1. Be able to specify different charsets ("All Unicode" or "Typeable")

Simple default case:
```go
password := qpass.NewPassword(length)
```

Power user case, with configuration:
```go
p := qpass.NewPasswordGenerator()
p.Charset = qpass.TypeableCharacters
p.NewPassword(length)
```