package email

import "fmt"

type EmailBuilder struct {
	Email Email
}

func (b *EmailBuilder) From(from string) *EmailBuilder {
	b.Email.From = from
	return b
}

func (b *EmailBuilder) To(to string) *EmailBuilder {
	b.Email.To = to
	return b
}

func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	b.Email.Subject = subject
	return b
}

func (b *EmailBuilder) Body(body string) *EmailBuilder {
	b.Email.Body = body
	return b
}

func SendEmailImpl(email *Email) {
	fmt.Println(email)
}

type build func(builder *EmailBuilder)

func SendEmail(action build) {
	builder := EmailBuilder{}
	action(&builder)
	SendEmailImpl(&builder.Email)
}
