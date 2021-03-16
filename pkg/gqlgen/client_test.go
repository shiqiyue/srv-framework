package gqlgen

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Result struct {
	GetById Enterprise `json:"getById"`
}
type Enterprise struct {
	EnterpriseId string `json:"enterprise_id"`
}

func TestNewClient(t *testing.T) {
	client := NewClient("http://localhost:8080/org/enterprise/query")
	request := &Request{
		q:      "query q1{\n  getById(enterpriseId:\"da\"){\n    enterprise_id\n  }\n}\n",
		vars:   nil,
		files:  nil,
		Header: nil,
	}
	result := &Result{}
	err := client.runWithJSON(context.Background(), request, result)
	fmt.Println(err)
	fmt.Println(result)
	assert.Nil(t, err)
}
