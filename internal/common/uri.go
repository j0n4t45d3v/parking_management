package common

import (
	"fmt"
	"net/http"
)


func BuildUriLocation(r http.Request, route string, id int) string {
  host := r.Host
  return fmt.Sprintf("%v/%v/%d", host, route, id)
}
