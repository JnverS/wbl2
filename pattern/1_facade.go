package main

import (
	"errors"
	"fmt"
)

/*
	Реализовать паттерн «фасад».
	Объяснить применимость паттерна, его плюсы и минусы,
	а также реальные примеры использования данного примера на практике.
*/

type WalletFacade struct {
	account      *Account
	wallet       *Wallet
	securityCode *SecurityCode
}

// когда создаем новый фасад внутри него создаем акк, код, кошелек
func newWalletFacade(accountId string, code int) *WalletFacade {
	fmt.Println("Starting create account")
	walletFacade := &WalletFacade{
		account:      newAccount(accountId),
		securityCode: newSecurityCode(code),
		wallet:       newWallet(),
	}
	fmt.Println("Account created")
	return walletFacade
}

// метод добавления в кошелек обращается к методам проверки акка, кода и добавляет
func (w *WalletFacade) addMoneyToWallet(accountId string, securityCode int, amount int) error {
	fmt.Println("Trying add money to wallet")
	err := w.account.checkAccount(accountId)
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	w.wallet.creditBalance(amount)
	fmt.Println("added!")
	return nil
}

//метод покупки аналогично, проверяет правильность данных и совершает списание
func (w *WalletFacade) Buy(accountId string, securityCode int, amount int) error {
	fmt.Println("Trying debit money from wallet")
	err := w.account.checkAccount(accountId)
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	err = w.wallet.debitBalance(amount)
	if err != nil {
		return err
	}
	fmt.Println("Сompleted!")
	return nil
}

type Account struct {
	accountId string
}

func newAccount(accountId string) *Account {
	return &Account{
		accountId: accountId,
	}
}

func (a *Account) checkAccount(accountId string) error {
	if a.accountId != accountId {
		return errors.New("Account does not exist")
	}
	return nil
}

type SecurityCode struct {
	securityCode int
}

func newSecurityCode(code int) *SecurityCode {
	return &SecurityCode{
		securityCode: code,
	}
}

func (s *SecurityCode) checkCode(code int) error {
	if s.securityCode != code {
		return errors.New("Wrong code!")
	}
	return nil
}

type Wallet struct {
	balance int
}

func newWallet() *Wallet {
	return &Wallet{
		balance: 0,
	}
}

func (w *Wallet) creditBalance(amount int) {
	w.balance += amount
}

func (w *Wallet) debitBalance(amount int) error {
	if w.balance < amount {
		return errors.New("Not enough money")
	}
	w.balance -= amount
	return nil
}

func main() {
	var err error

	// создадим фасад для работы с аккаунтом
	myAccount := newWalletFacade("Yuli", 123456)

	err = myAccount.addMoneyToWallet("Yuli", 123456, 500)
	err = myAccount.Buy("Yuli", 123456, 499)
	err = myAccount.Buy("yuli", 123456, 10)
	fmt.Println(err)
	err = myAccount.Buy("Yuli", 12345, 10)
	fmt.Println(err)
	err = myAccount.Buy("Yuli", 123456, 10)
	fmt.Println(err)
}

/*
	Фасад используется для упрощения взаимодействия со сложными подсистемами
	Применяется:
	- когда нужно предоставить простой или урезанный интерфейс к сложной подсистеме
	- когда нужно разложить подсистему на отдельные слои
	Преимущества: изоляция клиента от компонентов сложной подсистемы
	Недостатки: риск стать объектом, который хранит/делает сликом много и привязанн ко всем классам программы
*/
