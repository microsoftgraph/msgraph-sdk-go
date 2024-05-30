package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemParentnotebookParentNotebookRequestBuilder provides operations to manage the parentNotebook property of the microsoft.graph.onenotePage entity.
type ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemParentnotebookParentNotebookRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemParentnotebookParentNotebookRequestBuilderGetQueryParameters the notebook that contains the page.  Read-only.
type ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemParentnotebookParentNotebookRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemParentnotebookParentNotebookRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemParentnotebookParentNotebookRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemParentnotebookParentNotebookRequestBuilderGetQueryParameters
}
// NewItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemParentnotebookParentNotebookRequestBuilderInternal instantiates a new ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemParentnotebookParentNotebookRequestBuilder and sets the default values.
func NewItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemParentnotebookParentNotebookRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemParentnotebookParentNotebookRequestBuilder) {
    m := &ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemParentnotebookParentNotebookRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/onenote/notebooks/{notebook%2Did}/sectionGroups/{sectionGroup%2Did}/sections/{onenoteSection%2Did}/pages/{onenotePage%2Did}/parentNotebook{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemParentnotebookParentNotebookRequestBuilder instantiates a new ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemParentnotebookParentNotebookRequestBuilder and sets the default values.
func NewItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemParentnotebookParentNotebookRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemParentnotebookParentNotebookRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemParentnotebookParentNotebookRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the notebook that contains the page.  Read-only.
// returns a Notebookable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemParentnotebookParentNotebookRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemParentnotebookParentNotebookRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Notebookable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateNotebookFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Notebookable), nil
}
// ToGetRequestInformation the notebook that contains the page.  Read-only.
// returns a *RequestInformation when successful
func (m *ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemParentnotebookParentNotebookRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemParentnotebookParentNotebookRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemParentnotebookParentNotebookRequestBuilder when successful
func (m *ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemParentnotebookParentNotebookRequestBuilder) WithUrl(rawUrl string)(*ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemParentnotebookParentNotebookRequestBuilder) {
    return NewItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemParentnotebookParentNotebookRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
