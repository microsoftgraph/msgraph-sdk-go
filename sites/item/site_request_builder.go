package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i05bca8d67f58c6d9aa648acf49fac24e4d5361e451fb444c3868fc782faa3a3a "github.com/microsoftgraph/msgraph-sdk-go/sites/item/items"
    i0c26d90ad9ba1101a52431870fbb281c6a4f6b751415aa8750d1294df6fc9e2c "github.com/microsoftgraph/msgraph-sdk-go/sites/item/contenttypes"
    i10cc803300c5d47903ff6a5a1eb5b34159f7eb7927f500c05b63f005fcd7a3b7 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote"
    i1f30e998cc7f16739a41575c86813834123cba8d5d18b460b177e4d1033ca60d "github.com/microsoftgraph/msgraph-sdk-go/sites/item/drive"
    i339cee83997cfcbf197752f780095780d8a49c408ea09318bf2a269c76c80974 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/analytics"
    i367b3d6454445f7875bc897535d25edb404746299a8839bb9f87747de65f43fd "github.com/microsoftgraph/msgraph-sdk-go/sites/item/externalcolumns"
    i3ca97aef887804bcab28c6fc0da670ecf38816601cb616e2dd9231cb1cc3a18b "github.com/microsoftgraph/msgraph-sdk-go/sites/item/getapplicablecontenttypesforlistwithlistid"
    i3cd7ab6f3d06321d73344309d80fb03ef4dcccd5c3c20fc9e878db6dc8e33e51 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i4ffbbec24206153c780835b7a43a03ed6b772efe99d5161b43d3a4ba10302a0a "github.com/microsoftgraph/msgraph-sdk-go/sites/item/getbypathwithpath"
    i55cb8b4eba146099911d28cd6cbf1c2648018a786384c9ec7c40a8f9baf109f0 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists"
    i56d3808dd8956f908cd3292b0ec9af6c42c18061cbc621b53136a2fa051b6542 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstore"
    i5cf45aa21d0025c798a2868eb9d4091cba588f6ea645d747c04aba2a4591fe1b "github.com/microsoftgraph/msgraph-sdk-go/sites/item/permissions"
    i6dae79969c0cc01a1ae54240ed306a0bd1ad56e07a0515dc28cfab53e29bb580 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/sites"
    i8c3520c945a3499d24747186abcb37377ea0d8afa3d1bfa0cd8a7e6fcf9b2a83 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores"
    i939b9e822b399f2c722f94a48564ad5eeba57ffb61b7a071a9ff1916605e56e8 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/getactivitiesbyinterval"
    ic3e3392ffd4cfbeca5a615ab4821d60c9132bc0218c910f3bbb9967204677d15 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/drives"
    ifdd24c47f65a3313024bac4f3f3e730f886d4041edbb2510eda841cd16478038 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/columns"
    i209a587ff4beb5ee9c360ac0d6e4a3c7f62c54965cf6114ee02077fff23801c5 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/columns/item"
    i25b238e1bb04065b78bbf7b8af856ab5c032ddb3bf1f7e6a6dd2660de6f8217a "github.com/microsoftgraph/msgraph-sdk-go/sites/item/sites/item"
    i78f58c637039463a95013be84966d9b8b0483bfd65aa77130e076a7f7381cfd6 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item"
    i9bdc46bee3db40ed2ad5cb8d27cc47b07556e9f9bec575d1a403c32a5cce68ae "github.com/microsoftgraph/msgraph-sdk-go/sites/item/permissions/item"
    ia4734d1ec9f14b1703e1b22f66af2a51aea0b15668f690fb492176e062cae20e "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item"
    ia5bb104926264b975572fae9aaa50ca289316f014c3b08eb8446548e5146dedb "github.com/microsoftgraph/msgraph-sdk-go/sites/item/drives/item"
    id5c99b5d474798a011325ba892c9f2ef398ac5ca9126710bc745e020c22b60fd "github.com/microsoftgraph/msgraph-sdk-go/sites/item/items/item"
    if642ca98ef1cfea5e51e56e879548fde41cf1297f3ce162d09a6224e659bc1f8 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/contenttypes/item"
)

// Builds and executes requests for operations under \sites\{site-id}
type SiteRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type SiteRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type SiteRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SiteRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Get entity from sites by key
type SiteRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type SiteRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Site;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *SiteRequestBuilder) Analytics()(*i339cee83997cfcbf197752f780095780d8a49c408ea09318bf2a269c76c80974.AnalyticsRequestBuilder) {
    return i339cee83997cfcbf197752f780095780d8a49c408ea09318bf2a269c76c80974.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SiteRequestBuilder) Columns()(*ifdd24c47f65a3313024bac4f3f3e730f886d4041edbb2510eda841cd16478038.ColumnsRequestBuilder) {
    return ifdd24c47f65a3313024bac4f3f3e730f886d4041edbb2510eda841cd16478038.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.sites.item.columns.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *SiteRequestBuilder) ColumnsById(id string)(*i209a587ff4beb5ee9c360ac0d6e4a3c7f62c54965cf6114ee02077fff23801c5.ColumnDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i209a587ff4beb5ee9c360ac0d6e4a3c7f62c54965cf6114ee02077fff23801c5.NewColumnDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new SiteRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewSiteRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SiteRequestBuilder) {
    m := &SiteRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/sites/{site_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new SiteRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewSiteRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SiteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSiteRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *SiteRequestBuilder) ContentTypes()(*i0c26d90ad9ba1101a52431870fbb281c6a4f6b751415aa8750d1294df6fc9e2c.ContentTypesRequestBuilder) {
    return i0c26d90ad9ba1101a52431870fbb281c6a4f6b751415aa8750d1294df6fc9e2c.NewContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.sites.item.contentTypes.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *SiteRequestBuilder) ContentTypesById(id string)(*if642ca98ef1cfea5e51e56e879548fde41cf1297f3ce162d09a6224e659bc1f8.ContentTypeRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType_id"] = id
    }
    return if642ca98ef1cfea5e51e56e879548fde41cf1297f3ce162d09a6224e659bc1f8.NewContentTypeRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete entity from sites
// Parameters:
//  - options : Options for the request
func (m *SiteRequestBuilder) CreateDeleteRequestInformation(options *SiteRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get entity from sites by key
// Parameters:
//  - options : Options for the request
func (m *SiteRequestBuilder) CreateGetRequestInformation(options *SiteRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Update entity in sites
// Parameters:
//  - options : Options for the request
func (m *SiteRequestBuilder) CreatePatchRequestInformation(options *SiteRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete entity from sites
// Parameters:
//  - options : Options for the request
func (m *SiteRequestBuilder) Delete(options *SiteRequestBuilderDeleteOptions)(error) {
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
func (m *SiteRequestBuilder) Drive()(*i1f30e998cc7f16739a41575c86813834123cba8d5d18b460b177e4d1033ca60d.DriveRequestBuilder) {
    return i1f30e998cc7f16739a41575c86813834123cba8d5d18b460b177e4d1033ca60d.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SiteRequestBuilder) Drives()(*ic3e3392ffd4cfbeca5a615ab4821d60c9132bc0218c910f3bbb9967204677d15.DrivesRequestBuilder) {
    return ic3e3392ffd4cfbeca5a615ab4821d60c9132bc0218c910f3bbb9967204677d15.NewDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.sites.item.drives.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *SiteRequestBuilder) DrivesById(id string)(*ia5bb104926264b975572fae9aaa50ca289316f014c3b08eb8446548e5146dedb.DriveRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["drive_id"] = id
    }
    return ia5bb104926264b975572fae9aaa50ca289316f014c3b08eb8446548e5146dedb.NewDriveRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SiteRequestBuilder) ExternalColumns()(*i367b3d6454445f7875bc897535d25edb404746299a8839bb9f87747de65f43fd.ExternalColumnsRequestBuilder) {
    return i367b3d6454445f7875bc897535d25edb404746299a8839bb9f87747de65f43fd.NewExternalColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get entity from sites by key
// Parameters:
//  - options : Options for the request
func (m *SiteRequestBuilder) Get(options *SiteRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Site, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewSite() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Site), nil
}
// Builds and executes requests for operations under \sites\{site-id}\microsoft.graph.getActivitiesByInterval()
func (m *SiteRequestBuilder) GetActivitiesByInterval()(*i939b9e822b399f2c722f94a48564ad5eeba57ffb61b7a071a9ff1916605e56e8.GetActivitiesByIntervalRequestBuilder) {
    return i939b9e822b399f2c722f94a48564ad5eeba57ffb61b7a071a9ff1916605e56e8.NewGetActivitiesByIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \sites\{site-id}\microsoft.graph.getActivitiesByInterval(startDateTime='{startDateTime}',endDateTime='{endDateTime}',interval='{interval}')
// Parameters:
//  - endDateTime : Usage: endDateTime={endDateTime}
//  - interval : Usage: interval={interval}
//  - startDateTime : Usage: startDateTime={startDateTime}
func (m *SiteRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(startDateTime *string, endDateTime *string, interval *string)(*i3cd7ab6f3d06321d73344309d80fb03ef4dcccd5c3c20fc9e878db6dc8e33e51.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return i3cd7ab6f3d06321d73344309d80fb03ef4dcccd5c3c20fc9e878db6dc8e33e51.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, startDateTime, endDateTime, interval);
}
// Builds and executes requests for operations under \sites\{site-id}\microsoft.graph.getApplicableContentTypesForList(listId='{listId}')
// Parameters:
//  - listId : Usage: listId={listId}
func (m *SiteRequestBuilder) GetApplicableContentTypesForListWithListId(listId *string)(*i3ca97aef887804bcab28c6fc0da670ecf38816601cb616e2dd9231cb1cc3a18b.GetApplicableContentTypesForListWithListIdRequestBuilder) {
    return i3ca97aef887804bcab28c6fc0da670ecf38816601cb616e2dd9231cb1cc3a18b.NewGetApplicableContentTypesForListWithListIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, listId);
}
// Builds and executes requests for operations under \sites\{site-id}\microsoft.graph.getByPath(path='{path}')
// Parameters:
//  - path : Usage: path={path}
func (m *SiteRequestBuilder) GetByPathWithPath(path *string)(*i4ffbbec24206153c780835b7a43a03ed6b772efe99d5161b43d3a4ba10302a0a.GetByPathWithPathRequestBuilder) {
    return i4ffbbec24206153c780835b7a43a03ed6b772efe99d5161b43d3a4ba10302a0a.NewGetByPathWithPathRequestBuilderInternal(m.pathParameters, m.requestAdapter, path);
}
func (m *SiteRequestBuilder) Items()(*i05bca8d67f58c6d9aa648acf49fac24e4d5361e451fb444c3868fc782faa3a3a.ItemsRequestBuilder) {
    return i05bca8d67f58c6d9aa648acf49fac24e4d5361e451fb444c3868fc782faa3a3a.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.sites.item.items.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *SiteRequestBuilder) ItemsById(id string)(*id5c99b5d474798a011325ba892c9f2ef398ac5ca9126710bc745e020c22b60fd.BaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["baseItem_id"] = id
    }
    return id5c99b5d474798a011325ba892c9f2ef398ac5ca9126710bc745e020c22b60fd.NewBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SiteRequestBuilder) Lists()(*i55cb8b4eba146099911d28cd6cbf1c2648018a786384c9ec7c40a8f9baf109f0.ListsRequestBuilder) {
    return i55cb8b4eba146099911d28cd6cbf1c2648018a786384c9ec7c40a8f9baf109f0.NewListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.sites.item.lists.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *SiteRequestBuilder) ListsById(id string)(*i78f58c637039463a95013be84966d9b8b0483bfd65aa77130e076a7f7381cfd6.ListRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["list_id"] = id
    }
    return i78f58c637039463a95013be84966d9b8b0483bfd65aa77130e076a7f7381cfd6.NewListRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SiteRequestBuilder) Onenote()(*i10cc803300c5d47903ff6a5a1eb5b34159f7eb7927f500c05b63f005fcd7a3b7.OnenoteRequestBuilder) {
    return i10cc803300c5d47903ff6a5a1eb5b34159f7eb7927f500c05b63f005fcd7a3b7.NewOnenoteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Update entity in sites
// Parameters:
//  - options : Options for the request
func (m *SiteRequestBuilder) Patch(options *SiteRequestBuilderPatchOptions)(error) {
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
func (m *SiteRequestBuilder) Permissions()(*i5cf45aa21d0025c798a2868eb9d4091cba588f6ea645d747c04aba2a4591fe1b.PermissionsRequestBuilder) {
    return i5cf45aa21d0025c798a2868eb9d4091cba588f6ea645d747c04aba2a4591fe1b.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.sites.item.permissions.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *SiteRequestBuilder) PermissionsById(id string)(*i9bdc46bee3db40ed2ad5cb8d27cc47b07556e9f9bec575d1a403c32a5cce68ae.PermissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return i9bdc46bee3db40ed2ad5cb8d27cc47b07556e9f9bec575d1a403c32a5cce68ae.NewPermissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SiteRequestBuilder) Sites()(*i6dae79969c0cc01a1ae54240ed306a0bd1ad56e07a0515dc28cfab53e29bb580.SitesRequestBuilder) {
    return i6dae79969c0cc01a1ae54240ed306a0bd1ad56e07a0515dc28cfab53e29bb580.NewSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.sites.item.sites.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *SiteRequestBuilder) SitesById(id string)(*i25b238e1bb04065b78bbf7b8af856ab5c032ddb3bf1f7e6a6dd2660de6f8217a.SiteRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["site_id1"] = id
    }
    return i25b238e1bb04065b78bbf7b8af856ab5c032ddb3bf1f7e6a6dd2660de6f8217a.NewSiteRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SiteRequestBuilder) TermStore()(*i56d3808dd8956f908cd3292b0ec9af6c42c18061cbc621b53136a2fa051b6542.TermStoreRequestBuilder) {
    return i56d3808dd8956f908cd3292b0ec9af6c42c18061cbc621b53136a2fa051b6542.NewTermStoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SiteRequestBuilder) TermStores()(*i8c3520c945a3499d24747186abcb37377ea0d8afa3d1bfa0cd8a7e6fcf9b2a83.TermStoresRequestBuilder) {
    return i8c3520c945a3499d24747186abcb37377ea0d8afa3d1bfa0cd8a7e6fcf9b2a83.NewTermStoresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.sites.item.termStores.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *SiteRequestBuilder) TermStoresById(id string)(*ia4734d1ec9f14b1703e1b22f66af2a51aea0b15668f690fb492176e062cae20e.StoreRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["store_id"] = id
    }
    return ia4734d1ec9f14b1703e1b22f66af2a51aea0b15668f690fb492176e062cae20e.NewStoreRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
