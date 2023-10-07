#### API

- ```docker build -t into_backend_api_go:0.1.0 .```
- ```docker run -itd --network=into-background_into_backend -p 8000:8000 -v /wkspace:/project --name into_backend_api_go into_backend_api_go:0.1.0```

#### Task

- ```docker build -t into_backend_task_go:0.1.0 .```
- ```docker run -itd --network=into-background_into_backend -v /wkspace:/project --name into_backend_task_go into_backend_task_go:0.1.0```
