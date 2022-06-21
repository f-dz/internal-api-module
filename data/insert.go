package data

import (
	"context"
	"net/http"

	"github.com/api-abc/internal-api-module/model/request"
	"github.com/api-abc/internal-api-module/model/response"
)

func (cl *Client) Insert(ctx context.Context, req request.InsertRequest) (resp response.BodyResponse, err error) {
	err = cl.Call(ctx, http.MethodPost, "/internalapi/data", req, &resp)
	return
}

func (cl *Client) GetInserted(ctx context.Context) (resp response.BodyResponseGet, err error) {
	err = cl.CallWithoutBody(ctx, http.MethodGet, "/internalapi/data", &resp)
	return
}
