package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Payload 
type Payload struct {
    Entity
}
// NewPayload instantiates a new payload and sets the default values.
func NewPayload()(*Payload) {
    m := &Payload{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePayloadFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePayloadFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPayload(), nil
}
// GetBrand gets the brand property value. The brand property
func (m *Payload) GetBrand()(*PayloadBrand) {
    val, err := m.GetBackingStore().Get("brand")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PayloadBrand)
    }
    return nil
}
// GetComplexity gets the complexity property value. The complexity property
func (m *Payload) GetComplexity()(*PayloadComplexity) {
    val, err := m.GetBackingStore().Get("complexity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PayloadComplexity)
    }
    return nil
}
// GetCreatedBy gets the createdBy property value. The createdBy property
func (m *Payload) GetCreatedBy()(EmailIdentityable) {
    val, err := m.GetBackingStore().Get("createdBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EmailIdentityable)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *Payload) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDescription gets the description property value. The description property
func (m *Payload) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDetail gets the detail property value. The detail property
func (m *Payload) GetDetail()(PayloadDetailable) {
    val, err := m.GetBackingStore().Get("detail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PayloadDetailable)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *Payload) GetDisplayName()(*string) {
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
func (m *Payload) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["brand"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePayloadBrand)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBrand(val.(*PayloadBrand))
        }
        return nil
    }
    res["complexity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePayloadComplexity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplexity(val.(*PayloadComplexity))
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEmailIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(EmailIdentityable))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
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
    res["detail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePayloadDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetail(val.(PayloadDetailable))
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
    res["industry"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePayloadIndustry)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndustry(val.(*PayloadIndustry))
        }
        return nil
    }
    res["isAutomated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAutomated(val)
        }
        return nil
    }
    res["isControversial"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsControversial(val)
        }
        return nil
    }
    res["isCurrentEvent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCurrentEvent(val)
        }
        return nil
    }
    res["language"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguage(val)
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEmailIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(EmailIdentityable))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["payloadTags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetPayloadTags(res)
        }
        return nil
    }
    res["platform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePayloadDeliveryPlatform)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatform(val.(*PayloadDeliveryPlatform))
        }
        return nil
    }
    res["predictedCompromiseRate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPredictedCompromiseRate(val)
        }
        return nil
    }
    res["simulationAttackType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSimulationAttackType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSimulationAttackType(val.(*SimulationAttackType))
        }
        return nil
    }
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSimulationContentSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(*SimulationContentSource))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSimulationContentStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*SimulationContentStatus))
        }
        return nil
    }
    res["technique"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSimulationAttackTechnique)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTechnique(val.(*SimulationAttackTechnique))
        }
        return nil
    }
    res["theme"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePayloadTheme)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTheme(val.(*PayloadTheme))
        }
        return nil
    }
    return res
}
// GetIndustry gets the industry property value. The industry property
func (m *Payload) GetIndustry()(*PayloadIndustry) {
    val, err := m.GetBackingStore().Get("industry")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PayloadIndustry)
    }
    return nil
}
// GetIsAutomated gets the isAutomated property value. The isAutomated property
func (m *Payload) GetIsAutomated()(*bool) {
    val, err := m.GetBackingStore().Get("isAutomated")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsControversial gets the isControversial property value. The isControversial property
func (m *Payload) GetIsControversial()(*bool) {
    val, err := m.GetBackingStore().Get("isControversial")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsCurrentEvent gets the isCurrentEvent property value. The isCurrentEvent property
func (m *Payload) GetIsCurrentEvent()(*bool) {
    val, err := m.GetBackingStore().Get("isCurrentEvent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLanguage gets the language property value. The language property
func (m *Payload) GetLanguage()(*string) {
    val, err := m.GetBackingStore().Get("language")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLastModifiedBy gets the lastModifiedBy property value. The lastModifiedBy property
func (m *Payload) GetLastModifiedBy()(EmailIdentityable) {
    val, err := m.GetBackingStore().Get("lastModifiedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EmailIdentityable)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *Payload) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetPayloadTags gets the payloadTags property value. The payloadTags property
func (m *Payload) GetPayloadTags()([]string) {
    val, err := m.GetBackingStore().Get("payloadTags")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetPlatform gets the platform property value. The platform property
func (m *Payload) GetPlatform()(*PayloadDeliveryPlatform) {
    val, err := m.GetBackingStore().Get("platform")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PayloadDeliveryPlatform)
    }
    return nil
}
// GetPredictedCompromiseRate gets the predictedCompromiseRate property value. The predictedCompromiseRate property
func (m *Payload) GetPredictedCompromiseRate()(*float64) {
    val, err := m.GetBackingStore().Get("predictedCompromiseRate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetSimulationAttackType gets the simulationAttackType property value. The simulationAttackType property
func (m *Payload) GetSimulationAttackType()(*SimulationAttackType) {
    val, err := m.GetBackingStore().Get("simulationAttackType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SimulationAttackType)
    }
    return nil
}
// GetSource gets the source property value. The source property
func (m *Payload) GetSource()(*SimulationContentSource) {
    val, err := m.GetBackingStore().Get("source")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SimulationContentSource)
    }
    return nil
}
// GetStatus gets the status property value. The status property
func (m *Payload) GetStatus()(*SimulationContentStatus) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SimulationContentStatus)
    }
    return nil
}
// GetTechnique gets the technique property value. The technique property
func (m *Payload) GetTechnique()(*SimulationAttackTechnique) {
    val, err := m.GetBackingStore().Get("technique")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SimulationAttackTechnique)
    }
    return nil
}
// GetTheme gets the theme property value. The theme property
func (m *Payload) GetTheme()(*PayloadTheme) {
    val, err := m.GetBackingStore().Get("theme")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PayloadTheme)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Payload) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetBrand() != nil {
        cast := (*m.GetBrand()).String()
        err = writer.WriteStringValue("brand", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetComplexity() != nil {
        cast := (*m.GetComplexity()).String()
        err = writer.WriteStringValue("complexity", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
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
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("detail", m.GetDetail())
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
    if m.GetIndustry() != nil {
        cast := (*m.GetIndustry()).String()
        err = writer.WriteStringValue("industry", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAutomated", m.GetIsAutomated())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isControversial", m.GetIsControversial())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isCurrentEvent", m.GetIsCurrentEvent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("language", m.GetLanguage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPayloadTags() != nil {
        err = writer.WriteCollectionOfStringValues("payloadTags", m.GetPayloadTags())
        if err != nil {
            return err
        }
    }
    if m.GetPlatform() != nil {
        cast := (*m.GetPlatform()).String()
        err = writer.WriteStringValue("platform", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("predictedCompromiseRate", m.GetPredictedCompromiseRate())
        if err != nil {
            return err
        }
    }
    if m.GetSimulationAttackType() != nil {
        cast := (*m.GetSimulationAttackType()).String()
        err = writer.WriteStringValue("simulationAttackType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSource() != nil {
        cast := (*m.GetSource()).String()
        err = writer.WriteStringValue("source", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTechnique() != nil {
        cast := (*m.GetTechnique()).String()
        err = writer.WriteStringValue("technique", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTheme() != nil {
        cast := (*m.GetTheme()).String()
        err = writer.WriteStringValue("theme", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBrand sets the brand property value. The brand property
func (m *Payload) SetBrand(value *PayloadBrand)() {
    err := m.GetBackingStore().Set("brand", value)
    if err != nil {
        panic(err)
    }
}
// SetComplexity sets the complexity property value. The complexity property
func (m *Payload) SetComplexity(value *PayloadComplexity)() {
    err := m.GetBackingStore().Set("complexity", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedBy sets the createdBy property value. The createdBy property
func (m *Payload) SetCreatedBy(value EmailIdentityable)() {
    err := m.GetBackingStore().Set("createdBy", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *Payload) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description property
func (m *Payload) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDetail sets the detail property value. The detail property
func (m *Payload) SetDetail(value PayloadDetailable)() {
    err := m.GetBackingStore().Set("detail", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *Payload) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetIndustry sets the industry property value. The industry property
func (m *Payload) SetIndustry(value *PayloadIndustry)() {
    err := m.GetBackingStore().Set("industry", value)
    if err != nil {
        panic(err)
    }
}
// SetIsAutomated sets the isAutomated property value. The isAutomated property
func (m *Payload) SetIsAutomated(value *bool)() {
    err := m.GetBackingStore().Set("isAutomated", value)
    if err != nil {
        panic(err)
    }
}
// SetIsControversial sets the isControversial property value. The isControversial property
func (m *Payload) SetIsControversial(value *bool)() {
    err := m.GetBackingStore().Set("isControversial", value)
    if err != nil {
        panic(err)
    }
}
// SetIsCurrentEvent sets the isCurrentEvent property value. The isCurrentEvent property
func (m *Payload) SetIsCurrentEvent(value *bool)() {
    err := m.GetBackingStore().Set("isCurrentEvent", value)
    if err != nil {
        panic(err)
    }
}
// SetLanguage sets the language property value. The language property
func (m *Payload) SetLanguage(value *string)() {
    err := m.GetBackingStore().Set("language", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedBy sets the lastModifiedBy property value. The lastModifiedBy property
func (m *Payload) SetLastModifiedBy(value EmailIdentityable)() {
    err := m.GetBackingStore().Set("lastModifiedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *Payload) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetPayloadTags sets the payloadTags property value. The payloadTags property
func (m *Payload) SetPayloadTags(value []string)() {
    err := m.GetBackingStore().Set("payloadTags", value)
    if err != nil {
        panic(err)
    }
}
// SetPlatform sets the platform property value. The platform property
func (m *Payload) SetPlatform(value *PayloadDeliveryPlatform)() {
    err := m.GetBackingStore().Set("platform", value)
    if err != nil {
        panic(err)
    }
}
// SetPredictedCompromiseRate sets the predictedCompromiseRate property value. The predictedCompromiseRate property
func (m *Payload) SetPredictedCompromiseRate(value *float64)() {
    err := m.GetBackingStore().Set("predictedCompromiseRate", value)
    if err != nil {
        panic(err)
    }
}
// SetSimulationAttackType sets the simulationAttackType property value. The simulationAttackType property
func (m *Payload) SetSimulationAttackType(value *SimulationAttackType)() {
    err := m.GetBackingStore().Set("simulationAttackType", value)
    if err != nil {
        panic(err)
    }
}
// SetSource sets the source property value. The source property
func (m *Payload) SetSource(value *SimulationContentSource)() {
    err := m.GetBackingStore().Set("source", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The status property
func (m *Payload) SetStatus(value *SimulationContentStatus)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetTechnique sets the technique property value. The technique property
func (m *Payload) SetTechnique(value *SimulationAttackTechnique)() {
    err := m.GetBackingStore().Set("technique", value)
    if err != nil {
        panic(err)
    }
}
// SetTheme sets the theme property value. The theme property
func (m *Payload) SetTheme(value *PayloadTheme)() {
    err := m.GetBackingStore().Set("theme", value)
    if err != nil {
        panic(err)
    }
}
// Payloadable 
type Payloadable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBrand()(*PayloadBrand)
    GetComplexity()(*PayloadComplexity)
    GetCreatedBy()(EmailIdentityable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDetail()(PayloadDetailable)
    GetDisplayName()(*string)
    GetIndustry()(*PayloadIndustry)
    GetIsAutomated()(*bool)
    GetIsControversial()(*bool)
    GetIsCurrentEvent()(*bool)
    GetLanguage()(*string)
    GetLastModifiedBy()(EmailIdentityable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPayloadTags()([]string)
    GetPlatform()(*PayloadDeliveryPlatform)
    GetPredictedCompromiseRate()(*float64)
    GetSimulationAttackType()(*SimulationAttackType)
    GetSource()(*SimulationContentSource)
    GetStatus()(*SimulationContentStatus)
    GetTechnique()(*SimulationAttackTechnique)
    GetTheme()(*PayloadTheme)
    SetBrand(value *PayloadBrand)()
    SetComplexity(value *PayloadComplexity)()
    SetCreatedBy(value EmailIdentityable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDetail(value PayloadDetailable)()
    SetDisplayName(value *string)()
    SetIndustry(value *PayloadIndustry)()
    SetIsAutomated(value *bool)()
    SetIsControversial(value *bool)()
    SetIsCurrentEvent(value *bool)()
    SetLanguage(value *string)()
    SetLastModifiedBy(value EmailIdentityable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPayloadTags(value []string)()
    SetPlatform(value *PayloadDeliveryPlatform)()
    SetPredictedCompromiseRate(value *float64)()
    SetSimulationAttackType(value *SimulationAttackType)()
    SetSource(value *SimulationContentSource)()
    SetStatus(value *SimulationContentStatus)()
    SetTechnique(value *SimulationAttackTechnique)()
    SetTheme(value *PayloadTheme)()
}
