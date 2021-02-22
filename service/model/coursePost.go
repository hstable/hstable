package model

type PostCourse struct {
	PPylx        int    `json:"p_pylx"`
	PSfgldjr     int    `json:"p_sfgldjr"`
	PSfredis     int    `json:"p_sfredis"`
	PSfsyxkgwc   int    `json:"p_sfsyxkgwc"`
	PXktjz       string `json:"p_xktjz"`
	PChaxunxh    string `json:"p_chaxunxh"`
	PGjz         string `json:"p_gjz"`
	PXn          string `json:"p_xn"`
	PXq          int    `json:"p_xq"`
	PXnxq        string `json:"p_xnxq"`
	PDqxn        string `json:"p_dqxn"`
	PDqxq        int    `json:"p_dqxq"`
	PDqxnxq      string `json:"p_dqxnxq"`
	PXkfsdm      string `json:"p_xkfsdm"`
	PXiaoqu      string `json:"p_xiaoqu"`
	PKkyx        string `json:"p_kkyx"`
	PXkxs        string `json:"p_xkxs"`
	PID          string `json:"p_id"`
	PSfhlctkc    int    `json:"p_sfhlctkc"`
	PSfhllrlkc   int    `json:"p_sfhllrlkc"`
	PKxsjXqj     string `json:"p_kxsj_xqj"`
	PKxsjKsjc    string `json:"p_kxsj_ksjc"`
	PKxsjJsjc    string `json:"p_kxsj_jsjc"`
	PKcdmJs      string `json:"p_kcdm_js"`
	PKcdmCxrw    string `json:"p_kcdm_cxrw"`
	PKcGjz       string `json:"p_kc_gjz"`
	PXzcxtjzNj   string `json:"p_xzcxtjz_nj"`
	PXzcxtjzYx   string `json:"p_xzcxtjz_yx"`
	PXzcxtjzZy   string `json:"p_xzcxtjz_zy"`
	PXzcxtjzZyfx string `json:"p_xzcxtjz_zyfx"`
	PXzcxtjzBj   string `json:"p_xzcxtjz_bj"`
	PSfxsgwckb   int    `json:"p_sfxsgwckb"`
	PSkyy        string `json:"p_skyy"`
}

func Get_Post_Course() PostCourse {
	post_data := PostCourse {
		PPylx: 2,
		PSfgldjr: 0,
		PSfredis: 0,
		PSfsyxkgwc: 0,
		PXktjz: "",
		PChaxunxh: "",
		PGjz: "",
		PXn: "2020-2021",
		PXq: 2,
		PXnxq: "2020-20212",
		PDqxn: "2020-2021",
		PDqxq: 2,
		PDqxnxq: "2020-20212",
		PXkfsdm: "yixuan",
		PXiaoqu: "",
		PKkyx: "",
		PXkxs: "",
		PID: "",
		PSfhlctkc: 0,
		PSfhllrlkc: 0,
		PKxsjXqj: "",
		PKxsjKsjc: "",
		PKxsjJsjc: "",
		PKcdmJs: "",
		PKcdmCxrw: "",
		PKcGjz: "",
		PXzcxtjzNj: "",
		PXzcxtjzYx: "",
		PXzcxtjzZy: "",
		PXzcxtjzZyfx: "",
		PXzcxtjzBj: "",
		PSfxsgwckb: 1,
		PSkyy: "",
	}
	return post_data
}
