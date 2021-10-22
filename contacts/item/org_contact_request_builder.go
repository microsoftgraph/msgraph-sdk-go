package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i055ad2f09c6431b30e6d3528a3dd11b667dcf169a3943e92ebbdc9f40c102910 "github.com/microsoftgraph/msgraph-sdk-go/contacts/item/directreports"
    i15c3213b1d34c79d320391089a05683e39eb479cb5d4b030a5b53be13ac05712 "github.com/microsoftgraph/msgraph-sdk-go/contacts/item/getmembergroups"
    i4830f34009f28f830e2d6c1e84624d4bbfc5a6d035e6f1097558162d8a77592d "github.com/microsoftgraph/msgraph-sdk-go/contacts/item/memberof"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i56f89ff65e41535b67a714e8fdea36f3ad34961ef4ad2e4c91a286c72ea09b44 "github.com/microsoftgraph/msgraph-sdk-go/contacts/item/restore"
    i6a99a368f031e6c1f338c5fc4bcbdad85d242df8d3d5dfe1b7b0475df468b3dd "github.com/microsoftgraph/msgraph-sdk-go/contacts/item/manager"
    i79406b1f55eac2f939b44b7a8f42f4c1fcc6945c0eb2e6ff7fb5ee4c06796b4d "github.com/microsoftgraph/msgraph-sdk-go/contacts/item/transitivememberof"
    ib08660edc0c07cb1bfb954212b3de0dfac964791cee0a520dea1ff1fc58ce970 "github.com/microsoftgraph/msgraph-sdk-go/contacts/item/checkmemberobjects"
    id7379fe1d66c0786481cb810e571e9cf8ca3a1db19fb8596ae3ab26b168ebd4b "github.com/microsoftgraph/msgraph-sdk-go/contacts/item/checkmembergroups"
    ifff3f2949bd5013659a326ca1cde357b44996411a8f259036c89d4ee51051f34 "github.com/microsoftgraph/msgraph-sdk-go/contacts/item/getmemberobjects"
)

type OrgContactRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type OrgContactRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *OrgContactRequestBuilder) CheckMemberGroups()(*id7379fe1d66c0786481cb810e571e9cf8ca3a1db19fb8596ae3ab26b168ebd4b.CheckMemberGroupsRequestBuilder) {
    return id7379fe1d66c0786481cb810e571e9cf8ca3a1db19fb8596ae3ab26b168ebd4b.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrgContactRequestBuilder) CheckMemberObjects()(*ib08660edc0c07cb1bfb954212b3de0dfac964791cee0a520dea1ff1fc58ce970.CheckMemberObjectsRequestBuilder) {
    return ib08660edc0c07cb1bfb954212b3de0dfac964791cee0a520dea1ff1fc58ce970.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewOrgContactRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OrgContactRequestBuilder) {
    m := &OrgContactRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/contacts/{orgContact_id}{?select,expand}";
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
func NewOrgContactRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OrgContactRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOrgContactRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *OrgContactRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *OrgContactRequestBuilder) CreateGetRequestInformation(q func (value *OrgContactRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(OrgContactRequestBuilderGetQueryParameters)
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
func (m *OrgContactRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OrgContact, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *OrgContactRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *OrgContactRequestBuilder) DirectReports()(*i055ad2f09c6431b30e6d3528a3dd11b667dcf169a3943e92ebbdc9f40c102910.DirectReportsRequestBuilder) {
    return i055ad2f09c6431b30e6d3528a3dd11b667dcf169a3943e92ebbdc9f40c102910.NewDirectReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrgContactRequestBuilder) Get(q func (value *OrgContactRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OrgContact, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewOrgContact() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OrgContact), nil
}
func (m *OrgContactRequestBuilder) GetMemberGroups()(*i15c3213b1d34c79d320391089a05683e39eb479cb5d4b030a5b53be13ac05712.GetMemberGroupsRequestBuilder) {
    return i15c3213b1d34c79d320391089a05683e39eb479cb5d4b030a5b53be13ac05712.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrgContactRequestBuilder) GetMemberObjects()(*ifff3f2949bd5013659a326ca1cde357b44996411a8f259036c89d4ee51051f34.GetMemberObjectsRequestBuilder) {
    return ifff3f2949bd5013659a326ca1cde357b44996411a8f259036c89d4ee51051f34.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrgContactRequestBuilder) Manager()(*i6a99a368f031e6c1f338c5fc4bcbdad85d242df8d3d5dfe1b7b0475df468b3dd.ManagerRequestBuilder) {
    return i6a99a368f031e6c1f338c5fc4bcbdad85d242df8d3d5dfe1b7b0475df468b3dd.NewManagerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrgContactRequestBuilder) MemberOf()(*i4830f34009f28f830e2d6c1e84624d4bbfc5a6d035e6f1097558162d8a77592d.MemberOfRequestBuilder) {
    return i4830f34009f28f830e2d6c1e84624d4bbfc5a6d035e6f1097558162d8a77592d.NewMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrgContactRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OrgContact, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *OrgContactRequestBuilder) Restore()(*i56f89ff65e41535b67a714e8fdea36f3ad34961ef4ad2e4c91a286c72ea09b44.RestoreRequestBuilder) {
    return i56f89ff65e41535b67a714e8fdea36f3ad34961ef4ad2e4c91a286c72ea09b44.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrgContactRequestBuilder) TransitiveMemberOf()(*i79406b1f55eac2f939b44b7a8f42f4c1fcc6945c0eb2e6ff7fb5ee4c06796b4d.TransitiveMemberOfRequestBuilder) {
    return i79406b1f55eac2f939b44b7a8f42f4c1fcc6945c0eb2e6ff7fb5ee4c06796b4d.NewTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
