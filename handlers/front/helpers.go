package front

import (
	"github.com/kendellfab/publish/domain"
	"net/http"
	"strconv"
)

func GetPagination(r *http.Request, total, pageCount int) domain.Pagination {
	paginator := domain.Pagination{}
	paginator.Count = pageCount

	r.ParseForm()
	if pStr, ok := r.Form["page"]; ok {
		if pg, pgErr := strconv.Atoi(pStr[0]); pgErr == nil {
			paginator.Offset = pg - 1
		}
	}

	if paginator.Offset >= 1 {
		paginator.HasNewer = true
		paginator.NewerIndex = paginator.Offset
	}

	if (paginator.Offset+1)*paginator.Count < total {
		paginator.HasOlder = true
		paginator.OlderIndex = paginator.Offset + 2
		if paginator.OlderIndex == 1 {
			paginator.OlderIndex += 1
		}
	}

	paginator.Total = total

	return paginator
}
