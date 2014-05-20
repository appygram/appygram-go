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

type AppygramResult struct {
	OK bool
}
