# PortalSS
A proof of concept of a multiparty chat for Android on Swarm PSS

This PoC use a modified version of swarm developed by vocdoni: [SimpleSwarm](https://github.com/vocdoni/go-dvote/tree/master/swarm)

# Download Android APK

[Download](https://github.com/eduadiez/PortalSS/new/master?readme=1)

# Compile gomobile library
```
$ go get -v golang.org/x/mobile/cmd/gomobile
$ gomobile init
$ git clone https://github.com/eduadiez/PortalSS && cd PortalSS
$ gomobile bind -o gomobile/gomobile.aar -target=android portalss
```

# Run a desktop client
```
$ git clone https://github.com/vocdoni/go-dvote && go-dvote
$ GOPATH="" go run ./cmd/pssChat/psschat.go --datadir ./pssChat --light false --nick "DesktopClient" --topic "topic"
```
