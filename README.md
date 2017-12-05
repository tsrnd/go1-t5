# GoWeb5 Team
Developers: 
* Mr. Ut
* Mr. Phuc
* Mr. Duoc
* Mr. Trung
* Mr. Hieu
* Mr. Vu

## Getting started:
* download [docker-compose](https://docs.docker.com/compose/install/) if not already installed
  Then run the following commands:

```bash
$ cd {your_source_code_folder}
$ git clone git@github.com:tsrnd/goweb5.git goweb5
$ cd goweb5
$ cp .env.example .env
$ docker-compose up
```
Then you can open the frontend at localhost:5001, admin localhost:5002 at  and the RESTful GoLang API at localhost:5000

If you want to build each images run:
```bash
$ docker build ./api --build-arg
$ docker build ./frontend --build-arg
$ docker build ./admin --build-arg
$ docker build ./db
```
