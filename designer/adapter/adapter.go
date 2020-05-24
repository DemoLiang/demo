package main

import "log"

type Bank interface {
	Deposit(m int)
	Withdraw(m int)
}

type CBCBank struct {
	CBCMoney int
}

func (c CBCBank) CBCDeposit(m int) {
	c.CBCMoney += m
	log.Printf("CBCBank deposit %v balance:%v\n", m, c.CBCMoney)
	return
}

func (c CBCBank) CBCWithdraw(m int) {
	c.CBCMoney -= m
	log.Printf("CBCBank withdraw %v balance:%v\n", m, c.CBCMoney)
	return
}

type CBCBankAdapter struct {
	MyBank CBCBank
}

func (this CBCBankAdapter) Deposit(m int) {
	this.MyBank.CBCDeposit(m)
}

func (this CBCBankAdapter) Withdraw(m int) {
	this.MyBank.CBCWithdraw(m)
}

type ICBCBank struct {
	ICBCMoney int
}

func (c ICBCBank) ICBCDeposit(m int) {
	c.ICBCMoney += m
	log.Printf("ICBCBank deposit %v balance:%v\n", m, c.ICBCMoney)
	return
}

func (c ICBCBank) ICBCWithdraw(m int) {
	c.ICBCMoney -= m
	log.Printf("ICBCBank withdraw %v balance:%v\n", m, c.ICBCMoney)
	return
}

type ICBCBankAdapter struct {
	MyBank ICBCBank
}

func (this ICBCBankAdapter) Deposit(m int) {
	this.MyBank.ICBCDeposit(m)
}

func (this ICBCBankAdapter) Withdraw(m int) {
	this.MyBank.ICBCWithdraw(m)
}

func main() {
	cbc := CBCBank{CBCMoney: 100} //初始化
	cbcAdapter := CBCBankAdapter{MyBank: cbc}
	log.Printf("CBCBank Init Money Balance:%v\n", cbc.CBCMoney)

	myDeposit(cbcAdapter, 10)  //从CBC存
	myWithdraw(cbcAdapter, 20) //从cbc取

	icbc := ICBCBank{ICBCMoney: 200} //初始化icbc 200
	icbcAdater := ICBCBankAdapter{MyBank: icbc}
	log.Printf("ICBCBank Init Money Balance:%v\n", icbc.ICBCMoney)

	myDeposit(icbcAdater, 20)  //存icbc 20
	myWithdraw(icbcAdater, 50) //取icbc 50

	return
}

func myDeposit(b Bank, m int) { //银行只有存取接口
	b.Deposit(m)
}

func myWithdraw(b Bank, m int) {
	b.Withdraw(m)
}
