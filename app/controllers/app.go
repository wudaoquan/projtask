package controllers

import "github.com/robfig/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Hello(myName string) revel.Result {
	c.Validation.Required(myName).Message("请输入姓名!")
	c.Validation.MinSize(myName, 3).Message("姓名不少于三个字符!")

	if c.Validation.HasErrors() {
		//c.Flash.Error("输入信息有误!")
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}
	//c.Flash.Success("正常执行!")
	return c.Render(myName)
}
