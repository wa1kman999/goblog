package model

// Profile 个人信息
type Profile struct {
	ID        int    `cond:"primaryKey" json:"id"`
	Name      string `cond:"type:varchar(20)" json:"name"`
	Desc      string `cond:"type:varchar(200)" json:"desc"`
	QQ        string `cond:"type:varchar(200)" json:"qq"`
	Wechat    string `cond:"type:varchar(100)" json:"wechat"`
	Weibo     string `cond:"type:varchar(200)" json:"weibo"`
	Bili      string `cond:"type:varchar(200)" json:"bili"`
	Email     string `cond:"type:varchar(200)" json:"email"`
	Img       string `cond:"type:varchar(200)" json:"img"`
	Avatar    string `cond:"type:varchar(200)" json:"avatar"`
	IcpRecord string `cond:"type:varchar(200)" json:"icp_record"`
}
