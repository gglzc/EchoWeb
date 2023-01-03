/*跟前端溝通的地方 */
package request

import (
	
)

type RegisterRequest struct{
<<<<<<< HEAD:request/request.go
	UserID			uint64 `json:"userid"`
	Password  	string `json:"password"`
	Email		string `json:"Email"`
	Birthday  	string `json:"Birthday"`
=======
	UserName	string 
	Password  	string
	Email		string
	Birthday  	string
>>>>>>> 5971af5 (improve whole structure):model/request/request.go
}

type LoginRequest struct{
	ID		string `json:"UserID"`
	IP		string `json:"ip"`
	Token	string `json:"token"`
}

func (r *RegisterRequest) Register() {
	r=&RegisterRequest{
		ID:r.ID,
		Password:r.Password,
	} 
}



