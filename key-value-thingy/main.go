package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

var Database = make(map[string]string)

func process_get_request( context *gin.Context ) {
    key := context.Params.ByName("key")
    value, ok := Database[key]
    if ok {
        context.JSON(200, gin.H{"key": key, "value": value})
    } else {
        context.JSON(200, gin.H{"key": key, "value": "nil"})
    }
}

func process_post( context *gin.Context ) {
    // Parse JSON
    var json struct {
        Key string `json:"key" binding:"required"`
        Value string `json:"value" binding:"required"`
    }

    if error := context.ShouldBindJSON(&json); error == nil {
        fmt.Println("STUFF:", json)

        Database[json.Key] = json.Value
        context.JSON( 200, gin.H{"status": "ok"} )
    } else {
        context.JSON( http.StatusBadRequest, gin.H{"error": error.Error } )
    }
}

func process_delete( context *gin.Context ) {
    key := context.Params.ByName("key")

    delete( Database, key )

    context.JSON( 200, gin.H{"status": "ok"} )
}

 
func setupRouter() *gin.Engine {
    
    router := gin.Default()

    // Get the correct value for the key provided
    router.GET( "/api/objects/:key", process_get_request )

    // Insert value stored at form "value" at the key form "key"
    router.POST( "/api/objects", process_post ) 

    // Remove the entry named by form parameter "key"
    router.DELETE( "/api/objects/:key", process_delete )

    return router
}

func main() {
    router := setupRouter()
    // Listen on 0.0.0.0:8080
    router.Run(":8080")
}
