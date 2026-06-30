package main

import (
	"context"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/cloudwego/eino/schema"
)

type Game struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type InputParams struct {
	Name string `json:"name" jsonschema:"description=the name of game"`
}

func GetGame(_ context.Context, params *InputParams) (string, error) {
	GameSet := []Game{
		{Name: "原神", Url: "https://www.mihoyo.com/"},
		{Name: "鸣潮", Url: "https://www.kurogames.com/"},
		{Name: "明日方舟", Url: "https://www.hypergryph.com/"},
	}
	for _, game := range GameSet {
		if game.Name == params.Name {
			return game.Url, nil
		}
	}
	return "", nil
}

func CreateTool() tool.InvokableTool {
	getGameToool := utils.NewTool(
		&schema.ToolInfo{
			Name: "get_game",
			Desc: "get a game url by name",
			ParamsOneOf: schema.NewParamsOneOfByParams(
				map[string]*schema.ParameterInfo{
					"name": &schema.ParameterInfo{
						Type:     schema.String,
						Desc:     "game's name",
						Required: true,
					},
				},
			),
		}, GetGame)
	return getGameToool
}
