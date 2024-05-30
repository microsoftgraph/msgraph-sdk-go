package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UserflowattributesUserFlowAttributesRequestBuilder provides operations to manage the userFlowAttributes property of the microsoft.graph.identityContainer entity.
type UserflowattributesUserFlowAttributesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserflowattributesUserFlowAttributesRequestBuilderGetQueryParameters retrieve a list of identityUserFlowAttribute objects.
type UserflowattributesUserFlowAttributesRequestBuilderGetQueryParameters struct {
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
// UserflowattributesUserFlowAttributesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserflowattributesUserFlowAttributesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserflowattributesUserFlowAttributesRequestBuilderGetQueryParameters
}
// UserflowattributesUserFlowAttributesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserflowattributesUserFlowAttributesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByIdentityUserFlowAttributeId provides operations to manage the userFlowAttributes property of the microsoft.graph.identityContainer entity.
// returns a *UserflowattributesIdentityUserFlowAttributeItemRequestBuilder when successful
func (m *UserflowattributesUserFlowAttributesRequestBuilder) ByIdentityUserFlowAttributeId(identityUserFlowAttributeId string)(*UserflowattributesIdentityUserFlowAttributeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if identityUserFlowAttributeId != "" {
        urlTplParams["identityUserFlowAttribute%2Did"] = identityUserFlowAttributeId
    }
    return NewUserflowattributesIdentityUserFlowAttributeItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserflowattributesUserFlowAttributesRequestBuilderInternal instantiates a new UserflowattributesUserFlowAttributesRequestBuilder and sets the default values.
func NewUserflowattributesUserFlowAttributesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserflowattributesUserFlowAttributesRequestBuilder) {
    m := &UserflowattributesUserFlowAttributesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/userFlowAttributes{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewUserflowattributesUserFlowAttributesRequestBuilder instantiates a new UserflowattributesUserFlowAttributesRequestBuilder and sets the default values.
func NewUserflowattributesUserFlowAttributesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserflowattributesUserFlowAttributesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserflowattributesUserFlowAttributesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *UserflowattributesCountRequestBuilder when successful
func (m *UserflowattributesUserFlowAttributesRequestBuilder) Count()(*UserflowattributesCountRequestBuilder) {
    return NewUserflowattributesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve a list of identityUserFlowAttribute objects.
// returns a IdentityUserFlowAttributeCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identityuserflowattribute-list?view=graph-rest-1.0
func (m *UserflowattributesUserFlowAttributesRequestBuilder) Get(ctx context.Context, requestConfiguration *UserflowattributesUserFlowAttributesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityUserFlowAttributeCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateIdentityUserFlowAttributeCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityUserFlowAttributeCollectionResponseable), nil
}
// Post create a new custom identityUserFlowAttribute object.
// returns a IdentityUserFlowAttributeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identityuserflowattribute-post?view=graph-rest-1.0
func (m *UserflowattributesUserFlowAttributesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityUserFlowAttributeable, requestConfiguration *UserflowattributesUserFlowAttributesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityUserFlowAttributeable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateIdentityUserFlowAttributeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityUserFlowAttributeable), nil
}
// ToGetRequestInformation retrieve a list of identityUserFlowAttribute objects.
// returns a *RequestInformation when successful
func (m *UserflowattributesUserFlowAttributesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserflowattributesUserFlowAttributesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new custom identityUserFlowAttribute object.
// returns a *RequestInformation when successful
func (m *UserflowattributesUserFlowAttributesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityUserFlowAttributeable, requestConfiguration *UserflowattributesUserFlowAttributesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserflowattributesUserFlowAttributesRequestBuilder when successful
func (m *UserflowattributesUserFlowAttributesRequestBuilder) WithUrl(rawUrl string)(*UserflowattributesUserFlowAttributesRequestBuilder) {
    return NewUserflowattributesUserFlowAttributesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
