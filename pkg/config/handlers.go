package config

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/vartanbeno/go-reddit/reddit"
	"math/rand"
	"strings"
)



func (c *Config) help(s *discordgo.Session, m *discordgo.MessageCreate){
	helpBanner := `Gracias por usar esta wea `+"`%s`" + `!

Prefix `+"`;`"+`

Command list:
	` + "`;queryrandom <subreddit name> <optional: default:new | hot | rising | controversial | top`"+`
	` + "`;qr same as queryrandom`" + `
`
	helpBanner = fmt.Sprintf(helpBanner, m.Author)
	s.ChannelMessageSend(m.ChannelID,helpBanner)
}


/*
	Function to query random post from a given subreddit
*/
func (c *Config) queryRandomPost(s *discordgo.Session, m *discordgo.MessageCreate,
	query []string) {
	nPosts := 50
	var err error
	var posts []*reddit.Post
	if len(query) >= 3{
		switch strings.ToLower(query[2]) {
		case "hot":
			posts, _, err = c.RedditThings.RedditClient.Subreddit.HotPosts(
				context.Background(),
				query[1],
				&reddit.ListOptions{
					Limit:  nPosts,
				})
		case "new":
			posts,_,err = c.RedditThings.RedditClient.Subreddit.NewPosts(
				context.Background(),
				query[1],
				&reddit.ListOptions{
					Limit:  nPosts,
				})
		case "rising":
			posts, _, err = c.RedditThings.RedditClient.Subreddit.RisingPosts(
				context.Background(),
				query[1],
				&reddit.ListOptions{
					Limit:  nPosts,
				})
		case "controversial":
			posts, _, err = c.RedditThings.RedditClient.Subreddit.ControversialPosts(
				context.Background(),
				query[1],
				&reddit.ListPostOptions{
					ListOptions: reddit.ListOptions{
						Limit:  10,
					},
					Time:        "month",
				})
		case "top":
			posts, _, err = c.RedditThings.RedditClient.Subreddit.TopPosts(
				context.Background(),
				query[1],
				&reddit.ListPostOptions{
					ListOptions: reddit.ListOptions{
						Limit:  nPosts,
					},
					Time:        "month",
				})
		default:
			posts,_,err = c.RedditThings.RedditClient.Subreddit.NewPosts(
				context.Background(),
				query[1],
				&reddit.ListOptions{
					Limit:  nPosts,
				})
		}
	} else {
		posts,_,err = c.RedditThings.RedditClient.Subreddit.NewPosts(
			context.Background(),
			query[1],
			&reddit.ListOptions{
				Limit:  nPosts,
			})
	}
	if err != nil{
		return
	}
	// Get random post from it
	post := posts[rand.Intn(len(posts))]
	message := fmt.Sprintf("`%s`\n%s\n%s",post.Title, post.Body, post.URL)
	s.ChannelMessageSend(m.ChannelID, message)
	return
}

func (c *Config) GetRandomPost(s *discordgo.Session, m *discordgo.MessageCreate){
	// Ignore self bots messages
	if m.Author.ID == c.DiscordThings.DiscordSession.State.User.ID{
		return
	}
	user := m.Author
	content := m.Content
	if c.Verbose{
		fmt.Printf("Received a message from %s with the following content:%s\n",user,content)
	}
	//	Check if the incoming message is more than just a prefix
	if len(content) <= len(c.Prefix){
		return
	}
	args := strings.Fields(content)
	name := strings.ToLower(args[0])

	fmt.Println(args, name)
	// Case of each command
	switch name {
	case ";h", ";help":
		c.help(s, m)

	case ";qr", ";queryrandom":
		c.queryRandomPost(s, m, args)

	}

}