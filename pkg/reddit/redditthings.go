package reddit

import (
	"fmt"
	"github.com/vartanbeno/go-reddit/reddit"
)


type RedditThings struct {
	RedditCredentials *reddit.Credentials	//	Save credentials object
	RedditClient *reddit.Client	//	Save client reddit object
}


func NewRedditThings(v bool, id string, secret string, username string, password string)(r *RedditThings,e error){
	r = new(RedditThings)
	// Create reddit object credentials
	cred := new(reddit.Credentials)
	cred.ID = id
	cred.Secret = secret
	cred.Username = username
	cred.Password = password
	// Assign to our RedditThings
	r.RedditCredentials = cred
	if v{
		fmt.Printf(`
Creating new reddit credential object with:
ID: %s,
Secret: %s
Username: %s
Password: ****
`, cred.ID, cred.Secret, cred.Username)
	}
	// Create the client who's responsible to talk with the api
	client, err := reddit.NewClient(*r.RedditCredentials)
	r.RedditClient = client	// Pass the client to our object
	return r,err
}