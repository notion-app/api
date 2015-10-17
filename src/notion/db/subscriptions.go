package db

import (
	"notion/errors"
	"notion/log"
	"notion/model"
)

func GetSubscriptionsByUserId(id string) ([]model.DbSubscription, error) {
	log.Info("Getting all subscriptions for user %v", id)
	var subscriptions []model.DbSubscription
	err := GenericGetMultiple("subscriptions", "user_id", id, &subscriptions)
	return subscriptions, err
}

func CreateSubscription(sub model.DbSubscription) error {
	log.Info("Creating new subscription for user %v notebook %v", sub.UserId, sub.NotebookId)
	err := dbmap.Insert(&sub)
	if log.Error(err) {
		return errors.ISE()
	}
	return nil
}
