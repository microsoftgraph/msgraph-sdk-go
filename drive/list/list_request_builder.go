package list

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i4ceb3442437d25b4d65a5fb368c3ea2eebf89199b8598c6e3c12a21be842c870 "github.com/microsoftgraph/msgraph-sdk-go/drive/list/contenttypes"
    ibba07da9108e84b41a3d09d01bb4a6886a66e86df95f8f0730fa1c61201467ba "github.com/microsoftgraph/msgraph-sdk-go/drive/list/drive"
    ibf5d29d5462efb100f3ac502d1eba6060e878b0461baab8d107bba2657e634d5 "github.com/microsoftgraph/msgraph-sdk-go/drive/list/subscriptions"
    ibf8d356e7a219821b2879666575a0d558fe0723087550da01c802cb67c2ea4a9 "github.com/microsoftgraph/msgraph-sdk-go/drive/list/items"
    id137f289f6d68aa0ce87e57630353e341f93dbb3fb993d0ec54f3e9944c8865c "github.com/microsoftgraph/msgraph-sdk-go/drive/list/columns"
    i05799d6c1d82ecbf851acb48a1cb741eeecec1479585ec331eb151a1eb12fef7 "github.com/microsoftgraph/msgraph-sdk-go/drive/list/columns/item"
    i0cc711a2d9e9934918c44b76b7f313fb885f4cc911346a17e775fac04e5518ac "github.com/microsoftgraph/msgraph-sdk-go/drive/list/subscriptions/item"
    i666016b9de0fea6c6ca42db4cdc518e2bbe5b22999b3927c4fc4035901b10667 "github.com/microsoftgraph/msgraph-sdk-go/drive/list/contenttypes/item"
    if1844020acb9c67e872cff770020bec5e9fbe0ec4bdb729561240d6a686db31c "github.com/microsoftgraph/msgraph-sdk-go/drive/list/items/item"
)

// Builds and executes requests for operations under \drive\list
type ListRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type ListRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type ListRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ListRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// For drives in SharePoint, the underlying document library list. Read-only. Nullable.
type ListRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type ListRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.List;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ListRequestBuilder) Columns()(*id137f289f6d68aa0ce87e57630353e341f93dbb3fb993d0ec54f3e9944c8865c.ColumnsRequestBuilder) {
    return id137f289f6d68aa0ce87e57630353e341f93dbb3fb993d0ec54f3e9944c8865c.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.list.columns.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ListRequestBuilder) ColumnsById(id string)(*i05799d6c1d82ecbf851acb48a1cb741eeecec1479585ec331eb151a1eb12fef7.ColumnDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i05799d6c1d82ecbf851acb48a1cb741eeecec1479585ec331eb151a1eb12fef7.NewColumnDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new ListRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewListRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListRequestBuilder) {
    m := &ListRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drive/list{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ListRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewListRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewListRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ListRequestBuilder) ContentTypes()(*i4ceb3442437d25b4d65a5fb368c3ea2eebf89199b8598c6e3c12a21be842c870.ContentTypesRequestBuilder) {
    return i4ceb3442437d25b4d65a5fb368c3ea2eebf89199b8598c6e3c12a21be842c870.NewContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.list.contentTypes.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ListRequestBuilder) ContentTypesById(id string)(*i666016b9de0fea6c6ca42db4cdc518e2bbe5b22999b3927c4fc4035901b10667.ContentTypeRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType_id"] = id
    }
    return i666016b9de0fea6c6ca42db4cdc518e2bbe5b22999b3927c4fc4035901b10667.NewContentTypeRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// For drives in SharePoint, the underlying document library list. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *ListRequestBuilder) CreateDeleteRequestInformation(options *ListRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// For drives in SharePoint, the underlying document library list. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *ListRequestBuilder) CreateGetRequestInformation(options *ListRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// For drives in SharePoint, the underlying document library list. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *ListRequestBuilder) CreatePatchRequestInformation(options *ListRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// For drives in SharePoint, the underlying document library list. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *ListRequestBuilder) Delete(options *ListRequestBuilderDeleteOptions)(error) {
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
func (m *ListRequestBuilder) Drive()(*ibba07da9108e84b41a3d09d01bb4a6886a66e86df95f8f0730fa1c61201467ba.DriveRequestBuilder) {
    return ibba07da9108e84b41a3d09d01bb4a6886a66e86df95f8f0730fa1c61201467ba.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// For drives in SharePoint, the underlying document library list. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *ListRequestBuilder) Get(options *ListRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.List, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewList() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.List), nil
}
func (m *ListRequestBuilder) Items()(*ibf8d356e7a219821b2879666575a0d558fe0723087550da01c802cb67c2ea4a9.ItemsRequestBuilder) {
    return ibf8d356e7a219821b2879666575a0d558fe0723087550da01c802cb67c2ea4a9.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.list.items.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ListRequestBuilder) ItemsById(id string)(*if1844020acb9c67e872cff770020bec5e9fbe0ec4bdb729561240d6a686db31c.ListItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItem_id"] = id
    }
    return if1844020acb9c67e872cff770020bec5e9fbe0ec4bdb729561240d6a686db31c.NewListItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// For drives in SharePoint, the underlying document library list. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *ListRequestBuilder) Patch(options *ListRequestBuilderPatchOptions)(error) {
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
func (m *ListRequestBuilder) Subscriptions()(*ibf5d29d5462efb100f3ac502d1eba6060e878b0461baab8d107bba2657e634d5.SubscriptionsRequestBuilder) {
    return ibf5d29d5462efb100f3ac502d1eba6060e878b0461baab8d107bba2657e634d5.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.list.subscriptions.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ListRequestBuilder) SubscriptionsById(id string)(*i0cc711a2d9e9934918c44b76b7f313fb885f4cc911346a17e775fac04e5518ac.SubscriptionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return i0cc711a2d9e9934918c44b76b7f313fb885f4cc911346a17e775fac04e5518ac.NewSubscriptionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
