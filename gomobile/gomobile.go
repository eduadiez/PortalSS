package gomobile

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/log"
	flag "github.com/spf13/pflag"
	swarm "github.com/vocdoni/go-dvote/swarm"
)

var gSS *swarm.SimpleSwarm
var gTopic *string
var gEnc *string
var gKey *string
var gAddr *string
var gNick *string

type Message struct {
	Type int    `json:"type"`
	Nick string `json:"nick"`
	Data string `json:"message"`
}

// callback
var jc JavaCallback

type JavaCallback interface {
	SendString(string)
}

func RegisterJavaCallback(c JavaCallback) {
	jc = c
	jc.SendString("RegisterJavaCallback!")
}

func SendMessage(text string) string {

	if gSS = nil {
		return "Not connected"
	}

	var jmsg Message
	jmsg.Type = 0
	jmsg.Nick = *gNick

	jmsg.Data = text
	msg, err := json.Marshal(jmsg)
	if err != nil {
		log.Crit(err.Error())
	}
	err = gSS.PssPub(*gEnc, *gKey, *gTopic, fmt.Sprintf("%s", msg), *gAddr)
	if err != nil {
		log.Warn(err.Error())
	}

	jc.SendString(fmt.Sprintln("Me: ", text))
	return "Success!"
}

func Starting(datadir string, topic string, nick string) string {
	go start(datadir, topic, nick)
	return "Started!"
}

func start(_datadir string, _topic string, _nick string) {

	topic := &_topic
	nick := &_nick
	dir := &_datadir
	kind := flag.String("encryption", "sym", "encryption key schema (raw, sym, asym)")
	key := flag.String("key", "vocdoni", "encryption key (sym or asym)")
	addr := flag.String("address", "", "pss address to send messages")
	light := flag.Bool("light", true, "use light mode (less consumption)")
	logLevel := flag.String("log", "info", "log level (info, warn, crit)")
	flag.Parse()

	sn := new(swarm.SimpleSwarm)
	sn.LightNode = *light
	gSS = sn
	gTopic = topic
	gEnc = kind
	gKey = key
	gAddr = addr
	gNick = nick
	sn.SetDatadir(*dir)

	err := sn.InitPSS()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	err = sn.SetLog(*logLevel)
	if err != nil {
		fmt.Printf("Cannot set loglevel %v\n", err)
	}

	sn.PssSub(*kind, *key, *topic)
	defer sn.PssTopics[*topic].Unregister()

	log.Info("My PSS pubKey is %s", sn.PssPubKey)

	var jmsg Message
	jmsg.Type = 0
	jmsg.Nick = *nick
	jmsg.Data = "I'm connected!!!"
	msg, err := json.Marshal(jmsg)
	if err != nil {
		log.Crit(err.Error())
	}

	err = sn.PssPub(*gEnc, *gKey, *gTopic, fmt.Sprintf("%s", msg), *gAddr)
	if err != nil {
		jc.SendString("Error: " + err.Error())
		log.Warn(err.Error())
	}

	jc.SendString("Process started!!!!!")

	listener(*topic, *kind, *key, *nick, *addr, sn, *light)
}

func listener(topic, enc, key, mynick, addr string, sn *swarm.SimpleSwarm, light bool) {
	jc.SendString("Start listening...")
	var nick string
	var msg string
	var jmsg Message
	for {
		pmsg := <-sn.PssTopics[topic].Delivery

		err := json.Unmarshal(pmsg.Msg, &jmsg)

		if err != nil {
			nick = "raw"
			msg = fmt.Sprintf("%s", pmsg.Msg)
		} else {
			nick = jmsg.Nick
			msg = jmsg.Data
		}

		log.Info("%s <%s>: %s\n", time.Now().Format("3:04PM"), nick, msg)
		jc.SendString(fmt.Sprintf("%s <%s>: %s\n", time.Now().Format("3:04PM"), nick, msg))
	}

}