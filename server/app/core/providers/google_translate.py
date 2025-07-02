from urllib.parse import quote_plus

import requests

from app.core.providers.base_model import BaseTranslationModel
from app.core.providers.exceptions import NotFoundLanguage


class GoogleTranslate(BaseTranslationModel):
    def __init__(self):
        self.name = "google-translate"

    def get_model_name(self) -> str:
        return self.name

    def translate(self, text: str, source_lang: str, target_lang: str) -> str:
        keywords = quote_plus(text)
        url_tmpl = (
            "https://translate.googleapis.com/translate_a/"
            "single?client=gtx&sl={}&tl={}&dt=at&dt=bd&dt=ex&"
            "dt=ld&dt=md&dt=qca&dt=rw&dt=rm&dt=ss&dt=t&q={}"
        )
        response = requests.get(
            url_tmpl.format(source_lang, target_lang, keywords), timeout=10
        ).json()
        try:
            result = "".join(x[0] for x in response[0] if x[0] is not None)
        except Exception:
            raise ConnectionError("Failed to translate.")

        return result

    def validate_lang(self, lang: str):
        """
        Validates if the given language code is supported by the model.
        Returns True if supported, False otherwise.
        """
        supported_langs = [code for _, code in self.get_languages()]
        if lang not in supported_langs:
            raise NotFoundLanguage

    def get_languages(self):
        return [
            ("Afrikaans", "af"),
            ("Albanian", "sq"),
            ("Amharic", "am"),
            ("Arabic", "ar"),
            ("Armenian", "hy"),
            ("Assamese", "as"),
            ("Aymara", "ay"),
            ("Azerbaijani", "az"),
            ("Bambara", "bm"),
            ("Basque", "eu"),
            ("Belarusian", "be"),
            ("Bengali", "bn"),
            ("Bhojpuri", "bho"),
            ("Bosnian", "bs"),
            ("Bulgarian", "bg"),
            ("Catalan", "ca"),
            ("Cebuano", "ceb"),
            ("Chinese (Simplified)", "zh-CN"),
            ("Chinese (Traditional)", "zh-TW"),
            ("Croatian", "hr"),
            ("Corsican", "co"),
            ("Czech", "cs"),
            ("Danish", "da"),
            ("Dhivehi", "dv"),
            ("Dogri", "doi"),
            ("Dutch", "nl"),
            ("English", "en"),
            ("Esperanto", "eo"),
            ("Estonian", "et"),
            ("Ewe", "ee"),
            ("Filipino (Tagalog)", "fil"),
            ("Tagalog (Filipino)", "tl"),
            ("Finnish", "fi"),
            ("French", "fr"),
            ("Frisian", "fy"),
            ("Galician", "gl"),
            ("Georgian", "ka"),
            ("German", "de"),
            ("Greek", "el"),
            ("Guarani", "gn"),
            ("Gujarati", "gu"),
            ("Haitian (Creole)", "ht"),
            ("Hausa", "ha"),
            ("Hawaiian", "haw"),
            ("Hebrew", "iw"),
            ("Hindi", "hi"),
            ("Hmong", "hmn"),
            ("Hungarian", "hu"),
            ("Icelandic", "is"),
            ("Igbo", "ig"),
            ("Ilocano", "ilo"),
            ("Indonesian", "id"),
            ("Irish", "ga"),
            ("Italian", "it"),
            ("Japanese", "ja"),
            ("Javanese", "jv"),
            ("Kannada", "kn"),
            ("Kazakh", "kk"),
            ("Khmer", "km"),
            ("Kinyarwanda", "rw"),
            ("Konkani", "gom"),
            ("Korean", "ko"),
            ("Krio", "kri"),
            ("Kurdish", "ku"),
            ("Kurdish (Sorani)", "ckb"),
            ("Kyrgyz", "ky"),
            ("Lao", "lo"),
            ("Latin", "la"),
            ("Latvian", "lv"),
            ("Lingala", "ln"),
            ("Lithuanian", "lt"),
            ("Luganda", "lg"),
            ("Luxembourgish", "lb"),
            ("Macedonian", "mk"),
            ("Maithili", "mai"),
            ("Malagasy", "mg"),
            ("Malay", "ms"),
            ("Malayalam", "ml"),
            ("Maltese", "mt"),
            ("Maori", "mi"),
            ("Marathi", "mr"),
            ("Meiteilon (Manipuri)", "mni-Mtei"),
            ("Mizo", "lus"),
            ("Mongolian", "mn"),
            ("Myanmar (Burmese)", "my"),
            ("Nepali", "ne"),
            ("Norwegian", "no"),
            ("Nyanja (Chichewa)", "ny"),
            ("Odia (Oriya)", "or"),
            ("Oromo", "om"),
            ("Pashto", "ps"),
            ("Persian", "fa"),
            ("Polish", "pl"),
            ("Portuguese", "pt"),
            ("Punjabi", "pa"),
            ("Quechua", "qu"),
            ("Romanian", "ro"),
            ("Russian", "ru"),
            ("Samoan", "sm"),
            ("Sanskrit", "sa"),
            ("Scots Gaelic", "gd"),
            ("Sepedi", "nso"),
            ("Serbian", "sr"),
            ("Sesotho", "st"),
            ("Shona", "sn"),
            ("Sindhi", "sd"),
            ("Slovak", "sk"),
            ("Slovenian", "sl"),
            ("Somali", "so"),
            ("Spanish", "es"),
            ("Sundanese", "su"),
            ("Swahili", "sw"),
            ("Swedish", "sv"),
            ("Tajik", "tg"),
            ("Tamil", "ta"),
            ("Tatar", "tt"),
            ("Telugu", "te"),
            ("Thai", "th"),
            ("Tigrinya", "ti"),
            ("Tsonga", "ts"),
            ("Turkish", "tr"),
            ("Turkmen", "tk"),
            ("Twi (Akan)", "ak"),
            ("Ukrainian", "uk"),
            ("Urdu", "ur"),
            ("Uyghur", "ug"),
            ("Uzbek", "uz"),
            ("Vietnamese", "vi"),
            ("Welsh", "cy"),
            ("Xhosa", "xh"),
            ("Yiddish", "yi"),
            ("Yoruba", "yo"),
            ("Zulu", "zu"),
        ]
