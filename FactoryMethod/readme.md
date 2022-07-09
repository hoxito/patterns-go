# Factory Method Pattern

# Intent
Define an interface for creating an object, but let subclasses decide which class to instantiate. Factory Method lets a class defer instantiation to subclasses

## Motivation
Frameworks use abstract classes to define and maintain relationships between objects. A framework is often responsible for creating these objects as well. 

## Applicability
Use the Factory Method pattern when
• a class can't anticipate the class of objects it must create.
• a class wants its subclasses to specify the objectsit creates.
• classes delegate responsibility to one of several helper subclasses, and you
want to localize the knowledge of which helper subclass is the delegate.

## Structure



## Description
A factory is an object that is used to create other objects. In Factory pattern, we create object without exposing the creation logic to the client and refer to newly created object using a common interface.

A superclass specifies all standard and generic behavior (using pure virtual "placeholders" for creation steps), and then delegates the creation details to subclasses that are supplied by the client.

In a factory method pattern, a helper method (or function) is defined, to enable object creation without knowing the implementation class details. 

In the example we can see that the helper method is NewCharacter that does not need to know what type of character is being created.

## Pros and Cons

### Pros

 You avoid tight coupling between the creator and the concrete products.
 Single Responsibility Principle. You can move the product creation code into one place in the program, making the code easier to support.
 Open/Closed Principle. You can introduce new types of products into the program without breaking existing client code.

 ### Cons

  The code may become more complicated since you need to introduce a lot of new subclasses to implement the pattern. The best case scenario is when you’re introducing the pattern into an existing hierarchy of creator classes.