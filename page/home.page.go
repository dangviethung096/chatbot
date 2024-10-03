package page

import "github.com/dangviethung096/core"

func HomePage(ctx *core.HttpContext, request *core.PageRequest) (core.PageResponse, core.Error) {

	return core.PageResponse{
		PageFiles:    []string{"./html/home_page.html"},
		TemplateName: "home_page",
	}, nil
}
