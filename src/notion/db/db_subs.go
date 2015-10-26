
package db

import (
  "notion/model"
)

func GetUserSubscriptions(userId string) ([]model.DbSubscription, error) {
  subs := make([]model.DbSubscription, 0)
  err := GenericGet(&subs, `select * from subscriptions where user_id=$1`, userId)
  return subs, err
}
