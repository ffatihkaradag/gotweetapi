# gotweetapi
simple go tweet api


## Get Tweets
GET /tweet HTTP/1.1
Host: gotweetapi.herokuapp.com

## Get Single Tweet
GET /tweet/1 HTTP/1.1
Host: gotweetapi.herokuapp.com


## Get Create Tweet
POST /tweet/ HTTP/1.1
Host: gotweetapi.herokuapp.com
Content-Type: application/json
Content-Length: 47

{
    "id": 2,
    "title": "second tweet"
}

## Create Tweet
POST /tweet/ HTTP/1.1
Host: gotweetapi.herokuapp.com
Content-Type: application/json
Content-Length: 47

{
    "id": 2,
    "title": "second tweet"
}

## Delete Tweet
DELETE /tweet/2 HTTP/1.1
Host: gotweetapi.herokuapp.com