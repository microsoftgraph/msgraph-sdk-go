package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19 "github.com/microsoftgraph/msgraph-sdk-go/models/callrecords"
    i44905f86d7fc9263aedf143e8e39a3aa82a76eb05b10809a989cddd52d2125dd "github.com/microsoftgraph/msgraph-sdk-go/communications/callrecords/item/sessions/item/segments"
    ib91adf431219411bc95020ad3d58710ee71e71aba8aa4198d6e77a14ec635f0a "github.com/microsoftgraph/msgraph-sdk-go/communications/callrecords/item/sessions/item/segments/item"
)

// SessionItemRequestBuilder provides operations to manage the sessions property of the microsoft.graph.callRecords.callRecord entity.
type SessionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SessionItemRequestBuilderDeleteOptions options for Delete
type SessionItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// SessionItemRequestBuilderGetOptions options for Get
type SessionItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *SessionItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// SessionItemRequestBuilderGetQueryParameters list of sessions involved in the call. Peer-to-peer calls typically only have one session, whereas group calls typically have at least one session per participant. Read-only. Nullable.
type SessionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// SessionItemRequestBuilderPatchOptions options for Patch
type SessionItemRequestBuilderPatchOptions struct {
    // 
    Body iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.Sessionable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewSessionItemRequestBuilderInternal instantiates a new SessionItemRequestBuilder and sets the default values.
func NewSessionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SessionItemRequestBuilder) {
    m := &SessionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/callRecords/{callRecord_id}/sessions/{session_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSessionItemRequestBuilder instantiates a new SessionItemRequestBuilder and sets the default values.
func NewSessionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SessionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSessionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property sessions for communications
func (m *SessionItemRequestBuilder) CreateDeleteRequestInformation(options *SessionItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation list of sessions involved in the call. Peer-to-peer calls typically only have one session, whereas group calls typically have at least one session per participant. Read-only. Nullable.
func (m *SessionItemRequestBuilder) CreateGetRequestInformation(options *SessionItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property sessions in communications
func (m *SessionItemRequestBuilder) CreatePatchRequestInformation(options *SessionItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property sessions for communications
func (m *SessionItemRequestBuilder) Delete(options *SessionItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get list of sessions involved in the call. Peer-to-peer calls typically only have one session, whereas group calls typically have at least one session per participant. Read-only. Nullable.
func (m *SessionItemRequestBuilder) Get(options *SessionItemRequestBuilderGetOptions)(iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.Sessionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.CreateSessionFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.Sessionable), nil
}
// Patch update the navigation property sessions in communications
func (m *SessionItemRequestBuilder) Patch(options *SessionItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Segments the segments property
func (m *SessionItemRequestBuilder) Segments()(*i44905f86d7fc9263aedf143e8e39a3aa82a76eb05b10809a989cddd52d2125dd.SegmentsRequestBuilder) {
    return i44905f86d7fc9263aedf143e8e39a3aa82a76eb05b10809a989cddd52d2125dd.NewSegmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SegmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.communications.callRecords.item.sessions.item.segments.item collection
func (m *SessionItemRequestBuilder) SegmentsById(id string)(*ib91adf431219411bc95020ad3d58710ee71e71aba8aa4198d6e77a14ec635f0a.SegmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["segment_id"] = id
    }
    return ib91adf431219411bc95020ad3d58710ee71e71aba8aa4198d6e77a14ec635f0a.NewSegmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
