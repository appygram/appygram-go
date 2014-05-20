package appygram

type AppygramMessage struct {
	Name     string
	Email    string
	Phone    string
	Message  string
	Platform string
	Software string
	Summary  string
	Topic    string
	App_json string
}

type AppygramTrace struct {
	AppygramMessage
	Class     string
	Backtrace [][]string
}

type AppygramResult struct {
	OK bool
}

type AppygramTopic struct {
	Name string
	Id   string
}

type AppygramTopics struct {
	AppygramResult
	Topics []AppygramTopic
}
