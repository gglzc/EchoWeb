/*跟前端溝通的地方 */
package request

type RegisterRequest struct{
	ID			string 
	Password  	string
	Email		string
	Birthday  	string
	Post		string
}

type LoginRequest struct{
	ID		string
	IP		string
	Token	string
}

type PostRequest struct{

}

type LikeRequest struct{
	Like 	uint64
	
}

