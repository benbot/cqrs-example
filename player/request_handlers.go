package player

import (
	"encoding/json"

	"github.com/kataras/iris"
)

func AddPlayer(ctx *iris.Context) {
	id, err := player_add()
	if err != nil {
		ctx.SetStatusCode(500)
		return
	}

	p, err := player_projection(id)
	if err != nil {
		ctx.SetStatusCode(500)
		return
	}

	var resp []byte
	resp, err = json.Marshal(p)
	if err != nil {
		ctx.SetStatusCode(500)
		return
	}

	ctx.SetStatusCode(201)
	ctx.SetContentType("application/json")
	ctx.SetBody(resp)
}

func GetPlayer(ctx *iris.Context) {
	id := ctx.Param("id")

	p, err := player_projection(id)
	if err != nil {
		ctx.SetStatusCode(400)
		ctx.SetBodyString(err.Error())
		return
	}

	resp, err := json.Marshal(p)
	if err != nil {
		ctx.SetStatusCode(500)
		ctx.SetBodyString(err.Error())
		return
	}

	ctx.SetStatusCode(200)
	ctx.SetContentType("application/json")
	ctx.SetBodyString(string(resp))
}
