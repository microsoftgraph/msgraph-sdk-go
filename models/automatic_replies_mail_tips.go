package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AutomaticRepliesMailTips 
type AutomaticRepliesMailTips struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The automatic reply message.
    message *string
    // The language that the automatic reply message is in.
    messageLanguage LocaleInfoable
    // The OdataType property
    odataType *string
    // The date and time that automatic replies are set to end.
    scheduledEndTime DateTimeTimeZoneable
    // The date and time that automatic replies are set to begin.
    scheduledStartTime DateTimeTimeZoneable
}
// NewAutomaticRepliesMailTips instantiates a new automaticRepliesMailTips and sets the default values.
func NewAutomaticRepliesMailTips()(*AutomaticRepliesMailTips) {
    m := &AutomaticRepliesMailTips{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.automaticRepliesMailTips";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAutomaticRepliesMailTipsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAutomaticRepliesMailTipsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAutomaticRepliesMailTips(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AutomaticRepliesMailTips) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AutomaticRepliesMailTips) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["messageLanguage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLocaleInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessageLanguage(val.(LocaleInfoable))
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
    res["scheduledEndTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduledEndTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["scheduledStartTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduledStartTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. The automatic reply message.
func (m *AutomaticRepliesMailTips) GetMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.message
    }
}
// GetMessageLanguage gets the messageLanguage property value. The language that the automatic reply message is in.
func (m *AutomaticRepliesMailTips) GetMessageLanguage()(LocaleInfoable) {
    if m == nil {
        return nil
    } else {
        return m.messageLanguage
    }
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AutomaticRepliesMailTips) GetOdataType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.odataType
    }
}
// GetScheduledEndTime gets the scheduledEndTime property value. The date and time that automatic replies are set to end.
func (m *AutomaticRepliesMailTips) GetScheduledEndTime()(DateTimeTimeZoneable) {
    if m == nil {
        return nil
    } else {
        return m.scheduledEndTime
    }
}
// GetScheduledStartTime gets the scheduledStartTime property value. The date and time that automatic replies are set to begin.
func (m *AutomaticRepliesMailTips) GetScheduledStartTime()(DateTimeTimeZoneable) {
    if m == nil {
        return nil
    } else {
        return m.scheduledStartTime
    }
}
// Serialize serializes information the current object
func (m *AutomaticRepliesMailTips) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("messageLanguage", m.GetMessageLanguage())
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
        err := writer.WriteObjectValue("scheduledEndTime", m.GetScheduledEndTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("scheduledStartTime", m.GetScheduledStartTime())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AutomaticRepliesMailTips) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMessage sets the message property value. The automatic reply message.
func (m *AutomaticRepliesMailTips) SetMessage(value *string)() {
    if m != nil {
        m.message = value
    }
}
// SetMessageLanguage sets the messageLanguage property value. The language that the automatic reply message is in.
func (m *AutomaticRepliesMailTips) SetMessageLanguage(value LocaleInfoable)() {
    if m != nil {
        m.messageLanguage = value
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AutomaticRepliesMailTips) SetOdataType(value *string)() {
    if m != nil {
        m.odataType = value
    }
}
// SetScheduledEndTime sets the scheduledEndTime property value. The date and time that automatic replies are set to end.
func (m *AutomaticRepliesMailTips) SetScheduledEndTime(value DateTimeTimeZoneable)() {
    if m != nil {
        m.scheduledEndTime = value
    }
}
// SetScheduledStartTime sets the scheduledStartTime property value. The date and time that automatic replies are set to begin.
func (m *AutomaticRepliesMailTips) SetScheduledStartTime(value DateTimeTimeZoneable)() {
    if m != nil {
        m.scheduledStartTime = value
    }
}
