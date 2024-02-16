module github.com/myname/goodbot

go 1.16

replace github.com/myname/goodbot/languages/go => ./languages/go

require (
	github.com/jamesog/iptoasn v0.1.0
	github.com/yl2chen/cidranger v1.0.2
)
