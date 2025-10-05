from rest_framework import viewsets, generics,status
from rest_framework.permissions import AllowAny, IsAuthenticated
from .models import User
from rest_framework.response import Response
from rest_framework.decorators import action
from .serializers import UserSerializer
from rest_framework_simplejwt.tokens import RefreshToken
from datetime import timedelta
from rest_framework_simplejwt.views import TokenObtainPairView, TokenRefreshView
from rest_framework_simplejwt.authentication import JWTAuthentication
from django.utils import timezone
from .auth import CookieJWTAuthentication

class UserViewSet(viewsets.ModelViewSet):
    queryset = User.objects.all()
    serializer_class = UserSerializer
    authentication_classes = [CookieJWTAuthentication]  
    
    @action(detail=False, methods=["get"], url_path="me", permission_classes=[IsAuthenticated])
    def me(self, request):
        serializer = self.get_serializer(request.user)
        return Response(serializer.data)

    def list(self, request, *args, **kwargs):
        return Response({"detail": "Not allowed."}, status=status.HTTP_405_METHOD_NOT_ALLOWED)

    def retrieve(self, request, *args, **kwargs):
        return Response({"detail": "Not allowed."}, status=status.HTTP_405_METHOD_NOT_ALLOWED)




class SignupView(generics.CreateAPIView):
    serializer_class = UserSerializer
    permission_classes = [AllowAny]
    authentication_classes = []

    def create(self, request, *args, **kwargs):
        data = request.data.copy()

        serializer = self.get_serializer(data=request.data)
        serializer.is_valid(raise_exception=True)
        user = serializer.save()

        user.last_login = timezone.now()
        user.save(update_fields=['last_login'])

        refresh = RefreshToken.for_user(user)
        access_token = str(refresh.access_token)
        refresh_token = str(refresh)

        data = {
            "user": {
                "id": user.id,
                "email": user.email,
                "name": user.name,
            }
        }

        response = Response(data, status=status.HTTP_201_CREATED)

        cookie_max_age = 7 * 24 * 60 * 60  
        response.set_cookie(
            "access_token",
            access_token,
            max_age=cookie_max_age,
            httponly=True,
            secure=False,  
            samesite="Lax",
            path="/"
        )
        response.set_cookie(
            "refresh_token",
            refresh_token,
            max_age=cookie_max_age,
            httponly=True,
            secure=False,
            samesite="Lax",
            path="/"
        )

        return response

class CookieTokenRefreshView(TokenRefreshView):
    def post(self, request, *args, **kwargs):
        if "refresh" not in request.data and "refresh_token" in request.COOKIES:
            request.data["refresh"] = request.COOKIES.get("refresh_token")
        return super().post(request, *args, **kwargs)

    def finalize_response(self, request, response, *args, **kwargs):
        if response.status_code == 200 and "access" in response.data:
            access_token = response.data["access"]
            cookie_max_age = 7 * 24 * 60 * 60
            response.set_cookie(
                "access_token",
                access_token,
                max_age=cookie_max_age,
                httponly=True,
                secure=False,
                samesite="Lax",
                path="/"
            )
            response.data.pop("access", None)
        return super().finalize_response(request, response, *args, **kwargs)



class CookieTokenObtainPairView(TokenObtainPairView):
    def finalize_response(self, request, response, *args, **kwargs):
        if response.status_code == 200:
            data = response.data
            access_token = data.get("access")
            refresh_token = data.get("refresh")

            cookie_max_age = 7 * 24 * 60 * 60  

            response.set_cookie(
                "access_token",
                access_token,
                max_age=cookie_max_age,
                httponly=True,
                secure=False,
                samesite="Lax",
                path="/"
            )
            response.set_cookie(
                "refresh_token",
                refresh_token,
                max_age=cookie_max_age,
                httponly=True,
                secure=False,
                samesite="Lax",
                path="/"
            )

            response.data.pop("access", None)
            response.data.pop("refresh", None)

        return super().finalize_response(request, response, *args, **kwargs)
