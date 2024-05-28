package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemOnenoteSectiongroupsItemSectionsCountRequestBuilder provides operations to count the resources in the collection.
type ItemOnenoteSectiongroupsItemSectionsCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOnenoteSectiongroupsItemSectionsCountRequestBuilderGetQueryParameters get the number of the resource
type ItemOnenoteSectiongroupsItemSectionsCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// ItemOnenoteSectiongroupsItemSectionsCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnenoteSectiongroupsItemSectionsCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOnenoteSectiongroupsItemSectionsCountRequestBuilderGetQueryParameters
}
// NewItemOnenoteSectiongroupsItemSectionsCountRequestBuilderInternal instantiates a new ItemOnenoteSectiongroupsItemSectionsCountRequestBuilder and sets the default values.
func NewItemOnenoteSectiongroupsItemSectionsCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteSectiongroupsItemSectionsCountRequestBuilder) {
    m := &ItemOnenoteSectiongroupsItemSectionsCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/onenote/sectionGroups/{sectionGroup%2Did}/sections/$count{?%24filter,%24search}", pathParameters),
    }
    return m
}
// NewItemOnenoteSectiongroupsItemSectionsCountRequestBuilder instantiates a new ItemOnenoteSectiongroupsItemSectionsCountRequestBuilder and sets the default values.
func NewItemOnenoteSectiongroupsItemSectionsCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteSectiongroupsItemSectionsCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnenoteSectiongroupsItemSectionsCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
// returns a *int32 when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnenoteSectiongroupsItemSectionsCountRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOnenoteSectiongroupsItemSectionsCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
// returns a *RequestInformation when successful
func (m *ItemOnenoteSectiongroupsItemSectionsCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOnenoteSectiongroupsItemSectionsCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemOnenoteSectiongroupsItemSectionsCountRequestBuilder when successful
func (m *ItemOnenoteSectiongroupsItemSectionsCountRequestBuilder) WithUrl(rawUrl string)(*ItemOnenoteSectiongroupsItemSectionsCountRequestBuilder) {
    return NewItemOnenoteSectiongroupsItemSectionsCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
