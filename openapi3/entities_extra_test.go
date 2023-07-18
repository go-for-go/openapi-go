package openapi3_test

import (
	"testing"

	"github.com/go-for-go/openapi-go/openapi3"
	"github.com/stretchr/testify/assert"
)

func TestSpec_MarshalYAML(t *testing.T) {
	var s openapi3.Spec

	spec := `openapi: 3.0.3
info:
  description: description
  license:
    name: Apache-2.0
    url: https://www.apache.org/licenses/LICENSE-2.0.html
  title: title
  version: 2.0.0
servers:
  - url: /v2
paths:
  /user:
    put:
      summary: updates the user by id
      operationId: UpdateUser
      requestBody:
        content:
          application/json:
            schema:
              type: string
        description: Updated user object
        required: true
      responses:
        "404":
          description: User not found
components:
  securitySchemes:
    api_key:
      in: header
      name: x-api-key
      type: apiKey
    bearer_auth:
      type: http
      scheme: bearer
      bearerFormat: JWT`

	assert.NoError(t, s.UnmarshalYAML([]byte(spec)))
}
