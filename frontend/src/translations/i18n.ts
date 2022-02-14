import i18n from "i18next";
import { initReactI18next } from "react-i18next";
import HttpApi from "i18next-http-backend";

export const i18next = i18n

  .use(initReactI18next)
  .use(HttpApi) // passes i18n down to react-i18next
  .init({
    backend: {
      loadPath: "/locales/{{lng}}/{{ns}}.json",
    },
    defaultNS: "common",
    lng: "de",
    fallbackLng: "de",

    interpolation: {
      escapeValue: false,
    },
  });
