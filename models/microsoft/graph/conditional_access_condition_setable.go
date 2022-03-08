package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ConditionalAccessConditionSetable 
type ConditionalAccessConditionSetable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetApplications()(ConditionalAccessApplicationsable)
    GetClientAppTypes()([]ConditionalAccessClientApp)
    GetDevices()(ConditionalAccessDevicesable)
    GetLocations()(ConditionalAccessLocationsable)
    GetPlatforms()(ConditionalAccessPlatformsable)
    GetSignInRiskLevels()([]RiskLevel)
    GetUserRiskLevels()([]RiskLevel)
    GetUsers()(ConditionalAccessUsersable)
    SetApplications(value ConditionalAccessApplicationsable)()
    SetClientAppTypes(value []ConditionalAccessClientApp)()
    SetDevices(value ConditionalAccessDevicesable)()
    SetLocations(value ConditionalAccessLocationsable)()
    SetPlatforms(value ConditionalAccessPlatformsable)()
    SetSignInRiskLevels(value []RiskLevel)()
    SetUserRiskLevels(value []RiskLevel)()
    SetUsers(value ConditionalAccessUsersable)()
}
