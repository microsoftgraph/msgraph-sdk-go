package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FeatureRolloutPoliciesItemAppliesToRequestBuilder provides operations to manage the appliesTo property of the microsoft.graph.featureRolloutPolicy entity.
type FeatureRolloutPoliciesItemAppliesToRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FeatureRolloutPoliciesItemAppliesToRequestBuilderGetQueryParameters nullable. Specifies a list of directoryObjects that feature is enabled for.
type FeatureRolloutPoliciesItemAppliesToRequestBuilderGetQueryParameters struct {
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
// FeatureRolloutPoliciesItemAppliesToRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FeatureRolloutPoliciesItemAppliesToRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FeatureRolloutPoliciesItemAppliesToRequestBuilderGetQueryParameters
}
// FeatureRolloutPoliciesItemAppliesToRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FeatureRolloutPoliciesItemAppliesToRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDirectoryObjectId gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.policies.featureRolloutPolicies.item.appliesTo.item collection
func (m *FeatureRolloutPoliciesItemAppliesToRequestBuilder) ByDirectoryObjectId(directoryObjectId string)(*FeatureRolloutPoliciesItemAppliesToDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if directoryObjectId != "" {
        urlTplParams["directoryObject%2Did"] = directoryObjectId
    }
    return NewFeatureRolloutPoliciesItemAppliesToDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewFeatureRolloutPoliciesItemAppliesToRequestBuilderInternal instantiates a new AppliesToRequestBuilder and sets the default values.
func NewFeatureRolloutPoliciesItemAppliesToRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FeatureRolloutPoliciesItemAppliesToRequestBuilder) {
    m := &FeatureRolloutPoliciesItemAppliesToRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/featureRolloutPolicies/{featureRolloutPolicy%2Did}/appliesTo{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewFeatureRolloutPoliciesItemAppliesToRequestBuilder instantiates a new AppliesToRequestBuilder and sets the default values.
func NewFeatureRolloutPoliciesItemAppliesToRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FeatureRolloutPoliciesItemAppliesToRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFeatureRolloutPoliciesItemAppliesToRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *FeatureRolloutPoliciesItemAppliesToRequestBuilder) Count()(*FeatureRolloutPoliciesItemAppliesToCountRequestBuilder) {
    return NewFeatureRolloutPoliciesItemAppliesToCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delta provides operations to call the delta method.
func (m *FeatureRolloutPoliciesItemAppliesToRequestBuilder) Delta()(*FeatureRolloutPoliciesItemAppliesToDeltaRequestBuilder) {
    return NewFeatureRolloutPoliciesItemAppliesToDeltaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get nullable. Specifies a list of directoryObjects that feature is enabled for.
func (m *FeatureRolloutPoliciesItemAppliesToRequestBuilder) Get(ctx context.Context, requestConfiguration *FeatureRolloutPoliciesItemAppliesToRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
// GetAvailableExtensionProperties provides operations to call the getAvailableExtensionProperties method.
func (m *FeatureRolloutPoliciesItemAppliesToRequestBuilder) GetAvailableExtensionProperties()(*FeatureRolloutPoliciesItemAppliesToGetAvailableExtensionPropertiesRequestBuilder) {
    return NewFeatureRolloutPoliciesItemAppliesToGetAvailableExtensionPropertiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetByIds provides operations to call the getByIds method.
func (m *FeatureRolloutPoliciesItemAppliesToRequestBuilder) GetByIds()(*FeatureRolloutPoliciesItemAppliesToGetByIdsRequestBuilder) {
    return NewFeatureRolloutPoliciesItemAppliesToGetByIdsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post add an appliesTo on a featureRolloutPolicy object to specify the directoryObject to which the featureRolloutPolicy should be applied.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/featurerolloutpolicy-post-appliesto?view=graph-rest-1.0
func (m *FeatureRolloutPoliciesItemAppliesToRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, requestConfiguration *FeatureRolloutPoliciesItemAppliesToRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
func (m *FeatureRolloutPoliciesItemAppliesToRequestBuilder) Ref()(*FeatureRolloutPoliciesItemAppliesToRefRequestBuilder) {
    return NewFeatureRolloutPoliciesItemAppliesToRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation nullable. Specifies a list of directoryObjects that feature is enabled for.
func (m *FeatureRolloutPoliciesItemAppliesToRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FeatureRolloutPoliciesItemAppliesToRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation add an appliesTo on a featureRolloutPolicy object to specify the directoryObject to which the featureRolloutPolicy should be applied.
func (m *FeatureRolloutPoliciesItemAppliesToRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, requestConfiguration *FeatureRolloutPoliciesItemAppliesToRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ValidateProperties provides operations to call the validateProperties method.
func (m *FeatureRolloutPoliciesItemAppliesToRequestBuilder) ValidateProperties()(*FeatureRolloutPoliciesItemAppliesToValidatePropertiesRequestBuilder) {
    return NewFeatureRolloutPoliciesItemAppliesToValidatePropertiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *FeatureRolloutPoliciesItemAppliesToRequestBuilder) WithUrl(rawUrl string)(*FeatureRolloutPoliciesItemAppliesToRequestBuilder) {
    return NewFeatureRolloutPoliciesItemAppliesToRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
