appygram.go
===========

Appygram go http client


###Installation

First appygram.go must be in your GOPATH. Do this where you like.

```
go get github.com/wlaurance/appygram-go
```

###Usage

Import the package

```go
...
import (
  ...
  "github.com/wlaurance/appygram.go"
  ...
)
```

####Appygram Client

```go
appygram.Client(apiKey string, httpClient *http.Client) appygram.AppygramClient
```

To obtain an Appygram Client, use the above function. Example

```go
client := appygram.Client("myapikey12345", &http.Client{})
```

If you are using **Google App Engine**, pass a reference to a
http.Client obtained from the [appengine/urlfetch](https://developers.google.com/appengine/docs/go/urlfetch/)
package like this:

```go
c := appengine.NewContext(r) //r *http.Request
client := appygram.Client("myapikey12345", urlfetch.Client(c))
```

####Getting Appygram Topics

Appygram Topics are how you route different messages to different
places.

```go
var topics appygram.AppygramTopic
topics, err := client.GetTopics()
```

####Sending Appygram Message

An Appygram message is intended for human generated content. This will
suffice for pushing feedback through a web form or equivalent.

```go
appygramMessage := appygram.AppygramMessage{
  Name: "Gopher John", Topic: "Feedback",
  Message: "Your App is a lot of fun!!", Email: "john@gophernet.net",
}
var response appygram.AppygramResult
response, err := appygram.SendAppygramMessage(appygramMessage)
//handle err and response accordingly
```

####Sending Appygram Trace (In Progress)

An Appygram trace is intended for machine generated stack traces or
errors.

```go
appygramTrace := appygram.AppygramTrace{Class: "My File", Message: "Gophers"}
//appygramTrace.Backtrace = appygram.MakeErrorBackTrace(bytes)
//Optionally add AppygramMessage
appygramMessage := appygram.AppygramMessage{
  Name: "Gopher John", Topic: "Feedback",
  Message: "Your App is a lot of fun!!", Email: "john@gophernet.net",
}
appygramTraceWithMessage := ,
appygramAppygramTraceWithMessage{
  Trace: appygramTrace, AppygramMessage: appygramMessage,
}
var response appygram.AppygramResult
response, err := appygram.SendAppygramTrace(appygramAppygramTraceWithMessage)
//handle err and response accordingly
```

###Notes on Concurrency

This is a blocking http client. All requests will block the current
routine. It is up to the caller to handle their concurrency needs.

If you want the request to happen asynchronously do something like this:

```go
...
go func() {
  appygramMessage := appygram.AppygramMessage{
    Name: "Gopher John", Topic: "Feedback",
    Message: "Your App is a lot of fun!!", Email: "john@gophernet.net",
  }
  var response appygram.AppygramResult
  response, err := appygram.SendAppygramMessage(appygramMessage)
  //handle err and response accordingly
}()
```

###Developing and Testing!

We love PRs and issues! If you are having a problem feel free to open up
an issue.

If you need to test locally please do the following:

```
export API_KEY="SOME API KEY FOR AN APP YOU CONTROL"
go test
```

If you need to run a particular test case, use the `-run` option

Example: Running just the topics test

```
export API_KEY="SOME API KEY FOR AN APP YOU CONTROL"
go test -run GetTopics
```

