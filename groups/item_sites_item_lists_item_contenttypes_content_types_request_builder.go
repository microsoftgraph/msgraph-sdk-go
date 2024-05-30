package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemListsItemContenttypesContentTypesRequestBuilder provides operations to manage the contentTypes property of the microsoft.graph.list entity.
type ItemSitesItemListsItemContenttypesContentTypesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemListsItemContenttypesContentTypesRequestBuilderGetQueryParameters the collection of content types present in this list.
type ItemSitesItemListsItemContenttypesContentTypesRequestBuilderGetQueryParameters struct {
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
// ItemSitesItemListsItemContenttypesContentTypesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemListsItemContenttypesContentTypesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemListsItemContenttypesContentTypesRequestBuilderGetQueryParameters
}
// ItemSitesItemListsItemContenttypesContentTypesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemListsItemContenttypesContentTypesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AddCopy provides operations to call the addCopy method.
// returns a *ItemSitesItemListsItemContenttypesAddcopyAddCopyRequestBuilder when successful
func (m *ItemSitesItemListsItemContenttypesContentTypesRequestBuilder) AddCopy()(*ItemSitesItemListsItemContenttypesAddcopyAddCopyRequestBuilder) {
    return NewItemSitesItemListsItemContenttypesAddcopyAddCopyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AddCopyFromContentTypeHub provides operations to call the addCopyFromContentTypeHub method.
// returns a *ItemSitesItemListsItemContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilder when successful
func (m *ItemSitesItemListsItemContenttypesContentTypesRequestBuilder) AddCopyFromContentTypeHub()(*ItemSitesItemListsItemContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilder) {
    return NewItemSitesItemListsItemContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ByContentTypeId provides operations to manage the contentTypes property of the microsoft.graph.list entity.
// returns a *ItemSitesItemListsItemContenttypesContentTypeItemRequestBuilder when successful
func (m *ItemSitesItemListsItemContenttypesContentTypesRequestBuilder) ByContentTypeId(contentTypeId string)(*ItemSitesItemListsItemContenttypesContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if contentTypeId != "" {
        urlTplParams["contentType%2Did"] = contentTypeId
    }
    return NewItemSitesItemListsItemContenttypesContentTypeItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemSitesItemListsItemContenttypesContentTypesRequestBuilderInternal instantiates a new ItemSitesItemListsItemContenttypesContentTypesRequestBuilder and sets the default values.
func NewItemSitesItemListsItemContenttypesContentTypesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemContenttypesContentTypesRequestBuilder) {
    m := &ItemSitesItemListsItemContenttypesContentTypesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/lists/{list%2Did}/contentTypes{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemSitesItemListsItemContenttypesContentTypesRequestBuilder instantiates a new ItemSitesItemListsItemContenttypesContentTypesRequestBuilder and sets the default values.
func NewItemSitesItemListsItemContenttypesContentTypesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemContenttypesContentTypesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemListsItemContenttypesContentTypesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemSitesItemListsItemContenttypesCountRequestBuilder when successful
func (m *ItemSitesItemListsItemContenttypesContentTypesRequestBuilder) Count()(*ItemSitesItemListsItemContenttypesCountRequestBuilder) {
    return NewItemSitesItemListsItemContenttypesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the collection of content types present in this list.
// returns a ContentTypeCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemListsItemContenttypesContentTypesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemListsItemContenttypesContentTypesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeCollectionResponseable, error) {
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
// returns a *ItemSitesItemListsItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder when successful
func (m *ItemSitesItemListsItemContenttypesContentTypesRequestBuilder) GetCompatibleHubContentTypes()(*ItemSitesItemListsItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder) {
    return NewItemSitesItemListsItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create new navigation property to contentTypes for groups
// returns a ContentTypeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemListsItemContenttypesContentTypesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, requestConfiguration *ItemSitesItemListsItemContenttypesContentTypesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, error) {
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
func (m *ItemSitesItemListsItemContenttypesContentTypesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemListsItemContenttypesContentTypesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to contentTypes for groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemListsItemContenttypesContentTypesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, requestConfiguration *ItemSitesItemListsItemContenttypesContentTypesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemListsItemContenttypesContentTypesRequestBuilder when successful
func (m *ItemSitesItemListsItemContenttypesContentTypesRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemListsItemContenttypesContentTypesRequestBuilder) {
    return NewItemSitesItemListsItemContenttypesContentTypesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
