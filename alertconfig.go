package alertconfig

//package  main

import (
	"fmt"
	"github.com/robfig/config"
	"strconv"
)

type ServiceConfig struct {
	Video_dir     string
	Audio_file    string
	Email_server  string
	From_Numbertw string
	TW_site       string
}

type Authtokens struct {
	TwaccountSid string
	TwauthToken  string
	SlackToken   string
	WatsonToken  string
	WatsonPass   string
}
type CallMember struct {
	Name    string
	Group   string
	Phone   string
	Slackid string
	Email   string
}
type KeysMServices struct {
}

//var places = []*CallMember{}

const maxList = 1024

func BuildCallList(places *[]*CallMember, serviceconfig *ServiceConfig, authTokens *Authtokens) {
	loc := new(CallMember)
	c, _ := config.ReadDefault("alert.ini")
	fmt.Println("---------")
	serviceconfig.Video_dir, _ = c.String("Service-Config", "VIDEO_DIR")
	serviceconfig.Audio_file, _ = c.String("Service-Config", "AUDIO_FILE")
	serviceconfig.Email_server, _ = c.String("Service-Config", "EMAIL_SERVER")
	serviceconfig.From_Numbertw, _ = c.String("Service-Config", "TWFROM_NUMBER")
	serviceconfig.TW_site, _ = c.String("Service-Config", "TWSITE")
	fmt.Println("Server Config", serviceconfig)

	authTokens.TwaccountSid, _ = c.String("Authtokens", "TWILIO-AccountSid")
	authTokens.TwauthToken, _ = c.String("Authtokens", "TWILIO-AUTHTOKEN")
	authTokens.SlackToken, _ = c.String("Authtokens", "SLACKTOKEN")
	authTokens.WatsonToken, _ = c.String("Authtokens", "WATSONTOKEN")
	authTokens.WatsonPass, _ = c.String("Authtokens", "WATSONPASS")
	fmt.Println("Tokens", authTokens)
       
	loc = new(CallMember)
	for i := 1; i < maxList; i++ {
		loc.Name, _ = c.String("Member"+strconv.Itoa(i), "Name")
		if loc.Name == "" {
			continue
		}
		loc.Group, _ = c.String("Member"+strconv.Itoa(i), "Group")
		loc.Phone, _ = c.String("Member"+strconv.Itoa(i), "Phone")
		loc.Slackid, _ = c.String("Member"+strconv.Itoa(i), "Slackid")
		loc.Email, _ = c.String("Member"+strconv.Itoa(i), "Email")
		*places = append(*places, loc)
		fmt.Println(loc)
		loc = new(CallMember)
	}
	// Loop over all indexes in the slice.
	// ... Print all struct data.
	
	for i := range *places {
		place := *places
		fmt.Printf("Location:%s %s %s %d\n", place[i].Name, place[i].Phone, place[i].Group, i)
	}

}
func main() {
	var places = []*CallMember{}
	var serviceconfig ServiceConfig
	var authTokens Authtokens
	BuildCallList(&places, &serviceconfig, &authTokens)
	fmt.Println("Server Configuraiton", serviceconfig)
	fmt.Println("Tokens", authTokens)
	fmt.Println("LIST BEFORE")
	fmt.Println(places[0])
	fmt.Printf("LIST\n")
	for i := range places {
		place := places[i]
		fmt.Println("Location:", place.Phone)
	}
}
