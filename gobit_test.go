package main

import (
    "testing"
    "strings"
)

func TestFromBTC(t *testing.T) {
    usd := "USD"
    value := FromBTC(usd)
    if strings.Contains(value, usd) != true || strings.Contains(value, "BTC") != true {
        t.Errorf("FromBTC(%v) = %v...\n expected %v and BTC to both be present\n", usd, value, usd)
    }
}

func TestToBTC(t *testing.T) {
    btc := "1"
    usd := "USD"
    value := ToBTC(btc, usd)

    if strings.Contains(value, usd) != true || strings.Contains(value, "BTC") != true {
        t.Errorf("ToBTC(%v, %v) = %s...\n expected USD and BTC to both be present\n", btc, usd, value)
    }
}
