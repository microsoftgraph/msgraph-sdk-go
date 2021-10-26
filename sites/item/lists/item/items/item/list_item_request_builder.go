package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i27181275dde742d95e48c29d39848d85ed01e2d7b5fd1ef749b368ce4147130e "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/items/item/driveitem"
    i7f3e5620be51565e864cf4e746c9b26e2c47d2d2b1b1c97a4c4ef42c65645263 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/items/item/versions"
    i85333e43888dd43b991f03d537f5ae34a91ab18fe219c1aff65a43eec70df186 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/items/item/fields"
    i95d11a43632c0b7f2416167fd5578e5bf7c62439e1ebe37312bda2718b31a8fe "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/items/item/analytics"
    i9fb260dad120a024acc09a04318bb0b07fd9c07bb184296c494c7072ecc4b09a "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/items/item/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    ic0757d0a0ffa83f3296bfa19f2c157e45033523373412dda78c5df27ddc29bc6 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/items/item/getactivitiesbyinterval"
    ie4696deb90abb7ecec8da795c90a12b52f9218953b092733c93bbe7dcf472577 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/items/item/versions/item"
)

// Builds and executes requests for operations under \sites\{site-id}\lists\{list-id}\items\{listItem-id}
type ListItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// All items contained in the list.
type ListItemRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
func (m *ListItemRequestBuilder) Analytics()(*i95d11a43632c0b7f2416167fd5578e5bf7c62439e1ebe37312bda2718b31a8fe.AnalyticsRequestBuilder) {
    return i95d11a43632c0b7f2416167fd5578e5bf7c62439e1ebe37312bda2718b31a8fe.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new ListItemRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemRequestBuilder) {
    m := &ListItemRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/sites/{site_id}/lists/{list_id}/items/{listItem_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ListItemRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewListItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewListItemRequestBuilderInternal(urlParams, requestAdapter)
}
// All items contained in the list.
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *ListItemRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// All items contained in the list.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *ListItemRequestBuilder) CreateGetRequestInformation(q func (value *ListItemRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ListItemRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// All items contained in the list.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *ListItemRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ListItem, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// All items contained in the list.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *ListItemRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *ListItemRequestBuilder) DriveItem()(*i27181275dde742d95e48c29d39848d85ed01e2d7b5fd1ef749b368ce4147130e.DriveItemRequestBuilder) {
    return i27181275dde742d95e48c29d39848d85ed01e2d7b5fd1ef749b368ce4147130e.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Fields()(*i85333e43888dd43b991f03d537f5ae34a91ab18fe219c1aff65a43eec70df186.FieldsRequestBuilder) {
    return i85333e43888dd43b991f03d537f5ae34a91ab18fe219c1aff65a43eec70df186.NewFieldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// All items contained in the list.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *ListItemRequestBuilder) Get(q func (value *ListItemRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ListItem, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewListItem() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ListItem), nil
}
// Builds and executes requests for operations under \sites\{site-id}\lists\{list-id}\items\{listItem-id}\microsoft.graph.getActivitiesByInterval()
func (m *ListItemRequestBuilder) GetActivitiesByInterval()(*ic0757d0a0ffa83f3296bfa19f2c157e45033523373412dda78c5df27ddc29bc6.GetActivitiesByIntervalRequestBuilder) {
    return ic0757d0a0ffa83f3296bfa19f2c157e45033523373412dda78c5df27ddc29bc6.NewGetActivitiesByIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \sites\{site-id}\lists\{list-id}\items\{listItem-id}\microsoft.graph.getActivitiesByInterval(startDateTime='{startDateTime}',endDateTime='{endDateTime}',interval='{interval}')
// Parameters:
//  - endDateTime : Usage: endDateTime={endDateTime}
//  - interval : Usage: interval={interval}
//  - startDateTime : Usage: startDateTime={startDateTime}
func (m *ListItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(startDateTime *string, endDateTime *string, interval *string)(*i9fb260dad120a024acc09a04318bb0b07fd9c07bb184296c494c7072ecc4b09a.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return i9fb260dad120a024acc09a04318bb0b07fd9c07bb184296c494c7072ecc4b09a.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, startDateTime, endDateTime, interval);
}
// All items contained in the list.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *ListItemRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ListItem, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *ListItemRequestBuilder) Versions()(*i7f3e5620be51565e864cf4e746c9b26e2c47d2d2b1b1c97a4c4ef42c65645263.VersionsRequestBuilder) {
    return i7f3e5620be51565e864cf4e746c9b26e2c47d2d2b1b1c97a4c4ef42c65645263.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.lists.item.items.item.versions.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ListItemRequestBuilder) VersionsById(id string)(*ie4696deb90abb7ecec8da795c90a12b52f9218953b092733c93bbe7dcf472577.ListItemVersionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItemVersion_id"] = id
    }
    return ie4696deb90abb7ecec8da795c90a12b52f9218953b092733c93bbe7dcf472577.NewListItemVersionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
