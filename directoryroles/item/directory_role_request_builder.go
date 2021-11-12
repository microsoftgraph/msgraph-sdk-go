package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i0b751179e4031051436a326b2702d4abde4bb19175d1daaf554c829a18c29177 "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/item/checkmemberobjects"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i718fb7a72b38652faad082566b31b033df2c9f8d04cc0af99ee1e607f87e8d80 "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/item/scopedmembers"
    iaeb18d19aec9445180dbe3b2ba9b76bbfbcf9731777c886b4f0fa5501b97e58c "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/item/checkmembergroups"
    ibcb647a3bce4367afc71872d9ca7c9cc64edcc4e60fe4831e8fcab3de590a294 "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/item/restore"
    id49876e5f70243d663a4f93ec6d77d7f89f0c91f1eda2daad48f4f037781d7b0 "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/item/getmemberobjects"
    id63def69f45b85c0f48fc2880b865fc141de2bfccf786d2d6a9813abb7b316e9 "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/item/getmembergroups"
    if1385e9a44279e575f64436ffd1ff5c93bc83fbaa5a4051782729fbe1d3bdb3d "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/item/members"
    ic73a5746a64610f345c78c4114780b56c25d27fdb8734dbd3b6500ebf1ad197c "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/item/scopedmembers/item"
)

// Builds and executes requests for operations under \directoryRoles\{directoryRole-id}
type DirectoryRoleRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type DirectoryRoleRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type DirectoryRoleRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DirectoryRoleRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Get entity from directoryRoles by key
type DirectoryRoleRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type DirectoryRoleRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DirectoryRole;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DirectoryRoleRequestBuilder) CheckMemberGroups()(*iaeb18d19aec9445180dbe3b2ba9b76bbfbcf9731777c886b4f0fa5501b97e58c.CheckMemberGroupsRequestBuilder) {
    return iaeb18d19aec9445180dbe3b2ba9b76bbfbcf9731777c886b4f0fa5501b97e58c.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRoleRequestBuilder) CheckMemberObjects()(*i0b751179e4031051436a326b2702d4abde4bb19175d1daaf554c829a18c29177.CheckMemberObjectsRequestBuilder) {
    return i0b751179e4031051436a326b2702d4abde4bb19175d1daaf554c829a18c29177.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new DirectoryRoleRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewDirectoryRoleRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRoleRequestBuilder) {
    m := &DirectoryRoleRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directoryRoles/{directoryRole_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new DirectoryRoleRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewDirectoryRoleRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRoleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete entity from directoryRoles
// Parameters:
//  - options : Options for the request
func (m *DirectoryRoleRequestBuilder) CreateDeleteRequestInformation(options *DirectoryRoleRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get entity from directoryRoles by key
// Parameters:
//  - options : Options for the request
func (m *DirectoryRoleRequestBuilder) CreateGetRequestInformation(options *DirectoryRoleRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(options.Q)
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
// Update entity in directoryRoles
// Parameters:
//  - options : Options for the request
func (m *DirectoryRoleRequestBuilder) CreatePatchRequestInformation(options *DirectoryRoleRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete entity from directoryRoles
// Parameters:
//  - options : Options for the request
func (m *DirectoryRoleRequestBuilder) Delete(options *DirectoryRoleRequestBuilderDeleteOptions)(error) {
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
// Get entity from directoryRoles by key
// Parameters:
//  - options : Options for the request
func (m *DirectoryRoleRequestBuilder) Get(options *DirectoryRoleRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DirectoryRole, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewDirectoryRole() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DirectoryRole), nil
}
func (m *DirectoryRoleRequestBuilder) GetMemberGroups()(*id63def69f45b85c0f48fc2880b865fc141de2bfccf786d2d6a9813abb7b316e9.GetMemberGroupsRequestBuilder) {
    return id63def69f45b85c0f48fc2880b865fc141de2bfccf786d2d6a9813abb7b316e9.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRoleRequestBuilder) GetMemberObjects()(*id49876e5f70243d663a4f93ec6d77d7f89f0c91f1eda2daad48f4f037781d7b0.GetMemberObjectsRequestBuilder) {
    return id49876e5f70243d663a4f93ec6d77d7f89f0c91f1eda2daad48f4f037781d7b0.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRoleRequestBuilder) Members()(*if1385e9a44279e575f64436ffd1ff5c93bc83fbaa5a4051782729fbe1d3bdb3d.MembersRequestBuilder) {
    return if1385e9a44279e575f64436ffd1ff5c93bc83fbaa5a4051782729fbe1d3bdb3d.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Update entity in directoryRoles
// Parameters:
//  - options : Options for the request
func (m *DirectoryRoleRequestBuilder) Patch(options *DirectoryRoleRequestBuilderPatchOptions)(error) {
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
func (m *DirectoryRoleRequestBuilder) Restore()(*ibcb647a3bce4367afc71872d9ca7c9cc64edcc4e60fe4831e8fcab3de590a294.RestoreRequestBuilder) {
    return ibcb647a3bce4367afc71872d9ca7c9cc64edcc4e60fe4831e8fcab3de590a294.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRoleRequestBuilder) ScopedMembers()(*i718fb7a72b38652faad082566b31b033df2c9f8d04cc0af99ee1e607f87e8d80.ScopedMembersRequestBuilder) {
    return i718fb7a72b38652faad082566b31b033df2c9f8d04cc0af99ee1e607f87e8d80.NewScopedMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.directoryRoles.item.scopedMembers.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *DirectoryRoleRequestBuilder) ScopedMembersById(id string)(*ic73a5746a64610f345c78c4114780b56c25d27fdb8734dbd3b6500ebf1ad197c.ScopedRoleMembershipRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["scopedRoleMembership_id"] = id
    }
    return ic73a5746a64610f345c78c4114780b56c25d27fdb8734dbd3b6500ebf1ad197c.NewScopedRoleMembershipRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
