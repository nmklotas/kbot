package command

import (
	"errors"
	"fmt"
	"kbot/bot"
	"strings"
)

type ForgetOrderCommand struct {
	ordersStore *OrdersStore
	messages    *bot.Messages
	users       *bot.Users
}

func NewForgetOrderCommand(ordersStore *OrdersStore, messages *bot.Messages, users *bot.Users) *ForgetOrderCommand {
	return &ForgetOrderCommand{ordersStore, messages, users}
}

func (p ForgetOrderCommand) CanHandle(message Message) bool {
	_, removeErr := ParseOrderToRemove(message.Text)
	return removeErr == nil
}

func (p ForgetOrderCommand) Handle(message Message) error {
	return p.RemoveOrder(message)
}

func (p ForgetOrderCommand) Help() string {
	return "kbot forget {letter}. forgets order for the user"
}

func (p ForgetOrderCommand) RemoveOrder(message Message) error {
	orderToRemove, err := ParseOrderToRemove(message.Text)
	if err == nil {
		return p.RemoveOrderFromStore(message.UserId, *orderToRemove)
	}

	return errors.New("Order not saved")
}

func (p ForgetOrderCommand) RemoveOrderFromStore(userId string, order string) error {
	if err := p.ordersStore.Remove(userId, order); err != nil {
		return err
	}

	message := fmt.Sprintf("Order %s removed", strings.ToUpper(order))
	return p.messages.Send(message)
}
