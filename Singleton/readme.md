# Singleton

Sometimes, you may come across a need to restrict the number of objects of a specific time in the system. Singleton is the design pattern that restricts thecreation of objects to a single one. This might be useful, for example, when you want a single coordinator object across in multiple places of the code.

It should be noted that the singleton pattern is actually considered an anti-pattern because of the introduction of a global state. This causes a hidden coupling between components and can lead to difficult-to-debug situations. It should not be overused.