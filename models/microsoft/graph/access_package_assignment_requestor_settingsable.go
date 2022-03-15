package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageAssignmentRequestorSettingsable 
type AccessPackageAssignmentRequestorSettingsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAllowCustomAssignmentSchedule()(*bool)
    GetEnableOnBehalfRequestorsToAddAccess()(*bool)
    GetEnableOnBehalfRequestorsToRemoveAccess()(*bool)
    GetEnableOnBehalfRequestorsToUpdateAccess()(*bool)
    GetEnableTargetsToSelfAddAccess()(*bool)
    GetEnableTargetsToSelfRemoveAccess()(*bool)
    GetEnableTargetsToSelfUpdateAccess()(*bool)
    GetOnBehalfRequestors()([]SubjectSetable)
    SetAllowCustomAssignmentSchedule(value *bool)()
    SetEnableOnBehalfRequestorsToAddAccess(value *bool)()
    SetEnableOnBehalfRequestorsToRemoveAccess(value *bool)()
    SetEnableOnBehalfRequestorsToUpdateAccess(value *bool)()
    SetEnableTargetsToSelfAddAccess(value *bool)()
    SetEnableTargetsToSelfRemoveAccess(value *bool)()
    SetEnableTargetsToSelfUpdateAccess(value *bool)()
    SetOnBehalfRequestors(value []SubjectSetable)()
}
