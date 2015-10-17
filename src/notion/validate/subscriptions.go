package validate

import (
  "notion/errors"
  "notion/model"
)

func AddSubscriptionRequest(m model.SubscriptionRequest) error {
  if m.NotebookId == "" {
    return errors.BadRequest("Must provide notebook id")
  }
  return nil
}

func RemoveSubscriptionRequest(m model.SubscriptionRequest) error {
  return AddSubscriptionRequest(m)
}
