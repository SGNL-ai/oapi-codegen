package packageA

//go:generate go run github.com/sgnl-ai/oapi-codegen/cmd/oapi-codegen --old-config-style -generate types,skip-prune,spec --package=packageA -o externalref.gen.go --import-mapping=../packageB/spec.yaml:github.com/sgnl-ai/oapi-codegen/internal/test/externalref/packageB spec.yaml
