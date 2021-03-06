# Builder Pattern

Sometimes, object creation is not so straightforward. For example:
It might be necessary to have business rules to validate some parameters or derive some added attributes.

We might need some code to bring in efficiency—for example, retrieving an object from the cache rather than reading from the DB.

It might be necessary to have idempotency and thread safety in object creation. That is, multiple requests for object creation with the same parameters should give the same object.

The objects might have multiple constructor arguments (typically called telescopic constructors)

, and it is difficult to remember the order of parameters for the clients. Some of these parameters might be optional.
Such constructors frequently lead to bugs in client code.

Builder is a creational design pattern that lets you construct complex objects step by step instead of creating the object in a single call. The pattern allows you to produce different types and representations of an object using the same construction code and to run some additional construction steps before fetching the product.