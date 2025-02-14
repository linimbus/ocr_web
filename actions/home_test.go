package actions

import "net/http"

func (as *ActionSuite) Test_HomeHandler() {
	res := as.HTML("/").Get()

	as.Equal(http.StatusOK, res.Code)
	as.Contains(res.Body.String(), "Welcome to Buffalo")
}

func (as *ActionSuite) Test_Home_Index() {
	as.Fail("Not Implemented!")
}

