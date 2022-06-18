# Abstract factory

With real-life problems, there are many related (family of) objects that need to be created together without specifying their concrete classes.

The client code should be able to call the creation methods of the abstract factory object instead of doing so directly with a constructor call. 
Since a factory corresponds to a single product variant, all its products will be compatible.

Client code works with factories and products only through their abstract interfaces. This lets the client code work with any product variants, created by the factory object. You just create a new concrete factory class and pass it to the client code.
