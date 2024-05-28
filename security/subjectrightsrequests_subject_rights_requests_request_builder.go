package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// SubjectrightsrequestsSubjectRightsRequestsRequestBuilder provides operations to manage the subjectRightsRequests property of the microsoft.graph.security entity.
type SubjectrightsrequestsSubjectRightsRequestsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SubjectrightsrequestsSubjectRightsRequestsRequestBuilderGetQueryParameters get subjectRightsRequests from security
type SubjectrightsrequestsSubjectRightsRequestsRequestBuilderGetQueryParameters struct {
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
// SubjectrightsrequestsSubjectRightsRequestsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SubjectrightsrequestsSubjectRightsRequestsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SubjectrightsrequestsSubjectRightsRequestsRequestBuilderGetQueryParameters
}
// SubjectrightsrequestsSubjectRightsRequestsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SubjectrightsrequestsSubjectRightsRequestsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BySubjectRightsRequestId provides operations to manage the subjectRightsRequests property of the microsoft.graph.security entity.
// returns a *SubjectrightsrequestsSubjectRightsRequestItemRequestBuilder when successful
func (m *SubjectrightsrequestsSubjectRightsRequestsRequestBuilder) BySubjectRightsRequestId(subjectRightsRequestId string)(*SubjectrightsrequestsSubjectRightsRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if subjectRightsRequestId != "" {
        urlTplParams["subjectRightsRequest%2Did"] = subjectRightsRequestId
    }
    return NewSubjectrightsrequestsSubjectRightsRequestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewSubjectrightsrequestsSubjectRightsRequestsRequestBuilderInternal instantiates a new SubjectrightsrequestsSubjectRightsRequestsRequestBuilder and sets the default values.
func NewSubjectrightsrequestsSubjectRightsRequestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SubjectrightsrequestsSubjectRightsRequestsRequestBuilder) {
    m := &SubjectrightsrequestsSubjectRightsRequestsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/subjectRightsRequests{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewSubjectrightsrequestsSubjectRightsRequestsRequestBuilder instantiates a new SubjectrightsrequestsSubjectRightsRequestsRequestBuilder and sets the default values.
func NewSubjectrightsrequestsSubjectRightsRequestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SubjectrightsrequestsSubjectRightsRequestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSubjectrightsrequestsSubjectRightsRequestsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *SubjectrightsrequestsCountRequestBuilder when successful
func (m *SubjectrightsrequestsSubjectRightsRequestsRequestBuilder) Count()(*SubjectrightsrequestsCountRequestBuilder) {
    return NewSubjectrightsrequestsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get subjectRightsRequests from security
// returns a SubjectRightsRequestCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *SubjectrightsrequestsSubjectRightsRequestsRequestBuilder) Get(ctx context.Context, requestConfiguration *SubjectrightsrequestsSubjectRightsRequestsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SubjectRightsRequestCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSubjectRightsRequestCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SubjectRightsRequestCollectionResponseable), nil
}
// Post create new navigation property to subjectRightsRequests for security
// returns a SubjectRightsRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *SubjectrightsrequestsSubjectRightsRequestsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SubjectRightsRequestable, requestConfiguration *SubjectrightsrequestsSubjectRightsRequestsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SubjectRightsRequestable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSubjectRightsRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SubjectRightsRequestable), nil
}
// ToGetRequestInformation get subjectRightsRequests from security
// returns a *RequestInformation when successful
func (m *SubjectrightsrequestsSubjectRightsRequestsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SubjectrightsrequestsSubjectRightsRequestsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to subjectRightsRequests for security
// returns a *RequestInformation when successful
func (m *SubjectrightsrequestsSubjectRightsRequestsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SubjectRightsRequestable, requestConfiguration *SubjectrightsrequestsSubjectRightsRequestsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *SubjectrightsrequestsSubjectRightsRequestsRequestBuilder when successful
func (m *SubjectrightsrequestsSubjectRightsRequestsRequestBuilder) WithUrl(rawUrl string)(*SubjectrightsrequestsSubjectRightsRequestsRequestBuilder) {
    return NewSubjectrightsrequestsSubjectRightsRequestsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
