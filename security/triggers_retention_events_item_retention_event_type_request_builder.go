package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// TriggersRetentionEventsItemRetentionEventTypeRequestBuilder provides operations to manage the retentionEventType property of the microsoft.graph.security.retentionEvent entity.
type TriggersRetentionEventsItemRetentionEventTypeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TriggersRetentionEventsItemRetentionEventTypeRequestBuilderGetQueryParameters specifies the event that will start the retention period for labels that use this event type when an event is created.
type TriggersRetentionEventsItemRetentionEventTypeRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TriggersRetentionEventsItemRetentionEventTypeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TriggersRetentionEventsItemRetentionEventTypeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TriggersRetentionEventsItemRetentionEventTypeRequestBuilderGetQueryParameters
}
// NewTriggersRetentionEventsItemRetentionEventTypeRequestBuilderInternal instantiates a new RetentionEventTypeRequestBuilder and sets the default values.
func NewTriggersRetentionEventsItemRetentionEventTypeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TriggersRetentionEventsItemRetentionEventTypeRequestBuilder) {
    m := &TriggersRetentionEventsItemRetentionEventTypeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/triggers/retentionEvents/{retentionEvent%2Did}/retentionEventType{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewTriggersRetentionEventsItemRetentionEventTypeRequestBuilder instantiates a new RetentionEventTypeRequestBuilder and sets the default values.
func NewTriggersRetentionEventsItemRetentionEventTypeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TriggersRetentionEventsItemRetentionEventTypeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTriggersRetentionEventsItemRetentionEventTypeRequestBuilderInternal(urlParams, requestAdapter)
}
// Get specifies the event that will start the retention period for labels that use this event type when an event is created.
func (m *TriggersRetentionEventsItemRetentionEventTypeRequestBuilder) Get(ctx context.Context, requestConfiguration *TriggersRetentionEventsItemRetentionEventTypeRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.RetentionEventTypeable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateRetentionEventTypeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.RetentionEventTypeable), nil
}
// ToGetRequestInformation specifies the event that will start the retention period for labels that use this event type when an event is created.
func (m *TriggersRetentionEventsItemRetentionEventTypeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TriggersRetentionEventsItemRetentionEventTypeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
