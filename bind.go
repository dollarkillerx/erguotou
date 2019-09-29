/**
 * @Author: DollarKiller
 * @Description: 参数绑定
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 23:48 2019-09-29
 */
package erguotou

import "github.com/dollarkillerx/erguotou/formband"

func BandValue(ctx *Context, obj interface{}) error {
	err := formband.Band(ctx.Ctx, obj)
	return err
}

func BandFrom(ctx *Context, obj interface{}) error {
	err := formband.BindForm(ctx.Ctx, obj)
	return err
}

func BandJson(ctx *Context, obj interface{}) error {
	err := formband.BindJson(ctx.Ctx, obj)
	return err
}
