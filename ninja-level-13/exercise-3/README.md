# Exercise 3 for Ninja level 13

---

1. Run app:
    ```shell
    go run main.go
    ```
1. Run tests:
    ```shell
    go test
    ```
1. Run banchmark:
    ```shell
    go test -bench .
    ```
1. Run coverage:
    ```shell
    go test -cover
    ```
1. Run coverage and save result in file:
    ```shell
    go test -coverprofile c.out
    ```
1. Run coverage and show it in web browser:
    ```shell
    go tool cover -html=c.out
    ```
1. Run docs:
    ```shell
    godoc -http=:8080
    ```
