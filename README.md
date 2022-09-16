## following tutorial from logrocket

https://blog.logrocket.com/documenting-go-web-apis-with-swag/

## Demo
- Create and Read
<video src="https://user-images.githubusercontent.com/16322250/190625180-27e346a1-a7b5-4e40-8a2c-d714d0ebac1f.mov" controls="controls" style="max-width: 720px;">
  </video>

- Update and Delete
<video src="https://user-images.githubusercontent.com/16322250/190625314-d00c4f59-6fee-44cb-9c6b-34040b17cc78.mov" controls="controls" style="max-width: 720px;">
  </video>

## issues faced

- Go wasn't setup properly, so swag init didnt work

  - issue
    - GOBIN is empty
      - to check, `go env` or `go env GOBIN` - it will be empty
  - fixes i did
    - export GOBIN=$(go env GOPATH)/bin
    - this will set the current project(?, i think) to use the go bin folder
      - in my case its `/Users/zhenhaoc/go/bin`
    - afterwards `go install github.com/swaggo/swag/cmd/swag` <-- this will install swag into the go bin folder
  - source https://github.com/swaggo/swag/issues/197#issuecomment-1100847754 tyvm!~

- main.go doesn't recognise swaggerFiles

  - issue
    - `swaggerFiles "github.com/swaggo/files"` keeps getting deleted on save by gopls
  - fixes
    - save the main.go without the line first (lol)
    - `go get -u github.com/swaggo/files`
    - then add it back in lol
    - seems to resolve it?

- trying to re-order the routes doesnt work
  - issue
    - reordering the routes doesnt do anything
    - apparently the operations needs to be sorted/defined
    - need to look into how to do it
  - fixes
    - try another tutorial lol
