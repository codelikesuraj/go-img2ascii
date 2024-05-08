# go-img2ascii - An image-to-ascii converter
This is a cli tool that converts any jpeg, jpg or png image to an ascii representation.

## Usage
```go run main.go [path-to-image] [character-width]```

## Todo
- ✅ load and decode image file
- ✅ resize the image
- ✅ convert image to grayscale
- ✅ map grayscaled image to ascii characters
- ✅ save ascii'd image to file
- ✅ refactor with any cli library (https://github.com/urfave/cli)