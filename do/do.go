package do

import (
	"context"
	"time"
	"strings"
	"fmt"
	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

var client *godo.Client

type tokenSource struct {
	AccessToken string
}

func (t *tokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token {
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func initClient(apiKey string) {
	tokenSource := &tokenSource {
		AccessToken: apiKey,
	}

	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client = godo.NewClient(oauthClient)
}

func getRegion(region string) string {

	selectedRegion := region

	return selectedRegion
}

// This function is the main input for this part of the program
func GetServer(region string, apiKey string) {
	initClient(apiKey)

	dropletName := strings.Join([]string{"cpm", time.Now().String()}, "-")
	createRequest := &godo.DropletCreateRequest{
		Name:   dropletName,
		Region: getRegion(region),
		Size:   "s-1vcpu-1gb",
		Image: godo.DropletCreateImage{
			Slug: "ubuntu-14-04-x64",
		},
	}

	ctx := context.TODO()

	newDroplet, _, err := client.Droplets.Create(ctx, createRequest)

	if err != nil {
		fmt.Printf("Something bad happened: %s\n\n", err)
		return err
	}
}
