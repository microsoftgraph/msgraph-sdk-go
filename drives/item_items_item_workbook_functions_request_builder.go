package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookFunctionsRequestBuilder provides operations to manage the functions property of the microsoft.graph.workbook entity.
type ItemItemsItemWorkbookFunctionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemsItemWorkbookFunctionsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookFunctionsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemsItemWorkbookFunctionsRequestBuilderGetQueryParameters get functions from drives
type ItemItemsItemWorkbookFunctionsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemItemsItemWorkbookFunctionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookFunctionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemsItemWorkbookFunctionsRequestBuilderGetQueryParameters
}
// ItemItemsItemWorkbookFunctionsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookFunctionsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookFunctionsRequestBuilderInternal instantiates a new FunctionsRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookFunctionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookFunctionsRequestBuilder) {
    m := &ItemItemsItemWorkbookFunctionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/functions{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemItemsItemWorkbookFunctionsRequestBuilder instantiates a new FunctionsRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookFunctionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookFunctionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookFunctionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property functions for drives
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookFunctionsRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get functions from drives
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookFunctionsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFunctionsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookFunctionsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFunctionsable), nil
}
// MicrosoftGraphAbs provides operations to call the abs method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAbs()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAbsAbsRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAbsAbsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphAccrInt provides operations to call the accrInt method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAccrInt()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntAccrIntRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntAccrIntRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphAccrIntM provides operations to call the accrIntM method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAccrIntM()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntMAccrIntMRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntMAccrIntMRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphAcos provides operations to call the acos method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAcos()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAcosAcosRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAcosAcosRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphAcosh provides operations to call the acosh method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAcosh()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAcoshAcoshRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAcoshAcoshRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphAcot provides operations to call the acot method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAcot()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAcotAcotRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAcotAcotRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphAcoth provides operations to call the acoth method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAcoth()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAcothAcothRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAcothAcothRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphAmorDegrc provides operations to call the amorDegrc method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAmorDegrc()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAmorDegrcAmorDegrcRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAmorDegrcAmorDegrcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphAmorLinc provides operations to call the amorLinc method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAmorLinc()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAmorLincAmorLincRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAmorLincAmorLincRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphAnd provides operations to call the and method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAnd()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAndAndRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAndAndRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphArabic provides operations to call the arabic method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphArabic()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphArabicArabicRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphArabicArabicRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphAreas provides operations to call the areas method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAreas()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAreasAreasRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAreasAreasRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphAsc provides operations to call the asc method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAsc()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAscAscRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAscAscRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphAsin provides operations to call the asin method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAsin()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAsinAsinRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAsinAsinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphAsinh provides operations to call the asinh method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAsinh()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAsinhAsinhRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAsinhAsinhRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphAtan provides operations to call the atan method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAtan()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAtanAtanRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAtanAtanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphAtan2 provides operations to call the atan2 method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAtan2()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAtan2Atan2RequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAtan2Atan2RequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphAtanh provides operations to call the atanh method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAtanh()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAtanhAtanhRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAtanhAtanhRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphAveDev provides operations to call the aveDev method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAveDev()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAveDevAveDevRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAveDevAveDevRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphAverage provides operations to call the average method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAverage()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAverageAverageRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAverageAverageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphAverageA provides operations to call the averageA method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAverageA()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAverageAAverageARequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAverageAAverageARequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphAverageIf provides operations to call the averageIf method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAverageIf()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAverageIfAverageIfRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAverageIfAverageIfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphAverageIfs provides operations to call the averageIfs method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAverageIfs()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAverageIfsAverageIfsRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAverageIfsAverageIfsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphBahtText provides operations to call the bahtText method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBahtText()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBahtTextBahtTextRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBahtTextBahtTextRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphBase provides operations to call the base method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBase()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBaseBaseRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBaseBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphBesselI provides operations to call the besselI method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBesselI()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBesselIBesselIRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBesselIBesselIRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphBesselJ provides operations to call the besselJ method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBesselJ()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBesselJBesselJRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBesselJBesselJRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphBesselK provides operations to call the besselK method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBesselK()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBesselKBesselKRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBesselKBesselKRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphBesselY provides operations to call the besselY method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBesselY()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBesselYBesselYRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBesselYBesselYRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphBeta_Dist provides operations to call the beta_Dist method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBeta_Dist()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBeta_DistBeta_DistRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBeta_DistBeta_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphBeta_Inv provides operations to call the beta_Inv method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBeta_Inv()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBeta_InvBeta_InvRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBeta_InvBeta_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphBin2Dec provides operations to call the bin2Dec method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBin2Dec()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBin2DecBin2DecRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBin2DecBin2DecRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphBin2Hex provides operations to call the bin2Hex method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBin2Hex()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBin2HexBin2HexRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBin2HexBin2HexRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphBin2Oct provides operations to call the bin2Oct method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBin2Oct()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBin2OctBin2OctRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBin2OctBin2OctRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphBinom_Dist provides operations to call the binom_Dist method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBinom_Dist()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_DistBinom_DistRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_DistBinom_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphBinom_Dist_Range provides operations to call the binom_Dist_Range method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBinom_Dist_Range()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangeRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphBinom_Inv provides operations to call the binom_Inv method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBinom_Inv()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_InvBinom_InvRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_InvBinom_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphBitand provides operations to call the bitand method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBitand()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBitandBitandRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBitandBitandRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphBitlshift provides operations to call the bitlshift method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBitlshift()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBitlshiftBitlshiftRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBitlshiftBitlshiftRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphBitor provides operations to call the bitor method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBitor()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBitorBitorRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBitorBitorRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphBitrshift provides operations to call the bitrshift method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBitrshift()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBitrshiftBitrshiftRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBitrshiftBitrshiftRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphBitxor provides operations to call the bitxor method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBitxor()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBitxorBitxorRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBitxorBitxorRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphCeiling_Math provides operations to call the ceiling_Math method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCeiling_Math()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCeiling_MathCeiling_MathRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCeiling_MathCeiling_MathRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphCeiling_Precise provides operations to call the ceiling_Precise method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCeiling_Precise()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCeiling_PreciseCeiling_PreciseRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCeiling_PreciseCeiling_PreciseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphChar provides operations to call the char method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphChar()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCharCharRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCharCharRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphChiSq_Dist provides operations to call the chiSq_Dist method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphChiSq_Dist()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_DistChiSq_DistRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_DistChiSq_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphChiSq_Dist_RT provides operations to call the chiSq_Dist_RT method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphChiSq_Dist_RT()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Dist_RTChiSq_Dist_RTRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Dist_RTChiSq_Dist_RTRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphChiSq_Inv provides operations to call the chiSq_Inv method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphChiSq_Inv()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_InvChiSq_InvRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_InvChiSq_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphChiSq_Inv_RT provides operations to call the chiSq_Inv_RT method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphChiSq_Inv_RT()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Inv_RTChiSq_Inv_RTRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Inv_RTChiSq_Inv_RTRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphChoose provides operations to call the choose method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphChoose()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphChooseChooseRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphChooseChooseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphClean provides operations to call the clean method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphClean()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCleanCleanRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCleanCleanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphCode provides operations to call the code method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCode()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCodeCodeRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCodeCodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphColumns provides operations to call the columns method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphColumns()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphColumnsColumnsRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphColumnsColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphCombin provides operations to call the combin method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCombin()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCombinCombinRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCombinCombinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphCombina provides operations to call the combina method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCombina()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCombinaCombinaRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCombinaCombinaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphComplex provides operations to call the complex method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphComplex()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphComplexComplexRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphComplexComplexRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphConcatenate provides operations to call the concatenate method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphConcatenate()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphConcatenateConcatenateRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphConcatenateConcatenateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphConfidence_Norm provides operations to call the confidence_Norm method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphConfidence_Norm()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphConfidence_NormConfidence_NormRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphConfidence_NormConfidence_NormRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphConfidence_T provides operations to call the confidence_T method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphConfidence_T()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphConfidence_TConfidence_TRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphConfidence_TConfidence_TRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphConvert provides operations to call the convert method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphConvert()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphConvertConvertRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphConvertConvertRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphCos provides operations to call the cos method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCos()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCosCosRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCosCosRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphCosh provides operations to call the cosh method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCosh()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCoshCoshRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCoshCoshRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphCot provides operations to call the cot method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCot()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCotCotRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCotCotRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphCoth provides operations to call the coth method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCoth()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCothCothRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCothCothRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphCount provides operations to call the count method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCount()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCountCountRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCountCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphCountA provides operations to call the countA method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCountA()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCountACountARequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCountACountARequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphCountBlank provides operations to call the countBlank method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCountBlank()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCountBlankCountBlankRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCountBlankCountBlankRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphCountIf provides operations to call the countIf method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCountIf()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCountIfCountIfRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCountIfCountIfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphCountIfs provides operations to call the countIfs method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCountIfs()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCountIfsCountIfsRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCountIfsCountIfsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphCoupDayBs provides operations to call the coupDayBs method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCoupDayBs()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCoupDayBsCoupDayBsRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCoupDayBsCoupDayBsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphCoupDays provides operations to call the coupDays method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCoupDays()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCoupDaysCoupDaysRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCoupDaysCoupDaysRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphCoupDaysNc provides operations to call the coupDaysNc method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCoupDaysNc()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCoupDaysNcCoupDaysNcRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCoupDaysNcCoupDaysNcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphCoupNcd provides operations to call the coupNcd method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCoupNcd()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCoupNcdCoupNcdRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCoupNcdCoupNcdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphCoupNum provides operations to call the coupNum method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCoupNum()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCoupNumCoupNumRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCoupNumCoupNumRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphCoupPcd provides operations to call the coupPcd method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCoupPcd()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCoupPcdCoupPcdRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCoupPcdCoupPcdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphCsc provides operations to call the csc method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCsc()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCscCscRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCscCscRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphCsch provides operations to call the csch method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCsch()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCschCschRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCschCschRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphCumIPmt provides operations to call the cumIPmt method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCumIPmt()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCumIPmtCumIPmtRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCumIPmtCumIPmtRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphCumPrinc provides operations to call the cumPrinc method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCumPrinc()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCumPrincCumPrincRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCumPrincCumPrincRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphDate provides operations to call the date method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDate()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDateDateRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDateDateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphDatevalue provides operations to call the datevalue method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDatevalue()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDatevalueDatevalueRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDatevalueDatevalueRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphDaverage provides operations to call the daverage method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDaverage()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDaverageDaverageRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDaverageDaverageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphDay provides operations to call the day method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDay()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDayDayRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDayDayRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphDays provides operations to call the days method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDays()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDaysDaysRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDaysDaysRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphDays360 provides operations to call the days360 method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDays360()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDays360Days360RequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDays360Days360RequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphDb provides operations to call the db method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDb()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDbDbRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDbDbRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphDbcs provides operations to call the dbcs method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDbcs()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDbcsDbcsRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDbcsDbcsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphDcount provides operations to call the dcount method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDcount()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDcountDcountRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDcountDcountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphDcountA provides operations to call the dcountA method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDcountA()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDcountADcountARequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDcountADcountARequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphDdb provides operations to call the ddb method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDdb()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDdbDdbRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDdbDdbRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphDec2Bin provides operations to call the dec2Bin method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDec2Bin()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDec2BinDec2BinRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDec2BinDec2BinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphDec2Hex provides operations to call the dec2Hex method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDec2Hex()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDec2HexDec2HexRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDec2HexDec2HexRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphDec2Oct provides operations to call the dec2Oct method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDec2Oct()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDec2OctDec2OctRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDec2OctDec2OctRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphDecimal provides operations to call the decimal method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDecimal()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDecimalDecimalRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDecimalDecimalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphDegrees provides operations to call the degrees method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDegrees()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDegreesDegreesRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDegreesDegreesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphDelta provides operations to call the delta method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDelta()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDeltaDeltaRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDeltaDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphDevSq provides operations to call the devSq method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDevSq()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDevSqDevSqRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDevSqDevSqRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphDget provides operations to call the dget method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDget()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDgetDgetRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDgetDgetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphDisc provides operations to call the disc method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDisc()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDiscDiscRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDiscDiscRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphDmax provides operations to call the dmax method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDmax()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDmaxDmaxRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDmaxDmaxRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphDmin provides operations to call the dmin method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDmin()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDminDminRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDminDminRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphDollar provides operations to call the dollar method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDollar()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDollarDollarRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDollarDollarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphDollarDe provides operations to call the dollarDe method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDollarDe()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDollarDeDollarDeRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDollarDeDollarDeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphDollarFr provides operations to call the dollarFr method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDollarFr()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDollarFrDollarFrRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDollarFrDollarFrRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphDproduct provides operations to call the dproduct method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDproduct()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDproductDproductRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDproductDproductRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphDstDev provides operations to call the dstDev method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDstDev()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDstDevDstDevRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDstDevDstDevRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphDstDevP provides operations to call the dstDevP method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDstDevP()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDstDevPDstDevPRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDstDevPDstDevPRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphDsum provides operations to call the dsum method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDsum()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDsumDsumRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDsumDsumRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphDuration provides operations to call the duration method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDuration()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDurationDurationRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDurationDurationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphDvar provides operations to call the dvar method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDvar()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDvarDvarRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDvarDvarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphDvarP provides operations to call the dvarP method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDvarP()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDvarPDvarPRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDvarPDvarPRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphEcma_Ceiling provides operations to call the ecma_Ceiling method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphEcma_Ceiling()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphEcma_CeilingEcma_CeilingRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphEcma_CeilingEcma_CeilingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphEdate provides operations to call the edate method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphEdate()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphEdateEdateRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphEdateEdateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphEffect provides operations to call the effect method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphEffect()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphEffectEffectRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphEffectEffectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphEoMonth provides operations to call the eoMonth method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphEoMonth()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphEoMonthEoMonthRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphEoMonthEoMonthRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphErf provides operations to call the erf method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphErf()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphErfErfRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphErfErfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphErf_Precise provides operations to call the erf_Precise method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphErf_Precise()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphErf_PreciseErf_PreciseRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphErf_PreciseErf_PreciseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphErfC provides operations to call the erfC method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphErfC()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphErfCErfCRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphErfCErfCRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphErfC_Precise provides operations to call the erfC_Precise method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphErfC_Precise()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphErfC_PreciseErfC_PreciseRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphErfC_PreciseErfC_PreciseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphError_Type provides operations to call the error_Type method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphError_Type()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphError_TypeError_TypeRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphError_TypeError_TypeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphEven provides operations to call the even method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphEven()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphEvenEvenRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphEvenEvenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphExact provides operations to call the exact method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphExact()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphExactExactRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphExactExactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphExp provides operations to call the exp method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphExp()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphExpExpRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphExpExpRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphExpon_Dist provides operations to call the expon_Dist method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphExpon_Dist()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphExpon_DistExpon_DistRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphExpon_DistExpon_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphF_Dist provides operations to call the f_Dist method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphF_Dist()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphF_DistF_DistRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphF_DistF_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphF_Dist_RT provides operations to call the f_Dist_RT method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphF_Dist_RT()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphF_Dist_RTF_Dist_RTRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphF_Dist_RTF_Dist_RTRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphF_Inv provides operations to call the f_Inv method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphF_Inv()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphF_InvF_InvRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphF_InvF_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphF_Inv_RT provides operations to call the f_Inv_RT method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphF_Inv_RT()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphF_Inv_RTF_Inv_RTRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphF_Inv_RTF_Inv_RTRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphFact provides operations to call the fact method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphFact()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphFactFactRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphFactFactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphFactDouble provides operations to call the factDouble method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphFactDouble()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphFactDoubleFactDoubleRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphFactDoubleFactDoubleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphFalse provides operations to call the false method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphFalse()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphFalseFalseRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphFalseFalseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphFind provides operations to call the find method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphFind()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphFindFindRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphFindFindRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphFindB provides operations to call the findB method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphFindB()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphFindBFindBRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphFindBFindBRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphFisher provides operations to call the fisher method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphFisher()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphFisherFisherRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphFisherFisherRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphFisherInv provides operations to call the fisherInv method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphFisherInv()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphFisherInvFisherInvRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphFisherInvFisherInvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphFixed provides operations to call the fixed method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphFixed()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphFixedFixedRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphFixedFixedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphFloor_Math provides operations to call the floor_Math method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphFloor_Math()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphFloor_MathFloor_MathRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphFloor_MathFloor_MathRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphFloor_Precise provides operations to call the floor_Precise method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphFloor_Precise()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphFloor_PreciseFloor_PreciseRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphFloor_PreciseFloor_PreciseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphFv provides operations to call the fv method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphFv()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphFvFvRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphFvFvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphFvschedule provides operations to call the fvschedule method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphFvschedule()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphFvscheduleFvscheduleRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphFvscheduleFvscheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphGamma provides operations to call the gamma method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphGamma()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphGammaGammaRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphGammaGammaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphGamma_Dist provides operations to call the gamma_Dist method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphGamma_Dist()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphGamma_DistGamma_DistRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphGamma_DistGamma_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphGamma_Inv provides operations to call the gamma_Inv method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphGamma_Inv()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphGamma_InvGamma_InvRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphGamma_InvGamma_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphGammaLn provides operations to call the gammaLn method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphGammaLn()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphGammaLnGammaLnRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphGammaLnGammaLnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphGammaLn_Precise provides operations to call the gammaLn_Precise method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphGammaLn_Precise()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphGammaLn_PreciseGammaLn_PreciseRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphGammaLn_PreciseGammaLn_PreciseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphGauss provides operations to call the gauss method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphGauss()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphGaussGaussRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphGaussGaussRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphGcd provides operations to call the gcd method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphGcd()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphGcdGcdRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphGcdGcdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphGeoMean provides operations to call the geoMean method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphGeoMean()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphGeoMeanGeoMeanRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphGeoMeanGeoMeanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphGeStep provides operations to call the geStep method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphGeStep()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphGeStepGeStepRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphGeStepGeStepRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphHarMean provides operations to call the harMean method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphHarMean()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphHarMeanHarMeanRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphHarMeanHarMeanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphHex2Bin provides operations to call the hex2Bin method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphHex2Bin()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphHex2BinHex2BinRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphHex2BinHex2BinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphHex2Dec provides operations to call the hex2Dec method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphHex2Dec()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphHex2DecHex2DecRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphHex2DecHex2DecRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphHex2Oct provides operations to call the hex2Oct method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphHex2Oct()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphHex2OctHex2OctRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphHex2OctHex2OctRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphHlookup provides operations to call the hlookup method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphHlookup()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphHlookupHlookupRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphHlookupHlookupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphHour provides operations to call the hour method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphHour()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphHourHourRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphHourHourRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphHyperlink provides operations to call the hyperlink method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphHyperlink()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphHyperlinkHyperlinkRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphHyperlinkHyperlinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphHypGeom_Dist provides operations to call the hypGeom_Dist method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphHypGeom_Dist()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistHypGeom_DistRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistHypGeom_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphIf provides operations to call the if method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIf()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIfIfRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIfIfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphImAbs provides operations to call the imAbs method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImAbs()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImAbsImAbsRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImAbsImAbsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphImaginary provides operations to call the imaginary method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImaginary()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImaginaryImaginaryRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImaginaryImaginaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphImArgument provides operations to call the imArgument method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImArgument()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImArgumentImArgumentRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImArgumentImArgumentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphImConjugate provides operations to call the imConjugate method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImConjugate()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImConjugateImConjugateRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImConjugateImConjugateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphImCos provides operations to call the imCos method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImCos()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImCosImCosRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImCosImCosRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphImCosh provides operations to call the imCosh method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImCosh()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImCoshImCoshRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImCoshImCoshRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphImCot provides operations to call the imCot method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImCot()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImCotImCotRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImCotImCotRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphImCsc provides operations to call the imCsc method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImCsc()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImCscImCscRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImCscImCscRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphImCsch provides operations to call the imCsch method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImCsch()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImCschImCschRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImCschImCschRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphImDiv provides operations to call the imDiv method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImDiv()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImDivImDivRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImDivImDivRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphImExp provides operations to call the imExp method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImExp()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImExpImExpRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImExpImExpRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphImLn provides operations to call the imLn method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImLn()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImLnImLnRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImLnImLnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphImLog10 provides operations to call the imLog10 method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImLog10()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImLog10ImLog10RequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImLog10ImLog10RequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphImLog2 provides operations to call the imLog2 method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImLog2()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImLog2ImLog2RequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImLog2ImLog2RequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphImPower provides operations to call the imPower method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImPower()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImPowerImPowerRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImPowerImPowerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphImProduct provides operations to call the imProduct method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImProduct()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImProductImProductRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImProductImProductRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphImReal provides operations to call the imReal method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImReal()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImRealImRealRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImRealImRealRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphImSec provides operations to call the imSec method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImSec()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImSecImSecRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImSecImSecRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphImSech provides operations to call the imSech method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImSech()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImSechImSechRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImSechImSechRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphImSin provides operations to call the imSin method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImSin()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImSinImSinRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImSinImSinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphImSinh provides operations to call the imSinh method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImSinh()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImSinhImSinhRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImSinhImSinhRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphImSqrt provides operations to call the imSqrt method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImSqrt()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImSqrtImSqrtRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImSqrtImSqrtRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphImSub provides operations to call the imSub method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImSub()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImSubImSubRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImSubImSubRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphImSum provides operations to call the imSum method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImSum()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImSumImSumRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImSumImSumRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphImTan provides operations to call the imTan method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImTan()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImTanImTanRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImTanImTanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphInt provides operations to call the int method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphInt()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIntIntRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIntIntRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphIntRate provides operations to call the intRate method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIntRate()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIntRateIntRateRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIntRateIntRateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphIpmt provides operations to call the ipmt method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIpmt()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIpmtIpmtRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIpmtIpmtRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphIrr provides operations to call the irr method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIrr()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIrrIrrRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIrrIrrRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphIsErr provides operations to call the isErr method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIsErr()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIsErrIsErrRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIsErrIsErrRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphIsError provides operations to call the isError method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIsError()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIsErrorIsErrorRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIsErrorIsErrorRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphIsEven provides operations to call the isEven method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIsEven()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIsEvenIsEvenRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIsEvenIsEvenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphIsFormula provides operations to call the isFormula method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIsFormula()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIsFormulaIsFormulaRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIsFormulaIsFormulaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphIsLogical provides operations to call the isLogical method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIsLogical()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIsLogicalIsLogicalRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIsLogicalIsLogicalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphIsNA provides operations to call the isNA method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIsNA()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIsNAIsNARequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIsNAIsNARequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphIsNonText provides operations to call the isNonText method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIsNonText()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIsNonTextIsNonTextRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIsNonTextIsNonTextRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphIsNumber provides operations to call the isNumber method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIsNumber()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIsNumberIsNumberRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIsNumberIsNumberRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphIso_Ceiling provides operations to call the iso_Ceiling method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIso_Ceiling()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIso_CeilingIso_CeilingRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIso_CeilingIso_CeilingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphIsOdd provides operations to call the isOdd method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIsOdd()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIsOddIsOddRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIsOddIsOddRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphIsoWeekNum provides operations to call the isoWeekNum method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIsoWeekNum()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIsoWeekNumIsoWeekNumRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIsoWeekNumIsoWeekNumRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphIspmt provides operations to call the ispmt method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIspmt()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIspmtIspmtRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIspmtIspmtRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphIsref provides operations to call the isref method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIsref()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIsrefIsrefRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIsrefIsrefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphIsText provides operations to call the isText method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIsText()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIsTextIsTextRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIsTextIsTextRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphKurt provides operations to call the kurt method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphKurt()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphKurtKurtRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphKurtKurtRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphLarge provides operations to call the large method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphLarge()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphLargeLargeRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphLargeLargeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphLcm provides operations to call the lcm method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphLcm()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphLcmLcmRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphLcmLcmRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphLeft provides operations to call the left method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphLeft()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphLeftLeftRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphLeftLeftRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphLeftb provides operations to call the leftb method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphLeftb()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphLeftbLeftbRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphLeftbLeftbRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphLen provides operations to call the len method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphLen()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphLenLenRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphLenLenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphLenb provides operations to call the lenb method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphLenb()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphLenbLenbRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphLenbLenbRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphLn provides operations to call the ln method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphLn()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphLnLnRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphLnLnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphLog provides operations to call the log method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphLog()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphLogLogRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphLogLogRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphLog10 provides operations to call the log10 method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphLog10()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphLog10Log10RequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphLog10Log10RequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphLogNorm_Dist provides operations to call the logNorm_Dist method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphLogNorm_Dist()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphLogNorm_DistLogNorm_DistRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphLogNorm_DistLogNorm_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphLogNorm_Inv provides operations to call the logNorm_Inv method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphLogNorm_Inv()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphLogNorm_InvLogNorm_InvRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphLogNorm_InvLogNorm_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphLookup provides operations to call the lookup method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphLookup()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphLookupLookupRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphLookupLookupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphLower provides operations to call the lower method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphLower()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphLowerLowerRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphLowerLowerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphMatch provides operations to call the match method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphMatch()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphMatchMatchRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphMatchMatchRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphMax provides operations to call the max method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphMax()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphMaxMaxRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphMaxMaxRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphMaxA provides operations to call the maxA method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphMaxA()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphMaxAMaxARequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphMaxAMaxARequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphMduration provides operations to call the mduration method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphMduration()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphMdurationMdurationRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphMdurationMdurationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphMedian provides operations to call the median method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphMedian()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphMedianMedianRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphMedianMedianRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphMid provides operations to call the mid method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphMid()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphMidMidRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphMidMidRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphMidb provides operations to call the midb method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphMidb()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphMidbMidbRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphMidbMidbRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphMin provides operations to call the min method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphMin()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphMinMinRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphMinMinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphMinA provides operations to call the minA method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphMinA()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphMinAMinARequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphMinAMinARequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphMinute provides operations to call the minute method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphMinute()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphMinuteMinuteRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphMinuteMinuteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphMirr provides operations to call the mirr method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphMirr()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphMirrMirrRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphMirrMirrRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphMod provides operations to call the mod method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphMod()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphModModRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphModModRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphMonth provides operations to call the month method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphMonth()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphMonthMonthRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphMonthMonthRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphMround provides operations to call the mround method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphMround()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphMroundMroundRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphMroundMroundRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphMultiNomial provides operations to call the multiNomial method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphMultiNomial()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphMultiNomialMultiNomialRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphMultiNomialMultiNomialRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphN provides operations to call the n method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphN()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphNNRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphNNRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphNa provides operations to call the na method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphNa()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphNaNaRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphNaNaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphNegBinom_Dist provides operations to call the negBinom_Dist method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphNegBinom_Dist()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphNegBinom_DistNegBinom_DistRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphNegBinom_DistNegBinom_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphNetworkDays provides operations to call the networkDays method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphNetworkDays()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphNetworkDaysNetworkDaysRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphNetworkDaysNetworkDaysRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphNetworkDays_Intl provides operations to call the networkDays_Intl method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphNetworkDays_Intl()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphNetworkDays_IntlNetworkDays_IntlRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphNetworkDays_IntlNetworkDays_IntlRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphNominal provides operations to call the nominal method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphNominal()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphNominalNominalRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphNominalNominalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphNorm_Dist provides operations to call the norm_Dist method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphNorm_Dist()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_DistNorm_DistRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_DistNorm_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphNorm_Inv provides operations to call the norm_Inv method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphNorm_Inv()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_InvNorm_InvRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_InvNorm_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphNorm_S_Dist provides operations to call the norm_S_Dist method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphNorm_S_Dist()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_S_DistNorm_S_DistRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_S_DistNorm_S_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphNorm_S_Inv provides operations to call the norm_S_Inv method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphNorm_S_Inv()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_S_InvNorm_S_InvRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_S_InvNorm_S_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphNot provides operations to call the not method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphNot()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphNotNotRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphNotNotRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphNow provides operations to call the now method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphNow()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphNowNowRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphNowNowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphNper provides operations to call the nper method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphNper()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphNperNperRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphNperNperRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphNpv provides operations to call the npv method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphNpv()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphNpvNpvRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphNpvNpvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphNumberValue provides operations to call the numberValue method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphNumberValue()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphNumberValueNumberValueRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphNumberValueNumberValueRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphOct2Bin provides operations to call the oct2Bin method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphOct2Bin()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphOct2BinOct2BinRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphOct2BinOct2BinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphOct2Dec provides operations to call the oct2Dec method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphOct2Dec()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphOct2DecOct2DecRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphOct2DecOct2DecRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphOct2Hex provides operations to call the oct2Hex method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphOct2Hex()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphOct2HexOct2HexRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphOct2HexOct2HexRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphOdd provides operations to call the odd method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphOdd()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphOddOddRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphOddOddRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphOddFPrice provides operations to call the oddFPrice method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphOddFPrice()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphOddFPriceOddFPriceRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphOddFPriceOddFPriceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphOddFYield provides operations to call the oddFYield method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphOddFYield()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphOddFYieldOddFYieldRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphOddFYieldOddFYieldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphOddLPrice provides operations to call the oddLPrice method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphOddLPrice()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphOddLPriceOddLPriceRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphOddLPriceOddLPriceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphOddLYield provides operations to call the oddLYield method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphOddLYield()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphOddLYieldOddLYieldRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphOddLYieldOddLYieldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphOr provides operations to call the or method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphOr()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphOrOrRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphOrOrRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphPduration provides operations to call the pduration method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphPduration()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphPdurationPdurationRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphPdurationPdurationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphPercentile_Exc provides operations to call the percentile_Exc method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphPercentile_Exc()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphPercentile_ExcPercentile_ExcRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphPercentile_ExcPercentile_ExcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphPercentile_Inc provides operations to call the percentile_Inc method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphPercentile_Inc()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphPercentile_IncPercentile_IncRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphPercentile_IncPercentile_IncRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphPercentRank_Exc provides operations to call the percentRank_Exc method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphPercentRank_Exc()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphPercentRank_ExcPercentRank_ExcRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphPercentRank_ExcPercentRank_ExcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphPercentRank_Inc provides operations to call the percentRank_Inc method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphPercentRank_Inc()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphPercentRank_IncPercentRank_IncRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphPercentRank_IncPercentRank_IncRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphPermut provides operations to call the permut method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphPermut()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphPermutPermutRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphPermutPermutRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphPermutationa provides operations to call the permutationa method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphPermutationa()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphPermutationaPermutationaRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphPermutationaPermutationaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphPhi provides operations to call the phi method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphPhi()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphPhiPhiRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphPhiPhiRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphPi provides operations to call the pi method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphPi()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphPiPiRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphPiPiRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphPmt provides operations to call the pmt method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphPmt()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphPmtPmtRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphPmtPmtRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphPoisson_Dist provides operations to call the poisson_Dist method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphPoisson_Dist()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphPoisson_DistPoisson_DistRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphPoisson_DistPoisson_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphPower provides operations to call the power method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphPower()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphPowerPowerRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphPowerPowerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphPpmt provides operations to call the ppmt method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphPpmt()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphPpmtPpmtRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphPpmtPpmtRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphPrice provides operations to call the price method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphPrice()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphPricePriceRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphPricePriceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphPriceDisc provides operations to call the priceDisc method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphPriceDisc()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphPriceDiscPriceDiscRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphPriceDiscPriceDiscRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphPriceMat provides operations to call the priceMat method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphPriceMat()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphPriceMatPriceMatRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphPriceMatPriceMatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphProduct provides operations to call the product method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphProduct()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphProductProductRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphProductProductRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphProper provides operations to call the proper method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphProper()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphProperProperRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphProperProperRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphPv provides operations to call the pv method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphPv()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphPvPvRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphPvPvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphQuartile_Exc provides operations to call the quartile_Exc method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphQuartile_Exc()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphQuartile_ExcQuartile_ExcRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphQuartile_ExcQuartile_ExcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphQuartile_Inc provides operations to call the quartile_Inc method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphQuartile_Inc()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphQuartile_IncQuartile_IncRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphQuartile_IncQuartile_IncRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphQuotient provides operations to call the quotient method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphQuotient()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphQuotientQuotientRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphQuotientQuotientRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphRadians provides operations to call the radians method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphRadians()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphRadiansRadiansRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphRadiansRadiansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphRand provides operations to call the rand method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphRand()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphRandRandRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphRandRandRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphRandBetween provides operations to call the randBetween method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphRandBetween()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphRandBetweenRandBetweenRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphRandBetweenRandBetweenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphRank_Avg provides operations to call the rank_Avg method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphRank_Avg()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphRank_AvgRank_AvgRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphRank_AvgRank_AvgRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphRank_Eq provides operations to call the rank_Eq method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphRank_Eq()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphRank_EqRank_EqRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphRank_EqRank_EqRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphRate provides operations to call the rate method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphRate()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphRateRateRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphRateRateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphReceived provides operations to call the received method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphReceived()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphReceivedReceivedRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphReceivedReceivedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphReplace provides operations to call the replace method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphReplace()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceReplaceRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceReplaceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphReplaceB provides operations to call the replaceB method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphReplaceB()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceBReplaceBRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceBReplaceBRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphRept provides operations to call the rept method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphRept()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphReptReptRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphReptReptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphRight provides operations to call the right method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphRight()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphRightRightRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphRightRightRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphRightb provides operations to call the rightb method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphRightb()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphRightbRightbRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphRightbRightbRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphRoman provides operations to call the roman method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphRoman()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphRomanRomanRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphRomanRomanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphRound provides operations to call the round method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphRound()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphRoundRoundRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphRoundRoundRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphRoundDown provides operations to call the roundDown method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphRoundDown()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphRoundDownRoundDownRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphRoundDownRoundDownRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphRoundUp provides operations to call the roundUp method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphRoundUp()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphRoundUpRoundUpRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphRoundUpRoundUpRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphRows provides operations to call the rows method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphRows()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphRowsRowsRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphRowsRowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphRri provides operations to call the rri method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphRri()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphRriRriRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphRriRriRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphSec provides operations to call the sec method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSec()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSecSecRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSecSecRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphSech provides operations to call the sech method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSech()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSechSechRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSechSechRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphSecond provides operations to call the second method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSecond()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSecondSecondRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSecondSecondRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphSeriesSum provides operations to call the seriesSum method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSeriesSum()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSeriesSumSeriesSumRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSeriesSumSeriesSumRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphSheet provides operations to call the sheet method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSheet()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSheetSheetRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSheetSheetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphSheets provides operations to call the sheets method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSheets()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSheetsSheetsRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSheetsSheetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphSign provides operations to call the sign method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSign()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSignSignRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSignSignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphSin provides operations to call the sin method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSin()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSinSinRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSinSinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphSinh provides operations to call the sinh method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSinh()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSinhSinhRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSinhSinhRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphSkew provides operations to call the skew method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSkew()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSkewSkewRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSkewSkewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphSkew_p provides operations to call the skew_p method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSkew_p()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSkew_pSkew_pRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSkew_pSkew_pRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphSln provides operations to call the sln method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSln()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSlnSlnRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSlnSlnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphSmall provides operations to call the small method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSmall()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSmallSmallRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSmallSmallRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphSqrt provides operations to call the sqrt method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSqrt()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSqrtSqrtRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSqrtSqrtRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphSqrtPi provides operations to call the sqrtPi method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSqrtPi()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSqrtPiSqrtPiRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSqrtPiSqrtPiRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphStandardize provides operations to call the standardize method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphStandardize()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphStandardizeStandardizeRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphStandardizeStandardizeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphStDev_P provides operations to call the stDev_P method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphStDev_P()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphStDev_PStDev_PRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphStDev_PStDev_PRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphStDev_S provides operations to call the stDev_S method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphStDev_S()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphStDev_SStDev_SRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphStDev_SStDev_SRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphStDevA provides operations to call the stDevA method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphStDevA()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphStDevAStDevARequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphStDevAStDevARequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphStDevPA provides operations to call the stDevPA method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphStDevPA()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphStDevPAStDevPARequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphStDevPAStDevPARequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphSubstitute provides operations to call the substitute method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSubstitute()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSubstituteSubstituteRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSubstituteSubstituteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphSubtotal provides operations to call the subtotal method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSubtotal()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSubtotalSubtotalRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSubtotalSubtotalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphSum provides operations to call the sum method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSum()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSumSumRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSumSumRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphSumIf provides operations to call the sumIf method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSumIf()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSumIfSumIfRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSumIfSumIfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphSumIfs provides operations to call the sumIfs method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSumIfs()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSumIfsSumIfsRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSumIfsSumIfsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphSumSq provides operations to call the sumSq method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSumSq()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSumSqSumSqRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSumSqSumSqRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphSyd provides operations to call the syd method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSyd()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSydSydRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSydSydRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphT provides operations to call the t method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphT()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphTTRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphTTRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphT_Dist provides operations to call the t_Dist method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphT_Dist()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphT_DistT_DistRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphT_DistT_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphT_Dist_2T provides operations to call the t_Dist_2T method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphT_Dist_2T()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphT_Dist_2TT_Dist_2TRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphT_Dist_2TT_Dist_2TRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphT_Dist_RT provides operations to call the t_Dist_RT method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphT_Dist_RT()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphT_Dist_RTT_Dist_RTRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphT_Dist_RTT_Dist_RTRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphT_Inv provides operations to call the t_Inv method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphT_Inv()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphT_InvT_InvRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphT_InvT_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphT_Inv_2T provides operations to call the t_Inv_2T method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphT_Inv_2T()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphT_Inv_2TT_Inv_2TRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphT_Inv_2TT_Inv_2TRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphTan provides operations to call the tan method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphTan()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphTanTanRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphTanTanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphTanh provides operations to call the tanh method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphTanh()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphTanhTanhRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphTanhTanhRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphTbillEq provides operations to call the tbillEq method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphTbillEq()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphTbillEqTbillEqRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphTbillEqTbillEqRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphTbillPrice provides operations to call the tbillPrice method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphTbillPrice()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphTbillPriceTbillPriceRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphTbillPriceTbillPriceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphTbillYield provides operations to call the tbillYield method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphTbillYield()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphTbillYieldTbillYieldRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphTbillYieldTbillYieldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphText provides operations to call the text method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphText()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphTextTextRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphTextTextRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphTime provides operations to call the time method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphTime()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphTimeTimeRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphTimeTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphTimevalue provides operations to call the timevalue method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphTimevalue()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphTimevalueTimevalueRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphTimevalueTimevalueRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphToday provides operations to call the today method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphToday()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphTodayTodayRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphTodayTodayRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphTrim provides operations to call the trim method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphTrim()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphTrimTrimRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphTrimTrimRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphTrimMean provides operations to call the trimMean method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphTrimMean()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphTrimMeanTrimMeanRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphTrimMeanTrimMeanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphTrue provides operations to call the true method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphTrue()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphTrueTrueRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphTrueTrueRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphTrunc provides operations to call the trunc method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphTrunc()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphTruncTruncRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphTruncTruncRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphType provides operations to call the type method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphType()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphTypeTypeRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphTypeTypeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphUnichar provides operations to call the unichar method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphUnichar()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphUnicharUnicharRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphUnicharUnicharRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphUnicode provides operations to call the unicode method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphUnicode()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphUnicodeUnicodeRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphUnicodeUnicodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphUpper provides operations to call the upper method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphUpper()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphUpperUpperRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphUpperUpperRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphUsdollar provides operations to call the usdollar method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphUsdollar()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphUsdollarUsdollarRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphUsdollarUsdollarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphValue provides operations to call the value method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphValue()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphValueValueRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphValueValueRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphVar_P provides operations to call the var_P method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphVar_P()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphVar_PVar_PRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphVar_PVar_PRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphVar_S provides operations to call the var_S method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphVar_S()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphVar_SVar_SRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphVar_SVar_SRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphVarA provides operations to call the varA method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphVarA()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphVarAVarARequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphVarAVarARequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphVarPA provides operations to call the varPA method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphVarPA()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphVarPAVarPARequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphVarPAVarPARequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphVdb provides operations to call the vdb method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphVdb()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphVdbVdbRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphVdbVdbRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphVlookup provides operations to call the vlookup method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphVlookup()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphVlookupVlookupRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphVlookupVlookupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphWeekday provides operations to call the weekday method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphWeekday()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphWeekdayWeekdayRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphWeekdayWeekdayRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphWeekNum provides operations to call the weekNum method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphWeekNum()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphWeekNumWeekNumRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphWeekNumWeekNumRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphWeibull_Dist provides operations to call the weibull_Dist method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphWeibull_Dist()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphWeibull_DistWeibull_DistRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphWeibull_DistWeibull_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphWorkDay provides operations to call the workDay method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphWorkDay()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDayWorkDayRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDayWorkDayRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphWorkDay_Intl provides operations to call the workDay_Intl method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphWorkDay_Intl()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphXirr provides operations to call the xirr method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphXirr()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphXirrXirrRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphXirrXirrRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphXnpv provides operations to call the xnpv method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphXnpv()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphXnpvXnpvRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphXnpvXnpvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphXor provides operations to call the xor method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphXor()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphXorXorRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphXorXorRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphYear provides operations to call the year method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphYear()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphYearYearRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphYearYearRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphYearFrac provides operations to call the yearFrac method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphYearFrac()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphYearFracYearFracRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphYearFracYearFracRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphYield provides operations to call the yield method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphYield()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphYieldYieldRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphYieldYieldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphYieldDisc provides operations to call the yieldDisc method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphYieldDisc()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphYieldDiscYieldDiscRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphYieldDiscYieldDiscRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphYieldMat provides operations to call the yieldMat method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphYieldMat()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphYieldMatYieldMatRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphYieldMatYieldMatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftGraphZ_Test provides operations to call the z_Test method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphZ_Test()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphZ_TestZ_TestRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphZ_TestZ_TestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property functions in drives
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFunctionsable, requestConfiguration *ItemItemsItemWorkbookFunctionsRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFunctionsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookFunctionsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFunctionsable), nil
}
// ToDeleteRequestInformation delete navigation property functions for drives
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookFunctionsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation get functions from drives
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookFunctionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property functions in drives
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFunctionsable, requestConfiguration *ItemItemsItemWorkbookFunctionsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
