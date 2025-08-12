# G.A.T.O.R (Get A Ton Of RSS)

Welcome to **G.A.T.O.R**, a simple RSS blog aggregator I built as a portfolio project. It helps you gather, follow, and manage tons of RSS feeds all in one place — perfect for staying up-to-date with your favorite blogs and news sources.

## What it does

G.A.T.O.R lets you:

- **login** — get into your account
- **register** — create a new user
- **reset** — reset users and feeds
- **users** — list all users (for admin or fun)
- **agg** — aggregate the latest posts from feeds you follow
- **addfeed** — add a new RSS feed to the system
- **feeds** — list all available feeds
- **follow** — follow a specific feed to get updates
- **following** — see which feeds you’re currently following

## How to use

You will need Postgres and Go installed to run the program.

Install by running 'go install github.com/kaelbroersma/gator'

USAGE : gator <cmd> <..args>
Some commands have args.. others dont. 

command 'agg' is meant to be running in the background. It is what retrieves the post title, description, url, etc.
Run CLI in another instance to run browse commands and fetch posts

command list: 

```bash
login (login as user)
register (register a user)
reset (reset users and feeds)
users (list users)
agg (start aggregating) <duration> (1s,1m,1h)
addfeed (add feed to aggregate)
feeds (list feeds)
follow (follow a feed)
unfollow (unfollow a feed)
following (list feeds you are following)
browse (browse feeds of logged in user)

