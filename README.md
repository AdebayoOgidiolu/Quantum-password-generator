# Quantum-password-generator
Password generator based on a quantum source

Todos:

1. Use qrand
1. Use typeable character set
1. Be able to specify different charsets ("All Unicode" or "Typeable")

Simpler:
```go
// Base case
p := qpass.NewPassword(10)
...
// Power user case
g := qpass.NewPasswordGenerator(qpass.TypeableCharacters)
p := g.NewPassword(length)
```