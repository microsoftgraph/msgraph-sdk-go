package domains

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilder provides operations to manage the serviceConfigurationRecords property of the microsoft.graph.domain entity.
type ItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilderGetQueryParameters dNS records the customer adds to the DNS zone file of the domain before the domain can be used by Microsoft Online services. Read-only, Nullable. Supports $expand.
type ItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilderGetQueryParameters
}
// ItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilderInternal instantiates a new ItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilder and sets the default values.
func NewItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilder) {
    m := &ItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/domains/{domain%2Did}/serviceConfigurationRecords/{domainDnsRecord%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilder instantiates a new ItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilder and sets the default values.
func NewItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property serviceConfigurationRecords for domains
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get dNS records the customer adds to the DNS zone file of the domain before the domain can be used by Microsoft Online services. Read-only, Nullable. Supports $expand.
// returns a DomainDnsRecordable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DomainDnsRecordable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDomainDnsRecordFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DomainDnsRecordable), nil
}
// Patch update the navigation property serviceConfigurationRecords in domains
// returns a DomainDnsRecordable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DomainDnsRecordable, requestConfiguration *ItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DomainDnsRecordable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDomainDnsRecordFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DomainDnsRecordable), nil
}
// ToDeleteRequestInformation delete navigation property serviceConfigurationRecords for domains
// returns a *RequestInformation when successful
func (m *ItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation dNS records the customer adds to the DNS zone file of the domain before the domain can be used by Microsoft Online services. Read-only, Nullable. Supports $expand.
// returns a *RequestInformation when successful
func (m *ItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property serviceConfigurationRecords in domains
// returns a *RequestInformation when successful
func (m *ItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DomainDnsRecordable, requestConfiguration *ItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilder when successful
func (m *ItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilder) WithUrl(rawUrl string)(*ItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilder) {
    return NewItemServiceconfigurationrecordsDomainDnsRecordItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
