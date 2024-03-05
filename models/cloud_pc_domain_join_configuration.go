package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type CloudPcDomainJoinConfiguration struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewCloudPcDomainJoinConfiguration instantiates a new CloudPcDomainJoinConfiguration and sets the default values.
func NewCloudPcDomainJoinConfiguration()(*CloudPcDomainJoinConfiguration) {
    m := &CloudPcDomainJoinConfiguration{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCloudPcDomainJoinConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudPcDomainJoinConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcDomainJoinConfiguration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CloudPcDomainJoinConfiguration) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *CloudPcDomainJoinConfiguration) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDomainJoinType gets the domainJoinType property value. The domainJoinType property
// returns a *CloudPcDomainJoinType when successful
func (m *CloudPcDomainJoinConfiguration) GetDomainJoinType()(*CloudPcDomainJoinType) {
    val, err := m.GetBackingStore().Get("domainJoinType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcDomainJoinType)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloudPcDomainJoinConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["domainJoinType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcDomainJoinType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainJoinType(val.(*CloudPcDomainJoinType))
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
    res["onPremisesConnectionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesConnectionId(val)
        }
        return nil
    }
    res["regionGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcRegionGroup)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegionGroup(val.(*CloudPcRegionGroup))
        }
        return nil
    }
    res["regionName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegionName(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *CloudPcDomainJoinConfiguration) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOnPremisesConnectionId gets the onPremisesConnectionId property value. The onPremisesConnectionId property
// returns a *string when successful
func (m *CloudPcDomainJoinConfiguration) GetOnPremisesConnectionId()(*string) {
    val, err := m.GetBackingStore().Get("onPremisesConnectionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRegionGroup gets the regionGroup property value. The regionGroup property
// returns a *CloudPcRegionGroup when successful
func (m *CloudPcDomainJoinConfiguration) GetRegionGroup()(*CloudPcRegionGroup) {
    val, err := m.GetBackingStore().Get("regionGroup")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcRegionGroup)
    }
    return nil
}
// GetRegionName gets the regionName property value. The regionName property
// returns a *string when successful
func (m *CloudPcDomainJoinConfiguration) GetRegionName()(*string) {
    val, err := m.GetBackingStore().Get("regionName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CloudPcDomainJoinConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDomainJoinType() != nil {
        cast := (*m.GetDomainJoinType()).String()
        err := writer.WriteStringValue("domainJoinType", &cast)
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
        err := writer.WriteStringValue("onPremisesConnectionId", m.GetOnPremisesConnectionId())
        if err != nil {
            return err
        }
    }
    if m.GetRegionGroup() != nil {
        cast := (*m.GetRegionGroup()).String()
        err := writer.WriteStringValue("regionGroup", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("regionName", m.GetRegionName())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcDomainJoinConfiguration) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *CloudPcDomainJoinConfiguration) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDomainJoinType sets the domainJoinType property value. The domainJoinType property
func (m *CloudPcDomainJoinConfiguration) SetDomainJoinType(value *CloudPcDomainJoinType)() {
    err := m.GetBackingStore().Set("domainJoinType", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CloudPcDomainJoinConfiguration) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOnPremisesConnectionId sets the onPremisesConnectionId property value. The onPremisesConnectionId property
func (m *CloudPcDomainJoinConfiguration) SetOnPremisesConnectionId(value *string)() {
    err := m.GetBackingStore().Set("onPremisesConnectionId", value)
    if err != nil {
        panic(err)
    }
}
// SetRegionGroup sets the regionGroup property value. The regionGroup property
func (m *CloudPcDomainJoinConfiguration) SetRegionGroup(value *CloudPcRegionGroup)() {
    err := m.GetBackingStore().Set("regionGroup", value)
    if err != nil {
        panic(err)
    }
}
// SetRegionName sets the regionName property value. The regionName property
func (m *CloudPcDomainJoinConfiguration) SetRegionName(value *string)() {
    err := m.GetBackingStore().Set("regionName", value)
    if err != nil {
        panic(err)
    }
}
type CloudPcDomainJoinConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDomainJoinType()(*CloudPcDomainJoinType)
    GetOdataType()(*string)
    GetOnPremisesConnectionId()(*string)
    GetRegionGroup()(*CloudPcRegionGroup)
    GetRegionName()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDomainJoinType(value *CloudPcDomainJoinType)()
    SetOdataType(value *string)()
    SetOnPremisesConnectionId(value *string)()
    SetRegionGroup(value *CloudPcRegionGroup)()
    SetRegionName(value *string)()
}
