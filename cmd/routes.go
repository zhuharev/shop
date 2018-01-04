// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import(
  "github.com/zhuharev/tamework"
  "github.com/zhuharev/shop/routes/shops"
	"github.com/zhuharev/shop/routes/products"
	
)

func registreRoutes(tw *tamework.Tamework) {
    tw.Text("shops.add_director_", shops.AddDirector)
		tw.Cb("shops.create", shops.Create)
		tw.Cb("shops.list", shops.List)
		tw.Prefix("shops.set_geo_", shops.SetGeo)
		tw.Prefix("shops.withdraw_", shops.Withdraw)
		tw.Prefix("shops.stats", shops.Stats)
		tw.Prefix("shops.add_manager", shops.AddManager)
		tw.Prefix("shops.toggle_state", shops.ToggleState)
		tw.Prefix("products.upload_photo_", products.Upload)
		tw.Prefix("products.set_price_", products.SetPrice)
		tw.Prefix("products.set_title_", products.SetTitle)
		tw.Cb("products.search", products.Search)
		tw.Prefix("products.buy_", products.Buy)
		tw.Prefix("products.buy_qiwi_", products.BuyQiwi)
		
}