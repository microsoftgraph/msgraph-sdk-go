package shares

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemListContenttypesContentTypesRequestBuilder provides operations to manage the contentTypes property of the microsoft.graph.list entity.
type ItemListContenttypesContentTypesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemListContenttypesContentTypesRequestBuilderGetQueryParameters the collection of content types present in this list.
type ItemListContenttypesContentTypesRequestBuilderGetQueryParameters struct {
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
// ItemListContenttypesContentTypesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListContenttypesContentTypesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemListContenttypesContentTypesRequestBuilderGetQueryParameters
}
// ItemListContenttypesContentTypesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListContenttypesContentTypesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AddCopy provides operations to call the addCopy method.
// returns a *ItemListContenttypesAddcopyAddCopyRequestBuilder when successful
func (m *ItemListContenttypesContentTypesRequestBuilder) AddCopy()(*ItemListContenttypesAddcopyAddCopyRequestBuilder) {
    return NewItemListContenttypesAddcopyAddCopyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AddCopyFromContentTypeHub provides operations to call the addCopyFromContentTypeHub method.
// returns a *ItemListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilder when successful
func (m *ItemListContenttypesContentTypesRequestBuilder) AddCopyFromContentTypeHub()(*ItemListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilder) {
    return NewItemListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ByContentTypeId provides operations to manage the contentTypes property of the microsoft.graph.list entity.
// returns a *ItemListContenttypesContentTypeItemRequestBuilder when successful
func (m *ItemListContenttypesContentTypesRequestBuilder) ByContentTypeId(contentTypeId string)(*ItemListContenttypesContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if contentTypeId != "" {
        urlTplParams["contentType%2Did"] = contentTypeId
    }
    return NewItemListContenttypesContentTypeItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemListContenttypesContentTypesRequestBuilderInternal instantiates a new ItemListContenttypesContentTypesRequestBuilder and sets the default values.
func NewItemListContenttypesContentTypesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListContenttypesContentTypesRequestBuilder) {
    m := &ItemListContenttypesContentTypesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/shares/{sharedDriveItem%2Did}/list/contentTypes{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemListContenttypesContentTypesRequestBuilder instantiates a new ItemListContenttypesContentTypesRequestBuilder and sets the default values.
func NewItemListContenttypesContentTypesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListContenttypesContentTypesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemListContenttypesContentTypesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemListContenttypesCountRequestBuilder when successful
func (m *ItemListContenttypesContentTypesRequestBuilder) Count()(*ItemListContenttypesCountRequestBuilder) {
    return NewItemListContenttypesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the collection of content types present in this list.
// returns a ContentTypeCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemListContenttypesContentTypesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemListContenttypesContentTypesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateContentTypeCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeCollectionResponseable), nil
}
// GetCompatibleHubContentTypes provides operations to call the getCompatibleHubContentTypes method.
// returns a *ItemListContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder when successful
func (m *ItemListContenttypesContentTypesRequestBuilder) GetCompatibleHubContentTypes()(*ItemListContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder) {
    return NewItemListContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create new navigation property to contentTypes for shares
// returns a ContentTypeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemListContenttypesContentTypesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, requestConfiguration *ItemListContenttypesContentTypesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateContentTypeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable), nil
}
// ToGetRequestInformation the collection of content types present in this list.
// returns a *RequestInformation when successful
func (m *ItemListContenttypesContentTypesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemListContenttypesContentTypesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to contentTypes for shares
// returns a *RequestInformation when successful
func (m *ItemListContenttypesContentTypesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, requestConfiguration *ItemListContenttypesContentTypesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemListContenttypesContentTypesRequestBuilder when successful
func (m *ItemListContenttypesContentTypesRequestBuilder) WithUrl(rawUrl string)(*ItemListContenttypesContentTypesRequestBuilder) {
    return NewItemListContenttypesContentTypesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
