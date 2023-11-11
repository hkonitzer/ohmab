package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/hkonitzer/ohmab/ent"
	_ "github.com/hkonitzer/ohmab/ent/runtime"
	"github.com/hkonitzer/ohmab/ent/timetable"
	"github.com/hkonitzer/ohmab/internal/pkg/common/log"
	"github.com/hkonitzer/ohmab/internal/pkg/db"
	"github.com/hkonitzer/ohmab/internal/pkg/ghost"
	"github.com/hkonitzer/ohmab/internal/pkg/routes"
	"github.com/hkonitzer/ohmab/internal/pkg/utils"
	"github.com/hkonitzer/ohmab/templates"
	"strings"
	"time"
)

var logger = log.GetLoggerInstance()

func main() {

	// get the timetable data
	ctx := context.TODO()
	client, clientError := db.CreateClient(ctx)
	if clientError != nil {
		logger.Fatal().Msgf("Error creating client: %v", clientError)
	}
	defer client.Close()

	// determine ghost locale
	locale := configurations.Ghost.Locale // try config first
	if locale == "" {
		// request locale from ghost site meta data
		site, err := ghost.RequestSite()
		if err != nil {
			logger.Fatal().Msgf("Error requesting site: %v", err)
		}
		locale = site.Site.Locale
	}
	logger.Debug().Msgf("Using locale: %s", locale)

	ttdata := routes.DataTemplate{
		Locale: locale,
	}

	_, err := routes.GetStandardTimetableData(&ttdata, "EMERGENCYSERVICE", nil, client, ctx)
	if err != nil {
		logger.Fatal().Msgf("Error getting timetable data: %v", err)
	}
	var buff1 bytes.Buffer
	if err := templates.TableTimeTableHTMLTemplate.Execute(&buff1, ttdata); err != nil {
		logger.Fatal().Msgf("Error executing template: %v", err)
	}

	var postID, pageID, postTitle string
	for _, i := range configurations.Ghost.Content {
		for k, v := range i {
			if k == strings.ToLower(timetable.TimetableTypeEMERGENCYSERVICE.String()) {
				postID = v.PostID
				pageID = v.PageID
				postTitle = v.PostTitle
				break
			}
		}
	}
	if pageID == "" {
		logger.Fatal().Msgf("Error getting pageID, see config.yml")
	}

	// request pages first to get the updated_at date (needed for UpdatePage)
	sres := ghost.SingleRessource{
		ID: pageID,
	}
	pages := ghost.PageEntity{
		Pages: []ghost.SingleRessource{sres},
	}
	err = pages.Request()
	page := &pages.Pages[0]
	if err != nil {
		logger.Fatal().Msgf("Error requesting page: %v", err)
	}
	// replace mobiledoc
	page.MobileDoc = ghost.ReplaceMobileDoc(page.MobileDoc, buff1.String())
	page.PublishedAt = time.Now().Format(time.RFC3339)
	page.Status = "published"
	// update
	err = pages.UpdatePage()
	if err != nil {
		logger.Fatal().Msgf("Error updating pages with id '%s': %v", page.ID, err)
	}
	logger.Debug().Msgf("Updated page with id '%s'", page.ID)
	// update post
	if postID != "" {
		sres = ghost.SingleRessource{
			ID: postID,
		}
		posts := ghost.PostEntity{
			Posts: []ghost.SingleRessource{sres},
		}
		err = posts.Request()
		if err != nil {
			logger.Fatal().Msgf("Error requesting post: %v", err)
		}
		firstEntry := ttdata.TimetableTypeData[0].Data[0]
		ttdata.TimetableTypeData[0].Data = make([]*ent.Timetable, 0)
		ttdata.TimetableTypeData[0].Data = append(ttdata.TimetableTypeData[0].Data, firstEntry)
		var buff2 bytes.Buffer
		if err := templates.TableTimeTableHTMLTemplate.Execute(&buff2, ttdata); err != nil {
			logger.Fatal().Msgf("Error executing template: %v", err)
		}
		posts.Posts[0].MobileDoc = ghost.ReplaceMobileDoc(posts.Posts[0].MobileDoc, buff2.String())
		posts.Posts[0].PublishedAt = time.Now().Format(time.RFC3339)
		posts.Posts[0].Status = "published"
		posts.Posts[0].Title = fmt.Sprintf(postTitle, firstEntry.I18nDate(timetable.FieldDatetimeTo))

		posts.Posts[0].Excerpt = fmt.Sprintf("%s, %s, %s %s - %s",
			firstEntry.Edges.Address.Edges.Business.Name1,
			firstEntry.Edges.Address.Street,
			firstEntry.Edges.Address.Zip,
			firstEntry.Edges.Address.City,
			utils.ConvertTel(firstEntry.Edges.Address.Telephone),
		)
		err = posts.UpdatePost()
		if err != nil {
			logger.Fatal().Msgf("Error updating post with id '%s': %v", posts.Posts[0].ID, err)
		}
		logger.Debug().Msgf("Updated post with id '%s'", posts.Posts[0].ID)
	} else {
		logger.Debug().Msgf("No postID given, skipping post update")
	}

}
