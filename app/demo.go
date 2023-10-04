package app

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/guimassoqueto/go-web-rankings/website"
)

func RunRepositoryDemo(ctx context.Context, websiteRepository website.Repository) {
	fmt.Println("1. MIGRATE REPOSITORY")
	if err := websiteRepository.Migrate(ctx); err != nil {
		log.Fatal(err)
	}

	fmt.Println("2. CREATE RECORDS OF REPOSITORY")
	facebookWebsite := website.Website{
		Name: "FACEBOOK",
		URL: "https://facebook.com/",
		Rank: 2,
	}
	youtubeWebsite := website.Website{
		Name: "YOUTUBE",
		URL: "https://youtube.com/",
		Rank: 3,
	}
	googleWebsite := website.Website{
		Name: "GOOGLE",
		URL: "https://google.com/",
		Rank: 1,
	}

	// tenta adicionar facebook no banco
	createdFacebookWebSite, err := websiteRepository.Create(ctx, facebookWebsite)
	if err != nil {
		if errors.Is(err, website.ErrDuplicate) {
			fmt.Printf("record: %+v alread exists\n", facebookWebsite)
		}
		log.Fatal(err)
	}

	// tenta adicionar youtube no banco
	createdYoutubeWebSite, err := websiteRepository.Create(ctx, youtubeWebsite)
	if err != nil {
		if errors.Is(err, website.ErrDuplicate) {
			fmt.Printf("record: %+v alread exists\n", youtubeWebsite)
		}
		log.Fatal(err)
	}

	// tenta adicionar google no banco
	createdGoogleWebSite, err := websiteRepository.Create(ctx, googleWebsite)
	if err != nil {
		if errors.Is(err, website.ErrDuplicate) {
			fmt.Printf("record: %+v alread exists\n", youtubeWebsite)
		}
		log.Fatal(err)
	}

	// retorna os 4 websites inseridos em caso de sucesso
	fmt.Println("SUCCESS ON WEBSITE CREATION")
	fmt.Printf("%+v\n%+v\n%+v\n", createdFacebookWebSite, createdYoutubeWebSite, createdGoogleWebSite)

	// tenta buscar o facebook recém inserido no banco de dados
	fmt.Println("4. GET RECORD BY NAME (facebook)")
	gotFacebookWebsite, err := websiteRepository.GetByName(ctx, "FACEBOOK")
	if errors.Is(err, website.ErrNotExist) {
		fmt.Println("record: GOSAMPLES does not exist in the repository")
	} else if err != nil {
		log.Fatal(err)
	}
	fmt.Println("SUCCESS ON RETRIEVING FACEBOOK FROM DATABASE")
	fmt.Printf("%+v\n", gotFacebookWebsite)

	// tenta atualizar um website
	fmt.Println("4. UPDATE RECORD")
	createdFacebookWebSite.Name = "Instagram"
	if _, err := websiteRepository.Update(ctx, createdFacebookWebSite.ID, *createdFacebookWebSite); err != nil {
		if errors.Is(err, website.ErrDuplicate) {
			fmt.Printf("record: %+v already exists\n", createdFacebookWebSite)
		} else if errors.Is(err, website.ErrUpdateFailed) {
			fmt.Printf("update of record: %+v already exists\n", createdFacebookWebSite)
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println("SUCCESS ON UPDATING FACEBOOK -> INSTAGRAM ON DATABASE")

	// tenta pegar todos os websites do banco
	fmt.Println("5. GET ALL")
	all, err := websiteRepository.All(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, website := range all{
		fmt.Printf("%+v\n", website)
	}
	fmt.Println("SUCCESS ON GETTING ALL THE WEBSITES")
	
	// tenta excluir um website do banco
	fmt.Println("5. DELETE FACEBOOK (NOW INSTAGRAM) FROM DATABASE")
	if err := websiteRepository.Delete(ctx, createdFacebookWebSite.ID); err != nil {
		if errors.Is(err, website.ErrDeleteFailed) {
			fmt.Printf("delete of record: %d failed", createdFacebookWebSite.ID)
		} else {
			log.Fatal(err)
		}
	}

	// busca novamente todos os websites
	all, err = websiteRepository.All(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, website := range all {
		fmt.Printf("%+v\n", website)
	}
}