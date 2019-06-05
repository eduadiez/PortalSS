# PortalSS
A proof of concept of a multiparty chat for Android on Swarm PSS

This PoC use a modified version of swarm developed by vocdoni: [SimpleSwarm](https://github.com/vocdoni/go-dvote/tree/master/swarm)

# Download Android APK

[Download](https://github.com/eduadiez/PortalSS/new/master?readme=1)

# Compile gomobile library
```
$ go get golang.org/x/mobile/cmd/gomobile
$ gomobile init
$ go get github.com/ethereum/go-ethereum/log && go get github.com/spf13/pflag && go get github.com/vocdoni/go-dvote/swarm
$ mkdir -p $GOPATH/src/portalss
$ cp gomobile/gomobile.go $GOPATH/src/portalss
$ gomobile bind -o gomobile/gomobile.aar -target=android portalss
```

# Run a desktop client
```
$ go run gomobile/pssChat.go --datadir ./pssChat --light false --nick "DesktopClient" --topic "topic"
```
