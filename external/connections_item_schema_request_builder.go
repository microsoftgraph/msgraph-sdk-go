package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc "github.com/microsoftgraph/msgraph-sdk-go/models/externalconnectors"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ConnectionsItemSchemaRequestBuilder provides operations to manage the schema property of the microsoft.graph.externalConnectors.externalConnection entity.
type ConnectionsItemSchemaRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConnectionsItemSchemaRequestBuilderGetQueryParameters read the properties and relationships of a schema object. This API is available in the following national cloud deployments.
type ConnectionsItemSchemaRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ConnectionsItemSchemaRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectionsItemSchemaRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConnectionsItemSchemaRequestBuilderGetQueryParameters
}
// ConnectionsItemSchemaRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectionsItemSchemaRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewConnectionsItemSchemaRequestBuilderInternal instantiates a new SchemaRequestBuilder and sets the default values.
func NewConnectionsItemSchemaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectionsItemSchemaRequestBuilder) {
    m := &ConnectionsItemSchemaRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/connections/{externalConnection%2Did}/schema{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewConnectionsItemSchemaRequestBuilder instantiates a new SchemaRequestBuilder and sets the default values.
func NewConnectionsItemSchemaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectionsItemSchemaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConnectionsItemSchemaRequestBuilderInternal(urlParams, requestAdapter)
}
// Get read the properties and relationships of a schema object. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/externalconnectors-schema-get?view=graph-rest-1.0
func (m *ConnectionsItemSchemaRequestBuilder) Get(ctx context.Context, requestConfiguration *ConnectionsItemSchemaRequestBuilderGetRequestConfiguration)(i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.Schemaable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.CreateSchemaFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.Schemaable), nil
}
// Patch create a new schema object. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/externalconnectors-externalconnection-patch-schema?view=graph-rest-1.0
func (m *ConnectionsItemSchemaRequestBuilder) Patch(ctx context.Context, body i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.Schemaable, requestConfiguration *ConnectionsItemSchemaRequestBuilderPatchRequestConfiguration)(i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.Schemaable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.CreateSchemaFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.Schemaable), nil
}
// ToGetRequestInformation read the properties and relationships of a schema object. This API is available in the following national cloud deployments.
func (m *ConnectionsItemSchemaRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConnectionsItemSchemaRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// ToPatchRequestInformation create a new schema object. This API is available in the following national cloud deployments.
func (m *ConnectionsItemSchemaRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.Schemaable, requestConfiguration *ConnectionsItemSchemaRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ConnectionsItemSchemaRequestBuilder) WithUrl(rawUrl string)(*ConnectionsItemSchemaRequestBuilder) {
    return NewConnectionsItemSchemaRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
