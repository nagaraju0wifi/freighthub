# freighthub

## Manual Executiong steps 

1. Clone the Git repo to your local golang src folder ( example : /go/src/ )
2. Install the dependency for routing. "go get -u github.com/gorilla/mux"
3. Start the program ( go run main.go )
4. Verify the output by hitting the endpoint ( http://localhost:80/convert/FF0000 )

## Executing through Docker image

1. Pull my customized docker image ( docker pull nagaraju0wifi/hex_to_rgb )
2. Run the container from above image ( docker run -dti -p 80:80 nagaraju0wifi/hex_to_rgb )
3. Login into the container ( docker exec -ti <container_id> /bin/bash )
4. Start the application ( go run /go/src/main.go & )
5. Verify the output by hitting the endpoint ( http://localhost:80/convert/FF0000 )
