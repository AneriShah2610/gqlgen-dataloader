//go:generate gorunpkg github.com/aneri/gqlgen-dataloader -ids int github.com/aneri/gqlgen-dataloader.Application
package gqlgen_dataloader

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/aneri/gqlgen-authentication/models"
	"github.com/aneri/gqlgen-example/dal"
)

type ctxKetType struct {
	name string
}

var context context.Context
var ctxKey = ctxKetType{"applicationCtx"}

type loaders struct {
	applicationByID *ApplicationLoader
}

// MiddleWareHandler to handle db connection
func MiddleWareHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		crConn, err := dal.Connect()
		if err != nil {
			log.Fatal(err)
		}
		context = context.WithValue(request.Context(), "crConn", crConn)
		next.ServeHTTP(writer, request.WithContext(context))
	})
}

// ApplicationLoaderMiddleware for datloader
func ApplicationLoaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ldrs := loaders{}
		wait := 250 * time.Millisecond
		ldrs.applicationByID = &ApplicationLoader{
			wait:     wait,
			maxBatch: 100,
			fetch: func(ids []int) ([]*Application, []error) {
				var keySql []string
				for _, id := range ids {
					keySql = append(keySql, strconv.Itoa(id))
				}
				var application []Application
				crConn := context.Value("crConn").(*dal.DbConnection)
				rows := crConn.Db.Where("id IN (?)", strings.Join(keySql, ",")).Find(&application)
				defer rows.Close()
				time.Sleep(5 * time.Millisecond)

				applications := make([]*models.Application, len(ids))
				errors := make([]error, len(ids))
				for i, id := range ids {
					applications[i] = &models.Application{}
				}
				return applications, nil
			},
		}
		ctx := context.WithValue(request.Context(), ctxKey, ldrs)
		next.ServeHTTP(writer, request.WithContext(ctx))
	})
}
