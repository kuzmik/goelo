# Goelo 3.GO

Generally I learn a new programming language by writing an IRC bot. I would like to replace my ruby IRC bot with a statically compiled bot that is just one binary. So here we are.

## Dependencies
- IRC library: [irc-event](https://github.com/thoj/go-ircevent)
- SQL library: [go-sqlite3](https://github.com/mattn/go-sqlite3)
- Log library: [glog](http://github.com/golang/glog)
- Twitter library: [go-twitter](https://github.com/dghubble/go-twitter)

```
go get github.com/golang/glog
go get github.com/thoj/go-ircevent
go get github.com/mattn/go-sqlite3
go get github.com/dghubble/go-twitter/twitter
```

## MVP Requirements
- [x] IRC connectivity, obviously
- [x] Logging IRC messages to a database
- [ ] Twitter functionality
    - [ ] Post tweets from IRC
    - [ ] Monitor mentions of the bot and post them to IRC channel


## TODO later
* Bebot nonesense (markov chains, replies on mentions based on a word)
* The other plugins like last.fm, linkgrabber, etc
