package config

import (
	"context"
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/vartanbeno/go-reddit/reddit"
	"math/rand"
	"strings"
)




var helpBanner string = "" +
	"Gracias por usar esta wea `%s` !\n" +	// Name of the requester
	"Prefix `%s`\n" + // Prefix on config structure
	"Command list:\n" +
	"\t\t`;queryrandom <subreddit name> <optional: default:new | hot | rising | controversial | top>`\n" +
	"\t\t`;qr <same as queryrandom>`\n" +
	"\t\t`;explore <query to search for subreddits>`\n" +
	"\t\t`;ex <same as explore>`\n" +
	"\t\t`;x <same as ex>`\n\n\n" +
	"Give a star or a coffe here: `%s`\n"	//	github url :D





func (c *Config) help(s *discordgo.Session, m *discordgo.MessageCreate){

	hb := fmt.Sprintf(helpBanner, m.Author, c.Prefix, "github.com/jcatala/drgob")
	if c.Verbose{
		fmt.Printf("Sending help banner!\n%s", hb)
	}
	s.ChannelMessageSend(m.ChannelID,hb)
}


/*
	Function to query random post from a given subreddit
*/

func (c *Config) checkIfSubredditExists(s string)(error){
	sub,res,err := reddit.DefaultClient().Subreddit.Get(context.Background(), s)
	if c.Verbose{
		fmt.Println(res.StatusCode)
	}
	// If error from GET or statusCode not found or sub == nil (which is returned from a not-existed subreddit)
	if err != nil || res.StatusCode == 404 || sub == nil{
		err = errors.New("Reddit does not exists")
		return err
	}
	return nil
}




/*
*	TODO
*	Do the following functions to only return an slice of posts,
*	Do the logic of sorting, randomize and shit outside this function
*	Make a function to make the logic behind the auto-reaction on another functions
*/
// func (c *Config) getPosts(s *discordgo.Session, m *discordgo.MessageCreate, query []string) {
func (c *Config) queryRandomPost(s *discordgo.Session, m *discordgo.MessageCreate, query []string)([]*reddit.Post, error) {
	var err error
	var posts []*reddit.Post
	// The check if there's more than 2 arguments came before this function
	// So now, first we need to check if the subreddit exists
	if c.Verbose {
		fmt.Printf("Checking the following query: %s\n", query)
	}
	err = c.checkIfSubredditExists(query[1])
	// Just return if the subreddit does not exists
	if err != nil{
		message := fmt.Sprintf("`%s-kun`, el sub-reddit `%s` no existe, gil ", m.Author, query[1])
		s.ChannelMessageSend(m.ChannelID, message)
		return nil, err
	}
	if len(query) >= 3{
		switch strings.ToLower(query[2]) {
		case "hot":
			posts, _, err = c.RedditThings.RedditClient.Subreddit.HotPosts(
				context.Background(),
				query[1],
				&reddit.ListOptions{
					Limit:  c.Nposts,
				})
		case "new":
			posts,_,err = c.RedditThings.RedditClient.Subreddit.NewPosts(
				context.Background(),
				query[1],
				&reddit.ListOptions{
					Limit:  c.Nposts,
				})
		case "rising":
			posts, _, err = c.RedditThings.RedditClient.Subreddit.RisingPosts(
				context.Background(),
				query[1],
				&reddit.ListOptions{
					Limit:  c.Nposts,
				})
		case "controversial":
			posts, _, err = c.RedditThings.RedditClient.Subreddit.ControversialPosts(
				context.Background(),
				query[1],
				&reddit.ListPostOptions{
					ListOptions: reddit.ListOptions{
						Limit:  10,
					},
					Time:        "week",
				})
		case "top":
			posts, _, err = c.RedditThings.RedditClient.Subreddit.TopPosts(
				context.Background(),
				query[1],
				&reddit.ListPostOptions{
					ListOptions: reddit.ListOptions{
						Limit:  c.Nposts,
					},
					Time:        "week",
				})
		default:
			posts,_,err = c.RedditThings.RedditClient.Subreddit.NewPosts(
				context.Background(),
				query[1],
				&reddit.ListOptions{
					Limit:  c.Nposts,
				})
		}
	} else {
		posts,_,err = c.RedditThings.RedditClient.Subreddit.NewPosts(
			context.Background(),
			query[1],
			&reddit.ListOptions{
				Limit:  c.Nposts,
			})
	}
	// Any kind of error of the return value of the reddit api
	if err != nil{
		return nil, err
	}

	// Get random post from it
	if len(posts) < 1 {
		// Dunno why sometimes the API does return zero posts, just in case, lol
		noPosts := errors.New("No post returned")
		return	nil, noPosts
	}
	return posts, nil
	/*
	post := posts[rand.Intn(len(posts))]
	message := fmt.Sprintf("`%s`\n%s\n%s",post.Title, post.Body, post.URL)
	if c.Verbose{
		fmt.Printf("Sending the following message to the channel %s\n", message)
	}

	// Catch the new message
	mes ,err := s.ChannelMessageSend(m.ChannelID, message)
	if err != nil{
		return
	}
	// Get all the emojis of the current guild
	allEmojis, err := s.GuildEmojis(m.GuildID)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	// Select a random emoji from the current guild and react
	// First check if the guild have custom emojis, lol
	if len(allEmojis) < 1{
		if c.Verbose{
			fmt.Println("Error, the guild have no customs emoji's")
		}
		return
	}

	randEmoji := allEmojis[rand.Intn(len(allEmojis))]
	err = s.MessageReactionAdd(mes.ChannelID, mes.ID, randEmoji.APIName() )
	if err != nil{
		fmt.Println(err.Error())
		return
	}

	return
	*/
}

func (c *Config) toggleReaction(s* discordgo.Session, m *discordgo.MessageCreate, mes *discordgo.Message, )(error){
	// Get all the emojis from the current guild
	allEmojis, err := s.GuildEmojis(m.GuildID)
	if err != nil{
		return err
	}
	// Check if the current guild have a custom emoji
	if len(allEmojis) < 1 {
		if c.Verbose{
			fmt.Println("Error, the guild have no customs emoji's")
		}
		// Randomize the message of "create a emoji pls"
		// Pretty fucking bad decision tbh.
		err = errors.New("No emoji available")
		if rand.Intn(100) == 50{
			err = errors.New("No emoji available, send")
		}
		return err
	}
	// Select a random emoji
	randEmoji := allEmojis[rand.Intn(len(allEmojis))]

	err = s.MessageReactionAdd(mes.ChannelID, mes.ID, randEmoji.APIName())
	if err != nil{
		if c.Verbose{
			fmt.Println(err.Error())
		}
	}
	return err
}

/*
*	Use the slice of posts to search trhough the only media things
*	TODO: Investigate which things are useful to render on discord:
*	jpg, jpeg, png, gif, svg, gfycat, imgur
*	From now, we're just doing basic fucking greping here,
*	Need to search another method to discrimante between media and non-media post
*/

func (c *Config) findInside(post *reddit.Post)(bool){

	for i,word := range c.oracle{
		if c.Verbose{
			fmt.Printf("[ %d ] Looking for word: %s inside:\nBody: %s\nURL:%s\nPermalink: %s\n\n", i, word, post.Body,post.URL, post.Permalink)
		}
		if strings.Contains(post.Body, word) || strings.Contains(post.URL, word) {
			return true
		}
	}
	return false
}

func (c *Config) queryMedia(s *discordgo.Session, m *discordgo.MessageCreate, query []string){
	posts, err := c.queryRandomPost(s, m, query)
	if err != nil{
		return
	}
	// Shuffle the posts in order to iterate over it and grepping for our target
	// This is a stupid fucking implementation, lol.
	rand.Shuffle(len(posts), func(i, j int) { posts[i], posts[j] = posts[j], posts[i] })
	// After the shuffle, iterate over it and look up for our oracle of "media things"
	for _,p := range posts{
		// If we found our oracle, send the message and end the routine
		if c.findInside(p){
			message := fmt.Sprintf("`%s`\n%s\n%s",p.Title, p.Body, p.URL)
			mes, err := s.ChannelMessageSend(m.ChannelID, message)
			if err != nil{
				return
			}
			if c.Verbose{
				fmt.Printf("Sended the following message: %s", mes.Content)
			}
			_ = c.toggleReaction(s,m,mes)
			return
		}
	}
	return
}

func (c *Config) queryAll(s *discordgo.Session, m *discordgo.MessageCreate, query []string){
	posts, err := c.queryRandomPost(s, m, query);
	if err != nil{
		return
	}
	post := posts[rand.Intn(len(posts))]
	message := fmt.Sprintf("`%s`\n%s\n%s",post.Title, post.Body, post.URL)
	if c.Verbose{
		fmt.Printf("Sending the following message to the channel %s\n", message)
	}

	// Catch the message to make the reaction

	// Catch the new message
	mes ,err := s.ChannelMessageSend(m.ChannelID, message)
	if err != nil{
		return
	}

	// With the message sended and created, toggle a reaction
	err = c.toggleReaction(s, m, mes)
	if err != nil{
		if strings.Contains(err.Error(), "send"){
			s.ChannelMessageSend(m.ChannelID,"Please, create a custom emoji, 8=D")
		}
	}

}


func (c *Config) explore(s *discordgo.Session, m *discordgo.MessageCreate, arg []string){

	query := strings.Join(arg, " ")
	if c.Verbose{
		fmt.Printf("Starting to explore subrredit on %s", query)
	}
	sr, _, err := c.RedditThings.RedditClient.Subreddit.Search(context.Background(),query, &reddit.ListSubredditOptions{
		ListOptions: reddit.ListOptions{
			Limit:  c.Nsr,
		},
		Sort:        "activity",
	})
	if err != nil{
		fmt.Printf("Error geting subreddits for this query, %s", err.Error())
		return
	}
	// Start building the result, we need to add the backquote to end the markdown mode
	result := fmt.Sprintf("Result for search %s:\n\n```",query)
	for i, s := range sr {
		result = fmt.Sprintf("%s\t[ %d ] %s\n", result, i+1, s.Name )
	}
	result = fmt.Sprintf("%s\n```", result)

	if c.Verbose{
		fmt.Printf("Got the following subreddits: %s", result)
	}
	s.ChannelMessageSend(m.ChannelID, result)
	return
}

func (c *Config) Core(s *discordgo.Session, m *discordgo.MessageCreate){
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
	command := strings.ToLower(args[0])
	if c.Verbose {
		fmt.Printf("Prefix received: %s\n", string(command[0]) )
	}
	// If the the prefix does not correspond, quit
	if strings.Compare(string(command[0]), c.Prefix) != 0{
		if c.Verbose{
			fmt.Printf("Bad prefix!\nReceived: %s\nExpected: %s\n", command[0], c.Prefix)
		}
		return
	}

	// Delete the prefix for the incoming command
	command = strings.Replace(command, ";","", -1 )
	if c.Verbose {
		fmt.Printf("Received command: %s\n", command)
	}
	// Case of each command
	switch command {
	case "h", "help":
		go c.help(s, m)

	case "qr", "queryrandom":
		if len(args) >= 2{
			//go c.queryRandomPost(s, m, args)
			go c.queryAll(s, m, args)
		}
	case "qrm", "querymedia","qm":
		if len(args) >=2{
			go c.queryMedia(s, m, args)
		}
	case "explore", "ex", "x":
		if len(args) >= 2{
			go c.explore(s, m, args[1:])
		}
	}
}