## vanadium-demo

### General instructions

1. go mod init <directory i want to run .go file>
2. go build
3. go run <name>.go

### Set Principal & Blessings

#### Create Principal (server)

```
go run v.io/x/ref/cmd/principal create --with-passphrase=false --overwrite ~/vanadium-demo/chat/server/cred/dimi dimi
```

#### Give bless to teo (as dimi, the server owner) of type: dimi:friend:teo

```
go run v.io/x/ref/cmd/principal bless \
    --v23.credentials ~/vanadium-demo/chat/server/cred/dimi \
    --for=24h ~/vanadium-demo/chat/client/cred/teo friend:teo | \
        go run v.io/x/ref/cmd/principal \
            --v23.credentials ~/vanadium-demo/chat/client/cred/teo \
            set forpeer - dimi
```

### Testing

#### Start client with blessings (as teo)

```
go run ~/vanadium-demo/chat/client/main.go --v23.credentials ~/vanadium-demo/chat/client/cred/teo
```

#### Start server with creds (as dimi)

```
go run ~/vanadium-demo/chat/server/main.go --v23.credentials ~/vanadium-demo/chat/server/cred/dimi
```
