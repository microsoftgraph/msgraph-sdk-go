package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0527cdc9fd204c09fc094d9f1e7b55f4ccd84d958a7e0489f836475992dd87bb "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/iosmanagedappprotections/item/apps"
    i6ae3cac28f103d10b9e38003e0d0f3659395bfb3860d2aa342dff5dcd0e09b03 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/iosmanagedappprotections/item/deploymentsummary"
    i55f341823424fe64c759539b60c8dd7b2233b0db784b5c0c209dbdcc827b89fc "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/iosmanagedappprotections/item/apps/item"
)

// IosManagedAppProtectionRequestBuilder builds and executes requests for operations under \deviceAppManagement\iosManagedAppProtections\{iosManagedAppProtection-id}
type IosManagedAppProtectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// IosManagedAppProtectionRequestBuilderDeleteOptions options for Delete
type IosManagedAppProtectionRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// IosManagedAppProtectionRequestBuilderGetOptions options for Get
type IosManagedAppProtectionRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *IosManagedAppProtectionRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// IosManagedAppProtectionRequestBuilderGetQueryParameters iOS managed app policies.
type IosManagedAppProtectionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// IosManagedAppProtectionRequestBuilderPatchOptions options for Patch
type IosManagedAppProtectionRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.IosManagedAppProtection;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *IosManagedAppProtectionRequestBuilder) Apps()(*i0527cdc9fd204c09fc094d9f1e7b55f4ccd84d958a7e0489f836475992dd87bb.AppsRequestBuilder) {
    return i0527cdc9fd204c09fc094d9f1e7b55f4ccd84d958a7e0489f836475992dd87bb.NewAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.iosManagedAppProtections.item.apps.item collection
func (m *IosManagedAppProtectionRequestBuilder) AppsById(id string)(*i55f341823424fe64c759539b60c8dd7b2233b0db784b5c0c209dbdcc827b89fc.ManagedMobileAppRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedMobileApp_id"] = id
    }
    return i55f341823424fe64c759539b60c8dd7b2233b0db784b5c0c209dbdcc827b89fc.NewManagedMobileAppRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewIosManagedAppProtectionRequestBuilderInternal instantiates a new IosManagedAppProtectionRequestBuilder and sets the default values.
func NewIosManagedAppProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IosManagedAppProtectionRequestBuilder) {
    m := &IosManagedAppProtectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/iosManagedAppProtections/{iosManagedAppProtection_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIosManagedAppProtectionRequestBuilder instantiates a new IosManagedAppProtectionRequestBuilder and sets the default values.
func NewIosManagedAppProtectionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IosManagedAppProtectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIosManagedAppProtectionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation iOS managed app policies.
func (m *IosManagedAppProtectionRequestBuilder) CreateDeleteRequestInformation(options *IosManagedAppProtectionRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation iOS managed app policies.
func (m *IosManagedAppProtectionRequestBuilder) CreateGetRequestInformation(options *IosManagedAppProtectionRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation iOS managed app policies.
func (m *IosManagedAppProtectionRequestBuilder) CreatePatchRequestInformation(options *IosManagedAppProtectionRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete iOS managed app policies.
func (m *IosManagedAppProtectionRequestBuilder) Delete(options *IosManagedAppProtectionRequestBuilderDeleteOptions)(error) {
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
func (m *IosManagedAppProtectionRequestBuilder) DeploymentSummary()(*i6ae3cac28f103d10b9e38003e0d0f3659395bfb3860d2aa342dff5dcd0e09b03.DeploymentSummaryRequestBuilder) {
    return i6ae3cac28f103d10b9e38003e0d0f3659395bfb3860d2aa342dff5dcd0e09b03.NewDeploymentSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get iOS managed app policies.
func (m *IosManagedAppProtectionRequestBuilder) Get(options *IosManagedAppProtectionRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.IosManagedAppProtection, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewIosManagedAppProtection() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.IosManagedAppProtection), nil
}
// Patch iOS managed app policies.
func (m *IosManagedAppProtectionRequestBuilder) Patch(options *IosManagedAppProtectionRequestBuilderPatchOptions)(error) {
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
