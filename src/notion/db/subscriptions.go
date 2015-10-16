package db

import (
	"notion/model"
)

func GetSubscriptionsByUserId(id string) ([]model.DbSubscription, error) {
	var subscriptions []model.DbSubscription
	subscriptionsG, err := GenericGetMultiple("subscriptions", "user_id", id, &subscriptions)
	return *subscriptionsG.(*[]model.DbSubscription), err
}
