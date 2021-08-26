package main

import "./email"

func main() {
	email.SendEmail(func(b *email.EmailBuilder) {
		b.From("email@email.com").
			To("email@email.com").
			Subject("subject").
			Body("Body")
	})
}
