package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/rumyantseva/lowest-common-ancestor/pkg/lca"
	"github.com/rumyantseva/lowest-common-ancestor/pkg/middleware"
)

// Find handles request of finding lowest common manager between two employees.
func ClosestCommonManager(calc lca.Finder) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		emp := r.URL.Query().Get("employees")
		if len(emp) == 0 {
			msg := "The `employees` parameter is required and must contain 1 or 2 keys delimited by comma."
			middleware.WriteResponseError(w, msg, http.StatusBadRequest)
			return
		}

		employees := strings.Split(emp, ",")

		employeesCnt := len(employees)
		if employeesCnt == 0 || employeesCnt > 2 {
			msg := "The `employees` parameter must contain 1 or 2 keys delimited by comma."
			middleware.WriteResponseError(w, msg, http.StatusBadRequest)
			return
		} else if employeesCnt == 1 {
			employees = append(employees, employees[0])
		}

		var manager *lca.Key
		manager = calc.Find(lca.Key(employees[0]), lca.Key(employees[1]))

		if manager == nil {
			var msg string
			if employeesCnt == 2 {
				msg = fmt.Sprintf(
					"Couldn't find the lowest manager between `%s` and `%s`. Please, check what both employees exist.",
					employees[0], employees[1],
				)
			} else {
				msg = fmt.Sprintf(
					"Couldn't find the manager of `%s`. Please, check what the employee exists.",
					employees[0],
				)
			}

			middleware.WriteResponseError(w, msg, http.StatusBadRequest)
			return
		}

		data := struct {
			Manager lca.Key `json:"manager"`
		}{
			Manager: *manager,
		}
		middleware.WriteResponseSuccess(w, data, http.StatusOK)
	}
}
