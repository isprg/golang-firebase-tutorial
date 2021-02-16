## GolangでFirebase Realtime Databaseを取得するチュートリアル

<br>`※firebase.goでは、data/flag配下のflagの状態を取得しています。`

```go
flags := map[string]int
// key  (string) -> ギミックの名前
// value(int)    -> ギミックの状態
```

### 依存パッケージの導入

`$ go get firebase.google.com/go`
`$ go get google.golang.org/api/option`