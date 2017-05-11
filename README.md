# Bitcoin price ticker

## Task

----
Develop a simple bitcoin realtime ticker in *** dev lang that shows in a single text line bitcoin value in Euro based on auto selection between 2 or more bitcoin price feeds and 2 and more currency rates feeds. Selection can be best value for user.
Feel free to use any feed you like. Completed code we expect to see in GitHub or BitBucket.

Example output line:

BTC/USD: 600   EUR/USD: 1.05   BTC/EUR: 550 Active sources: BTC/USD (3 of 3)  EUR/USD (2 of 3)

Technical requirements.
The expected solution must handle feeds in parallel, consistently, work with isolated state, care of resources and shutdown the subscriber gracefully. It must also handle network latencies, frozen requests and corrupted data.

Business requirements.
Aggregate unlimited number of feeds.
Given every feed data expiration time, the solution must not use obsolete data in comparisons and calculations.
----

Hints:
How responsive is the solution to
1) the new price available in the feed?
2) program shutdown?

Is it possible to attach multiple tick viewers?
How does one of the broken feed affect ticker behaviour?
Do you understand the concept of Bid/Ask for the price?
How old data is you final ticker BTC/EUR giving out to end user?
What is the theoretical limit of feed sources in your solution?
Did you run build with -race detection?
Is it possible to test the solution without real access to the feed sources?