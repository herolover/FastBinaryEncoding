// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// Version: 1.1.0.0

package test

import "../fbe"
import "../proto"

// Workaround for Go unused imports issue
var _ = fbe.Version
var _ = proto.Version

// Fast Binary Encoding StructNested final model
type FinalModelStructNested struct {
    buffer *fbe.Buffer  // Final model buffer
    offset int          // Final model buffer offset

    FinalModelStructOptional
    F1000 *FinalModelEnumSimple
    F1001 *FinalModelOptionalEnumSimple
    F1002 *FinalModelEnumTyped
    F1003 *FinalModelOptionalEnumTyped
    F1004 *FinalModelFlagsSimple
    F1005 *FinalModelOptionalFlagsSimple
    F1006 *FinalModelFlagsTyped
    F1007 *FinalModelOptionalFlagsTyped
    F1008 *FinalModelStructSimple
    F1009 *FinalModelOptionalStructSimple
    F1010 *FinalModelStructOptional
    F1011 *FinalModelOptionalStructOptional
}

// Create a new StructNested final model
func NewFinalModelStructNested(buffer *fbe.Buffer, offset int) *FinalModelStructNested {
    fbeResult := FinalModelStructNested{buffer: buffer, offset: offset}
    fbeResult.FinalModelStructOptional = *NewFinalModelStructOptional(buffer, 0)
    fbeResult.F1000 = NewFinalModelEnumSimple(buffer, 0)
    fbeResult.F1001 = NewFinalModelOptionalEnumSimple(buffer, 0)
    fbeResult.F1002 = NewFinalModelEnumTyped(buffer, 0)
    fbeResult.F1003 = NewFinalModelOptionalEnumTyped(buffer, 0)
    fbeResult.F1004 = NewFinalModelFlagsSimple(buffer, 0)
    fbeResult.F1005 = NewFinalModelOptionalFlagsSimple(buffer, 0)
    fbeResult.F1006 = NewFinalModelFlagsTyped(buffer, 0)
    fbeResult.F1007 = NewFinalModelOptionalFlagsTyped(buffer, 0)
    fbeResult.F1008 = NewFinalModelStructSimple(buffer, 0)
    fbeResult.F1009 = NewFinalModelOptionalStructSimple(buffer, 0)
    fbeResult.F1010 = NewFinalModelStructOptional(buffer, 0)
    fbeResult.F1011 = NewFinalModelOptionalStructOptional(buffer, 0)
    return &fbeResult
}

// Get the allocation size
func (fm *FinalModelStructNested) FBEAllocationSize(fbeValue *StructNested) int {
    fbeResult := 0 +
        fm.FinalModelStructOptional.FBEAllocationSize(&fbeValue.StructOptional) + 
        fm.F1000.FBEAllocationSize(&fbeValue.F1000) +
        fm.F1001.FBEAllocationSize(fbeValue.F1001) +
        fm.F1002.FBEAllocationSize(&fbeValue.F1002) +
        fm.F1003.FBEAllocationSize(fbeValue.F1003) +
        fm.F1004.FBEAllocationSize(&fbeValue.F1004) +
        fm.F1005.FBEAllocationSize(fbeValue.F1005) +
        fm.F1006.FBEAllocationSize(&fbeValue.F1006) +
        fm.F1007.FBEAllocationSize(fbeValue.F1007) +
        fm.F1008.FBEAllocationSize(&fbeValue.F1008) +
        fm.F1009.FBEAllocationSize(fbeValue.F1009) +
        fm.F1010.FBEAllocationSize(&fbeValue.F1010) +
        fm.F1011.FBEAllocationSize(fbeValue.F1011) +
        0
    return fbeResult
}

// Get the final size
func (fm *FinalModelStructNested) FBESize() int { return 0 }

// Get the final extra size
func (fm *FinalModelStructNested) FBEExtra() int { return 0 }

// Get the final type
func (fm *FinalModelStructNested) FBEType() int { return 112 }

// Get the final offset
func (fm *FinalModelStructNested) FBEOffset() int { return fm.offset }
// Set the final offset
func (fm *FinalModelStructNested) SetFBEOffset(value int) { fm.offset = value }

// Shift the current final offset
func (fm *FinalModelStructNested) FBEShift(size int) { fm.offset += size }
// Unshift the current final offset
func (fm *FinalModelStructNested) FBEUnshift(size int) { fm.offset -= size }

// Check if the struct value is valid
func (fm *FinalModelStructNested) Verify() int {
    fm.buffer.Shift(fm.FBEOffset())
    fbeResult := fm.VerifyFields()
    fm.buffer.Unshift(fm.FBEOffset())
    return fbeResult
}

// Check if the struct fields are valid
func (fm *FinalModelStructNested) VerifyFields() int {
    fbeCurrentOffset := 0
    fbeFieldSize := 0


    fm.FinalModelStructOptional.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.FinalModelStructOptional.VerifyFields(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F1000.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F1000.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F1001.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F1001.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F1002.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F1002.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F1003.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F1003.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F1004.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F1004.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F1005.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F1005.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F1006.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F1006.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F1007.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F1007.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F1008.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F1008.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F1009.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F1009.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F1010.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F1010.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F1011.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F1011.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    return fbeCurrentOffset
}

// Get the struct value
func (fm *FinalModelStructNested) Get() (*StructNested, int, error) {
    fbeResult := NewStructNested()
    fbeSize, err := fm.GetValue(fbeResult)
    return fbeResult, fbeSize, err
}

// Get the struct value by the given pointer
func (fm *FinalModelStructNested) GetValue(fbeValue *StructNested) (int, error) {
    fm.buffer.Shift(fm.FBEOffset())
    fbeSize, err := fm.GetFields(fbeValue)
    fm.buffer.Unshift(fm.FBEOffset())
    return fbeSize, err
}

// Get the struct fields values
func (fm *FinalModelStructNested) GetFields(fbeValue *StructNested) (int, error) {
    var err error = nil
    fbeCurrentOffset := 0
    fbeCurrentSize := 0
    fbeFieldSize := 0

    fm.FinalModelStructOptional.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.FinalModelStructOptional.GetFields(&fbeValue.StructOptional); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F1000.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F1000.GetValue(&fbeValue.F1000); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F1001.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F1001, fbeFieldSize, err = fm.F1001.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F1002.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F1002.GetValue(&fbeValue.F1002); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F1003.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F1003, fbeFieldSize, err = fm.F1003.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F1004.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F1004.GetValue(&fbeValue.F1004); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F1005.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F1005, fbeFieldSize, err = fm.F1005.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F1006.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F1006.GetValue(&fbeValue.F1006); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F1007.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F1007, fbeFieldSize, err = fm.F1007.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F1008.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F1008.GetValue(&fbeValue.F1008); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F1009.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F1009, fbeFieldSize, err = fm.F1009.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F1010.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F1010.GetValue(&fbeValue.F1010); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F1011.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F1011, fbeFieldSize, err = fm.F1011.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    return fbeCurrentSize, nil
}

// Set the struct value
func (fm *FinalModelStructNested) Set(fbeValue *StructNested) (int, error) {
    fm.buffer.Shift(fm.FBEOffset())
    fbeResult, err := fm.SetFields(fbeValue)
    fm.buffer.Unshift(fm.FBEOffset())
    return fbeResult, err
}

// Set the struct fields values
func (fm *FinalModelStructNested) SetFields(fbeValue *StructNested) (int, error) {
    var err error = nil
    fbeCurrentOffset := 0
    fbeCurrentSize := 0
    fbeFieldSize := 0

    fm.FinalModelStructOptional.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.FinalModelStructOptional.SetFields(&fbeValue.StructOptional); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F1000.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F1000.Set(&fbeValue.F1000); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F1001.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F1001.Set(fbeValue.F1001); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F1002.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F1002.Set(&fbeValue.F1002); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F1003.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F1003.Set(fbeValue.F1003); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F1004.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F1004.Set(&fbeValue.F1004); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F1005.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F1005.Set(fbeValue.F1005); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F1006.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F1006.Set(&fbeValue.F1006); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F1007.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F1007.Set(fbeValue.F1007); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F1008.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F1008.Set(&fbeValue.F1008); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F1009.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F1009.Set(fbeValue.F1009); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F1010.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F1010.Set(&fbeValue.F1010); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F1011.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F1011.Set(fbeValue.F1011); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    return fbeCurrentSize, nil
}
