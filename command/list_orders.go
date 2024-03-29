package command

import (
	"fmt"
	"kbot/bot"
	"strings"

	. "github.com/ahmetb/go-linq/v3"
)

type ListOrdersCommand struct {
	ordersStore *OrdersStore
	messages    *bot.Messages
}

func NewListOrdersCommand(ordersStore *OrdersStore, messages *bot.Messages) *ListOrdersCommand {
	return &ListOrdersCommand{ordersStore, messages}
}

func (c ListOrdersCommand) CanHandle(message Message) bool {
	return ParseListOrders(message.Text)
}

func (c ListOrdersCommand) Handle(message Message) error {
	ordersMessage, err := c.CreateAllOrdersMessage()
	if err != nil {
		return err
	}

	return c.messages.Send(*ordersMessage)
}

func (c ListOrdersCommand) Help() string {
	return "kbot list. lists all orders"
}

func (c ListOrdersCommand) CreateAllOrdersMessage() (*string, error) {
	var orderDetails []string

	orders, err := c.ordersStore.List()
	if err != nil {
		return nil, err
	}

	if len(orders) == 0 {
		result := "No orders found for today"
		return &result, nil
	}

	From(orders).
		SelectT(func(o Order) string {
			return fmt.Sprintf("User: %s, Order: %s", o.UserName, o.Value)
		}).
		ToSlice(&orderDetails)

	result := strings.Join(orderDetails, "\r\n")
	return &result, nil
}
