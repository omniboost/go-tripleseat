module github.com/omniboost/go-tripleseat

go 1.23.2

require (
	github.com/gorilla/schema v1.1.0
	github.com/joefitzgerald/passwordcredentials v0.2.0
	github.com/pkg/errors v0.9.1
	golang.org/x/oauth2 v0.11.0
	gopkg.in/guregu/null.v3 v3.5.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.14.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

replace github.com/gorilla/schema => github.com/omniboost/schema v1.1.1-0.20211111150515-2e872025e306
