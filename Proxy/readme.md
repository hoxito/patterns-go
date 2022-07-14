Every Pattern´s Readme must follow the next structure:

# Facade Pattern

## Description
When you need to perform several high cost operations with a database for example and you must 
The Proxy pattern suggests that you create a new proxy class with the same interface as an original service object. Then you update your app so that it passes the proxy object to all of the original object’s clients. Upon receiving a request from a client, the proxy creates a real service object and delegates all the work to it.
## Motivation
What problem solves the pattern
## Code Example
The example in this repository
## Applicability
Lazy initialization (virtual proxy). This is when you have a heavyweight service object that wastes system resources by being always up, even though you only need it from time to time.
Access control (protection proxy). This is when you want only specific clients to be able to use the service object; for instance, when your objects are crucial parts of an operating system and clients are various launched applications (including malicious ones).
Local execution of a remote service (remote proxy). This is when the service object is located on a remote server.
 Logging requests (logging proxy). This is when you want to keep a history of requests to the service object.
  Caching request results (caching proxy). This is when you need to cache results of client requests and manage the life cycle of this cache, especially if results are quite large.
   Smart reference. This is when you need to be able to dismiss a heavyweight object once there are no clients that use it.
## Structure
the pattern´s structure 
## Pros and Cons

### Pros
 You can control the service object without clients knowing about it.
 You can manage the lifecycle of the service object when clients don’t care about it.
 The proxy works even if the service object isn’t ready or is not available.
 Open/Closed Principle. You can introduce new proxies without changing the service or clients.
### Cons
 The code may become more complicated since you need to introduce a lot of new classes.
 The response from the service might get delayed.