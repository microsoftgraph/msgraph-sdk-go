package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilder provides operations to manage the authenticationContextClassReferences property of the microsoft.graph.conditionalAccessRoot entity.
type ConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilderGetQueryParameters retrieve the properties and relationships of a authenticationContextClassReference object.
type ConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilderGetQueryParameters
}
// ConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilderInternal instantiates a new ConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilder and sets the default values.
func NewConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilder) {
    m := &ConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/conditionalAccess/authenticationContextClassReferences/{authenticationContextClassReference%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilder instantiates a new ConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilder and sets the default values.
func NewConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete an authenticationContextClassReference object that's not published or used by a conditional access policy.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationcontextclassreference-delete?view=graph-rest-1.0
func (m *ConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get retrieve the properties and relationships of a authenticationContextClassReference object.
// returns a AuthenticationContextClassReferenceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationcontextclassreference-get?view=graph-rest-1.0
func (m *ConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationContextClassReferenceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuthenticationContextClassReferenceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationContextClassReferenceable), nil
}
// Patch create an authenticationContextClassReference object, if the ID has not been used. If ID has been used, this call updates the authenticationContextClassReference object.
// returns a AuthenticationContextClassReferenceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationcontextclassreference-update?view=graph-rest-1.0
func (m *ConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationContextClassReferenceable, requestConfiguration *ConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationContextClassReferenceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuthenticationContextClassReferenceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationContextClassReferenceable), nil
}
// ToDeleteRequestInformation delete an authenticationContextClassReference object that's not published or used by a conditional access policy.
// returns a *RequestInformation when successful
func (m *ConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve the properties and relationships of a authenticationContextClassReference object.
// returns a *RequestInformation when successful
func (m *ConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation create an authenticationContextClassReference object, if the ID has not been used. If ID has been used, this call updates the authenticationContextClassReference object.
// returns a *RequestInformation when successful
func (m *ConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationContextClassReferenceable, requestConfiguration *ConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilder when successful
func (m *ConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilder) WithUrl(rawUrl string)(*ConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilder) {
    return NewConditionalaccessAuthenticationcontextclassreferencesAuthenticationContextClassReferenceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
