package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i23081904cae97cc914ee1e96256b8b201e9e89aae88137434ed8befc80c525b4 "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roleassignments/item/roledefinition"
    i517df9bb1eb43ab5f2a7c664e7b284783f8e5b9ebdc5a4f713fea7a6238655fa "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roleassignments/item/principal"
    ic4c721a6468a7f4e9de2a19e2983ef41cf8472cc41c3aef1dcd61d1ce5a3e997 "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roleassignments/item/appscope"
    ie74dee05c5a2bc30388fa3bc35898754f6666a2e7747f4d074fce1ef16796a68 "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roleassignments/item/directoryscope"
)

// UnifiedRoleAssignmentRequestBuilder builds and executes requests for operations under \roleManagement\entitlementManagement\roleAssignments\{unifiedRoleAssignment-id}
type UnifiedRoleAssignmentRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UnifiedRoleAssignmentRequestBuilderDeleteOptions options for Delete
type UnifiedRoleAssignmentRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UnifiedRoleAssignmentRequestBuilderGetOptions options for Get
type UnifiedRoleAssignmentRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UnifiedRoleAssignmentRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UnifiedRoleAssignmentRequestBuilderGetQueryParameters resource to grant access to users or groups.
type UnifiedRoleAssignmentRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// UnifiedRoleAssignmentRequestBuilderPatchOptions options for Patch
type UnifiedRoleAssignmentRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.UnifiedRoleAssignment;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *UnifiedRoleAssignmentRequestBuilder) AppScope()(*ic4c721a6468a7f4e9de2a19e2983ef41cf8472cc41c3aef1dcd61d1ce5a3e997.AppScopeRequestBuilder) {
    return ic4c721a6468a7f4e9de2a19e2983ef41cf8472cc41c3aef1dcd61d1ce5a3e997.NewAppScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUnifiedRoleAssignmentRequestBuilderInternal instantiates a new UnifiedRoleAssignmentRequestBuilder and sets the default values.
func NewUnifiedRoleAssignmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnifiedRoleAssignmentRequestBuilder) {
    m := &UnifiedRoleAssignmentRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/entitlementManagement/roleAssignments/{unifiedRoleAssignment_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUnifiedRoleAssignmentRequestBuilder instantiates a new UnifiedRoleAssignmentRequestBuilder and sets the default values.
func NewUnifiedRoleAssignmentRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnifiedRoleAssignmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUnifiedRoleAssignmentRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation resource to grant access to users or groups.
func (m *UnifiedRoleAssignmentRequestBuilder) CreateDeleteRequestInformation(options *UnifiedRoleAssignmentRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation resource to grant access to users or groups.
func (m *UnifiedRoleAssignmentRequestBuilder) CreateGetRequestInformation(options *UnifiedRoleAssignmentRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation resource to grant access to users or groups.
func (m *UnifiedRoleAssignmentRequestBuilder) CreatePatchRequestInformation(options *UnifiedRoleAssignmentRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete resource to grant access to users or groups.
func (m *UnifiedRoleAssignmentRequestBuilder) Delete(options *UnifiedRoleAssignmentRequestBuilderDeleteOptions)(error) {
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
func (m *UnifiedRoleAssignmentRequestBuilder) DirectoryScope()(*ie74dee05c5a2bc30388fa3bc35898754f6666a2e7747f4d074fce1ef16796a68.DirectoryScopeRequestBuilder) {
    return ie74dee05c5a2bc30388fa3bc35898754f6666a2e7747f4d074fce1ef16796a68.NewDirectoryScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get resource to grant access to users or groups.
func (m *UnifiedRoleAssignmentRequestBuilder) Get(options *UnifiedRoleAssignmentRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.UnifiedRoleAssignment, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewUnifiedRoleAssignment() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.UnifiedRoleAssignment), nil
}
// Patch resource to grant access to users or groups.
func (m *UnifiedRoleAssignmentRequestBuilder) Patch(options *UnifiedRoleAssignmentRequestBuilderPatchOptions)(error) {
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
func (m *UnifiedRoleAssignmentRequestBuilder) Principal()(*i517df9bb1eb43ab5f2a7c664e7b284783f8e5b9ebdc5a4f713fea7a6238655fa.PrincipalRequestBuilder) {
    return i517df9bb1eb43ab5f2a7c664e7b284783f8e5b9ebdc5a4f713fea7a6238655fa.NewPrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UnifiedRoleAssignmentRequestBuilder) RoleDefinition()(*i23081904cae97cc914ee1e96256b8b201e9e89aae88137434ed8befc80c525b4.RoleDefinitionRequestBuilder) {
    return i23081904cae97cc914ee1e96256b8b201e9e89aae88137434ed8befc80c525b4.NewRoleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
