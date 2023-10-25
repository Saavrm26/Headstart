package Controllers

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
	"hs/headstart/Database"
	"hs/headstart/Utils"
	"mime/multipart"
	"net/http"
)

func GetAllCatalouges(ctx *gin.Context) {
	data := make([]map[string]any, 0)
	iter := Database.Fs.Collection("Prebuilt").Documents(ctx)
	var err error
	for {
		doc, itrErr := iter.Next()
		if itrErr == iterator.Done {
			break
		}
		if itrErr != nil {
			err = itrErr
		}
		data = append(data, doc.Data())
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}

type CatalougeObject struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"imageURL"`
	ZipLocation string `json:"zipLocation"`
}

func handleUpload(fileHeader *multipart.FileHeader, uploadpath string) error {
	fileBuffer, err := Utils.MultipartFileBuffer(fileHeader)
	err = Database.UploadToCloudStorage(uploadpath+fileHeader.Filename, fileBuffer)
	if err != nil {
		return err
	}
	return nil
}

// TODO: Implement cleanUp
func cleanUp() {

}

func PostACatalouge(ctx *gin.Context) {
	// TODO: do sanitization
	name := ctx.PostForm("name")
	description := ctx.PostForm("description")
	imageHeader, _ := ctx.FormFile("image")
	zipHeader, _ := ctx.FormFile("zip")

	// TODO: add go routine for each of them
	err := handleUpload(imageHeader, "image/")
	err = handleUpload(zipHeader, "Prebuilt/")
	imagePublicUrl := Database.PublicURL("images/" + imageHeader.Filename)

	firestoreObject := CatalougeObject{name, description, imagePublicUrl, "Prebuilt/" + zipHeader.Filename}
	_, _, err = Database.Fs.Collection("Prebuilt").Add(ctx, firestoreObject)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		cleanUp()
		return
	}
	ctx.JSON(http.StatusCreated, firestoreObject)
}
