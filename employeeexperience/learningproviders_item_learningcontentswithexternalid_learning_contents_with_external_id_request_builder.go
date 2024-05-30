package employeeexperience

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// LearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilder provides operations to manage the learningContents property of the microsoft.graph.learningProvider entity.
type LearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// LearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilderGetQueryParameters get the specified learningContent resource which represents the metadata of the specified provider's ingested content.
type LearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilderGetQueryParameters
}
// LearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewLearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilderInternal instantiates a new LearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilder and sets the default values.
func NewLearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, externalId *string)(*LearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilder) {
    m := &LearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/employeeExperience/learningProviders/{learningProvider%2Did}/learningContents(externalId='{externalId}'){?%24expand,%24select}", pathParameters),
    }
    if externalId != nil {
        m.BaseRequestBuilder.PathParameters["externalId"] = *externalId
    }
    return m
}
// NewLearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilder instantiates a new LearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilder and sets the default values.
func NewLearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Delete delete the specified learningContent resource that represents the metadata of the specified provider's ingested content.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/learningprovider-delete-learningcontents?view=graph-rest-1.0
func (m *LearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilder) Delete(ctx context.Context, requestConfiguration *LearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get the specified learningContent resource which represents the metadata of the specified provider's ingested content.
// returns a LearningContentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/learningcontent-get?view=graph-rest-1.0
func (m *LearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilder) Get(ctx context.Context, requestConfiguration *LearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LearningContentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateLearningContentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LearningContentable), nil
}
// Patch update the navigation property learningContents in employeeExperience
// returns a LearningContentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LearningContentable, requestConfiguration *LearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LearningContentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateLearningContentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LearningContentable), nil
}
// ToDeleteRequestInformation delete the specified learningContent resource that represents the metadata of the specified provider's ingested content.
// returns a *RequestInformation when successful
func (m *LearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *LearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get the specified learningContent resource which represents the metadata of the specified provider's ingested content.
// returns a *RequestInformation when successful
func (m *LearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property learningContents in employeeExperience
// returns a *RequestInformation when successful
func (m *LearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LearningContentable, requestConfiguration *LearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilder when successful
func (m *LearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilder) WithUrl(rawUrl string)(*LearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilder) {
    return NewLearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
