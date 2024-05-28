package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemOnenoteSectiongroupsItemSectionsItemPagesItemPreviewRequestBuilder provides operations to call the preview method.
type ItemOnenoteSectiongroupsItemSectionsItemPagesItemPreviewRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOnenoteSectiongroupsItemSectionsItemPagesItemPreviewRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnenoteSectiongroupsItemSectionsItemPagesItemPreviewRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemOnenoteSectiongroupsItemSectionsItemPagesItemPreviewRequestBuilderInternal instantiates a new ItemOnenoteSectiongroupsItemSectionsItemPagesItemPreviewRequestBuilder and sets the default values.
func NewItemOnenoteSectiongroupsItemSectionsItemPagesItemPreviewRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteSectiongroupsItemSectionsItemPagesItemPreviewRequestBuilder) {
    m := &ItemOnenoteSectiongroupsItemSectionsItemPagesItemPreviewRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/onenote/sectionGroups/{sectionGroup%2Did}/sections/{onenoteSection%2Did}/pages/{onenotePage%2Did}/preview()", pathParameters),
    }
    return m
}
// NewItemOnenoteSectiongroupsItemSectionsItemPagesItemPreviewRequestBuilder instantiates a new ItemOnenoteSectiongroupsItemSectionsItemPagesItemPreviewRequestBuilder and sets the default values.
func NewItemOnenoteSectiongroupsItemSectionsItemPagesItemPreviewRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteSectiongroupsItemSectionsItemPagesItemPreviewRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnenoteSectiongroupsItemSectionsItemPagesItemPreviewRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function preview
// returns a OnenotePagePreviewable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnenoteSectiongroupsItemSectionsItemPagesItemPreviewRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOnenoteSectiongroupsItemSectionsItemPagesItemPreviewRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePagePreviewable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOnenotePagePreviewFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePagePreviewable), nil
}
// ToGetRequestInformation invoke function preview
// returns a *RequestInformation when successful
func (m *ItemOnenoteSectiongroupsItemSectionsItemPagesItemPreviewRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOnenoteSectiongroupsItemSectionsItemPagesItemPreviewRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemOnenoteSectiongroupsItemSectionsItemPagesItemPreviewRequestBuilder when successful
func (m *ItemOnenoteSectiongroupsItemSectionsItemPagesItemPreviewRequestBuilder) WithUrl(rawUrl string)(*ItemOnenoteSectiongroupsItemSectionsItemPagesItemPreviewRequestBuilder) {
    return NewItemOnenoteSectiongroupsItemSectionsItemPagesItemPreviewRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
