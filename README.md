## Hello, World! <img src="https://raw.githubusercontent.com/iampavangandhi/iampavangandhi/master/gifs/Hi.gif" width="30px"></h2>
<img align="right" alt="GIF" src="https://media.giphy.com/media/13HgwGsXF0aiGY/giphy.gif" />

### Slowly это:
- Максимально простое API
- Один роут
- Один тест
- И, конечто же, CI : ![CI](https://img.shields.io/github/workflow/status/LeftTwixWand/Slowly/Go)
### Как заставить это работать?
- Открыть консоль
- git clone https://github.com/LeftTwixWand/Slowly
- go get -v -t -d ./...
- go run ./
- Test: go test -v ./...
- Проверить через Postman: http://localhost:8080/api/slow?timeout=1000
