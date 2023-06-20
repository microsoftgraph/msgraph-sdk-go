package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSynchronizationSecretsRequestBuilder builds and executes requests for operations under \servicePrincipals\{servicePrincipal-id}\synchronization\secrets
type ItemSynchronizationSecretsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSynchronizationSecretsRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSynchronizationSecretsRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSynchronizationSecretsRequestBuilderInternal instantiates a new SecretsRequestBuilder and sets the default values.
func NewItemSynchronizationSecretsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationSecretsRequestBuilder) {
    m := &ItemSynchronizationSecretsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/synchronization/secrets", pathParameters),
    }
    return m
}
// NewItemSynchronizationSecretsRequestBuilder instantiates a new SecretsRequestBuilder and sets the default values.
func NewItemSynchronizationSecretsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationSecretsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSynchronizationSecretsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemSynchronizationSecretsRequestBuilder) Count()(*ItemSynchronizationSecretsCountRequestBuilder) {
    return NewItemSynchronizationSecretsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Put update property secrets value.
func (m *ItemSynchronizationSecretsRequestBuilder) Put(ctx context.Context, body []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SynchronizationSecretKeyStringValuePairable, requestConfiguration *ItemSynchronizationSecretsRequestBuilderPutRequestConfiguration)([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SynchronizationSecretKeyStringValuePairable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSynchronizationSecretKeyStringValuePairFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SynchronizationSecretKeyStringValuePairable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SynchronizationSecretKeyStringValuePairable)
        }
    }
    return val, nil
}
// ToPutRequestInformation update property secrets value.
func (m *ItemSynchronizationSecretsRequestBuilder) ToPutRequestInformation(ctx context.Context, body []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SynchronizationSecretKeyStringValuePairable, requestConfiguration *ItemSynchronizationSecretsRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
    requestInfo.Headers.Add("Accept", "application/json")
    cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(body))
    for i, v := range body {
        if v != nil {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
    }
    err := requestInfo.SetContentFromParsableCollection(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", cast)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
