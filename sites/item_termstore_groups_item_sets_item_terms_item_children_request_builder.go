package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3 "github.com/microsoftgraph/msgraph-sdk-go/models/termstore"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemTermstoreGroupsItemSetsItemTermsItemChildrenRequestBuilder provides operations to manage the children property of the microsoft.graph.termStore.term entity.
type ItemTermstoreGroupsItemSetsItemTermsItemChildrenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTermstoreGroupsItemSetsItemTermsItemChildrenRequestBuilderGetQueryParameters children of current term.
type ItemTermstoreGroupsItemSetsItemTermsItemChildrenRequestBuilderGetQueryParameters struct {
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
// ItemTermstoreGroupsItemSetsItemTermsItemChildrenRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermstoreGroupsItemSetsItemTermsItemChildrenRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTermstoreGroupsItemSetsItemTermsItemChildrenRequestBuilderGetQueryParameters
}
// ItemTermstoreGroupsItemSetsItemTermsItemChildrenRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermstoreGroupsItemSetsItemTermsItemChildrenRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByTermId1 provides operations to manage the children property of the microsoft.graph.termStore.term entity.
// returns a *ItemTermstoreGroupsItemSetsItemTermsItemChildrenTermItemRequestBuilder when successful
func (m *ItemTermstoreGroupsItemSetsItemTermsItemChildrenRequestBuilder) ByTermId1(termId1 string)(*ItemTermstoreGroupsItemSetsItemTermsItemChildrenTermItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if termId1 != "" {
        urlTplParams["term%2Did1"] = termId1
    }
    return NewItemTermstoreGroupsItemSetsItemTermsItemChildrenTermItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemTermstoreGroupsItemSetsItemTermsItemChildrenRequestBuilderInternal instantiates a new ItemTermstoreGroupsItemSetsItemTermsItemChildrenRequestBuilder and sets the default values.
func NewItemTermstoreGroupsItemSetsItemTermsItemChildrenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTermstoreGroupsItemSetsItemTermsItemChildrenRequestBuilder) {
    m := &ItemTermstoreGroupsItemSetsItemTermsItemChildrenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/termStore/groups/{group%2Did}/sets/{set%2Did}/terms/{term%2Did}/children{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemTermstoreGroupsItemSetsItemTermsItemChildrenRequestBuilder instantiates a new ItemTermstoreGroupsItemSetsItemTermsItemChildrenRequestBuilder and sets the default values.
func NewItemTermstoreGroupsItemSetsItemTermsItemChildrenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTermstoreGroupsItemSetsItemTermsItemChildrenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTermstoreGroupsItemSetsItemTermsItemChildrenRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemTermstoreGroupsItemSetsItemTermsItemChildrenCountRequestBuilder when successful
func (m *ItemTermstoreGroupsItemSetsItemTermsItemChildrenRequestBuilder) Count()(*ItemTermstoreGroupsItemSetsItemTermsItemChildrenCountRequestBuilder) {
    return NewItemTermstoreGroupsItemSetsItemTermsItemChildrenCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get children of current term.
// returns a TermCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTermstoreGroupsItemSetsItemTermsItemChildrenRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTermstoreGroupsItemSetsItemTermsItemChildrenRequestBuilderGetRequestConfiguration)(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.TermCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.CreateTermCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.TermCollectionResponseable), nil
}
// Post create new navigation property to children for sites
// returns a Termable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTermstoreGroupsItemSetsItemTermsItemChildrenRequestBuilder) Post(ctx context.Context, body ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable, requestConfiguration *ItemTermstoreGroupsItemSetsItemTermsItemChildrenRequestBuilderPostRequestConfiguration)(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.CreateTermFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable), nil
}
// ToGetRequestInformation children of current term.
// returns a *RequestInformation when successful
func (m *ItemTermstoreGroupsItemSetsItemTermsItemChildrenRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTermstoreGroupsItemSetsItemTermsItemChildrenRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to children for sites
// returns a *RequestInformation when successful
func (m *ItemTermstoreGroupsItemSetsItemTermsItemChildrenRequestBuilder) ToPostRequestInformation(ctx context.Context, body ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable, requestConfiguration *ItemTermstoreGroupsItemSetsItemTermsItemChildrenRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTermstoreGroupsItemSetsItemTermsItemChildrenRequestBuilder when successful
func (m *ItemTermstoreGroupsItemSetsItemTermsItemChildrenRequestBuilder) WithUrl(rawUrl string)(*ItemTermstoreGroupsItemSetsItemTermsItemChildrenRequestBuilder) {
    return NewItemTermstoreGroupsItemSetsItemTermsItemChildrenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
