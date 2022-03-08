package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageAssignmentRequestable 
type AccessPackageAssignmentRequestable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetAccessPackage()(AccessPackageable)
    GetAssignment()(AccessPackageAssignmentable)
    GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRequestor()(AccessPackageSubjectable)
    GetRequestType()(*AccessPackageRequestType)
    GetSchedule()(EntitlementManagementScheduleable)
    GetState()(*AccessPackageRequestState)
    GetStatus()(*string)
    SetAccessPackage(value AccessPackageable)()
    SetAssignment(value AccessPackageAssignmentable)()
    SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRequestor(value AccessPackageSubjectable)()
    SetRequestType(value *AccessPackageRequestType)()
    SetSchedule(value EntitlementManagementScheduleable)()
    SetState(value *AccessPackageRequestState)()
    SetStatus(value *string)()
}
