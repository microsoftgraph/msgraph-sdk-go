package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FeaturerolloutpoliciesItemAppliestoAppliesToRequestBuilder provides operations to manage the appliesTo property of the microsoft.graph.featureRolloutPolicy entity.
type FeaturerolloutpoliciesItemAppliestoAppliesToRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FeaturerolloutpoliciesItemAppliestoAppliesToRequestBuilderGetQueryParameters nullable. Specifies a list of directoryObject resources that feature is enabled for.
type FeaturerolloutpoliciesItemAppliestoAppliesToRequestBuilderGetQueryParameters struct {
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
// FeaturerolloutpoliciesItemAppliestoAppliesToRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FeaturerolloutpoliciesItemAppliestoAppliesToRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FeaturerolloutpoliciesItemAppliestoAppliesToRequestBuilderGetQueryParameters
}
// FeaturerolloutpoliciesItemAppliestoAppliesToRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FeaturerolloutpoliciesItemAppliestoAppliesToRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDirectoryObjectId gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.policies.featureRolloutPolicies.item.appliesTo.item collection
// returns a *FeaturerolloutpoliciesItemAppliestoDirectoryObjectItemRequestBuilder when successful
func (m *FeaturerolloutpoliciesItemAppliestoAppliesToRequestBuilder) ByDirectoryObjectId(directoryObjectId string)(*FeaturerolloutpoliciesItemAppliestoDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if directoryObjectId != "" {
        urlTplParams["directoryObject%2Did"] = directoryObjectId
    }
    return NewFeaturerolloutpoliciesItemAppliestoDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewFeaturerolloutpoliciesItemAppliestoAppliesToRequestBuilderInternal instantiates a new FeaturerolloutpoliciesItemAppliestoAppliesToRequestBuilder and sets the default values.
func NewFeaturerolloutpoliciesItemAppliestoAppliesToRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FeaturerolloutpoliciesItemAppliestoAppliesToRequestBuilder) {
    m := &FeaturerolloutpoliciesItemAppliestoAppliesToRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/featureRolloutPolicies/{featureRolloutPolicy%2Did}/appliesTo{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewFeaturerolloutpoliciesItemAppliestoAppliesToRequestBuilder instantiates a new FeaturerolloutpoliciesItemAppliestoAppliesToRequestBuilder and sets the default values.
func NewFeaturerolloutpoliciesItemAppliestoAppliesToRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FeaturerolloutpoliciesItemAppliestoAppliesToRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFeaturerolloutpoliciesItemAppliestoAppliesToRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *FeaturerolloutpoliciesItemAppliestoCountRequestBuilder when successful
func (m *FeaturerolloutpoliciesItemAppliestoAppliesToRequestBuilder) Count()(*FeaturerolloutpoliciesItemAppliestoCountRequestBuilder) {
    return NewFeaturerolloutpoliciesItemAppliestoCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get nullable. Specifies a list of directoryObject resources that feature is enabled for.
// returns a DirectoryObjectCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FeaturerolloutpoliciesItemAppliestoAppliesToRequestBuilder) Get(ctx context.Context, requestConfiguration *FeaturerolloutpoliciesItemAppliestoAppliesToRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable), nil
}
// Post add an appliesTo on a featureRolloutPolicy object to specify the directoryObject to which the featureRolloutPolicy should be applied.
// returns a DirectoryObjectable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/featurerolloutpolicy-post-appliesto?view=graph-rest-1.0
func (m *FeaturerolloutpoliciesItemAppliestoAppliesToRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, requestConfiguration *FeaturerolloutpoliciesItemAppliestoAppliesToRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable), nil
}
// Ref provides operations to manage the collection of policyRoot entities.
// returns a *FeaturerolloutpoliciesItemAppliestoRefRequestBuilder when successful
func (m *FeaturerolloutpoliciesItemAppliestoAppliesToRequestBuilder) Ref()(*FeaturerolloutpoliciesItemAppliestoRefRequestBuilder) {
    return NewFeaturerolloutpoliciesItemAppliestoRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation nullable. Specifies a list of directoryObject resources that feature is enabled for.
// returns a *RequestInformation when successful
func (m *FeaturerolloutpoliciesItemAppliestoAppliesToRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FeaturerolloutpoliciesItemAppliestoAppliesToRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation add an appliesTo on a featureRolloutPolicy object to specify the directoryObject to which the featureRolloutPolicy should be applied.
// returns a *RequestInformation when successful
func (m *FeaturerolloutpoliciesItemAppliestoAppliesToRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, requestConfiguration *FeaturerolloutpoliciesItemAppliestoAppliesToRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FeaturerolloutpoliciesItemAppliestoAppliesToRequestBuilder when successful
func (m *FeaturerolloutpoliciesItemAppliestoAppliesToRequestBuilder) WithUrl(rawUrl string)(*FeaturerolloutpoliciesItemAppliestoAppliesToRequestBuilder) {
    return NewFeaturerolloutpoliciesItemAppliestoAppliesToRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
