package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageAssignmentable 
type AccessPackageAssignmentable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAccessPackage()(AccessPackageable)
    GetExpiredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSchedule()(EntitlementManagementScheduleable)
    GetState()(*AccessPackageAssignmentState)
    GetStatus()(*string)
    GetTarget()(AccessPackageSubjectable)
    SetAccessPackage(value AccessPackageable)()
    SetExpiredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSchedule(value EntitlementManagementScheduleable)()
    SetState(value *AccessPackageAssignmentState)()
    SetStatus(value *string)()
    SetTarget(value AccessPackageSubjectable)()
}
