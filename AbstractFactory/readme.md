# Abstract factory

With real-life problems, there are many related (family of) objects that need to be created together without specifying their concrete classes.

The client code should be able to call the creation methods of the abstract factory object instead of doing so directly with a constructor call (new operator). Since a factory corresponds to a single product variant, all its products will be compatible.
