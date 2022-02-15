package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i414bcfa81cf62cb492a10891e6ade220e793a207338da61bbf717173d599c787 "github.com/microsoftgraph/msgraph-sdk-go/directoryroletemplates/item/checkmemberobjects"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    iaa01455e2185fdd184cd4993202a2e1f3eec805d38e4ba431cce0c3144cca8f4 "github.com/microsoftgraph/msgraph-sdk-go/directoryroletemplates/item/getmemberobjects"
    ic28c47ff94ea6e8ffffd08e5644e712f6344e2bc49deb132bfac7a2d04ceb662 "github.com/microsoftgraph/msgraph-sdk-go/directoryroletemplates/item/restore"
    iec1a8ab976accbeb3945c4eeab0a0c4f560960c9c1f9f5ebd942d29059b40629 "github.com/microsoftgraph/msgraph-sdk-go/directoryroletemplates/item/getmembergroups"
    iff1e45e5e14f79eaaea630c1948f4ffdd5e56ce522e6fc233cd011734be4178b "github.com/microsoftgraph/msgraph-sdk-go/directoryroletemplates/item/checkmembergroups"
)

// DirectoryRoleTemplateRequestBuilder builds and executes requests for operations under \directoryRoleTemplates\{directoryRoleTemplate-id}
type DirectoryRoleTemplateRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DirectoryRoleTemplateRequestBuilderDeleteOptions options for Delete
type DirectoryRoleTemplateRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DirectoryRoleTemplateRequestBuilderGetOptions options for Get
type DirectoryRoleTemplateRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DirectoryRoleTemplateRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DirectoryRoleTemplateRequestBuilderGetQueryParameters get entity from directoryRoleTemplates by key
type DirectoryRoleTemplateRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DirectoryRoleTemplateRequestBuilderPatchOptions options for Patch
type DirectoryRoleTemplateRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DirectoryRoleTemplate;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DirectoryRoleTemplateRequestBuilder) CheckMemberGroups()(*iff1e45e5e14f79eaaea630c1948f4ffdd5e56ce522e6fc233cd011734be4178b.CheckMemberGroupsRequestBuilder) {
    return iff1e45e5e14f79eaaea630c1948f4ffdd5e56ce522e6fc233cd011734be4178b.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRoleTemplateRequestBuilder) CheckMemberObjects()(*i414bcfa81cf62cb492a10891e6ade220e793a207338da61bbf717173d599c787.CheckMemberObjectsRequestBuilder) {
    return i414bcfa81cf62cb492a10891e6ade220e793a207338da61bbf717173d599c787.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDirectoryRoleTemplateRequestBuilderInternal instantiates a new DirectoryRoleTemplateRequestBuilder and sets the default values.
func NewDirectoryRoleTemplateRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRoleTemplateRequestBuilder) {
    m := &DirectoryRoleTemplateRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directoryRoleTemplates/{directoryRoleTemplate_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectoryRoleTemplateRequestBuilder instantiates a new DirectoryRoleTemplateRequestBuilder and sets the default values.
func NewDirectoryRoleTemplateRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRoleTemplateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleTemplateRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from directoryRoleTemplates
func (m *DirectoryRoleTemplateRequestBuilder) CreateDeleteRequestInformation(options *DirectoryRoleTemplateRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from directoryRoleTemplates by key
func (m *DirectoryRoleTemplateRequestBuilder) CreateGetRequestInformation(options *DirectoryRoleTemplateRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in directoryRoleTemplates
func (m *DirectoryRoleTemplateRequestBuilder) CreatePatchRequestInformation(options *DirectoryRoleTemplateRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete entity from directoryRoleTemplates
func (m *DirectoryRoleTemplateRequestBuilder) Delete(options *DirectoryRoleTemplateRequestBuilderDeleteOptions)(error) {
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
// Get get entity from directoryRoleTemplates by key
func (m *DirectoryRoleTemplateRequestBuilder) Get(options *DirectoryRoleTemplateRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DirectoryRoleTemplate, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewDirectoryRoleTemplate() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DirectoryRoleTemplate), nil
}
func (m *DirectoryRoleTemplateRequestBuilder) GetMemberGroups()(*iec1a8ab976accbeb3945c4eeab0a0c4f560960c9c1f9f5ebd942d29059b40629.GetMemberGroupsRequestBuilder) {
    return iec1a8ab976accbeb3945c4eeab0a0c4f560960c9c1f9f5ebd942d29059b40629.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRoleTemplateRequestBuilder) GetMemberObjects()(*iaa01455e2185fdd184cd4993202a2e1f3eec805d38e4ba431cce0c3144cca8f4.GetMemberObjectsRequestBuilder) {
    return iaa01455e2185fdd184cd4993202a2e1f3eec805d38e4ba431cce0c3144cca8f4.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in directoryRoleTemplates
func (m *DirectoryRoleTemplateRequestBuilder) Patch(options *DirectoryRoleTemplateRequestBuilderPatchOptions)(error) {
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
func (m *DirectoryRoleTemplateRequestBuilder) Restore()(*ic28c47ff94ea6e8ffffd08e5644e712f6344e2bc49deb132bfac7a2d04ceb662.RestoreRequestBuilder) {
    return ic28c47ff94ea6e8ffffd08e5644e712f6344e2bc49deb132bfac7a2d04ceb662.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
