# (D)iscord (r)eddit (Go) (b)ot weaaaaa


Hi everyone!

This is a `discord bot` that will respond with `posts` from an arbitrary subreddit.


# Commands

```bash
;help Display the help message
;h (same as help)
;queryrandom <subreddit name> <optional: default:new | hot | rising | controversial | top>
;qr (same as queryrandom)
```


# Installation

If you want to run it by yourself (instead of invite the official (my bot) to your server), you the following:

* Discord bot `token`
* Reddit account 
* Create Reddit application (app token and secret)

* Install the bot 

```bash
go get github.com/jcatala/drgob
```

Create the `.env` file with the following content

```bash
discord_key=discord_key
reddit_username=reddit_username
reddit_password=reddit_password
reddit_id=reddit_app_id
reddit_secret=reddit_app_secret
```


Run the bot

```bash
drgob -verbose
```



