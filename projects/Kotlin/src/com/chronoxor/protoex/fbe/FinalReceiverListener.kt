// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: protoex.fbe
// Version: 1.3.0.0

@file:Suppress("UnusedImport", "unused")

package com.chronoxor.protoex.fbe

// Fast Binary Encoding com.chronoxor.protoex final receiver listener
interface FinalReceiverListener : com.chronoxor.proto.fbe.FinalReceiverListener
{
    fun onReceive(value: com.chronoxor.protoex.OrderMessage) {}
    fun onReceive(value: com.chronoxor.protoex.BalanceMessage) {}
    fun onReceive(value: com.chronoxor.protoex.AccountMessage) {}
}
