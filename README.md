# URL-Shortener

## Url Shortener

### Docker Build

To build the service using Docker Compose:

```
docker-compose -f docker-compose.yml up -d --build 
```

# Curl For Testing
#### Create Short URL
#### Generate a short URL for a long URL:

```
curl --location '127.0.0.1:8080/shorten' \
--header 'Content-Type: application/json' \
--data '{
    "long_url":"https://meet.google.com"
}'
```

# Get Short URL
#### Retrieve the long URL associated with a short URL:

```
curl --location 'localhost:8080/get_url/wTzkFNxt1'
```
Replace `wTzkFNxt1` with the actual short URL you want to retrieve the long URL for.

