package db

import (
	"notion/log"
	"notion/model"
)

func GetSubscriptionsByUserId(id string) ([]model.DbSubscription, error) {
	log.Info("Getting all subscriptions for user %v", id)
	var subscriptions []model.DbSubscription
	subscriptionsG, err := GenericGetMultiple("subscriptions", "user_id", id, &subscriptions)
	return *subscriptionsG.(*[]model.DbSubscription), err
}

func CreateSubscription(sub model.DbSubscription) error {
	log.Info("Creating new subscription for user %v notebook %v", sub.UserId, sub.NotebookId)
	return dbmap.Insert(&sub)
}

func RemoveSubscription(sub model.DbSubscription) error {
	log.Info("Removing subscription for user %v notebook %v", sub.UserId, sub.NotebookId)
	_, err := dbmap.Delete(&sub)
	return err
}
