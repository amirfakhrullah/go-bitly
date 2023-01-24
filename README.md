# GO-BITLY

## Description

Simple link-shortener service written in Go

## Functionalities

- Basic JWT authentication for users (login, signup and logout)
- Users can create, read, update and delete their shortened links, and also read their user info. The endpoints are protected via JWT middleware
- The shortened link endpoint is made public. The count for how many times the links are clicked is persisted

## Tech-stacks

- [Go](https://go.dev/)
- [Go Fiber](https://gofiber.io/)
- [Gorm](https://gorm.io/)

## License

License under the [MIT License](./LICENSE)
