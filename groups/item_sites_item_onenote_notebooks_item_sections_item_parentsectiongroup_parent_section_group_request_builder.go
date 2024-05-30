package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemOnenoteNotebooksItemSectionsItemParentsectiongroupParentSectionGroupRequestBuilder provides operations to manage the parentSectionGroup property of the microsoft.graph.onenoteSection entity.
type ItemSitesItemOnenoteNotebooksItemSectionsItemParentsectiongroupParentSectionGroupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemOnenoteNotebooksItemSectionsItemParentsectiongroupParentSectionGroupRequestBuilderGetQueryParameters the section group that contains the section.  Read-only.
type ItemSitesItemOnenoteNotebooksItemSectionsItemParentsectiongroupParentSectionGroupRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemOnenoteNotebooksItemSectionsItemParentsectiongroupParentSectionGroupRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemOnenoteNotebooksItemSectionsItemParentsectiongroupParentSectionGroupRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemOnenoteNotebooksItemSectionsItemParentsectiongroupParentSectionGroupRequestBuilderGetQueryParameters
}
// NewItemSitesItemOnenoteNotebooksItemSectionsItemParentsectiongroupParentSectionGroupRequestBuilderInternal instantiates a new ItemSitesItemOnenoteNotebooksItemSectionsItemParentsectiongroupParentSectionGroupRequestBuilder and sets the default values.
func NewItemSitesItemOnenoteNotebooksItemSectionsItemParentsectiongroupParentSectionGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemOnenoteNotebooksItemSectionsItemParentsectiongroupParentSectionGroupRequestBuilder) {
    m := &ItemSitesItemOnenoteNotebooksItemSectionsItemParentsectiongroupParentSectionGroupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/onenote/notebooks/{notebook%2Did}/sections/{onenoteSection%2Did}/parentSectionGroup{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSitesItemOnenoteNotebooksItemSectionsItemParentsectiongroupParentSectionGroupRequestBuilder instantiates a new ItemSitesItemOnenoteNotebooksItemSectionsItemParentsectiongroupParentSectionGroupRequestBuilder and sets the default values.
func NewItemSitesItemOnenoteNotebooksItemSectionsItemParentsectiongroupParentSectionGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemOnenoteNotebooksItemSectionsItemParentsectiongroupParentSectionGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemOnenoteNotebooksItemSectionsItemParentsectiongroupParentSectionGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the section group that contains the section.  Read-only.
// returns a SectionGroupable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemOnenoteNotebooksItemSectionsItemParentsectiongroupParentSectionGroupRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteNotebooksItemSectionsItemParentsectiongroupParentSectionGroupRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SectionGroupable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSectionGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SectionGroupable), nil
}
// ToGetRequestInformation the section group that contains the section.  Read-only.
// returns a *RequestInformation when successful
func (m *ItemSitesItemOnenoteNotebooksItemSectionsItemParentsectiongroupParentSectionGroupRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteNotebooksItemSectionsItemParentsectiongroupParentSectionGroupRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemOnenoteNotebooksItemSectionsItemParentsectiongroupParentSectionGroupRequestBuilder when successful
func (m *ItemSitesItemOnenoteNotebooksItemSectionsItemParentsectiongroupParentSectionGroupRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemOnenoteNotebooksItemSectionsItemParentsectiongroupParentSectionGroupRequestBuilder) {
    return NewItemSitesItemOnenoteNotebooksItemSectionsItemParentsectiongroupParentSectionGroupRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
