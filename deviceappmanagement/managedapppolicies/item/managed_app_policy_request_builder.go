package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i3683ba7627b8cfd708db5823cf7d88bd969d10b57d9160ad4fa9a4c365901758 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedapppolicies/item/windowsinformationprotection"
    ib18aa17be07a21575f3dc1461c70fa5b8c65d4d6a8fd95cc3441984879664e20 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedapppolicies/item/targetedmanagedappprotection"
    ibab033c1a8b48bed14f99450a0356d81369bc62e3c8a8198133c8b5ead2c5d4a "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedapppolicies/item/managedappprotection"
    iece8e7ddbf399451514e3763c1358d8aada93c5bf96f4c2a68c0c6db1146ba76 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedapppolicies/item/targetapps"
)

// Builds and executes requests for operations under \deviceAppManagement\managedAppPolicies\{managedAppPolicy-id}
type ManagedAppPolicyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Managed app policies.
type ManagedAppPolicyRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Instantiates a new ManagedAppPolicyRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewManagedAppPolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedAppPolicyRequestBuilder) {
    m := &ManagedAppPolicyRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/deviceAppManagement/managedAppPolicies/{managedAppPolicy_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ManagedAppPolicyRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewManagedAppPolicyRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedAppPolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedAppPolicyRequestBuilderInternal(urlParams, requestAdapter)
}
// Managed app policies.
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *ManagedAppPolicyRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Managed app policies.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *ManagedAppPolicyRequestBuilder) CreateGetRequestInformation(q func (value *ManagedAppPolicyRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ManagedAppPolicyRequestBuilderGetQueryParameters)
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
// Managed app policies.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *ManagedAppPolicyRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ManagedAppPolicy, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Managed app policies.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *ManagedAppPolicyRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
// Managed app policies.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *ManagedAppPolicyRequestBuilder) Get(q func (value *ManagedAppPolicyRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ManagedAppPolicy, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewManagedAppPolicy() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ManagedAppPolicy), nil
}
func (m *ManagedAppPolicyRequestBuilder) ManagedAppProtection()(*ibab033c1a8b48bed14f99450a0356d81369bc62e3c8a8198133c8b5ead2c5d4a.ManagedAppProtectionRequestBuilder) {
    return ibab033c1a8b48bed14f99450a0356d81369bc62e3c8a8198133c8b5ead2c5d4a.NewManagedAppProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Managed app policies.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *ManagedAppPolicyRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ManagedAppPolicy, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ManagedAppPolicyRequestBuilder) TargetApps()(*iece8e7ddbf399451514e3763c1358d8aada93c5bf96f4c2a68c0c6db1146ba76.TargetAppsRequestBuilder) {
    return iece8e7ddbf399451514e3763c1358d8aada93c5bf96f4c2a68c0c6db1146ba76.NewTargetAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedAppPolicyRequestBuilder) TargetedManagedAppProtection()(*ib18aa17be07a21575f3dc1461c70fa5b8c65d4d6a8fd95cc3441984879664e20.TargetedManagedAppProtectionRequestBuilder) {
    return ib18aa17be07a21575f3dc1461c70fa5b8c65d4d6a8fd95cc3441984879664e20.NewTargetedManagedAppProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedAppPolicyRequestBuilder) WindowsInformationProtection()(*i3683ba7627b8cfd708db5823cf7d88bd969d10b57d9160ad4fa9a4c365901758.WindowsInformationProtectionRequestBuilder) {
    return i3683ba7627b8cfd708db5823cf7d88bd969d10b57d9160ad4fa9a4c365901758.NewWindowsInformationProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
