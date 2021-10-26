package listitem

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i23f1cda7bfe631f0ba0ca42cbc617403778bad2a3375d879652516aa07a93fb1 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/listitem/versions"
    i2ca2c646c48632855b77497ed7674e192faec4dd86004f8544a203f4e572a581 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/listitem/analytics"
    i8f9c43b2560e8e6a51853fb4608862b6166cac2a5714fc3dca44939ea0ec52e4 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/listitem/fields"
    ida206328e87a4a7537e1818e2581887f623f7805a74e9741d4585bead51c8055 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/listitem/driveitem"
    idd2e57cc57bdeb40f198e4baff301e5864036c0df0d2b539f4c9a14017a132cd "github.com/microsoftgraph/msgraph-sdk-go/shares/item/listitem/getactivitiesbyinterval"
    if94999b8a3e6e5a5347772ebfe7cd633500318f71903f991e104b29eb38314d3 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/listitem/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    i72269e45229a663a9977d48d0cd12eff60a4fe9dee7fb89be37703026e55fa41 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/listitem/versions/item"
)

// Builds and executes requests for operations under \shares\{sharedDriveItem-id}\listItem
type ListItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Used to access the underlying listItem
type ListItemRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
func (m *ListItemRequestBuilder) Analytics()(*i2ca2c646c48632855b77497ed7674e192faec4dd86004f8544a203f4e572a581.AnalyticsRequestBuilder) {
    return i2ca2c646c48632855b77497ed7674e192faec4dd86004f8544a203f4e572a581.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new ListItemRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemRequestBuilder) {
    m := &ListItemRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/shares/{sharedDriveItem_id}/listItem{?select,expand}";
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
// Used to access the underlying listItem
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
// Used to access the underlying listItem
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
// Used to access the underlying listItem
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
// Used to access the underlying listItem
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
func (m *ListItemRequestBuilder) DriveItem()(*ida206328e87a4a7537e1818e2581887f623f7805a74e9741d4585bead51c8055.DriveItemRequestBuilder) {
    return ida206328e87a4a7537e1818e2581887f623f7805a74e9741d4585bead51c8055.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Fields()(*i8f9c43b2560e8e6a51853fb4608862b6166cac2a5714fc3dca44939ea0ec52e4.FieldsRequestBuilder) {
    return i8f9c43b2560e8e6a51853fb4608862b6166cac2a5714fc3dca44939ea0ec52e4.NewFieldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Used to access the underlying listItem
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
// Builds and executes requests for operations under \shares\{sharedDriveItem-id}\listItem\microsoft.graph.getActivitiesByInterval()
func (m *ListItemRequestBuilder) GetActivitiesByInterval()(*idd2e57cc57bdeb40f198e4baff301e5864036c0df0d2b539f4c9a14017a132cd.GetActivitiesByIntervalRequestBuilder) {
    return idd2e57cc57bdeb40f198e4baff301e5864036c0df0d2b539f4c9a14017a132cd.NewGetActivitiesByIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \shares\{sharedDriveItem-id}\listItem\microsoft.graph.getActivitiesByInterval(startDateTime='{startDateTime}',endDateTime='{endDateTime}',interval='{interval}')
// Parameters:
//  - endDateTime : Usage: endDateTime={endDateTime}
//  - interval : Usage: interval={interval}
//  - startDateTime : Usage: startDateTime={startDateTime}
func (m *ListItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(startDateTime *string, endDateTime *string, interval *string)(*if94999b8a3e6e5a5347772ebfe7cd633500318f71903f991e104b29eb38314d3.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return if94999b8a3e6e5a5347772ebfe7cd633500318f71903f991e104b29eb38314d3.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, startDateTime, endDateTime, interval);
}
// Used to access the underlying listItem
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
func (m *ListItemRequestBuilder) Versions()(*i23f1cda7bfe631f0ba0ca42cbc617403778bad2a3375d879652516aa07a93fb1.VersionsRequestBuilder) {
    return i23f1cda7bfe631f0ba0ca42cbc617403778bad2a3375d879652516aa07a93fb1.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.shares.item.listItem.versions.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ListItemRequestBuilder) VersionsById(id string)(*i72269e45229a663a9977d48d0cd12eff60a4fe9dee7fb89be37703026e55fa41.ListItemVersionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItemVersion_id"] = id
    }
    return i72269e45229a663a9977d48d0cd12eff60a4fe9dee7fb89be37703026e55fa41.NewListItemVersionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
