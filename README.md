# go-img2ascii - An image-to-ascii converter
This is a cli tool that converts any jpeg, jpg or png image to an ascii representation.

## Usage
```go run main.go [path-to-image] [character-width]```

## Sample
|Original|Ascii'd|
|-|-|
| ![image](https://github.com/codelikesuraj/go-img2ascii/assets/70463535/32a519a6-fd2d-4c0c-ba29-6f16430a7713) | ![image](https://github.com/codelikesuraj/go-img2ascii/assets/70463535/e9da71eb-97c6-4942-9419-4f33a67887db) |
| ![image](https://github.com/codelikesuraj/go-img2ascii/assets/70463535/1e4f629b-f69b-48b4-a1a4-7dd1a857f377) | ![image](https://github.com/codelikesuraj/go-img2ascii/assets/70463535/55f79285-91b8-47ed-af4f-79247abf98ec) |

## Todo
- ✅ load and decode image file
- ✅ resize the image
- ✅ convert image to grayscale
- ✅ map grayscaled image to ascii characters
- ✅ save ascii'd image to file
- ✅ refactor with any cli library (https://github.com/urfave/cli)
