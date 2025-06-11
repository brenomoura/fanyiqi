import time
from app.services.translation import TranslationService


def main():
    start_time = time.time()
    transalation_service = TranslationService("facebook/m2m100_418M")
    result = transalation_service.translate("Teste", "pt", "en")
    print(result)

    print(f"Time elapsed {time.time() - start_time:.2f} seconds")


if __name__ == "__main__":
    main()
