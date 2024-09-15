## Error logs

`var` is not used with card & is directly assigned
```go
func main() {

        card = "Ace of spades"
```
```bash
# command-line-arguments
./main.go:7:2: undefined: card
```

---
variable declared once using `:=` need not be used again for assigning new value
```go
        cardOne := "Five of hearts"
        cardOne := "Five of diamonds" // use cardOne = "Five of diamonds"
```
```bash
# command-line-arguments
./main.go:22:10: no new variables on left side of :=
```
