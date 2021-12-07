package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UriClickSecurityState 
type UriClickSecurityState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    clickAction *string;
    // 
    clickDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    id *string;
    // 
    sourceId *string;
    // 
    uriDomain *string;
    // 
    verdict *string;
}
// NewUriClickSecurityState instantiates a new uriClickSecurityState and sets the default values.
func NewUriClickSecurityState()(*UriClickSecurityState) {
    m := &UriClickSecurityState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UriClickSecurityState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetClickAction gets the clickAction property value. 
func (m *UriClickSecurityState) GetClickAction()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clickAction
    }
}
// GetClickDateTime gets the clickDateTime property value. 
func (m *UriClickSecurityState) GetClickDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.clickDateTime
    }
}
// GetId gets the id property value. 
func (m *UriClickSecurityState) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetSourceId gets the sourceId property value. 
func (m *UriClickSecurityState) GetSourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceId
    }
}
// GetUriDomain gets the uriDomain property value. 
func (m *UriClickSecurityState) GetUriDomain()(*string) {
    if m == nil {
        return nil
    } else {
        return m.uriDomain
    }
}
// GetVerdict gets the verdict property value. 
func (m *UriClickSecurityState) GetVerdict()(*string) {
    if m == nil {
        return nil
    } else {
        return m.verdict
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UriClickSecurityState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["clickAction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClickAction(val)
        }
        return nil
    }
    res["clickDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClickDateTime(val)
        }
        return nil
    }
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["sourceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceId(val)
        }
        return nil
    }
    res["uriDomain"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUriDomain(val)
        }
        return nil
    }
    res["verdict"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerdict(val)
        }
        return nil
    }
    return res
}
func (m *UriClickSecurityState) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UriClickSecurityState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("clickAction", m.GetClickAction())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("clickDateTime", m.GetClickDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sourceId", m.GetSourceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("uriDomain", m.GetUriDomain())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("verdict", m.GetVerdict())
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
func (m *UriClickSecurityState) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetClickAction sets the clickAction property value. 
func (m *UriClickSecurityState) SetClickAction(value *string)() {
    if m != nil {
        m.clickAction = value
    }
}
// SetClickDateTime sets the clickDateTime property value. 
func (m *UriClickSecurityState) SetClickDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.clickDateTime = value
    }
}
// SetId sets the id property value. 
func (m *UriClickSecurityState) SetId(value *string)() {
    if m != nil {
        m.id = value
    }
}
// SetSourceId sets the sourceId property value. 
func (m *UriClickSecurityState) SetSourceId(value *string)() {
    if m != nil {
        m.sourceId = value
    }
}
// SetUriDomain sets the uriDomain property value. 
func (m *UriClickSecurityState) SetUriDomain(value *string)() {
    if m != nil {
        m.uriDomain = value
    }
}
// SetVerdict sets the verdict property value. 
func (m *UriClickSecurityState) SetVerdict(value *string)() {
    if m != nil {
        m.verdict = value
    }
}
