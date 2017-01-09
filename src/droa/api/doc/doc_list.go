package doc

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/widuu/gomysql"
)

/*type Document struct {
	id            string
	name          string
	url           string
	uploader      string
	uploadDate    string
	reApprover    string
	reApproveDate string
	approve       string
	approveDate   string
	status        string
}
*/
func DocList(q string) (str string, errs error) {
	c, err := gomysql.SetConfig("src/config/dbConfig.ini")
	if err != nil {
		err = errors.New("数据库连接错误！")
	}
	_sql := setSql(q)
	data := c.Query(_sql)

	mapdata := data.(map[int]map[string]string)
	droa := make([]interface{}, 0, len(mapdata))
	for _, v := range mapdata {
		droa = append(droa, v)
	}
	resultJson := make(map[string]interface{})
	resultJson["doc"] = droa
	resultJson["count"] = len(mapdata)

	s, errs := json.Marshal(resultJson)

	return string(s), errs
}

func setSql(con string) (_sql string) {
	s := `SELECT 
		    h.docId,
		    d.name,
		    h.operator,
		    h.operateTime,
		    CASE d.status
		        WHEN 0 THEN '提交'
		        WHEN 1 THEN '复核通过'
		        WHEN 2 THEN '审核通过'
		        WHEN 3 THEN '未通过复核'
		        WHEN 4 THEN '未通过审核'
		        WHEN 5 THEN '已发出'
		        ELSE '未知'
		    END AS status,
		    d.url
		FROM
		    ((SELECT 
		        *
		    FROM
		        (SELECT 
		        docId, operator, operateTime
		    FROM
		        history
		    WHERE
		        status = 0
		    ORDER BY operateTime ASC) f
		    GROUP BY docId) h
		    LEFT JOIN document d ON (d.id = h.docId))
		WHERE 
		d.name`
	nn := " is not null"
	like_s := " like '%"
	like_e := "%'"
	conti := " ORDER BY operateTime DESC LIMIT 0 , 30;"
	var rsql string
	if len(con) > 0 {
		rsql = fmt.Sprintf("%v%v%v%v%v", s, like_s, con, like_e, conti)
	} else {
		rsql = fmt.Sprintf("%v%v%v", s, nn, conti)
	}

	return rsql
}

/*
SELECT
			h.docId,
		    d.name,
		    h.operator,
		    h.operateTime,
		    CASE d.status
		        WHEN 0 THEN '提交'
		        WHEN 1 THEN '复核通过'
		        WHEN 2 THEN '审核通过'
		        WHEN 3 THEN '未通过复核'
		        WHEN 4 THEN '未通过审核'
		        WHEN 5 THEN '已发出'
		        ELSE '未知'
		    END AS status,
		    d.url
		FROM
		    ((SELECT
		        *
		    FROM
		        (SELECT
		        docId, operator, operateTime
		    FROM
		        history
		    WHERE
		        status = 0
		    ORDER BY operateTime ASC) f
		    GROUP BY docId) h
		    LEFT JOIN document d ON (d.id = h.docId))
		WHERE
		    d.name IS NOT NULL
		ORDER BY operateTime DESC
		LIMIT 0 , 20;
*/
