// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: protoex.fbe
// Version: 1.3.0.0

package com.chronoxor.protoex;

public class Account implements Comparable<Object>
{
    public int id = 0;
    public String name = "";
    public StateEx state = StateEx.fromSet(java.util.EnumSet.of(StateEx.initialized.getEnum(), StateEx.bad.getEnum(), StateEx.sad.getEnum()));
    public Balance wallet = new Balance();
    public Balance asset = null;
    public java.util.ArrayList<Order> orders = new java.util.ArrayList<Order>();

    public static final long fbeTypeConst = 3;
    public long fbeType() { return fbeTypeConst; }

    public Account() {}

    public Account(int id, String name, StateEx state, Balance wallet, Balance asset, java.util.ArrayList<Order> orders)
    {
        this.id = id;
        this.name = name;
        this.state = state;
        this.wallet = wallet;
        this.asset = asset;
        this.orders = orders;
    }

    public Account(Account other)
    {
        this.id = other.id;
        this.name = other.name;
        this.state = other.state;
        this.wallet = other.wallet;
        this.asset = other.asset;
        this.orders = other.orders;
    }

    public Account clone()
    {
        // Serialize the struct to the FBE stream
        var writer = new com.chronoxor.protoex.fbe.AccountModel();
        writer.serialize(this);

        // Deserialize the struct from the FBE stream
        var reader = new com.chronoxor.protoex.fbe.AccountModel();
        reader.attach(writer.getBuffer());
        return reader.deserialize();
    }

    @Override
    public int compareTo(Object other)
    {
        if (other == null)
            return -1;

        if (!Account.class.isAssignableFrom(other.getClass()))
            return -1;

        final Account obj = (Account)other;

        int result = 0;
        result = Integer.compare(id, obj.id);
        if (result != 0)
            return result;
        return result;
    }

    @Override
    public boolean equals(Object other)
    {
        if (other == null)
            return false;

        if (!Account.class.isAssignableFrom(other.getClass()))
            return false;

        final Account obj = (Account)other;

        if (id != obj.id)
            return false;
        return true;
    }

    @Override
    public int hashCode()
    {
        int hash = 17;
        hash = hash * 31 + Integer.hashCode(id);
        return hash;
    }

    @Override
    public String toString()
    {
        var sb = new StringBuilder();
        sb.append("Account(");
        sb.append("id="); sb.append(id);
        sb.append(",name="); if (name != null) sb.append("\"").append(name).append("\""); else sb.append("null");
        sb.append(",state="); sb.append(state);
        sb.append(",wallet="); sb.append(wallet);
        sb.append(",asset="); if (asset != null) sb.append(asset); else sb.append("null");
        if (orders != null)
        {
            boolean first = true;
            sb.append(",orders=[").append(orders.size()).append("][");
            for (var item : orders)
            {
                sb.append(first ? "" : ",").append(item);
                first = false;
            }
            sb.append("]");
        }
        else
            sb.append(",orders=[0][]");
        sb.append(")");
        return sb.toString();
    }

    public String toJson() { return com.chronoxor.protoex.fbe.Json.getEngine().toJson(this); }
    public static Account fromJson(String json) { return com.chronoxor.protoex.fbe.Json.getEngine().fromJson(json, Account.class); }
}
