package segments

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19 "github.com/microsoftgraph/msgraph-sdk-go/models/callrecords"
    i3788037479c68f720877816caa90c67142919b35fd06303b8e9fd77903f8bcb9 "github.com/microsoftgraph/msgraph-sdk-go/communications/callrecords/item/sessions/item/segments/count"
)

// SegmentsRequestBuilder provides operations to manage the segments property of the microsoft.graph.callRecords.session entity.
type SegmentsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SegmentsRequestBuilderGetOptions options for Get
type SegmentsRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *SegmentsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// SegmentsRequestBuilderGetQueryParameters the list of segments involved in the session. Read-only. Nullable.
type SegmentsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// SegmentsRequestBuilderPostOptions options for Post
type SegmentsRequestBuilderPostOptions struct {
    // 
    Body iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.Segmentable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewSegmentsRequestBuilderInternal instantiates a new SegmentsRequestBuilder and sets the default values.
func NewSegmentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SegmentsRequestBuilder) {
    m := &SegmentsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/callRecords/{callRecord_id}/sessions/{session_id}/segments{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSegmentsRequestBuilder instantiates a new SegmentsRequestBuilder and sets the default values.
func NewSegmentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SegmentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSegmentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *SegmentsRequestBuilder) Count()(*i3788037479c68f720877816caa90c67142919b35fd06303b8e9fd77903f8bcb9.CountRequestBuilder) {
    return i3788037479c68f720877816caa90c67142919b35fd06303b8e9fd77903f8bcb9.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the list of segments involved in the session. Read-only. Nullable.
func (m *SegmentsRequestBuilder) CreateGetRequestInformation(options *SegmentsRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to segments for communications
func (m *SegmentsRequestBuilder) CreatePostRequestInformation(options *SegmentsRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
// Get the list of segments involved in the session. Read-only. Nullable.
func (m *SegmentsRequestBuilder) Get(options *SegmentsRequestBuilderGetOptions)(iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.SegmentCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.CreateSegmentCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.SegmentCollectionResponseable), nil
}
// Post create new navigation property to segments for communications
func (m *SegmentsRequestBuilder) Post(options *SegmentsRequestBuilderPostOptions)(iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.Segmentable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.CreateSegmentFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.Segmentable), nil
}
