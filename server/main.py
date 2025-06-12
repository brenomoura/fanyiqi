import time

from fastapi import FastAPI
from fastapi.routing import APIRoute

from app.core.config import settings
from app.services.translation import TranslationService
from app.api.routes import routes


def main():
    start_time = time.time()
    transalation_service = TranslationService("michaelfeil/ct2fast-m2m100_418M")
    result = transalation_service.translate(
        "Super teste meu caro, será que isso aqui vai funcionar?", "pt", "en"
    )
    print(result)

    print(f"Time elapsed {time.time() - start_time:.2f} seconds")


def custom_generate_unique_id(route: APIRoute) -> str:
    return f"{route.tags[0]}-{route.name}"


app = FastAPI(
    title=settings.PROJECT_NAME,
    openapi_url=f"{settings.API_V1_STR}/openapi.json",
    generate_unique_id_function=custom_generate_unique_id,
)


app.include_router(routes.api_router, prefix=settings.API_V1_STR)
