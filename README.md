# Go Race Condition with Closed Channel

This repository demonstrates a common race condition in Go programs involving a closed channel. The program creates multiple goroutines that attempt to receive values from a channel that's closed prematurely.

## Problem

The issue arises because multiple goroutines concurrently attempt to read from the `ch` channel after it has been closed. This leads to undefined behavior; the output will vary on each run.

## Solution

The solution involves buffering the channel and ensuring all goroutines complete before closing the channel.  This prevents race conditions by ensuring that the read operations on the channel happen in a predictable order.