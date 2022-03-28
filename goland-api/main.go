package main

import (
	"goland-api/configs"
	"goland-api/routes"
	"log"
	"net/http"

	//"github.com/aws/aws-sdk-go/aws/session"
	//"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/gin-gonic/gin"
)

type DiscordMessage struct {
	//channel_id string `json:"channel_id"`
	content string `json:"content"`
}

// func initDb() {
// 	sess := session.Must(session.NewSessionWithOptions(session.Options{
// 		SharedConfigState: session.SharedConfigEnable,
// 	}))
// 	svc := dynamodb.New(sess)
// }
// func discordHandler(c *gin.Context) {
// 	//content := c.Param("name")
// 	var (
// 		webhook_id    string = "956936242178695238"
// 		webhook_token        = "uT_Hq2DKq7yln38E9stRtToQp_UNBDu3LNgwlMsNlGBUpHH9hf3AAp3l6IpwTfrUfHuW"
// 		base_url      string = "https://discord.com/api/webhooks"
// 	)

// 	body := &DiscordMessage{
// 		content: "Hi",
// 	}
// 	var url string = base_url + "/" + webhook_id + "/" + webhook_token
// 	var jsonByte = []byte(`
// 		{
// 			"name": "test webhook",
// 			"type": 1,
// 			"content": "test a dyingapple webhook 2",
// 			"channel_id": "199737254929760256",
// 			"token": "3d89bb7572e0fb30d8128367b3b1b44fecd1726de135cbe28a41f8b2f777c372ba2939e72279b94526ff5d1bd4358d65cf11",
// 			"avatar": null,
// 			"guild_id": "199737254929760256",
// 			"id": "223704706495545344",
// 			"application_id": null,
// 			"user": {
// 			"username": "test",
// 			"discriminator": "7479",
// 			"id": "190320984123768832",
// 			"avatar": "b004ec1740a63ca06ae2e14c5cee11f3",
// 			"public_flags": 131328
// 			}
// 		}
// 	`)
// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonByte))
// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)

// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()

// 	fmt.Println("response Status:", resp.Status)
// 	fmt.Println("response Headers:", resp.Header)
// 	// body, _ := ioutil.ReadAll(resp.Body)
// 	// fmt.Println("response Body:", string(body))
// }

func helloHandler(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}

func welcomeHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello World from Go")
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"text": "Welcome to gin lambda server.",
	})
}

// func boardHandler(c *gin.Context) {
// 	mongoconn := "mongodb://dyingapple:dyingapple@localhost:27017/"
// }

func routerEngine() *gin.Engine {
	// set server mode
	gin.SetMode(gin.DebugMode)

	r := gin.New()

	configs.ConnectDB()

	// Global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/welcome", welcomeHandler)
	r.GET("/user/:name", helloHandler)
	r.GET("/", rootHandler)
	// r.GET("/discord/:message", discordHandler)

	routes.BoardRoutes(r) //add this

	return r
}

func main() {
	//addr := ":" + os.Getenv("PORT")
	// log.Fatal(gateway.ListenAndServe("8080", routerEngine()))
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", routerEngine()))
}
