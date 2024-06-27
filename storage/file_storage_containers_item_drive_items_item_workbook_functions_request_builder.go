package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder provides operations to manage the functions property of the microsoft.graph.workbook entity.
type FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilderGetQueryParameters get functions from storage
type FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilderGetQueryParameters
}
// FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Abs provides operations to call the abs method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsAbsRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Abs()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsAbsRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsAbsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccrInt provides operations to call the accrInt method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsAccrIntRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) AccrInt()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsAccrIntRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsAccrIntRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccrIntM provides operations to call the accrIntM method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsAccrIntMRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) AccrIntM()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsAccrIntMRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsAccrIntMRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Acos provides operations to call the acos method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsAcosRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Acos()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsAcosRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsAcosRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Acosh provides operations to call the acosh method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsAcoshRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Acosh()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsAcoshRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsAcoshRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Acot provides operations to call the acot method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsAcotRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Acot()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsAcotRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsAcotRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Acoth provides operations to call the acoth method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsAcothRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Acoth()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsAcothRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsAcothRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AmorDegrc provides operations to call the amorDegrc method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsAmorDegrcRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) AmorDegrc()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsAmorDegrcRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsAmorDegrcRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AmorLinc provides operations to call the amorLinc method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsAmorLincRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) AmorLinc()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsAmorLincRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsAmorLincRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// And provides operations to call the and method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsAndRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) And()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsAndRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsAndRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Arabic provides operations to call the arabic method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsArabicRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Arabic()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsArabicRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsArabicRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Areas provides operations to call the areas method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsAreasRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Areas()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsAreasRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsAreasRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Asc provides operations to call the asc method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsAscRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Asc()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsAscRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsAscRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Asin provides operations to call the asin method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsAsinRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Asin()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsAsinRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsAsinRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Asinh provides operations to call the asinh method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsAsinhRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Asinh()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsAsinhRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsAsinhRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Atan provides operations to call the atan method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsAtanRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Atan()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsAtanRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsAtanRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Atan2 provides operations to call the atan2 method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsAtan2RequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Atan2()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsAtan2RequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsAtan2RequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Atanh provides operations to call the atanh method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsAtanhRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Atanh()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsAtanhRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsAtanhRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AveDev provides operations to call the aveDev method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsAveDevRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) AveDev()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsAveDevRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsAveDevRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Average provides operations to call the average method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsAverageRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Average()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsAverageRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsAverageRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AverageA provides operations to call the averageA method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsAverageARequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) AverageA()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsAverageARequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsAverageARequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AverageIf provides operations to call the averageIf method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsAverageIfRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) AverageIf()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsAverageIfRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsAverageIfRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AverageIfs provides operations to call the averageIfs method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsAverageIfsRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) AverageIfs()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsAverageIfsRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsAverageIfsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BahtText provides operations to call the bahtText method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsBahtTextRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) BahtText()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsBahtTextRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsBahtTextRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Base provides operations to call the base method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsBaseRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Base()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsBaseRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsBaseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BesselI provides operations to call the besselI method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsBesselIRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) BesselI()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsBesselIRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsBesselIRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BesselJ provides operations to call the besselJ method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsBesselJRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) BesselJ()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsBesselJRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsBesselJRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BesselK provides operations to call the besselK method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsBesselKRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) BesselK()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsBesselKRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsBesselKRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BesselY provides operations to call the besselY method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsBesselYRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) BesselY()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsBesselYRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsBesselYRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Beta_Dist provides operations to call the beta_Dist method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsBeta_DistRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Beta_Dist()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsBeta_DistRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsBeta_DistRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Beta_Inv provides operations to call the beta_Inv method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsBeta_InvRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Beta_Inv()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsBeta_InvRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsBeta_InvRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Bin2Dec provides operations to call the bin2Dec method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsBin2DecRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Bin2Dec()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsBin2DecRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsBin2DecRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Bin2Hex provides operations to call the bin2Hex method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsBin2HexRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Bin2Hex()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsBin2HexRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsBin2HexRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Bin2Oct provides operations to call the bin2Oct method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsBin2OctRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Bin2Oct()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsBin2OctRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsBin2OctRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Binom_Dist provides operations to call the binom_Dist method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsBinom_DistRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Binom_Dist()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsBinom_DistRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsBinom_DistRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Binom_Dist_Range provides operations to call the binom_Dist_Range method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsBinom_Dist_RangeRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Binom_Dist_Range()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsBinom_Dist_RangeRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsBinom_Dist_RangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Binom_Inv provides operations to call the binom_Inv method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsBinom_InvRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Binom_Inv()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsBinom_InvRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsBinom_InvRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Bitand provides operations to call the bitand method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsBitandRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Bitand()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsBitandRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsBitandRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Bitlshift provides operations to call the bitlshift method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsBitlshiftRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Bitlshift()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsBitlshiftRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsBitlshiftRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Bitor provides operations to call the bitor method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsBitorRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Bitor()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsBitorRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsBitorRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Bitrshift provides operations to call the bitrshift method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsBitrshiftRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Bitrshift()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsBitrshiftRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsBitrshiftRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Bitxor provides operations to call the bitxor method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsBitxorRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Bitxor()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsBitxorRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsBitxorRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Ceiling_Math provides operations to call the ceiling_Math method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsCeiling_MathRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Ceiling_Math()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsCeiling_MathRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsCeiling_MathRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Ceiling_Precise provides operations to call the ceiling_Precise method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsCeiling_PreciseRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Ceiling_Precise()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsCeiling_PreciseRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsCeiling_PreciseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Char provides operations to call the char method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsCharRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Char()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsCharRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsCharRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChiSq_Dist provides operations to call the chiSq_Dist method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsChiSq_DistRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) ChiSq_Dist()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsChiSq_DistRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsChiSq_DistRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChiSq_Dist_RT provides operations to call the chiSq_Dist_RT method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsChiSq_Dist_RTRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) ChiSq_Dist_RT()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsChiSq_Dist_RTRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsChiSq_Dist_RTRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChiSq_Inv provides operations to call the chiSq_Inv method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsChiSq_InvRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) ChiSq_Inv()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsChiSq_InvRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsChiSq_InvRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChiSq_Inv_RT provides operations to call the chiSq_Inv_RT method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsChiSq_Inv_RTRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) ChiSq_Inv_RT()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsChiSq_Inv_RTRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsChiSq_Inv_RTRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Choose provides operations to call the choose method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsChooseRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Choose()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsChooseRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsChooseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Clean provides operations to call the clean method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsCleanRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Clean()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsCleanRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsCleanRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Code provides operations to call the code method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsCodeRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Code()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsCodeRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsCodeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Columns provides operations to call the columns method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsColumnsRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Columns()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsColumnsRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsColumnsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Combin provides operations to call the combin method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsCombinRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Combin()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsCombinRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsCombinRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Combina provides operations to call the combina method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsCombinaRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Combina()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsCombinaRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsCombinaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Complex provides operations to call the complex method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsComplexRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Complex()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsComplexRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsComplexRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Concatenate provides operations to call the concatenate method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsConcatenateRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Concatenate()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsConcatenateRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsConcatenateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Confidence_Norm provides operations to call the confidence_Norm method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsConfidence_NormRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Confidence_Norm()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsConfidence_NormRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsConfidence_NormRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Confidence_T provides operations to call the confidence_T method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsConfidence_TRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Confidence_T()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsConfidence_TRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsConfidence_TRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilderInternal instantiates a new FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) {
    m := &FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/functions{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder instantiates a new FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Convert provides operations to call the convert method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsConvertRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Convert()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsConvertRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsConvertRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cos provides operations to call the cos method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsCosRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Cos()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsCosRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsCosRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cosh provides operations to call the cosh method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsCoshRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Cosh()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsCoshRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsCoshRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cot provides operations to call the cot method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsCotRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Cot()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsCotRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsCotRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Coth provides operations to call the coth method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsCothRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Coth()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsCothRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsCothRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Count provides operations to call the count method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsCountRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Count()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsCountRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CountA provides operations to call the countA method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsCountARequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) CountA()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsCountARequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsCountARequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CountBlank provides operations to call the countBlank method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsCountBlankRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) CountBlank()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsCountBlankRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsCountBlankRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CountIf provides operations to call the countIf method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsCountIfRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) CountIf()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsCountIfRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsCountIfRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CountIfs provides operations to call the countIfs method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsCountIfsRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) CountIfs()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsCountIfsRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsCountIfsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CoupDayBs provides operations to call the coupDayBs method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsCoupDayBsRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) CoupDayBs()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsCoupDayBsRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsCoupDayBsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CoupDays provides operations to call the coupDays method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsCoupDaysRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) CoupDays()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsCoupDaysRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsCoupDaysRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CoupDaysNc provides operations to call the coupDaysNc method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsCoupDaysNcRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) CoupDaysNc()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsCoupDaysNcRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsCoupDaysNcRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CoupNcd provides operations to call the coupNcd method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsCoupNcdRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) CoupNcd()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsCoupNcdRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsCoupNcdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CoupNum provides operations to call the coupNum method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsCoupNumRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) CoupNum()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsCoupNumRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsCoupNumRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CoupPcd provides operations to call the coupPcd method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsCoupPcdRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) CoupPcd()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsCoupPcdRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsCoupPcdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Csc provides operations to call the csc method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsCscRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Csc()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsCscRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsCscRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Csch provides operations to call the csch method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsCschRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Csch()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsCschRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsCschRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CumIPmt provides operations to call the cumIPmt method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsCumIPmtRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) CumIPmt()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsCumIPmtRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsCumIPmtRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CumPrinc provides operations to call the cumPrinc method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsCumPrincRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) CumPrinc()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsCumPrincRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsCumPrincRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Date provides operations to call the date method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsDateRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Date()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsDateRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsDateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Datevalue provides operations to call the datevalue method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsDatevalueRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Datevalue()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsDatevalueRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsDatevalueRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Daverage provides operations to call the daverage method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsDaverageRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Daverage()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsDaverageRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsDaverageRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Day provides operations to call the day method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsDayRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Day()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsDayRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsDayRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Days provides operations to call the days method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsDaysRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Days()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsDaysRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsDaysRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Days360 provides operations to call the days360 method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsDays360RequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Days360()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsDays360RequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsDays360RequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Db provides operations to call the db method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsDbRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Db()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsDbRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsDbRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Dbcs provides operations to call the dbcs method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsDbcsRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Dbcs()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsDbcsRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsDbcsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Dcount provides operations to call the dcount method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsDcountRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Dcount()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsDcountRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsDcountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DcountA provides operations to call the dcountA method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsDcountARequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) DcountA()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsDcountARequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsDcountARequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Ddb provides operations to call the ddb method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsDdbRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Ddb()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsDdbRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsDdbRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Dec2Bin provides operations to call the dec2Bin method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsDec2BinRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Dec2Bin()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsDec2BinRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsDec2BinRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Dec2Hex provides operations to call the dec2Hex method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsDec2HexRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Dec2Hex()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsDec2HexRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsDec2HexRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Dec2Oct provides operations to call the dec2Oct method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsDec2OctRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Dec2Oct()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsDec2OctRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsDec2OctRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Decimal provides operations to call the decimal method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsDecimalRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Decimal()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsDecimalRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsDecimalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Degrees provides operations to call the degrees method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsDegreesRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Degrees()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsDegreesRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsDegreesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property functions for storage
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Delete(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Delta provides operations to call the delta method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsDeltaRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Delta()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsDeltaRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsDeltaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DevSq provides operations to call the devSq method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsDevSqRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) DevSq()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsDevSqRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsDevSqRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Dget provides operations to call the dget method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsDgetRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Dget()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsDgetRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsDgetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Disc provides operations to call the disc method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsDiscRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Disc()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsDiscRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsDiscRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Dmax provides operations to call the dmax method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsDmaxRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Dmax()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsDmaxRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsDmaxRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Dmin provides operations to call the dmin method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsDminRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Dmin()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsDminRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsDminRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Dollar provides operations to call the dollar method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsDollarRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Dollar()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsDollarRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsDollarRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DollarDe provides operations to call the dollarDe method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsDollarDeRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) DollarDe()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsDollarDeRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsDollarDeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DollarFr provides operations to call the dollarFr method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsDollarFrRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) DollarFr()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsDollarFrRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsDollarFrRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Dproduct provides operations to call the dproduct method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsDproductRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Dproduct()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsDproductRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsDproductRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DstDev provides operations to call the dstDev method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsDstDevRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) DstDev()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsDstDevRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsDstDevRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DstDevP provides operations to call the dstDevP method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsDstDevPRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) DstDevP()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsDstDevPRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsDstDevPRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Dsum provides operations to call the dsum method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsDsumRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Dsum()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsDsumRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsDsumRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Duration provides operations to call the duration method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsDurationRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Duration()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsDurationRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsDurationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Dvar provides operations to call the dvar method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsDvarRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Dvar()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsDvarRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsDvarRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DvarP provides operations to call the dvarP method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsDvarPRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) DvarP()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsDvarPRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsDvarPRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Ecma_Ceiling provides operations to call the ecma_Ceiling method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsEcma_CeilingRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Ecma_Ceiling()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsEcma_CeilingRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsEcma_CeilingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Edate provides operations to call the edate method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsEdateRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Edate()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsEdateRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsEdateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Effect provides operations to call the effect method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsEffectRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Effect()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsEffectRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsEffectRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EoMonth provides operations to call the eoMonth method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsEoMonthRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) EoMonth()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsEoMonthRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsEoMonthRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Erf provides operations to call the erf method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsErfRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Erf()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsErfRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsErfRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Erf_Precise provides operations to call the erf_Precise method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsErf_PreciseRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Erf_Precise()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsErf_PreciseRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsErf_PreciseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ErfC provides operations to call the erfC method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsErfCRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) ErfC()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsErfCRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsErfCRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ErfC_Precise provides operations to call the erfC_Precise method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsErfC_PreciseRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) ErfC_Precise()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsErfC_PreciseRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsErfC_PreciseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Error_Type provides operations to call the error_Type method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsError_TypeRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Error_Type()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsError_TypeRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsError_TypeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Even provides operations to call the even method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsEvenRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Even()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsEvenRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsEvenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Exact provides operations to call the exact method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsExactRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Exact()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsExactRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsExactRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Exp provides operations to call the exp method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsExpRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Exp()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsExpRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsExpRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Expon_Dist provides operations to call the expon_Dist method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsExpon_DistRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Expon_Dist()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsExpon_DistRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsExpon_DistRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// F_Dist provides operations to call the f_Dist method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsF_DistRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) F_Dist()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsF_DistRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsF_DistRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// F_Dist_RT provides operations to call the f_Dist_RT method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsF_Dist_RTRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) F_Dist_RT()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsF_Dist_RTRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsF_Dist_RTRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// F_Inv provides operations to call the f_Inv method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsF_InvRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) F_Inv()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsF_InvRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsF_InvRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// F_Inv_RT provides operations to call the f_Inv_RT method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsF_Inv_RTRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) F_Inv_RT()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsF_Inv_RTRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsF_Inv_RTRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Fact provides operations to call the fact method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsFactRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Fact()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsFactRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsFactRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FactDouble provides operations to call the factDouble method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsFactDoubleRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) FactDouble()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsFactDoubleRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsFactDoubleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// False provides operations to call the false method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsFalseRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) False()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsFalseRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsFalseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Find provides operations to call the find method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsFindRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Find()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsFindRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsFindRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FindB provides operations to call the findB method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsFindBRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) FindB()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsFindBRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsFindBRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Fisher provides operations to call the fisher method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsFisherRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Fisher()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsFisherRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsFisherRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FisherInv provides operations to call the fisherInv method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsFisherInvRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) FisherInv()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsFisherInvRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsFisherInvRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Fixed provides operations to call the fixed method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsFixedRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Fixed()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsFixedRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsFixedRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Floor_Math provides operations to call the floor_Math method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsFloor_MathRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Floor_Math()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsFloor_MathRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsFloor_MathRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Floor_Precise provides operations to call the floor_Precise method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsFloor_PreciseRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Floor_Precise()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsFloor_PreciseRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsFloor_PreciseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Fv provides operations to call the fv method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsFvRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Fv()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsFvRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsFvRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Fvschedule provides operations to call the fvschedule method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsFvscheduleRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Fvschedule()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsFvscheduleRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsFvscheduleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Gamma provides operations to call the gamma method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsGammaRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Gamma()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsGammaRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsGammaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Gamma_Dist provides operations to call the gamma_Dist method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsGamma_DistRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Gamma_Dist()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsGamma_DistRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsGamma_DistRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Gamma_Inv provides operations to call the gamma_Inv method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsGamma_InvRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Gamma_Inv()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsGamma_InvRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsGamma_InvRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GammaLn provides operations to call the gammaLn method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsGammaLnRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) GammaLn()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsGammaLnRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsGammaLnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GammaLn_Precise provides operations to call the gammaLn_Precise method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsGammaLn_PreciseRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) GammaLn_Precise()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsGammaLn_PreciseRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsGammaLn_PreciseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Gauss provides operations to call the gauss method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsGaussRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Gauss()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsGaussRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsGaussRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Gcd provides operations to call the gcd method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsGcdRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Gcd()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsGcdRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsGcdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GeoMean provides operations to call the geoMean method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsGeoMeanRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) GeoMean()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsGeoMeanRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsGeoMeanRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GeStep provides operations to call the geStep method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsGeStepRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) GeStep()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsGeStepRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsGeStepRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get functions from storage
// returns a WorkbookFunctionsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFunctionsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookFunctionsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFunctionsable), nil
}
// HarMean provides operations to call the harMean method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsHarMeanRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) HarMean()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsHarMeanRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsHarMeanRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Hex2Bin provides operations to call the hex2Bin method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsHex2BinRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Hex2Bin()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsHex2BinRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsHex2BinRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Hex2Dec provides operations to call the hex2Dec method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsHex2DecRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Hex2Dec()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsHex2DecRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsHex2DecRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Hex2Oct provides operations to call the hex2Oct method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsHex2OctRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Hex2Oct()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsHex2OctRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsHex2OctRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Hlookup provides operations to call the hlookup method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsHlookupRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Hlookup()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsHlookupRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsHlookupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Hour provides operations to call the hour method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsHourRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Hour()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsHourRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsHourRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Hyperlink provides operations to call the hyperlink method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsHyperlinkRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Hyperlink()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsHyperlinkRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsHyperlinkRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// HypGeom_Dist provides operations to call the hypGeom_Dist method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsHypGeom_DistRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) HypGeom_Dist()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsHypGeom_DistRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsHypGeom_DistRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IfEscaped provides operations to call the if method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsIfRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) IfEscaped()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsIfRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsIfRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImAbs provides operations to call the imAbs method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsImAbsRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) ImAbs()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsImAbsRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsImAbsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Imaginary provides operations to call the imaginary method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsImaginaryRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Imaginary()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsImaginaryRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsImaginaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImArgument provides operations to call the imArgument method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsImArgumentRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) ImArgument()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsImArgumentRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsImArgumentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImConjugate provides operations to call the imConjugate method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsImConjugateRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) ImConjugate()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsImConjugateRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsImConjugateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImCos provides operations to call the imCos method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsImCosRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) ImCos()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsImCosRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsImCosRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImCosh provides operations to call the imCosh method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsImCoshRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) ImCosh()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsImCoshRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsImCoshRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImCot provides operations to call the imCot method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsImCotRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) ImCot()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsImCotRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsImCotRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImCsc provides operations to call the imCsc method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsImCscRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) ImCsc()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsImCscRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsImCscRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImCsch provides operations to call the imCsch method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsImCschRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) ImCsch()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsImCschRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsImCschRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImDiv provides operations to call the imDiv method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsImDivRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) ImDiv()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsImDivRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsImDivRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImExp provides operations to call the imExp method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsImExpRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) ImExp()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsImExpRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsImExpRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImLn provides operations to call the imLn method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsImLnRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) ImLn()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsImLnRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsImLnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImLog10 provides operations to call the imLog10 method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsImLog10RequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) ImLog10()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsImLog10RequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsImLog10RequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImLog2 provides operations to call the imLog2 method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsImLog2RequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) ImLog2()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsImLog2RequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsImLog2RequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImPower provides operations to call the imPower method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsImPowerRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) ImPower()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsImPowerRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsImPowerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImProduct provides operations to call the imProduct method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsImProductRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) ImProduct()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsImProductRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsImProductRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImReal provides operations to call the imReal method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsImRealRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) ImReal()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsImRealRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsImRealRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImSec provides operations to call the imSec method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsImSecRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) ImSec()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsImSecRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsImSecRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImSech provides operations to call the imSech method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsImSechRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) ImSech()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsImSechRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsImSechRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImSin provides operations to call the imSin method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsImSinRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) ImSin()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsImSinRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsImSinRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImSinh provides operations to call the imSinh method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsImSinhRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) ImSinh()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsImSinhRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsImSinhRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImSqrt provides operations to call the imSqrt method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsImSqrtRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) ImSqrt()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsImSqrtRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsImSqrtRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImSub provides operations to call the imSub method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsImSubRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) ImSub()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsImSubRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsImSubRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImSum provides operations to call the imSum method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsImSumRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) ImSum()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsImSumRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsImSumRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImTan provides operations to call the imTan method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsImTanRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) ImTan()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsImTanRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsImTanRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Int provides operations to call the int method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsIntRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Int()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsIntRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsIntRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IntRate provides operations to call the intRate method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsIntRateRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) IntRate()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsIntRateRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsIntRateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Ipmt provides operations to call the ipmt method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsIpmtRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Ipmt()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsIpmtRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsIpmtRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Irr provides operations to call the irr method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsIrrRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Irr()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsIrrRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsIrrRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IsErr provides operations to call the isErr method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsIsErrRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) IsErr()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsIsErrRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsIsErrRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IsError provides operations to call the isError method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsIsErrorRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) IsError()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsIsErrorRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsIsErrorRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IsEven provides operations to call the isEven method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsIsEvenRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) IsEven()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsIsEvenRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsIsEvenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IsFormula provides operations to call the isFormula method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsIsFormulaRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) IsFormula()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsIsFormulaRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsIsFormulaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IsLogical provides operations to call the isLogical method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsIsLogicalRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) IsLogical()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsIsLogicalRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsIsLogicalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IsNA provides operations to call the isNA method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsIsNARequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) IsNA()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsIsNARequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsIsNARequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IsNonText provides operations to call the isNonText method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsIsNonTextRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) IsNonText()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsIsNonTextRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsIsNonTextRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IsNumber provides operations to call the isNumber method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsIsNumberRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) IsNumber()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsIsNumberRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsIsNumberRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Iso_Ceiling provides operations to call the iso_Ceiling method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsIso_CeilingRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Iso_Ceiling()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsIso_CeilingRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsIso_CeilingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IsOdd provides operations to call the isOdd method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsIsOddRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) IsOdd()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsIsOddRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsIsOddRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IsoWeekNum provides operations to call the isoWeekNum method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsIsoWeekNumRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) IsoWeekNum()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsIsoWeekNumRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsIsoWeekNumRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Ispmt provides operations to call the ispmt method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsIspmtRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Ispmt()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsIspmtRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsIspmtRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Isref provides operations to call the isref method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsIsrefRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Isref()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsIsrefRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsIsrefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IsText provides operations to call the isText method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsIsTextRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) IsText()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsIsTextRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsIsTextRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Kurt provides operations to call the kurt method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsKurtRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Kurt()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsKurtRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsKurtRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Large provides operations to call the large method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsLargeRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Large()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsLargeRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsLargeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Lcm provides operations to call the lcm method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsLcmRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Lcm()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsLcmRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsLcmRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Left provides operations to call the left method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsLeftRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Left()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsLeftRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsLeftRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Leftb provides operations to call the leftb method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsLeftbRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Leftb()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsLeftbRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsLeftbRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Len provides operations to call the len method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsLenRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Len()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsLenRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsLenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Lenb provides operations to call the lenb method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsLenbRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Lenb()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsLenbRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsLenbRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Ln provides operations to call the ln method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsLnRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Ln()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsLnRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsLnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Log provides operations to call the log method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsLogRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Log()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsLogRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsLogRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Log10 provides operations to call the log10 method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsLog10RequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Log10()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsLog10RequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsLog10RequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LogNorm_Dist provides operations to call the logNorm_Dist method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsLogNorm_DistRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) LogNorm_Dist()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsLogNorm_DistRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsLogNorm_DistRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LogNorm_Inv provides operations to call the logNorm_Inv method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) LogNorm_Inv()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Lookup provides operations to call the lookup method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsLookupRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Lookup()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsLookupRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsLookupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Lower provides operations to call the lower method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsLowerRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Lower()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsLowerRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsLowerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Match provides operations to call the match method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsMatchRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Match()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsMatchRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsMatchRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Max provides operations to call the max method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsMaxRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Max()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsMaxRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsMaxRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MaxA provides operations to call the maxA method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsMaxARequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) MaxA()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsMaxARequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsMaxARequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Mduration provides operations to call the mduration method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsMdurationRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Mduration()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsMdurationRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsMdurationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Median provides operations to call the median method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsMedianRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Median()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsMedianRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsMedianRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Mid provides operations to call the mid method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsMidRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Mid()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsMidRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsMidRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Midb provides operations to call the midb method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsMidbRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Midb()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsMidbRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsMidbRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Min provides operations to call the min method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsMinRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Min()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsMinRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsMinRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MinA provides operations to call the minA method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsMinARequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) MinA()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsMinARequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsMinARequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Minute provides operations to call the minute method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsMinuteRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Minute()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsMinuteRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsMinuteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Mirr provides operations to call the mirr method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsMirrRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Mirr()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsMirrRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsMirrRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Mod provides operations to call the mod method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsModRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Mod()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsModRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsModRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Month provides operations to call the month method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsMonthRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Month()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsMonthRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsMonthRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Mround provides operations to call the mround method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsMroundRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Mround()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsMroundRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsMroundRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MultiNomial provides operations to call the multiNomial method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsMultiNomialRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) MultiNomial()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsMultiNomialRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsMultiNomialRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// N provides operations to call the n method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsNRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) N()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsNRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsNRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Na provides operations to call the na method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsNaRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Na()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsNaRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsNaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NegBinom_Dist provides operations to call the negBinom_Dist method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsNegBinom_DistRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) NegBinom_Dist()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsNegBinom_DistRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsNegBinom_DistRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NetworkDays provides operations to call the networkDays method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsNetworkDaysRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) NetworkDays()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsNetworkDaysRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsNetworkDaysRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NetworkDays_Intl provides operations to call the networkDays_Intl method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsNetworkDays_IntlRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) NetworkDays_Intl()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsNetworkDays_IntlRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsNetworkDays_IntlRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Nominal provides operations to call the nominal method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsNominalRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Nominal()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsNominalRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsNominalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Norm_Dist provides operations to call the norm_Dist method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsNorm_DistRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Norm_Dist()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsNorm_DistRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsNorm_DistRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Norm_Inv provides operations to call the norm_Inv method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsNorm_InvRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Norm_Inv()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsNorm_InvRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsNorm_InvRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Norm_S_Dist provides operations to call the norm_S_Dist method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsNorm_S_DistRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Norm_S_Dist()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsNorm_S_DistRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsNorm_S_DistRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Norm_S_Inv provides operations to call the norm_S_Inv method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsNorm_S_InvRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Norm_S_Inv()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsNorm_S_InvRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsNorm_S_InvRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Not provides operations to call the not method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsNotRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Not()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsNotRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsNotRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Now provides operations to call the now method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsNowRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Now()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsNowRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsNowRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Nper provides operations to call the nper method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsNperRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Nper()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsNperRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsNperRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Npv provides operations to call the npv method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsNpvRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Npv()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsNpvRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsNpvRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NumberValue provides operations to call the numberValue method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsNumberValueRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) NumberValue()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsNumberValueRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsNumberValueRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Oct2Bin provides operations to call the oct2Bin method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsOct2BinRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Oct2Bin()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsOct2BinRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsOct2BinRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Oct2Dec provides operations to call the oct2Dec method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsOct2DecRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Oct2Dec()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsOct2DecRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsOct2DecRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Oct2Hex provides operations to call the oct2Hex method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsOct2HexRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Oct2Hex()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsOct2HexRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsOct2HexRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Odd provides operations to call the odd method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsOddRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Odd()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsOddRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsOddRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OddFPrice provides operations to call the oddFPrice method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsOddFPriceRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) OddFPrice()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsOddFPriceRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsOddFPriceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OddFYield provides operations to call the oddFYield method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsOddFYieldRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) OddFYield()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsOddFYieldRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsOddFYieldRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OddLPrice provides operations to call the oddLPrice method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsOddLPriceRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) OddLPrice()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsOddLPriceRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsOddLPriceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OddLYield provides operations to call the oddLYield method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsOddLYieldRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) OddLYield()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsOddLYieldRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsOddLYieldRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Or provides operations to call the or method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsOrRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Or()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsOrRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsOrRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property functions in storage
// returns a WorkbookFunctionsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFunctionsable, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFunctionsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookFunctionsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFunctionsable), nil
}
// Pduration provides operations to call the pduration method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsPdurationRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Pduration()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsPdurationRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsPdurationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Percentile_Exc provides operations to call the percentile_Exc method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsPercentile_ExcRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Percentile_Exc()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsPercentile_ExcRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsPercentile_ExcRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Percentile_Inc provides operations to call the percentile_Inc method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsPercentile_IncRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Percentile_Inc()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsPercentile_IncRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsPercentile_IncRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PercentRank_Exc provides operations to call the percentRank_Exc method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsPercentRank_ExcRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) PercentRank_Exc()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsPercentRank_ExcRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsPercentRank_ExcRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PercentRank_Inc provides operations to call the percentRank_Inc method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsPercentRank_IncRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) PercentRank_Inc()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsPercentRank_IncRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsPercentRank_IncRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Permut provides operations to call the permut method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsPermutRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Permut()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsPermutRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsPermutRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Permutationa provides operations to call the permutationa method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsPermutationaRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Permutationa()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsPermutationaRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsPermutationaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Phi provides operations to call the phi method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsPhiRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Phi()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsPhiRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsPhiRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Pi provides operations to call the pi method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsPiRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Pi()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsPiRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsPiRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Pmt provides operations to call the pmt method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsPmtRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Pmt()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsPmtRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsPmtRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Poisson_Dist provides operations to call the poisson_Dist method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsPoisson_DistRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Poisson_Dist()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsPoisson_DistRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsPoisson_DistRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Power provides operations to call the power method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsPowerRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Power()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsPowerRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsPowerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Ppmt provides operations to call the ppmt method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsPpmtRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Ppmt()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsPpmtRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsPpmtRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Price provides operations to call the price method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsPriceRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Price()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsPriceRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsPriceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PriceDisc provides operations to call the priceDisc method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsPriceDiscRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) PriceDisc()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsPriceDiscRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsPriceDiscRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PriceMat provides operations to call the priceMat method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsPriceMatRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) PriceMat()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsPriceMatRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsPriceMatRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Product provides operations to call the product method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsProductRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Product()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsProductRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsProductRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Proper provides operations to call the proper method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsProperRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Proper()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsProperRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsProperRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Pv provides operations to call the pv method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsPvRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Pv()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsPvRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsPvRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Quartile_Exc provides operations to call the quartile_Exc method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsQuartile_ExcRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Quartile_Exc()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsQuartile_ExcRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsQuartile_ExcRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Quartile_Inc provides operations to call the quartile_Inc method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsQuartile_IncRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Quartile_Inc()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsQuartile_IncRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsQuartile_IncRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Quotient provides operations to call the quotient method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsQuotientRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Quotient()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsQuotientRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsQuotientRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Radians provides operations to call the radians method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRadiansRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Radians()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsRadiansRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsRadiansRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rand provides operations to call the rand method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRandRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Rand()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsRandRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsRandRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RandBetween provides operations to call the randBetween method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRandBetweenRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) RandBetween()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsRandBetweenRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsRandBetweenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rank_Avg provides operations to call the rank_Avg method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRank_AvgRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Rank_Avg()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsRank_AvgRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsRank_AvgRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rank_Eq provides operations to call the rank_Eq method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRank_EqRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Rank_Eq()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsRank_EqRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsRank_EqRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rate provides operations to call the rate method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRateRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Rate()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsRateRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsRateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Received provides operations to call the received method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsReceivedRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Received()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsReceivedRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsReceivedRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Replace provides operations to call the replace method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsReplaceRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Replace()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsReplaceRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsReplaceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ReplaceB provides operations to call the replaceB method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsReplaceBRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) ReplaceB()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsReplaceBRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsReplaceBRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rept provides operations to call the rept method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsReptRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Rept()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsReptRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsReptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Right provides operations to call the right method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRightRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Right()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsRightRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsRightRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rightb provides operations to call the rightb method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRightbRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Rightb()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsRightbRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsRightbRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Roman provides operations to call the roman method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRomanRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Roman()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsRomanRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsRomanRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Round provides operations to call the round method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRoundRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Round()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsRoundRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsRoundRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoundDown provides operations to call the roundDown method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRoundDownRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) RoundDown()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsRoundDownRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsRoundDownRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoundUp provides operations to call the roundUp method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRoundUpRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) RoundUp()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsRoundUpRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsRoundUpRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rows provides operations to call the rows method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRowsRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Rows()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsRowsRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsRowsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rri provides operations to call the rri method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRriRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Rri()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsRriRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsRriRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sec provides operations to call the sec method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsSecRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Sec()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsSecRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsSecRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sech provides operations to call the sech method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsSechRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Sech()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsSechRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsSechRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Second provides operations to call the second method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsSecondRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Second()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsSecondRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsSecondRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SeriesSum provides operations to call the seriesSum method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsSeriesSumRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) SeriesSum()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsSeriesSumRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsSeriesSumRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sheet provides operations to call the sheet method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsSheetRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Sheet()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsSheetRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsSheetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sheets provides operations to call the sheets method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsSheetsRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Sheets()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsSheetsRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsSheetsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sign provides operations to call the sign method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsSignRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Sign()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsSignRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsSignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sin provides operations to call the sin method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsSinRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Sin()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsSinRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsSinRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sinh provides operations to call the sinh method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsSinhRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Sinh()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsSinhRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsSinhRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Skew provides operations to call the skew method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsSkewRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Skew()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsSkewRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsSkewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Skew_p provides operations to call the skew_p method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsSkew_pRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Skew_p()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsSkew_pRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsSkew_pRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sln provides operations to call the sln method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsSlnRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Sln()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsSlnRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsSlnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Small provides operations to call the small method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsSmallRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Small()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsSmallRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsSmallRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sqrt provides operations to call the sqrt method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsSqrtRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Sqrt()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsSqrtRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsSqrtRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SqrtPi provides operations to call the sqrtPi method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsSqrtPiRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) SqrtPi()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsSqrtPiRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsSqrtPiRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Standardize provides operations to call the standardize method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsStandardizeRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Standardize()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsStandardizeRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsStandardizeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// StDev_P provides operations to call the stDev_P method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsStDev_PRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) StDev_P()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsStDev_PRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsStDev_PRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// StDev_S provides operations to call the stDev_S method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsStDev_SRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) StDev_S()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsStDev_SRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsStDev_SRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// StDevA provides operations to call the stDevA method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsStDevARequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) StDevA()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsStDevARequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsStDevARequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// StDevPA provides operations to call the stDevPA method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsStDevPARequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) StDevPA()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsStDevPARequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsStDevPARequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Substitute provides operations to call the substitute method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsSubstituteRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Substitute()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsSubstituteRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsSubstituteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Subtotal provides operations to call the subtotal method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsSubtotalRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Subtotal()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsSubtotalRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsSubtotalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sum provides operations to call the sum method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsSumRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Sum()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsSumRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsSumRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SumIf provides operations to call the sumIf method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsSumIfRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) SumIf()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsSumIfRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsSumIfRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SumIfs provides operations to call the sumIfs method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsSumIfsRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) SumIfs()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsSumIfsRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsSumIfsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SumSq provides operations to call the sumSq method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsSumSqRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) SumSq()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsSumSqRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsSumSqRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Syd provides operations to call the syd method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsSydRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Syd()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsSydRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsSydRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// T provides operations to call the t method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsTRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) T()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsTRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsTRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// T_Dist provides operations to call the t_Dist method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsT_DistRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) T_Dist()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsT_DistRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsT_DistRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// T_Dist_2T provides operations to call the t_Dist_2T method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsT_Dist_2TRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) T_Dist_2T()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsT_Dist_2TRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsT_Dist_2TRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// T_Dist_RT provides operations to call the t_Dist_RT method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsT_Dist_RTRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) T_Dist_RT()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsT_Dist_RTRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsT_Dist_RTRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// T_Inv provides operations to call the t_Inv method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsT_InvRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) T_Inv()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsT_InvRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsT_InvRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// T_Inv_2T provides operations to call the t_Inv_2T method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsT_Inv_2TRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) T_Inv_2T()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsT_Inv_2TRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsT_Inv_2TRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tan provides operations to call the tan method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsTanRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Tan()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsTanRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsTanRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tanh provides operations to call the tanh method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsTanhRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Tanh()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsTanhRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsTanhRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TbillEq provides operations to call the tbillEq method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsTbillEqRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) TbillEq()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsTbillEqRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsTbillEqRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TbillPrice provides operations to call the tbillPrice method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsTbillPriceRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) TbillPrice()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsTbillPriceRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsTbillPriceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TbillYield provides operations to call the tbillYield method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsTbillYieldRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) TbillYield()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsTbillYieldRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsTbillYieldRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Text provides operations to call the text method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsTextRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Text()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsTextRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsTextRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Time provides operations to call the time method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsTimeRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Time()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsTimeRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsTimeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Timevalue provides operations to call the timevalue method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsTimevalueRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Timevalue()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsTimevalueRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsTimevalueRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Today provides operations to call the today method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsTodayRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Today()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsTodayRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsTodayRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property functions for storage
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get functions from storage
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property functions in storage
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFunctionsable, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// Trim provides operations to call the trim method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsTrimRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Trim()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsTrimRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsTrimRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TrimMean provides operations to call the trimMean method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsTrimMeanRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) TrimMean()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsTrimMeanRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsTrimMeanRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// True provides operations to call the true method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsTrueRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) True()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsTrueRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsTrueRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Trunc provides operations to call the trunc method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsTruncRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Trunc()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsTruncRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsTruncRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TypeEscaped provides operations to call the type method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsTypeRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) TypeEscaped()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsTypeRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsTypeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Unichar provides operations to call the unichar method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsUnicharRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Unichar()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsUnicharRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsUnicharRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Unicode provides operations to call the unicode method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsUnicodeRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Unicode()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsUnicodeRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsUnicodeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Upper provides operations to call the upper method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsUpperRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Upper()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsUpperRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsUpperRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Usdollar provides operations to call the usdollar method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsUsdollarRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Usdollar()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsUsdollarRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsUsdollarRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Value provides operations to call the value method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsValueRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Value()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsValueRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsValueRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Var_P provides operations to call the var_P method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsVar_PRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Var_P()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsVar_PRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsVar_PRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Var_S provides operations to call the var_S method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsVar_SRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Var_S()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsVar_SRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsVar_SRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// VarA provides operations to call the varA method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsVarARequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) VarA()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsVarARequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsVarARequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// VarPA provides operations to call the varPA method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsVarPARequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) VarPA()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsVarPARequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsVarPARequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Vdb provides operations to call the vdb method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsVdbRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Vdb()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsVdbRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsVdbRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Vlookup provides operations to call the vlookup method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsVlookupRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Vlookup()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsVlookupRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsVlookupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Weekday provides operations to call the weekday method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsWeekdayRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Weekday()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsWeekdayRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsWeekdayRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WeekNum provides operations to call the weekNum method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsWeekNumRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) WeekNum()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsWeekNumRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsWeekNumRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Weibull_Dist provides operations to call the weibull_Dist method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsWeibull_DistRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Weibull_Dist()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsWeibull_DistRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsWeibull_DistRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
// WorkDay provides operations to call the workDay method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsWorkDayRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) WorkDay()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsWorkDayRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsWorkDayRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WorkDay_Intl provides operations to call the workDay_Intl method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsWorkDay_IntlRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) WorkDay_Intl()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsWorkDay_IntlRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsWorkDay_IntlRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Xirr provides operations to call the xirr method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsXirrRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Xirr()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsXirrRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsXirrRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Xnpv provides operations to call the xnpv method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsXnpvRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Xnpv()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsXnpvRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsXnpvRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Xor provides operations to call the xor method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsXorRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Xor()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsXorRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsXorRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Year provides operations to call the year method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsYearRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Year()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsYearRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsYearRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// YearFrac provides operations to call the yearFrac method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsYearFracRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) YearFrac()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsYearFracRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsYearFracRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Yield provides operations to call the yield method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsYieldRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Yield()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsYieldRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsYieldRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// YieldDisc provides operations to call the yieldDisc method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsYieldDiscRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) YieldDisc()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsYieldDiscRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsYieldDiscRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// YieldMat provides operations to call the yieldMat method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsYieldMatRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) YieldMat()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsYieldMatRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsYieldMatRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Z_Test provides operations to call the z_Test method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsZ_TestRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) Z_Test()(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsZ_TestRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsZ_TestRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
