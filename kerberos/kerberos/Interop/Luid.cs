﻿using System;
using System.Runtime.InteropServices;

namespace BloodHound.lib.Interop
{
    [StructLayout(LayoutKind.Sequential)]
    public struct LUID
    {
        public UInt32 LowPart;
        public Int32 HighPart;

        public LUID(UInt64 value)
        {
            LowPart = (UInt32)(value & 0xffffffffL);
            HighPart = (Int32)(value >> 32);
        }

        public LUID(LUID value)
        {
            LowPart = value.LowPart;
            HighPart = value.HighPart;
        }

        public LUID(string value)
        {
            if (System.Text.RegularExpressions.Regex.IsMatch(value, @"^0x[0-9A-Fa-f]+$"))
            {
                // if the passed LUID string is of form 0xABC123
                UInt64 uintVal = Convert.ToUInt64(value, 16);
                LowPart = (UInt32)(uintVal & 0xffffffffL);
                HighPart = (Int32)(uintVal >> 32);
            }
            else if (System.Text.RegularExpressions.Regex.IsMatch(value, @"^\d+$"))
            {
                // if the passed LUID string is a decimal form
                UInt64 uintVal = UInt64.Parse(value);
                LowPart = (UInt32)(uintVal & 0xffffffffL);
                HighPart = (Int32)(uintVal >> 32);
            }
            else
            {
                ArgumentException argEx = new ArgumentException("Passed LUID string value is not in a hex or decimal form", value);
                throw argEx;
            }
        }

        public override int GetHashCode()
        {
            UInt64 Value = ((UInt64)HighPart << 32) + LowPart;
            return Value.GetHashCode();
        }

        public override bool Equals(object obj)
        {
            return obj is LUID && (((ulong)this) == (LUID)obj);
        }

        public override string ToString()
        {
            UInt64 Value = ((UInt64)HighPart << 32) + LowPart;
            return $"0x{Value:x}";
        }

        public static bool operator ==(LUID x, LUID y)
        {
            return (x == ((ulong)y));
        }

        public static bool operator !=(LUID x, LUID y)
        {
            return (x != ((ulong)y));
        }

        public static implicit operator ulong(LUID luid)
        {
            // enable casting to a ulong
            UInt64 Value = ((UInt64)luid.HighPart << 32);
            return Value + luid.LowPart;
        }
    }
}
