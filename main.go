package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/log"
)

func main() {

	r := gin.Default()
	//解决跨域
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))
	r.Use(gin.Recovery())
	r.Use(static.Serve("/", static.LocalFile(".", true)))

	rootApi := r.Group("/")

	// Ping test
	rootApi.GET("ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	usersGroup := rootApi.Group("api")
	{
		usersGroup.POST("/ActivityCenter/ActivityList", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"Result": []map[string]interface{}{
					{
						"AccountId":        "44-9731211B2286C6PWMA2BA0BS570",
						"ActivityName":     "活动名称1",
						"ActivityType":     "体育",
						"ActivityStartDT":  "2018-06-30 12:01:31",
						"ActivityEndDT":    "2018-06-30 12:01:31",
						"ContactUser":      "业主1",
						"ContactPhone":     "15365895645",
						"Description":      "为鼓励业主锻炼身体，建立良好邻里关系，申请举办羽毛球友谊赛。",
						"UserQty":          30,
						"Status":           "待审批",
						"AgreeRule":        "1,2,3",
						"TotalAmount":      3000,
						"ApproverStep":     "初审",
						"ApproverUser1":    "初审人员1",
						"ApproverContent1": "初审内容",
						"ApproverResult1":  0,
						"ApproverDate1":    "2018-06-30 12:01:31",
						"ApproverUser2":    "复审人员1",
						"ApproverContent2": "复审内容",
						"ApproverResult2":  0,
						"ApproverDate2":    "2018-06-30 12:01:31",
						"ApproverUser3":    "终审人员1",
						"ApproverContent3": "终审内容",
						"ApproverResult3":  0,
						"ApproverDate3":    "2018-06-30 12:01:31",
						"CreatedDT":        "2018-06-30 12:01:31",
						"CreatedBy":        "业主1", "Id": 2,
					},
					{
						"AccountId":        "44-9731211B2286C6PWMA2BA0BS570",
						"ActivityName":     "活动名称1",
						"ActivityType":     "体育",
						"ActivityStartDT":  "2018-06-30 12:01:31",
						"ActivityEndDT":    "2018-06-30 12:01:31",
						"ContactUser":      "业主1",
						"ContactPhone":     "15365895645",
						"Description":      "为鼓励业主锻炼身体，建立良好邻里关系，申请举办羽毛球友谊赛。",
						"UserQty":          30,
						"Status":           "待审批",
						"AgreeRule":        "1,2,3",
						"TotalAmount":      3000,
						"ApproverStep":     "初审",
						"ApproverUser1":    "初审人员1",
						"ApproverContent1": "初审内容",
						"ApproverResult1":  0,
						"ApproverDate1":    "2018-06-30 12:01:31",
						"ApproverUser2":    "复审人员1",
						"ApproverContent2": "复审内容",
						"ApproverResult2":  0,
						"ApproverDate2":    "2018-06-30 12:01:31",
						"ApproverUser3":    "终审人员1",
						"ApproverContent3": "终审内容",
						"ApproverResult3":  0,
						"ApproverDate3":    "2018-06-30 12:01:31",
						"CreatedDT":        "2018-06-30 12:01:31",
						"CreatedBy":        "业主1",
						"Id":               2,
					},
				},
				"Pagination": map[string]interface{} {
					"Current": 1,
					"Total": 10,
					"PageSize":10,
				}
				"Message": "操作成功",
				"Code":    200,
			})
		})

		usersGroup.GET("/ActivityCenter/ActivityTypes", func(c *gin.Context) {
			c.JSON(http.StatusOK, []map[string]interface{}{
				{"Id": 1, "Name": "体育"},
			})
			return
		})

		usersGroup.GET("/ActivityCenter/ActivityGoods", func(c *gin.Context) {
			c.JSON(http.StatusOK, []map[string]interface{}{
				{
					"Id":       1,
					"Name":     "全部",
					"IconName": "images/u257.png",
					"ActivityGoods": []map[string]interface{}{
						{
							"GoodName": "物资名称1",
							"UnitName": "场",
							"GoodType": "活动场地"},
						{
							"GoodName": "物资名称2",
							"UnitName": "个",
							"GoodType": "活动场地",
						},
					},
				},
				{"Id": 2, "Name": "活动场地", "IconName": "images/u143.png",
					"ActivityGoods": []map[string]interface{}{
						{
							"GoodName": "物资名称1",
							"UnitName": "场",
							"GoodType": "活动场地"},
						{
							"GoodName": "物资名称2",
							"UnitName": "个",
							"GoodType": "活动场地",
						},
					},
				},
				{"Id": 3, "Name": "宣传物料", "IconName": "images/u153.png",
					"ActivityGoods": []map[string]interface{}{
						{
							"GoodName": "宣传物资名称1",
							"UnitName": "场",
							"GoodType": "宣传物料"},
						{
							"GoodName": "宣传物资名称2",
							"UnitName": "个",
							"GoodType": "宣传物料",
						},
					},
				},
				{"Id": 4, "Name": "设备道具", "IconName": "images/u155.png",
					"ActivityGoods": []map[string]interface{}{
						{
							"GoodName": "设备物资名称1",
							"UnitName": "场",
							"GoodType": "设备道具"},
						{
							"GoodName": "设备物资名称2",
							"UnitName": "个",
							"GoodType": "设备道具",
						},
					},
				},
				{"Id": 5, "Name": "活动物资", "IconName": "images/u141.png",
					"ActivityGoods": []map[string]interface{}{
						{
							"GoodName": "活动物资名称1",
							"UnitName": "场",
							"GoodType": "活动物资"},
						{
							"GoodName": "活动物资名称2",
							"UnitName": "个",
							"GoodType": "活动物资",
						},
					},
				},
				{"Id": 6, "Name": "其他", "IconName": "images/u161.png",
					"ActivityGoods": []map[string]interface{}{
						{
							"GoodName": "其他物资名称1",
							"UnitName": "场",
							"GoodType": "其他"},
						{
							"GoodName": "其他物资名称2",
							"UnitName": "个",
							"GoodType": "其他",
						},
					},
				},
			})
			return
		})

		usersGroup.POST("/ActivityCenter/ActivityApplyAdd", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"Message": "操作成功",
				"Code":    200,
			})
			return
		})

		//usersGroup.GET(":userId", srv.GetUser)
		//		usersGroup.PUT(":userId/jiafen", srv.PutUserJiaFen)
		//		usersGroup.PUT(":userId/sign", srv.PutUserSign)
		//		usersGroup.PUT(":userId/games", srv.PutUserGames)
		//		usersGroup.GET(":userId/gamesRecords", srv.GetUserGamesRecords)
		//		usersGroup.DELETE(":userId", srv.DeleteUser)
		//		usersGroup.PUT(":userId/exchange", srv.PutUserExchange) //积分兑换

	}

	//pprof.Register(r, nil)

	log.Fatal(r.Run("0.0.0.0:8007"))
}
