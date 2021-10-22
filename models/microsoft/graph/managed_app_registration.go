package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ManagedAppRegistration struct {
    Entity
    appIdentifier *MobileAppIdentifier;
    applicationVersion *string;
    appliedPolicies []ManagedAppPolicy;
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    deviceName *string;
    deviceTag *string;
    deviceType *string;
    flaggedReasons []ManagedAppFlaggedReason;
    intendedPolicies []ManagedAppPolicy;
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    managementSdkVersion *string;
    operations []ManagedAppOperation;
    platformVersion *string;
    userId *string;
    version *string;
}
func NewManagedAppRegistration()(*ManagedAppRegistration) {
    m := &ManagedAppRegistration{
        Entity: *NewEntity(),
    }
    return m
}
func (m *ManagedAppRegistration) GetAppIdentifier()(*MobileAppIdentifier) {
    if m == nil {
        return nil
    } else {
        return m.appIdentifier
    }
}
func (m *ManagedAppRegistration) GetApplicationVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationVersion
    }
}
func (m *ManagedAppRegistration) GetAppliedPolicies()([]ManagedAppPolicy) {
    if m == nil {
        return nil
    } else {
        return m.appliedPolicies
    }
}
func (m *ManagedAppRegistration) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
func (m *ManagedAppRegistration) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
func (m *ManagedAppRegistration) GetDeviceTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceTag
    }
}
func (m *ManagedAppRegistration) GetDeviceType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceType
    }
}
func (m *ManagedAppRegistration) GetFlaggedReasons()([]ManagedAppFlaggedReason) {
    if m == nil {
        return nil
    } else {
        return m.flaggedReasons
    }
}
func (m *ManagedAppRegistration) GetIntendedPolicies()([]ManagedAppPolicy) {
    if m == nil {
        return nil
    } else {
        return m.intendedPolicies
    }
}
func (m *ManagedAppRegistration) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncDateTime
    }
}
func (m *ManagedAppRegistration) GetManagementSdkVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managementSdkVersion
    }
}
func (m *ManagedAppRegistration) GetOperations()([]ManagedAppOperation) {
    if m == nil {
        return nil
    } else {
        return m.operations
    }
}
func (m *ManagedAppRegistration) GetPlatformVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.platformVersion
    }
}
func (m *ManagedAppRegistration) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
func (m *ManagedAppRegistration) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
func (m *ManagedAppRegistration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobileAppIdentifier() })
        if err != nil {
            return err
        }
        m.SetAppIdentifier(val.(*MobileAppIdentifier))
        return nil
    }
    res["applicationVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetApplicationVersion(val)
        return nil
    }
    res["appliedPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedAppPolicy() })
        if err != nil {
            return err
        }
        res := make([]ManagedAppPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagedAppPolicy))
        }
        m.SetAppliedPolicies(res)
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["deviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceName(val)
        return nil
    }
    res["deviceTag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceTag(val)
        return nil
    }
    res["deviceType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceType(val)
        return nil
    }
    res["flaggedReasons"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseManagedAppFlaggedReason)
        if err != nil {
            return err
        }
        res := make([]ManagedAppFlaggedReason, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagedAppFlaggedReason))
        }
        m.SetFlaggedReasons(res)
        return nil
    }
    res["intendedPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedAppPolicy() })
        if err != nil {
            return err
        }
        res := make([]ManagedAppPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagedAppPolicy))
        }
        m.SetIntendedPolicies(res)
        return nil
    }
    res["lastSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastSyncDateTime(val)
        return nil
    }
    res["managementSdkVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManagementSdkVersion(val)
        return nil
    }
    res["operations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedAppOperation() })
        if err != nil {
            return err
        }
        res := make([]ManagedAppOperation, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagedAppOperation))
        }
        m.SetOperations(res)
        return nil
    }
    res["platformVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPlatformVersion(val)
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserId(val)
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetVersion(val)
        return nil
    }
    return res
}
func (m *ManagedAppRegistration) IsNil()(bool) {
    return m == nil
}
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
func (m *ManagedAppRegistration) SetAppIdentifier(value *MobileAppIdentifier)() {
    m.appIdentifier = value
}
func (m *ManagedAppRegistration) SetApplicationVersion(value *string)() {
    m.applicationVersion = value
}
func (m *ManagedAppRegistration) SetAppliedPolicies(value []ManagedAppPolicy)() {
    m.appliedPolicies = value
}
func (m *ManagedAppRegistration) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
func (m *ManagedAppRegistration) SetDeviceName(value *string)() {
    m.deviceName = value
}
func (m *ManagedAppRegistration) SetDeviceTag(value *string)() {
    m.deviceTag = value
}
func (m *ManagedAppRegistration) SetDeviceType(value *string)() {
    m.deviceType = value
}
func (m *ManagedAppRegistration) SetFlaggedReasons(value []ManagedAppFlaggedReason)() {
    m.flaggedReasons = value
}
func (m *ManagedAppRegistration) SetIntendedPolicies(value []ManagedAppPolicy)() {
    m.intendedPolicies = value
}
func (m *ManagedAppRegistration) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncDateTime = value
}
func (m *ManagedAppRegistration) SetManagementSdkVersion(value *string)() {
    m.managementSdkVersion = value
}
func (m *ManagedAppRegistration) SetOperations(value []ManagedAppOperation)() {
    m.operations = value
}
func (m *ManagedAppRegistration) SetPlatformVersion(value *string)() {
    m.platformVersion = value
}
func (m *ManagedAppRegistration) SetUserId(value *string)() {
    m.userId = value
}
func (m *ManagedAppRegistration) SetVersion(value *string)() {
    m.version = value
}
