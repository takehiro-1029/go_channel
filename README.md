# 概要
## 目的
　メッセージ処理でgo channelの簡単な実装をしてポーリングや非同期処理に関しての理解を深めたかった

## 内容
　post /messsageでchannelにmessageIDを格納
　channelが受信したら、別処理を動かしてmessage内容を取得してコンソール画面に表示

## ディレクトリ構成

```bash
.
├── README.md                   <-- This instructions file
├── main.go                     <-- DynamoDBとの連携
├── chat.go                     <-- Cognitoとの連携
├── go.mod                      <-- Cognitoとの連携
├── go.sum                      <-- Cognitoとの連携
```

## テスト方法
1. go run *.go
2. 別コンソールを開き、以下コマンドを入力
　```
　curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"body":"Hello World"}' \
  http://localhost:8080/messages
  ```
3. コンソール画面にmessege.IDが出力される

