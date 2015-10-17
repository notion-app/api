
// Every single error that any of my OWN functions define in Notion should return
// one of these error types. All of these types descend from echo's HttpError type,
// and are thus safe to return back to the client. But they also contain
// additional information about the error which will be logged when an error
// happens
package errors

import (
  "github.com/labstack/echo"
)

type BadRequest struct {
  echo.HTTPError
}

type Unauthorized struct {
  echo.HTTPError
}

type Forbidden struct {
  echo.HTTPError
}

type NotFound struct {
  echo.HTTPError
}

type ISE struct {
  echo.HTTPError
}
