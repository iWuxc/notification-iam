package notification

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	"github.com/marmotedu/component-base/pkg/json"
	"github.com/spf13/viper"
)

func (n *NotificationController) Post(c *gin.Context) {
	buf := make([]byte, 1024)
	read, err := c.Request.Body.Read(buf)
	if err != nil {
		core.WriteResponse(c, err, nil)
	}
	for i := 0; i < read; i += 1 {
		if buf[i] == ' ' || buf[i] == '\x00' {
			buf = append(buf[:i], buf[(i+1):]...)
			read -= 1
		}
	}
	var p_json interface{}
	_ = json.Unmarshal(buf[:read], &p_json)

	//P_type := c.Param("type")
	//P_tpl := c.Param("tpl")
	P_ddurl := c.Param("ddurl")

	if P_ddurl == "" {
		P_ddurl = viper.GetString("webhook.dd-url")
	}

	//if P_tpl != "" && P_type != "" {
	//	tpltext, err := models.GetTplOne(P_tpl)
	//	if err != nil {
	//		logs.Error(logsign, err)
	//	}
	//	buf := new(bytes.Buffer)
	//	//tpl, err := template.New("").Funcs(template.FuncMap{"GetCSTtime": GetCSTtime}).Parse(tpltext.Tpl)
	//	tpl, err := template.New("").Funcs(funcMap).Parse(tpltext.Tpl)
	//	if err != nil {
	//		logs.Error(logsign, err.Error())
	//		message = err.Error()
	//	} else {
	//		tpl.Execute(buf, p_json)
	//		message = SendMessagePrometheusAlert(c, buf.String(), P_type, P_ddurl, P_wxurl, P_fsurl, P_webhookurl, P_phone, P_email, P_touser, P_toparty, P_totag, P_groupid, P_atsomeone, P_rr, logsign)
	//	}
	//} else {
	//	message = "接口参数缺失！"
	//	logs.Error(logsign, message)
	//}

	core.WriteResponse(c, nil, P_ddurl)
}
