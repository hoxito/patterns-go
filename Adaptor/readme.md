# Adaptor

Many times when you code, you come across situations where you have a new requirement, and a component that almost meets that requirement.
 A nonsoftware example of this situation is the power adapter: a three-legged plug from India can't be connected to a two-pronged outlet in the US. You need to use a power adapter to enable compatibility and use both entities.

 In theory, there are two ways of implementing this pattern:

 Object Adaptor: Here, the adaptor class has an instance of the Adaptee class and the adaptor method delegates the work to the wrapped instance.

 Class Adaptor: The adaptor is a mix-in (a class with multiple inheritance) and inherits from both the expected interface and the adaptee interface
