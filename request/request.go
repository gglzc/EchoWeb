/*跟前端溝通的地方 */
package request

type RegisterRequest struct{
	UserID			uint64 `json:"userid"`
	Password  	string `json:"password"`
	Email		string `json:"Email"`
	Birthday  	string `json:"Birthday"`
}

type LoginRequest struct{
	ID		string `json:"UserID"`
	IP		string `json:"ip"`
	Token	string `json:"token"`
}

type PostRequest struct{

}

type LikeRequest struct{
	Like 	uint64
	
}

