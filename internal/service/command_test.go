package service

import (
	"testing"
	"users-app/domain"

	"github.com/stretchr/testify/assert"
)

func TestModifyUserCommand_fieldsToUpdate(t *testing.T) {
	strPtr := func(s string) *string { return &s }
	tests := []struct {
		name    string
		command ModifyUserCommand
		want    domain.Fields
	}{
		{"empty", ModifyUserCommand{}, domain.Fields{}},
		{
			"FirstName",
			ModifyUserCommand{FirstName: strPtr("John")},
			domain.Fields{"first_name": "John"},
		},
		{
			"LastName",
			ModifyUserCommand{LastName: strPtr("Doe")},
			domain.Fields{
				"last_name": "Doe",
			},
		},
		{
			"Nickname",
			ModifyUserCommand{Nickname: strPtr("jdoe")},
			domain.Fields{"nickname": "jdoe"},
		},
		{
			"Email",
			ModifyUserCommand{Email: strPtr("j@doe.com")},
			domain.Fields{"email": "j@doe.com"},
		},
		{
			"Country",
			ModifyUserCommand{Country: strPtr("UK")},
			domain.Fields{
				"country": "UK",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualValues(t, tt.want, tt.command.fieldsToUpdate())
		})
	}
}
