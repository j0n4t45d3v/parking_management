package middlerware

import "net/http"

func ContentTypeMiddlerware(next http.Handler) http.Handler {
  handler := http.HandlerFunc(func (res http.ResponseWriter, req *http.Request)  {
    res.Header().Set("Content-type", "application/json")
    next.ServeHTTP(res, req)
  }) 

  return handler
}

