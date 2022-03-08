package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i0b751179e4031051436a326b2702d4abde4bb19175d1daaf554c829a18c29177 "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/item/checkmemberobjects"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i718fb7a72b38652faad082566b31b033df2c9f8d04cc0af99ee1e607f87e8d80 "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/item/scopedmembers"
    iaeb18d19aec9445180dbe3b2ba9b76bbfbcf9731777c886b4f0fa5501b97e58c "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/item/checkmembergroups"
    ibcb647a3bce4367afc71872d9ca7c9cc64edcc4e60fe4831e8fcab3de590a294 "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/item/restore"
    id49876e5f70243d663a4f93ec6d77d7f89f0c91f1eda2daad48f4f037781d7b0 "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/item/getmemberobjects"
    id63def69f45b85c0f48fc2880b865fc141de2bfccf786d2d6a9813abb7b316e9 "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/item/getmembergroups"
    if1385e9a44279e575f64436ffd1ff5c93bc83fbaa5a4051782729fbe1d3bdb3d "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/item/members"
    i4dc00e738a43cdc3dbc464cba96748a3f629a6e7f36a76cbb15bf379872ebba9 "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/item/members/item"
    ic73a5746a64610f345c78c4114780b56c25d27fdb8734dbd3b6500ebf1ad197c "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/item/scopedmembers/item"
)

// DirectoryRoleItemRequestBuilder provides operations to manage the collection of directoryRole entities.
type DirectoryRoleItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DirectoryRoleItemRequestBuilderDeleteOptions options for Delete
type DirectoryRoleItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DirectoryRoleItemRequestBuilderGetOptions options for Get
type DirectoryRoleItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DirectoryRoleItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DirectoryRoleItemRequestBuilderGetQueryParameters get entity from directoryRoles by key
type DirectoryRoleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DirectoryRoleItemRequestBuilderPatchOptions options for Patch
type DirectoryRoleItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DirectoryRoleable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DirectoryRoleItemRequestBuilder) CheckMemberGroups()(*iaeb18d19aec9445180dbe3b2ba9b76bbfbcf9731777c886b4f0fa5501b97e58c.CheckMemberGroupsRequestBuilder) {
    return iaeb18d19aec9445180dbe3b2ba9b76bbfbcf9731777c886b4f0fa5501b97e58c.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRoleItemRequestBuilder) CheckMemberObjects()(*i0b751179e4031051436a326b2702d4abde4bb19175d1daaf554c829a18c29177.CheckMemberObjectsRequestBuilder) {
    return i0b751179e4031051436a326b2702d4abde4bb19175d1daaf554c829a18c29177.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDirectoryRoleItemRequestBuilderInternal instantiates a new DirectoryRoleItemRequestBuilder and sets the default values.
func NewDirectoryRoleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRoleItemRequestBuilder) {
    m := &DirectoryRoleItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directoryRoles/{directoryRole_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectoryRoleItemRequestBuilder instantiates a new DirectoryRoleItemRequestBuilder and sets the default values.
func NewDirectoryRoleItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRoleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from directoryRoles
func (m *DirectoryRoleItemRequestBuilder) CreateDeleteRequestInformation(options *DirectoryRoleItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from directoryRoles by key
func (m *DirectoryRoleItemRequestBuilder) CreateGetRequestInformation(options *DirectoryRoleItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in directoryRoles
func (m *DirectoryRoleItemRequestBuilder) CreatePatchRequestInformation(options *DirectoryRoleItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete entity from directoryRoles
func (m *DirectoryRoleItemRequestBuilder) Delete(options *DirectoryRoleItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get entity from directoryRoles by key
func (m *DirectoryRoleItemRequestBuilder) Get(options *DirectoryRoleItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DirectoryRoleable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateDirectoryRoleFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DirectoryRoleable), nil
}
func (m *DirectoryRoleItemRequestBuilder) GetMemberGroups()(*id63def69f45b85c0f48fc2880b865fc141de2bfccf786d2d6a9813abb7b316e9.GetMemberGroupsRequestBuilder) {
    return id63def69f45b85c0f48fc2880b865fc141de2bfccf786d2d6a9813abb7b316e9.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRoleItemRequestBuilder) GetMemberObjects()(*id49876e5f70243d663a4f93ec6d77d7f89f0c91f1eda2daad48f4f037781d7b0.GetMemberObjectsRequestBuilder) {
    return id49876e5f70243d663a4f93ec6d77d7f89f0c91f1eda2daad48f4f037781d7b0.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRoleItemRequestBuilder) Members()(*if1385e9a44279e575f64436ffd1ff5c93bc83fbaa5a4051782729fbe1d3bdb3d.MembersRequestBuilder) {
    return if1385e9a44279e575f64436ffd1ff5c93bc83fbaa5a4051782729fbe1d3bdb3d.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.directoryRoles.item.members.item collection
func (m *DirectoryRoleItemRequestBuilder) MembersById(id string)(*i4dc00e738a43cdc3dbc464cba96748a3f629a6e7f36a76cbb15bf379872ebba9.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return i4dc00e738a43cdc3dbc464cba96748a3f629a6e7f36a76cbb15bf379872ebba9.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update entity in directoryRoles
func (m *DirectoryRoleItemRequestBuilder) Patch(options *DirectoryRoleItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *DirectoryRoleItemRequestBuilder) Restore()(*ibcb647a3bce4367afc71872d9ca7c9cc64edcc4e60fe4831e8fcab3de590a294.RestoreRequestBuilder) {
    return ibcb647a3bce4367afc71872d9ca7c9cc64edcc4e60fe4831e8fcab3de590a294.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRoleItemRequestBuilder) ScopedMembers()(*i718fb7a72b38652faad082566b31b033df2c9f8d04cc0af99ee1e607f87e8d80.ScopedMembersRequestBuilder) {
    return i718fb7a72b38652faad082566b31b033df2c9f8d04cc0af99ee1e607f87e8d80.NewScopedMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScopedMembersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.directoryRoles.item.scopedMembers.item collection
func (m *DirectoryRoleItemRequestBuilder) ScopedMembersById(id string)(*ic73a5746a64610f345c78c4114780b56c25d27fdb8734dbd3b6500ebf1ad197c.ScopedRoleMembershipItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["scopedRoleMembership_id"] = id
    }
    return ic73a5746a64610f345c78c4114780b56c25d27fdb8734dbd3b6500ebf1ad197c.NewScopedRoleMembershipItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
