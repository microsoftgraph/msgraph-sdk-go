module github.com/microsoftgraph/msgraph-sdk-go

go 1.18

require (
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.14.0
	github.com/google/uuid v1.6.0
	github.com/microsoft/kiota-abstractions-go v1.6.1
	github.com/microsoft/kiota-serialization-form-go v1.0.0
	github.com/microsoft/kiota-serialization-json-go v1.0.8
	github.com/microsoft/kiota-serialization-multipart-go v1.0.0
	github.com/microsoft/kiota-serialization-text-go v1.0.0
	github.com/microsoftgraph/msgraph-sdk-go-core v1.2.0
	github.com/stretchr/testify v1.9.0
)

require (
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.10.0 // indirect
	github.com/cjlapao/common-go v0.0.39 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/microsoft/kiota-authentication-azure-go v1.0.2 // indirect
	github.com/microsoft/kiota-http-go v1.4.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/std-uritemplate/std-uritemplate/go v0.0.57 // indirect
	go.opentelemetry.io/otel v1.24.0 // indirect
	go.opentelemetry.io/otel/metric v1.24.0 // indirect
	go.opentelemetry.io/otel/trace v1.24.0 // indirect
	golang.org/x/net v0.27.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract (
	// release contains a unintentined breaking change in name of classes
	v1.45.0
	v1.44.0
// release contains a unintentined breaking change in name of classes
)
