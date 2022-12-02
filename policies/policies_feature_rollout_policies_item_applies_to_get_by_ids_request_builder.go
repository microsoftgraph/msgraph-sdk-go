package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsRequestBuilder provides operations to call the getByIds method.
type PoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsRequestBuilderInternal instantiates a new GetByIdsRequestBuilder and sets the default values.
func NewPoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsRequestBuilder) {
    m := &PoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies/featureRolloutPolicies/{featureRolloutPolicy%2Did}/appliesTo/microsoft.graph.getByIds";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsRequestBuilder instantiates a new GetByIdsRequestBuilder and sets the default values.
func NewPoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation return the directory objects specified in a list of IDs. Some common uses for this function are to:
func (m *PoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsRequestBuilder) CreatePostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsPostRequestBodyable, requestConfiguration *PoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post return the directory objects specified in a list of IDs. Some common uses for this function are to:
func (m *PoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsPostRequestBodyable, requestConfiguration *PoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsResponseable), nil
}
