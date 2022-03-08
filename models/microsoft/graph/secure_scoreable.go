package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SecureScoreable 
type SecureScoreable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetActiveUserCount()(*int32)
    GetAverageComparativeScores()([]AverageComparativeScoreable)
    GetAzureTenantId()(*string)
    GetControlScores()([]ControlScoreable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCurrentScore()(*float64)
    GetEnabledServices()([]string)
    GetLicensedUserCount()(*int32)
    GetMaxScore()(*float64)
    GetVendorInformation()(SecurityVendorInformationable)
    SetActiveUserCount(value *int32)()
    SetAverageComparativeScores(value []AverageComparativeScoreable)()
    SetAzureTenantId(value *string)()
    SetControlScores(value []ControlScoreable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCurrentScore(value *float64)()
    SetEnabledServices(value []string)()
    SetLicensedUserCount(value *int32)()
    SetMaxScore(value *float64)()
    SetVendorInformation(value SecurityVendorInformationable)()
}
