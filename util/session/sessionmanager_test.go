package session

import (
	"context"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/argoproj/argo-cd/common"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"

	"github.com/argoproj/argo-cd/errors"
	p "github.com/argoproj/argo-cd/util/password"
	"github.com/argoproj/argo-cd/util/settings"
)

const password = "password"

func getKubeClient() *fake.Clientset {
	const defaultSecretKey = "Hello, world!"

	bcrypt, err := p.HashPassword(password)
	errors.CheckError(err)

	return fake.NewSimpleClientset(&corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "argocd-cm",
			Namespace: "argocd",
			Labels: map[string]string{
				"app.kubernetes.io/part-of": "argocd",
			},
		},
	}, &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "argocd-secret",
			Namespace: "argocd",
		},
		Data: map[string][]byte{
			"admin.password":   []byte(bcrypt),
			"server.secretkey": []byte(defaultSecretKey),
		},
	})
}

func TestSessionManager(t *testing.T) {
	const (
		defaultSubject = "argo"
	)
	settingsMgr := settings.NewSettingsManager(context.Background(), getKubeClient(), "argocd", false)
	mgr := NewSessionManager(settingsMgr, "")

	token, err := mgr.Create(defaultSubject, 0)
	if err != nil {
		t.Errorf("Could not create token: %v", err)
	}

	claims, err := mgr.Parse(token)
	if err != nil {
		t.Errorf("Could not parse token: %v", err)
	}

	mapClaims := *(claims.(*jwt.MapClaims))
	subject := mapClaims["sub"].(string)
	if subject != "argo" {
		t.Errorf("Token claim subject \"%s\" does not match expected subject \"%s\".", subject, defaultSubject)
	}
}

var loggedOutContext = context.Background()
var loggedInContext = context.WithValue(context.Background(), "claims", &jwt.MapClaims{"iss": "qux", "sub": "foo", "email": "bar", "groups": []string{"baz"}})

func TestIss(t *testing.T) {
	assert.Empty(t, Iss(loggedOutContext))
	assert.Equal(t, "qux", Iss(loggedInContext))
}
func TestLoggedIn(t *testing.T) {
	assert.False(t, LoggedIn(loggedOutContext))
	assert.True(t, LoggedIn(loggedInContext))
}

func TestUsername(t *testing.T) {
	assert.Empty(t, Username(loggedOutContext))
	assert.Equal(t, "bar", Username(loggedInContext))
}

func TestSub(t *testing.T) {
	assert.Empty(t, Sub(loggedOutContext))
	assert.Equal(t, "foo", Sub(loggedInContext))
}

func TestGroups(t *testing.T) {
	assert.Empty(t, Groups(loggedOutContext))
	assert.Equal(t, []string{"baz"}, Groups(loggedInContext))
}

func TestVerifyUsernamePassword(t *testing.T) {
	for _, tc := range []struct {
		name         string
		disableAdmin bool
		userName     string
		password     string
		expected     error
	}{
		{
			name:         "Success if userName and password is correct",
			disableAdmin: false,
			userName:     common.ArgoCDAdminUsername,
			password:     password,
			expected:     nil,
		},
		{
			name:         "Return error if userName is not admin",
			disableAdmin: false,
			userName:     "foo",
			password:     password,
			expected:     status.Errorf(codes.Unauthenticated, badUserError),
		},
		{
			name:         "Return error if password is empty",
			disableAdmin: false,
			userName:     common.ArgoCDAdminUsername,
			password:     "",
			expected:     status.Errorf(codes.Unauthenticated, blankPasswordError),
		},
		{
			name:         "Return error if password is not correct",
			disableAdmin: false,
			userName:     common.ArgoCDAdminUsername,
			password:     "foo",
			expected:     status.Errorf(codes.Unauthenticated, invalidLoginError),
		},
		{
			name:         "Return error if disableAdmin is true",
			disableAdmin: true,
			userName:     common.ArgoCDAdminUsername,
			password:     password,
			expected:     status.Errorf(codes.Unauthenticated, adminDisable),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			settingsMgr := settings.NewSettingsManager(context.Background(), getKubeClient(), "argocd", tc.disableAdmin)
			mgr := NewSessionManager(settingsMgr, "")

			err := mgr.VerifyUsernamePassword(tc.userName, tc.password)

			if tc.expected == nil {
				assert.Nil(t, err)
			} else {
				assert.EqualError(t, err, tc.expected.Error())
			}

		})
	}

}