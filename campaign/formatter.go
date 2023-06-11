package campaign

import "strings"

type (
	CampaignFormatter struct {
		ID               int    `json:"id"`
		UserID           int    `json:"user_id"`
		Name             string `json:"name"`
		ShortDescription string `json:"short_description"`
		ImageUrl         string `json:"image_url"`
		Slug             string `json:"slug"`
		GoalAmount       int    `json:"goal_amount"`
		CurrentAmount    int    `json:"current_amount"`
	}

	CampaignDetailFormatter struct {
		ID               int                      `json:"id"`
		Name             string                   `json:"name"`
		ShortDescription string                   `json:"short_description"`
		Description      string                   `json:"description"`
		ImageUrl         string                   `json:"image_url"`
		GoalAmount       int                      `json:"goal_amount"`
		CurrentAmount    int                      `json:"current_amount"`
		UserID           int                      `json:"user_id"`
		Slug             string                   `json:"slug"`
		Perks            []string                 `json:"perks"`
		User             CampaignUserFormatter    `json:"user"`
		Images           []CampaignImageFormatter `json:"images"`
	}

	CampaignUserFormatter struct {
		Name     string `json:"name"`
		ImageUrl string `json:"image_url"`
	}

	CampaignImageFormatter struct {
		ImageUrl  string `json:"image_url"`
		IsPrimary bool   `json:"is_primary"`
	}
)

func FormatCampaign(campaign Campaign) CampaignFormatter {
	formatter := CampaignFormatter{}
	formatter.ID = campaign.ID
	formatter.UserID = campaign.UserID
	formatter.Name = campaign.Name
	formatter.ShortDescription = campaign.ShortDescription
	formatter.Slug = campaign.Slug
	formatter.GoalAmount = campaign.GoalAmount
	formatter.CurrentAmount = campaign.CurrentAmount
	formatter.ImageUrl = ""

	if len(campaign.CampaignImages) > 0 {
		formatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	return formatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}
	return campaignsFormatter
}

func FormatDetailCampaign(campaign Campaign) CampaignDetailFormatter {
	campaignDetailFormatter := CampaignDetailFormatter{}
	campaignDetailFormatter.ID = campaign.ID
	campaignDetailFormatter.Name = campaign.Name
	campaignDetailFormatter.Description = campaign.Description
	campaignDetailFormatter.ShortDescription = campaign.ShortDescription
	campaignDetailFormatter.Slug = campaign.Slug
	campaignDetailFormatter.UserID = campaign.UserID
	campaignDetailFormatter.GoalAmount = campaign.GoalAmount
	campaignDetailFormatter.CurrentAmount = campaign.CurrentAmount
	campaignDetailFormatter.ImageUrl = ""

	if len(campaign.CampaignImages) > 0 {
		campaignDetailFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	campaignDetailFormatter.Perks = perks

	user := campaign.User
	campaignUserFormater := CampaignUserFormatter{}
	campaignUserFormater.Name = user.Name
	campaignUserFormater.ImageUrl = user.AvatarFileName

	images := []CampaignImageFormatter{}
	for _, image := range campaign.CampaignImages {
		campaignImageFormatter := CampaignImageFormatter{}
		campaignImageFormatter.ImageUrl = image.FileName

		isPrimary := false
		if image.IsPrimary == 1 {
			isPrimary = true
		}
		campaignImageFormatter.IsPrimary = isPrimary
		images = append(images, campaignImageFormatter)
	}

	campaignDetailFormatter.Images = images
	campaignDetailFormatter.User = campaignUserFormater

	return campaignDetailFormatter
}
