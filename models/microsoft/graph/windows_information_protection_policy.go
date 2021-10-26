package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WindowsInformationProtectionPolicy struct {
    WindowsInformationProtection
    // Offline interval before app data is wiped (days)
    daysWithoutContactBeforeUnenroll *int32;
    // Enrollment url for the MDM
    mdmEnrollmentUrl *string;
    // Specifies the maximum amount of time (in minutes) allowed after the device is idle that will cause the device to become PIN or password locked.   Range is an integer X where 0 <= X <= 999.
    minutesOfInactivityBeforeDeviceLock *int32;
    // Integer value that specifies the number of past PINs that can be associated to a user account that can't be reused. The largest number you can configure for this policy setting is 50. The lowest number you can configure for this policy setting is 0. If this policy is set to 0, then storage of previous PINs is not required. This node was added in Windows 10, version 1511. Default is 0.
    numberOfPastPinsRemembered *int32;
    // The number of authentication failures allowed before the device will be wiped. A value of 0 disables device wipe functionality. Range is an integer X where 4 <= X <= 16 for desktop and 0 <= X <= 999 for mobile devices.
    passwordMaximumAttemptCount *int32;
    // Integer value specifies the period of time (in days) that a PIN can be used before the system requires the user to change it. The largest number you can configure for this policy setting is 730. The lowest number you can configure for this policy setting is 0. If this policy is set to 0, then the user's PIN will never expire. This node was added in Windows 10, version 1511. Default is 0.
    pinExpirationDays *int32;
    // Integer value that configures the use of lowercase letters in the Windows Hello for Business PIN. Default is NotAllow. Possible values are: notAllow, requireAtLeastOne, allow.
    pinLowercaseLetters *WindowsInformationProtectionPinCharacterRequirements;
    // Integer value that sets the minimum number of characters required for the PIN. Default value is 4. The lowest number you can configure for this policy setting is 4. The largest number you can configure must be less than the number configured in the Maximum PIN length policy setting or the number 127, whichever is the lowest.
    pinMinimumLength *int32;
    // Integer value that configures the use of special characters in the Windows Hello for Business PIN. Valid special characters for Windows Hello for Business PIN gestures include: ! ' # $ % & ' ( )  + , - . / : ; < = > ? @ [ / ] ^  ` {
    pinSpecialCharacters *WindowsInformationProtectionPinCharacterRequirements;
    // Integer value that configures the use of uppercase letters in the Windows Hello for Business PIN. Default is NotAllow. Possible values are: notAllow, requireAtLeastOne, allow.
    pinUppercaseLetters *WindowsInformationProtectionPinCharacterRequirements;
    // New property in RS2, pending documentation
    revokeOnMdmHandoffDisabled *bool;
    // Boolean value that sets Windows Hello for Business as a method for signing into Windows.
    windowsHelloForBusinessBlocked *bool;
}
// Instantiates a new windowsInformationProtectionPolicy and sets the default values.
func NewWindowsInformationProtectionPolicy()(*WindowsInformationProtectionPolicy) {
    m := &WindowsInformationProtectionPolicy{
        WindowsInformationProtection: *NewWindowsInformationProtection(),
    }
    return m
}
// Gets the daysWithoutContactBeforeUnenroll property value. Offline interval before app data is wiped (days)
func (m *WindowsInformationProtectionPolicy) GetDaysWithoutContactBeforeUnenroll()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.daysWithoutContactBeforeUnenroll
    }
}
// Gets the mdmEnrollmentUrl property value. Enrollment url for the MDM
func (m *WindowsInformationProtectionPolicy) GetMdmEnrollmentUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mdmEnrollmentUrl
    }
}
// Gets the minutesOfInactivityBeforeDeviceLock property value. Specifies the maximum amount of time (in minutes) allowed after the device is idle that will cause the device to become PIN or password locked.   Range is an integer X where 0 <= X <= 999.
func (m *WindowsInformationProtectionPolicy) GetMinutesOfInactivityBeforeDeviceLock()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.minutesOfInactivityBeforeDeviceLock
    }
}
// Gets the numberOfPastPinsRemembered property value. Integer value that specifies the number of past PINs that can be associated to a user account that can't be reused. The largest number you can configure for this policy setting is 50. The lowest number you can configure for this policy setting is 0. If this policy is set to 0, then storage of previous PINs is not required. This node was added in Windows 10, version 1511. Default is 0.
func (m *WindowsInformationProtectionPolicy) GetNumberOfPastPinsRemembered()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfPastPinsRemembered
    }
}
// Gets the passwordMaximumAttemptCount property value. The number of authentication failures allowed before the device will be wiped. A value of 0 disables device wipe functionality. Range is an integer X where 4 <= X <= 16 for desktop and 0 <= X <= 999 for mobile devices.
func (m *WindowsInformationProtectionPolicy) GetPasswordMaximumAttemptCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passwordMaximumAttemptCount
    }
}
// Gets the pinExpirationDays property value. Integer value specifies the period of time (in days) that a PIN can be used before the system requires the user to change it. The largest number you can configure for this policy setting is 730. The lowest number you can configure for this policy setting is 0. If this policy is set to 0, then the user's PIN will never expire. This node was added in Windows 10, version 1511. Default is 0.
func (m *WindowsInformationProtectionPolicy) GetPinExpirationDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pinExpirationDays
    }
}
// Gets the pinLowercaseLetters property value. Integer value that configures the use of lowercase letters in the Windows Hello for Business PIN. Default is NotAllow. Possible values are: notAllow, requireAtLeastOne, allow.
func (m *WindowsInformationProtectionPolicy) GetPinLowercaseLetters()(*WindowsInformationProtectionPinCharacterRequirements) {
    if m == nil {
        return nil
    } else {
        return m.pinLowercaseLetters
    }
}
// Gets the pinMinimumLength property value. Integer value that sets the minimum number of characters required for the PIN. Default value is 4. The lowest number you can configure for this policy setting is 4. The largest number you can configure must be less than the number configured in the Maximum PIN length policy setting or the number 127, whichever is the lowest.
func (m *WindowsInformationProtectionPolicy) GetPinMinimumLength()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pinMinimumLength
    }
}
// Gets the pinSpecialCharacters property value. Integer value that configures the use of special characters in the Windows Hello for Business PIN. Valid special characters for Windows Hello for Business PIN gestures include: ! ' # $ % & ' ( )  + , - . / : ; < = > ? @ [ / ] ^  ` {
func (m *WindowsInformationProtectionPolicy) GetPinSpecialCharacters()(*WindowsInformationProtectionPinCharacterRequirements) {
    if m == nil {
        return nil
    } else {
        return m.pinSpecialCharacters
    }
}
// Gets the pinUppercaseLetters property value. Integer value that configures the use of uppercase letters in the Windows Hello for Business PIN. Default is NotAllow. Possible values are: notAllow, requireAtLeastOne, allow.
func (m *WindowsInformationProtectionPolicy) GetPinUppercaseLetters()(*WindowsInformationProtectionPinCharacterRequirements) {
    if m == nil {
        return nil
    } else {
        return m.pinUppercaseLetters
    }
}
// Gets the revokeOnMdmHandoffDisabled property value. New property in RS2, pending documentation
func (m *WindowsInformationProtectionPolicy) GetRevokeOnMdmHandoffDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.revokeOnMdmHandoffDisabled
    }
}
// Gets the windowsHelloForBusinessBlocked property value. Boolean value that sets Windows Hello for Business as a method for signing into Windows.
func (m *WindowsInformationProtectionPolicy) GetWindowsHelloForBusinessBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.windowsHelloForBusinessBlocked
    }
}
// The deserialization information for the current model
func (m *WindowsInformationProtectionPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.WindowsInformationProtection.GetFieldDeserializers()
    res["daysWithoutContactBeforeUnenroll"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDaysWithoutContactBeforeUnenroll(val)
        return nil
    }
    res["mdmEnrollmentUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMdmEnrollmentUrl(val)
        return nil
    }
    res["minutesOfInactivityBeforeDeviceLock"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetMinutesOfInactivityBeforeDeviceLock(val)
        return nil
    }
    res["numberOfPastPinsRemembered"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNumberOfPastPinsRemembered(val)
        return nil
    }
    res["passwordMaximumAttemptCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetPasswordMaximumAttemptCount(val)
        return nil
    }
    res["pinExpirationDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetPinExpirationDays(val)
        return nil
    }
    res["pinLowercaseLetters"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsInformationProtectionPinCharacterRequirements)
        if err != nil {
            return err
        }
        cast := val.(WindowsInformationProtectionPinCharacterRequirements)
        m.SetPinLowercaseLetters(&cast)
        return nil
    }
    res["pinMinimumLength"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetPinMinimumLength(val)
        return nil
    }
    res["pinSpecialCharacters"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsInformationProtectionPinCharacterRequirements)
        if err != nil {
            return err
        }
        cast := val.(WindowsInformationProtectionPinCharacterRequirements)
        m.SetPinSpecialCharacters(&cast)
        return nil
    }
    res["pinUppercaseLetters"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsInformationProtectionPinCharacterRequirements)
        if err != nil {
            return err
        }
        cast := val.(WindowsInformationProtectionPinCharacterRequirements)
        m.SetPinUppercaseLetters(&cast)
        return nil
    }
    res["revokeOnMdmHandoffDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetRevokeOnMdmHandoffDisabled(val)
        return nil
    }
    res["windowsHelloForBusinessBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetWindowsHelloForBusinessBlocked(val)
        return nil
    }
    return res
}
func (m *WindowsInformationProtectionPolicy) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WindowsInformationProtectionPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.WindowsInformationProtection.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("daysWithoutContactBeforeUnenroll", m.GetDaysWithoutContactBeforeUnenroll())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mdmEnrollmentUrl", m.GetMdmEnrollmentUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("minutesOfInactivityBeforeDeviceLock", m.GetMinutesOfInactivityBeforeDeviceLock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("numberOfPastPinsRemembered", m.GetNumberOfPastPinsRemembered())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMaximumAttemptCount", m.GetPasswordMaximumAttemptCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("pinExpirationDays", m.GetPinExpirationDays())
        if err != nil {
            return err
        }
    }
    if m.GetPinLowercaseLetters() != nil {
        cast := m.GetPinLowercaseLetters().String()
        err = writer.WriteStringValue("pinLowercaseLetters", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("pinMinimumLength", m.GetPinMinimumLength())
        if err != nil {
            return err
        }
    }
    if m.GetPinSpecialCharacters() != nil {
        cast := m.GetPinSpecialCharacters().String()
        err = writer.WriteStringValue("pinSpecialCharacters", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPinUppercaseLetters() != nil {
        cast := m.GetPinUppercaseLetters().String()
        err = writer.WriteStringValue("pinUppercaseLetters", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("revokeOnMdmHandoffDisabled", m.GetRevokeOnMdmHandoffDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("windowsHelloForBusinessBlocked", m.GetWindowsHelloForBusinessBlocked())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the daysWithoutContactBeforeUnenroll property value. Offline interval before app data is wiped (days)
// Parameters:
//  - value : Value to set for the daysWithoutContactBeforeUnenroll property.
func (m *WindowsInformationProtectionPolicy) SetDaysWithoutContactBeforeUnenroll(value *int32)() {
    m.daysWithoutContactBeforeUnenroll = value
}
// Sets the mdmEnrollmentUrl property value. Enrollment url for the MDM
// Parameters:
//  - value : Value to set for the mdmEnrollmentUrl property.
func (m *WindowsInformationProtectionPolicy) SetMdmEnrollmentUrl(value *string)() {
    m.mdmEnrollmentUrl = value
}
// Sets the minutesOfInactivityBeforeDeviceLock property value. Specifies the maximum amount of time (in minutes) allowed after the device is idle that will cause the device to become PIN or password locked.   Range is an integer X where 0 <= X <= 999.
// Parameters:
//  - value : Value to set for the minutesOfInactivityBeforeDeviceLock property.
func (m *WindowsInformationProtectionPolicy) SetMinutesOfInactivityBeforeDeviceLock(value *int32)() {
    m.minutesOfInactivityBeforeDeviceLock = value
}
// Sets the numberOfPastPinsRemembered property value. Integer value that specifies the number of past PINs that can be associated to a user account that can't be reused. The largest number you can configure for this policy setting is 50. The lowest number you can configure for this policy setting is 0. If this policy is set to 0, then storage of previous PINs is not required. This node was added in Windows 10, version 1511. Default is 0.
// Parameters:
//  - value : Value to set for the numberOfPastPinsRemembered property.
func (m *WindowsInformationProtectionPolicy) SetNumberOfPastPinsRemembered(value *int32)() {
    m.numberOfPastPinsRemembered = value
}
// Sets the passwordMaximumAttemptCount property value. The number of authentication failures allowed before the device will be wiped. A value of 0 disables device wipe functionality. Range is an integer X where 4 <= X <= 16 for desktop and 0 <= X <= 999 for mobile devices.
// Parameters:
//  - value : Value to set for the passwordMaximumAttemptCount property.
func (m *WindowsInformationProtectionPolicy) SetPasswordMaximumAttemptCount(value *int32)() {
    m.passwordMaximumAttemptCount = value
}
// Sets the pinExpirationDays property value. Integer value specifies the period of time (in days) that a PIN can be used before the system requires the user to change it. The largest number you can configure for this policy setting is 730. The lowest number you can configure for this policy setting is 0. If this policy is set to 0, then the user's PIN will never expire. This node was added in Windows 10, version 1511. Default is 0.
// Parameters:
//  - value : Value to set for the pinExpirationDays property.
func (m *WindowsInformationProtectionPolicy) SetPinExpirationDays(value *int32)() {
    m.pinExpirationDays = value
}
// Sets the pinLowercaseLetters property value. Integer value that configures the use of lowercase letters in the Windows Hello for Business PIN. Default is NotAllow. Possible values are: notAllow, requireAtLeastOne, allow.
// Parameters:
//  - value : Value to set for the pinLowercaseLetters property.
func (m *WindowsInformationProtectionPolicy) SetPinLowercaseLetters(value *WindowsInformationProtectionPinCharacterRequirements)() {
    m.pinLowercaseLetters = value
}
// Sets the pinMinimumLength property value. Integer value that sets the minimum number of characters required for the PIN. Default value is 4. The lowest number you can configure for this policy setting is 4. The largest number you can configure must be less than the number configured in the Maximum PIN length policy setting or the number 127, whichever is the lowest.
// Parameters:
//  - value : Value to set for the pinMinimumLength property.
func (m *WindowsInformationProtectionPolicy) SetPinMinimumLength(value *int32)() {
    m.pinMinimumLength = value
}
// Sets the pinSpecialCharacters property value. Integer value that configures the use of special characters in the Windows Hello for Business PIN. Valid special characters for Windows Hello for Business PIN gestures include: ! ' # $ % & ' ( )  + , - . / : ; < = > ? @ [ / ] ^  ` {
// Parameters:
//  - value : Value to set for the pinSpecialCharacters property.
func (m *WindowsInformationProtectionPolicy) SetPinSpecialCharacters(value *WindowsInformationProtectionPinCharacterRequirements)() {
    m.pinSpecialCharacters = value
}
// Sets the pinUppercaseLetters property value. Integer value that configures the use of uppercase letters in the Windows Hello for Business PIN. Default is NotAllow. Possible values are: notAllow, requireAtLeastOne, allow.
// Parameters:
//  - value : Value to set for the pinUppercaseLetters property.
func (m *WindowsInformationProtectionPolicy) SetPinUppercaseLetters(value *WindowsInformationProtectionPinCharacterRequirements)() {
    m.pinUppercaseLetters = value
}
// Sets the revokeOnMdmHandoffDisabled property value. New property in RS2, pending documentation
// Parameters:
//  - value : Value to set for the revokeOnMdmHandoffDisabled property.
func (m *WindowsInformationProtectionPolicy) SetRevokeOnMdmHandoffDisabled(value *bool)() {
    m.revokeOnMdmHandoffDisabled = value
}
// Sets the windowsHelloForBusinessBlocked property value. Boolean value that sets Windows Hello for Business as a method for signing into Windows.
// Parameters:
//  - value : Value to set for the windowsHelloForBusinessBlocked property.
func (m *WindowsInformationProtectionPolicy) SetWindowsHelloForBusinessBlocked(value *bool)() {
    m.windowsHelloForBusinessBlocked = value
}
