// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// Version: 1.3.0.0

package test

import "strings"
import "errors"
import "../fbe"
import "../proto"

// Workaround for Go unused imports issue
var _ = errors.New
var _ = fbe.Version
var _ = proto.Version

// FlagsSimple flags key
type FlagsSimpleKey int32

// Convert FlagsSimple flags key to string
func (k FlagsSimpleKey) String() string {
    return FlagsSimple(k).String()
}

// FlagsSimple flags
type FlagsSimple int32

// FlagsSimple flags values
//noinspection GoSnakeCaseUsage
const (
    FlagsSimple_FLAG_VALUE_0 = FlagsSimple(0x0)
    FlagsSimple_FLAG_VALUE_1 = FlagsSimple(0x1)
    FlagsSimple_FLAG_VALUE_2 = FlagsSimple(0x2)
    FlagsSimple_FLAG_VALUE_3 = FlagsSimple(0x4)
    FlagsSimple_FLAG_VALUE_4 = FlagsSimple(FlagsSimple_FLAG_VALUE_3)
    FlagsSimple_FLAG_VALUE_5 = FlagsSimple(FlagsSimple_FLAG_VALUE_1 | FlagsSimple_FLAG_VALUE_3)
)

// Create a new FlagsSimple flags
func NewFlagsSimple() *FlagsSimple {
    return new(FlagsSimple)
}

// Create a new FlagsSimple flags from the given flags
func NewFlagsSimpleFromValue(flags int32) *FlagsSimple {
    result := FlagsSimple(flags)
    return &result
}

// Is flags set?
func (f FlagsSimple) HasFlags(flags FlagsSimple) bool {
    return ((f & flags) != 0) && ((f & flags) == flags)
}

// Set flags
func (f *FlagsSimple) SetFlags(flags FlagsSimple) *FlagsSimple {
    *f |= flags
    return f
}

// Remove flags
func (f *FlagsSimple) RemoveFlags(flags FlagsSimple) *FlagsSimple {
    *f &^= flags
    return f
}

// Get the flags key
func (f FlagsSimple) Key() FlagsSimpleKey {
    return FlagsSimpleKey(f)
}

// Convert flags to optional
func (f *FlagsSimple) Optional() *FlagsSimple {
    return f
}

// Convert flags to string
//noinspection GoBoolExpressions
func (f FlagsSimple) String() string {
    var sb strings.Builder
    first := true
    if ((f & FlagsSimple_FLAG_VALUE_0) != 0) && ((f & FlagsSimple_FLAG_VALUE_0) == FlagsSimple_FLAG_VALUE_0) {
        if first {
            first = false
        } else {
            sb.WriteRune('|')
        }
        sb.WriteString("FLAG_VALUE_0")
    }
    if ((f & FlagsSimple_FLAG_VALUE_1) != 0) && ((f & FlagsSimple_FLAG_VALUE_1) == FlagsSimple_FLAG_VALUE_1) {
        if first {
            first = false
        } else {
            sb.WriteRune('|')
        }
        sb.WriteString("FLAG_VALUE_1")
    }
    if ((f & FlagsSimple_FLAG_VALUE_2) != 0) && ((f & FlagsSimple_FLAG_VALUE_2) == FlagsSimple_FLAG_VALUE_2) {
        if first {
            first = false
        } else {
            sb.WriteRune('|')
        }
        sb.WriteString("FLAG_VALUE_2")
    }
    if ((f & FlagsSimple_FLAG_VALUE_3) != 0) && ((f & FlagsSimple_FLAG_VALUE_3) == FlagsSimple_FLAG_VALUE_3) {
        if first {
            first = false
        } else {
            sb.WriteRune('|')
        }
        sb.WriteString("FLAG_VALUE_3")
    }
    if ((f & FlagsSimple_FLAG_VALUE_4) != 0) && ((f & FlagsSimple_FLAG_VALUE_4) == FlagsSimple_FLAG_VALUE_4) {
        if first {
            first = false
        } else {
            sb.WriteRune('|')
        }
        sb.WriteString("FLAG_VALUE_4")
    }
    if ((f & FlagsSimple_FLAG_VALUE_5) != 0) && ((f & FlagsSimple_FLAG_VALUE_5) == FlagsSimple_FLAG_VALUE_5) {
        if first {
            first = false
        } else {
            sb.WriteRune('|')
        }
        sb.WriteString("FLAG_VALUE_5")
    }
    return sb.String()
}

// Convert flags to JSON
func (f FlagsSimple) MarshalJSON() ([]byte, error) {
    value := int32(f)
    return fbe.Json.Marshal(&value)
}

// Convert JSON to flags
func (f *FlagsSimple) UnmarshalJSON(buffer []byte) error {
    var result int32
    err := fbe.Json.Unmarshal(buffer, &result)
    if err != nil {
        return err
    }
    *f = FlagsSimple(result)
    return nil
}
