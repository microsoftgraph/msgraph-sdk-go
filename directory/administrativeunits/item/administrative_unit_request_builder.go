package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i776e11034305b48e5290b2e53dd2f7575359c486273d18f2353ff65faa9063b1 "github.com/microsoftgraph/msgraph-sdk-go/directory/administrativeunits/item/scopedrolemembers"
    i94253098475db605e990260e13febd658ee5c0709ceaf573e3f903be950c52b6 "github.com/microsoftgraph/msgraph-sdk-go/directory/administrativeunits/item/extensions"
    ia454d456d99bb0ead0e1bfb50c53f817055b0e68ca4356296e7f406ee25d3d77 "github.com/microsoftgraph/msgraph-sdk-go/directory/administrativeunits/item/members"
    ia9df9f379ea5fc92e0383f29dfd050c3b88591b8f475f586b856fc130a4fd053 "github.com/microsoftgraph/msgraph-sdk-go/directory/administrativeunits/item/scopedrolemembers/item"
    if71e90a642a2c3f3e90bf7d0f3cbe5d719d6ca30c2312110465fc6e587f8ec38 "github.com/microsoftgraph/msgraph-sdk-go/directory/administrativeunits/item/extensions/item"
)

type AdministrativeUnitRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type AdministrativeUnitRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func NewAdministrativeUnitRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AdministrativeUnitRequestBuilder) {
    m := &AdministrativeUnitRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/directory/administrativeUnits/{administrativeUnit_id}{?select,expand}";
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
func NewAdministrativeUnitRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AdministrativeUnitRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdministrativeUnitRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *AdministrativeUnitRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *AdministrativeUnitRequestBuilder) CreateGetRequestInformation(q func (value *AdministrativeUnitRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(AdministrativeUnitRequestBuilderGetQueryParameters)
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
func (m *AdministrativeUnitRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AdministrativeUnit, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *AdministrativeUnitRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *AdministrativeUnitRequestBuilder) Extensions()(*i94253098475db605e990260e13febd658ee5c0709ceaf573e3f903be950c52b6.ExtensionsRequestBuilder) {
    return i94253098475db605e990260e13febd658ee5c0709ceaf573e3f903be950c52b6.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AdministrativeUnitRequestBuilder) ExtensionsById(id string)(*if71e90a642a2c3f3e90bf7d0f3cbe5d719d6ca30c2312110465fc6e587f8ec38.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return if71e90a642a2c3f3e90bf7d0f3cbe5d719d6ca30c2312110465fc6e587f8ec38.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AdministrativeUnitRequestBuilder) Get(q func (value *AdministrativeUnitRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AdministrativeUnit, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewAdministrativeUnit() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AdministrativeUnit), nil
}
func (m *AdministrativeUnitRequestBuilder) Members()(*ia454d456d99bb0ead0e1bfb50c53f817055b0e68ca4356296e7f406ee25d3d77.MembersRequestBuilder) {
    return ia454d456d99bb0ead0e1bfb50c53f817055b0e68ca4356296e7f406ee25d3d77.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AdministrativeUnitRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AdministrativeUnit, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *AdministrativeUnitRequestBuilder) ScopedRoleMembers()(*i776e11034305b48e5290b2e53dd2f7575359c486273d18f2353ff65faa9063b1.ScopedRoleMembersRequestBuilder) {
    return i776e11034305b48e5290b2e53dd2f7575359c486273d18f2353ff65faa9063b1.NewScopedRoleMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AdministrativeUnitRequestBuilder) ScopedRoleMembersById(id string)(*ia9df9f379ea5fc92e0383f29dfd050c3b88591b8f475f586b856fc130a4fd053.ScopedRoleMembershipRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["scopedRoleMembership_id"] = id
    }
    return ia9df9f379ea5fc92e0383f29dfd050c3b88591b8f475f586b856fc130a4fd053.NewScopedRoleMembershipRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
