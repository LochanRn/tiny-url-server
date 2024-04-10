# tiny-url-server
URL shortner service 


###### You can start the server using the command 
```BASH
go run cmd/main.go
```

#### CREATE API: 

```BASH
curl --location 'http://127.0.0.1:2112/v1/tinyurl' --header 'Content-Type: application/json' \
--data '{
    "url": "https://www.google.com/"
}'
```

Response 

```JSON
{
    "creationTimestamp": "2024-04-08T13:02:53.504+05:30",
    "tinyurl": "http://127.0.0.1:2112/v1/tinyurl/3417810",
    "url": "https://www.google.com/"
}
```

#### Redirect API: 
```BASH
curl --location 'http://127.0.0.1:2112/v1/tinyurl/3417810'
```

#### Get the Domains with highest count API: 

```BASH
curl --location 'http://127.0.0.1:2112/v1/maxdomainsabbrev'
```

Response 

```JSON
{
    "domains": [
        {
            "count": 2,
            "domains": "https://www.google.com/"
        },
        {
            "count": 2,
            "domains": "https://www.udemy.com/"
        },
        {
            "count": 1,
            "domains": "https://www.spectrocloud.com/"
        }
    ]
}
```