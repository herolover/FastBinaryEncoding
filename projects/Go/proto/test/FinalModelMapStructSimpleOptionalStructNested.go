// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: fbe
// Version: 1.3.0.0

package test

import "errors"
import "../fbe"
import "../proto"

// Workaround for Go unused imports issue
var _ = errors.New
var _ = fbe.Version
var _ = proto.Version

// Fast Binary Encoding StructSimple->OptionalStructNested map final model
type FinalModelMapStructSimpleOptionalStructNested struct {
    // Final model buffer
    buffer *fbe.Buffer
    // Final model buffer offset
    offset int

    // Map key final model
    modelKey *FinalModelStructSimple
    // Map value final model
    modelValue *FinalModelOptionalStructNested
}

// Create a new StructSimple->OptionalStructNested map final model
func NewFinalModelMapStructSimpleOptionalStructNested(buffer *fbe.Buffer, offset int) *FinalModelMapStructSimpleOptionalStructNested {
    fbeResult := FinalModelMapStructSimpleOptionalStructNested{buffer: buffer, offset: offset}
    fbeResult.modelKey = NewFinalModelStructSimple(buffer, offset)
    fbeResult.modelValue = NewFinalModelOptionalStructNested(buffer, offset)
    return &fbeResult
}

// Get the allocation size
func (fm *FinalModelMapStructSimpleOptionalStructNested) FBEAllocationSize(values map[StructSimpleKey]struct{Key StructSimple; Value *StructNested}) int {
    size := 4
    for _, value := range values {
        size += fm.modelKey.FBEAllocationSize(&value.Key)
        size += fm.modelValue.FBEAllocationSize(value.Value)
    }
    return size
}

// Get the final offset
func (fm *FinalModelMapStructSimpleOptionalStructNested) FBEOffset() int { return fm.offset }
// Set the final offset
func (fm *FinalModelMapStructSimpleOptionalStructNested) SetFBEOffset(value int) { fm.offset = value }

// Shift the current final offset
func (fm *FinalModelMapStructSimpleOptionalStructNested) FBEShift(size int) { fm.offset += size }
// Unshift the current final offset
func (fm *FinalModelMapStructSimpleOptionalStructNested) FBEUnshift(size int) { fm.offset -= size }

// Check if the map is valid
func (fm *FinalModelMapStructSimpleOptionalStructNested) Verify() int {
    if (fm.buffer.Offset() + fm.FBEOffset() + 4) > fm.buffer.Size() {
        return fbe.MaxInt
    }

    fbeSetSize := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))

    size := 4
    fm.modelKey.SetFBEOffset(fm.FBEOffset() + 4)
    fm.modelValue.SetFBEOffset(fm.FBEOffset() + 4)
    for i := fbeSetSize; i > 0; i-- {
        offsetKey := fm.modelKey.Verify()
        if offsetKey == fbe.MaxInt {
            return fbe.MaxInt
        }
        fm.modelKey.FBEShift(offsetKey)
        fm.modelValue.FBEShift(offsetKey)
        size += offsetKey
        offsetValue := fm.modelValue.Verify()
        if offsetValue == fbe.MaxInt {
            return fbe.MaxInt
        }
        fm.modelKey.FBEShift(offsetValue)
        fm.modelValue.FBEShift(offsetValue)
        size += offsetValue
    }
    return size
}

// Get the map
func (fm *FinalModelMapStructSimpleOptionalStructNested) Get() (map[StructSimpleKey]struct{Key StructSimple; Value *StructNested}, int, error) {
    values := make(map[StructSimpleKey]struct{Key StructSimple; Value *StructNested})

    if (fm.buffer.Offset() + fm.FBEOffset() + 4) > fm.buffer.Size() {
        return values, 0, errors.New("model is broken")
    }

    fbeSetSize := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    if fbeSetSize == 0 {
        return values, 4, nil
    }

    size := 4
    fm.modelKey.SetFBEOffset(fm.FBEOffset() + 4)
    fm.modelValue.SetFBEOffset(fm.FBEOffset() + 4)
    for i := 0; i < fbeSetSize; i++ {
        key, offset, err := fm.modelKey.Get()
        if err != nil {
            return values, size, err
        }
        fm.modelKey.FBEShift(offset)
        fm.modelValue.FBEShift(offset)
        size += offset
        value, offset, err := fm.modelValue.Get()
        if err != nil {
            return values, size, err
        }
        fm.modelKey.FBEShift(offset)
        fm.modelValue.FBEShift(offset)
        size += offset
        values[key.Key()] = struct{Key StructSimple; Value *StructNested}{*key, value}
    }
    return values, size, nil
}

// Set the map
func (fm *FinalModelMapStructSimpleOptionalStructNested) Set(values map[StructSimpleKey]struct{Key StructSimple; Value *StructNested}) (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + 4) > fm.buffer.Size() {
        return 0, errors.New("model is broken")
    }

    fbe.WriteUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), uint32(len(values)))

    size := 4
    fm.modelKey.SetFBEOffset(fm.FBEOffset() + 4)
    fm.modelValue.SetFBEOffset(fm.FBEOffset() + 4)
    for _, value := range values {
        offsetKey, err := fm.modelKey.Set(&value.Key)
        if err != nil {
            return size, err
        }
        fm.modelKey.FBEShift(offsetKey)
        fm.modelValue.FBEShift(offsetKey)
        offsetValue, err := fm.modelValue.Set(value.Value)
        if err != nil {
            return size, err
        }
        fm.modelKey.FBEShift(offsetValue)
        fm.modelValue.FBEShift(offsetValue)
        size += offsetKey + offsetValue
    }
    return size, nil
}
