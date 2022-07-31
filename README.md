# Consistent Hash load balancer

This project is a sample proof of concept on using consistent hash as a load balancer.

## Excellent articles / videos about consistent hashing

* [A Guide to Consistent Hashing](https://www.toptal.com/big-data/consistent-hashing#:~:text=according%20to%20Wikipedia).-,Consistent%20Hashing%20is%20a%20distributed%20hashing%20scheme%20that%20operates%20independently,without%20affecting%20the%20overall%20system.)
* [A Brief Introduction to Consistent Hash - YouTube](https://www.youtube.com/watch?v=tHEyzVbl4bg)

This project is a simple implementation of a look aside load balancer using consistent hash.
![look aside](docs/assets/lookaside.png)

## The use case

Our use case is to be able to route a key to a particular node in order to make use of the local cache held in that node.
The process that needs to be performed by this node for this particular key must be in order.  

As an example: __Calculating an account balance__

Using Consistent hash, we can direct the request from an __account number__ to a specific node.  

