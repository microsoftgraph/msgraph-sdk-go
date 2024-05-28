package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenRequestBuilder provides operations to call the delta method.
type FilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenRequestBuilderGetQueryParameters invoke function delta
type FilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// FilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenRequestBuilderGetQueryParameters
}
// NewFilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenRequestBuilderInternal instantiates a new FilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, token *string)(*FilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenRequestBuilder) {
    m := &FilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/delta(token='{token}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if token != nil {
        m.BaseRequestBuilder.PathParameters["token"] = *token
    }
    return m
}
// NewFilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenRequestBuilder instantiates a new FilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function delta
// Deprecated: This method is obsolete. Use GetAsDeltaWithTokenGetResponse instead.
// returns a FilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenRequestBuilderGetRequestConfiguration)(FilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenResponseable), nil
}
// GetAsDeltaWithTokenGetResponse invoke function delta
// returns a FilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenRequestBuilder) GetAsDeltaWithTokenGetResponse(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenRequestBuilderGetRequestConfiguration)(FilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenGetResponseable), nil
}
// ToGetRequestInformation invoke function delta
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemDeltawithtokenDeltaWithTokenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
