package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemInferenceClassificationOverridesRequestBuilder provides operations to manage the overrides property of the microsoft.graph.inferenceClassification entity.
type ItemInferenceClassificationOverridesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInferenceClassificationOverridesRequestBuilderGetQueryParameters get the overrides that a user has set up to always classify messages from certain senders in specific ways. Each override corresponds to an SMTP address of a sender. Initially, a user does not have any overrides.
type ItemInferenceClassificationOverridesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemInferenceClassificationOverridesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInferenceClassificationOverridesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemInferenceClassificationOverridesRequestBuilderGetQueryParameters
}
// ItemInferenceClassificationOverridesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInferenceClassificationOverridesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByInferenceClassificationOverrideId provides operations to manage the overrides property of the microsoft.graph.inferenceClassification entity.
func (m *ItemInferenceClassificationOverridesRequestBuilder) ByInferenceClassificationOverrideId(inferenceClassificationOverrideId string)(*ItemInferenceClassificationOverridesInferenceClassificationOverrideItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if inferenceClassificationOverrideId != "" {
        urlTplParams["inferenceClassificationOverride%2Did"] = inferenceClassificationOverrideId
    }
    return NewItemInferenceClassificationOverridesInferenceClassificationOverrideItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemInferenceClassificationOverridesRequestBuilderInternal instantiates a new OverridesRequestBuilder and sets the default values.
func NewItemInferenceClassificationOverridesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInferenceClassificationOverridesRequestBuilder) {
    m := &ItemInferenceClassificationOverridesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/inferenceClassification/overrides{?%24top,%24skip,%24filter,%24count,%24orderby,%24select}", pathParameters),
    }
    return m
}
// NewItemInferenceClassificationOverridesRequestBuilder instantiates a new OverridesRequestBuilder and sets the default values.
func NewItemInferenceClassificationOverridesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInferenceClassificationOverridesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInferenceClassificationOverridesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemInferenceClassificationOverridesRequestBuilder) Count()(*ItemInferenceClassificationOverridesCountRequestBuilder) {
    return NewItemInferenceClassificationOverridesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the overrides that a user has set up to always classify messages from certain senders in specific ways. Each override corresponds to an SMTP address of a sender. Initially, a user does not have any overrides.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/inferenceclassification-list-overrides?view=graph-rest-1.0
func (m *ItemInferenceClassificationOverridesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemInferenceClassificationOverridesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.InferenceClassificationOverrideCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateInferenceClassificationOverrideCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.InferenceClassificationOverrideCollectionResponseable), nil
}
// Post create an override for a sender identified by an SMTP address. Future messages from that SMTP address will be consistently classifiedas specified in the override. Note
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/inferenceclassification-post-overrides?view=graph-rest-1.0
func (m *ItemInferenceClassificationOverridesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.InferenceClassificationOverrideable, requestConfiguration *ItemInferenceClassificationOverridesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.InferenceClassificationOverrideable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateInferenceClassificationOverrideFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.InferenceClassificationOverrideable), nil
}
// ToGetRequestInformation get the overrides that a user has set up to always classify messages from certain senders in specific ways. Each override corresponds to an SMTP address of a sender. Initially, a user does not have any overrides.
func (m *ItemInferenceClassificationOverridesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemInferenceClassificationOverridesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPostRequestInformation create an override for a sender identified by an SMTP address. Future messages from that SMTP address will be consistently classifiedas specified in the override. Note
func (m *ItemInferenceClassificationOverridesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.InferenceClassificationOverrideable, requestConfiguration *ItemInferenceClassificationOverridesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemInferenceClassificationOverridesRequestBuilder) WithUrl(rawUrl string)(*ItemInferenceClassificationOverridesRequestBuilder) {
    return NewItemInferenceClassificationOverridesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
