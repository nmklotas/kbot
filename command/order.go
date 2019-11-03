package command

import (
	"errors"
	"fmt"
	"kbot/bot"
	"strings"
)

type OrderCommand struct {
	ordersStore *OrdersStore
	messages    *bot.Messages
	users       *bot.Users
}

func NewOrderCommand(ordersStore *OrdersStore, messages *bot.Messages, users *bot.Users) *OrderCommand {
	return &OrderCommand{ordersStore, messages, users}
}

func (p OrderCommand) CanHandle(message Message) bool {
	_, addErr := ParseOrderToAdd(message.Text)
	return addErr == nil
}

func (p OrderCommand) Handle(message Message) error {
	return p.SaveOrder(message)
}

func (p OrderCommand) Help() string {
	return "kbot order {letter}. saves order for the user"
}

func (p OrderCommand) SaveOrder(message Message) error {
	orderToAdd, err := ParseOrderToAdd(message.Text)
	if err == nil {
		return p.AddOrderToStore(message.UserId, orderToAdd)
	}

	return errors.New("Order not saved")
}

func (p OrderCommand) AddOrderToStore(userId string, order *string) error {
	newOrder, err := p.createOrder(userId, *order)
	if err != nil {
		return err
	}

	if err := p.ordersStore.Add(newOrder); err != nil {
		return err
	}

	message := fmt.Sprintf("Order %s saved", strings.ToUpper(*order))
	return p.messages.Send(message)
}

func (p OrderCommand) createOrder(userId string, order string) (Order, error) {
	userName, err := p.users.Name(userId)
	if err != nil {
		return Order{}, err
	}

	return Order{UserId: userId, Value: order, UserName: *userName}, nil
}
