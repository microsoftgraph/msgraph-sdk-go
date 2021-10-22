package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WindowsInformationProtectionPolicy struct {
    WindowsInformationProtection
    daysWithoutContactBeforeUnenroll *int32;
    mdmEnrollmentUrl *string;
    minutesOfInactivityBeforeDeviceLock *int32;
    numberOfPastPinsRemembered *int32;
    passwordMaximumAttemptCount *int32;
    pinExpirationDays *int32;
    pinLowercaseLetters *WindowsInformationProtectionPinCharacterRequirements;
    pinMinimumLength *int32;
    pinSpecialCharacters *WindowsInformationProtectionPinCharacterRequirements;
    pinUppercaseLetters *WindowsInformationProtectionPinCharacterRequirements;
    revokeOnMdmHandoffDisabled *bool;
    windowsHelloForBusinessBlocked *bool;
}
func NewWindowsInformationProtectionPolicy()(*WindowsInformationProtectionPolicy) {
    m := &WindowsInformationProtectionPolicy{
        WindowsInformationProtection: *NewWindowsInformationProtection(),
    }
    return m
}
func (m *WindowsInformationProtectionPolicy) GetDaysWithoutContactBeforeUnenroll()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.daysWithoutContactBeforeUnenroll
    }
}
func (m *WindowsInformationProtectionPolicy) GetMdmEnrollmentUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mdmEnrollmentUrl
    }
}
func (m *WindowsInformationProtectionPolicy) GetMinutesOfInactivityBeforeDeviceLock()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.minutesOfInactivityBeforeDeviceLock
    }
}
func (m *WindowsInformationProtectionPolicy) GetNumberOfPastPinsRemembered()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfPastPinsRemembered
    }
}
func (m *WindowsInformationProtectionPolicy) GetPasswordMaximumAttemptCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passwordMaximumAttemptCount
    }
}
func (m *WindowsInformationProtectionPolicy) GetPinExpirationDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pinExpirationDays
    }
}
func (m *WindowsInformationProtectionPolicy) GetPinLowercaseLetters()(*WindowsInformationProtectionPinCharacterRequirements) {
    if m == nil {
        return nil
    } else {
        return m.pinLowercaseLetters
    }
}
func (m *WindowsInformationProtectionPolicy) GetPinMinimumLength()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pinMinimumLength
    }
}
func (m *WindowsInformationProtectionPolicy) GetPinSpecialCharacters()(*WindowsInformationProtectionPinCharacterRequirements) {
    if m == nil {
        return nil
    } else {
        return m.pinSpecialCharacters
    }
}
func (m *WindowsInformationProtectionPolicy) GetPinUppercaseLetters()(*WindowsInformationProtectionPinCharacterRequirements) {
    if m == nil {
        return nil
    } else {
        return m.pinUppercaseLetters
    }
}
func (m *WindowsInformationProtectionPolicy) GetRevokeOnMdmHandoffDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.revokeOnMdmHandoffDisabled
    }
}
func (m *WindowsInformationProtectionPolicy) GetWindowsHelloForBusinessBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.windowsHelloForBusinessBlocked
    }
}
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
func (m *WindowsInformationProtectionPolicy) SetDaysWithoutContactBeforeUnenroll(value *int32)() {
    m.daysWithoutContactBeforeUnenroll = value
}
func (m *WindowsInformationProtectionPolicy) SetMdmEnrollmentUrl(value *string)() {
    m.mdmEnrollmentUrl = value
}
func (m *WindowsInformationProtectionPolicy) SetMinutesOfInactivityBeforeDeviceLock(value *int32)() {
    m.minutesOfInactivityBeforeDeviceLock = value
}
func (m *WindowsInformationProtectionPolicy) SetNumberOfPastPinsRemembered(value *int32)() {
    m.numberOfPastPinsRemembered = value
}
func (m *WindowsInformationProtectionPolicy) SetPasswordMaximumAttemptCount(value *int32)() {
    m.passwordMaximumAttemptCount = value
}
func (m *WindowsInformationProtectionPolicy) SetPinExpirationDays(value *int32)() {
    m.pinExpirationDays = value
}
func (m *WindowsInformationProtectionPolicy) SetPinLowercaseLetters(value *WindowsInformationProtectionPinCharacterRequirements)() {
    m.pinLowercaseLetters = value
}
func (m *WindowsInformationProtectionPolicy) SetPinMinimumLength(value *int32)() {
    m.pinMinimumLength = value
}
func (m *WindowsInformationProtectionPolicy) SetPinSpecialCharacters(value *WindowsInformationProtectionPinCharacterRequirements)() {
    m.pinSpecialCharacters = value
}
func (m *WindowsInformationProtectionPolicy) SetPinUppercaseLetters(value *WindowsInformationProtectionPinCharacterRequirements)() {
    m.pinUppercaseLetters = value
}
func (m *WindowsInformationProtectionPolicy) SetRevokeOnMdmHandoffDisabled(value *bool)() {
    m.revokeOnMdmHandoffDisabled = value
}
func (m *WindowsInformationProtectionPolicy) SetWindowsHelloForBusinessBlocked(value *bool)() {
    m.windowsHelloForBusinessBlocked = value
}
