package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i67e41a929042c50c03aaffa467dfbb883460f0386b1bbbaf0c67841405db5350 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/lists/item/items/item/versions"
    icbaaf60777717606ab34d69e5cf2a34cf755e1f7cedab6950007a9ffb84d7008 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/lists/item/items/item/getactivitiesbyinterval"
    icdd597e52b660be79cd5ae3dcabe4d5def49be59cbfe26fccee978568262dd29 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/lists/item/items/item/driveitem"
    id8d804de7d48a69f9443f09eaf49a569e6a03d16aef4ff2ae5f50fec29025877 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/lists/item/items/item/analytics"
    idf96688e335c8e438ce9407a133e7cd3b1cafc5a2f1bc2b302f7c48445588c8b "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/lists/item/items/item/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    ieff756b62a651236a62880261d2917b918b53cd47801295415a45625143d4c06 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/lists/item/items/item/fields"
    i29f00f317ab4934d3d28db324d28b5fb20f67fb04a50c6b38d4489274ff85efb "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/lists/item/items/item/versions/item"
)

// ListItemItemRequestBuilder provides operations to manage the items property of the microsoft.graph.list entity.
type ListItemItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ListItemItemRequestBuilderDeleteOptions options for Delete
type ListItemItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// ListItemItemRequestBuilderGetOptions options for Get
type ListItemItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *ListItemItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// ListItemItemRequestBuilderGetQueryParameters all items contained in the list.
type ListItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ListItemItemRequestBuilderPatchOptions options for Patch
type ListItemItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ListItemable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// Analytics the analytics property
func (m *ListItemItemRequestBuilder) Analytics()(*id8d804de7d48a69f9443f09eaf49a569e6a03d16aef4ff2ae5f50fec29025877.AnalyticsRequestBuilder) {
    return id8d804de7d48a69f9443f09eaf49a569e6a03d16aef4ff2ae5f50fec29025877.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewListItemItemRequestBuilderInternal instantiates a new ListItemItemRequestBuilder and sets the default values.
func NewListItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ListItemItemRequestBuilder) {
    m := &ListItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/sites/{site_id}/lists/{list_id}/items/{listItem_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewListItemItemRequestBuilder instantiates a new ListItemItemRequestBuilder and sets the default values.
func NewListItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ListItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewListItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property items for groups
func (m *ListItemItemRequestBuilder) CreateDeleteRequestInformation(options *ListItemItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation all items contained in the list.
func (m *ListItemItemRequestBuilder) CreateGetRequestInformation(options *ListItemItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property items in groups
func (m *ListItemItemRequestBuilder) CreatePatchRequestInformation(options *ListItemItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property items for groups
func (m *ListItemItemRequestBuilder) Delete(options *ListItemItemRequestBuilderDeleteOptions)(error) {
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
// DriveItem the driveItem property
func (m *ListItemItemRequestBuilder) DriveItem()(*icdd597e52b660be79cd5ae3dcabe4d5def49be59cbfe26fccee978568262dd29.DriveItemRequestBuilder) {
    return icdd597e52b660be79cd5ae3dcabe4d5def49be59cbfe26fccee978568262dd29.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Fields the fields property
func (m *ListItemItemRequestBuilder) Fields()(*ieff756b62a651236a62880261d2917b918b53cd47801295415a45625143d4c06.FieldsRequestBuilder) {
    return ieff756b62a651236a62880261d2917b918b53cd47801295415a45625143d4c06.NewFieldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get all items contained in the list.
func (m *ListItemItemRequestBuilder) Get(options *ListItemItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ListItemable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateListItemFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ListItemable), nil
}
// GetActivitiesByInterval provides operations to call the getActivitiesByInterval method.
func (m *ListItemItemRequestBuilder) GetActivitiesByInterval()(*icbaaf60777717606ab34d69e5cf2a34cf755e1f7cedab6950007a9ffb84d7008.GetActivitiesByIntervalRequestBuilder) {
    return icbaaf60777717606ab34d69e5cf2a34cf755e1f7cedab6950007a9ffb84d7008.NewGetActivitiesByIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
func (m *ListItemItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*idf96688e335c8e438ce9407a133e7cd3b1cafc5a2f1bc2b302f7c48445588c8b.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return idf96688e335c8e438ce9407a133e7cd3b1cafc5a2f1bc2b302f7c48445588c8b.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// Patch update the navigation property items in groups
func (m *ListItemItemRequestBuilder) Patch(options *ListItemItemRequestBuilderPatchOptions)(error) {
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
// Versions the versions property
func (m *ListItemItemRequestBuilder) Versions()(*i67e41a929042c50c03aaffa467dfbb883460f0386b1bbbaf0c67841405db5350.VersionsRequestBuilder) {
    return i67e41a929042c50c03aaffa467dfbb883460f0386b1bbbaf0c67841405db5350.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.sites.item.lists.item.items.item.versions.item collection
func (m *ListItemItemRequestBuilder) VersionsById(id string)(*i29f00f317ab4934d3d28db324d28b5fb20f67fb04a50c6b38d4489274ff85efb.ListItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItemVersion_id"] = id
    }
    return i29f00f317ab4934d3d28db324d28b5fb20f67fb04a50c6b38d4489274ff85efb.NewListItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
