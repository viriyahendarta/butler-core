package service

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/viriyahendarta/butler-core/config"
	"github.com/viriyahendarta/butler-core/infra/constant"

	_context "github.com/viriyahendarta/butler-core/infra/context"
	e "github.com/viriyahendarta/butler-core/infra/error"
)

const (
	defaultSuccessMessage = "success"
	defaultErrorMessage   = "error"
)

func (r *Resource) RenderJSON(ctx context.Context, w http.ResponseWriter, data interface{}, successHTTPCode int, err error) error {
	var response interface{}

	httpCode := successHTTPCode

	if oErr := e.Cast(err); oErr != nil {
		httpCode = oErr.GetHTTPCodeEquivalent()
		response = r.buildErrorResponse(ctx, oErr)
	} else {
		response = APISuccessResponse{
			Meta: r.buildMeta(ctx),
			Data: data,
		}
	}

	byteData, err := json.Marshal(response)
	if err != nil {
		return e.New(ctx, e.CodeRenderResponse, "failed to marshal response", err)
	}

	w.Header().Add("Content-Type", constant.ContentTypeApplicationJSON)
	w.WriteHeader(httpCode)
	_, err = w.Write(byteData)

	return err
}

func (r *Resource) buildErrorResponse(ctx context.Context, err *e.Error) APIErrorResponse {
	errResponse := APIError{
		Code: err.GetCode(),
		Type: err.GetTitle(),
	}
	if err.GetError() != nil {
		errResponse.Messages = err.GetMessages()
	}

	if config.Get().Debug || !err.IsInternalServerError() {
		errResponse.Reason = err.GetError().Error()
	}

	return APIErrorResponse{
		Meta:  r.buildMeta(ctx),
		Error: errResponse,
	}
}

func (r *Resource) buildMeta(ctx context.Context) APIMeta {
	return APIMeta{
		ProcessTime: _context.GetElapsedTime(ctx),
	}
}
