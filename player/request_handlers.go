package player

import "github.com/kataras/iris"

func AddPlayer(ctx *iris.Context) {
	id, err := player_add()
	if err != nil {
		ctx.SetStatusCode(500)
		return
	}

	ctx.SetStatusCode(201)
	ctx.SetContentType("application/json")
	ctx.SetBody([]byte(id))
}
