package gpt

import gogpt "github.com/sashabaranov/go-openai"

type Status int

const (
	OK Status = iota
	Banned
	OutOfService
	OutOfQuota
)

func (s Status) String() string {
	switch s {
	case OK:
		return "OK"
	case Banned:
		return "BANNED"
	case OutOfService:
		return "OOS"
	case OutOfQuota:
		return "OOQ"
	default:
		return "Unknown"
	}
}

type Client struct {
	cli    *gogpt.Client
	Status Status
	Token  string
}

func New(token string) (g *Client) {
	return &Client{
		cli:    gogpt.NewClient(token),
		Status: OK,
		Token:  token,
	}
}

func (g *Client) IsOk() bool {
	return g.Status == OK
}

func (g *Client) IsBanned() bool {
	return g.Status == Banned
}

func (g *Client) IsOutOfService() bool {
	return g.Status == OutOfService
}
