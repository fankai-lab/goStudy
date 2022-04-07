package otherMain

import (
	"fmt"
)
type acount struct{
	name string
	pass string
	money float64
}
func NewAcount(name, pass string,money float64) *acount {
	if len(name) < 6 || len(name) > 10 {
		fmt.Println("用户名不正确")
		return nil
	}
	if len(pass) != 6 {
		fmt.Println("密码格式不正确")
		return nil
	}
	if money < 20 {
		fmt.Println("金额不能低于20元")
		return nil
	}
	return &acount{
		name,
		pass,
		money,
	}
}
func (a *acount) SetMoney(money float64) {
	if (a.money + money) < 20 {
		fmt.Println("存入金额与总金额不能小于20")
	}else {
		a.money = money
	}
}
func (a *acount) GetMoney() float64 {
	return a.money
}
