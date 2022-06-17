# Factory Method Pattern

A factory is an object that is used to create other objects. In Factory pattern, we create object without exposing the creation logic to the client and refer to newly created object using a common interface.

A superclass specifies all standard and generic behavior (using pure virtual "placeholders" for creation steps), and then delegates the creation details to subclasses that are supplied by the client.

In a factory method pattern, a helper method (or function) is defined, to enable object creation without knowing the implementation class details. 
