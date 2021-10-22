package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i20a49af14049cce13bbcbf77639ba345f23a979ba735144f96caf7cc46ac52d4 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileappconfigurations/item/userstatuses"
    i6bbd09e352a0e87235e9c9775845f897603974356216a5236acfbea5f16155b0 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileappconfigurations/item/assignments"
    i8d9da7a408610e6b2da98617dc04ec0862fb37929855ca63ba17948a3e7c6869 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileappconfigurations/item/devicestatussummary"
    ia2a6efbe6de075696db4a436ca6bc9745df0eb596dd5de9281c86bf2bebcd940 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileappconfigurations/item/assign"
    id36e539f78070cadf5b4f1fcefe8eeff25266761b7e3541b45788010e5cdfb9b "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileappconfigurations/item/userstatussummary"
    iebd1b47b377fdbd0e32c261a4ebc852aebb3266e58bde42f93309b4e2edc6d16 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileappconfigurations/item/devicestatuses"
    i5644f6cefb6e37ba63ec8582d3a158ec48759832b4146d5ffd2d350288b5e5ca "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileappconfigurations/item/devicestatuses/item"
    i72a4318a7e8f28adce62f7520318129090631e70fa23a9371561c34fac1f0582 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileappconfigurations/item/userstatuses/item"
    i88ff03a861b0c98567d5cfa6d93d8d9d0de34bb73fdd948e55765d9858a275c3 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileappconfigurations/item/assignments/item"
)

type ManagedDeviceMobileAppConfigurationRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type ManagedDeviceMobileAppConfigurationRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) Assign()(*ia2a6efbe6de075696db4a436ca6bc9745df0eb596dd5de9281c86bf2bebcd940.AssignRequestBuilder) {
    return ia2a6efbe6de075696db4a436ca6bc9745df0eb596dd5de9281c86bf2bebcd940.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) Assignments()(*i6bbd09e352a0e87235e9c9775845f897603974356216a5236acfbea5f16155b0.AssignmentsRequestBuilder) {
    return i6bbd09e352a0e87235e9c9775845f897603974356216a5236acfbea5f16155b0.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) AssignmentsById(id string)(*i88ff03a861b0c98567d5cfa6d93d8d9d0de34bb73fdd948e55765d9858a275c3.ManagedDeviceMobileAppConfigurationAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationAssignment_id"] = id
    }
    return i88ff03a861b0c98567d5cfa6d93d8d9d0de34bb73fdd948e55765d9858a275c3.NewManagedDeviceMobileAppConfigurationAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewManagedDeviceMobileAppConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDeviceMobileAppConfigurationRequestBuilder) {
    m := &ManagedDeviceMobileAppConfigurationRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/deviceAppManagement/mobileAppConfigurations/{managedDeviceMobileAppConfiguration_id}{?select,expand}";
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
func NewManagedDeviceMobileAppConfigurationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDeviceMobileAppConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDeviceMobileAppConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) CreateGetRequestInformation(q func (value *ManagedDeviceMobileAppConfigurationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ManagedDeviceMobileAppConfigurationRequestBuilderGetQueryParameters)
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
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ManagedDeviceMobileAppConfiguration, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) DeviceStatuses()(*iebd1b47b377fdbd0e32c261a4ebc852aebb3266e58bde42f93309b4e2edc6d16.DeviceStatusesRequestBuilder) {
    return iebd1b47b377fdbd0e32c261a4ebc852aebb3266e58bde42f93309b4e2edc6d16.NewDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) DeviceStatusesById(id string)(*i5644f6cefb6e37ba63ec8582d3a158ec48759832b4146d5ffd2d350288b5e5ca.ManagedDeviceMobileAppConfigurationDeviceStatusRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationDeviceStatus_id"] = id
    }
    return i5644f6cefb6e37ba63ec8582d3a158ec48759832b4146d5ffd2d350288b5e5ca.NewManagedDeviceMobileAppConfigurationDeviceStatusRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) DeviceStatusSummary()(*i8d9da7a408610e6b2da98617dc04ec0862fb37929855ca63ba17948a3e7c6869.DeviceStatusSummaryRequestBuilder) {
    return i8d9da7a408610e6b2da98617dc04ec0862fb37929855ca63ba17948a3e7c6869.NewDeviceStatusSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) Get(q func (value *ManagedDeviceMobileAppConfigurationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ManagedDeviceMobileAppConfiguration, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewManagedDeviceMobileAppConfiguration() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ManagedDeviceMobileAppConfiguration), nil
}
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ManagedDeviceMobileAppConfiguration, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) UserStatuses()(*i20a49af14049cce13bbcbf77639ba345f23a979ba735144f96caf7cc46ac52d4.UserStatusesRequestBuilder) {
    return i20a49af14049cce13bbcbf77639ba345f23a979ba735144f96caf7cc46ac52d4.NewUserStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) UserStatusesById(id string)(*i72a4318a7e8f28adce62f7520318129090631e70fa23a9371561c34fac1f0582.ManagedDeviceMobileAppConfigurationUserStatusRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationUserStatus_id"] = id
    }
    return i72a4318a7e8f28adce62f7520318129090631e70fa23a9371561c34fac1f0582.NewManagedDeviceMobileAppConfigurationUserStatusRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) UserStatusSummary()(*id36e539f78070cadf5b4f1fcefe8eeff25266761b7e3541b45788010e5cdfb9b.UserStatusSummaryRequestBuilder) {
    return id36e539f78070cadf5b4f1fcefe8eeff25266761b7e3541b45788010e5cdfb9b.NewUserStatusSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
