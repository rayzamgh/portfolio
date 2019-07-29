package service

import (
	"net/http"

	"gitlab.com/standard-go/project/internal/app/project"
)

func printErrorMessage(err error) *project.SingleResponse {
	var ret *project.SingleResponse
	
	switch err.Error() {
	case "400":
		ret = &project.SingleResponse {
			Meta: project.SimpleMeta{
				StatusCode: http.StatusBadRequest,
				Message: project.Message{
					"errors": []string{"Bad Request"},
				},	
			},
		}
		break
	case "400-Token-Parse":
		ret = &project.SingleResponse {
			Meta: project.SimpleMeta{
				StatusCode: http.StatusBadRequest,
				Message: project.Message{
					"errors": []string{"Cannot Read JWT Payload"},
				},	
			},
		}
		break
	case "400-Signature":
		ret = &project.SingleResponse {
			Meta: project.SimpleMeta{
				StatusCode: http.StatusBadRequest,
				Message: project.Message{
					"errors": []string{"Invalid JWT Signature"},
				},	
			},
		}
		break
	case "401-Expired":
		ret = &project.SingleResponse {
			Meta: project.SimpleMeta{
				StatusCode: http.StatusBadRequest,
				Message: project.Message{
					"errors": []string{"Invalid JWT Signature"},
				},	
			},
		}
		break
	case "401":
		ret = &project.SingleResponse {
			Meta: project.SimpleMeta{
				StatusCode: http.StatusUnauthorized,
				Message: project.Message{
					"errors": []string{"JWT Token Is Required"},
				},	
			},
		}
		break
	case "404":
		ret = &project.SingleResponse {
			Meta: project.SimpleMeta{
				StatusCode: http.StatusNotFound,
				Message: project.Message{
					"errors": []string{"Not Found"},
				},	
			},
		}
		break
	case "405":
		ret = &project.SingleResponse {
			Meta: project.SimpleMeta{
				StatusCode: http.StatusMethodNotAllowed,
				Message: project.Message{
					"errors": []string{"Method Not Allowed"},
				},	
			},
		}
		break
	case "500":
		ret = &project.SingleResponse {
			Meta: project.SimpleMeta{
				StatusCode: http.StatusInternalServerError,
				Message: project.Message{
					"errors": []string{err.Error()},
				},	
			},
		}
		break
	default:
		ret = &project.SingleResponse {
			Meta: project.SimpleMeta{
				StatusCode: http.StatusBadRequest,
				Message: project.Message{
					"errors": []string{err.Error()},
				},	
			},
		}
		break
	}

	return ret
}

func ResponseValidation(messages project.Message) *project.SingleResponse {
	response := &project.SingleResponse {
		Meta: project.SimpleMeta{
			StatusCode: http.StatusBadRequest,
			Message: messages,
		},
	}

	return response	
}