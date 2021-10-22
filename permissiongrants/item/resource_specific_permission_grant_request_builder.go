package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a7063fe0d09620e5d6522f433c1e9cf3e5d296667e87f9dfecfc4af1b3aae9d "github.com/microsoftgraph/msgraph-sdk-go/permissiongrants/item/getmemberobjects"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i8eeb73da3d89f8a760cd4adc6d53215b2c688043311f9d987c3f034cc4014646 "github.com/microsoftgraph/msgraph-sdk-go/permissiongrants/item/checkmembergroups"
    ia2efcc0d1f3eacb9f6edf91b52e6f227f398341a4af363e1ddc708ad6cade478 "github.com/microsoftgraph/msgraph-sdk-go/permissiongrants/item/restore"
    ia62a23d20d63f22d8eaaf57757dfb7d2ffbe1f8f51f38621c112384a083cf3a8 "github.com/microsoftgraph/msgraph-sdk-go/permissiongrants/item/getmembergroups"
    if0a1607357c7ac08b0f7369dd31ca3dae44d9d88c78c88afa398d08a8a0fe97e "github.com/microsoftgraph/msgraph-sdk-go/permissiongrants/item/checkmemberobjects"
)

type ResourceSpecificPermissionGrantRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type ResourceSpecificPermissionGrantRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *ResourceSpecificPermissionGrantRequestBuilder) CheckMemberGroups()(*i8eeb73da3d89f8a760cd4adc6d53215b2c688043311f9d987c3f034cc4014646.CheckMemberGroupsRequestBuilder) {
    return i8eeb73da3d89f8a760cd4adc6d53215b2c688043311f9d987c3f034cc4014646.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceSpecificPermissionGrantRequestBuilder) CheckMemberObjects()(*if0a1607357c7ac08b0f7369dd31ca3dae44d9d88c78c88afa398d08a8a0fe97e.CheckMemberObjectsRequestBuilder) {
    return if0a1607357c7ac08b0f7369dd31ca3dae44d9d88c78c88afa398d08a8a0fe97e.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewResourceSpecificPermissionGrantRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ResourceSpecificPermissionGrantRequestBuilder) {
    m := &ResourceSpecificPermissionGrantRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/permissionGrants/{resourceSpecificPermissionGrant_id}{?select,expand}";
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
func NewResourceSpecificPermissionGrantRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ResourceSpecificPermissionGrantRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewResourceSpecificPermissionGrantRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ResourceSpecificPermissionGrantRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ResourceSpecificPermissionGrantRequestBuilder) CreateGetRequestInformation(q func (value *ResourceSpecificPermissionGrantRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ResourceSpecificPermissionGrantRequestBuilderGetQueryParameters)
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
func (m *ResourceSpecificPermissionGrantRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ResourceSpecificPermissionGrant, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ResourceSpecificPermissionGrantRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ResourceSpecificPermissionGrantRequestBuilder) Get(q func (value *ResourceSpecificPermissionGrantRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ResourceSpecificPermissionGrant, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewResourceSpecificPermissionGrant() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ResourceSpecificPermissionGrant), nil
}
func (m *ResourceSpecificPermissionGrantRequestBuilder) GetMemberGroups()(*ia62a23d20d63f22d8eaaf57757dfb7d2ffbe1f8f51f38621c112384a083cf3a8.GetMemberGroupsRequestBuilder) {
    return ia62a23d20d63f22d8eaaf57757dfb7d2ffbe1f8f51f38621c112384a083cf3a8.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceSpecificPermissionGrantRequestBuilder) GetMemberObjects()(*i4a7063fe0d09620e5d6522f433c1e9cf3e5d296667e87f9dfecfc4af1b3aae9d.GetMemberObjectsRequestBuilder) {
    return i4a7063fe0d09620e5d6522f433c1e9cf3e5d296667e87f9dfecfc4af1b3aae9d.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceSpecificPermissionGrantRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ResourceSpecificPermissionGrant, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ResourceSpecificPermissionGrantRequestBuilder) Restore()(*ia2efcc0d1f3eacb9f6edf91b52e6f227f398341a4af363e1ddc708ad6cade478.RestoreRequestBuilder) {
    return ia2efcc0d1f3eacb9f6edf91b52e6f227f398341a4af363e1ddc708ad6cade478.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
