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

// OrgContactItemRequestBuilder builds and executes requests for operations under \contacts\{orgContact-id}
type OrgContactItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OrgContactItemRequestBuilderDeleteOptions options for Delete
type OrgContactItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OrgContactItemRequestBuilderGetOptions options for Get
type OrgContactItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OrgContactItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OrgContactItemRequestBuilderGetQueryParameters get entity from contacts by key
type OrgContactItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// OrgContactItemRequestBuilderPatchOptions options for Patch
type OrgContactItemRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OrgContact;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *OrgContactItemRequestBuilder) CheckMemberGroups()(*id7379fe1d66c0786481cb810e571e9cf8ca3a1db19fb8596ae3ab26b168ebd4b.CheckMemberGroupsRequestBuilder) {
    return id7379fe1d66c0786481cb810e571e9cf8ca3a1db19fb8596ae3ab26b168ebd4b.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrgContactItemRequestBuilder) CheckMemberObjects()(*ib08660edc0c07cb1bfb954212b3de0dfac964791cee0a520dea1ff1fc58ce970.CheckMemberObjectsRequestBuilder) {
    return ib08660edc0c07cb1bfb954212b3de0dfac964791cee0a520dea1ff1fc58ce970.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewOrgContactItemRequestBuilderInternal instantiates a new OrgContactItemRequestBuilder and sets the default values.
func NewOrgContactItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OrgContactItemRequestBuilder) {
    m := &OrgContactItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/contacts/{orgContact_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOrgContactItemRequestBuilder instantiates a new OrgContactItemRequestBuilder and sets the default values.
func NewOrgContactItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OrgContactItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOrgContactItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from contacts
func (m *OrgContactItemRequestBuilder) CreateDeleteRequestInformation(options *OrgContactItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from contacts by key
func (m *OrgContactItemRequestBuilder) CreateGetRequestInformation(options *OrgContactItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in contacts
func (m *OrgContactItemRequestBuilder) CreatePatchRequestInformation(options *OrgContactItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete entity from contacts
func (m *OrgContactItemRequestBuilder) Delete(options *OrgContactItemRequestBuilderDeleteOptions)(error) {
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
func (m *OrgContactItemRequestBuilder) DirectReports()(*i055ad2f09c6431b30e6d3528a3dd11b667dcf169a3943e92ebbdc9f40c102910.DirectReportsRequestBuilder) {
    return i055ad2f09c6431b30e6d3528a3dd11b667dcf169a3943e92ebbdc9f40c102910.NewDirectReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get entity from contacts by key
func (m *OrgContactItemRequestBuilder) Get(options *OrgContactItemRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OrgContact, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewOrgContact() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OrgContact), nil
}
func (m *OrgContactItemRequestBuilder) GetMemberGroups()(*i15c3213b1d34c79d320391089a05683e39eb479cb5d4b030a5b53be13ac05712.GetMemberGroupsRequestBuilder) {
    return i15c3213b1d34c79d320391089a05683e39eb479cb5d4b030a5b53be13ac05712.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrgContactItemRequestBuilder) GetMemberObjects()(*ifff3f2949bd5013659a326ca1cde357b44996411a8f259036c89d4ee51051f34.GetMemberObjectsRequestBuilder) {
    return ifff3f2949bd5013659a326ca1cde357b44996411a8f259036c89d4ee51051f34.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrgContactItemRequestBuilder) Manager()(*i6a99a368f031e6c1f338c5fc4bcbdad85d242df8d3d5dfe1b7b0475df468b3dd.ManagerRequestBuilder) {
    return i6a99a368f031e6c1f338c5fc4bcbdad85d242df8d3d5dfe1b7b0475df468b3dd.NewManagerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrgContactItemRequestBuilder) MemberOf()(*i4830f34009f28f830e2d6c1e84624d4bbfc5a6d035e6f1097558162d8a77592d.MemberOfRequestBuilder) {
    return i4830f34009f28f830e2d6c1e84624d4bbfc5a6d035e6f1097558162d8a77592d.NewMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in contacts
func (m *OrgContactItemRequestBuilder) Patch(options *OrgContactItemRequestBuilderPatchOptions)(error) {
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
func (m *OrgContactItemRequestBuilder) Restore()(*i56f89ff65e41535b67a714e8fdea36f3ad34961ef4ad2e4c91a286c72ea09b44.RestoreRequestBuilder) {
    return i56f89ff65e41535b67a714e8fdea36f3ad34961ef4ad2e4c91a286c72ea09b44.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrgContactItemRequestBuilder) TransitiveMemberOf()(*i79406b1f55eac2f939b44b7a8f42f4c1fcc6945c0eb2e6ff7fb5ee4c06796b4d.TransitiveMemberOfRequestBuilder) {
    return i79406b1f55eac2f939b44b7a8f42f4c1fcc6945c0eb2e6ff7fb5ee4c06796b4d.NewTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
