# Go hexagonal architecture example

This is an example illustrating how to structure code in go using
hexagonal architecture

Reference :
<https://youtu.be/rQnTtQZGpg8?list=RDCMUCYqCZOwHbnPwyjawKfE21wg>

Here are some articles about hexagonal architecture
* <https://blog.ndepend.com/hexagonal-architecture/>
* <https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)>
* <https://medium.com/idealo-tech-blog/hexagonal-ports-adapters-architecture-e3617bcf00a0>
* <https://dzone.com/articles/hexagonal-architecture-what-is-it-and-how-does-it>


### Folder structure

* api -> contains api handler logic
* repository -> logic to intact with different storage mechanisms (mongo, redis etc)
  - memory -> implementation of repository logic using in-memory store (hashmap)
* serializer -> logic to encode/decode objects to bytes
  - json -> serializer implementation using json
  - msgpack -> serializer implementation using msgpack
* shortner -> app and domain logic for shortner
  - logic.go -> business logic written as a service
  - model.go -> struct
  - repository.go -> interface to deal with repository
  - serializer.go -> interface to encode/decode Redirect struct
  - service.go -> interface for business logic

## Build and run

Run the following command to run the app

```
go run ./...
```
