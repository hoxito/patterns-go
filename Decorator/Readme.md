#Decorator

##intent
Attach additional responsibilities to an object dynamically. Decorators provide a
flexible alternative to subclassing for extending functionality.


## Motivation
The decorator pattern allows the extension of a function of an existing object dynamically without the need to alter the original object. This is achieved by wrapping the original object and function into a new function.
Sometimes we want to add responsibilities to individual objects, not to an entire class.
One way to add responsibilities is with inheritance.But in most cases, the client can't control how and when to decorate the component with a functionality.
A more flexible approach is to enclose the component in another object that adds the functionality.
The enclosing object is called a decorator. The decorator conforms to the interface of the component it decorates so that its presence is transparent to the component's clients.
The decorator forwards requests to the component and may perform additional actions before or after forwarding.
Transparency lets you nest decorators recursively, thereby allowing an unlimited number of added responsibilities.

## Code Example
In this example, you can observe a function called "profileDecorator", that wraps a function that uses type T  as argument and R as a result implemented with generics (with the keyword any), this allows this wrapper to receive any function that uses an argument and returns a value and what this wrapper does to the callback function is a time measurement for debugging purposes for example.
Then we define a function called "duplicate" that sums a number twice just as an example function to be the callback function of the decorator.
This way we isolate any function from the wrapper, being able to extend its behavior fulfilling the solid principles.

## Applicability
Use Decorator
• to add responsibilities to individual objects dynamically and transparently, that is, without affecting other objects.
• for responsibilities that can be withdrawn.
• when extension by subclassing is impractical. Sometimes a large number of independent extensions are possible and would produce an explosion of subclasses to support every combination. Or a class definition maybe hidden or otherwise unavailable for subclassing.

## Structure

image

## Pros and Cons

### pros
 You can extend an object’s behavior without making a new subclass.
 You can add or remove responsibilities from an object at runtime.
 You can combine several behaviors by wrapping an object into multiple decorators.
 Single Responsibility Principle. You can divide a monolithic class that implements many possible variants of behavior into several smaller classes.

 ### Cons
  It’s hard to remove a specific wrapper from the wrappers stack.
 It’s hard to implement a decorator in such a way that its behavior doesn’t depend on the order in the decorators stack.
 The initial configuration code of layers might look pretty ugly.