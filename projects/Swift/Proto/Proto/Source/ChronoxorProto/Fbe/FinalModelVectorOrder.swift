// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: proto.fbe
// Version: 1.3.0.0

import Foundation
import ChronoxorFbe

// Fast Binary Encoding Order vector final model
class FinalModelVectorOrder: FinalModel {
    var _buffer: Buffer = Buffer()
    var _offset: Int = 0

    private var _model: FinalModelOrder

    init(buffer: Buffer, offset: Int) {
        _buffer = buffer
        _offset = offset

        _model = FinalModelOrder(buffer: buffer, offset: offset)
    }

    // Get the allocation size
    func fbeAllocationSize(value values: Array<ChronoxorProto.Order>) -> Int {
        var size: Int = 4
        for i in 0..<values.count {
            size += _model.fbeAllocationSize(value: values[i])
        }

        return size
    }

    // Check if the vector is valid
    public func verify() -> Int {
        if _buffer.offset + fbeOffset + 4 > _buffer.size {
            return Int.max
        }

        let fbeVectorSize = Int(readUInt32(offset: fbeOffset))

        var size: Int = 4
        _model.fbeOffset = fbeOffset + 4
        var i = fbeVectorSize
        while i > 0 {
            let offset = _model.verify()
            if offset == Int.max { return Int.max }
            _model.fbeShift(size: offset)
            size += offset
            i -= 1
        }
        return size
    }

    public func get(values: inout Array<ChronoxorProto.Order>) -> Int {
        values.removeAll()

        if _buffer.offset + fbeOffset + 4 > _buffer.size {
            assertionFailure("Model is broken!")
            return 0
        }

        let fbeVectorSize = Int(readUInt32(offset: fbeOffset))
        if fbeVectorSize == 0 {
            return 4
        }

        //values.ensureCapacity(fbeVectorSize.toInt())

        var size: Int = 4
        var offset = Size()
        _model.fbeOffset = fbeOffset + 4
        for _ in 1...fbeVectorSize {
            offset.value = 0
            let value = _model.get(size: &offset)
            values.append(value)
            _model.fbeShift(size: offset.value)
            size += offset.value
        }
        return size
    }

    public func set(value values: Array<ChronoxorProto.Order>) throws -> Int {
        if _buffer.offset + fbeOffset + 4 > _buffer.size {
            assertionFailure("Model is broken!")
            return 0
        }

        write(offset: fbeOffset, value: UInt32(values.count))

        var size: Int = 4
        _model.fbeOffset = fbeOffset + 4
        for i in 0..<values.count {
            let offset = try _model.set(value: values[i])
            _model.fbeShift(size: offset)
            size += offset
        }
        return size
    }
}