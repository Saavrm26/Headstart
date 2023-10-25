package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hs/headstart/Database"
	"hs/headstart/Utils"
	"log"
	"net/http"
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

	imgBuffer, err := Utils.MultipartFileBuffer(imageHeader)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	zipBuffer, err := Utils.MultipartFileBuffer(zipHeader)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err :=
		Database.UploadToCloudStorage("images/"+imageHeader.Filename, imgBuffer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err :=
		Database.UploadToCloudStorage("Prebuilt/"+zipHeader.Filename, zipBuffer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	imagePublicUrl := Database.PublicURL("images/" + imageHeader.Filename)


	fmt.Printf("%s uploaded to bucket.\n", imageHeader.Filename)
	log.Printf("%s %s %s %s", name, description, imageHeader.Filename, zipHeader.Filename)
	log.Printf("%s\n", imagePublicUrl)

}
