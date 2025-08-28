module github.com/microsoftgraph/msgraph-sdk-go

go 1.23.0

toolchain go1.24.1

require (
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.18.2
	github.com/google/uuid v1.6.0
	github.com/microsoft/kiota-abstractions-go v1.9.3
	github.com/microsoft/kiota-serialization-form-go v1.1.2
	github.com/microsoft/kiota-serialization-json-go v1.1.2
	github.com/microsoft/kiota-serialization-multipart-go v1.1.2
	github.com/microsoft/kiota-serialization-text-go v1.1.2
	github.com/microsoftgraph/msgraph-sdk-go-core v1.3.2
	github.com/stretchr/testify v1.11.1
)

require (
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.11.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/microsoft/kiota-authentication-azure-go v1.3.0 // indirect
	github.com/microsoft/kiota-http-go v1.5.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/std-uritemplate/std-uritemplate/go/v2 v2.0.3 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/otel v1.35.0 // indirect
	go.opentelemetry.io/otel/metric v1.35.0 // indirect
	go.opentelemetry.io/otel/trace v1.35.0 // indirect
	golang.org/x/net v0.42.0 // indirect
	golang.org/x/text v0.27.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract (
	// release contains a unintentined breaking change in name of classes
	v1.45.0
	v1.44.0
// release contains a unintentined breaking change in name of classes
)
