// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package context

import (
	"time"

	"gitlab.com/zhuha/cashroulette/models"
	"gitlab.com/zhuha/cashroulette/pkg/btn"
	"gitlab.com/zhuha/cashroulette/pkg/setting"

	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
	base62 "github.com/pilu/go-base62"
	"github.com/zhuharev/tamework"
)

type Context struct {
	*tamework.Context

	User *models.User
}

func Contexter() tamework.Handler {
	return func(c *tamework.Context) {
		color.Cyan("middleware context")
		ctx := &Context{
			Context: c,
		}

		user, err := models.UsersGetByTelegramID(c.UserID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				registreFlow(ctx)
			}
			// TODO:
			color.Red("%s", err)
		} else {
			ctx.Data["user"] = user
			ctx.User = user
		}

		c.Map(ctx)
	}
}

func registreFlow(c *Context) {
	referID := base62.Decode(c.Text) / setting.App.SecretNumber
	user, err := models.UsersCreate(c.UserID, uint(referID))
	if err != nil {
		// TODO:
		return
	}
	//bonus
	user.BonusBalance += setting.App.RegBonus
	c.Data["user"] = user
	c.User = user

	c.Keyboard = btn.MenuKeyboard("en-US")
	c.Markdown(c.T("wi_give_you", setting.App.RegBonus))

	time.Sleep(1 * time.Second)
}
