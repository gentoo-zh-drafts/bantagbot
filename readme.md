删除群组中以 #GentooZh 开头的消息

编译:

```
$ CGO_ENABLED=0 go build -o bot
$ strip ./bot
```

运行:

```
$ tgbot_token="xxxx" ./bot
2025/03/23 11:42:15 Authorized on account deleletag_bot
2025/03/23 11:42:31 已删除消息 ID 138xxxx
```
