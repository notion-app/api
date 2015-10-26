
package db

import (
  "notion/model"
)

func GetUserSubscriptions(userId string) ([]model.DbSubscription, error) {
  subs := make([]model.DbSubscription, 0)
  err := GenericGet(&subs, `select * from subscriptions where user_id=$1`, userId)
  return subs, err
}

func CreateSubscription(s model.DbSubscription) error {
  return dbmap.Insert(&s)
}

func UpdateSubscription(s model.DbSubscription) error {
  _, err := dbmap.Update(&s)
  return err
}

func DeleteSubscription(s model.DbSubscription) error {
  _, err := dbmap.Delete(&s)
  return err
}
