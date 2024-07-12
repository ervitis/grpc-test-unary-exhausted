package data

import (
	"github.com/ervitis/grpc-test-unary-exhausted/pb_impl"
	"strconv"
)

type Chat struct {
	Content          string
	Personalizations []*ChatPersonalization
}

type ChatPersonalization struct {
	Email         string
	Substitutions map[string]string
}

type ChatPersonalizations []*ChatPersonalization

func MockData(n int) Chat {
	perss := make(ChatPersonalizations, 0)
	for i := 0; i < n; i++ {
		p := &ChatPersonalization{
			Email:         "email" + strconv.Itoa(i) + "@aafsda.es",
			Substitutions: map[string]string{"fefe": "eeeee"},
		}
		perss = append(perss, p)
	}
	return Chat{
		Content:          "Hello World",
		Personalizations: perss,
	}
}

func NewRpcMessage(n int) []*pb_impl.Personalization {
	mockData := MockData(n)
	personalizations := make([]*pb_impl.Personalization, len(mockData.Personalizations))
	for i, p := range mockData.Personalizations {
		substitutions := p.Substitutions
		personalizations[i] = &pb_impl.Personalization{Email: p.Email, Substitutions: substitutions}
	}
	return personalizations
}

func MockMessage(n int) string {
	var msg string
	for i := 0; i < n; i++ {
		msg += "a"
	}
	return msg
}
