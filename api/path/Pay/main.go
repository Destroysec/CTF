package pay

import (
	db "api/db"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Getgv(c *gin.Context, s db.Db_mongo, mobile_number string) {
	file := c.PostForm("link")
	campaign_code := strings.Replace(file, "https://gift.truemoney.com/compaign?v=", "", -1)
	redeem_url := "https://gift.truemoney.com/campaign/vouchers/" + campaign_code + "/redeem"
	payload, _ := json.Marshal(map[string]string{"mobile": mobile_number})
	response_campaign, err := http.Post(redeem_url, "application/json", bytes.NewBuffer(payload))
	if err != nil {

	}
	defer response_campaign.Body.Close()
	body, _ := ioutil.ReadAll(response_campaign.Body)
	c.JSON(http.StatusOK, gin.H{
		"message": body,
	})

}
