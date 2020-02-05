## Get Together

### Local deployment
#### Requirements:
* Go version 1.13.7+
* Git version 2.17.1+
* Node version 13.7.0+
* Docker version 18.09.9+
* Docker-compose version 1.23.2+

#### Clone repository:
```bash
$ git clone https://github.com/ilya-mezentsev/get_together.git && cd get_together
```

#### Prepare local workspace (in project directory):
```bash
$ bash prepare_workspace.sh $(pwd)
```

#### Check by running integration tests:
```bash
$ bash run.sh integration_tests
```
