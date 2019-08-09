// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// Version: 1.3.0.0

@file:Suppress("UnusedImport", "unused")

package com.chronoxor.test

@Suppress("MemberVisibilityCanBePrivate", "RemoveRedundantCallsOfConversionMethods")
open class StructHashEx : Comparable<Any?>
{
    var f1: java.util.HashMap<StructSimple, StructNested> = java.util.HashMap()
    var f2: java.util.HashMap<StructSimple, StructNested?> = java.util.HashMap()

    @Transient open var fbeType: Long = 142

    constructor()

    constructor(f1: java.util.HashMap<StructSimple, StructNested>, f2: java.util.HashMap<StructSimple, StructNested?>)
    {
        this.f1 = f1
        this.f2 = f2
    }

    @Suppress("UNUSED_PARAMETER")
    constructor(other: StructHashEx)
    {
        this.f1 = other.f1
        this.f2 = other.f2
    }

    open fun clone(): StructHashEx
    {
        // Serialize the struct to the FBE stream
        val writer = com.chronoxor.test.fbe.StructHashExModel()
        writer.serialize(this)

        // Deserialize the struct from the FBE stream
        val reader = com.chronoxor.test.fbe.StructHashExModel()
        reader.attach(writer.buffer)
        return reader.deserialize()
    }

    override fun compareTo(other: Any?): Int
    {
        if (other == null)
            return -1

        if (!StructHashEx::class.java.isAssignableFrom(other.javaClass))
            return -1

        @Suppress("UNUSED_VARIABLE")
        val obj = other as StructHashEx? ?: return -1

        @Suppress("VARIABLE_WITH_REDUNDANT_INITIALIZER", "CanBeVal", "UnnecessaryVariable")
        var result = 0
        return result
    }

    override fun equals(other: Any?): Boolean
    {
        if (other == null)
            return false

        if (!StructHashEx::class.java.isAssignableFrom(other.javaClass))
            return false

        @Suppress("UNUSED_VARIABLE")
        val obj = other as StructHashEx? ?: return false

        return true
    }

    override fun hashCode(): Int
    {
        @Suppress("CanBeVal", "UnnecessaryVariable")
        var hash = 17
        return hash
    }

    override fun toString(): String
    {
        val sb = StringBuilder()
        sb.append("StructHashEx(")
        @Suppress("ConstantConditionIf")
        if (true)
        {
            var first = true
            sb.append("f1=[").append(f1.size).append("][{")
            for (item in f1.entries)
            {
                sb.append(if (first) "" else ",").append(item.key)
                sb.append("->")
                sb.append(item.value)
                first = false
            }
            sb.append("}]")
        }
        @Suppress("ConstantConditionIf")
        if (true)
        {
            var first = true
            sb.append(",f2=[").append(f2.size).append("][{")
            for (item in f2.entries)
            {
                sb.append(if (first) "" else ",").append(item.key)
                sb.append("->")
                if (item.value != null) sb.append(item.value!!); else sb.append("null")
                first = false
            }
            sb.append("}]")
        }
        sb.append(")")
        return sb.toString()
    }

    open fun toJson(): String = com.chronoxor.test.fbe.Json.engine.toJson(this)

    companion object
    {
        const val fbeTypeConst: Long = 142
        fun fromJson(json: String): StructHashEx = com.chronoxor.test.fbe.Json.engine.fromJson(json, StructHashEx::class.java)
    }
}