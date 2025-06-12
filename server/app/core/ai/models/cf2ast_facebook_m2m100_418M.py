from huggingface_hub import snapshot_download
from app.core.ai.base_model import BaseTranslationModel
import ctranslate2
import transformers
from app.core.ai.models.exceptions import NotFoundLanguage


class CF2FastM2M100(BaseTranslationModel):
    def __init__(self):
        self.ctranslate2_name = "m2m100_418M"
        self.name = "michaelfeil/ct2fast-m2m100_418M"

    def get_model_name(self) -> str:
        return self.name

    def translate(self, text: str, source_lang: str, target_lang: str) -> str:
        for lang in [source_lang, target_lang]:
            self.validate_lang(lang)

        model_path = snapshot_download(self.name)
        print(model_path)
        translator = ctranslate2.Translator(model_path, compute_type="auto")
        # Use the original HuggingFace model for the tokenizer
        tokenizer = transformers.AutoTokenizer.from_pretrained("facebook/m2m100_418M")
        tokenizer.src_lang = source_lang

        source = tokenizer.convert_ids_to_tokens(tokenizer.encode(text))
        target_prefix = [tokenizer.lang_code_to_token[target_lang]]
        results = translator.translate_batch([source], target_prefix=[target_prefix])
        target = results[0].hypotheses[0][1:]

        return tokenizer.decode(tokenizer.convert_tokens_to_ids(target))

    def validate_lang(self, lang: str):
        """
        Validates if the given language code is supported by the model.
        Returns True if supported, False otherwise.
        """
        supported_langs = [code for _, code in self.get_languages()]
        if lang not in supported_langs:
            raise NotFoundLanguage

    def get_languages(self):
        """
        List of covered languages:
        https://huggingface.co/facebook/m2m100_418M#languages-covered
        """
        return [
            ("Afrikaans", "af"),
            ("Albanian", "sq"),
            ("Amharic", "am"),
            ("Arabic", "ar"),
            ("Armenian", "hy"),
            ("Asturian", "ast"),
            ("Azerbaijani", "az"),
            ("Bashkir", "ba"),
            ("Belarusian", "be"),
            ("Bengali", "bn"),
            ("Bosnian", "bs"),
            ("Breton", "br"),
            ("Bulgarian", "bg"),
            ("Burmese", "my"),
            ("Catalan; Valencian", "ca"),
            ("Cebuano", "ceb"),
            ("Central Khmer", "km"),
            ("Chinese", "zh"),
            ("Croatian", "hr"),
            ("Czech", "cs"),
            ("Danish", "da"),
            ("Dutch; Flemish", "nl"),
            ("English", "en"),
            ("Estonian", "et"),
            ("Finnish", "fi"),
            ("French", "fr"),
            ("Fulah", "ff"),
            ("Galician", "gl"),
            ("Ganda", "lg"),
            ("Gaelic; Scottish Gaelic", "gd"),
            ("Georgian", "ka"),
            ("German", "de"),
            ("Greeek", "el"),
            ("Gujarati", "gu"),
            ("Haitian; Haitian Creole", "ht"),
            ("Hausa", "ha"),
            ("Hebrew", "he"),
            ("Hindi", "hi"),
            ("Hungarian", "hu"),
            ("Icelandic", "is"),
            ("Igbo", "ig"),
            ("Iloko", "ilo"),
            ("Indonesian", "id"),
            ("Irish", "ga"),
            ("Italian", "it"),
            ("Japanese", "ja"),
            ("Javanese", "jv"),
            ("Kannada", "kn"),
            ("Kazakh", "kk"),
            ("Korean", "ko"),
            ("Lao", "lo"),
            ("Latvian", "lv"),
            ("Lingala", "ln"),
            ("Lithuanian", "lt"),
            ("Luxembourgish; Letzeburgesch", "lb"),
            ("Macedonian", "mk"),
            ("Malagasy", "mg"),
            ("Malay", "ms"),
            ("Malayalam", "ml"),
            ("Marathi", "mr"),
            ("Mongolian", "mn"),
            ("Nepali", "ne"),
            ("Northern Sotho", "ns"),
            ("Norwegian", "no"),
            ("Occitan (post 1500)", "oc"),
            ("Oriya", "or"),
            ("Panjabi; Punjabi", "pa"),
            ("Persian", "fa"),
            ("Polish", "pl"),
            ("Portuguese", "pt"),
            ("Pushto; Pashto", "ps"),
            ("Romanian; Moldavian; Moldovan", "ro"),
            ("Russian", "ru"),
            ("Sindhi", "sd"),
            ("Sinhala; Sinhalese", "si"),
            ("Slovak", "sk"),
            ("Slovenian", "sl"),
            ("Somali", "so"),
            ("Spanish", "es"),
            ("Sundanese", "su"),
            ("Swahili", "sw"),
            ("Swati", "ss"),
            ("Swedish", "sv"),
            ("Tagalog", "tl"),
            ("Tamil", "ta"),
            ("Thai", "th"),
            ("Tswana", "tn"),
            ("Turkish", "tr"),
            ("Ukrainian", "uk"),
            ("Urdu", "ur"),
            ("Uzbek", "uz"),
            ("Vietnamese", "vi"),
            ("Welsh", "cy"),
            ("Western Frisian", "fy"),
            ("Wolof", "wo"),
            ("Xhosa", "xh"),
            ("Yiddish", "yi"),
            ("Yoruba", "yo"),
            ("Zulu", "zu"),
        ]
