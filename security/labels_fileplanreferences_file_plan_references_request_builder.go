package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// LabelsFileplanreferencesFilePlanReferencesRequestBuilder provides operations to manage the filePlanReferences property of the microsoft.graph.security.labelsRoot entity.
type LabelsFileplanreferencesFilePlanReferencesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LabelsFileplanreferencesFilePlanReferencesRequestBuilderGetQueryParameters get a list of the filePlanReferenceTemplate objects and their properties.
type LabelsFileplanreferencesFilePlanReferencesRequestBuilderGetQueryParameters struct {
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
// LabelsFileplanreferencesFilePlanReferencesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LabelsFileplanreferencesFilePlanReferencesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LabelsFileplanreferencesFilePlanReferencesRequestBuilderGetQueryParameters
}
// LabelsFileplanreferencesFilePlanReferencesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LabelsFileplanreferencesFilePlanReferencesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByFilePlanReferenceTemplateId provides operations to manage the filePlanReferences property of the microsoft.graph.security.labelsRoot entity.
// returns a *LabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilder when successful
func (m *LabelsFileplanreferencesFilePlanReferencesRequestBuilder) ByFilePlanReferenceTemplateId(filePlanReferenceTemplateId string)(*LabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if filePlanReferenceTemplateId != "" {
        urlTplParams["filePlanReferenceTemplate%2Did"] = filePlanReferenceTemplateId
    }
    return NewLabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewLabelsFileplanreferencesFilePlanReferencesRequestBuilderInternal instantiates a new LabelsFileplanreferencesFilePlanReferencesRequestBuilder and sets the default values.
func NewLabelsFileplanreferencesFilePlanReferencesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LabelsFileplanreferencesFilePlanReferencesRequestBuilder) {
    m := &LabelsFileplanreferencesFilePlanReferencesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/labels/filePlanReferences{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewLabelsFileplanreferencesFilePlanReferencesRequestBuilder instantiates a new LabelsFileplanreferencesFilePlanReferencesRequestBuilder and sets the default values.
func NewLabelsFileplanreferencesFilePlanReferencesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LabelsFileplanreferencesFilePlanReferencesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLabelsFileplanreferencesFilePlanReferencesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *LabelsFileplanreferencesCountRequestBuilder when successful
func (m *LabelsFileplanreferencesFilePlanReferencesRequestBuilder) Count()(*LabelsFileplanreferencesCountRequestBuilder) {
    return NewLabelsFileplanreferencesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the filePlanReferenceTemplate objects and their properties.
// returns a FilePlanReferenceTemplateCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-labelsroot-list-fileplanreferences?view=graph-rest-1.0
func (m *LabelsFileplanreferencesFilePlanReferencesRequestBuilder) Get(ctx context.Context, requestConfiguration *LabelsFileplanreferencesFilePlanReferencesRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.FilePlanReferenceTemplateCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateFilePlanReferenceTemplateCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.FilePlanReferenceTemplateCollectionResponseable), nil
}
// Post create a new filePlanReferenceTemplate object.
// returns a FilePlanReferenceTemplateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-labelsroot-post-fileplanreferences?view=graph-rest-1.0
func (m *LabelsFileplanreferencesFilePlanReferencesRequestBuilder) Post(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.FilePlanReferenceTemplateable, requestConfiguration *LabelsFileplanreferencesFilePlanReferencesRequestBuilderPostRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.FilePlanReferenceTemplateable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateFilePlanReferenceTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.FilePlanReferenceTemplateable), nil
}
// ToGetRequestInformation get a list of the filePlanReferenceTemplate objects and their properties.
// returns a *RequestInformation when successful
func (m *LabelsFileplanreferencesFilePlanReferencesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LabelsFileplanreferencesFilePlanReferencesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new filePlanReferenceTemplate object.
// returns a *RequestInformation when successful
func (m *LabelsFileplanreferencesFilePlanReferencesRequestBuilder) ToPostRequestInformation(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.FilePlanReferenceTemplateable, requestConfiguration *LabelsFileplanreferencesFilePlanReferencesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LabelsFileplanreferencesFilePlanReferencesRequestBuilder when successful
func (m *LabelsFileplanreferencesFilePlanReferencesRequestBuilder) WithUrl(rawUrl string)(*LabelsFileplanreferencesFilePlanReferencesRequestBuilder) {
    return NewLabelsFileplanreferencesFilePlanReferencesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
