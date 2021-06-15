# Fullstack Code Challenge
A small stack consisting of:
- MySQL database (found in `/database`)
- Golang API (using grpc-gateway, found in `/api`) 
- React UI

It allows a user to see a list of classrooms and the students contained therein.

## Dependencies
- `git`
- `docker-compose` (for running these services locally)
- `protoc` (for translating `.proto` files into go code)

## Challenge
Users are interested in seeing additional metadata about each student. We've decided to include an additional field, "email" (assume that we can backfill this information easily -- i.e. we'll make them up during this exercise) on each student object. In order to display emails to the end user, we must:
- Update the database schema to include email address per student
  - Backfill email addresses for our existing students (we can make these up)
- Update the API layer to retrieve the emails and return them
- Update the UI to display the additional information

For this exercise, assume that your interviewer is an engineer familiar with this stack, and you have been assigned to a ticket to add this functionality while pair programming with them. Use any and all resources that you might find useful, and ask as many questions as you like.

## Resources
[gRPC](https://grpc.io/)
[grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)
[React](https://reactjs.org/)
