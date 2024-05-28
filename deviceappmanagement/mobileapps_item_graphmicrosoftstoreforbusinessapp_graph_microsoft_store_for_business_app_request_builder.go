package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder casts the previous resource to microsoftStoreForBusinessApp.
type MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilderGetQueryParameters get the item of type microsoft.graph.mobileApp as microsoft.graph.microsoftStoreForBusinessApp
type MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilderGetQueryParameters
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.mobileApp entity.
// returns a *MobileappsItemGraphmicrosoftstoreforbusinessappAssignmentsRequestBuilder when successful
func (m *MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder) Assignments()(*MobileappsItemGraphmicrosoftstoreforbusinessappAssignmentsRequestBuilder) {
    return NewMobileappsItemGraphmicrosoftstoreforbusinessappAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Categories provides operations to manage the categories property of the microsoft.graph.mobileApp entity.
// returns a *MobileappsItemGraphmicrosoftstoreforbusinessappCategoriesRequestBuilder when successful
func (m *MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder) Categories()(*MobileappsItemGraphmicrosoftstoreforbusinessappCategoriesRequestBuilder) {
    return NewMobileappsItemGraphmicrosoftstoreforbusinessappCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilderInternal instantiates a new MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder and sets the default values.
func NewMobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder) {
    m := &MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.microsoftStoreForBusinessApp{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder instantiates a new MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder and sets the default values.
func NewMobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the item of type microsoft.graph.mobileApp as microsoft.graph.microsoftStoreForBusinessApp
// returns a MicrosoftStoreForBusinessAppable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MicrosoftStoreForBusinessAppable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMicrosoftStoreForBusinessAppFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MicrosoftStoreForBusinessAppable), nil
}
// ToGetRequestInformation get the item of type microsoft.graph.mobileApp as microsoft.graph.microsoftStoreForBusinessApp
// returns a *RequestInformation when successful
func (m *MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder when successful
func (m *MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder) WithUrl(rawUrl string)(*MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder) {
    return NewMobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
