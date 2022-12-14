package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AnalyzedMessageEvidence 
type AnalyzedMessageEvidence struct {
    AlertEvidence
    // The antiSpamDirection property
    antiSpamDirection *string
    // The attachmentsCount property
    attachmentsCount *int64
    // The deliveryAction property
    deliveryAction *string
    // The deliveryLocation property
    deliveryLocation *string
    // The internetMessageId property
    internetMessageId *string
    // The language property
    language *string
    // The networkMessageId property
    networkMessageId *string
    // The p1Sender property
    p1Sender EmailSenderable
    // The p2Sender property
    p2Sender EmailSenderable
    // The receivedDateTime property
    receivedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The recipientEmailAddress property
    recipientEmailAddress *string
    // The senderIp property
    senderIp *string
    // The subject property
    subject *string
    // The threatDetectionMethods property
    threatDetectionMethods []string
    // The threats property
    threats []string
    // The urlCount property
    urlCount *int64
    // The urls property
    urls []string
    // The urn property
    urn *string
}
// NewAnalyzedMessageEvidence instantiates a new AnalyzedMessageEvidence and sets the default values.
func NewAnalyzedMessageEvidence()(*AnalyzedMessageEvidence) {
    m := &AnalyzedMessageEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    return m
}
// CreateAnalyzedMessageEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAnalyzedMessageEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAnalyzedMessageEvidence(), nil
}
// GetAntiSpamDirection gets the antiSpamDirection property value. The antiSpamDirection property
func (m *AnalyzedMessageEvidence) GetAntiSpamDirection()(*string) {
    return m.antiSpamDirection
}
// GetAttachmentsCount gets the attachmentsCount property value. The attachmentsCount property
func (m *AnalyzedMessageEvidence) GetAttachmentsCount()(*int64) {
    return m.attachmentsCount
}
// GetDeliveryAction gets the deliveryAction property value. The deliveryAction property
func (m *AnalyzedMessageEvidence) GetDeliveryAction()(*string) {
    return m.deliveryAction
}
// GetDeliveryLocation gets the deliveryLocation property value. The deliveryLocation property
func (m *AnalyzedMessageEvidence) GetDeliveryLocation()(*string) {
    return m.deliveryLocation
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AnalyzedMessageEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["antiSpamDirection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAntiSpamDirection(val)
        }
        return nil
    }
    res["attachmentsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttachmentsCount(val)
        }
        return nil
    }
    res["deliveryAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeliveryAction(val)
        }
        return nil
    }
    res["deliveryLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeliveryLocation(val)
        }
        return nil
    }
    res["internetMessageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternetMessageId(val)
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
    res["networkMessageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkMessageId(val)
        }
        return nil
    }
    res["p1Sender"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEmailSenderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetP1Sender(val.(EmailSenderable))
        }
        return nil
    }
    res["p2Sender"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEmailSenderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetP2Sender(val.(EmailSenderable))
        }
        return nil
    }
    res["receivedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReceivedDateTime(val)
        }
        return nil
    }
    res["recipientEmailAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecipientEmailAddress(val)
        }
        return nil
    }
    res["senderIp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSenderIp(val)
        }
        return nil
    }
    res["subject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val)
        }
        return nil
    }
    res["threatDetectionMethods"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetThreatDetectionMethods(res)
        }
        return nil
    }
    res["threats"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetThreats(res)
        }
        return nil
    }
    res["urlCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrlCount(val)
        }
        return nil
    }
    res["urls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetUrls(res)
        }
        return nil
    }
    res["urn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrn(val)
        }
        return nil
    }
    return res
}
// GetInternetMessageId gets the internetMessageId property value. The internetMessageId property
func (m *AnalyzedMessageEvidence) GetInternetMessageId()(*string) {
    return m.internetMessageId
}
// GetLanguage gets the language property value. The language property
func (m *AnalyzedMessageEvidence) GetLanguage()(*string) {
    return m.language
}
// GetNetworkMessageId gets the networkMessageId property value. The networkMessageId property
func (m *AnalyzedMessageEvidence) GetNetworkMessageId()(*string) {
    return m.networkMessageId
}
// GetP1Sender gets the p1Sender property value. The p1Sender property
func (m *AnalyzedMessageEvidence) GetP1Sender()(EmailSenderable) {
    return m.p1Sender
}
// GetP2Sender gets the p2Sender property value. The p2Sender property
func (m *AnalyzedMessageEvidence) GetP2Sender()(EmailSenderable) {
    return m.p2Sender
}
// GetReceivedDateTime gets the receivedDateTime property value. The receivedDateTime property
func (m *AnalyzedMessageEvidence) GetReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.receivedDateTime
}
// GetRecipientEmailAddress gets the recipientEmailAddress property value. The recipientEmailAddress property
func (m *AnalyzedMessageEvidence) GetRecipientEmailAddress()(*string) {
    return m.recipientEmailAddress
}
// GetSenderIp gets the senderIp property value. The senderIp property
func (m *AnalyzedMessageEvidence) GetSenderIp()(*string) {
    return m.senderIp
}
// GetSubject gets the subject property value. The subject property
func (m *AnalyzedMessageEvidence) GetSubject()(*string) {
    return m.subject
}
// GetThreatDetectionMethods gets the threatDetectionMethods property value. The threatDetectionMethods property
func (m *AnalyzedMessageEvidence) GetThreatDetectionMethods()([]string) {
    return m.threatDetectionMethods
}
// GetThreats gets the threats property value. The threats property
func (m *AnalyzedMessageEvidence) GetThreats()([]string) {
    return m.threats
}
// GetUrlCount gets the urlCount property value. The urlCount property
func (m *AnalyzedMessageEvidence) GetUrlCount()(*int64) {
    return m.urlCount
}
// GetUrls gets the urls property value. The urls property
func (m *AnalyzedMessageEvidence) GetUrls()([]string) {
    return m.urls
}
// GetUrn gets the urn property value. The urn property
func (m *AnalyzedMessageEvidence) GetUrn()(*string) {
    return m.urn
}
// Serialize serializes information the current object
func (m *AnalyzedMessageEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("antiSpamDirection", m.GetAntiSpamDirection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("attachmentsCount", m.GetAttachmentsCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deliveryAction", m.GetDeliveryAction())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deliveryLocation", m.GetDeliveryLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("internetMessageId", m.GetInternetMessageId())
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
        err = writer.WriteStringValue("networkMessageId", m.GetNetworkMessageId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("p1Sender", m.GetP1Sender())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("p2Sender", m.GetP2Sender())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("receivedDateTime", m.GetReceivedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("recipientEmailAddress", m.GetRecipientEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("senderIp", m.GetSenderIp())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    if m.GetThreatDetectionMethods() != nil {
        err = writer.WriteCollectionOfStringValues("threatDetectionMethods", m.GetThreatDetectionMethods())
        if err != nil {
            return err
        }
    }
    if m.GetThreats() != nil {
        err = writer.WriteCollectionOfStringValues("threats", m.GetThreats())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("urlCount", m.GetUrlCount())
        if err != nil {
            return err
        }
    }
    if m.GetUrls() != nil {
        err = writer.WriteCollectionOfStringValues("urls", m.GetUrls())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("urn", m.GetUrn())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAntiSpamDirection sets the antiSpamDirection property value. The antiSpamDirection property
func (m *AnalyzedMessageEvidence) SetAntiSpamDirection(value *string)() {
    m.antiSpamDirection = value
}
// SetAttachmentsCount sets the attachmentsCount property value. The attachmentsCount property
func (m *AnalyzedMessageEvidence) SetAttachmentsCount(value *int64)() {
    m.attachmentsCount = value
}
// SetDeliveryAction sets the deliveryAction property value. The deliveryAction property
func (m *AnalyzedMessageEvidence) SetDeliveryAction(value *string)() {
    m.deliveryAction = value
}
// SetDeliveryLocation sets the deliveryLocation property value. The deliveryLocation property
func (m *AnalyzedMessageEvidence) SetDeliveryLocation(value *string)() {
    m.deliveryLocation = value
}
// SetInternetMessageId sets the internetMessageId property value. The internetMessageId property
func (m *AnalyzedMessageEvidence) SetInternetMessageId(value *string)() {
    m.internetMessageId = value
}
// SetLanguage sets the language property value. The language property
func (m *AnalyzedMessageEvidence) SetLanguage(value *string)() {
    m.language = value
}
// SetNetworkMessageId sets the networkMessageId property value. The networkMessageId property
func (m *AnalyzedMessageEvidence) SetNetworkMessageId(value *string)() {
    m.networkMessageId = value
}
// SetP1Sender sets the p1Sender property value. The p1Sender property
func (m *AnalyzedMessageEvidence) SetP1Sender(value EmailSenderable)() {
    m.p1Sender = value
}
// SetP2Sender sets the p2Sender property value. The p2Sender property
func (m *AnalyzedMessageEvidence) SetP2Sender(value EmailSenderable)() {
    m.p2Sender = value
}
// SetReceivedDateTime sets the receivedDateTime property value. The receivedDateTime property
func (m *AnalyzedMessageEvidence) SetReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.receivedDateTime = value
}
// SetRecipientEmailAddress sets the recipientEmailAddress property value. The recipientEmailAddress property
func (m *AnalyzedMessageEvidence) SetRecipientEmailAddress(value *string)() {
    m.recipientEmailAddress = value
}
// SetSenderIp sets the senderIp property value. The senderIp property
func (m *AnalyzedMessageEvidence) SetSenderIp(value *string)() {
    m.senderIp = value
}
// SetSubject sets the subject property value. The subject property
func (m *AnalyzedMessageEvidence) SetSubject(value *string)() {
    m.subject = value
}
// SetThreatDetectionMethods sets the threatDetectionMethods property value. The threatDetectionMethods property
func (m *AnalyzedMessageEvidence) SetThreatDetectionMethods(value []string)() {
    m.threatDetectionMethods = value
}
// SetThreats sets the threats property value. The threats property
func (m *AnalyzedMessageEvidence) SetThreats(value []string)() {
    m.threats = value
}
// SetUrlCount sets the urlCount property value. The urlCount property
func (m *AnalyzedMessageEvidence) SetUrlCount(value *int64)() {
    m.urlCount = value
}
// SetUrls sets the urls property value. The urls property
func (m *AnalyzedMessageEvidence) SetUrls(value []string)() {
    m.urls = value
}
// SetUrn sets the urn property value. The urn property
func (m *AnalyzedMessageEvidence) SetUrn(value *string)() {
    m.urn = value
}
