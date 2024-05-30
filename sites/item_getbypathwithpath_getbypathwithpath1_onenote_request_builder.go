package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilder provides operations to manage the onenote property of the microsoft.graph.site entity.
type ItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilderGetQueryParameters calls the OneNote service for notebook related operations.
type ItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilderGetQueryParameters
}
// ItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilderInternal instantiates a new ItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilder and sets the default values.
func NewItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilder) {
    m := &ItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/getByPath(path='{path}')/getByPath(path='{path1}')/onenote{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilder instantiates a new ItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilder and sets the default values.
func NewItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property onenote for sites
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get calls the OneNote service for notebook related operations.
// returns a Onenoteable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Onenoteable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOnenoteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Onenoteable), nil
}
// Patch update the navigation property onenote in sites
// returns a Onenoteable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Onenoteable, requestConfiguration *ItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Onenoteable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOnenoteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Onenoteable), nil
}
// ToDeleteRequestInformation delete navigation property onenote for sites
// returns a *RequestInformation when successful
func (m *ItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation calls the OneNote service for notebook related operations.
// returns a *RequestInformation when successful
func (m *ItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property onenote in sites
// returns a *RequestInformation when successful
func (m *ItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Onenoteable, requestConfiguration *ItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilder when successful
func (m *ItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilder) WithUrl(rawUrl string)(*ItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilder) {
    return NewItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
