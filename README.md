# Wait For Elasticsearch

This is a container that waits for your Elasticsearch client to start up and ping the elasticsearch without being stopping.

Idea and base code for this is influenced by a repo that does the same try wait for MySQL [here](https://github.com/jimmysawczuk/wait-for-mysql).

## Run With Container

```
docker run --network host jongschneider/wait-for-elasticsearch '[connectionstring]'
```

## Pushing to Docker Hub

[Docker Hub](https://hub.docker.com/r/jongschneider/wait-for-elasticsearch)

execute the following command in your terminal:
```bash
make
```
