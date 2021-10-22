package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0035901ffc213fc4b59f4993863d1daec79b347f97e62b284d79c7f4e51e541a "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/childfolders"
    ib939870f831b5d1e6f646bd6624ff59df0efd3d63e40ad5d6874229710fb5a0f "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/contacts"
    ibe30f52eff3d7782e030050a6a70ab017912ac09c1eaaafeb1f5bc12d8b3a5c0 "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/singlevalueextendedproperties"
    ic1c22c979ab4e96bd99861c36b59000262714ad01fb89b523d57edd0e2710294 "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/multivalueextendedproperties"
    i6249342797ad051a9fde0138fc7a4de2b582b010423ef26c264a2a10b971def7 "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/multivalueextendedproperties/item"
    i73617fc20b6e8870c7f350b9a031e29042cc9e456837fc4b6df34bbf3918f396 "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/contacts/item"
    i7b320e3aa4ff049100b60cd68585037930958ce6c760db10289e1a7baef49df1 "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/childfolders/item"
    i82ebac0048f2c4d4678fe54780699f0f1230c24f422393c29fcc9bccf43bf443 "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/singlevalueextendedproperties/item"
)

type ContactFolderRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type ContactFolderRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *ContactFolderRequestBuilder) ChildFolders()(*i0035901ffc213fc4b59f4993863d1daec79b347f97e62b284d79c7f4e51e541a.ChildFoldersRequestBuilder) {
    return i0035901ffc213fc4b59f4993863d1daec79b347f97e62b284d79c7f4e51e541a.NewChildFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContactFolderRequestBuilder) ChildFoldersById(id string)(*i7b320e3aa4ff049100b60cd68585037930958ce6c760db10289e1a7baef49df1.ContactFolderRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["contactFolder_id1"] = id
    }
    return i7b320e3aa4ff049100b60cd68585037930958ce6c760db10289e1a7baef49df1.NewContactFolderRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewContactFolderRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContactFolderRequestBuilder) {
    m := &ContactFolderRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/contactFolders/{contactFolder_id}{?select,expand}";
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
func NewContactFolderRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContactFolderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContactFolderRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ContactFolderRequestBuilder) Contacts()(*ib939870f831b5d1e6f646bd6624ff59df0efd3d63e40ad5d6874229710fb5a0f.ContactsRequestBuilder) {
    return ib939870f831b5d1e6f646bd6624ff59df0efd3d63e40ad5d6874229710fb5a0f.NewContactsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContactFolderRequestBuilder) ContactsById(id string)(*i73617fc20b6e8870c7f350b9a031e29042cc9e456837fc4b6df34bbf3918f396.ContactRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["contact_id"] = id
    }
    return i73617fc20b6e8870c7f350b9a031e29042cc9e456837fc4b6df34bbf3918f396.NewContactRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContactFolderRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ContactFolderRequestBuilder) CreateGetRequestInformation(q func (value *ContactFolderRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ContactFolderRequestBuilderGetQueryParameters)
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
func (m *ContactFolderRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ContactFolder, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ContactFolderRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ContactFolderRequestBuilder) Get(q func (value *ContactFolderRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ContactFolder, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewContactFolder() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ContactFolder), nil
}
func (m *ContactFolderRequestBuilder) MultiValueExtendedProperties()(*ic1c22c979ab4e96bd99861c36b59000262714ad01fb89b523d57edd0e2710294.MultiValueExtendedPropertiesRequestBuilder) {
    return ic1c22c979ab4e96bd99861c36b59000262714ad01fb89b523d57edd0e2710294.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContactFolderRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i6249342797ad051a9fde0138fc7a4de2b582b010423ef26c264a2a10b971def7.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i6249342797ad051a9fde0138fc7a4de2b582b010423ef26c264a2a10b971def7.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContactFolderRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ContactFolder, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ContactFolderRequestBuilder) SingleValueExtendedProperties()(*ibe30f52eff3d7782e030050a6a70ab017912ac09c1eaaafeb1f5bc12d8b3a5c0.SingleValueExtendedPropertiesRequestBuilder) {
    return ibe30f52eff3d7782e030050a6a70ab017912ac09c1eaaafeb1f5bc12d8b3a5c0.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContactFolderRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i82ebac0048f2c4d4678fe54780699f0f1230c24f422393c29fcc9bccf43bf443.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i82ebac0048f2c4d4678fe54780699f0f1230c24f422393c29fcc9bccf43bf443.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
