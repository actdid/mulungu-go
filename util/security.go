package util

import (
	"net/http"
	"strings"

	"github.com/actdid/mulungu-go/constant"
	"github.com/actdid/mulungu-go/logger"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

//IsAuthorised checks if Request is authorised
func IsAuthorised(ctx context.Context, r *http.Request) bool {
	requestAuthorised := r.Header.Get(constant.HeaderAuthorised)
	logger.Debugf(ctx, "security util", "checking if request has been authorised, authorised %s", requestAuthorised)

	if requestAuthorised != "" && requestAuthorised == "true" {
		logger.Debugf(ctx, "security util", "request contains authorization information, authorised %s", requestAuthorised)
		return true
	}

	logger.Debugf(ctx, "security util", "request not authorised")
	return false
}

//HasAnyRole checks if this request has any of the roles specified in roles
func HasAnyRole(ctx context.Context, roles []string, r *http.Request) bool {
	requestRoles := strings.Split(r.Header.Get(constant.HeaderAuthorisedRole), ",")
	logger.Debugf(ctx, "security util", "request roles %#v ", requestRoles)

	if len(requestRoles) > 0 {
		for _, role := range roles {
			for _, requestRole := range requestRoles {
				if strings.Compare(role, requestRole) == 0 {
					logger.Debugf(ctx, "security util", "request has one required role required:%s found:%s", role, requestRole)
					return true
				}
			}
		}

	}

	return false
}

//EncryptPassword encrypts a password
func EncryptPassword(subject string) ([]byte, error) {

	if IsEncryptedPassword(subject) == false {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(subject), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		return hashedPassword, nil
	}

	//this password is already encrypted, skip encryption
	return []byte(subject), nil
}

//IsEncryptedPassword determins if password is sncrypted by calculating its cost
func IsEncryptedPassword(subject string) bool {
	//is this password encrypted? determine by calculating encryption cost
	encryptionCost, _ := bcrypt.Cost([]byte(subject))
	if encryptionCost == 0 {
		return false
	}

	return true
}
