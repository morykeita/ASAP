/**
 * @author Mory Keita on 1/20/20
 */
package v1

import (
	"encoding/json"
	"github.com/morykeita/ASAP/auth-service/internal/api/utils"
	"github.com/morykeita/ASAP/auth-service/internal/database"
	"github.com/morykeita/ASAP/auth-service/internal/model"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type UserAPI struct {
	DB database.Database
}

type UserParams struct {
	model.User
	Password string `json:"password"`
}

func  (api *UserAPI) Create (w http.ResponseWriter , r *http.Request) {
	logger := log.WithField("func" , "user.go -> created()")

	// Load parameters
	var userParams UserParams
	if err := json.NewDecoder(r.Body).Decode(&UserParams{});err != nil{
		log.WithError(err).Warn("Could not decode request parameters.")
		utils.WriteError(w,http.StatusBadRequest, "Could not decode request parameters", map[string] string{
			"error":err.Error(),
		})
		return
	}
	logger = logger.WithFields(log.Fields{
		"email": userParams.Email,
		//"password" : userParams.Password,
	})

	if err := userParams.Verify(); err != nil{
		log.WithError(err).Warn("Not all field found.")
		utils.WriteError(w,http.StatusInternalServerError,"Not all field found.", map[string] string{
			"error":err.Error(),
		})
		return
	}

	pwHash , err := model.HashPassword(userParams.Password)
	if err != nil{
		log.WithError(err).Warn("Could not hash password.")
		utils.WriteError(w,http.StatusInternalServerError,"Could not hash password.",nil)
		return
	}

	newUser := &model.User{
		Email:        userParams.Email,
		PasswordHash: &pwHash,
	}

	ctx := r.Context()
	if err := api.DB.CreateUser(ctx,newUser); err == database.ErrorUserExists{
		log.WithError(err).Warn("User already exists.")
		utils.WriteError(w,http.StatusConflict,"User already exists",nil)
	} else if err!=nil{
		log.WithError(err).Warn("Error creating user.")
		utils.WriteError(w,http.StatusConflict,"Error creating user.",nil)
	}
	// user has been created

	createdUser,err := api.DB.GetUserByID(ctx,&newUser.ID)
	if err != nil {
		log.WithError(err).Warn("Error creating user.")
		utils.WriteError(w,http.StatusConflict,"Error creating user.",nil)
		return
	}
	log.Info("User created!")
	utils.WriteJSON(w,http.StatusCreated,createdUser)

}