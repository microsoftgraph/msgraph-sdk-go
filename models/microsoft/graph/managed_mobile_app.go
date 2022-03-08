package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagedMobileApp provides operations to manage the deviceAppManagement singleton.
type ManagedMobileApp struct {
    Entity
    // The identifier for an app with it's operating system type.
    mobileAppIdentifier MobileAppIdentifierable;
    // Version of the entity.
    version *string;
}
// NewManagedMobileApp instantiates a new managedMobileApp and sets the default values.
func NewManagedMobileApp()(*ManagedMobileApp) {
    m := &ManagedMobileApp{
        Entity: *NewEntity(),
    }
    return m
}
// CreateManagedMobileAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedMobileAppFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewManagedMobileApp(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedMobileApp) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["mobileAppIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateMobileAppIdentifierFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMobileAppIdentifier(val.(MobileAppIdentifierable))
        }
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetMobileAppIdentifier gets the mobileAppIdentifier property value. The identifier for an app with it's operating system type.
func (m *ManagedMobileApp) GetMobileAppIdentifier()(MobileAppIdentifierable) {
    if m == nil {
        return nil
    } else {
        return m.mobileAppIdentifier
    }
}
// GetVersion gets the version property value. Version of the entity.
func (m *ManagedMobileApp) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
func (m *ManagedMobileApp) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ManagedMobileApp) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("mobileAppIdentifier", m.GetMobileAppIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMobileAppIdentifier sets the mobileAppIdentifier property value. The identifier for an app with it's operating system type.
func (m *ManagedMobileApp) SetMobileAppIdentifier(value MobileAppIdentifierable)() {
    if m != nil {
        m.mobileAppIdentifier = value
    }
}
// SetVersion sets the version property value. Version of the entity.
func (m *ManagedMobileApp) SetVersion(value *string)() {
    if m != nil {
        m.version = value
    }
}
