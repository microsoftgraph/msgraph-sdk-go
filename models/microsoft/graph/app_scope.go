package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AppScope 
type AppScope struct {
    Entity
    // Provides the display name of the app-specific resource represented by the app scope. Provided for display purposes since appScopeId is often an immutable, non-human-readable id. Read-only.
    displayName *string;
    // Describes the type of app-specific resource represented by the app scope. Provided for display purposes, so a user interface can convey to the user the kind of app specific resource represented by the app scope. Read-only.
    type_escaped *string;
}
// NewAppScope instantiates a new appScope and sets the default values.
func NewAppScope()(*AppScope) {
    m := &AppScope{
        Entity: *NewEntity(),
    }
    return m
}
// GetDisplayName gets the displayName property value. Provides the display name of the app-specific resource represented by the app scope. Provided for display purposes since appScopeId is often an immutable, non-human-readable id. Read-only.
func (m *AppScope) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetType_escaped gets the type_escaped property value. Describes the type of app-specific resource represented by the app scope. Provided for display purposes, so a user interface can convey to the user the kind of app specific resource represented by the app scope. Read-only.
func (m *AppScope) GetType_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppScope) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType_escaped(val)
        }
        return nil
    }
    return res
}
func (m *AppScope) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AppScope) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("type_escaped", m.GetType_escaped())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. Provides the display name of the app-specific resource represented by the app scope. Provided for display purposes since appScopeId is often an immutable, non-human-readable id. Read-only.
func (m *AppScope) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetType_escaped sets the type_escaped property value. Describes the type of app-specific resource represented by the app scope. Provided for display purposes, so a user interface can convey to the user the kind of app specific resource represented by the app scope. Read-only.
func (m *AppScope) SetType_escaped(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
