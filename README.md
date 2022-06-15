# 3lab-kpi
### Software architecture, laboratory work #3
### Variant 1
A system of forums where users communicate at random
topics. Each of the forums has a keyword-defined topic.
Users are subscribed to updates in the forums in which they are interested

Example
```bash
{

 “name”: “Політика в Україні”,
 “topicKeyword”: “ukraine-politics”,
 “users”: [ “user1”, “user2” ]
 
}
```
## The server starts by a command

```bash
go run ./cmd/server
```

## A client runs by a command

```bash
go run cmd/example/main.go
```
