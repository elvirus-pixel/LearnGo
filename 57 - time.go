// Go offers extensive support for times and durations; here are some examples.
package main

import (
    "fmt"
    "time"
)

func main() {
    p := fmt.Println

    // We’ll start by getting the current time.
    now := time.Now()
    p(now)

    // You can build a `time` struct by providing the year, month, day, etc. Times are always associated with a `Location`, i.e. time zone.
    then := time.Date(
        2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
    p(then)

    // You can extract the various components of the time value as expected.
    p(then.Year())
    p(then.Month())
    p(then.Day())
    p(then.Hour())
    p(then.Minute())
    p(then.Second())
    p(then.Nanosecond())
    p(then.Location())

    // The Monday-Sunday `Weekday` is also available.
    p(then.Weekday())

    // These methods compare two times, testing if the first occurs before, after, or at the same time as the second, respectively.
    p(then.Before(now))
    p(then.After(now))
    p(then.Equal(now))

    // The `Sub` methods returns a `Duration` representing the interval between two times.
    diff := now.Sub(then)
    p(diff)

    // We can compute the length of the duration in various units.0
    p(diff.Hours())
    p(diff.Minutes())
    p(diff.Seconds())
    p(diff.Nanoseconds())

    // You can use `Add` to advance a time by a given duration, or with a `-` to move backwards by a duration.
    p(then.Add(diff))
    p(then.Add(-diff))
}

/*
$ go run '57 - time.go' 
2026-05-27 11:24:22.710421154 -0400 EDT m=+0.000011384
2009-11-17 20:34:58.651387237 +0000 UTC
2009
November
17
20
34
58
651387237
UTC
Tuesday
true
false
false
144834h49m24.059033917s
144834.82334973165
8.690089400983898e+06
5.2140536405903393e+08
521405364059033917
2026-05-27 15:24:22.710421154 +0000 UTC
1993-05-11 01:45:34.59235332 +0000 UTC
*/