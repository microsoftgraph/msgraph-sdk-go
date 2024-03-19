package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CloudPcOnPremisesConnection struct {
    Entity
}
// NewCloudPcOnPremisesConnection instantiates a new CloudPcOnPremisesConnection and sets the default values.
func NewCloudPcOnPremisesConnection()(*CloudPcOnPremisesConnection) {
    m := &CloudPcOnPremisesConnection{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudPcOnPremisesConnectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudPcOnPremisesConnectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcOnPremisesConnection(), nil
}
// GetAdDomainName gets the adDomainName property value. The adDomainName property
// returns a *string when successful
func (m *CloudPcOnPremisesConnection) GetAdDomainName()(*string) {
    val, err := m.GetBackingStore().Get("adDomainName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAdDomainPassword gets the adDomainPassword property value. The adDomainPassword property
// returns a *string when successful
func (m *CloudPcOnPremisesConnection) GetAdDomainPassword()(*string) {
    val, err := m.GetBackingStore().Get("adDomainPassword")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAdDomainUsername gets the adDomainUsername property value. The adDomainUsername property
// returns a *string when successful
func (m *CloudPcOnPremisesConnection) GetAdDomainUsername()(*string) {
    val, err := m.GetBackingStore().Get("adDomainUsername")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAlternateResourceUrl gets the alternateResourceUrl property value. The alternateResourceUrl property
// returns a *string when successful
func (m *CloudPcOnPremisesConnection) GetAlternateResourceUrl()(*string) {
    val, err := m.GetBackingStore().Get("alternateResourceUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetConnectionType gets the connectionType property value. The connectionType property
// returns a *CloudPcOnPremisesConnectionType when successful
func (m *CloudPcOnPremisesConnection) GetConnectionType()(*CloudPcOnPremisesConnectionType) {
    val, err := m.GetBackingStore().Get("connectionType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcOnPremisesConnectionType)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The displayName property
// returns a *string when successful
func (m *CloudPcOnPremisesConnection) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloudPcOnPremisesConnection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["adDomainName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdDomainName(val)
        }
        return nil
    }
    res["adDomainPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdDomainPassword(val)
        }
        return nil
    }
    res["adDomainUsername"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdDomainUsername(val)
        }
        return nil
    }
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
    res["connectionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcOnPremisesConnectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectionType(val.(*CloudPcOnPremisesConnectionType))
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
    res["healthCheckStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcOnPremisesConnectionStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealthCheckStatus(val.(*CloudPcOnPremisesConnectionStatus))
        }
        return nil
    }
    res["healthCheckStatusDetail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcOnPremisesConnectionStatusDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealthCheckStatusDetail(val.(CloudPcOnPremisesConnectionStatusDetailable))
        }
        return nil
    }
    res["inUse"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInUse(val)
        }
        return nil
    }
    res["organizationalUnit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationalUnit(val)
        }
        return nil
    }
    res["resourceGroupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceGroupId(val)
        }
        return nil
    }
    res["subnetId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubnetId(val)
        }
        return nil
    }
    res["subscriptionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriptionId(val)
        }
        return nil
    }
    res["subscriptionName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriptionName(val)
        }
        return nil
    }
    res["virtualNetworkId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVirtualNetworkId(val)
        }
        return nil
    }
    res["virtualNetworkLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVirtualNetworkLocation(val)
        }
        return nil
    }
    return res
}
// GetHealthCheckStatus gets the healthCheckStatus property value. The healthCheckStatus property
// returns a *CloudPcOnPremisesConnectionStatus when successful
func (m *CloudPcOnPremisesConnection) GetHealthCheckStatus()(*CloudPcOnPremisesConnectionStatus) {
    val, err := m.GetBackingStore().Get("healthCheckStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcOnPremisesConnectionStatus)
    }
    return nil
}
// GetHealthCheckStatusDetail gets the healthCheckStatusDetail property value. The healthCheckStatusDetail property
// returns a CloudPcOnPremisesConnectionStatusDetailable when successful
func (m *CloudPcOnPremisesConnection) GetHealthCheckStatusDetail()(CloudPcOnPremisesConnectionStatusDetailable) {
    val, err := m.GetBackingStore().Get("healthCheckStatusDetail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CloudPcOnPremisesConnectionStatusDetailable)
    }
    return nil
}
// GetInUse gets the inUse property value. The inUse property
// returns a *bool when successful
func (m *CloudPcOnPremisesConnection) GetInUse()(*bool) {
    val, err := m.GetBackingStore().Get("inUse")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOrganizationalUnit gets the organizationalUnit property value. The organizationalUnit property
// returns a *string when successful
func (m *CloudPcOnPremisesConnection) GetOrganizationalUnit()(*string) {
    val, err := m.GetBackingStore().Get("organizationalUnit")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetResourceGroupId gets the resourceGroupId property value. The resourceGroupId property
// returns a *string when successful
func (m *CloudPcOnPremisesConnection) GetResourceGroupId()(*string) {
    val, err := m.GetBackingStore().Get("resourceGroupId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSubnetId gets the subnetId property value. The subnetId property
// returns a *string when successful
func (m *CloudPcOnPremisesConnection) GetSubnetId()(*string) {
    val, err := m.GetBackingStore().Get("subnetId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSubscriptionId gets the subscriptionId property value. The subscriptionId property
// returns a *string when successful
func (m *CloudPcOnPremisesConnection) GetSubscriptionId()(*string) {
    val, err := m.GetBackingStore().Get("subscriptionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSubscriptionName gets the subscriptionName property value. The subscriptionName property
// returns a *string when successful
func (m *CloudPcOnPremisesConnection) GetSubscriptionName()(*string) {
    val, err := m.GetBackingStore().Get("subscriptionName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetVirtualNetworkId gets the virtualNetworkId property value. The virtualNetworkId property
// returns a *string when successful
func (m *CloudPcOnPremisesConnection) GetVirtualNetworkId()(*string) {
    val, err := m.GetBackingStore().Get("virtualNetworkId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetVirtualNetworkLocation gets the virtualNetworkLocation property value. The virtualNetworkLocation property
// returns a *string when successful
func (m *CloudPcOnPremisesConnection) GetVirtualNetworkLocation()(*string) {
    val, err := m.GetBackingStore().Get("virtualNetworkLocation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CloudPcOnPremisesConnection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("adDomainName", m.GetAdDomainName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("adDomainPassword", m.GetAdDomainPassword())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("adDomainUsername", m.GetAdDomainUsername())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("alternateResourceUrl", m.GetAlternateResourceUrl())
        if err != nil {
            return err
        }
    }
    if m.GetConnectionType() != nil {
        cast := (*m.GetConnectionType()).String()
        err = writer.WriteStringValue("connectionType", &cast)
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
    if m.GetHealthCheckStatus() != nil {
        cast := (*m.GetHealthCheckStatus()).String()
        err = writer.WriteStringValue("healthCheckStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("healthCheckStatusDetail", m.GetHealthCheckStatusDetail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("inUse", m.GetInUse())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("organizationalUnit", m.GetOrganizationalUnit())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceGroupId", m.GetResourceGroupId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subnetId", m.GetSubnetId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subscriptionId", m.GetSubscriptionId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subscriptionName", m.GetSubscriptionName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("virtualNetworkId", m.GetVirtualNetworkId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("virtualNetworkLocation", m.GetVirtualNetworkLocation())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdDomainName sets the adDomainName property value. The adDomainName property
func (m *CloudPcOnPremisesConnection) SetAdDomainName(value *string)() {
    err := m.GetBackingStore().Set("adDomainName", value)
    if err != nil {
        panic(err)
    }
}
// SetAdDomainPassword sets the adDomainPassword property value. The adDomainPassword property
func (m *CloudPcOnPremisesConnection) SetAdDomainPassword(value *string)() {
    err := m.GetBackingStore().Set("adDomainPassword", value)
    if err != nil {
        panic(err)
    }
}
// SetAdDomainUsername sets the adDomainUsername property value. The adDomainUsername property
func (m *CloudPcOnPremisesConnection) SetAdDomainUsername(value *string)() {
    err := m.GetBackingStore().Set("adDomainUsername", value)
    if err != nil {
        panic(err)
    }
}
// SetAlternateResourceUrl sets the alternateResourceUrl property value. The alternateResourceUrl property
func (m *CloudPcOnPremisesConnection) SetAlternateResourceUrl(value *string)() {
    err := m.GetBackingStore().Set("alternateResourceUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectionType sets the connectionType property value. The connectionType property
func (m *CloudPcOnPremisesConnection) SetConnectionType(value *CloudPcOnPremisesConnectionType)() {
    err := m.GetBackingStore().Set("connectionType", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *CloudPcOnPremisesConnection) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetHealthCheckStatus sets the healthCheckStatus property value. The healthCheckStatus property
func (m *CloudPcOnPremisesConnection) SetHealthCheckStatus(value *CloudPcOnPremisesConnectionStatus)() {
    err := m.GetBackingStore().Set("healthCheckStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetHealthCheckStatusDetail sets the healthCheckStatusDetail property value. The healthCheckStatusDetail property
func (m *CloudPcOnPremisesConnection) SetHealthCheckStatusDetail(value CloudPcOnPremisesConnectionStatusDetailable)() {
    err := m.GetBackingStore().Set("healthCheckStatusDetail", value)
    if err != nil {
        panic(err)
    }
}
// SetInUse sets the inUse property value. The inUse property
func (m *CloudPcOnPremisesConnection) SetInUse(value *bool)() {
    err := m.GetBackingStore().Set("inUse", value)
    if err != nil {
        panic(err)
    }
}
// SetOrganizationalUnit sets the organizationalUnit property value. The organizationalUnit property
func (m *CloudPcOnPremisesConnection) SetOrganizationalUnit(value *string)() {
    err := m.GetBackingStore().Set("organizationalUnit", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceGroupId sets the resourceGroupId property value. The resourceGroupId property
func (m *CloudPcOnPremisesConnection) SetResourceGroupId(value *string)() {
    err := m.GetBackingStore().Set("resourceGroupId", value)
    if err != nil {
        panic(err)
    }
}
// SetSubnetId sets the subnetId property value. The subnetId property
func (m *CloudPcOnPremisesConnection) SetSubnetId(value *string)() {
    err := m.GetBackingStore().Set("subnetId", value)
    if err != nil {
        panic(err)
    }
}
// SetSubscriptionId sets the subscriptionId property value. The subscriptionId property
func (m *CloudPcOnPremisesConnection) SetSubscriptionId(value *string)() {
    err := m.GetBackingStore().Set("subscriptionId", value)
    if err != nil {
        panic(err)
    }
}
// SetSubscriptionName sets the subscriptionName property value. The subscriptionName property
func (m *CloudPcOnPremisesConnection) SetSubscriptionName(value *string)() {
    err := m.GetBackingStore().Set("subscriptionName", value)
    if err != nil {
        panic(err)
    }
}
// SetVirtualNetworkId sets the virtualNetworkId property value. The virtualNetworkId property
func (m *CloudPcOnPremisesConnection) SetVirtualNetworkId(value *string)() {
    err := m.GetBackingStore().Set("virtualNetworkId", value)
    if err != nil {
        panic(err)
    }
}
// SetVirtualNetworkLocation sets the virtualNetworkLocation property value. The virtualNetworkLocation property
func (m *CloudPcOnPremisesConnection) SetVirtualNetworkLocation(value *string)() {
    err := m.GetBackingStore().Set("virtualNetworkLocation", value)
    if err != nil {
        panic(err)
    }
}
type CloudPcOnPremisesConnectionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdDomainName()(*string)
    GetAdDomainPassword()(*string)
    GetAdDomainUsername()(*string)
    GetAlternateResourceUrl()(*string)
    GetConnectionType()(*CloudPcOnPremisesConnectionType)
    GetDisplayName()(*string)
    GetHealthCheckStatus()(*CloudPcOnPremisesConnectionStatus)
    GetHealthCheckStatusDetail()(CloudPcOnPremisesConnectionStatusDetailable)
    GetInUse()(*bool)
    GetOrganizationalUnit()(*string)
    GetResourceGroupId()(*string)
    GetSubnetId()(*string)
    GetSubscriptionId()(*string)
    GetSubscriptionName()(*string)
    GetVirtualNetworkId()(*string)
    GetVirtualNetworkLocation()(*string)
    SetAdDomainName(value *string)()
    SetAdDomainPassword(value *string)()
    SetAdDomainUsername(value *string)()
    SetAlternateResourceUrl(value *string)()
    SetConnectionType(value *CloudPcOnPremisesConnectionType)()
    SetDisplayName(value *string)()
    SetHealthCheckStatus(value *CloudPcOnPremisesConnectionStatus)()
    SetHealthCheckStatusDetail(value CloudPcOnPremisesConnectionStatusDetailable)()
    SetInUse(value *bool)()
    SetOrganizationalUnit(value *string)()
    SetResourceGroupId(value *string)()
    SetSubnetId(value *string)()
    SetSubscriptionId(value *string)()
    SetSubscriptionName(value *string)()
    SetVirtualNetworkId(value *string)()
    SetVirtualNetworkLocation(value *string)()
}
