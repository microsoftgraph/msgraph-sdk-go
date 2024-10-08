package callrecords

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type MediaStream struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// MediaStream_MediaStream_averageAudioDegradation composed type wrapper for classes float32, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric, string
type MediaStream_MediaStream_averageAudioDegradation struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewMediaStream_MediaStream_averageAudioDegradation instantiates a new MediaStream_MediaStream_averageAudioDegradation and sets the default values.
func NewMediaStream_MediaStream_averageAudioDegradation()(*MediaStream_MediaStream_averageAudioDegradation) {
    m := &MediaStream_MediaStream_averageAudioDegradation{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateMediaStream_MediaStream_averageAudioDegradationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMediaStream_MediaStream_averageAudioDegradationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewMediaStream_MediaStream_averageAudioDegradation()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseReferenceNumeric); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetReferenceNumeric(val)
    } else if val, err := parseNode.GetFloat32Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetFloat(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *MediaStream_MediaStream_averageAudioDegradation) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MediaStream_MediaStream_averageAudioDegradation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetFloat gets the float property value. Composed type representation for type float32
// returns a *float32 when successful
func (m *MediaStream_MediaStream_averageAudioDegradation) GetFloat()(*float32) {
    val, err := m.GetBackingStore().Get("float")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float32)
    }
    return nil
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *MediaStream_MediaStream_averageAudioDegradation) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *MediaStream_MediaStream_averageAudioDegradation) GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *MediaStream_MediaStream_averageAudioDegradation) GetString()(*string) {
    val, err := m.GetBackingStore().Get("string")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MediaStream_MediaStream_averageAudioDegradation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetFloat() != nil {
        err := writer.WriteFloat32Value("", m.GetFloat())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *MediaStream_MediaStream_averageAudioDegradation) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFloat sets the float property value. Composed type representation for type float32
func (m *MediaStream_MediaStream_averageAudioDegradation) SetFloat(value *float32)() {
    err := m.GetBackingStore().Set("float", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
func (m *MediaStream_MediaStream_averageAudioDegradation) SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *MediaStream_MediaStream_averageAudioDegradation) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// MediaStream_MediaStream_averagePacketLossRate composed type wrapper for classes float32, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric, string
type MediaStream_MediaStream_averagePacketLossRate struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewMediaStream_MediaStream_averagePacketLossRate instantiates a new MediaStream_MediaStream_averagePacketLossRate and sets the default values.
func NewMediaStream_MediaStream_averagePacketLossRate()(*MediaStream_MediaStream_averagePacketLossRate) {
    m := &MediaStream_MediaStream_averagePacketLossRate{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateMediaStream_MediaStream_averagePacketLossRateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMediaStream_MediaStream_averagePacketLossRateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewMediaStream_MediaStream_averagePacketLossRate()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseReferenceNumeric); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetReferenceNumeric(val)
    } else if val, err := parseNode.GetFloat32Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetFloat(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *MediaStream_MediaStream_averagePacketLossRate) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MediaStream_MediaStream_averagePacketLossRate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetFloat gets the float property value. Composed type representation for type float32
// returns a *float32 when successful
func (m *MediaStream_MediaStream_averagePacketLossRate) GetFloat()(*float32) {
    val, err := m.GetBackingStore().Get("float")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float32)
    }
    return nil
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *MediaStream_MediaStream_averagePacketLossRate) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *MediaStream_MediaStream_averagePacketLossRate) GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *MediaStream_MediaStream_averagePacketLossRate) GetString()(*string) {
    val, err := m.GetBackingStore().Get("string")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MediaStream_MediaStream_averagePacketLossRate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetFloat() != nil {
        err := writer.WriteFloat32Value("", m.GetFloat())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *MediaStream_MediaStream_averagePacketLossRate) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFloat sets the float property value. Composed type representation for type float32
func (m *MediaStream_MediaStream_averagePacketLossRate) SetFloat(value *float32)() {
    err := m.GetBackingStore().Set("float", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
func (m *MediaStream_MediaStream_averagePacketLossRate) SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *MediaStream_MediaStream_averagePacketLossRate) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// MediaStream_MediaStream_averageRatioOfConcealedSamples composed type wrapper for classes float32, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric, string
type MediaStream_MediaStream_averageRatioOfConcealedSamples struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewMediaStream_MediaStream_averageRatioOfConcealedSamples instantiates a new MediaStream_MediaStream_averageRatioOfConcealedSamples and sets the default values.
func NewMediaStream_MediaStream_averageRatioOfConcealedSamples()(*MediaStream_MediaStream_averageRatioOfConcealedSamples) {
    m := &MediaStream_MediaStream_averageRatioOfConcealedSamples{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateMediaStream_MediaStream_averageRatioOfConcealedSamplesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMediaStream_MediaStream_averageRatioOfConcealedSamplesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewMediaStream_MediaStream_averageRatioOfConcealedSamples()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseReferenceNumeric); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetReferenceNumeric(val)
    } else if val, err := parseNode.GetFloat32Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetFloat(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *MediaStream_MediaStream_averageRatioOfConcealedSamples) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MediaStream_MediaStream_averageRatioOfConcealedSamples) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetFloat gets the float property value. Composed type representation for type float32
// returns a *float32 when successful
func (m *MediaStream_MediaStream_averageRatioOfConcealedSamples) GetFloat()(*float32) {
    val, err := m.GetBackingStore().Get("float")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float32)
    }
    return nil
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *MediaStream_MediaStream_averageRatioOfConcealedSamples) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *MediaStream_MediaStream_averageRatioOfConcealedSamples) GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *MediaStream_MediaStream_averageRatioOfConcealedSamples) GetString()(*string) {
    val, err := m.GetBackingStore().Get("string")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MediaStream_MediaStream_averageRatioOfConcealedSamples) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetFloat() != nil {
        err := writer.WriteFloat32Value("", m.GetFloat())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *MediaStream_MediaStream_averageRatioOfConcealedSamples) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFloat sets the float property value. Composed type representation for type float32
func (m *MediaStream_MediaStream_averageRatioOfConcealedSamples) SetFloat(value *float32)() {
    err := m.GetBackingStore().Set("float", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
func (m *MediaStream_MediaStream_averageRatioOfConcealedSamples) SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *MediaStream_MediaStream_averageRatioOfConcealedSamples) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// MediaStream_MediaStream_averageReceivedFrameRate composed type wrapper for classes float32, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric, string
type MediaStream_MediaStream_averageReceivedFrameRate struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewMediaStream_MediaStream_averageReceivedFrameRate instantiates a new MediaStream_MediaStream_averageReceivedFrameRate and sets the default values.
func NewMediaStream_MediaStream_averageReceivedFrameRate()(*MediaStream_MediaStream_averageReceivedFrameRate) {
    m := &MediaStream_MediaStream_averageReceivedFrameRate{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateMediaStream_MediaStream_averageReceivedFrameRateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMediaStream_MediaStream_averageReceivedFrameRateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewMediaStream_MediaStream_averageReceivedFrameRate()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseReferenceNumeric); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetReferenceNumeric(val)
    } else if val, err := parseNode.GetFloat32Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetFloat(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *MediaStream_MediaStream_averageReceivedFrameRate) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MediaStream_MediaStream_averageReceivedFrameRate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetFloat gets the float property value. Composed type representation for type float32
// returns a *float32 when successful
func (m *MediaStream_MediaStream_averageReceivedFrameRate) GetFloat()(*float32) {
    val, err := m.GetBackingStore().Get("float")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float32)
    }
    return nil
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *MediaStream_MediaStream_averageReceivedFrameRate) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *MediaStream_MediaStream_averageReceivedFrameRate) GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *MediaStream_MediaStream_averageReceivedFrameRate) GetString()(*string) {
    val, err := m.GetBackingStore().Get("string")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MediaStream_MediaStream_averageReceivedFrameRate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetFloat() != nil {
        err := writer.WriteFloat32Value("", m.GetFloat())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *MediaStream_MediaStream_averageReceivedFrameRate) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFloat sets the float property value. Composed type representation for type float32
func (m *MediaStream_MediaStream_averageReceivedFrameRate) SetFloat(value *float32)() {
    err := m.GetBackingStore().Set("float", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
func (m *MediaStream_MediaStream_averageReceivedFrameRate) SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *MediaStream_MediaStream_averageReceivedFrameRate) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// MediaStream_MediaStream_averageVideoFrameLossPercentage composed type wrapper for classes float32, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric, string
type MediaStream_MediaStream_averageVideoFrameLossPercentage struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewMediaStream_MediaStream_averageVideoFrameLossPercentage instantiates a new MediaStream_MediaStream_averageVideoFrameLossPercentage and sets the default values.
func NewMediaStream_MediaStream_averageVideoFrameLossPercentage()(*MediaStream_MediaStream_averageVideoFrameLossPercentage) {
    m := &MediaStream_MediaStream_averageVideoFrameLossPercentage{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateMediaStream_MediaStream_averageVideoFrameLossPercentageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMediaStream_MediaStream_averageVideoFrameLossPercentageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewMediaStream_MediaStream_averageVideoFrameLossPercentage()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseReferenceNumeric); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetReferenceNumeric(val)
    } else if val, err := parseNode.GetFloat32Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetFloat(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *MediaStream_MediaStream_averageVideoFrameLossPercentage) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MediaStream_MediaStream_averageVideoFrameLossPercentage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetFloat gets the float property value. Composed type representation for type float32
// returns a *float32 when successful
func (m *MediaStream_MediaStream_averageVideoFrameLossPercentage) GetFloat()(*float32) {
    val, err := m.GetBackingStore().Get("float")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float32)
    }
    return nil
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *MediaStream_MediaStream_averageVideoFrameLossPercentage) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *MediaStream_MediaStream_averageVideoFrameLossPercentage) GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *MediaStream_MediaStream_averageVideoFrameLossPercentage) GetString()(*string) {
    val, err := m.GetBackingStore().Get("string")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MediaStream_MediaStream_averageVideoFrameLossPercentage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetFloat() != nil {
        err := writer.WriteFloat32Value("", m.GetFloat())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *MediaStream_MediaStream_averageVideoFrameLossPercentage) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFloat sets the float property value. Composed type representation for type float32
func (m *MediaStream_MediaStream_averageVideoFrameLossPercentage) SetFloat(value *float32)() {
    err := m.GetBackingStore().Set("float", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
func (m *MediaStream_MediaStream_averageVideoFrameLossPercentage) SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *MediaStream_MediaStream_averageVideoFrameLossPercentage) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// MediaStream_MediaStream_averageVideoFrameRate composed type wrapper for classes float32, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric, string
type MediaStream_MediaStream_averageVideoFrameRate struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewMediaStream_MediaStream_averageVideoFrameRate instantiates a new MediaStream_MediaStream_averageVideoFrameRate and sets the default values.
func NewMediaStream_MediaStream_averageVideoFrameRate()(*MediaStream_MediaStream_averageVideoFrameRate) {
    m := &MediaStream_MediaStream_averageVideoFrameRate{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateMediaStream_MediaStream_averageVideoFrameRateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMediaStream_MediaStream_averageVideoFrameRateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewMediaStream_MediaStream_averageVideoFrameRate()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseReferenceNumeric); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetReferenceNumeric(val)
    } else if val, err := parseNode.GetFloat32Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetFloat(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *MediaStream_MediaStream_averageVideoFrameRate) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MediaStream_MediaStream_averageVideoFrameRate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetFloat gets the float property value. Composed type representation for type float32
// returns a *float32 when successful
func (m *MediaStream_MediaStream_averageVideoFrameRate) GetFloat()(*float32) {
    val, err := m.GetBackingStore().Get("float")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float32)
    }
    return nil
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *MediaStream_MediaStream_averageVideoFrameRate) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *MediaStream_MediaStream_averageVideoFrameRate) GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *MediaStream_MediaStream_averageVideoFrameRate) GetString()(*string) {
    val, err := m.GetBackingStore().Get("string")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MediaStream_MediaStream_averageVideoFrameRate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetFloat() != nil {
        err := writer.WriteFloat32Value("", m.GetFloat())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *MediaStream_MediaStream_averageVideoFrameRate) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFloat sets the float property value. Composed type representation for type float32
func (m *MediaStream_MediaStream_averageVideoFrameRate) SetFloat(value *float32)() {
    err := m.GetBackingStore().Set("float", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
func (m *MediaStream_MediaStream_averageVideoFrameRate) SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *MediaStream_MediaStream_averageVideoFrameRate) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// MediaStream_MediaStream_averageVideoPacketLossRate composed type wrapper for classes float32, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric, string
type MediaStream_MediaStream_averageVideoPacketLossRate struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewMediaStream_MediaStream_averageVideoPacketLossRate instantiates a new MediaStream_MediaStream_averageVideoPacketLossRate and sets the default values.
func NewMediaStream_MediaStream_averageVideoPacketLossRate()(*MediaStream_MediaStream_averageVideoPacketLossRate) {
    m := &MediaStream_MediaStream_averageVideoPacketLossRate{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateMediaStream_MediaStream_averageVideoPacketLossRateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMediaStream_MediaStream_averageVideoPacketLossRateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewMediaStream_MediaStream_averageVideoPacketLossRate()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseReferenceNumeric); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetReferenceNumeric(val)
    } else if val, err := parseNode.GetFloat32Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetFloat(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *MediaStream_MediaStream_averageVideoPacketLossRate) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MediaStream_MediaStream_averageVideoPacketLossRate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetFloat gets the float property value. Composed type representation for type float32
// returns a *float32 when successful
func (m *MediaStream_MediaStream_averageVideoPacketLossRate) GetFloat()(*float32) {
    val, err := m.GetBackingStore().Get("float")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float32)
    }
    return nil
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *MediaStream_MediaStream_averageVideoPacketLossRate) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *MediaStream_MediaStream_averageVideoPacketLossRate) GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *MediaStream_MediaStream_averageVideoPacketLossRate) GetString()(*string) {
    val, err := m.GetBackingStore().Get("string")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MediaStream_MediaStream_averageVideoPacketLossRate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetFloat() != nil {
        err := writer.WriteFloat32Value("", m.GetFloat())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *MediaStream_MediaStream_averageVideoPacketLossRate) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFloat sets the float property value. Composed type representation for type float32
func (m *MediaStream_MediaStream_averageVideoPacketLossRate) SetFloat(value *float32)() {
    err := m.GetBackingStore().Set("float", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
func (m *MediaStream_MediaStream_averageVideoPacketLossRate) SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *MediaStream_MediaStream_averageVideoPacketLossRate) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// MediaStream_MediaStream_lowFrameRateRatio composed type wrapper for classes float32, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric, string
type MediaStream_MediaStream_lowFrameRateRatio struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewMediaStream_MediaStream_lowFrameRateRatio instantiates a new MediaStream_MediaStream_lowFrameRateRatio and sets the default values.
func NewMediaStream_MediaStream_lowFrameRateRatio()(*MediaStream_MediaStream_lowFrameRateRatio) {
    m := &MediaStream_MediaStream_lowFrameRateRatio{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateMediaStream_MediaStream_lowFrameRateRatioFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMediaStream_MediaStream_lowFrameRateRatioFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewMediaStream_MediaStream_lowFrameRateRatio()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseReferenceNumeric); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetReferenceNumeric(val)
    } else if val, err := parseNode.GetFloat32Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetFloat(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *MediaStream_MediaStream_lowFrameRateRatio) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MediaStream_MediaStream_lowFrameRateRatio) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetFloat gets the float property value. Composed type representation for type float32
// returns a *float32 when successful
func (m *MediaStream_MediaStream_lowFrameRateRatio) GetFloat()(*float32) {
    val, err := m.GetBackingStore().Get("float")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float32)
    }
    return nil
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *MediaStream_MediaStream_lowFrameRateRatio) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *MediaStream_MediaStream_lowFrameRateRatio) GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *MediaStream_MediaStream_lowFrameRateRatio) GetString()(*string) {
    val, err := m.GetBackingStore().Get("string")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MediaStream_MediaStream_lowFrameRateRatio) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetFloat() != nil {
        err := writer.WriteFloat32Value("", m.GetFloat())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *MediaStream_MediaStream_lowFrameRateRatio) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFloat sets the float property value. Composed type representation for type float32
func (m *MediaStream_MediaStream_lowFrameRateRatio) SetFloat(value *float32)() {
    err := m.GetBackingStore().Set("float", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
func (m *MediaStream_MediaStream_lowFrameRateRatio) SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *MediaStream_MediaStream_lowFrameRateRatio) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// MediaStream_MediaStream_lowVideoProcessingCapabilityRatio composed type wrapper for classes float32, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric, string
type MediaStream_MediaStream_lowVideoProcessingCapabilityRatio struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewMediaStream_MediaStream_lowVideoProcessingCapabilityRatio instantiates a new MediaStream_MediaStream_lowVideoProcessingCapabilityRatio and sets the default values.
func NewMediaStream_MediaStream_lowVideoProcessingCapabilityRatio()(*MediaStream_MediaStream_lowVideoProcessingCapabilityRatio) {
    m := &MediaStream_MediaStream_lowVideoProcessingCapabilityRatio{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateMediaStream_MediaStream_lowVideoProcessingCapabilityRatioFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMediaStream_MediaStream_lowVideoProcessingCapabilityRatioFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewMediaStream_MediaStream_lowVideoProcessingCapabilityRatio()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseReferenceNumeric); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetReferenceNumeric(val)
    } else if val, err := parseNode.GetFloat32Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetFloat(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *MediaStream_MediaStream_lowVideoProcessingCapabilityRatio) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MediaStream_MediaStream_lowVideoProcessingCapabilityRatio) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetFloat gets the float property value. Composed type representation for type float32
// returns a *float32 when successful
func (m *MediaStream_MediaStream_lowVideoProcessingCapabilityRatio) GetFloat()(*float32) {
    val, err := m.GetBackingStore().Get("float")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float32)
    }
    return nil
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *MediaStream_MediaStream_lowVideoProcessingCapabilityRatio) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *MediaStream_MediaStream_lowVideoProcessingCapabilityRatio) GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *MediaStream_MediaStream_lowVideoProcessingCapabilityRatio) GetString()(*string) {
    val, err := m.GetBackingStore().Get("string")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MediaStream_MediaStream_lowVideoProcessingCapabilityRatio) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetFloat() != nil {
        err := writer.WriteFloat32Value("", m.GetFloat())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *MediaStream_MediaStream_lowVideoProcessingCapabilityRatio) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFloat sets the float property value. Composed type representation for type float32
func (m *MediaStream_MediaStream_lowVideoProcessingCapabilityRatio) SetFloat(value *float32)() {
    err := m.GetBackingStore().Set("float", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
func (m *MediaStream_MediaStream_lowVideoProcessingCapabilityRatio) SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *MediaStream_MediaStream_lowVideoProcessingCapabilityRatio) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// MediaStream_MediaStream_maxPacketLossRate composed type wrapper for classes float32, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric, string
type MediaStream_MediaStream_maxPacketLossRate struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewMediaStream_MediaStream_maxPacketLossRate instantiates a new MediaStream_MediaStream_maxPacketLossRate and sets the default values.
func NewMediaStream_MediaStream_maxPacketLossRate()(*MediaStream_MediaStream_maxPacketLossRate) {
    m := &MediaStream_MediaStream_maxPacketLossRate{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateMediaStream_MediaStream_maxPacketLossRateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMediaStream_MediaStream_maxPacketLossRateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewMediaStream_MediaStream_maxPacketLossRate()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseReferenceNumeric); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetReferenceNumeric(val)
    } else if val, err := parseNode.GetFloat32Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetFloat(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *MediaStream_MediaStream_maxPacketLossRate) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MediaStream_MediaStream_maxPacketLossRate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetFloat gets the float property value. Composed type representation for type float32
// returns a *float32 when successful
func (m *MediaStream_MediaStream_maxPacketLossRate) GetFloat()(*float32) {
    val, err := m.GetBackingStore().Get("float")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float32)
    }
    return nil
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *MediaStream_MediaStream_maxPacketLossRate) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *MediaStream_MediaStream_maxPacketLossRate) GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *MediaStream_MediaStream_maxPacketLossRate) GetString()(*string) {
    val, err := m.GetBackingStore().Get("string")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MediaStream_MediaStream_maxPacketLossRate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetFloat() != nil {
        err := writer.WriteFloat32Value("", m.GetFloat())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *MediaStream_MediaStream_maxPacketLossRate) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFloat sets the float property value. Composed type representation for type float32
func (m *MediaStream_MediaStream_maxPacketLossRate) SetFloat(value *float32)() {
    err := m.GetBackingStore().Set("float", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
func (m *MediaStream_MediaStream_maxPacketLossRate) SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *MediaStream_MediaStream_maxPacketLossRate) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// MediaStream_MediaStream_maxRatioOfConcealedSamples composed type wrapper for classes float32, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric, string
type MediaStream_MediaStream_maxRatioOfConcealedSamples struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewMediaStream_MediaStream_maxRatioOfConcealedSamples instantiates a new MediaStream_MediaStream_maxRatioOfConcealedSamples and sets the default values.
func NewMediaStream_MediaStream_maxRatioOfConcealedSamples()(*MediaStream_MediaStream_maxRatioOfConcealedSamples) {
    m := &MediaStream_MediaStream_maxRatioOfConcealedSamples{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateMediaStream_MediaStream_maxRatioOfConcealedSamplesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMediaStream_MediaStream_maxRatioOfConcealedSamplesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewMediaStream_MediaStream_maxRatioOfConcealedSamples()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseReferenceNumeric); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetReferenceNumeric(val)
    } else if val, err := parseNode.GetFloat32Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetFloat(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *MediaStream_MediaStream_maxRatioOfConcealedSamples) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MediaStream_MediaStream_maxRatioOfConcealedSamples) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetFloat gets the float property value. Composed type representation for type float32
// returns a *float32 when successful
func (m *MediaStream_MediaStream_maxRatioOfConcealedSamples) GetFloat()(*float32) {
    val, err := m.GetBackingStore().Get("float")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float32)
    }
    return nil
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *MediaStream_MediaStream_maxRatioOfConcealedSamples) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *MediaStream_MediaStream_maxRatioOfConcealedSamples) GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *MediaStream_MediaStream_maxRatioOfConcealedSamples) GetString()(*string) {
    val, err := m.GetBackingStore().Get("string")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MediaStream_MediaStream_maxRatioOfConcealedSamples) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetFloat() != nil {
        err := writer.WriteFloat32Value("", m.GetFloat())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *MediaStream_MediaStream_maxRatioOfConcealedSamples) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFloat sets the float property value. Composed type representation for type float32
func (m *MediaStream_MediaStream_maxRatioOfConcealedSamples) SetFloat(value *float32)() {
    err := m.GetBackingStore().Set("float", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
func (m *MediaStream_MediaStream_maxRatioOfConcealedSamples) SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *MediaStream_MediaStream_maxRatioOfConcealedSamples) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// MediaStream_MediaStream_postForwardErrorCorrectionPacketLossRate composed type wrapper for classes float32, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric, string
type MediaStream_MediaStream_postForwardErrorCorrectionPacketLossRate struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewMediaStream_MediaStream_postForwardErrorCorrectionPacketLossRate instantiates a new MediaStream_MediaStream_postForwardErrorCorrectionPacketLossRate and sets the default values.
func NewMediaStream_MediaStream_postForwardErrorCorrectionPacketLossRate()(*MediaStream_MediaStream_postForwardErrorCorrectionPacketLossRate) {
    m := &MediaStream_MediaStream_postForwardErrorCorrectionPacketLossRate{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateMediaStream_MediaStream_postForwardErrorCorrectionPacketLossRateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMediaStream_MediaStream_postForwardErrorCorrectionPacketLossRateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewMediaStream_MediaStream_postForwardErrorCorrectionPacketLossRate()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseReferenceNumeric); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetReferenceNumeric(val)
    } else if val, err := parseNode.GetFloat32Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetFloat(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *MediaStream_MediaStream_postForwardErrorCorrectionPacketLossRate) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MediaStream_MediaStream_postForwardErrorCorrectionPacketLossRate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetFloat gets the float property value. Composed type representation for type float32
// returns a *float32 when successful
func (m *MediaStream_MediaStream_postForwardErrorCorrectionPacketLossRate) GetFloat()(*float32) {
    val, err := m.GetBackingStore().Get("float")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float32)
    }
    return nil
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *MediaStream_MediaStream_postForwardErrorCorrectionPacketLossRate) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *MediaStream_MediaStream_postForwardErrorCorrectionPacketLossRate) GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *MediaStream_MediaStream_postForwardErrorCorrectionPacketLossRate) GetString()(*string) {
    val, err := m.GetBackingStore().Get("string")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MediaStream_MediaStream_postForwardErrorCorrectionPacketLossRate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetFloat() != nil {
        err := writer.WriteFloat32Value("", m.GetFloat())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *MediaStream_MediaStream_postForwardErrorCorrectionPacketLossRate) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFloat sets the float property value. Composed type representation for type float32
func (m *MediaStream_MediaStream_postForwardErrorCorrectionPacketLossRate) SetFloat(value *float32)() {
    err := m.GetBackingStore().Set("float", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
func (m *MediaStream_MediaStream_postForwardErrorCorrectionPacketLossRate) SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *MediaStream_MediaStream_postForwardErrorCorrectionPacketLossRate) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
type MediaStream_MediaStream_averageAudioDegradationable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFloat()(*float32)
    GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFloat(value *float32)()
    SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)()
    SetString(value *string)()
}
type MediaStream_MediaStream_averagePacketLossRateable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFloat()(*float32)
    GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFloat(value *float32)()
    SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)()
    SetString(value *string)()
}
type MediaStream_MediaStream_averageRatioOfConcealedSamplesable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFloat()(*float32)
    GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFloat(value *float32)()
    SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)()
    SetString(value *string)()
}
type MediaStream_MediaStream_averageReceivedFrameRateable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFloat()(*float32)
    GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFloat(value *float32)()
    SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)()
    SetString(value *string)()
}
type MediaStream_MediaStream_averageVideoFrameLossPercentageable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFloat()(*float32)
    GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFloat(value *float32)()
    SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)()
    SetString(value *string)()
}
type MediaStream_MediaStream_averageVideoFrameRateable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFloat()(*float32)
    GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFloat(value *float32)()
    SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)()
    SetString(value *string)()
}
type MediaStream_MediaStream_averageVideoPacketLossRateable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFloat()(*float32)
    GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFloat(value *float32)()
    SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)()
    SetString(value *string)()
}
type MediaStream_MediaStream_lowFrameRateRatioable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFloat()(*float32)
    GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFloat(value *float32)()
    SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)()
    SetString(value *string)()
}
type MediaStream_MediaStream_lowVideoProcessingCapabilityRatioable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFloat()(*float32)
    GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFloat(value *float32)()
    SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)()
    SetString(value *string)()
}
type MediaStream_MediaStream_maxPacketLossRateable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFloat()(*float32)
    GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFloat(value *float32)()
    SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)()
    SetString(value *string)()
}
type MediaStream_MediaStream_maxRatioOfConcealedSamplesable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFloat()(*float32)
    GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFloat(value *float32)()
    SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)()
    SetString(value *string)()
}
type MediaStream_MediaStream_postForwardErrorCorrectionPacketLossRateable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFloat()(*float32)
    GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFloat(value *float32)()
    SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)()
    SetString(value *string)()
}
// NewMediaStream instantiates a new MediaStream and sets the default values.
func NewMediaStream()(*MediaStream) {
    m := &MediaStream{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMediaStreamFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMediaStreamFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMediaStream(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *MediaStream) GetAdditionalData()(map[string]any) {
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
// GetAudioCodec gets the audioCodec property value. Codec name used to encode audio for transmission on the network. Possible values are: unknown, invalid, cn, pcma, pcmu, amrWide, g722, g7221, g7221c, g729, multiChannelAudio, muchv2, opus, satin, satinFullband, rtAudio8, rtAudio16, silk, silkNarrow, silkWide, siren, xmsRta, unknownFutureValue.
// returns a *AudioCodec when successful
func (m *MediaStream) GetAudioCodec()(*AudioCodec) {
    val, err := m.GetBackingStore().Get("audioCodec")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AudioCodec)
    }
    return nil
}
// GetAverageAudioDegradation gets the averageAudioDegradation property value. Average Network Mean Opinion Score degradation for stream. Represents how much the network loss and jitter has impacted the quality of received audio.
// returns a MediaStream_MediaStream_averageAudioDegradationable when successful
func (m *MediaStream) GetAverageAudioDegradation()(MediaStream_MediaStream_averageAudioDegradationable) {
    val, err := m.GetBackingStore().Get("averageAudioDegradation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MediaStream_MediaStream_averageAudioDegradationable)
    }
    return nil
}
// GetAverageAudioNetworkJitter gets the averageAudioNetworkJitter property value. Average jitter for the stream computed as specified in RFC 3550, denoted in ISO 8601 format. For example, 1 second is denoted as 'PT1S', where 'P' is the duration designator, 'T' is the time designator, and 'S' is the second designator.
// returns a *ISODuration when successful
func (m *MediaStream) GetAverageAudioNetworkJitter()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("averageAudioNetworkJitter")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetAverageBandwidthEstimate gets the averageBandwidthEstimate property value. Average estimated bandwidth available between two endpoints in bits per second.
// returns a *int64 when successful
func (m *MediaStream) GetAverageBandwidthEstimate()(*int64) {
    val, err := m.GetBackingStore().Get("averageBandwidthEstimate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetAverageFreezeDuration gets the averageFreezeDuration property value. Average duration of the received freezing time in the video stream.
// returns a *ISODuration when successful
func (m *MediaStream) GetAverageFreezeDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("averageFreezeDuration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetAverageJitter gets the averageJitter property value. Average jitter for the stream computed as specified in RFC 3550, denoted in ISO 8601 format. For example, 1 second is denoted as 'PT1S', where 'P' is the duration designator, 'T' is the time designator, and 'S' is the second designator.
// returns a *ISODuration when successful
func (m *MediaStream) GetAverageJitter()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("averageJitter")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetAveragePacketLossRate gets the averagePacketLossRate property value. Average packet loss rate for stream.
// returns a MediaStream_MediaStream_averagePacketLossRateable when successful
func (m *MediaStream) GetAveragePacketLossRate()(MediaStream_MediaStream_averagePacketLossRateable) {
    val, err := m.GetBackingStore().Get("averagePacketLossRate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MediaStream_MediaStream_averagePacketLossRateable)
    }
    return nil
}
// GetAverageRatioOfConcealedSamples gets the averageRatioOfConcealedSamples property value. Ratio of the number of audio frames with samples generated by packet loss concealment to the total number of audio frames.
// returns a MediaStream_MediaStream_averageRatioOfConcealedSamplesable when successful
func (m *MediaStream) GetAverageRatioOfConcealedSamples()(MediaStream_MediaStream_averageRatioOfConcealedSamplesable) {
    val, err := m.GetBackingStore().Get("averageRatioOfConcealedSamples")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MediaStream_MediaStream_averageRatioOfConcealedSamplesable)
    }
    return nil
}
// GetAverageReceivedFrameRate gets the averageReceivedFrameRate property value. Average frames per second received for all video streams computed over the duration of the session.
// returns a MediaStream_MediaStream_averageReceivedFrameRateable when successful
func (m *MediaStream) GetAverageReceivedFrameRate()(MediaStream_MediaStream_averageReceivedFrameRateable) {
    val, err := m.GetBackingStore().Get("averageReceivedFrameRate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MediaStream_MediaStream_averageReceivedFrameRateable)
    }
    return nil
}
// GetAverageRoundTripTime gets the averageRoundTripTime property value. Average network propagation round-trip time computed as specified in RFC 3550, denoted in ISO 8601 format. For example, 1 second is denoted as 'PT1S', where 'P' is the duration designator, 'T' is the time designator, and 'S' is the second designator.
// returns a *ISODuration when successful
func (m *MediaStream) GetAverageRoundTripTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("averageRoundTripTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetAverageVideoFrameLossPercentage gets the averageVideoFrameLossPercentage property value. Average percentage of video frames lost as displayed to the user.
// returns a MediaStream_MediaStream_averageVideoFrameLossPercentageable when successful
func (m *MediaStream) GetAverageVideoFrameLossPercentage()(MediaStream_MediaStream_averageVideoFrameLossPercentageable) {
    val, err := m.GetBackingStore().Get("averageVideoFrameLossPercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MediaStream_MediaStream_averageVideoFrameLossPercentageable)
    }
    return nil
}
// GetAverageVideoFrameRate gets the averageVideoFrameRate property value. Average frames per second received for a video stream, computed over the duration of the session.
// returns a MediaStream_MediaStream_averageVideoFrameRateable when successful
func (m *MediaStream) GetAverageVideoFrameRate()(MediaStream_MediaStream_averageVideoFrameRateable) {
    val, err := m.GetBackingStore().Get("averageVideoFrameRate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MediaStream_MediaStream_averageVideoFrameRateable)
    }
    return nil
}
// GetAverageVideoPacketLossRate gets the averageVideoPacketLossRate property value. Average fraction of packets lost, as specified in RFC 3550, computed over the duration of the session.
// returns a MediaStream_MediaStream_averageVideoPacketLossRateable when successful
func (m *MediaStream) GetAverageVideoPacketLossRate()(MediaStream_MediaStream_averageVideoPacketLossRateable) {
    val, err := m.GetBackingStore().Get("averageVideoPacketLossRate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MediaStream_MediaStream_averageVideoPacketLossRateable)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *MediaStream) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetEndDateTime gets the endDateTime property value. UTC time when the stream ended. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. This field is only available for streams that use the SIP protocol.
// returns a *Time when successful
func (m *MediaStream) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("endDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MediaStream) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["audioCodec"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAudioCodec)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAudioCodec(val.(*AudioCodec))
        }
        return nil
    }
    res["averageAudioDegradation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMediaStream_MediaStream_averageAudioDegradationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageAudioDegradation(val.(*MediaStream_MediaStream_averageAudioDegradation))
        }
        return nil
    }
    res["averageAudioNetworkJitter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageAudioNetworkJitter(val)
        }
        return nil
    }
    res["averageBandwidthEstimate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageBandwidthEstimate(val)
        }
        return nil
    }
    res["averageFreezeDuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageFreezeDuration(val)
        }
        return nil
    }
    res["averageJitter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageJitter(val)
        }
        return nil
    }
    res["averagePacketLossRate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMediaStream_MediaStream_averagePacketLossRateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAveragePacketLossRate(val.(*MediaStream_MediaStream_averagePacketLossRate))
        }
        return nil
    }
    res["averageRatioOfConcealedSamples"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMediaStream_MediaStream_averageRatioOfConcealedSamplesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageRatioOfConcealedSamples(val.(*MediaStream_MediaStream_averageRatioOfConcealedSamples))
        }
        return nil
    }
    res["averageReceivedFrameRate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMediaStream_MediaStream_averageReceivedFrameRateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageReceivedFrameRate(val.(*MediaStream_MediaStream_averageReceivedFrameRate))
        }
        return nil
    }
    res["averageRoundTripTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageRoundTripTime(val)
        }
        return nil
    }
    res["averageVideoFrameLossPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMediaStream_MediaStream_averageVideoFrameLossPercentageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageVideoFrameLossPercentage(val.(*MediaStream_MediaStream_averageVideoFrameLossPercentage))
        }
        return nil
    }
    res["averageVideoFrameRate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMediaStream_MediaStream_averageVideoFrameRateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageVideoFrameRate(val.(*MediaStream_MediaStream_averageVideoFrameRate))
        }
        return nil
    }
    res["averageVideoPacketLossRate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMediaStream_MediaStream_averageVideoPacketLossRateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageVideoPacketLossRate(val.(*MediaStream_MediaStream_averageVideoPacketLossRate))
        }
        return nil
    }
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["isAudioForwardErrorCorrectionUsed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAudioForwardErrorCorrectionUsed(val)
        }
        return nil
    }
    res["lowFrameRateRatio"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMediaStream_MediaStream_lowFrameRateRatioFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLowFrameRateRatio(val.(*MediaStream_MediaStream_lowFrameRateRatio))
        }
        return nil
    }
    res["lowVideoProcessingCapabilityRatio"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMediaStream_MediaStream_lowVideoProcessingCapabilityRatioFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLowVideoProcessingCapabilityRatio(val.(*MediaStream_MediaStream_lowVideoProcessingCapabilityRatio))
        }
        return nil
    }
    res["maxAudioNetworkJitter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxAudioNetworkJitter(val)
        }
        return nil
    }
    res["maxJitter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxJitter(val)
        }
        return nil
    }
    res["maxPacketLossRate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMediaStream_MediaStream_maxPacketLossRateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxPacketLossRate(val.(*MediaStream_MediaStream_maxPacketLossRate))
        }
        return nil
    }
    res["maxRatioOfConcealedSamples"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMediaStream_MediaStream_maxRatioOfConcealedSamplesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxRatioOfConcealedSamples(val.(*MediaStream_MediaStream_maxRatioOfConcealedSamples))
        }
        return nil
    }
    res["maxRoundTripTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxRoundTripTime(val)
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
    res["packetUtilization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPacketUtilization(val)
        }
        return nil
    }
    res["postForwardErrorCorrectionPacketLossRate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMediaStream_MediaStream_postForwardErrorCorrectionPacketLossRateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostForwardErrorCorrectionPacketLossRate(val.(*MediaStream_MediaStream_postForwardErrorCorrectionPacketLossRate))
        }
        return nil
    }
    res["rmsFreezeDuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRmsFreezeDuration(val)
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    res["streamDirection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMediaStreamDirection)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStreamDirection(val.(*MediaStreamDirection))
        }
        return nil
    }
    res["streamId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStreamId(val)
        }
        return nil
    }
    res["videoCodec"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVideoCodec)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVideoCodec(val.(*VideoCodec))
        }
        return nil
    }
    res["wasMediaBypassed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWasMediaBypassed(val)
        }
        return nil
    }
    return res
}
// GetIsAudioForwardErrorCorrectionUsed gets the isAudioForwardErrorCorrectionUsed property value. Indicates whether the forward error correction (FEC) was used at some point during the session. The default value is null.
// returns a *bool when successful
func (m *MediaStream) GetIsAudioForwardErrorCorrectionUsed()(*bool) {
    val, err := m.GetBackingStore().Get("isAudioForwardErrorCorrectionUsed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLowFrameRateRatio gets the lowFrameRateRatio property value. Fraction of the call where frame rate is less than 7.5 frames per second.
// returns a MediaStream_MediaStream_lowFrameRateRatioable when successful
func (m *MediaStream) GetLowFrameRateRatio()(MediaStream_MediaStream_lowFrameRateRatioable) {
    val, err := m.GetBackingStore().Get("lowFrameRateRatio")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MediaStream_MediaStream_lowFrameRateRatioable)
    }
    return nil
}
// GetLowVideoProcessingCapabilityRatio gets the lowVideoProcessingCapabilityRatio property value. Fraction of the call that the client is running less than 70% expected video processing capability.
// returns a MediaStream_MediaStream_lowVideoProcessingCapabilityRatioable when successful
func (m *MediaStream) GetLowVideoProcessingCapabilityRatio()(MediaStream_MediaStream_lowVideoProcessingCapabilityRatioable) {
    val, err := m.GetBackingStore().Get("lowVideoProcessingCapabilityRatio")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MediaStream_MediaStream_lowVideoProcessingCapabilityRatioable)
    }
    return nil
}
// GetMaxAudioNetworkJitter gets the maxAudioNetworkJitter property value. Maximum of audio network jitter computed over each of the 20 second windows during the session, denoted in ISO 8601 format. For example, 1 second is denoted as 'PT1S', where 'P' is the duration designator, 'T' is the time designator, and 'S' is the second designator.
// returns a *ISODuration when successful
func (m *MediaStream) GetMaxAudioNetworkJitter()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("maxAudioNetworkJitter")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetMaxJitter gets the maxJitter property value. Maximum jitter for the stream computed as specified in RFC 3550, denoted in ISO 8601 format. For example, 1 second is denoted as 'PT1S', where 'P' is the duration designator, 'T' is the time designator, and 'S' is the second designator.
// returns a *ISODuration when successful
func (m *MediaStream) GetMaxJitter()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("maxJitter")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetMaxPacketLossRate gets the maxPacketLossRate property value. Maximum packet loss rate for the stream.
// returns a MediaStream_MediaStream_maxPacketLossRateable when successful
func (m *MediaStream) GetMaxPacketLossRate()(MediaStream_MediaStream_maxPacketLossRateable) {
    val, err := m.GetBackingStore().Get("maxPacketLossRate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MediaStream_MediaStream_maxPacketLossRateable)
    }
    return nil
}
// GetMaxRatioOfConcealedSamples gets the maxRatioOfConcealedSamples property value. Maximum ratio of packets concealed by the healer.
// returns a MediaStream_MediaStream_maxRatioOfConcealedSamplesable when successful
func (m *MediaStream) GetMaxRatioOfConcealedSamples()(MediaStream_MediaStream_maxRatioOfConcealedSamplesable) {
    val, err := m.GetBackingStore().Get("maxRatioOfConcealedSamples")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MediaStream_MediaStream_maxRatioOfConcealedSamplesable)
    }
    return nil
}
// GetMaxRoundTripTime gets the maxRoundTripTime property value. Maximum network propagation round-trip time computed as specified in RFC 3550, denoted in ISO 8601 format. For example, 1 second is denoted as 'PT1S', where 'P' is the duration designator, 'T' is the time designator, and 'S' is the second designator.
// returns a *ISODuration when successful
func (m *MediaStream) GetMaxRoundTripTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("maxRoundTripTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *MediaStream) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPacketUtilization gets the packetUtilization property value. Packet count for the stream.
// returns a *int64 when successful
func (m *MediaStream) GetPacketUtilization()(*int64) {
    val, err := m.GetBackingStore().Get("packetUtilization")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetPostForwardErrorCorrectionPacketLossRate gets the postForwardErrorCorrectionPacketLossRate property value. Packet loss rate after FEC has been applied aggregated across all video streams and codecs.
// returns a MediaStream_MediaStream_postForwardErrorCorrectionPacketLossRateable when successful
func (m *MediaStream) GetPostForwardErrorCorrectionPacketLossRate()(MediaStream_MediaStream_postForwardErrorCorrectionPacketLossRateable) {
    val, err := m.GetBackingStore().Get("postForwardErrorCorrectionPacketLossRate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MediaStream_MediaStream_postForwardErrorCorrectionPacketLossRateable)
    }
    return nil
}
// GetRmsFreezeDuration gets the rmsFreezeDuration property value. Average duration of the received freezing time in the video stream represented in root mean square.
// returns a *ISODuration when successful
func (m *MediaStream) GetRmsFreezeDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("rmsFreezeDuration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetStartDateTime gets the startDateTime property value. UTC time when the stream started. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. This field is only available for streams that use the SIP protocol.
// returns a *Time when successful
func (m *MediaStream) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("startDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetStreamDirection gets the streamDirection property value. The streamDirection property
// returns a *MediaStreamDirection when successful
func (m *MediaStream) GetStreamDirection()(*MediaStreamDirection) {
    val, err := m.GetBackingStore().Get("streamDirection")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MediaStreamDirection)
    }
    return nil
}
// GetStreamId gets the streamId property value. Unique identifier for the stream.
// returns a *string when successful
func (m *MediaStream) GetStreamId()(*string) {
    val, err := m.GetBackingStore().Get("streamId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetVideoCodec gets the videoCodec property value. Codec name used to encode video for transmission on the network. Possible values are: unknown, invalid, av1, h263, h264, h264s, h264uc, h265, rtvc1, rtVideo, xrtvc1, unknownFutureValue.
// returns a *VideoCodec when successful
func (m *MediaStream) GetVideoCodec()(*VideoCodec) {
    val, err := m.GetBackingStore().Get("videoCodec")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*VideoCodec)
    }
    return nil
}
// GetWasMediaBypassed gets the wasMediaBypassed property value. True if the media stream bypassed the Mediation Server and went straight between client and PSTN Gateway/PBX, false otherwise.
// returns a *bool when successful
func (m *MediaStream) GetWasMediaBypassed()(*bool) {
    val, err := m.GetBackingStore().Get("wasMediaBypassed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MediaStream) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAudioCodec() != nil {
        cast := (*m.GetAudioCodec()).String()
        err := writer.WriteStringValue("audioCodec", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("averageAudioDegradation", m.GetAverageAudioDegradation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("averageAudioNetworkJitter", m.GetAverageAudioNetworkJitter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("averageBandwidthEstimate", m.GetAverageBandwidthEstimate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("averageFreezeDuration", m.GetAverageFreezeDuration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("averageJitter", m.GetAverageJitter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("averagePacketLossRate", m.GetAveragePacketLossRate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("averageRatioOfConcealedSamples", m.GetAverageRatioOfConcealedSamples())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("averageReceivedFrameRate", m.GetAverageReceivedFrameRate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("averageRoundTripTime", m.GetAverageRoundTripTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("averageVideoFrameLossPercentage", m.GetAverageVideoFrameLossPercentage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("averageVideoFrameRate", m.GetAverageVideoFrameRate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("averageVideoPacketLossRate", m.GetAverageVideoPacketLossRate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isAudioForwardErrorCorrectionUsed", m.GetIsAudioForwardErrorCorrectionUsed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("lowFrameRateRatio", m.GetLowFrameRateRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("lowVideoProcessingCapabilityRatio", m.GetLowVideoProcessingCapabilityRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("maxAudioNetworkJitter", m.GetMaxAudioNetworkJitter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("maxJitter", m.GetMaxJitter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("maxPacketLossRate", m.GetMaxPacketLossRate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("maxRatioOfConcealedSamples", m.GetMaxRatioOfConcealedSamples())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("maxRoundTripTime", m.GetMaxRoundTripTime())
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
        err := writer.WriteInt64Value("packetUtilization", m.GetPacketUtilization())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("postForwardErrorCorrectionPacketLossRate", m.GetPostForwardErrorCorrectionPacketLossRate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("rmsFreezeDuration", m.GetRmsFreezeDuration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetStreamDirection() != nil {
        cast := (*m.GetStreamDirection()).String()
        err := writer.WriteStringValue("streamDirection", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("streamId", m.GetStreamId())
        if err != nil {
            return err
        }
    }
    if m.GetVideoCodec() != nil {
        cast := (*m.GetVideoCodec()).String()
        err := writer.WriteStringValue("videoCodec", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("wasMediaBypassed", m.GetWasMediaBypassed())
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
func (m *MediaStream) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAudioCodec sets the audioCodec property value. Codec name used to encode audio for transmission on the network. Possible values are: unknown, invalid, cn, pcma, pcmu, amrWide, g722, g7221, g7221c, g729, multiChannelAudio, muchv2, opus, satin, satinFullband, rtAudio8, rtAudio16, silk, silkNarrow, silkWide, siren, xmsRta, unknownFutureValue.
func (m *MediaStream) SetAudioCodec(value *AudioCodec)() {
    err := m.GetBackingStore().Set("audioCodec", value)
    if err != nil {
        panic(err)
    }
}
// SetAverageAudioDegradation sets the averageAudioDegradation property value. Average Network Mean Opinion Score degradation for stream. Represents how much the network loss and jitter has impacted the quality of received audio.
func (m *MediaStream) SetAverageAudioDegradation(value MediaStream_MediaStream_averageAudioDegradationable)() {
    err := m.GetBackingStore().Set("averageAudioDegradation", value)
    if err != nil {
        panic(err)
    }
}
// SetAverageAudioNetworkJitter sets the averageAudioNetworkJitter property value. Average jitter for the stream computed as specified in RFC 3550, denoted in ISO 8601 format. For example, 1 second is denoted as 'PT1S', where 'P' is the duration designator, 'T' is the time designator, and 'S' is the second designator.
func (m *MediaStream) SetAverageAudioNetworkJitter(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("averageAudioNetworkJitter", value)
    if err != nil {
        panic(err)
    }
}
// SetAverageBandwidthEstimate sets the averageBandwidthEstimate property value. Average estimated bandwidth available between two endpoints in bits per second.
func (m *MediaStream) SetAverageBandwidthEstimate(value *int64)() {
    err := m.GetBackingStore().Set("averageBandwidthEstimate", value)
    if err != nil {
        panic(err)
    }
}
// SetAverageFreezeDuration sets the averageFreezeDuration property value. Average duration of the received freezing time in the video stream.
func (m *MediaStream) SetAverageFreezeDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("averageFreezeDuration", value)
    if err != nil {
        panic(err)
    }
}
// SetAverageJitter sets the averageJitter property value. Average jitter for the stream computed as specified in RFC 3550, denoted in ISO 8601 format. For example, 1 second is denoted as 'PT1S', where 'P' is the duration designator, 'T' is the time designator, and 'S' is the second designator.
func (m *MediaStream) SetAverageJitter(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("averageJitter", value)
    if err != nil {
        panic(err)
    }
}
// SetAveragePacketLossRate sets the averagePacketLossRate property value. Average packet loss rate for stream.
func (m *MediaStream) SetAveragePacketLossRate(value MediaStream_MediaStream_averagePacketLossRateable)() {
    err := m.GetBackingStore().Set("averagePacketLossRate", value)
    if err != nil {
        panic(err)
    }
}
// SetAverageRatioOfConcealedSamples sets the averageRatioOfConcealedSamples property value. Ratio of the number of audio frames with samples generated by packet loss concealment to the total number of audio frames.
func (m *MediaStream) SetAverageRatioOfConcealedSamples(value MediaStream_MediaStream_averageRatioOfConcealedSamplesable)() {
    err := m.GetBackingStore().Set("averageRatioOfConcealedSamples", value)
    if err != nil {
        panic(err)
    }
}
// SetAverageReceivedFrameRate sets the averageReceivedFrameRate property value. Average frames per second received for all video streams computed over the duration of the session.
func (m *MediaStream) SetAverageReceivedFrameRate(value MediaStream_MediaStream_averageReceivedFrameRateable)() {
    err := m.GetBackingStore().Set("averageReceivedFrameRate", value)
    if err != nil {
        panic(err)
    }
}
// SetAverageRoundTripTime sets the averageRoundTripTime property value. Average network propagation round-trip time computed as specified in RFC 3550, denoted in ISO 8601 format. For example, 1 second is denoted as 'PT1S', where 'P' is the duration designator, 'T' is the time designator, and 'S' is the second designator.
func (m *MediaStream) SetAverageRoundTripTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("averageRoundTripTime", value)
    if err != nil {
        panic(err)
    }
}
// SetAverageVideoFrameLossPercentage sets the averageVideoFrameLossPercentage property value. Average percentage of video frames lost as displayed to the user.
func (m *MediaStream) SetAverageVideoFrameLossPercentage(value MediaStream_MediaStream_averageVideoFrameLossPercentageable)() {
    err := m.GetBackingStore().Set("averageVideoFrameLossPercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetAverageVideoFrameRate sets the averageVideoFrameRate property value. Average frames per second received for a video stream, computed over the duration of the session.
func (m *MediaStream) SetAverageVideoFrameRate(value MediaStream_MediaStream_averageVideoFrameRateable)() {
    err := m.GetBackingStore().Set("averageVideoFrameRate", value)
    if err != nil {
        panic(err)
    }
}
// SetAverageVideoPacketLossRate sets the averageVideoPacketLossRate property value. Average fraction of packets lost, as specified in RFC 3550, computed over the duration of the session.
func (m *MediaStream) SetAverageVideoPacketLossRate(value MediaStream_MediaStream_averageVideoPacketLossRateable)() {
    err := m.GetBackingStore().Set("averageVideoPacketLossRate", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *MediaStream) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetEndDateTime sets the endDateTime property value. UTC time when the stream ended. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. This field is only available for streams that use the SIP protocol.
func (m *MediaStream) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("endDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetIsAudioForwardErrorCorrectionUsed sets the isAudioForwardErrorCorrectionUsed property value. Indicates whether the forward error correction (FEC) was used at some point during the session. The default value is null.
func (m *MediaStream) SetIsAudioForwardErrorCorrectionUsed(value *bool)() {
    err := m.GetBackingStore().Set("isAudioForwardErrorCorrectionUsed", value)
    if err != nil {
        panic(err)
    }
}
// SetLowFrameRateRatio sets the lowFrameRateRatio property value. Fraction of the call where frame rate is less than 7.5 frames per second.
func (m *MediaStream) SetLowFrameRateRatio(value MediaStream_MediaStream_lowFrameRateRatioable)() {
    err := m.GetBackingStore().Set("lowFrameRateRatio", value)
    if err != nil {
        panic(err)
    }
}
// SetLowVideoProcessingCapabilityRatio sets the lowVideoProcessingCapabilityRatio property value. Fraction of the call that the client is running less than 70% expected video processing capability.
func (m *MediaStream) SetLowVideoProcessingCapabilityRatio(value MediaStream_MediaStream_lowVideoProcessingCapabilityRatioable)() {
    err := m.GetBackingStore().Set("lowVideoProcessingCapabilityRatio", value)
    if err != nil {
        panic(err)
    }
}
// SetMaxAudioNetworkJitter sets the maxAudioNetworkJitter property value. Maximum of audio network jitter computed over each of the 20 second windows during the session, denoted in ISO 8601 format. For example, 1 second is denoted as 'PT1S', where 'P' is the duration designator, 'T' is the time designator, and 'S' is the second designator.
func (m *MediaStream) SetMaxAudioNetworkJitter(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("maxAudioNetworkJitter", value)
    if err != nil {
        panic(err)
    }
}
// SetMaxJitter sets the maxJitter property value. Maximum jitter for the stream computed as specified in RFC 3550, denoted in ISO 8601 format. For example, 1 second is denoted as 'PT1S', where 'P' is the duration designator, 'T' is the time designator, and 'S' is the second designator.
func (m *MediaStream) SetMaxJitter(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("maxJitter", value)
    if err != nil {
        panic(err)
    }
}
// SetMaxPacketLossRate sets the maxPacketLossRate property value. Maximum packet loss rate for the stream.
func (m *MediaStream) SetMaxPacketLossRate(value MediaStream_MediaStream_maxPacketLossRateable)() {
    err := m.GetBackingStore().Set("maxPacketLossRate", value)
    if err != nil {
        panic(err)
    }
}
// SetMaxRatioOfConcealedSamples sets the maxRatioOfConcealedSamples property value. Maximum ratio of packets concealed by the healer.
func (m *MediaStream) SetMaxRatioOfConcealedSamples(value MediaStream_MediaStream_maxRatioOfConcealedSamplesable)() {
    err := m.GetBackingStore().Set("maxRatioOfConcealedSamples", value)
    if err != nil {
        panic(err)
    }
}
// SetMaxRoundTripTime sets the maxRoundTripTime property value. Maximum network propagation round-trip time computed as specified in RFC 3550, denoted in ISO 8601 format. For example, 1 second is denoted as 'PT1S', where 'P' is the duration designator, 'T' is the time designator, and 'S' is the second designator.
func (m *MediaStream) SetMaxRoundTripTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("maxRoundTripTime", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MediaStream) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPacketUtilization sets the packetUtilization property value. Packet count for the stream.
func (m *MediaStream) SetPacketUtilization(value *int64)() {
    err := m.GetBackingStore().Set("packetUtilization", value)
    if err != nil {
        panic(err)
    }
}
// SetPostForwardErrorCorrectionPacketLossRate sets the postForwardErrorCorrectionPacketLossRate property value. Packet loss rate after FEC has been applied aggregated across all video streams and codecs.
func (m *MediaStream) SetPostForwardErrorCorrectionPacketLossRate(value MediaStream_MediaStream_postForwardErrorCorrectionPacketLossRateable)() {
    err := m.GetBackingStore().Set("postForwardErrorCorrectionPacketLossRate", value)
    if err != nil {
        panic(err)
    }
}
// SetRmsFreezeDuration sets the rmsFreezeDuration property value. Average duration of the received freezing time in the video stream represented in root mean square.
func (m *MediaStream) SetRmsFreezeDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("rmsFreezeDuration", value)
    if err != nil {
        panic(err)
    }
}
// SetStartDateTime sets the startDateTime property value. UTC time when the stream started. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. This field is only available for streams that use the SIP protocol.
func (m *MediaStream) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("startDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetStreamDirection sets the streamDirection property value. The streamDirection property
func (m *MediaStream) SetStreamDirection(value *MediaStreamDirection)() {
    err := m.GetBackingStore().Set("streamDirection", value)
    if err != nil {
        panic(err)
    }
}
// SetStreamId sets the streamId property value. Unique identifier for the stream.
func (m *MediaStream) SetStreamId(value *string)() {
    err := m.GetBackingStore().Set("streamId", value)
    if err != nil {
        panic(err)
    }
}
// SetVideoCodec sets the videoCodec property value. Codec name used to encode video for transmission on the network. Possible values are: unknown, invalid, av1, h263, h264, h264s, h264uc, h265, rtvc1, rtVideo, xrtvc1, unknownFutureValue.
func (m *MediaStream) SetVideoCodec(value *VideoCodec)() {
    err := m.GetBackingStore().Set("videoCodec", value)
    if err != nil {
        panic(err)
    }
}
// SetWasMediaBypassed sets the wasMediaBypassed property value. True if the media stream bypassed the Mediation Server and went straight between client and PSTN Gateway/PBX, false otherwise.
func (m *MediaStream) SetWasMediaBypassed(value *bool)() {
    err := m.GetBackingStore().Set("wasMediaBypassed", value)
    if err != nil {
        panic(err)
    }
}
type MediaStreamable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAudioCodec()(*AudioCodec)
    GetAverageAudioDegradation()(MediaStream_MediaStream_averageAudioDegradationable)
    GetAverageAudioNetworkJitter()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetAverageBandwidthEstimate()(*int64)
    GetAverageFreezeDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetAverageJitter()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetAveragePacketLossRate()(MediaStream_MediaStream_averagePacketLossRateable)
    GetAverageRatioOfConcealedSamples()(MediaStream_MediaStream_averageRatioOfConcealedSamplesable)
    GetAverageReceivedFrameRate()(MediaStream_MediaStream_averageReceivedFrameRateable)
    GetAverageRoundTripTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetAverageVideoFrameLossPercentage()(MediaStream_MediaStream_averageVideoFrameLossPercentageable)
    GetAverageVideoFrameRate()(MediaStream_MediaStream_averageVideoFrameRateable)
    GetAverageVideoPacketLossRate()(MediaStream_MediaStream_averageVideoPacketLossRateable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIsAudioForwardErrorCorrectionUsed()(*bool)
    GetLowFrameRateRatio()(MediaStream_MediaStream_lowFrameRateRatioable)
    GetLowVideoProcessingCapabilityRatio()(MediaStream_MediaStream_lowVideoProcessingCapabilityRatioable)
    GetMaxAudioNetworkJitter()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetMaxJitter()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetMaxPacketLossRate()(MediaStream_MediaStream_maxPacketLossRateable)
    GetMaxRatioOfConcealedSamples()(MediaStream_MediaStream_maxRatioOfConcealedSamplesable)
    GetMaxRoundTripTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetOdataType()(*string)
    GetPacketUtilization()(*int64)
    GetPostForwardErrorCorrectionPacketLossRate()(MediaStream_MediaStream_postForwardErrorCorrectionPacketLossRateable)
    GetRmsFreezeDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStreamDirection()(*MediaStreamDirection)
    GetStreamId()(*string)
    GetVideoCodec()(*VideoCodec)
    GetWasMediaBypassed()(*bool)
    SetAudioCodec(value *AudioCodec)()
    SetAverageAudioDegradation(value MediaStream_MediaStream_averageAudioDegradationable)()
    SetAverageAudioNetworkJitter(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetAverageBandwidthEstimate(value *int64)()
    SetAverageFreezeDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetAverageJitter(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetAveragePacketLossRate(value MediaStream_MediaStream_averagePacketLossRateable)()
    SetAverageRatioOfConcealedSamples(value MediaStream_MediaStream_averageRatioOfConcealedSamplesable)()
    SetAverageReceivedFrameRate(value MediaStream_MediaStream_averageReceivedFrameRateable)()
    SetAverageRoundTripTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetAverageVideoFrameLossPercentage(value MediaStream_MediaStream_averageVideoFrameLossPercentageable)()
    SetAverageVideoFrameRate(value MediaStream_MediaStream_averageVideoFrameRateable)()
    SetAverageVideoPacketLossRate(value MediaStream_MediaStream_averageVideoPacketLossRateable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIsAudioForwardErrorCorrectionUsed(value *bool)()
    SetLowFrameRateRatio(value MediaStream_MediaStream_lowFrameRateRatioable)()
    SetLowVideoProcessingCapabilityRatio(value MediaStream_MediaStream_lowVideoProcessingCapabilityRatioable)()
    SetMaxAudioNetworkJitter(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetMaxJitter(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetMaxPacketLossRate(value MediaStream_MediaStream_maxPacketLossRateable)()
    SetMaxRatioOfConcealedSamples(value MediaStream_MediaStream_maxRatioOfConcealedSamplesable)()
    SetMaxRoundTripTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetOdataType(value *string)()
    SetPacketUtilization(value *int64)()
    SetPostForwardErrorCorrectionPacketLossRate(value MediaStream_MediaStream_postForwardErrorCorrectionPacketLossRateable)()
    SetRmsFreezeDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStreamDirection(value *MediaStreamDirection)()
    SetStreamId(value *string)()
    SetVideoCodec(value *VideoCodec)()
    SetWasMediaBypassed(value *bool)()
}
