package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i3e25ccdc9a9a70e0ce653da565907342697e58d64e18b519cd2888ef09c7aec2 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/drive"
    ib0bb6c285d2b89517632b84321fc9413604706b96beb6806a4afeeb55142414b "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/subscriptions"
    iba44b7436ee578f0cdcf0bef23ebf91e573df950188cd5c09348cdb272313d22 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/columns"
    ibc5e8f990dad4f95b8b4bd579dded9ea1e8c076148f1d4cfb4f1744ebe20ab8d "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes"
    ie75397ed565a327126aae5b8e7b19c6bc3de18a156dba53f2b72b89e0a7e00b1 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/items"
    i2d9cdb676c9c26e70e774261d0a85daad5586103454cee0b289bf9656e13ffc9 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/items/item"
    i6c9dcd133fd23d5a310dbad4862c8aaa2a6609247eff036993808a9040099f13 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item"
    i9013361a5eaf684747310250468b31adbabe01728c088561aa2e8074c10fa175 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/subscriptions/item"
    if9bcea652e382c241c7fa4514e528bbc37f0207bc686524780f29a74dd5ca836 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/columns/item"
)

// ListRequestBuilder builds and executes requests for operations under \sites\{site-id}\lists\{list-id}
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
// ListRequestBuilderGetQueryParameters the collection of lists under this site.
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
func (m *ListRequestBuilder) Columns()(*iba44b7436ee578f0cdcf0bef23ebf91e573df950188cd5c09348cdb272313d22.ColumnsRequestBuilder) {
    return iba44b7436ee578f0cdcf0bef23ebf91e573df950188cd5c09348cdb272313d22.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.lists.item.columns.item collection
func (m *ListRequestBuilder) ColumnsById(id string)(*if9bcea652e382c241c7fa4514e528bbc37f0207bc686524780f29a74dd5ca836.ColumnDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return if9bcea652e382c241c7fa4514e528bbc37f0207bc686524780f29a74dd5ca836.NewColumnDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewListRequestBuilderInternal instantiates a new ListRequestBuilder and sets the default values.
func NewListRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListRequestBuilder) {
    m := &ListRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site_id}/lists/{list_id}{?select,expand}";
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
func (m *ListRequestBuilder) ContentTypes()(*ibc5e8f990dad4f95b8b4bd579dded9ea1e8c076148f1d4cfb4f1744ebe20ab8d.ContentTypesRequestBuilder) {
    return ibc5e8f990dad4f95b8b4bd579dded9ea1e8c076148f1d4cfb4f1744ebe20ab8d.NewContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContentTypesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.lists.item.contentTypes.item collection
func (m *ListRequestBuilder) ContentTypesById(id string)(*i6c9dcd133fd23d5a310dbad4862c8aaa2a6609247eff036993808a9040099f13.ContentTypeRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType_id"] = id
    }
    return i6c9dcd133fd23d5a310dbad4862c8aaa2a6609247eff036993808a9040099f13.NewContentTypeRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation the collection of lists under this site.
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
// CreateGetRequestInformation the collection of lists under this site.
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
// CreatePatchRequestInformation the collection of lists under this site.
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
// Delete the collection of lists under this site.
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
func (m *ListRequestBuilder) Drive()(*i3e25ccdc9a9a70e0ce653da565907342697e58d64e18b519cd2888ef09c7aec2.DriveRequestBuilder) {
    return i3e25ccdc9a9a70e0ce653da565907342697e58d64e18b519cd2888ef09c7aec2.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the collection of lists under this site.
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
func (m *ListRequestBuilder) Items()(*ie75397ed565a327126aae5b8e7b19c6bc3de18a156dba53f2b72b89e0a7e00b1.ItemsRequestBuilder) {
    return ie75397ed565a327126aae5b8e7b19c6bc3de18a156dba53f2b72b89e0a7e00b1.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.lists.item.items.item collection
func (m *ListRequestBuilder) ItemsById(id string)(*i2d9cdb676c9c26e70e774261d0a85daad5586103454cee0b289bf9656e13ffc9.ListItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItem_id"] = id
    }
    return i2d9cdb676c9c26e70e774261d0a85daad5586103454cee0b289bf9656e13ffc9.NewListItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the collection of lists under this site.
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
func (m *ListRequestBuilder) Subscriptions()(*ib0bb6c285d2b89517632b84321fc9413604706b96beb6806a4afeeb55142414b.SubscriptionsRequestBuilder) {
    return ib0bb6c285d2b89517632b84321fc9413604706b96beb6806a4afeeb55142414b.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.lists.item.subscriptions.item collection
func (m *ListRequestBuilder) SubscriptionsById(id string)(*i9013361a5eaf684747310250468b31adbabe01728c088561aa2e8074c10fa175.SubscriptionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return i9013361a5eaf684747310250468b31adbabe01728c088561aa2e8074c10fa175.NewSubscriptionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
