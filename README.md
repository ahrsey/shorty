# Shorty - Url shortener'ed

## TODO
- [x] Setup Dockerfile
- [x] Setup tests
- [x] Decide which database to use (relational)
- [x] Add endpoint for creating a shortened link
- [x] Add endpoint for fetching a shortened link
- [x] Look into hosting options
- [x] Create Makefile
- [ ] Add help to print out Makefile commands and go over the commands I've
      added to see what they do
- [ ] Setup docker compose with postgres
- [ ] Setup repo and CI
- [ ] Setup dev docker with hot module reload (not this one, but I don't
      remember the correct name now)
- [ ] Update readme
- [ ] Add changelog
- [ ] Setup logging
- [ ] Setup monitoring
- [ ] Add endpoint for healthchecking
- [ ] Setup authentication

## Roadmap
- [ ] User flow with frontend ui

## Hosting
- [hetzner](https://www.hetzner.com/cloud/) Hosted shared server 
- [fly.io](https://fly.io/docs/about/pricing/) Cheap free tier supported as a service
- [back4app](https://www.back4app.com/) Cheap free tier supported as a service

## Makefile
https://gist.github.com/alexedwards/3b40775846535d0014ab1ff477e4a568
https://news.ycombinator.com/item?id=21735176
https://waltercode.medium.com/building-and-pushing-images-using-docker-and-makefiles-2d520b17f97e
https://earthly.dev/blog/docker-and-makefiles/


## Domain names
Short domain name search
https://micro.domains/

## Notes
https://developer.hashicorp.com/terraform/tutorials/docker-get-started/docker-build
https://levelup.gitconnected.com/dockerized-crud-restful-api-with-go-gorm-jwt-postgresql-mysql-and-testing-61d731430bd8 

Go over this repo to see if there are any useful things I can steal
https://github.com/eldad87/go-boilerplate

Look into copying binary into container if the only changes are in the code
and not in the infrastructure. I could have this setup on a watch.
