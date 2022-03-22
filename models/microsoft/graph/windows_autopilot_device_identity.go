package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsAutopilotDeviceIdentity 
type WindowsAutopilotDeviceIdentity struct {
    Entity
    // Addressable user name.
    addressableUserName *string;
    // AAD Device ID - to be deprecated
    azureActiveDirectoryDeviceId *string;
    // Display Name
    displayName *string;
    // Intune enrollment state of the Windows autopilot device. Possible values are: unknown, enrolled, pendingReset, failed, notContacted.
    enrollmentState *EnrollmentState;
    // Group Tag of the Windows autopilot device.
    groupTag *string;
    // Intune Last Contacted Date Time of the Windows autopilot device.
    lastContactedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Managed Device ID
    managedDeviceId *string;
    // Oem manufacturer of the Windows autopilot device.
    manufacturer *string;
    // Model name of the Windows autopilot device.
    model *string;
    // Product Key of the Windows autopilot device.
    productKey *string;
    // Purchase Order Identifier of the Windows autopilot device.
    purchaseOrderIdentifier *string;
    // Resource Name.
    resourceName *string;
    // Serial number of the Windows autopilot device.
    serialNumber *string;
    // SKU Number
    skuNumber *string;
    // System Family
    systemFamily *string;
    // User Principal Name.
    userPrincipalName *string;
}
// NewWindowsAutopilotDeviceIdentity instantiates a new windowsAutopilotDeviceIdentity and sets the default values.
func NewWindowsAutopilotDeviceIdentity()(*WindowsAutopilotDeviceIdentity) {
    m := &WindowsAutopilotDeviceIdentity{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWindowsAutopilotDeviceIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsAutopilotDeviceIdentityFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWindowsAutopilotDeviceIdentity(), nil
}
// GetAddressableUserName gets the addressableUserName property value. Addressable user name.
func (m *WindowsAutopilotDeviceIdentity) GetAddressableUserName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.addressableUserName
    }
}
// GetAzureActiveDirectoryDeviceId gets the azureActiveDirectoryDeviceId property value. AAD Device ID - to be deprecated
func (m *WindowsAutopilotDeviceIdentity) GetAzureActiveDirectoryDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureActiveDirectoryDeviceId
    }
}
// GetDisplayName gets the displayName property value. Display Name
func (m *WindowsAutopilotDeviceIdentity) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEnrollmentState gets the enrollmentState property value. Intune enrollment state of the Windows autopilot device. Possible values are: unknown, enrolled, pendingReset, failed, notContacted.
func (m *WindowsAutopilotDeviceIdentity) GetEnrollmentState()(*EnrollmentState) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentState
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsAutopilotDeviceIdentity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["addressableUserName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddressableUserName(val)
        }
        return nil
    }
    res["azureActiveDirectoryDeviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureActiveDirectoryDeviceId(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["enrollmentState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnrollmentState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentState(val.(*EnrollmentState))
        }
        return nil
    }
    res["groupTag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupTag(val)
        }
        return nil
    }
    res["lastContactedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastContactedDateTime(val)
        }
        return nil
    }
    res["managedDeviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceId(val)
        }
        return nil
    }
    res["manufacturer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["model"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["productKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductKey(val)
        }
        return nil
    }
    res["purchaseOrderIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPurchaseOrderIdentifier(val)
        }
        return nil
    }
    res["resourceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceName(val)
        }
        return nil
    }
    res["serialNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSerialNumber(val)
        }
        return nil
    }
    res["skuNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkuNumber(val)
        }
        return nil
    }
    res["systemFamily"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemFamily(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetGroupTag gets the groupTag property value. Group Tag of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) GetGroupTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupTag
    }
}
// GetLastContactedDateTime gets the lastContactedDateTime property value. Intune Last Contacted Date Time of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) GetLastContactedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastContactedDateTime
    }
}
// GetManagedDeviceId gets the managedDeviceId property value. Managed Device ID
func (m *WindowsAutopilotDeviceIdentity) GetManagedDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceId
    }
}
// GetManufacturer gets the manufacturer property value. Oem manufacturer of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
// GetModel gets the model property value. Model name of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
// GetProductKey gets the productKey property value. Product Key of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) GetProductKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.productKey
    }
}
// GetPurchaseOrderIdentifier gets the purchaseOrderIdentifier property value. Purchase Order Identifier of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) GetPurchaseOrderIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.purchaseOrderIdentifier
    }
}
// GetResourceName gets the resourceName property value. Resource Name.
func (m *WindowsAutopilotDeviceIdentity) GetResourceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceName
    }
}
// GetSerialNumber gets the serialNumber property value. Serial number of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) GetSerialNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serialNumber
    }
}
// GetSkuNumber gets the skuNumber property value. SKU Number
func (m *WindowsAutopilotDeviceIdentity) GetSkuNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.skuNumber
    }
}
// GetSystemFamily gets the systemFamily property value. System Family
func (m *WindowsAutopilotDeviceIdentity) GetSystemFamily()(*string) {
    if m == nil {
        return nil
    } else {
        return m.systemFamily
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. User Principal Name.
func (m *WindowsAutopilotDeviceIdentity) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// Serialize serializes information the current object
func (m *WindowsAutopilotDeviceIdentity) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("addressableUserName", m.GetAddressableUserName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureActiveDirectoryDeviceId", m.GetAzureActiveDirectoryDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetEnrollmentState() != nil {
        cast := (*m.GetEnrollmentState()).String()
        err = writer.WriteStringValue("enrollmentState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupTag", m.GetGroupTag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastContactedDateTime", m.GetLastContactedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedDeviceId", m.GetManagedDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("model", m.GetModel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("productKey", m.GetProductKey())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("purchaseOrderIdentifier", m.GetPurchaseOrderIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceName", m.GetResourceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serialNumber", m.GetSerialNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("skuNumber", m.GetSkuNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("systemFamily", m.GetSystemFamily())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddressableUserName sets the addressableUserName property value. Addressable user name.
func (m *WindowsAutopilotDeviceIdentity) SetAddressableUserName(value *string)() {
    if m != nil {
        m.addressableUserName = value
    }
}
// SetAzureActiveDirectoryDeviceId sets the azureActiveDirectoryDeviceId property value. AAD Device ID - to be deprecated
func (m *WindowsAutopilotDeviceIdentity) SetAzureActiveDirectoryDeviceId(value *string)() {
    if m != nil {
        m.azureActiveDirectoryDeviceId = value
    }
}
// SetDisplayName sets the displayName property value. Display Name
func (m *WindowsAutopilotDeviceIdentity) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetEnrollmentState sets the enrollmentState property value. Intune enrollment state of the Windows autopilot device. Possible values are: unknown, enrolled, pendingReset, failed, notContacted.
func (m *WindowsAutopilotDeviceIdentity) SetEnrollmentState(value *EnrollmentState)() {
    if m != nil {
        m.enrollmentState = value
    }
}
// SetGroupTag sets the groupTag property value. Group Tag of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) SetGroupTag(value *string)() {
    if m != nil {
        m.groupTag = value
    }
}
// SetLastContactedDateTime sets the lastContactedDateTime property value. Intune Last Contacted Date Time of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) SetLastContactedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastContactedDateTime = value
    }
}
// SetManagedDeviceId sets the managedDeviceId property value. Managed Device ID
func (m *WindowsAutopilotDeviceIdentity) SetManagedDeviceId(value *string)() {
    if m != nil {
        m.managedDeviceId = value
    }
}
// SetManufacturer sets the manufacturer property value. Oem manufacturer of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) SetManufacturer(value *string)() {
    if m != nil {
        m.manufacturer = value
    }
}
// SetModel sets the model property value. Model name of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) SetModel(value *string)() {
    if m != nil {
        m.model = value
    }
}
// SetProductKey sets the productKey property value. Product Key of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) SetProductKey(value *string)() {
    if m != nil {
        m.productKey = value
    }
}
// SetPurchaseOrderIdentifier sets the purchaseOrderIdentifier property value. Purchase Order Identifier of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) SetPurchaseOrderIdentifier(value *string)() {
    if m != nil {
        m.purchaseOrderIdentifier = value
    }
}
// SetResourceName sets the resourceName property value. Resource Name.
func (m *WindowsAutopilotDeviceIdentity) SetResourceName(value *string)() {
    if m != nil {
        m.resourceName = value
    }
}
// SetSerialNumber sets the serialNumber property value. Serial number of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) SetSerialNumber(value *string)() {
    if m != nil {
        m.serialNumber = value
    }
}
// SetSkuNumber sets the skuNumber property value. SKU Number
func (m *WindowsAutopilotDeviceIdentity) SetSkuNumber(value *string)() {
    if m != nil {
        m.skuNumber = value
    }
}
// SetSystemFamily sets the systemFamily property value. System Family
func (m *WindowsAutopilotDeviceIdentity) SetSystemFamily(value *string)() {
    if m != nil {
        m.systemFamily = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. User Principal Name.
func (m *WindowsAutopilotDeviceIdentity) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
