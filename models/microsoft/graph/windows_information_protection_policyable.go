package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsInformationProtectionPolicyable 
type WindowsInformationProtectionPolicyable interface {
    WindowsInformationProtectionable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDaysWithoutContactBeforeUnenroll()(*int32)
    GetMdmEnrollmentUrl()(*string)
    GetMinutesOfInactivityBeforeDeviceLock()(*int32)
    GetNumberOfPastPinsRemembered()(*int32)
    GetPasswordMaximumAttemptCount()(*int32)
    GetPinExpirationDays()(*int32)
    GetPinLowercaseLetters()(*WindowsInformationProtectionPinCharacterRequirements)
    GetPinMinimumLength()(*int32)
    GetPinSpecialCharacters()(*WindowsInformationProtectionPinCharacterRequirements)
    GetPinUppercaseLetters()(*WindowsInformationProtectionPinCharacterRequirements)
    GetRevokeOnMdmHandoffDisabled()(*bool)
    GetWindowsHelloForBusinessBlocked()(*bool)
    SetDaysWithoutContactBeforeUnenroll(value *int32)()
    SetMdmEnrollmentUrl(value *string)()
    SetMinutesOfInactivityBeforeDeviceLock(value *int32)()
    SetNumberOfPastPinsRemembered(value *int32)()
    SetPasswordMaximumAttemptCount(value *int32)()
    SetPinExpirationDays(value *int32)()
    SetPinLowercaseLetters(value *WindowsInformationProtectionPinCharacterRequirements)()
    SetPinMinimumLength(value *int32)()
    SetPinSpecialCharacters(value *WindowsInformationProtectionPinCharacterRequirements)()
    SetPinUppercaseLetters(value *WindowsInformationProtectionPinCharacterRequirements)()
    SetRevokeOnMdmHandoffDisabled(value *bool)()
    SetWindowsHelloForBusinessBlocked(value *bool)()
}
