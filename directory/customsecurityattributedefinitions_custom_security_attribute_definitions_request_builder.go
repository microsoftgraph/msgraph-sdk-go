package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilder provides operations to manage the customSecurityAttributeDefinitions property of the microsoft.graph.directory entity.
type CustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilderGetQueryParameters get a list of the customSecurityAttributeDefinition objects and their properties.
type CustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilderGetQueryParameters struct {
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
// CustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilderGetQueryParameters
}
// CustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByCustomSecurityAttributeDefinitionId provides operations to manage the customSecurityAttributeDefinitions property of the microsoft.graph.directory entity.
// returns a *CustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionItemRequestBuilder when successful
func (m *CustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilder) ByCustomSecurityAttributeDefinitionId(customSecurityAttributeDefinitionId string)(*CustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if customSecurityAttributeDefinitionId != "" {
        urlTplParams["customSecurityAttributeDefinition%2Did"] = customSecurityAttributeDefinitionId
    }
    return NewCustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilderInternal instantiates a new CustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilder and sets the default values.
func NewCustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilder) {
    m := &CustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/customSecurityAttributeDefinitions{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilder instantiates a new CustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilder and sets the default values.
func NewCustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CustomsecurityattributedefinitionsCountRequestBuilder when successful
func (m *CustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilder) Count()(*CustomsecurityattributedefinitionsCountRequestBuilder) {
    return NewCustomsecurityattributedefinitionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the customSecurityAttributeDefinition objects and their properties.
// returns a CustomSecurityAttributeDefinitionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/directory-list-customsecurityattributedefinitions?view=graph-rest-1.0
func (m *CustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilder) Get(ctx context.Context, requestConfiguration *CustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CustomSecurityAttributeDefinitionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCustomSecurityAttributeDefinitionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CustomSecurityAttributeDefinitionCollectionResponseable), nil
}
// Post create a new customSecurityAttributeDefinition object.
// returns a CustomSecurityAttributeDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/directory-post-customsecurityattributedefinitions?view=graph-rest-1.0
func (m *CustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CustomSecurityAttributeDefinitionable, requestConfiguration *CustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CustomSecurityAttributeDefinitionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCustomSecurityAttributeDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CustomSecurityAttributeDefinitionable), nil
}
// ToGetRequestInformation get a list of the customSecurityAttributeDefinition objects and their properties.
// returns a *RequestInformation when successful
func (m *CustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new customSecurityAttributeDefinition object.
// returns a *RequestInformation when successful
func (m *CustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CustomSecurityAttributeDefinitionable, requestConfiguration *CustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilder when successful
func (m *CustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilder) WithUrl(rawUrl string)(*CustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilder) {
    return NewCustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
