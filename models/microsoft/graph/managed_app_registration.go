package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ManagedAppRegistration struct {
    Entity
    // The app package Identifier
    appIdentifier *MobileAppIdentifier;
    // App version
    applicationVersion *string;
    // Zero or more policys already applied on the registered app when it last synchronized with managment service.
    appliedPolicies []ManagedAppPolicy;
    // Date and time of creation
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Host device name
    deviceName *string;
    // App management SDK generated tag, which helps relate apps hosted on the same device. Not guaranteed to relate apps in all conditions.
    deviceTag *string;
    // Host device type
    deviceType *string;
    // Zero or more reasons an app registration is flagged. E.g. app running on rooted device
    flaggedReasons []ManagedAppFlaggedReason;
    // Zero or more policies admin intended for the app as of now.
    intendedPolicies []ManagedAppPolicy;
    // Date and time of last the app synced with management service.
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // App management SDK version
    managementSdkVersion *string;
    // Zero or more long running operations triggered on the app registration.
    operations []ManagedAppOperation;
    // Operating System version
    platformVersion *string;
    // The user Id to who this app registration belongs.
    userId *string;
    // Version of the entity.
    version *string;
}
// Instantiates a new managedAppRegistration and sets the default values.
func NewManagedAppRegistration()(*ManagedAppRegistration) {
    m := &ManagedAppRegistration{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the appIdentifier property value. The app package Identifier
func (m *ManagedAppRegistration) GetAppIdentifier()(*MobileAppIdentifier) {
    if m == nil {
        return nil
    } else {
        return m.appIdentifier
    }
}
// Gets the applicationVersion property value. App version
func (m *ManagedAppRegistration) GetApplicationVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationVersion
    }
}
// Gets the appliedPolicies property value. Zero or more policys already applied on the registered app when it last synchronized with managment service.
func (m *ManagedAppRegistration) GetAppliedPolicies()([]ManagedAppPolicy) {
    if m == nil {
        return nil
    } else {
        return m.appliedPolicies
    }
}
// Gets the createdDateTime property value. Date and time of creation
func (m *ManagedAppRegistration) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the deviceName property value. Host device name
func (m *ManagedAppRegistration) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// Gets the deviceTag property value. App management SDK generated tag, which helps relate apps hosted on the same device. Not guaranteed to relate apps in all conditions.
func (m *ManagedAppRegistration) GetDeviceTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceTag
    }
}
// Gets the deviceType property value. Host device type
func (m *ManagedAppRegistration) GetDeviceType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceType
    }
}
// Gets the flaggedReasons property value. Zero or more reasons an app registration is flagged. E.g. app running on rooted device
func (m *ManagedAppRegistration) GetFlaggedReasons()([]ManagedAppFlaggedReason) {
    if m == nil {
        return nil
    } else {
        return m.flaggedReasons
    }
}
// Gets the intendedPolicies property value. Zero or more policies admin intended for the app as of now.
func (m *ManagedAppRegistration) GetIntendedPolicies()([]ManagedAppPolicy) {
    if m == nil {
        return nil
    } else {
        return m.intendedPolicies
    }
}
// Gets the lastSyncDateTime property value. Date and time of last the app synced with management service.
func (m *ManagedAppRegistration) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncDateTime
    }
}
// Gets the managementSdkVersion property value. App management SDK version
func (m *ManagedAppRegistration) GetManagementSdkVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managementSdkVersion
    }
}
// Gets the operations property value. Zero or more long running operations triggered on the app registration.
func (m *ManagedAppRegistration) GetOperations()([]ManagedAppOperation) {
    if m == nil {
        return nil
    } else {
        return m.operations
    }
}
// Gets the platformVersion property value. Operating System version
func (m *ManagedAppRegistration) GetPlatformVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.platformVersion
    }
}
// Gets the userId property value. The user Id to who this app registration belongs.
func (m *ManagedAppRegistration) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// Gets the version property value. Version of the entity.
func (m *ManagedAppRegistration) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// The deserialization information for the current model
func (m *ManagedAppRegistration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobileAppIdentifier() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppIdentifier(val.(*MobileAppIdentifier))
        }
        return nil
    }
    res["applicationVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationVersion(val)
        }
        return nil
    }
    res["appliedPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedAppPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAppPolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagedAppPolicy))
            }
            m.SetAppliedPolicies(res)
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["deviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["deviceTag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceTag(val)
        }
        return nil
    }
    res["deviceType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceType(val)
        }
        return nil
    }
    res["flaggedReasons"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseManagedAppFlaggedReason)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAppFlaggedReason, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagedAppFlaggedReason))
            }
            m.SetFlaggedReasons(res)
        }
        return nil
    }
    res["intendedPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedAppPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAppPolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagedAppPolicy))
            }
            m.SetIntendedPolicies(res)
        }
        return nil
    }
    res["lastSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncDateTime(val)
        }
        return nil
    }
    res["managementSdkVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementSdkVersion(val)
        }
        return nil
    }
    res["operations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedAppOperation() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAppOperation, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagedAppOperation))
            }
            m.SetOperations(res)
        }
        return nil
    }
    res["platformVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatformVersion(val)
        }
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
func (m *ManagedAppRegistration) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ManagedAppRegistration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("appIdentifier", m.GetAppIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("applicationVersion", m.GetApplicationVersion())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAppliedPolicies()))
        for i, v := range m.GetAppliedPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("appliedPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceTag", m.GetDeviceTag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceType", m.GetDeviceType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("flaggedReasons", SerializeManagedAppFlaggedReason(m.GetFlaggedReasons()))
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIntendedPolicies()))
        for i, v := range m.GetIntendedPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("intendedPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSyncDateTime", m.GetLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managementSdkVersion", m.GetManagementSdkVersion())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("operations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("platformVersion", m.GetPlatformVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the appIdentifier property value. The app package Identifier
// Parameters:
//  - value : Value to set for the appIdentifier property.
func (m *ManagedAppRegistration) SetAppIdentifier(value *MobileAppIdentifier)() {
    m.appIdentifier = value
}
// Sets the applicationVersion property value. App version
// Parameters:
//  - value : Value to set for the applicationVersion property.
func (m *ManagedAppRegistration) SetApplicationVersion(value *string)() {
    m.applicationVersion = value
}
// Sets the appliedPolicies property value. Zero or more policys already applied on the registered app when it last synchronized with managment service.
// Parameters:
//  - value : Value to set for the appliedPolicies property.
func (m *ManagedAppRegistration) SetAppliedPolicies(value []ManagedAppPolicy)() {
    m.appliedPolicies = value
}
// Sets the createdDateTime property value. Date and time of creation
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *ManagedAppRegistration) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the deviceName property value. Host device name
// Parameters:
//  - value : Value to set for the deviceName property.
func (m *ManagedAppRegistration) SetDeviceName(value *string)() {
    m.deviceName = value
}
// Sets the deviceTag property value. App management SDK generated tag, which helps relate apps hosted on the same device. Not guaranteed to relate apps in all conditions.
// Parameters:
//  - value : Value to set for the deviceTag property.
func (m *ManagedAppRegistration) SetDeviceTag(value *string)() {
    m.deviceTag = value
}
// Sets the deviceType property value. Host device type
// Parameters:
//  - value : Value to set for the deviceType property.
func (m *ManagedAppRegistration) SetDeviceType(value *string)() {
    m.deviceType = value
}
// Sets the flaggedReasons property value. Zero or more reasons an app registration is flagged. E.g. app running on rooted device
// Parameters:
//  - value : Value to set for the flaggedReasons property.
func (m *ManagedAppRegistration) SetFlaggedReasons(value []ManagedAppFlaggedReason)() {
    m.flaggedReasons = value
}
// Sets the intendedPolicies property value. Zero or more policies admin intended for the app as of now.
// Parameters:
//  - value : Value to set for the intendedPolicies property.
func (m *ManagedAppRegistration) SetIntendedPolicies(value []ManagedAppPolicy)() {
    m.intendedPolicies = value
}
// Sets the lastSyncDateTime property value. Date and time of last the app synced with management service.
// Parameters:
//  - value : Value to set for the lastSyncDateTime property.
func (m *ManagedAppRegistration) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncDateTime = value
}
// Sets the managementSdkVersion property value. App management SDK version
// Parameters:
//  - value : Value to set for the managementSdkVersion property.
func (m *ManagedAppRegistration) SetManagementSdkVersion(value *string)() {
    m.managementSdkVersion = value
}
// Sets the operations property value. Zero or more long running operations triggered on the app registration.
// Parameters:
//  - value : Value to set for the operations property.
func (m *ManagedAppRegistration) SetOperations(value []ManagedAppOperation)() {
    m.operations = value
}
// Sets the platformVersion property value. Operating System version
// Parameters:
//  - value : Value to set for the platformVersion property.
func (m *ManagedAppRegistration) SetPlatformVersion(value *string)() {
    m.platformVersion = value
}
// Sets the userId property value. The user Id to who this app registration belongs.
// Parameters:
//  - value : Value to set for the userId property.
func (m *ManagedAppRegistration) SetUserId(value *string)() {
    m.userId = value
}
// Sets the version property value. Version of the entity.
// Parameters:
//  - value : Value to set for the version property.
func (m *ManagedAppRegistration) SetVersion(value *string)() {
    m.version = value
}
