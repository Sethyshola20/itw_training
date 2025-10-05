from rest_framework_simplejwt.authentication import JWTAuthentication
from rest_framework_simplejwt.exceptions import InvalidToken, TokenError
from rest_framework.exceptions import AuthenticationFailed
import logging

logger = logging.getLogger(__name__)

class CookieJWTAuthentication(JWTAuthentication):

    def authenticate(self, request):
        header = self.get_header(request)
        if header is not None:
            logger.debug("CookieJWTAuthentication: header present, delegating to parent")
            return super().authenticate(request)

        # 2) Cookie-based auth
        raw_token = request.COOKIES.get("access_token")
        logger.debug("CookieJWTAuthentication: raw_token from cookies: %s", bool(raw_token))
        if not raw_token:
            return None

        try:
            validated_token = self.get_validated_token(raw_token)
            user = self.get_user(validated_token)
        except TokenError as e:
            logger.warning("CookieJWTAuthentication: token invalid/expired: %s", e)
            raise AuthenticationFailed("Invalid or expired token")
        except Exception as e:
            logger.exception("CookieJWTAuthentication: unexpected error")
            raise AuthenticationFailed("Authentication error")

        return (user, validated_token)
