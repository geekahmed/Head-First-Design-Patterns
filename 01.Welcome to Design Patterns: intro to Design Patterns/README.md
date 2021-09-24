# Chapter 1. Welcome to Design Patterns: intro to Design Patterns
## Notes
- Design patterns are typical solutions to commonly occurring problems in software design. They are like pre-made blueprints that you can customize to solve a recurring design problem in your code.
- Identify the aspects of your application that vary and separate them from what stays the same.
	- take the parts that vary and encapsulate them, so that later you can alter or extend the parts that vary without affecting those that don’t.
- Program to an interface, not an implementation.
- Favor composition over inheritance.
- The Strategy Pattern defines a family of algorithms, encapsulates each one, and makes them interchangeable. Strategy lets the algorithm vary independently from clients that use it.

## Examples

### Building In-Memory-Cache

Suppose you are building an In-Memory-Cache. Since it’s in memory, it has a limited size. Whenever it reaches its maximum size, some entries have to be evicted to free-up space. This can happen via several algorithms. Some of the popular algorithms are:
 - Least Recently Used (LRU): remove an entry that has been used least recently.
 - First In, First Out (FIFO): remove an entry that was created first.
 - Least Frequently Used (LFU): remove an entry that was least frequently used.

The problem is how to decouple our cache class from these algorithms so that we can change the algorithm at run time. Also, the cache class should not change when a new algorithm is being added.

### Payment method in an e-commerce app

In this example, the Strategy pattern is used to implement the various payment methods in an e-commerce application. After selecting a product to purchase, a customer picks a payment method: either Paypal or credit card.
