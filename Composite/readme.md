#Composite Pattern

##Intent
Compose objects into tree structures to represent part-whole hierarchies. Composite lets clients treat individual objects and compositions of objects uniformly.

##Motivation
Graphics applications like drawing editors and schematic capture systems let users build complex diagrams out of simple components. 
The user can group componentsto form larger components,which in turn can be grouped to form still larger components.
A simple implementation could define classes for graphical primitives such as Text and Lines plus other classes that act as containers for these primitives.
But there's a problem with this approach: Code that usesthese classes must treat primitive and container objects differently, even if most of the time the user treats them identically.
Having to distinguish these objects makes the application more complex. The Composite pattern describes how to use recursive composition so that clients don't have to make this distinction.
The greatest benefit of this approach is that you don’t need to care about the concrete classes of objects that compose the tree. 
You don’t need to know whether an object is a simple leaf or a sophisticated branch. 
You can treat them all the same via the common interface. When you call a method, the objects themselves pass the request down the tree.

##Applicability
Use the Composite pattern when
• you want to represent part-whole hierarchies of objects.
• you want clients to be able to ignore the difference between compositions of objects and individual objects.Clients will treat all objects in the composite structure uniformly.

##Structure

image

A typical Composite object structure might look like this:

##Known Uses
Examples of the Composite pattern can be found in almost all object-oriented systems.
The original View class of Smalltalk Model/View/Controller [KP88] was a Composite, and nearly every user interface toolkit or framework has followed in its steps, including ET++  and Interviews.
It's interesting to note that the original View of Model/View/Controller had a set of subviews; in other words, View was both the Component class and the Composite class.

##Pros and cons

### Pros
 You can work with complex tree structures more conveniently: use polymorphism and recursion to your advantage.
 Open/Closed Principle. You can introduce new element types into the app without breaking the existing code, which now works with the object tree.
 
### Cons
It might be difficult to provide a common interface for classes whose functionality differs too much. In certain scenarios, you’d need to overgeneralize the component interface, making it harder to comprehend.
