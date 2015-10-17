package validate

import (
  "notion/errors"
  "notion/model"
)

func AddSubscriptionRequest(m model.AddSubscriptionRequest) error {
  if m.NotebookId == "" {
    return errors.BadRequest("Must provide notebook id")
  }
  return nil
}
