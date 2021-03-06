package web

import (
	"OnlineJudge-RearEnd/api/database"
	"errors"

	"gorm.io/gorm"
)

/*
	@Title
	problem

	@Description
	题目相关

	@Func List

	Class name: problem

	| func name           | develop | unit test |  bug  |

	|---------------------------------------------------|

	| insert单            |   yes   |    no	    |  no   |

	| insert多            |   no    |    no	    |  no   |

	| Delete单            |   yes   |    no	    |  no   |

	| Delete多            |   no    |    no	    |  no   |

	| Update              |   yes   |    no	    |  no   |

	| Detail              |   yes   |    no	    |  no   |

	| List                |   yes   |    no	    |  no   |
*/

/*
	bug list
	没有做权限管理
*/

func (problem Problem) Insert() HTTPStatus {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "problem.insert",
			Method:      "post",
		}
	}
	err = mdb.Create(&problem).Error
	if err != nil {
		return HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql insert fail",
			RequestPath: "problem.insert",
			Method:      "post",
		}
	}

	return HTTPStatus{
		Message:     "题目添加成功",
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "problem.insert",
		Method:      "post",
	}
}

/*
	@input
	problem.ID

	bug list
	可能会触发批量delete
*/
func (problem Problem) Delete() HTTPStatus {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "problem.delete",
			Method:      "delete",
		}
	}
	err = mdb.Delete(&problem).Error
	if err != nil {
		return HTTPStatus{
			Message:     "删除出错！",
			IsError:     true,
			ErrorCode:   403,
			SubMessage:  "problem delete error, error code is error",
			RequestPath: "problem.delete",
			Method:      "delete",
		}
	}

	return HTTPStatus{
		Message:     "删除成功",
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "problem.delete",
		Method:      "delete",
	}
}

/*
	偶为修改前，奇数为修改后
	输入的为修改后的
*/
func (problem Problem) Update() HTTPStatus {
	if problem.ID <= 0 {
		return HTTPStatus{
			Message:     "输入的什么鬼东西",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "id error",
			RequestPath: "problem.QueryDetail",
			Method:      "get",
		}
	}
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "problem.delete",
			Method:      "delete",
		}
	}
	err = mdb.Save(&problem).Error
	if err != nil {
		return HTTPStatus{
			Message:     "更新失败",
			IsError:     true,
			ErrorCode:   1,
			SubMessage:  "update error, error code is error",
			RequestPath: "problem.update",
			Method:      "put",
		}
	}

	return HTTPStatus{
		Message:     "更新成功",
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "problem.update",
		Method:      "put",
	}
}

/*
	查询可以优化

	bug
	查过头会报错
*/
func (problem Problem) List(pageIndex int, pageSize int) ([]Problem, HTTPStatus) {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return []Problem{}, HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "problem.list",
			Method:      "get",
		}
	}

	//分页查询
	if pageIndex <= 0 || pageSize <= 0 {
		return []Problem{}, HTTPStatus{
			Message:     "非法输入",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "page index or page size input error, error code is error",
			RequestPath: "problem.list",
			Method:      "get",
		}
	}

	var problems []Problem
	err = mdb.Debug().Offset((pageIndex-1)*pageSize).Limit(pageSize).Select("id", "name", "is_hide_to_user").Find(&problems).Error
	if err != nil {
		return []Problem{}, HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "query error",
			RequestPath: "problem.list",
			Method:      "get",
		}
	}

	return problems, HTTPStatus{}
}

/*
	需要输入id
*/
func (problem Problem) Detail() (Problem, HTTPStatus) {
	if problem.ID <= 0 {
		return Problem{}, HTTPStatus{
			Message:     "输入的什么鬼东西",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "id error",
			RequestPath: "problem.detail",
			Method:      "get",
		}
	}

	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return Problem{}, HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "problem.detail",
			Method:      "get",
		}
	}

	if errors.Is(mdb.First(&problem).Error, gorm.ErrRecordNotFound) {
		return Problem{}, HTTPStatus{
			Message:     "没有这个题目",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "id error",
			RequestPath: "problem.detail",
			Method:      "get",
		}
	}

	return problem, HTTPStatus{
		Message:     "查询成功",
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "problem.detail",
		Method:      "get",
	}
}