package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagedAppRegistration 
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
// NewManagedAppRegistration instantiates a new managedAppRegistration and sets the default values.
func NewManagedAppRegistration()(*ManagedAppRegistration) {
    m := &ManagedAppRegistration{
        Entity: *NewEntity(),
    }
    return m
}
// GetAppIdentifier gets the appIdentifier property value. The app package Identifier
func (m *ManagedAppRegistration) GetAppIdentifier()(*MobileAppIdentifier) {
    if m == nil {
        return nil
    } else {
        return m.appIdentifier
    }
}
// GetApplicationVersion gets the applicationVersion property value. App version
func (m *ManagedAppRegistration) GetApplicationVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationVersion
    }
}
// GetAppliedPolicies gets the appliedPolicies property value. Zero or more policys already applied on the registered app when it last synchronized with managment service.
func (m *ManagedAppRegistration) GetAppliedPolicies()([]ManagedAppPolicy) {
    if m == nil {
        return nil
    } else {
        return m.appliedPolicies
    }
}
// GetCreatedDateTime gets the createdDateTime property value. Date and time of creation
func (m *ManagedAppRegistration) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDeviceName gets the deviceName property value. Host device name
func (m *ManagedAppRegistration) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// GetDeviceTag gets the deviceTag property value. App management SDK generated tag, which helps relate apps hosted on the same device. Not guaranteed to relate apps in all conditions.
func (m *ManagedAppRegistration) GetDeviceTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceTag
    }
}
// GetDeviceType gets the deviceType property value. Host device type
func (m *ManagedAppRegistration) GetDeviceType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceType
    }
}
// GetFlaggedReasons gets the flaggedReasons property value. Zero or more reasons an app registration is flagged. E.g. app running on rooted device
func (m *ManagedAppRegistration) GetFlaggedReasons()([]ManagedAppFlaggedReason) {
    if m == nil {
        return nil
    } else {
        return m.flaggedReasons
    }
}
// GetIntendedPolicies gets the intendedPolicies property value. Zero or more policies admin intended for the app as of now.
func (m *ManagedAppRegistration) GetIntendedPolicies()([]ManagedAppPolicy) {
    if m == nil {
        return nil
    } else {
        return m.intendedPolicies
    }
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. Date and time of last the app synced with management service.
func (m *ManagedAppRegistration) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncDateTime
    }
}
// GetManagementSdkVersion gets the managementSdkVersion property value. App management SDK version
func (m *ManagedAppRegistration) GetManagementSdkVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managementSdkVersion
    }
}
// GetOperations gets the operations property value. Zero or more long running operations triggered on the app registration.
func (m *ManagedAppRegistration) GetOperations()([]ManagedAppOperation) {
    if m == nil {
        return nil
    } else {
        return m.operations
    }
}
// GetPlatformVersion gets the platformVersion property value. Operating System version
func (m *ManagedAppRegistration) GetPlatformVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.platformVersion
    }
}
// GetUserId gets the userId property value. The user Id to who this app registration belongs.
func (m *ManagedAppRegistration) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// GetVersion gets the version property value. Version of the entity.
func (m *ManagedAppRegistration) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetAppIdentifier sets the appIdentifier property value. The app package Identifier
func (m *ManagedAppRegistration) SetAppIdentifier(value *MobileAppIdentifier)() {
    if m != nil {
        m.appIdentifier = value
    }
}
// SetApplicationVersion sets the applicationVersion property value. App version
func (m *ManagedAppRegistration) SetApplicationVersion(value *string)() {
    if m != nil {
        m.applicationVersion = value
    }
}
// SetAppliedPolicies sets the appliedPolicies property value. Zero or more policys already applied on the registered app when it last synchronized with managment service.
func (m *ManagedAppRegistration) SetAppliedPolicies(value []ManagedAppPolicy)() {
    if m != nil {
        m.appliedPolicies = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Date and time of creation
func (m *ManagedAppRegistration) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDeviceName sets the deviceName property value. Host device name
func (m *ManagedAppRegistration) SetDeviceName(value *string)() {
    if m != nil {
        m.deviceName = value
    }
}
// SetDeviceTag sets the deviceTag property value. App management SDK generated tag, which helps relate apps hosted on the same device. Not guaranteed to relate apps in all conditions.
func (m *ManagedAppRegistration) SetDeviceTag(value *string)() {
    if m != nil {
        m.deviceTag = value
    }
}
// SetDeviceType sets the deviceType property value. Host device type
func (m *ManagedAppRegistration) SetDeviceType(value *string)() {
    if m != nil {
        m.deviceType = value
    }
}
// SetFlaggedReasons sets the flaggedReasons property value. Zero or more reasons an app registration is flagged. E.g. app running on rooted device
func (m *ManagedAppRegistration) SetFlaggedReasons(value []ManagedAppFlaggedReason)() {
    if m != nil {
        m.flaggedReasons = value
    }
}
// SetIntendedPolicies sets the intendedPolicies property value. Zero or more policies admin intended for the app as of now.
func (m *ManagedAppRegistration) SetIntendedPolicies(value []ManagedAppPolicy)() {
    if m != nil {
        m.intendedPolicies = value
    }
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. Date and time of last the app synced with management service.
func (m *ManagedAppRegistration) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastSyncDateTime = value
    }
}
// SetManagementSdkVersion sets the managementSdkVersion property value. App management SDK version
func (m *ManagedAppRegistration) SetManagementSdkVersion(value *string)() {
    if m != nil {
        m.managementSdkVersion = value
    }
}
// SetOperations sets the operations property value. Zero or more long running operations triggered on the app registration.
func (m *ManagedAppRegistration) SetOperations(value []ManagedAppOperation)() {
    if m != nil {
        m.operations = value
    }
}
// SetPlatformVersion sets the platformVersion property value. Operating System version
func (m *ManagedAppRegistration) SetPlatformVersion(value *string)() {
    if m != nil {
        m.platformVersion = value
    }
}
// SetUserId sets the userId property value. The user Id to who this app registration belongs.
func (m *ManagedAppRegistration) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
// SetVersion sets the version property value. Version of the entity.
func (m *ManagedAppRegistration) SetVersion(value *string)() {
    if m != nil {
        m.version = value
    }
}
