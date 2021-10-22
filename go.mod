module github.com/microsoftgraph/msgraph-sdk-go

replace github.com/microsoft/kiota/abstractions/go => ./core/kiota/abstractions/go

replace github.com/microsoft/kiota/serialization/go/json => ./core/kiota/serialization/go/json

replace github.com/microsoft/kiota/http/go/nethttp => ./core/kiota/http/go/nethttp

replace github.com/microsoft/kiota/authentication/go/azure => ./core/kiota/authentication/go/azure

replace github.com/microsoftgraph/msgraph-sdk-go-core => ./core

go 1.17

require (
	github.com/Azure/azure-sdk-for-go/sdk/azcore v0.19.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/internal v0.7.1 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/microsoft/kiota/abstractions/go v0.0.0-20211022102944-75e03c022f36 // indirect
	github.com/microsoft/kiota/authentication/go/azure v0.0.0-00010101000000-000000000000 // indirect
	github.com/microsoft/kiota/http/go/nethttp v0.0.0-00010101000000-000000000000 // indirect
	github.com/microsoft/kiota/serialization/go/json v0.0.0-20211022102944-75e03c022f36 // indirect
	github.com/microsoftgraph/msgraph-sdk-go-core v0.0.0-00010101000000-000000000000 // indirect
	github.com/yosida95/uritemplate/v3 v3.0.1 // indirect
	golang.org/x/net v0.0.0-20211011170408-caeb26a5c8c0 // indirect
	golang.org/x/text v0.3.7 // indirect
)
