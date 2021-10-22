package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AuditActivityInitiator struct {
    additionalData map[string]interface{};
    app *AppIdentity;
    user *UserIdentity;
}
func NewAuditActivityInitiator()(*AuditActivityInitiator) {
    m := &AuditActivityInitiator{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AuditActivityInitiator) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AuditActivityInitiator) GetApp()(*AppIdentity) {
    if m == nil {
        return nil
    } else {
        return m.app
    }
}
func (m *AuditActivityInitiator) GetUser()(*UserIdentity) {
    if m == nil {
        return nil
    } else {
        return m.user
    }
}
func (m *AuditActivityInitiator) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["app"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAppIdentity() })
        if err != nil {
            return err
        }
        m.SetApp(val.(*AppIdentity))
        return nil
    }
    res["user"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserIdentity() })
        if err != nil {
            return err
        }
        m.SetUser(val.(*UserIdentity))
        return nil
    }
    return res
}
func (m *AuditActivityInitiator) IsNil()(bool) {
    return m == nil
}
func (m *AuditActivityInitiator) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("app", m.GetApp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("user", m.GetUser())
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
func (m *AuditActivityInitiator) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AuditActivityInitiator) SetApp(value *AppIdentity)() {
    m.app = value
}
func (m *AuditActivityInitiator) SetUser(value *UserIdentity)() {
    m.user = value
}
