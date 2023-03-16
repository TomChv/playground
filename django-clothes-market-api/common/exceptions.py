from django.http import JsonResponse


class BaseError:
    def __init__(self, code: int, message: str):
        self.status_code = code
        self.message = message

    def res(self) -> JsonResponse:
        return JsonResponse({
            "code": self.status_code,
            "error_message": self.message
        }, status=self.status_code)


class BadRequest(BaseError):
    def __init__(self, code=400, message='request is not correctly formatted.'):
        super().__init__(code, message)
