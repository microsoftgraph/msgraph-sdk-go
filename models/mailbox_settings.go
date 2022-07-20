package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MailboxSettings 
type MailboxSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Folder ID of an archive folder for the user. Read only.
    archiveFolder *string
    // Configuration settings to automatically notify the sender of an incoming email with a message from the signed-in user.
    automaticRepliesSetting AutomaticRepliesSettingable
    // The date format for the user's mailbox.
    dateFormat *string
    // If the user has a calendar delegate, this specifies whether the delegate, mailbox owner, or both receive meeting messages and meeting responses. Possible values are: sendToDelegateAndInformationToPrincipal, sendToDelegateAndPrincipal, sendToDelegateOnly. The default is sendToDelegateOnly.
    delegateMeetingMessageDeliveryOptions *DelegateMeetingMessageDeliveryOptions
    // The locale information for the user, including the preferred language and country/region.
    language LocaleInfoable
    // The OdataType property
    odataType *string
    // The time format for the user's mailbox.
    timeFormat *string
    // The default time zone for the user's mailbox.
    timeZone *string
    // The purpose of the mailbox. Used to differentiate a mailbox for a single user from a shared mailbox and equipment mailbox in Exchange Online. Read only.
    userPurpose *UserPurpose
    // The days of the week and hours in a specific time zone that the user works.
    workingHours WorkingHoursable
}
// NewMailboxSettings instantiates a new mailboxSettings and sets the default values.
func NewMailboxSettings()(*MailboxSettings) {
    m := &MailboxSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.mailboxSettings";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMailboxSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMailboxSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMailboxSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MailboxSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetArchiveFolder gets the archiveFolder property value. Folder ID of an archive folder for the user. Read only.
func (m *MailboxSettings) GetArchiveFolder()(*string) {
    if m == nil {
        return nil
    } else {
        return m.archiveFolder
    }
}
// GetAutomaticRepliesSetting gets the automaticRepliesSetting property value. Configuration settings to automatically notify the sender of an incoming email with a message from the signed-in user.
func (m *MailboxSettings) GetAutomaticRepliesSetting()(AutomaticRepliesSettingable) {
    if m == nil {
        return nil
    } else {
        return m.automaticRepliesSetting
    }
}
// GetDateFormat gets the dateFormat property value. The date format for the user's mailbox.
func (m *MailboxSettings) GetDateFormat()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dateFormat
    }
}
// GetDelegateMeetingMessageDeliveryOptions gets the delegateMeetingMessageDeliveryOptions property value. If the user has a calendar delegate, this specifies whether the delegate, mailbox owner, or both receive meeting messages and meeting responses. Possible values are: sendToDelegateAndInformationToPrincipal, sendToDelegateAndPrincipal, sendToDelegateOnly. The default is sendToDelegateOnly.
func (m *MailboxSettings) GetDelegateMeetingMessageDeliveryOptions()(*DelegateMeetingMessageDeliveryOptions) {
    if m == nil {
        return nil
    } else {
        return m.delegateMeetingMessageDeliveryOptions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MailboxSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["archiveFolder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArchiveFolder(val)
        }
        return nil
    }
    res["automaticRepliesSetting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAutomaticRepliesSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutomaticRepliesSetting(val.(AutomaticRepliesSettingable))
        }
        return nil
    }
    res["dateFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateFormat(val)
        }
        return nil
    }
    res["delegateMeetingMessageDeliveryOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDelegateMeetingMessageDeliveryOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDelegateMeetingMessageDeliveryOptions(val.(*DelegateMeetingMessageDeliveryOptions))
        }
        return nil
    }
    res["language"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLocaleInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguage(val.(LocaleInfoable))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["timeFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeFormat(val)
        }
        return nil
    }
    res["timeZone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeZone(val)
        }
        return nil
    }
    res["userPurpose"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserPurpose)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPurpose(val.(*UserPurpose))
        }
        return nil
    }
    res["workingHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkingHoursFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkingHours(val.(WorkingHoursable))
        }
        return nil
    }
    return res
}
// GetLanguage gets the language property value. The locale information for the user, including the preferred language and country/region.
func (m *MailboxSettings) GetLanguage()(LocaleInfoable) {
    if m == nil {
        return nil
    } else {
        return m.language
    }
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MailboxSettings) GetOdataType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.odataType
    }
}
// GetTimeFormat gets the timeFormat property value. The time format for the user's mailbox.
func (m *MailboxSettings) GetTimeFormat()(*string) {
    if m == nil {
        return nil
    } else {
        return m.timeFormat
    }
}
// GetTimeZone gets the timeZone property value. The default time zone for the user's mailbox.
func (m *MailboxSettings) GetTimeZone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.timeZone
    }
}
// GetUserPurpose gets the userPurpose property value. The purpose of the mailbox. Used to differentiate a mailbox for a single user from a shared mailbox and equipment mailbox in Exchange Online. Read only.
func (m *MailboxSettings) GetUserPurpose()(*UserPurpose) {
    if m == nil {
        return nil
    } else {
        return m.userPurpose
    }
}
// GetWorkingHours gets the workingHours property value. The days of the week and hours in a specific time zone that the user works.
func (m *MailboxSettings) GetWorkingHours()(WorkingHoursable) {
    if m == nil {
        return nil
    } else {
        return m.workingHours
    }
}
// Serialize serializes information the current object
func (m *MailboxSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("archiveFolder", m.GetArchiveFolder())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("automaticRepliesSetting", m.GetAutomaticRepliesSetting())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("dateFormat", m.GetDateFormat())
        if err != nil {
            return err
        }
    }
    if m.GetDelegateMeetingMessageDeliveryOptions() != nil {
        cast := (*m.GetDelegateMeetingMessageDeliveryOptions()).String()
        err := writer.WriteStringValue("delegateMeetingMessageDeliveryOptions", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("language", m.GetLanguage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("timeFormat", m.GetTimeFormat())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("timeZone", m.GetTimeZone())
        if err != nil {
            return err
        }
    }
    if m.GetUserPurpose() != nil {
        cast := (*m.GetUserPurpose()).String()
        err := writer.WriteStringValue("userPurpose", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("workingHours", m.GetWorkingHours())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MailboxSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetArchiveFolder sets the archiveFolder property value. Folder ID of an archive folder for the user. Read only.
func (m *MailboxSettings) SetArchiveFolder(value *string)() {
    if m != nil {
        m.archiveFolder = value
    }
}
// SetAutomaticRepliesSetting sets the automaticRepliesSetting property value. Configuration settings to automatically notify the sender of an incoming email with a message from the signed-in user.
func (m *MailboxSettings) SetAutomaticRepliesSetting(value AutomaticRepliesSettingable)() {
    if m != nil {
        m.automaticRepliesSetting = value
    }
}
// SetDateFormat sets the dateFormat property value. The date format for the user's mailbox.
func (m *MailboxSettings) SetDateFormat(value *string)() {
    if m != nil {
        m.dateFormat = value
    }
}
// SetDelegateMeetingMessageDeliveryOptions sets the delegateMeetingMessageDeliveryOptions property value. If the user has a calendar delegate, this specifies whether the delegate, mailbox owner, or both receive meeting messages and meeting responses. Possible values are: sendToDelegateAndInformationToPrincipal, sendToDelegateAndPrincipal, sendToDelegateOnly. The default is sendToDelegateOnly.
func (m *MailboxSettings) SetDelegateMeetingMessageDeliveryOptions(value *DelegateMeetingMessageDeliveryOptions)() {
    if m != nil {
        m.delegateMeetingMessageDeliveryOptions = value
    }
}
// SetLanguage sets the language property value. The locale information for the user, including the preferred language and country/region.
func (m *MailboxSettings) SetLanguage(value LocaleInfoable)() {
    if m != nil {
        m.language = value
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MailboxSettings) SetOdataType(value *string)() {
    if m != nil {
        m.odataType = value
    }
}
// SetTimeFormat sets the timeFormat property value. The time format for the user's mailbox.
func (m *MailboxSettings) SetTimeFormat(value *string)() {
    if m != nil {
        m.timeFormat = value
    }
}
// SetTimeZone sets the timeZone property value. The default time zone for the user's mailbox.
func (m *MailboxSettings) SetTimeZone(value *string)() {
    if m != nil {
        m.timeZone = value
    }
}
// SetUserPurpose sets the userPurpose property value. The purpose of the mailbox. Used to differentiate a mailbox for a single user from a shared mailbox and equipment mailbox in Exchange Online. Read only.
func (m *MailboxSettings) SetUserPurpose(value *UserPurpose)() {
    if m != nil {
        m.userPurpose = value
    }
}
// SetWorkingHours sets the workingHours property value. The days of the week and hours in a specific time zone that the user works.
func (m *MailboxSettings) SetWorkingHours(value WorkingHoursable)() {
    if m != nil {
        m.workingHours = value
    }
}
