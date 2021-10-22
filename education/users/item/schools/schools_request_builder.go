package schools

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i31c55a4d38e851a2de14c96cf6d2fede7b768f082f53a4448032a19bba7c7c1e "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/schools/delta"
    i59640742c9e94ac289bae1f030f8b678df7249c049506acce82ced5002851d43 "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/schools/ref"
)

type SchoolsRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type SchoolsRequestBuilderGetQueryParameters struct {
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
func NewSchoolsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SchoolsRequestBuilder) {
    m := &SchoolsRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/education/users/{educationUser_id}/schools{?top,skip,search,filter,count,orderby,select,expand}";
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
func NewSchoolsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SchoolsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSchoolsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *SchoolsRequestBuilder) CreateGetRequestInformation(q func (value *SchoolsRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(SchoolsRequestBuilderGetQueryParameters)
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
func (m *SchoolsRequestBuilder) Delta()(*i31c55a4d38e851a2de14c96cf6d2fede7b768f082f53a4448032a19bba7c7c1e.DeltaRequestBuilder) {
    return i31c55a4d38e851a2de14c96cf6d2fede7b768f082f53a4448032a19bba7c7c1e.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SchoolsRequestBuilder) Get(q func (value *SchoolsRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*SchoolsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSchoolsResponse() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*SchoolsResponse), nil
}
func (m *SchoolsRequestBuilder) Ref()(*i59640742c9e94ac289bae1f030f8b678df7249c049506acce82ced5002851d43.RefRequestBuilder) {
    return i59640742c9e94ac289bae1f030f8b678df7249c049506acce82ced5002851d43.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
