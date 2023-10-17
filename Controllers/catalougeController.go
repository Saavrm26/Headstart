package Controllers

import (
	"github.com/gin-gonic/gin"
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
