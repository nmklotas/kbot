package command

import (
	"errors"
	"fmt"
	"kbot/bot"
	"strings"
)

type SaveOrderCommand struct {
	ordersStore *OrdersStore
    posts       *bot.Posts
    users *bot.Users
}

func NewSaveOrderCommand(ordersStore *OrdersStore, posts *bot.Posts, users *bot.Users) *SaveOrderCommand {
	return &SaveOrderCommand{ordersStore, posts, users}
}

func (p SaveOrderCommand) CanHandle(message Message) bool {
	_, addErr := ParseOrderToAdd(message.Text)
	_, removeErr := ParseOrderToRemove(message.Text)
	return addErr != nil || removeErr != nil
}

func (p SaveOrderCommand) Handle(message Message) error {
	return p.SaveOrder(message)
}

func (p SaveOrderCommand) Help() string {
	return "kbot order {letter}. saves order for the user"
}

func (p SaveOrderCommand) SaveOrder(message Message) error {
	orderToAdd, err := ParseOrderToAdd(message.Text)
	if err == nil {
		return p.AddOrderToStore(message.UserId, orderToAdd)
	}

	orderToRemove, err := ParseOrderToRemove(message.Text)
	if err == nil {
		return p.RemoveOrderFromStore(message.UserId, *orderToRemove)
	}

	return errors.New("Order not saved")
}

func (p SaveOrderCommand) AddOrderToStore(userId string, order *string) error {
    newOrder, err := p.createOrder(userId, *order);
    if err != nil {
        return err
    }
    
	if err := p.ordersStore.Add(newOrder); err != nil {
		return err
	}

	message := fmt.Sprintf("Order %s saved", strings.ToUpper(*order))
	return p.posts.Create(message, "")
}

func (p SaveOrderCommand) RemoveOrderFromStore(userId string, order string) error {
	if err := p.ordersStore.Remove(userId, order); err != nil {
		return err
	}

	message := fmt.Sprintf("Order %s removed", strings.ToUpper(order))
	return p.posts.Create(message, "")
}

func (p SaveOrderCommand) createOrder(userId string, order string) (Order, error) {
    userName, err := p.users.Name(userId); 
    if err != nil {
		return Order{}, err
	}

	return Order{UserId: userId, Value: order, UserName: *userName}, nil
}
