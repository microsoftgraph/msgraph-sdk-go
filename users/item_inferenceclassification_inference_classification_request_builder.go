package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemInferenceclassificationInferenceClassificationRequestBuilder provides operations to manage the inferenceClassification property of the microsoft.graph.user entity.
type ItemInferenceclassificationInferenceClassificationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInferenceclassificationInferenceClassificationRequestBuilderGetQueryParameters relevance classification of the user's messages based on explicit designations that override inferred relevance or importance.
type ItemInferenceclassificationInferenceClassificationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemInferenceclassificationInferenceClassificationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInferenceclassificationInferenceClassificationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemInferenceclassificationInferenceClassificationRequestBuilderGetQueryParameters
}
// ItemInferenceclassificationInferenceClassificationRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInferenceclassificationInferenceClassificationRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemInferenceclassificationInferenceClassificationRequestBuilderInternal instantiates a new ItemInferenceclassificationInferenceClassificationRequestBuilder and sets the default values.
func NewItemInferenceclassificationInferenceClassificationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInferenceclassificationInferenceClassificationRequestBuilder) {
    m := &ItemInferenceclassificationInferenceClassificationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/inferenceClassification{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemInferenceclassificationInferenceClassificationRequestBuilder instantiates a new ItemInferenceclassificationInferenceClassificationRequestBuilder and sets the default values.
func NewItemInferenceclassificationInferenceClassificationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInferenceclassificationInferenceClassificationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInferenceclassificationInferenceClassificationRequestBuilderInternal(urlParams, requestAdapter)
}
// Get relevance classification of the user's messages based on explicit designations that override inferred relevance or importance.
// returns a InferenceClassificationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemInferenceclassificationInferenceClassificationRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemInferenceclassificationInferenceClassificationRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.InferenceClassificationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateInferenceClassificationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.InferenceClassificationable), nil
}
// Overrides provides operations to manage the overrides property of the microsoft.graph.inferenceClassification entity.
// returns a *ItemInferenceclassificationOverridesRequestBuilder when successful
func (m *ItemInferenceclassificationInferenceClassificationRequestBuilder) Overrides()(*ItemInferenceclassificationOverridesRequestBuilder) {
    return NewItemInferenceclassificationOverridesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property inferenceClassification in users
// returns a InferenceClassificationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemInferenceclassificationInferenceClassificationRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.InferenceClassificationable, requestConfiguration *ItemInferenceclassificationInferenceClassificationRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.InferenceClassificationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateInferenceClassificationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.InferenceClassificationable), nil
}
// ToGetRequestInformation relevance classification of the user's messages based on explicit designations that override inferred relevance or importance.
// returns a *RequestInformation when successful
func (m *ItemInferenceclassificationInferenceClassificationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemInferenceclassificationInferenceClassificationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property inferenceClassification in users
// returns a *RequestInformation when successful
func (m *ItemInferenceclassificationInferenceClassificationRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.InferenceClassificationable, requestConfiguration *ItemInferenceclassificationInferenceClassificationRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemInferenceclassificationInferenceClassificationRequestBuilder when successful
func (m *ItemInferenceclassificationInferenceClassificationRequestBuilder) WithUrl(rawUrl string)(*ItemInferenceclassificationInferenceClassificationRequestBuilder) {
    return NewItemInferenceclassificationInferenceClassificationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
