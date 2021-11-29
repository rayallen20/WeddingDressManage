package kind

import "WeddingDressManage/business/v1/dress"

type Response struct {
	Id int `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

// Generate 根据biz层的Kind对象集合 生成响应信息
func (r *Response) Generate(kinds []*dress.Kind) []*Response {
	responses := make([]*Response, 0, len(kinds))

	for _, kind := range kinds {
		response := &Response{
			Id:   kind.Id,
			Code: kind.Code,
			Name: kind.Name,
		}

		responses = append(responses, response)
	}

	return responses
}
