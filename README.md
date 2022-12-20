# API Design

This project is setup for the discussion of API design as well as the demostrations for the different approaches. The code is initially generated by the https://openai.com/blog/chatgpt/, I'm also trying to explore the potential of how AI can assist engineers.

We gonna talk about Rest, OpenAPI, gRPC and GraphQL which are popular design or standard or framework that related to API design, meanwhile, they are prone to make people confused for the use case. 

Simple Put:
- REST is a software **architectural** and common communication **standard**.
- OpenAPI is a **specification**, which defines a standard, programming language-agnostic interface description for HTTP APIs.
- gRPC is a high performance Remote Procedure Call (RPC) **framework**.
- GraphQL is a **query language** for your API. 

## Rest
Representational state transfer (REST) is a software architectural style that describes a uniform interface between physically separate components, often across the Internet in a client-server architecture.

REST is the most common communication standard between computers over the internet, an API that follows the REST standard is called a RESTful API.

### Architectural Constraints
- Client–server architecture
- Statelessness
- Cacheability
- Layered system
- Code on demand (optional)
- Uniform interface
> refer to wikipedia(https://en.wikipedia.org/wiki/Representational_state_transfer#Architectural_constraints) for more details

### Applied to web services
- a base URI, such as http://api.example.com/
- standard HTTP methods (e.g., GET, POST, PUT, and DELETE)
- a media type that defines state transition data elements 

#### CRUD
- POST   --> CREATE *(non-idempotent)*
- GET    --> READ   *(idempotent)*
- PUT    --> UPDATE *(idempotent)*
- DELETE --> DELETE *(idempotent)*

#### Version
Versioning allows an implementation to provide backward compatibility so that if we introduce breaking changes from one version to another, consumers get enough time to move to the next version.

There are many ways to version an API. The most straightforward is to prefix the version before the resource on the URI. For instance:

- https://shopee.com/v1/products
- https://shopee.com/v2/products

## gRPC


## Open API


### Reference
- https://blog.bytebytego.com/p/why-is-restful-api-so-popular
- https://en.wikipedia.org/wiki/Representational_state_transfer
- [gRPC vs REST: Understanding gRPC, OpenAPI and REST and when to use them in API design](https://cloud.google.com/blog/products/api-management/understanding-grpc-openapi-and-rest-and-when-to-use-them)
- https://spec.openapis.org/oas/v3.1.0
