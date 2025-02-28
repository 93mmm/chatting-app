# chatting-app

realtime web chatting app

### Requirements:
* Docker latest

### Tech stack:
* Go
* docker, docker-compose
* PostgreSQL
* REST-api
* API-testing

### Run application:
1. Configure .env file using .env.template file
2. ```make run_app```

### Run tests:
> [!NOTE]
> The database will be cleared
1. Run application
2. ```make run_tests```

### TODO in project:
- [ ] Improve logging in entire project
- [ ] Write Unit-tests
- [ ] Write Swagger documentation
- [ ] Implement all api methods in db service
- [ ] Implement all api methods in main backend server
- [ ] Implement authentication in main backend server, add session management with JWT
- [ ] Rewrite api between main backend and db service in gRPC
- [ ] Write an HTMX+Go frontend server
- [ ] Add WebSockets (only for real time messaging)
- [ ] Use Redis to cache frequent requests

### Development Standards that I plan to follow:
- Logging 
- Unit-testing, Api-testing
- Input validation
- Error handling
- Caching
- Encrypting
