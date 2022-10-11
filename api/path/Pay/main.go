package pay

import (
	db "api/db"
	d "api/jwt/service"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type data struct {
	Status struct {
		Message string
		Code    string
	}
	Data struct {
		Voucher struct {
			Amount_baht string
		}
	}
}

func Getgv(c *gin.Context, s db.Db_mongo, mobile_number, reql string) {
	var b []byte
	tokenHeader := c.Request.Header.Get("jwt")
	ssss := make(chan primitive.D)
	key, err := d.DecodeToken(tokenHeader)
	if err != nil {
		c.AbortWithStatus(505)
		return
	}
	go s.Db_FindtOne("email", key.Claims.(jwt.MapClaims)["iss"].(string), ssss)
	file := c.PostForm("link")
	campaign_code := strings.Replace(file, "https://gift.truemoney.com/campaign/?v=", "", -1)
	redeem_url := "http://" + reql + "/" + mobile_number + "/" + campaign_code

	a, err := http.Get(redeem_url)
	if err != nil {
		c.JSON(500, err)
	}
	defer a.Body.Close()
	asfd := data{}

	b, err = ioutil.ReadAll(a.Body)
	json.Unmarshal(b, &asfd)

	if asfd.Status.Code != "SUCCESS" {
		c.JSON(500, "useed ")
		return
	}

	asdad, err := strconv.ParseFloat(asfd.Data.Voucher.Amount_baht, 64)
	if err != nil {
		c.JSON(500, "err")
		return
	}
	go s.Db_FixOneStuck(bson.M{"email": key.Claims.(jwt.MapClaims)["iss"].(string)},
		bson.M{"$inc": bson.M{"cradit": asdad}})

	c.JSON(http.StatusOK, asfd)

}
