
# Chain Of Responsibility Pattern

## Description
Chain of Responsibility is a behavioral design pattern that lets you pass requests along a chain of handlers. Upon receiving a request, each handler decides either to process the request or to pass it to the next handler in the chain.

## Motivation
Chain of Responsibility relies on transforming particular behaviors into stand-alone objects called handlers for action processing. If you have to check lots of things before executing some action, you should transform them into this handlers. Each linked handler has a field for storing a reference to the next handler in the chain. In addition to processing a request, handlers pass the request further along the chain. The request travels along the chain until all handlers have had a chance to process it and one final action is finally executed.
## Code Example

## Applicability

## Structure

## Pros and Cons

### Pros
 You can control the order of request handling.
 Single Responsibility Principle. You can decouple classes that invoke operations from classes that perform operations.
 Open/Closed Principle. You can introduce new handlers into the app without breaking the existing client code.
### Cons
 Some requests may end up unhandled.