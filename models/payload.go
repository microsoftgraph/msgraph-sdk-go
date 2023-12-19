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
// GetBrand gets the brand property value. The branch of a payload. Possible values are: unknown, other, americanExpress, capitalOne, dhl, docuSign, dropbox, facebook, firstAmerican, microsoft, netflix, scotiabank, sendGrid, stewartTitle, tesco, wellsFargo, syrinxCloud, adobe, teams, zoom, unknownFutureValue.
func (m *Payload) GetBrand()(*Payload_brand) {
    val, err := m.GetBackingStore().Get("brand")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Payload_brand)
    }
    return nil
}
// GetComplexity gets the complexity property value. The complexity of a payload. Possible values are: unknown, low, medium, high, unknownFutureValue.
func (m *Payload) GetComplexity()(*Payload_complexity) {
    val, err := m.GetBackingStore().Get("complexity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Payload_complexity)
    }
    return nil
}
// GetCreatedBy gets the createdBy property value. Identity of the user who created the attack simulation and training campaign payload.
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
// GetCreatedDateTime gets the createdDateTime property value. Date and time when the attack simulation and training campaign payload. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
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
// GetDescription gets the description property value. Description of the attack simulation and training campaign payload.
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
// GetDetail gets the detail property value. Additional details about the payload.
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
// GetDisplayName gets the displayName property value. Display name of the attack simulation and training campaign payload. Supports $filter and $orderby.
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
        val, err := n.GetEnumValue(ParsePayload_brand)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBrand(val.(*Payload_brand))
        }
        return nil
    }
    res["complexity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePayload_complexity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplexity(val.(*Payload_complexity))
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
        val, err := n.GetEnumValue(ParsePayload_industry)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndustry(val.(*Payload_industry))
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
        val, err := n.GetEnumValue(ParsePayload_platform)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatform(val.(*Payload_platform))
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
        val, err := n.GetEnumValue(ParsePayload_simulationAttackType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSimulationAttackType(val.(*Payload_simulationAttackType))
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
        val, err := n.GetEnumValue(ParsePayload_status)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*Payload_status))
        }
        return nil
    }
    res["technique"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePayload_technique)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTechnique(val.(*Payload_technique))
        }
        return nil
    }
    res["theme"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePayload_theme)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTheme(val.(*Payload_theme))
        }
        return nil
    }
    return res
}
// GetIndustry gets the industry property value. Industry of a payload. Possible values are: unknown, other, banking, businessServices, consumerServices, education, energy, construction, consulting, financialServices, government, hospitality, insurance, legal, courierServices, IT, healthcare, manufacturing, retail, telecom, realEstate, unknownFutureValue.
func (m *Payload) GetIndustry()(*Payload_industry) {
    val, err := m.GetBackingStore().Get("industry")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Payload_industry)
    }
    return nil
}
// GetIsAutomated gets the isAutomated property value. Indicates whether the attack simulation and training campaign payload was created from an automation flow. Supports $filter and $orderby.
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
// GetIsControversial gets the isControversial property value. Indicates whether the payload is controversial.
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
// GetIsCurrentEvent gets the isCurrentEvent property value. Indicates whether the payload is from any recent event.
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
// GetLanguage gets the language property value. Payload language.
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
// GetLastModifiedBy gets the lastModifiedBy property value. Identity of the user who most recently modified the attack simulation and training campaign payload.
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
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Date and time when the attack simulation and training campaign payload was last modified. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
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
// GetPayloadTags gets the payloadTags property value. Free text tags for a payload.
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
// GetPlatform gets the platform property value. The payload delivery platform for a simulation. Possible values are: unknown, sms, email, teams, unknownFutureValue.
func (m *Payload) GetPlatform()(*Payload_platform) {
    val, err := m.GetBackingStore().Get("platform")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Payload_platform)
    }
    return nil
}
// GetPredictedCompromiseRate gets the predictedCompromiseRate property value. Predicted probability for a payload to phish a targeted user.
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
// GetSimulationAttackType gets the simulationAttackType property value. Attack type of the attack simulation and training campaign. Supports $filter and $orderby. Possible values are: unknown, social, cloud, endpoint, unknownFutureValue.
func (m *Payload) GetSimulationAttackType()(*Payload_simulationAttackType) {
    val, err := m.GetBackingStore().Get("simulationAttackType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Payload_simulationAttackType)
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
// GetStatus gets the status property value. Simulation content status. Supports $filter and $orderby. Possible values are: unknown, draft, ready, archive, delete, unknownFutureValue.
func (m *Payload) GetStatus()(*Payload_status) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Payload_status)
    }
    return nil
}
// GetTechnique gets the technique property value. The social engineering technique used in the attack simulation and training campaign. Supports $filter and $orderby. Possible values are: unknown, credentialHarvesting, attachmentMalware, driveByUrl, linkInAttachment, linkToMalwareFile, unknownFutureValue, oAuthConsentGrant. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values from this evolvable enum: oAuthConsentGrant. For more information on the types of social engineering attack techniques, see simulations.
func (m *Payload) GetTechnique()(*Payload_technique) {
    val, err := m.GetBackingStore().Get("technique")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Payload_technique)
    }
    return nil
}
// GetTheme gets the theme property value. The theme of a payload. Possible values are: unknown, other, accountActivation, accountVerification, billing, cleanUpMail, controversial, documentReceived, expense, fax, financeReport, incomingMessages, invoice, itemReceived, loginAlert, mailReceived, password, payment, payroll, personalizedOffer, quarantine, remoteWork, reviewMessage, securityUpdate, serviceSuspended, signatureRequired, upgradeMailboxStorage, verifyMailbox, voicemail, advertisement, employeeEngagement, unknownFutureValue.
func (m *Payload) GetTheme()(*Payload_theme) {
    val, err := m.GetBackingStore().Get("theme")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Payload_theme)
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
// SetBrand sets the brand property value. The branch of a payload. Possible values are: unknown, other, americanExpress, capitalOne, dhl, docuSign, dropbox, facebook, firstAmerican, microsoft, netflix, scotiabank, sendGrid, stewartTitle, tesco, wellsFargo, syrinxCloud, adobe, teams, zoom, unknownFutureValue.
func (m *Payload) SetBrand(value *Payload_brand)() {
    err := m.GetBackingStore().Set("brand", value)
    if err != nil {
        panic(err)
    }
}
// SetComplexity sets the complexity property value. The complexity of a payload. Possible values are: unknown, low, medium, high, unknownFutureValue.
func (m *Payload) SetComplexity(value *Payload_complexity)() {
    err := m.GetBackingStore().Set("complexity", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedBy sets the createdBy property value. Identity of the user who created the attack simulation and training campaign payload.
func (m *Payload) SetCreatedBy(value EmailIdentityable)() {
    err := m.GetBackingStore().Set("createdBy", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Date and time when the attack simulation and training campaign payload. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *Payload) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. Description of the attack simulation and training campaign payload.
func (m *Payload) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDetail sets the detail property value. Additional details about the payload.
func (m *Payload) SetDetail(value PayloadDetailable)() {
    err := m.GetBackingStore().Set("detail", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. Display name of the attack simulation and training campaign payload. Supports $filter and $orderby.
func (m *Payload) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetIndustry sets the industry property value. Industry of a payload. Possible values are: unknown, other, banking, businessServices, consumerServices, education, energy, construction, consulting, financialServices, government, hospitality, insurance, legal, courierServices, IT, healthcare, manufacturing, retail, telecom, realEstate, unknownFutureValue.
func (m *Payload) SetIndustry(value *Payload_industry)() {
    err := m.GetBackingStore().Set("industry", value)
    if err != nil {
        panic(err)
    }
}
// SetIsAutomated sets the isAutomated property value. Indicates whether the attack simulation and training campaign payload was created from an automation flow. Supports $filter and $orderby.
func (m *Payload) SetIsAutomated(value *bool)() {
    err := m.GetBackingStore().Set("isAutomated", value)
    if err != nil {
        panic(err)
    }
}
// SetIsControversial sets the isControversial property value. Indicates whether the payload is controversial.
func (m *Payload) SetIsControversial(value *bool)() {
    err := m.GetBackingStore().Set("isControversial", value)
    if err != nil {
        panic(err)
    }
}
// SetIsCurrentEvent sets the isCurrentEvent property value. Indicates whether the payload is from any recent event.
func (m *Payload) SetIsCurrentEvent(value *bool)() {
    err := m.GetBackingStore().Set("isCurrentEvent", value)
    if err != nil {
        panic(err)
    }
}
// SetLanguage sets the language property value. Payload language.
func (m *Payload) SetLanguage(value *string)() {
    err := m.GetBackingStore().Set("language", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedBy sets the lastModifiedBy property value. Identity of the user who most recently modified the attack simulation and training campaign payload.
func (m *Payload) SetLastModifiedBy(value EmailIdentityable)() {
    err := m.GetBackingStore().Set("lastModifiedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Date and time when the attack simulation and training campaign payload was last modified. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *Payload) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetPayloadTags sets the payloadTags property value. Free text tags for a payload.
func (m *Payload) SetPayloadTags(value []string)() {
    err := m.GetBackingStore().Set("payloadTags", value)
    if err != nil {
        panic(err)
    }
}
// SetPlatform sets the platform property value. The payload delivery platform for a simulation. Possible values are: unknown, sms, email, teams, unknownFutureValue.
func (m *Payload) SetPlatform(value *Payload_platform)() {
    err := m.GetBackingStore().Set("platform", value)
    if err != nil {
        panic(err)
    }
}
// SetPredictedCompromiseRate sets the predictedCompromiseRate property value. Predicted probability for a payload to phish a targeted user.
func (m *Payload) SetPredictedCompromiseRate(value *float64)() {
    err := m.GetBackingStore().Set("predictedCompromiseRate", value)
    if err != nil {
        panic(err)
    }
}
// SetSimulationAttackType sets the simulationAttackType property value. Attack type of the attack simulation and training campaign. Supports $filter and $orderby. Possible values are: unknown, social, cloud, endpoint, unknownFutureValue.
func (m *Payload) SetSimulationAttackType(value *Payload_simulationAttackType)() {
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
// SetStatus sets the status property value. Simulation content status. Supports $filter and $orderby. Possible values are: unknown, draft, ready, archive, delete, unknownFutureValue.
func (m *Payload) SetStatus(value *Payload_status)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetTechnique sets the technique property value. The social engineering technique used in the attack simulation and training campaign. Supports $filter and $orderby. Possible values are: unknown, credentialHarvesting, attachmentMalware, driveByUrl, linkInAttachment, linkToMalwareFile, unknownFutureValue, oAuthConsentGrant. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values from this evolvable enum: oAuthConsentGrant. For more information on the types of social engineering attack techniques, see simulations.
func (m *Payload) SetTechnique(value *Payload_technique)() {
    err := m.GetBackingStore().Set("technique", value)
    if err != nil {
        panic(err)
    }
}
// SetTheme sets the theme property value. The theme of a payload. Possible values are: unknown, other, accountActivation, accountVerification, billing, cleanUpMail, controversial, documentReceived, expense, fax, financeReport, incomingMessages, invoice, itemReceived, loginAlert, mailReceived, password, payment, payroll, personalizedOffer, quarantine, remoteWork, reviewMessage, securityUpdate, serviceSuspended, signatureRequired, upgradeMailboxStorage, verifyMailbox, voicemail, advertisement, employeeEngagement, unknownFutureValue.
func (m *Payload) SetTheme(value *Payload_theme)() {
    err := m.GetBackingStore().Set("theme", value)
    if err != nil {
        panic(err)
    }
}
// Payloadable 
type Payloadable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBrand()(*Payload_brand)
    GetComplexity()(*Payload_complexity)
    GetCreatedBy()(EmailIdentityable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDetail()(PayloadDetailable)
    GetDisplayName()(*string)
    GetIndustry()(*Payload_industry)
    GetIsAutomated()(*bool)
    GetIsControversial()(*bool)
    GetIsCurrentEvent()(*bool)
    GetLanguage()(*string)
    GetLastModifiedBy()(EmailIdentityable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPayloadTags()([]string)
    GetPlatform()(*Payload_platform)
    GetPredictedCompromiseRate()(*float64)
    GetSimulationAttackType()(*Payload_simulationAttackType)
    GetSource()(*SimulationContentSource)
    GetStatus()(*Payload_status)
    GetTechnique()(*Payload_technique)
    GetTheme()(*Payload_theme)
    SetBrand(value *Payload_brand)()
    SetComplexity(value *Payload_complexity)()
    SetCreatedBy(value EmailIdentityable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDetail(value PayloadDetailable)()
    SetDisplayName(value *string)()
    SetIndustry(value *Payload_industry)()
    SetIsAutomated(value *bool)()
    SetIsControversial(value *bool)()
    SetIsCurrentEvent(value *bool)()
    SetLanguage(value *string)()
    SetLastModifiedBy(value EmailIdentityable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPayloadTags(value []string)()
    SetPlatform(value *Payload_platform)()
    SetPredictedCompromiseRate(value *float64)()
    SetSimulationAttackType(value *Payload_simulationAttackType)()
    SetSource(value *SimulationContentSource)()
    SetStatus(value *Payload_status)()
    SetTechnique(value *Payload_technique)()
    SetTheme(value *Payload_theme)()
}
