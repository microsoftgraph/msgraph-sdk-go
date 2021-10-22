package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ApplicationServicePrincipal struct {
    additionalData map[string]interface{};
    application *Application;
    servicePrincipal *ServicePrincipal;
}
func NewApplicationServicePrincipal()(*ApplicationServicePrincipal) {
    m := &ApplicationServicePrincipal{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ApplicationServicePrincipal) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ApplicationServicePrincipal) GetApplication()(*Application) {
    if m == nil {
        return nil
    } else {
        return m.application
    }
}
func (m *ApplicationServicePrincipal) GetServicePrincipal()(*ServicePrincipal) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipal
    }
}
func (m *ApplicationServicePrincipal) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["application"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewApplication() })
        if err != nil {
            return err
        }
        m.SetApplication(val.(*Application))
        return nil
    }
    res["servicePrincipal"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewServicePrincipal() })
        if err != nil {
            return err
        }
        m.SetServicePrincipal(val.(*ServicePrincipal))
        return nil
    }
    return res
}
func (m *ApplicationServicePrincipal) IsNil()(bool) {
    return m == nil
}
func (m *ApplicationServicePrincipal) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("application", m.GetApplication())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("servicePrincipal", m.GetServicePrincipal())
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
func (m *ApplicationServicePrincipal) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ApplicationServicePrincipal) SetApplication(value *Application)() {
    m.application = value
}
func (m *ApplicationServicePrincipal) SetServicePrincipal(value *ServicePrincipal)() {
    m.servicePrincipal = value
}
