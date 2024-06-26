package handlers

import (
	"github.com/frhnfrnk/blog-platform-microservices/api-gateway/internal/graphql/user"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func NewUserHandler(schema graphql.Schema) http.Handler {
	return handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
}

func UserHandler(resolver *user.Resolver) http.Handler {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    user.NewRootQuery(resolver),
		Mutation: user.NewRootMutation(resolver),
	})
	if err != nil {
		panic(err)
	}

	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ContextHandler(r.Context(), w, r)
	})
}
