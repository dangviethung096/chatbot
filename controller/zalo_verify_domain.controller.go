package controller

import (
	"io"
	"net/http"
	"os"

	"github.com/dangviethung096/core"
)

type ZaloVerifyDomainRequest struct {
}

func ZaloVerifyDomain(ctx *core.HttpContext, request ZaloVerifyDomainRequest) (core.HttpResponse, core.HttpError) {

	filePath := "./html/zalo_verifierNSUpDw2mA3vSku0KduiFI6kZXnU7rLTAD34q.html"
	file, err := os.Open(filePath)
	if err != nil {
		return nil, core.NewHttpError(http.StatusInternalServerError, 500, "Unable to open file", err)
	}
	defer file.Close()

	body, err := io.ReadAll(file)
	if err != nil {
		return nil, core.NewHttpError(http.StatusInternalServerError, 500, "Unable to read file", err)
	}
	header := http.Header{}

	ctx.EndResponse(http.StatusOK, &header, body)

	return nil, nil
}
