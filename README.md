# Storyly

- For installation, you have to create *conf.yml* file. This file stores
database credentials. I created *story-conf-sample.yml* file as a sample.
You can use that file. After create the config file, you have to change *viper.SetConfigName("story-conf-sample")*
and *viper.AddConfigPath("/Users/sedat/go/src/story")* in *config.go* and:

    ` go run main.go`

- Application listen port *8080*

- I sed *gin* for handling and routing. You can check from here:
[https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)

- For getting to .yml, **viper** was used. You can check from here: 
[https://github.com/spf13/viper](https://github.com/spf13/viper)

- I used *Postgresql* for a relational database. I used *gorm* for
connection and sql commands. You can check form here :
[https://github.com/go-gorm/gorm](https://github.com/go-gorm/gorm)

- I created 3 tables named Events, Tokens and Stories. 

- Events table stores : app_id, story_id,type and count.

- Tokens table stores : app_id and token.

- Stories table stores : app_id,story_id and metadata.

-  For logging to errors, **logrus** was used because it provides 
sentry implementation easily. You can check it from here : 
[https://github.com/sirupsen/logrus](https://github.com/sirupsen/logrus)

-  APIs return **StatusInternalServer** error, If there is a server 
related problem.
    
-  APIs return **StatusNotFound** error, If record is not found.
    
-  Added comments for exportable functions.
    
-  I used **go.mod** for dependency management.

###First Assignment

 - In this assignment you can send GET request with *app_token* as parameter 
and get stories information about application. 

-   **API** :
    
	**Method** : GET

	**Path** : $ curl "[http://localhost:8080/stories/token_1](http://localhost:8080/stories/token_1)"
	
	**Response** :
```json
{
    "app_id": 1,
    "ts": 1612638478,
    "metadata": [
        {
            "id": 1,
            "metadata": {
                "img": "image1.png"
            }
        },
        {
            "id": 2,
            "metadata": {
                "img": "image2.png"
            }
        },
        {
            "id": 3,
            "metadata": {
                "img": "image3.png"
            }
        }
    ]
}
```

###Second Assignment

- Firstly, I used gin. Because it provides some convenience. But then
I used native net/http package for better performance. 

- I used this command for Apache ab tool :
*ab -k -n 100000 -c 100 http://localhost:8080/stories/token_1*

- - k : keep alive, don't close connection
- - n 100000 : I send 100000 requests totally
- - c 100 : concurrency 
- - with this command, 100 user sends 100000 request

- Gin Results :

<p align="center"><img src="https://raw.githubusercontent.com/sedataktas/story/main/screenshots/ab-test-with-gin.png?raw=true"/></p>

- Again, at the beginning, I reach to db 2 times. First to check app token exists. Second
is to take metadatas. I also check is app token exist in second query but with first query
I can return to user, *app_token not found* information. For improving performance, I
waive this so now I reach to db 1 time.

- I also created index for token field in tokens table because I always query 
to token.

- net/http, 1 reach to db and with app_token index results: 

<p align="center"><img src="https://raw.githubusercontent.com/sedataktas/story/main/screenshots/ab-test-with-net/http.png?raw=true"/></p>


- Additionally, if there are not many applications, tokens can be store memory.

###Third Assignment

- In this assignment api take *POST* request if there is a record matching the values
in request body, record's count increments 1. If there is no record  matching the values
in request body, creates new record and set count value 1.

- If there is no record matching to story_id, app_id and event type, returns error.

- For type field in events table, *valdiator* tag can be use(impression/close ...)

-   **API** :
    
	**Method** : POST

	**Path** : $ curl "[http://localhost:8080/event/token_1](http://localhost:8080/event/token_1)"
	
	**Request Body** :
```json
{
    "story_id": 1,
    "event_type": "impression"
}
```


