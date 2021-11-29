package kind

import "WeddingDressManage/business/v1/dress"

type ShowResponse struct {
	Id int `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

// Generate 根据biz层的Kind对象集合 生成响应信息
func (r *ShowResponse) Generate(kinds []*dress.Kind) []*ShowResponse {
	responses := make([]*ShowResponse, 0, len(kinds))

	for _, kind := range kinds {
		response := &ShowResponse{
			Id:   kind.Id,
			Code: kind.Code,
			Name: kind.Name,
		}

		responses = append(responses, response)
	}

	return responses
}
