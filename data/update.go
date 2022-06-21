package data

import (
	"context"
	"net/http"

	"github.com/api-abc/internal-api-module/model/request"
	"github.com/api-abc/internal-api-module/model/response"
)

func (cl *Client) Update(ctx context.Context, req request.UpdateRequest, name string) (resp response.BodyResponse, err error) {
	err = cl.Call(ctx, http.MethodPut, "/internalapi/data/"+name, req, &resp)
	return
}

func (cl *Client) GetUpdated(ctx context.Context) (resp response.BodyResponseGet, err error) {
	err = cl.CallWithoutBody(ctx, http.MethodGet, "/internalapi/data", &resp)
	return
}
