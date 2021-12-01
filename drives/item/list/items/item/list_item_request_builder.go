package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i3512ac0de8a2f03f53c759c6626c35e234f70725f44e86450b8a4ebdb7648618 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/items/item/fields"
    i41fe9922169e567eac85f2759e67712ffa2957866f1746d066768c109c96a013 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/items/item/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    i59cec2d7c2cbda5e5b69c827dceb517387e69a57f2d5107676864d82ab2f4bfa "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/items/item/analytics"
    i7f11fe17492440df5d51d7aa1a4338cdd51adaa74030bf12f4efec4ad1041354 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/items/item/versions"
    ia387bfdbe6d76d0774c7bd8d256acb6489f5abdd6eb8267ac5999332294dc3f8 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/items/item/driveitem"
    iab650811575ccc52d1d3e05cf12dba4b155aa2c78edc416d4b3da5ffb2b7777d "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/items/item/getactivitiesbyinterval"
    i284a388448822a877cc80b553ec4328aefece3f18982b7d6f35ab5db6328ab0d "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/items/item/versions/item"
)

// ListItemRequestBuilder builds and executes requests for operations under \drives\{drive-id}\list\items\{listItem-id}
type ListItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ListItemRequestBuilderDeleteOptions options for Delete
type ListItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ListItemRequestBuilderGetOptions options for Get
type ListItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ListItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ListItemRequestBuilderGetQueryParameters all items contained in the list.
type ListItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ListItemRequestBuilderPatchOptions options for Patch
type ListItemRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ListItem;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ListItemRequestBuilder) Analytics()(*i59cec2d7c2cbda5e5b69c827dceb517387e69a57f2d5107676864d82ab2f4bfa.AnalyticsRequestBuilder) {
    return i59cec2d7c2cbda5e5b69c827dceb517387e69a57f2d5107676864d82ab2f4bfa.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewListItemRequestBuilderInternal instantiates a new ListItemRequestBuilder and sets the default values.
func NewListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemRequestBuilder) {
    m := &ListItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive_id}/list/items/{listItem_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewListItemRequestBuilder instantiates a new ListItemRequestBuilder and sets the default values.
func NewListItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewListItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation all items contained in the list.
func (m *ListItemRequestBuilder) CreateDeleteRequestInformation(options *ListItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation all items contained in the list.
func (m *ListItemRequestBuilder) CreateGetRequestInformation(options *ListItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation all items contained in the list.
func (m *ListItemRequestBuilder) CreatePatchRequestInformation(options *ListItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete all items contained in the list.
func (m *ListItemRequestBuilder) Delete(options *ListItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ListItemRequestBuilder) DriveItem()(*ia387bfdbe6d76d0774c7bd8d256acb6489f5abdd6eb8267ac5999332294dc3f8.DriveItemRequestBuilder) {
    return ia387bfdbe6d76d0774c7bd8d256acb6489f5abdd6eb8267ac5999332294dc3f8.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Fields()(*i3512ac0de8a2f03f53c759c6626c35e234f70725f44e86450b8a4ebdb7648618.FieldsRequestBuilder) {
    return i3512ac0de8a2f03f53c759c6626c35e234f70725f44e86450b8a4ebdb7648618.NewFieldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get all items contained in the list.
func (m *ListItemRequestBuilder) Get(options *ListItemRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ListItem, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewListItem() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ListItem), nil
}
// GetActivitiesByInterval builds and executes requests for operations under \drives\{drive-id}\list\items\{listItem-id}\microsoft.graph.getActivitiesByInterval()
func (m *ListItemRequestBuilder) GetActivitiesByInterval()(*iab650811575ccc52d1d3e05cf12dba4b155aa2c78edc416d4b3da5ffb2b7777d.GetActivitiesByIntervalRequestBuilder) {
    return iab650811575ccc52d1d3e05cf12dba4b155aa2c78edc416d4b3da5ffb2b7777d.NewGetActivitiesByIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval builds and executes requests for operations under \drives\{drive-id}\list\items\{listItem-id}\microsoft.graph.getActivitiesByInterval(startDateTime='{startDateTime}',endDateTime='{endDateTime}',interval='{interval}')
func (m *ListItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(startDateTime *string, endDateTime *string, interval *string)(*i41fe9922169e567eac85f2759e67712ffa2957866f1746d066768c109c96a013.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return i41fe9922169e567eac85f2759e67712ffa2957866f1746d066768c109c96a013.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, startDateTime, endDateTime, interval);
}
// Patch all items contained in the list.
func (m *ListItemRequestBuilder) Patch(options *ListItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ListItemRequestBuilder) Versions()(*i7f11fe17492440df5d51d7aa1a4338cdd51adaa74030bf12f4efec4ad1041354.VersionsRequestBuilder) {
    return i7f11fe17492440df5d51d7aa1a4338cdd51adaa74030bf12f4efec4ad1041354.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.list.items.item.versions.item collection
func (m *ListItemRequestBuilder) VersionsById(id string)(*i284a388448822a877cc80b553ec4328aefece3f18982b7d6f35ab5db6328ab0d.ListItemVersionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItemVersion_id"] = id
    }
    return i284a388448822a877cc80b553ec4328aefece3f18982b7d6f35ab5db6328ab0d.NewListItemVersionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
