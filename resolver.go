//go:generate gorunpkg github.com/99designs/gqlgen

package gqlgen_dataloader

import (
	"github.com/aneri/gqlgen-authentication/models"
	"github.com/aneri/gqlgen-example/dal"
)

var context context.Context

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) NewApplication(ctx context.Context, input NewApplication) (Application, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteApplication(ctx context.Context, id string) (string, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateApplication(ctx context.Context, id string, input UpdateApplication) (Application, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Applications(ctx context.Context) ([]models.Application, error) {
	crConn := context.Value("crConn").(*dal.DbConnection)
	var application []models.Application
	row := crConn.Db.Where("id=?").Find(&application)
	defer row.Close()
	return application, nil
}
