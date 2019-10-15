package command

import (
	"fmt"
	"kbot/bot"
	"strings"

	. "github.com/ahmetb/go-linq/v3"
)

type ListOrdersCommand struct {
	ordersStore *OrdersStore
	posts       *bot.Posts
}

func NewListOrdersCommand(ordersStore *OrdersStore, posts *bot.Posts) *ListOrdersCommand {
	return &ListOrdersCommand{ordersStore, posts}
}

func (c ListOrdersCommand) CanHandle(message Message) bool {
	return ParseListOrders(message.Text)
}

func (c ListOrdersCommand) Handle(message Message) error {
	ordersMessage, err := c.CreateAllOrdersMessage()
	if err != nil {
		return err
	}

	return c.posts.Create(*ordersMessage, "")
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

	From(orders).
		SelectT(func(o Order) string {
			return fmt.Sprintf("User: %s, Order: %s", o.UserId, o.Order)
		}).
		ToSlice(&orderDetails)

	result := strings.Join(orderDetails, "\r\n")
	return &result, nil
}
