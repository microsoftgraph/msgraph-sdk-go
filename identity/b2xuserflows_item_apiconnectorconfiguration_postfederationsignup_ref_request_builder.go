package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// B2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilder provides operations to manage the collection of identityContainer entities.
type B2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// B2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// B2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// B2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewB2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilderInternal instantiates a new B2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilder and sets the default values.
func NewB2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilder) {
    m := &B2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow%2Did}/apiConnectorConfiguration/postFederationSignup/$ref", pathParameters),
    }
    return m
}
// NewB2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilder instantiates a new B2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilder and sets the default values.
func NewB2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete ref of navigation property postFederationSignup for identity
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *B2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *B2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get ref of postFederationSignup from identity
// returns a *string when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *B2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilder) Get(ctx context.Context, requestConfiguration *B2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilderGetRequestConfiguration)(*string, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "string", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*string), nil
}
// Put update the ref of navigation property postFederationSignup in identity
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *B2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilder) Put(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceUpdateable, requestConfiguration *B2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilderPutRequestConfiguration)(error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete ref of navigation property postFederationSignup for identity
// returns a *RequestInformation when successful
func (m *B2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *B2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get ref of postFederationSignup from identity
// returns a *RequestInformation when successful
func (m *B2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *B2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPutRequestInformation update the ref of navigation property postFederationSignup in identity
// returns a *RequestInformation when successful
func (m *B2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilder) ToPutRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceUpdateable, requestConfiguration *B2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *B2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilder when successful
func (m *B2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilder) WithUrl(rawUrl string)(*B2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilder) {
    return NewB2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
