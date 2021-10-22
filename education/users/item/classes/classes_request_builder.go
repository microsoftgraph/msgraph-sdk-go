package classes

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i356788470e8b5dfebc108d91aa675a76dc77b9a6b30cb13f018abab6327e9aae "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/classes/delta"
    id3aa218f65fd74b2e00676c38fb1cdabdfd0681f7c5e3504a3a3fdb452b26bf1 "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/classes/ref"
)

type ClassesRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type ClassesRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Count *bool;
    Expand []string;
    Filter *string;
    Orderby []string;
    Search *string;
    Select_escpaped []string;
    Skip *int32;
    Top *int32;
}
func NewClassesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ClassesRequestBuilder) {
    m := &ClassesRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/education/users/{educationUser_id}/classes{?top,skip,search,filter,count,orderby,select,expand}";
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
func NewClassesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ClassesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewClassesRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ClassesRequestBuilder) CreateGetRequestInformation(q func (value *ClassesRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ClassesRequestBuilderGetQueryParameters)
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
func (m *ClassesRequestBuilder) Delta()(*i356788470e8b5dfebc108d91aa675a76dc77b9a6b30cb13f018abab6327e9aae.DeltaRequestBuilder) {
    return i356788470e8b5dfebc108d91aa675a76dc77b9a6b30cb13f018abab6327e9aae.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ClassesRequestBuilder) Get(q func (value *ClassesRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*ClassesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewClassesResponse() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*ClassesResponse), nil
}
func (m *ClassesRequestBuilder) Ref()(*id3aa218f65fd74b2e00676c38fb1cdabdfd0681f7c5e3504a3a3fdb452b26bf1.RefRequestBuilder) {
    return id3aa218f65fd74b2e00676c38fb1cdabdfd0681f7c5e3504a3a3fdb452b26bf1.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
