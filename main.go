package lambdaGinContext

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gin-gonic/gin"
	req "github.com/grokify/go-awslambda"
)

func Bind(e events.APIGatewayProxyRequest) (*gin.Context, error) {
	form, err := req.NewReaderMultipart(e)
	if err != nil {
		return nil, err
	}
	multipartForm, err := form.ReadForm(0)
	if err != nil {
		return nil, err
	}

	var c gin.Context
	headers := http.Header{}
	for k, v := range e.Headers {
		headers.Add(k, v)
	}
	c.Request = &http.Request{MultipartForm: multipartForm, Header: headers}
	return &c, nil
}
