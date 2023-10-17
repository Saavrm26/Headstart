package Controllers

import (
	"bytes"
	"context"
	"fmt"
	"hs/headstart/Database"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllCatalouges(ctx *gin.Context) {
	// TODO: fetch info from firestore
	data := make([]map[string]any, 0)
	data = append(data, gin.H{
		"downloadURL": "abc.com",
		"imageURL":    "https://nextjstemplates.com/_next/image?url=https%3A%2F%2Fapi.nextjstemplates.com%2Fimage%2FStartup-Pro---Next.js-Starter-Template-for-SaaS-Startups-282e26f7-f543-4ae4-a777-ac306c08cce8.png&w=1920&q=100",
		"name":        "Lorem ipsum dolor sit amet",
		"description": "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim labore culpa sint ad nisi Lorem pariatur mollit ex esse exercitation amet. Nisi anim cupidatat excepteur officia. Reprehenderit nostrud nostrud ipsum Lorem est aliquip amet voluptate voluptate dolor minim nulla est proident. Nostrud officia pariatur ut officia. Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem duis laboris cupidatat officia voluptate. Culpa proident adipisicing id nulla nisi laboris ex in Lorem sunt duis officia eiusmod. Aliqua reprehenderit commodo ex non excepteur duis sunt velit enim. Voluptate laboris sint cupidatat ullamco ut ea consectetur et est culpa et culpa duis.",
	})
	data = append(data, gin.H{
		"downloadURL": "abc.com",
		"imageURL":    "https://nextjstemplates.com/_next/image?url=https%3A%2F%2Fapi.nextjstemplates.com%2Fimage%2FTaxonomy---Next.js-13-Boilerplate-and-Template-f40f4bbe-ef9b-4335-b019-69d014801670.png&w=1920&q=10https://firebasestorage.googleapis.com/v0/b/clickcraft-ac100.appspot.com/o/images%2FStartup---Free-Next.webp?alt=media&token=6776a4d4-80a9-4519-92ca-d6e27e77610b0",
		"name":        "Lorem ipsum dolor sit amet",
		"description": "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim labore culpa sint ad nisi Lorem pariatur mollit ex esse exercitation amet. Nisi anim cupidatat excepteur officia. Reprehenderit nostrud nostrud ipsum Lorem est aliquip amet voluptate voluptate dolor minim nulla est proident. Nostrud officia pariatur ut officia. Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem duis laboris cupidatat officia voluptate. Culpa proident adipisicing id nulla nisi laboris ex in Lorem sunt duis officia eiusmod. Aliqua reprehenderit commodo ex non excepteur duis sunt velit enim. Voluptate laboris sint cupidatat ullamco ut ea consectetur et est culpa et culpa duis.",
	})
	ctx.JSON(http.StatusOK, data)
}

// type CatalougeForm struct {
// 	name        string                `form:"name" binding:"required"`
// 	description string                `form:"description" binding:"required"`
// 	image       *multipart.FileHeader `form:"image" binding:"required"`
// 	zip         *multipart.FileHeader `form:"zip" binding:"required"`
// }

func PostACatalouge(ctx *gin.Context) {
	// TODO: do sanitization

	name := ctx.PostForm("name")
	description := ctx.PostForm("description")
	imageHeader, _ := ctx.FormFile("image")
	zipHeader, _ := ctx.FormFile("zip")
	imageReader, err := imageHeader.Open()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	imageBytes := make([]byte, imageHeader.Size)
	_, err = imageReader.Read(imageBytes)

	buf := bytes.NewBuffer(imageBytes)
	c, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	wc := Database.Bucket.Object("Prebuilt/" + imageHeader.Filename).NewWriter(c)
	if _, err = io.Copy(wc, buf); err != nil {
		_ = fmt.Errorf("io.Copy: %w", err)
		return
	}
	// Data can continue to be added to the file until the writer is closed.
	if err := wc.Close(); err != nil {
		_ = fmt.Errorf("Writer.Close: %w", err)
		return
	}

	fmt.Printf("%s uploaded to bucket.\n", imageHeader.Filename)
	log.Printf("%v", imageReader)

	log.Printf("%s %s %s %s", name, description, imageHeader.Filename, zipHeader.Filename)

}
