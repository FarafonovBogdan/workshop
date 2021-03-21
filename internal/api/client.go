package api

//Client ...
type Client interface {
	GetJoke() (*JokeResponse, error)
}
