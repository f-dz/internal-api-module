package data

import (
	"context"
	"net/http"

	"github.com/api-abc/internal-api-module/model/response"
)

func (cl *Client) Delete(ctx context.Context, name string) (resp response.BodyResponse, err error) {
	err = cl.CallWithoutBody(ctx, http.MethodDelete, "/internalapi/data/"+name, &resp)
	return
}

func (cl *Client) GetDeleted(ctx context.Context) (resp response.BodyResponseGet, err error) {
	err = cl.CallWithoutBody(ctx, http.MethodGet, "/internalapi/data", &resp)
	return
}
