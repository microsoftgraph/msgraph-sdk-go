package list

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i6cc54c104a1f556165cc44589ee14207379ed8c37fa85729d5c314e13f2214d7 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/columns"
    i7a0c6962c8b1589339031c077d13646778071c443927dc1046bc04df31627bf6 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/items"
    i8b1352d9ad171bd1e78c354f011b512c0faed7a8b44a2d0af7a87499c0e9f548 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes"
    iaad175a704d2fb8c88af7ffd0f8ebd7bfd4f12f115e3413d8b3c39946957d581 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/subscriptions"
    id412f117e08220004924a69577571fc52ee563dcb4aa3bb6992790f4a56a8672 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/drive"
    i6f7c6ed0403d181f0f9df70f035fc911fc0235b004fca66648da29a67534a2e2 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item"
    ia2c40a1ff086507585f951d28494158116a4fcb5951ee49a77ded46131865f37 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/subscriptions/item"
    ia5941405950ea17199c6d656adb76e8bf88f3f605415a64747a31ce95e0c07b7 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/items/item"
    ifd1fbca74ca99c59f8151a23008189eb5fa814dfb2bdb94e0592368b94d84ebe "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/columns/item"
)

// ListRequestBuilder builds and executes requests for operations under \shares\{sharedDriveItem-id}\list
type ListRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ListRequestBuilderDeleteOptions options for Delete
type ListRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ListRequestBuilderGetOptions options for Get
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
// ListRequestBuilderGetQueryParameters used to access the underlying list
type ListRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ListRequestBuilderPatchOptions options for Patch
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
func (m *ListRequestBuilder) Columns()(*i6cc54c104a1f556165cc44589ee14207379ed8c37fa85729d5c314e13f2214d7.ColumnsRequestBuilder) {
    return i6cc54c104a1f556165cc44589ee14207379ed8c37fa85729d5c314e13f2214d7.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.shares.item.list.columns.item collection
func (m *ListRequestBuilder) ColumnsById(id string)(*ifd1fbca74ca99c59f8151a23008189eb5fa814dfb2bdb94e0592368b94d84ebe.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return ifd1fbca74ca99c59f8151a23008189eb5fa814dfb2bdb94e0592368b94d84ebe.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewListRequestBuilderInternal instantiates a new ListRequestBuilder and sets the default values.
func NewListRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListRequestBuilder) {
    m := &ListRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/shares/{sharedDriveItem_id}/list{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewListRequestBuilder instantiates a new ListRequestBuilder and sets the default values.
func NewListRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewListRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ListRequestBuilder) ContentTypes()(*i8b1352d9ad171bd1e78c354f011b512c0faed7a8b44a2d0af7a87499c0e9f548.ContentTypesRequestBuilder) {
    return i8b1352d9ad171bd1e78c354f011b512c0faed7a8b44a2d0af7a87499c0e9f548.NewContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContentTypesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.shares.item.list.contentTypes.item collection
func (m *ListRequestBuilder) ContentTypesById(id string)(*i6f7c6ed0403d181f0f9df70f035fc911fc0235b004fca66648da29a67534a2e2.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType_id"] = id
    }
    return i6f7c6ed0403d181f0f9df70f035fc911fc0235b004fca66648da29a67534a2e2.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation used to access the underlying list
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
// CreateGetRequestInformation used to access the underlying list
func (m *ListRequestBuilder) CreateGetRequestInformation(options *ListRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation used to access the underlying list
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
// Delete used to access the underlying list
func (m *ListRequestBuilder) Delete(options *ListRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ListRequestBuilder) Drive()(*id412f117e08220004924a69577571fc52ee563dcb4aa3bb6992790f4a56a8672.DriveRequestBuilder) {
    return id412f117e08220004924a69577571fc52ee563dcb4aa3bb6992790f4a56a8672.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get used to access the underlying list
func (m *ListRequestBuilder) Get(options *ListRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.List, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewList() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.List), nil
}
func (m *ListRequestBuilder) Items()(*i7a0c6962c8b1589339031c077d13646778071c443927dc1046bc04df31627bf6.ItemsRequestBuilder) {
    return i7a0c6962c8b1589339031c077d13646778071c443927dc1046bc04df31627bf6.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.shares.item.list.items.item collection
func (m *ListRequestBuilder) ItemsById(id string)(*ia5941405950ea17199c6d656adb76e8bf88f3f605415a64747a31ce95e0c07b7.ListItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItem_id"] = id
    }
    return ia5941405950ea17199c6d656adb76e8bf88f3f605415a64747a31ce95e0c07b7.NewListItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch used to access the underlying list
func (m *ListRequestBuilder) Patch(options *ListRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ListRequestBuilder) Subscriptions()(*iaad175a704d2fb8c88af7ffd0f8ebd7bfd4f12f115e3413d8b3c39946957d581.SubscriptionsRequestBuilder) {
    return iaad175a704d2fb8c88af7ffd0f8ebd7bfd4f12f115e3413d8b3c39946957d581.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.shares.item.list.subscriptions.item collection
func (m *ListRequestBuilder) SubscriptionsById(id string)(*ia2c40a1ff086507585f951d28494158116a4fcb5951ee49a77ded46131865f37.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return ia2c40a1ff086507585f951d28494158116a4fcb5951ee49a77ded46131865f37.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
