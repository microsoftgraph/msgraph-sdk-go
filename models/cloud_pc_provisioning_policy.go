package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CloudPcProvisioningPolicy struct {
    Entity
}
// NewCloudPcProvisioningPolicy instantiates a new CloudPcProvisioningPolicy and sets the default values.
func NewCloudPcProvisioningPolicy()(*CloudPcProvisioningPolicy) {
    m := &CloudPcProvisioningPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudPcProvisioningPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudPcProvisioningPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcProvisioningPolicy(), nil
}
// GetAlternateResourceUrl gets the alternateResourceUrl property value. The alternateResourceUrl property
// returns a *string when successful
func (m *CloudPcProvisioningPolicy) GetAlternateResourceUrl()(*string) {
    val, err := m.GetBackingStore().Get("alternateResourceUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAssignments gets the assignments property value. The assignments property
// returns a []CloudPcProvisioningPolicyAssignmentable when successful
func (m *CloudPcProvisioningPolicy) GetAssignments()([]CloudPcProvisioningPolicyAssignmentable) {
    val, err := m.GetBackingStore().Get("assignments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CloudPcProvisioningPolicyAssignmentable)
    }
    return nil
}
// GetCloudPcGroupDisplayName gets the cloudPcGroupDisplayName property value. The cloudPcGroupDisplayName property
// returns a *string when successful
func (m *CloudPcProvisioningPolicy) GetCloudPcGroupDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("cloudPcGroupDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCloudPcNamingTemplate gets the cloudPcNamingTemplate property value. The cloudPcNamingTemplate property
// returns a *string when successful
func (m *CloudPcProvisioningPolicy) GetCloudPcNamingTemplate()(*string) {
    val, err := m.GetBackingStore().Get("cloudPcNamingTemplate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *CloudPcProvisioningPolicy) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The displayName property
// returns a *string when successful
func (m *CloudPcProvisioningPolicy) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDomainJoinConfigurations gets the domainJoinConfigurations property value. The domainJoinConfigurations property
// returns a []CloudPcDomainJoinConfigurationable when successful
func (m *CloudPcProvisioningPolicy) GetDomainJoinConfigurations()([]CloudPcDomainJoinConfigurationable) {
    val, err := m.GetBackingStore().Get("domainJoinConfigurations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CloudPcDomainJoinConfigurationable)
    }
    return nil
}
// GetEnableSingleSignOn gets the enableSingleSignOn property value. The enableSingleSignOn property
// returns a *bool when successful
func (m *CloudPcProvisioningPolicy) GetEnableSingleSignOn()(*bool) {
    val, err := m.GetBackingStore().Get("enableSingleSignOn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloudPcProvisioningPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alternateResourceUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlternateResourceUrl(val)
        }
        return nil
    }
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcProvisioningPolicyAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcProvisioningPolicyAssignmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CloudPcProvisioningPolicyAssignmentable)
                }
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["cloudPcGroupDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudPcGroupDisplayName(val)
        }
        return nil
    }
    res["cloudPcNamingTemplate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudPcNamingTemplate(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["domainJoinConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcDomainJoinConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcDomainJoinConfigurationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CloudPcDomainJoinConfigurationable)
                }
            }
            m.SetDomainJoinConfigurations(res)
        }
        return nil
    }
    res["enableSingleSignOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableSingleSignOn(val)
        }
        return nil
    }
    res["gracePeriodInHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGracePeriodInHours(val)
        }
        return nil
    }
    res["imageDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImageDisplayName(val)
        }
        return nil
    }
    res["imageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImageId(val)
        }
        return nil
    }
    res["imageType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcProvisioningPolicyImageType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImageType(val.(*CloudPcProvisioningPolicyImageType))
        }
        return nil
    }
    res["localAdminEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalAdminEnabled(val)
        }
        return nil
    }
    res["microsoftManagedDesktop"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMicrosoftManagedDesktopFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftManagedDesktop(val.(MicrosoftManagedDesktopable))
        }
        return nil
    }
    res["provisioningType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcProvisioningType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningType(val.(*CloudPcProvisioningType))
        }
        return nil
    }
    res["windowsSetting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcWindowsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsSetting(val.(CloudPcWindowsSettingable))
        }
        return nil
    }
    return res
}
// GetGracePeriodInHours gets the gracePeriodInHours property value. The gracePeriodInHours property
// returns a *int32 when successful
func (m *CloudPcProvisioningPolicy) GetGracePeriodInHours()(*int32) {
    val, err := m.GetBackingStore().Get("gracePeriodInHours")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetImageDisplayName gets the imageDisplayName property value. The imageDisplayName property
// returns a *string when successful
func (m *CloudPcProvisioningPolicy) GetImageDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("imageDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetImageId gets the imageId property value. The imageId property
// returns a *string when successful
func (m *CloudPcProvisioningPolicy) GetImageId()(*string) {
    val, err := m.GetBackingStore().Get("imageId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetImageType gets the imageType property value. The imageType property
// returns a *CloudPcProvisioningPolicyImageType when successful
func (m *CloudPcProvisioningPolicy) GetImageType()(*CloudPcProvisioningPolicyImageType) {
    val, err := m.GetBackingStore().Get("imageType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcProvisioningPolicyImageType)
    }
    return nil
}
// GetLocalAdminEnabled gets the localAdminEnabled property value. The localAdminEnabled property
// returns a *bool when successful
func (m *CloudPcProvisioningPolicy) GetLocalAdminEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("localAdminEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMicrosoftManagedDesktop gets the microsoftManagedDesktop property value. The microsoftManagedDesktop property
// returns a MicrosoftManagedDesktopable when successful
func (m *CloudPcProvisioningPolicy) GetMicrosoftManagedDesktop()(MicrosoftManagedDesktopable) {
    val, err := m.GetBackingStore().Get("microsoftManagedDesktop")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MicrosoftManagedDesktopable)
    }
    return nil
}
// GetProvisioningType gets the provisioningType property value. The provisioningType property
// returns a *CloudPcProvisioningType when successful
func (m *CloudPcProvisioningPolicy) GetProvisioningType()(*CloudPcProvisioningType) {
    val, err := m.GetBackingStore().Get("provisioningType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcProvisioningType)
    }
    return nil
}
// GetWindowsSetting gets the windowsSetting property value. The windowsSetting property
// returns a CloudPcWindowsSettingable when successful
func (m *CloudPcProvisioningPolicy) GetWindowsSetting()(CloudPcWindowsSettingable) {
    val, err := m.GetBackingStore().Get("windowsSetting")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CloudPcWindowsSettingable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CloudPcProvisioningPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("alternateResourceUrl", m.GetAlternateResourceUrl())
        if err != nil {
            return err
        }
    }
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("cloudPcGroupDisplayName", m.GetCloudPcGroupDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("cloudPcNamingTemplate", m.GetCloudPcNamingTemplate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
    if m.GetDomainJoinConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDomainJoinConfigurations()))
        for i, v := range m.GetDomainJoinConfigurations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("domainJoinConfigurations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableSingleSignOn", m.GetEnableSingleSignOn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("gracePeriodInHours", m.GetGracePeriodInHours())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("imageDisplayName", m.GetImageDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("imageId", m.GetImageId())
        if err != nil {
            return err
        }
    }
    if m.GetImageType() != nil {
        cast := (*m.GetImageType()).String()
        err = writer.WriteStringValue("imageType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localAdminEnabled", m.GetLocalAdminEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("microsoftManagedDesktop", m.GetMicrosoftManagedDesktop())
        if err != nil {
            return err
        }
    }
    if m.GetProvisioningType() != nil {
        cast := (*m.GetProvisioningType()).String()
        err = writer.WriteStringValue("provisioningType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("windowsSetting", m.GetWindowsSetting())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlternateResourceUrl sets the alternateResourceUrl property value. The alternateResourceUrl property
func (m *CloudPcProvisioningPolicy) SetAlternateResourceUrl(value *string)() {
    err := m.GetBackingStore().Set("alternateResourceUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignments sets the assignments property value. The assignments property
func (m *CloudPcProvisioningPolicy) SetAssignments(value []CloudPcProvisioningPolicyAssignmentable)() {
    err := m.GetBackingStore().Set("assignments", value)
    if err != nil {
        panic(err)
    }
}
// SetCloudPcGroupDisplayName sets the cloudPcGroupDisplayName property value. The cloudPcGroupDisplayName property
func (m *CloudPcProvisioningPolicy) SetCloudPcGroupDisplayName(value *string)() {
    err := m.GetBackingStore().Set("cloudPcGroupDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetCloudPcNamingTemplate sets the cloudPcNamingTemplate property value. The cloudPcNamingTemplate property
func (m *CloudPcProvisioningPolicy) SetCloudPcNamingTemplate(value *string)() {
    err := m.GetBackingStore().Set("cloudPcNamingTemplate", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description property
func (m *CloudPcProvisioningPolicy) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *CloudPcProvisioningPolicy) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetDomainJoinConfigurations sets the domainJoinConfigurations property value. The domainJoinConfigurations property
func (m *CloudPcProvisioningPolicy) SetDomainJoinConfigurations(value []CloudPcDomainJoinConfigurationable)() {
    err := m.GetBackingStore().Set("domainJoinConfigurations", value)
    if err != nil {
        panic(err)
    }
}
// SetEnableSingleSignOn sets the enableSingleSignOn property value. The enableSingleSignOn property
func (m *CloudPcProvisioningPolicy) SetEnableSingleSignOn(value *bool)() {
    err := m.GetBackingStore().Set("enableSingleSignOn", value)
    if err != nil {
        panic(err)
    }
}
// SetGracePeriodInHours sets the gracePeriodInHours property value. The gracePeriodInHours property
func (m *CloudPcProvisioningPolicy) SetGracePeriodInHours(value *int32)() {
    err := m.GetBackingStore().Set("gracePeriodInHours", value)
    if err != nil {
        panic(err)
    }
}
// SetImageDisplayName sets the imageDisplayName property value. The imageDisplayName property
func (m *CloudPcProvisioningPolicy) SetImageDisplayName(value *string)() {
    err := m.GetBackingStore().Set("imageDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetImageId sets the imageId property value. The imageId property
func (m *CloudPcProvisioningPolicy) SetImageId(value *string)() {
    err := m.GetBackingStore().Set("imageId", value)
    if err != nil {
        panic(err)
    }
}
// SetImageType sets the imageType property value. The imageType property
func (m *CloudPcProvisioningPolicy) SetImageType(value *CloudPcProvisioningPolicyImageType)() {
    err := m.GetBackingStore().Set("imageType", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalAdminEnabled sets the localAdminEnabled property value. The localAdminEnabled property
func (m *CloudPcProvisioningPolicy) SetLocalAdminEnabled(value *bool)() {
    err := m.GetBackingStore().Set("localAdminEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrosoftManagedDesktop sets the microsoftManagedDesktop property value. The microsoftManagedDesktop property
func (m *CloudPcProvisioningPolicy) SetMicrosoftManagedDesktop(value MicrosoftManagedDesktopable)() {
    err := m.GetBackingStore().Set("microsoftManagedDesktop", value)
    if err != nil {
        panic(err)
    }
}
// SetProvisioningType sets the provisioningType property value. The provisioningType property
func (m *CloudPcProvisioningPolicy) SetProvisioningType(value *CloudPcProvisioningType)() {
    err := m.GetBackingStore().Set("provisioningType", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsSetting sets the windowsSetting property value. The windowsSetting property
func (m *CloudPcProvisioningPolicy) SetWindowsSetting(value CloudPcWindowsSettingable)() {
    err := m.GetBackingStore().Set("windowsSetting", value)
    if err != nil {
        panic(err)
    }
}
type CloudPcProvisioningPolicyable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlternateResourceUrl()(*string)
    GetAssignments()([]CloudPcProvisioningPolicyAssignmentable)
    GetCloudPcGroupDisplayName()(*string)
    GetCloudPcNamingTemplate()(*string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetDomainJoinConfigurations()([]CloudPcDomainJoinConfigurationable)
    GetEnableSingleSignOn()(*bool)
    GetGracePeriodInHours()(*int32)
    GetImageDisplayName()(*string)
    GetImageId()(*string)
    GetImageType()(*CloudPcProvisioningPolicyImageType)
    GetLocalAdminEnabled()(*bool)
    GetMicrosoftManagedDesktop()(MicrosoftManagedDesktopable)
    GetProvisioningType()(*CloudPcProvisioningType)
    GetWindowsSetting()(CloudPcWindowsSettingable)
    SetAlternateResourceUrl(value *string)()
    SetAssignments(value []CloudPcProvisioningPolicyAssignmentable)()
    SetCloudPcGroupDisplayName(value *string)()
    SetCloudPcNamingTemplate(value *string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetDomainJoinConfigurations(value []CloudPcDomainJoinConfigurationable)()
    SetEnableSingleSignOn(value *bool)()
    SetGracePeriodInHours(value *int32)()
    SetImageDisplayName(value *string)()
    SetImageId(value *string)()
    SetImageType(value *CloudPcProvisioningPolicyImageType)()
    SetLocalAdminEnabled(value *bool)()
    SetMicrosoftManagedDesktop(value MicrosoftManagedDesktopable)()
    SetProvisioningType(value *CloudPcProvisioningType)()
    SetWindowsSetting(value CloudPcWindowsSettingable)()
}
