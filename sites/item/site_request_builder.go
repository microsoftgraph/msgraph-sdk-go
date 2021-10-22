package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i05bca8d67f58c6d9aa648acf49fac24e4d5361e451fb444c3868fc782faa3a3a "github.com/microsoftgraph/msgraph-sdk-go/sites/item/items"
    i0c26d90ad9ba1101a52431870fbb281c6a4f6b751415aa8750d1294df6fc9e2c "github.com/microsoftgraph/msgraph-sdk-go/sites/item/contenttypes"
    i10cc803300c5d47903ff6a5a1eb5b34159f7eb7927f500c05b63f005fcd7a3b7 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote"
    i1f30e998cc7f16739a41575c86813834123cba8d5d18b460b177e4d1033ca60d "github.com/microsoftgraph/msgraph-sdk-go/sites/item/drive"
    i339cee83997cfcbf197752f780095780d8a49c408ea09318bf2a269c76c80974 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/analytics"
    i3cd7ab6f3d06321d73344309d80fb03ef4dcccd5c3c20fc9e878db6dc8e33e51 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i4ffbbec24206153c780835b7a43a03ed6b772efe99d5161b43d3a4ba10302a0a "github.com/microsoftgraph/msgraph-sdk-go/sites/item/getbypathwithpath"
    i55cb8b4eba146099911d28cd6cbf1c2648018a786384c9ec7c40a8f9baf109f0 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists"
    i5cf45aa21d0025c798a2868eb9d4091cba588f6ea645d747c04aba2a4591fe1b "github.com/microsoftgraph/msgraph-sdk-go/sites/item/permissions"
    i6dae79969c0cc01a1ae54240ed306a0bd1ad56e07a0515dc28cfab53e29bb580 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/sites"
    i939b9e822b399f2c722f94a48564ad5eeba57ffb61b7a071a9ff1916605e56e8 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/getactivitiesbyinterval"
    ic3e3392ffd4cfbeca5a615ab4821d60c9132bc0218c910f3bbb9967204677d15 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/drives"
    ifdd24c47f65a3313024bac4f3f3e730f886d4041edbb2510eda841cd16478038 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/columns"
    i209a587ff4beb5ee9c360ac0d6e4a3c7f62c54965cf6114ee02077fff23801c5 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/columns/item"
    i25b238e1bb04065b78bbf7b8af856ab5c032ddb3bf1f7e6a6dd2660de6f8217a "github.com/microsoftgraph/msgraph-sdk-go/sites/item/sites/item"
    i78f58c637039463a95013be84966d9b8b0483bfd65aa77130e076a7f7381cfd6 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item"
    i9bdc46bee3db40ed2ad5cb8d27cc47b07556e9f9bec575d1a403c32a5cce68ae "github.com/microsoftgraph/msgraph-sdk-go/sites/item/permissions/item"
    ia5bb104926264b975572fae9aaa50ca289316f014c3b08eb8446548e5146dedb "github.com/microsoftgraph/msgraph-sdk-go/sites/item/drives/item"
    id5c99b5d474798a011325ba892c9f2ef398ac5ca9126710bc745e020c22b60fd "github.com/microsoftgraph/msgraph-sdk-go/sites/item/items/item"
    if642ca98ef1cfea5e51e56e879548fde41cf1297f3ce162d09a6224e659bc1f8 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/contenttypes/item"
)

type SiteRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type SiteRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *SiteRequestBuilder) Analytics()(*i339cee83997cfcbf197752f780095780d8a49c408ea09318bf2a269c76c80974.AnalyticsRequestBuilder) {
    return i339cee83997cfcbf197752f780095780d8a49c408ea09318bf2a269c76c80974.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SiteRequestBuilder) Columns()(*ifdd24c47f65a3313024bac4f3f3e730f886d4041edbb2510eda841cd16478038.ColumnsRequestBuilder) {
    return ifdd24c47f65a3313024bac4f3f3e730f886d4041edbb2510eda841cd16478038.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SiteRequestBuilder) ColumnsById(id string)(*i209a587ff4beb5ee9c360ac0d6e4a3c7f62c54965cf6114ee02077fff23801c5.ColumnDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i209a587ff4beb5ee9c360ac0d6e4a3c7f62c54965cf6114ee02077fff23801c5.NewColumnDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewSiteRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SiteRequestBuilder) {
    m := &SiteRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/sites/{site_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewSiteRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SiteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSiteRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *SiteRequestBuilder) ContentTypes()(*i0c26d90ad9ba1101a52431870fbb281c6a4f6b751415aa8750d1294df6fc9e2c.ContentTypesRequestBuilder) {
    return i0c26d90ad9ba1101a52431870fbb281c6a4f6b751415aa8750d1294df6fc9e2c.NewContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SiteRequestBuilder) ContentTypesById(id string)(*if642ca98ef1cfea5e51e56e879548fde41cf1297f3ce162d09a6224e659bc1f8.ContentTypeRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["contentType_id"] = id
    }
    return if642ca98ef1cfea5e51e56e879548fde41cf1297f3ce162d09a6224e659bc1f8.NewContentTypeRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SiteRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *SiteRequestBuilder) CreateGetRequestInformation(q func (value *SiteRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(SiteRequestBuilderGetQueryParameters)
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
func (m *SiteRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Site, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *SiteRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *SiteRequestBuilder) Drive()(*i1f30e998cc7f16739a41575c86813834123cba8d5d18b460b177e4d1033ca60d.DriveRequestBuilder) {
    return i1f30e998cc7f16739a41575c86813834123cba8d5d18b460b177e4d1033ca60d.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SiteRequestBuilder) Drives()(*ic3e3392ffd4cfbeca5a615ab4821d60c9132bc0218c910f3bbb9967204677d15.DrivesRequestBuilder) {
    return ic3e3392ffd4cfbeca5a615ab4821d60c9132bc0218c910f3bbb9967204677d15.NewDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SiteRequestBuilder) DrivesById(id string)(*ia5bb104926264b975572fae9aaa50ca289316f014c3b08eb8446548e5146dedb.DriveRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["drive_id"] = id
    }
    return ia5bb104926264b975572fae9aaa50ca289316f014c3b08eb8446548e5146dedb.NewDriveRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SiteRequestBuilder) Get(q func (value *SiteRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Site, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewSite() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Site), nil
}
func (m *SiteRequestBuilder) GetActivitiesByInterval()(*i939b9e822b399f2c722f94a48564ad5eeba57ffb61b7a071a9ff1916605e56e8.GetActivitiesByIntervalRequestBuilder) {
    return i939b9e822b399f2c722f94a48564ad5eeba57ffb61b7a071a9ff1916605e56e8.NewGetActivitiesByIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SiteRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(startDateTime *string, endDateTime *string, interval *string)(*i3cd7ab6f3d06321d73344309d80fb03ef4dcccd5c3c20fc9e878db6dc8e33e51.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return i3cd7ab6f3d06321d73344309d80fb03ef4dcccd5c3c20fc9e878db6dc8e33e51.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, startDateTime, endDateTime, interval);
}
func (m *SiteRequestBuilder) GetByPathWithPath(path *string)(*i4ffbbec24206153c780835b7a43a03ed6b772efe99d5161b43d3a4ba10302a0a.GetByPathWithPathRequestBuilder) {
    return i4ffbbec24206153c780835b7a43a03ed6b772efe99d5161b43d3a4ba10302a0a.NewGetByPathWithPathRequestBuilderInternal(m.pathParameters, m.requestAdapter, path);
}
func (m *SiteRequestBuilder) Items()(*i05bca8d67f58c6d9aa648acf49fac24e4d5361e451fb444c3868fc782faa3a3a.ItemsRequestBuilder) {
    return i05bca8d67f58c6d9aa648acf49fac24e4d5361e451fb444c3868fc782faa3a3a.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SiteRequestBuilder) ItemsById(id string)(*id5c99b5d474798a011325ba892c9f2ef398ac5ca9126710bc745e020c22b60fd.BaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["baseItem_id"] = id
    }
    return id5c99b5d474798a011325ba892c9f2ef398ac5ca9126710bc745e020c22b60fd.NewBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SiteRequestBuilder) Lists()(*i55cb8b4eba146099911d28cd6cbf1c2648018a786384c9ec7c40a8f9baf109f0.ListsRequestBuilder) {
    return i55cb8b4eba146099911d28cd6cbf1c2648018a786384c9ec7c40a8f9baf109f0.NewListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SiteRequestBuilder) ListsById(id string)(*i78f58c637039463a95013be84966d9b8b0483bfd65aa77130e076a7f7381cfd6.ListRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["list_id"] = id
    }
    return i78f58c637039463a95013be84966d9b8b0483bfd65aa77130e076a7f7381cfd6.NewListRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SiteRequestBuilder) Onenote()(*i10cc803300c5d47903ff6a5a1eb5b34159f7eb7927f500c05b63f005fcd7a3b7.OnenoteRequestBuilder) {
    return i10cc803300c5d47903ff6a5a1eb5b34159f7eb7927f500c05b63f005fcd7a3b7.NewOnenoteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SiteRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Site, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *SiteRequestBuilder) Permissions()(*i5cf45aa21d0025c798a2868eb9d4091cba588f6ea645d747c04aba2a4591fe1b.PermissionsRequestBuilder) {
    return i5cf45aa21d0025c798a2868eb9d4091cba588f6ea645d747c04aba2a4591fe1b.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SiteRequestBuilder) PermissionsById(id string)(*i9bdc46bee3db40ed2ad5cb8d27cc47b07556e9f9bec575d1a403c32a5cce68ae.PermissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return i9bdc46bee3db40ed2ad5cb8d27cc47b07556e9f9bec575d1a403c32a5cce68ae.NewPermissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SiteRequestBuilder) Sites()(*i6dae79969c0cc01a1ae54240ed306a0bd1ad56e07a0515dc28cfab53e29bb580.SitesRequestBuilder) {
    return i6dae79969c0cc01a1ae54240ed306a0bd1ad56e07a0515dc28cfab53e29bb580.NewSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SiteRequestBuilder) SitesById(id string)(*i25b238e1bb04065b78bbf7b8af856ab5c032ddb3bf1f7e6a6dd2660de6f8217a.SiteRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["site_id1"] = id
    }
    return i25b238e1bb04065b78bbf7b8af856ab5c032ddb3bf1f7e6a6dd2660de6f8217a.NewSiteRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
