//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	log.Println("Generating entutils/entgql schema files...")
	ex, err := entgql.NewExtension(
		entgql.WithConfigPath("gqlgen.yml"),
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("./ent.graphql"),
		entgql.WithWhereInputs(true),
		entgql.WithNodeDescriptor(true),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	opts := []entc.Option{
		entc.Extensions(ex),
	}

	if err := entc.Generate("./ent/schema", &gen.Config{
		Templates: []*gen.Template{},
		Header: `
			// OHMAB
			// Code generated by entc, DO NOT EDIT.
		`}, opts...); err != nil {
		log.Fatalf("running entutils codegen: %v", err)
	}
}
