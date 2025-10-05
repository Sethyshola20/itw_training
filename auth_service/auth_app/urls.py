from django.urls import path, include
from rest_framework.routers import DefaultRouter
from .views import (
    UserViewSet,
    SignupView,
    CookieTokenObtainPairView,
    CookieTokenRefreshView,
)

router = DefaultRouter()
router.register(r"users", UserViewSet, basename="user")

urlpatterns = [
    path("api/", include(router.urls)),
    path("api/signup/", SignupView.as_view(), name="signup"),
    path("api/token/", CookieTokenObtainPairView.as_view(), name="token_obtain_pair"),
    path("api/token/refresh/", CookieTokenRefreshView.as_view(), name="token_refresh"),
]
