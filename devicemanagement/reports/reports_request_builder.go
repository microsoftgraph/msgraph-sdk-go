package reports

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i062d4a86380aad1b2a18fcdf2ae3039afe4fb3340e8c61cd8a0118f64e2961f4 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/reports/exportjobs"
    i0a839f3beac17f684c85ef687fb80c09048b3dcae7c1ecbde1af40e48df57f93 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/reports/getdevicemanagementintentsettingsreport"
    i2f4861dbc7fdcc004c73b6f3a193e531630645ae4e3dba4a06598c5ac995763e "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/reports/getcompliancepolicynoncompliancereport"
    i412da535edde14d19b0f7a70e3391905918a35901cecec9071168452048c7acc "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/reports/getpolicynoncompliancereport"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i53f2ba02a8f5254d99ef3d211f9756166f93939a36f554a78d1688617cbe7a5f "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/reports/getreportfilters"
    i63f7c0aabdb80575f559f70509dbec8f47b990899881ff53424a1b846dd52597 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/reports/gethistoricalreport"
    i69e11ae317e8e814a0b7be5423ee2af053bfe7bd422319ac2db1f93ee18dc6f5 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/reports/getsettingnoncompliancereport"
    ia0c903425c40dceef2013d71020f9eecda163cc98d6a70d6643a96e52d28d782 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/reports/getpolicynoncompliancesummaryreport"
    ia8bd370622efb0565c68883e987555ee843b6553925fd83251129c0ab1a330bf "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/reports/getdevicemanagementintentpersettingcontributingprofiles"
    ib215a0b2b7d6a201fa215cdae39964212be5dea6ccfa14879498271f5d9820d1 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/reports/getconfigurationpolicynoncompliancereport"
    ib550761bd4a7b3f5c6fbf730879368bdf6cf1f2e2a17f428ce1aff9b9d56b7bb "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/reports/getconfigurationpolicynoncompliancesummaryreport"
    ib8a4803ec3b35355bac5ee1f09413f3cedea9e96c350c5a3b7a72a1ce2329cd8 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/reports/getpolicynoncompliancemetadata"
    ic72dce9d0b6163f9d68d5eda111ccd3bd3a8ad43383cda3febb59c81ff2bb55f "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/reports/getcachedreport"
    id4ecac28ab783e0d39a17b453b0ef71f4a50bd1409017afcb0ca9e6b503faa03 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/reports/getconfigurationsettingnoncompliancereport"
    idf9e4ac66c87266487fcd200b5d2bd9b6e5236a426d072ab78c2b7d2b87f8fcd "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/reports/getdevicenoncompliancereport"
    ie35af60f2ad5a0746fb8d0fd9890e6a8832b0296a222da44702a59e069097f08 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/reports/getcompliancesettingnoncompliancereport"
    ie9d9c171f1948a762a7fe58817fdfd08ca15474b245a8d61a3965d150dd4a6df "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/reports/getcompliancepolicynoncompliancesummaryreport"
    i485f9fd192f0985070e6776ebbebebbeab22aaa9df64be7131a3e622b77e1c4e "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/reports/exportjobs/item"
)

// Builds and executes requests for operations under \deviceManagement\reports
type ReportsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type ReportsRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type ReportsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ReportsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Reports singleton
type ReportsRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type ReportsRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceManagementReports;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new ReportsRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ReportsRequestBuilder) {
    m := &ReportsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ReportsRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewReportsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// Reports singleton
// Parameters:
//  - options : Options for the request
func (m *ReportsRequestBuilder) CreateDeleteRequestInformation(options *ReportsRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Reports singleton
// Parameters:
//  - options : Options for the request
func (m *ReportsRequestBuilder) CreateGetRequestInformation(options *ReportsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// Reports singleton
// Parameters:
//  - options : Options for the request
func (m *ReportsRequestBuilder) CreatePatchRequestInformation(options *ReportsRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Reports singleton
// Parameters:
//  - options : Options for the request
func (m *ReportsRequestBuilder) Delete(options *ReportsRequestBuilderDeleteOptions)(error) {
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
func (m *ReportsRequestBuilder) ExportJobs()(*i062d4a86380aad1b2a18fcdf2ae3039afe4fb3340e8c61cd8a0118f64e2961f4.ExportJobsRequestBuilder) {
    return i062d4a86380aad1b2a18fcdf2ae3039afe4fb3340e8c61cd8a0118f64e2961f4.NewExportJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.deviceManagement.reports.exportJobs.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ReportsRequestBuilder) ExportJobsById(id string)(*i485f9fd192f0985070e6776ebbebebbeab22aaa9df64be7131a3e622b77e1c4e.DeviceManagementExportJobRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementExportJob_id"] = id
    }
    return i485f9fd192f0985070e6776ebbebebbeab22aaa9df64be7131a3e622b77e1c4e.NewDeviceManagementExportJobRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Reports singleton
// Parameters:
//  - options : Options for the request
func (m *ReportsRequestBuilder) Get(options *ReportsRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceManagementReports, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewDeviceManagementReports() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceManagementReports), nil
}
func (m *ReportsRequestBuilder) GetCachedReport()(*ic72dce9d0b6163f9d68d5eda111ccd3bd3a8ad43383cda3febb59c81ff2bb55f.GetCachedReportRequestBuilder) {
    return ic72dce9d0b6163f9d68d5eda111ccd3bd3a8ad43383cda3febb59c81ff2bb55f.NewGetCachedReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetCompliancePolicyNonComplianceReport()(*i2f4861dbc7fdcc004c73b6f3a193e531630645ae4e3dba4a06598c5ac995763e.GetCompliancePolicyNonComplianceReportRequestBuilder) {
    return i2f4861dbc7fdcc004c73b6f3a193e531630645ae4e3dba4a06598c5ac995763e.NewGetCompliancePolicyNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetCompliancePolicyNonComplianceSummaryReport()(*ie9d9c171f1948a762a7fe58817fdfd08ca15474b245a8d61a3965d150dd4a6df.GetCompliancePolicyNonComplianceSummaryReportRequestBuilder) {
    return ie9d9c171f1948a762a7fe58817fdfd08ca15474b245a8d61a3965d150dd4a6df.NewGetCompliancePolicyNonComplianceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetComplianceSettingNonComplianceReport()(*ie35af60f2ad5a0746fb8d0fd9890e6a8832b0296a222da44702a59e069097f08.GetComplianceSettingNonComplianceReportRequestBuilder) {
    return ie35af60f2ad5a0746fb8d0fd9890e6a8832b0296a222da44702a59e069097f08.NewGetComplianceSettingNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetConfigurationPolicyNonComplianceReport()(*ib215a0b2b7d6a201fa215cdae39964212be5dea6ccfa14879498271f5d9820d1.GetConfigurationPolicyNonComplianceReportRequestBuilder) {
    return ib215a0b2b7d6a201fa215cdae39964212be5dea6ccfa14879498271f5d9820d1.NewGetConfigurationPolicyNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetConfigurationPolicyNonComplianceSummaryReport()(*ib550761bd4a7b3f5c6fbf730879368bdf6cf1f2e2a17f428ce1aff9b9d56b7bb.GetConfigurationPolicyNonComplianceSummaryReportRequestBuilder) {
    return ib550761bd4a7b3f5c6fbf730879368bdf6cf1f2e2a17f428ce1aff9b9d56b7bb.NewGetConfigurationPolicyNonComplianceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetConfigurationSettingNonComplianceReport()(*id4ecac28ab783e0d39a17b453b0ef71f4a50bd1409017afcb0ca9e6b503faa03.GetConfigurationSettingNonComplianceReportRequestBuilder) {
    return id4ecac28ab783e0d39a17b453b0ef71f4a50bd1409017afcb0ca9e6b503faa03.NewGetConfigurationSettingNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetDeviceManagementIntentPerSettingContributingProfiles()(*ia8bd370622efb0565c68883e987555ee843b6553925fd83251129c0ab1a330bf.GetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder) {
    return ia8bd370622efb0565c68883e987555ee843b6553925fd83251129c0ab1a330bf.NewGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetDeviceManagementIntentSettingsReport()(*i0a839f3beac17f684c85ef687fb80c09048b3dcae7c1ecbde1af40e48df57f93.GetDeviceManagementIntentSettingsReportRequestBuilder) {
    return i0a839f3beac17f684c85ef687fb80c09048b3dcae7c1ecbde1af40e48df57f93.NewGetDeviceManagementIntentSettingsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetDeviceNonComplianceReport()(*idf9e4ac66c87266487fcd200b5d2bd9b6e5236a426d072ab78c2b7d2b87f8fcd.GetDeviceNonComplianceReportRequestBuilder) {
    return idf9e4ac66c87266487fcd200b5d2bd9b6e5236a426d072ab78c2b7d2b87f8fcd.NewGetDeviceNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetHistoricalReport()(*i63f7c0aabdb80575f559f70509dbec8f47b990899881ff53424a1b846dd52597.GetHistoricalReportRequestBuilder) {
    return i63f7c0aabdb80575f559f70509dbec8f47b990899881ff53424a1b846dd52597.NewGetHistoricalReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetPolicyNonComplianceMetadata()(*ib8a4803ec3b35355bac5ee1f09413f3cedea9e96c350c5a3b7a72a1ce2329cd8.GetPolicyNonComplianceMetadataRequestBuilder) {
    return ib8a4803ec3b35355bac5ee1f09413f3cedea9e96c350c5a3b7a72a1ce2329cd8.NewGetPolicyNonComplianceMetadataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetPolicyNonComplianceReport()(*i412da535edde14d19b0f7a70e3391905918a35901cecec9071168452048c7acc.GetPolicyNonComplianceReportRequestBuilder) {
    return i412da535edde14d19b0f7a70e3391905918a35901cecec9071168452048c7acc.NewGetPolicyNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetPolicyNonComplianceSummaryReport()(*ia0c903425c40dceef2013d71020f9eecda163cc98d6a70d6643a96e52d28d782.GetPolicyNonComplianceSummaryReportRequestBuilder) {
    return ia0c903425c40dceef2013d71020f9eecda163cc98d6a70d6643a96e52d28d782.NewGetPolicyNonComplianceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetReportFilters()(*i53f2ba02a8f5254d99ef3d211f9756166f93939a36f554a78d1688617cbe7a5f.GetReportFiltersRequestBuilder) {
    return i53f2ba02a8f5254d99ef3d211f9756166f93939a36f554a78d1688617cbe7a5f.NewGetReportFiltersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ReportsRequestBuilder) GetSettingNonComplianceReport()(*i69e11ae317e8e814a0b7be5423ee2af053bfe7bd422319ac2db1f93ee18dc6f5.GetSettingNonComplianceReportRequestBuilder) {
    return i69e11ae317e8e814a0b7be5423ee2af053bfe7bd422319ac2db1f93ee18dc6f5.NewGetSettingNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Reports singleton
// Parameters:
//  - options : Options for the request
func (m *ReportsRequestBuilder) Patch(options *ReportsRequestBuilderPatchOptions)(error) {
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
